// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.1.8"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "cce2f5e2f283c917bb9e66e669ce82b780f4ecc7531d796fb71d4271761cde58",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "31df62d28447280f51dd3312a24434272636316b3eb0f7ea9330b60cfaa4fec4",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "fcd8deb26c57adbabfdcd06d765ebcd2c6dc0119004f15f9c6af655a5a66d82b",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "f373d1f6ed74021c99c257b6b680b826f512b9b0208c2a5b7a525a4167f87080",
	"code_graph-x86_64-apple-darwin.tar.gz":           "87b3e703cdba4242140f088c0b4f43d2dee8a3b6392702dc9c49d3fefb9b32ff",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "7e692089bee09a6434e50b56f60e0938cfb8a53dbec315ed0a783ef2203ad02d",
	"code_graph-x86_64-pc-windows-msvc.zip":           "43dd89905c7c9ddde0706c6001deda006dde9d55990f3af556b15226b240b337",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "925530d33b6033b3321628643e91bd399a885d4660ee29b4462d9b5499a968bd",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "c543043615f6c75c0371d4114111ab0021f44e6d326eec8549574397cab687b6",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "8906da47ff4ce2dc77229851fbcf7544d9269203df904f61fe8493f67db5817a",
}
