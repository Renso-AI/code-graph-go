// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.0.12"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "578ca89ba23c730dc40452d155a2420abc0374577c9fb93b316ecb14dba3ece8",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "0833b793b8ffb2981ea5a2297854628197be074698f4f4d1e38ae3e0231345ce",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "2d3f16267f6920205c4b2c0b221580c16db821d61783f79dff7d223c09e18769",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "ccd85622e24085db4c2ef3dbe655c51bc55cd672994729536d5a943ab7d62bb8",
	"code_graph-x86_64-apple-darwin.tar.gz":           "ba41594b7f92c258663faf7fbe6755783c478a03c3fc7d289ce1d6ba13170c5c",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "6fa1b3df1117bc342a474225334b439408b9a003dc2dc80e40fbf18b5ef14f57",
	"code_graph-x86_64-pc-windows-msvc.zip":           "8c554b9e53bf78c8e7dfb41c7687a0864f90f62a709ca3ee49888414f5510046",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "05269ce58a93a5e942cd672406e1a58c537b9429a6e0078251df8f82ddfd3841",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "fd37c549ad7c3e3cea768dbf5095c2346574f9ac8fcca9d9855977a21789fa80",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "ea1e8ed3eace83ac0f9a166da0e556d4f5fc480c2b3f3c709966ba529fac077e",
}
