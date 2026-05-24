// Package launcher fetches and execs the prebuilt code_graph binaries.
//
// On first run for a given (binary, version, target) tuple we:
//   1. Build the URL: <RELEASE_URL>/v<VERSION>/<binary>-<target>.<ext>.
//   2. Download the archive.
//   3. SHA256-verify it against the manifest baked into manifest.go.
//   4. Extract the binary into XDG_CACHE_HOME/code-graph/<version>/.
//   5. chmod +x (Unix), then exec it forwarding os.Args + env.
//
// Subsequent runs hit the cached binary directly.

package launcher

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"syscall"

	"github.com/renso-ai/code-graph-go/internal/manifest"
)

// Run is the entry point for cmd/<binary>/main.go.
func Run(name string) {
	if err := run(name); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s: %v\n", name, err)
		os.Exit(126)
	}
}

func run(name string) error {
	target, ext, err := hostTarget()
	if err != nil {
		return err
	}

	bin, err := ensureCached(name, target, ext)
	if err != nil {
		return err
	}

	return execv(bin, os.Args[1:])
}

func hostTarget() (target, ext string, err error) {
	switch runtime.GOOS {
	case "linux":
		switch runtime.GOARCH {
		case "amd64":
			return "x86_64-unknown-linux-gnu", "tar.gz", nil
		case "arm64":
			return "aarch64-unknown-linux-gnu", "tar.gz", nil
		}
	case "darwin":
		switch runtime.GOARCH {
		case "amd64":
			return "x86_64-apple-darwin", "tar.gz", nil
		case "arm64":
			return "aarch64-apple-darwin", "tar.gz", nil
		}
	case "windows":
		if runtime.GOARCH == "amd64" {
			return "x86_64-pc-windows-msvc", "zip", nil
		}
	}
	return "", "", fmt.Errorf("unsupported platform: %s/%s", runtime.GOOS, runtime.GOARCH)
}

func cacheDir(version string) (string, error) {
	if runtime.GOOS == "windows" {
		if p := os.Getenv("LOCALAPPDATA"); p != "" {
			return filepath.Join(p, "code-graph", version), nil
		}
	}
	if p := os.Getenv("XDG_CACHE_HOME"); p != "" {
		return filepath.Join(p, "code-graph", version), nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".cache", "code-graph", version), nil
}

func ensureCached(name, target, ext string) (string, error) {
	dir, err := cacheDir(manifest.Version)
	if err != nil {
		return "", err
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", err
	}
	exe := name
	if runtime.GOOS == "windows" {
		exe = name + ".exe"
	}
	dst := filepath.Join(dir, exe)
	if _, err := os.Stat(dst); err == nil {
		return dst, nil
	}

	url := fmt.Sprintf("%s/v%s/%s-%s.%s",
		manifest.ReleaseURLBase, manifest.Version, name, target, ext)
	wantSHA, ok := manifest.SHA256SUMS[fmt.Sprintf("%s-%s.%s", name, target, ext)]
	if !ok {
		return "", fmt.Errorf("manifest has no SHA for %s-%s.%s", name, target, ext)
	}

	body, err := download(url)
	if err != nil {
		return "", fmt.Errorf("download %s: %w", url, err)
	}

	gotSHA := hashHex(body)
	if gotSHA != wantSHA {
		return "", fmt.Errorf("SHA256 mismatch for %s: want %s, got %s", url, wantSHA, gotSHA)
	}

	bin, err := extract(body, ext, exe)
	if err != nil {
		return "", fmt.Errorf("extract %s: %w", url, err)
	}

	tmp := dst + ".partial"
	if err := os.WriteFile(tmp, bin, 0o755); err != nil {
		return "", err
	}
	if err := os.Rename(tmp, dst); err != nil {
		return "", err
	}
	return dst, nil
}

func download(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode/100 != 2 {
		return nil, fmt.Errorf("HTTP %d", resp.StatusCode)
	}
	return io.ReadAll(resp.Body)
}

func hashHex(b []byte) string {
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:])
}

func extract(body []byte, ext, exe string) ([]byte, error) {
	switch ext {
	case "tar.gz":
		return extractTarGz(body, exe)
	case "zip":
		return extractZip(body, exe)
	}
	return nil, fmt.Errorf("unknown ext: %s", ext)
}

func extractTarGz(body []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytesReader(body))
	if err != nil {
		return nil, err
	}
	defer gz.Close()
	tr := tar.NewReader(gz)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if filepath.Base(hdr.Name) == name {
			return io.ReadAll(tr)
		}
	}
	return nil, fmt.Errorf("%s not found in archive", name)
}

func extractZip(body []byte, name string) ([]byte, error) {
	zr, err := zip.NewReader(bytesReaderAt(body), int64(len(body)))
	if err != nil {
		return nil, err
	}
	for _, f := range zr.File {
		if filepath.Base(f.Name) == name {
			rc, err := f.Open()
			if err != nil {
				return nil, err
			}
			defer rc.Close()
			return io.ReadAll(rc)
		}
	}
	return nil, fmt.Errorf("%s not found in archive", name)
}

func execv(bin string, args []string) error {
	if runtime.GOOS == "windows" {
		// No execv on Windows; spawn + wait + propagate.
		return spawnWait(bin, args)
	}
	argv := append([]string{bin}, args...)
	return syscall.Exec(bin, argv, os.Environ())
}
