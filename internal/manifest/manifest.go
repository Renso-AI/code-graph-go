// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.1.6"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "4fac8832403536451d3db9149433b7c430addfef831cfe0d29ee30691f976090",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "cffebb494619b373ee081a71143aa2e138bb01957fa9449a8c795602d95d9763",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "aaf2fa377e20669570b7d8528be589a12e97f5a6b26306a8a0f806171159394a",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "3e59d57f3d41f9dafe24162ded7a1cc00cb0d78316d16489b14b61b79e9e0854",
	"code_graph-x86_64-apple-darwin.tar.gz":           "3e2f66398df57b3ff3d97e046581c7a4aa8d8031d881d004d33e6a31eaa573f8",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "6fadaf216a06c4233da606ea3b155786095ad515298b071e6bf46ab13419b566",
	"code_graph-x86_64-pc-windows-msvc.zip":           "d3d39db463a40c92f74513626b6455450d8c09641c0011a63cc387c8e5ea99d3",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "25b2d79710d3338c17fb0766807f96a98ff3385a714cdc23f3f3ed755578db4e",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "306fd6bc86bf8c10476aa450b45ce723f0ae75adbc991abf0de996bb005880f5",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "8402bd12880f3195f05235be07f624ddc06ca9dba188cc0d8d8f402f1236e19a",
}
