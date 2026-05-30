// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.0.17"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "69e7bb67a166f68467c462b70f0786402aca87953b844a12e7c1828eb576c4e6",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "51eb92c156a789e45ca092c648d238c8fe171239f4028beab11cc39b6198e335",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "c7231d2a66d174fa197653673bf593b528b07d853b107657c9291c364971aaa0",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "e09fa96008d4983e3d5d9d115d7bd3acb47f3223b886cde125c3890d41128af8",
	"code_graph-x86_64-apple-darwin.tar.gz":           "ad125a5c254c31a967bd5a1e4e0cd70293671e251ddcb8107217e43897b5bf3c",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "cb7b496529d8fdfadb33a6e3429a40bf3107effb2ab1798b1ef829d9bc877896",
	"code_graph-x86_64-pc-windows-msvc.zip":           "e71bd5afc59f91727104b9687cd8377ee6b8d4aaa7588716de63b71f338f83c6",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "07cf021e2f158bf690af732dc2d35fac488c436137c3ee545821449e1f420197",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "f5964e4138bd16aa1bd054f292c649b5419abdad69d14e68631bf4ec54ce3fd0",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "9ad4f033ee4bcd7fde24cb59cb64564cdfef5abd608967095ea676c5eea5cb8b",
}
