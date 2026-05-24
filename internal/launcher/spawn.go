package launcher

import (
	"bytes"
	"io"
	"os"
	"os/exec"
)

// bytesReader returns an io.Reader over `b` without an extra copy.
func bytesReader(b []byte) io.Reader { return bytes.NewReader(b) }

// bytesReaderAt returns an io.ReaderAt over `b` (needed by zip.NewReader).
func bytesReaderAt(b []byte) io.ReaderAt { return bytes.NewReader(b) }

// spawnWait runs `bin args...` to completion, forwarding stdio and
// the parent's environment. Windows path (no syscall.Exec available).
func spawnWait(bin string, args []string) error {
	cmd := exec.Command(bin, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()
	err := cmd.Run()
	if exitErr, ok := err.(*exec.ExitError); ok {
		os.Exit(exitErr.ExitCode())
	}
	if err != nil {
		return err
	}
	os.Exit(0)
	return nil
}
