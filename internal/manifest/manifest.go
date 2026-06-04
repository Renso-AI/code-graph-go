// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.0.24"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "3c3ee9ea99705a6ffd8f090223f060677567d2b734eef3159b3ee47363120a2e",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "c4ede7fd4f62f73676db087675199e92b287e2b0f429731a81b9ad7ae48b8371",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "ee384d0aef628ff3345c51e1e2bfee41a5995d6a713d47ebe5a5c46818ea9fb8",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "b26637dcbf6ce471446bd82b7b56f7556306a4fcaabef6e7767c6648cf2bd8fd",
	"code_graph-x86_64-apple-darwin.tar.gz":           "bf3d0954145671918996eb5e8421d445ebaf693b7fe356a65abdd81bfce69886",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "8e38bed24207fd296bdb44e2e8ae1552ae178cddbe005673cfbe0cded1b0ed3c",
	"code_graph-x86_64-pc-windows-msvc.zip":           "a0478a630e8770673f23550cc126806d78ee678e70aafe243d97fd0cd1161c08",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "34cbd23403a41d2416ead26e35fb7a695531db0c8bea64e39e3fe2cb9547627d",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "5461e5ab17fc708efc7b578e9124c079abeb1fb65a962dc7a5523e555dfadbe7",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "96ce4b651b8ee05ab790cbafa1f0c57da6b7a7235f9625917430b1eacb470504",
}
