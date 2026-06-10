// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.1.10"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "a65673f1df5a2432dfb0c8aeb588777f2b331880d88c245acef9f39d755ed1a7",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "4b569c5870dd15332433fa504898f47a3233a102dd7e9826ea2080fce54b1e93",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "64ed56ae2ac2a7c2795f4f746ac22185864991668d2e6833fc206174ea13338b",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "ecc23be9a24590376a1abb7c2b7178d2d154db46d40f8c06396581d0465385ea",
	"code_graph-x86_64-apple-darwin.tar.gz":           "0bc97c0f5b9330c445967e731ac7b3de555df56c0bae3d417c25d02e7fc74bc4",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "d651cfc1bf240e0cf015439736968ef997a2bfc405ffc66fa67cd10c2f1da4e0",
	"code_graph-x86_64-pc-windows-msvc.zip":           "deffaed73f6bf3048890d66cdd780f6ebf08486aec3cf5950cb127da44d1f6bd",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "abd95c2c00b50fea093220d9f21fe74bd3a4ebe0e874f6a75ee670cff1007fb0",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "6d987ea43b088faab6e161018962c584f5994e266b8617d15ee4f1416db49571",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "8fd10f655109a22b6333295c4cceee0b07a5bd06ecd92ca1e0b8a6f9f6402ffa",
}
