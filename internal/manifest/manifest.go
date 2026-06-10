// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.1.11"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "b85bf797a1dce342b0fe999c0ab3e42d9b6250c4fd7cbd850632a3b534af871a",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "54906140d4757736b0844e260764c396b2bd7b8913be502ea044d705603900cd",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "a7ae87433e3fc3efd8dd0680403960796e248a3ee17f8c7c404cd7014159abce",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "c36cce884528955f76c5d9a87cc4bb176700016aa8cc2146b14f77c40c39fea0",
	"code_graph-x86_64-apple-darwin.tar.gz":           "b09c6e9341b417ff8cded2c65624fd0ac62e965686f2369d1e95aee2710e0bbc",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "13c3d3cc2af14c2d46e22b4c07751989ddb59627d46993736f99840a95d1e384",
	"code_graph-x86_64-pc-windows-msvc.zip":           "52306cec24894ee8ba92152f42a299353ed787f313016534e368de2f005619e8",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "b338f566255df3c5668f70b88219294f5485ba20e7c22e62f8e933480e31845e",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "836f6ae4a79d405a261dc4b69ddf3e66b156721b347dfd1b3bea00c5b7f76b3a",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "58a342be73f19d5d8caff49fff10ed94598c054e073b368b34deeba29a2f6590",
}
