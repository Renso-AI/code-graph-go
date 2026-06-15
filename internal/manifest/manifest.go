// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.3.0"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "a1477358c62576b2fef131136743c0fa8f76b6b27b7eb792c2ce42e13b20959d",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "14c6bb8ca39e5a586cdea5a266fb7c0d033520516743bac5d0e21eb9b36882aa",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "fb17f62634d84e8bb7bc58a30c0711aa61b75970b634ef370dbbe73c1b128c95",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "85c695446a860f3a33514c3ee3dfcd6ed1ddd70697f530b46fd5e0b12c46a011",
	"code_graph-x86_64-apple-darwin.tar.gz":           "821e86257310707887a7dcc2d473fde3a925f9f5b223498eda760d127ca45cec",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "84b49d8f677abd7df5bd6901a188c647bb4995c06c132d50310d38b756b1e782",
	"code_graph-x86_64-pc-windows-msvc.zip":           "57fea9f0a9201231a13484b81b60c210fb71a70857c423b7cd963c39dc577ffb",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "f5f6d50cbf0d4ea615bab7f70e2a29d2187f1152ced054019a0aa8da5fdd5cff",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "5843da175d4f31dc289543acb4d4fb42d507a4d6ba8f89057b6f5c253a5ad49c",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "33f98fb5241e071264dd9d25ebe3d1469848263cd17e3151d059fb3eccb42a6c",
}
