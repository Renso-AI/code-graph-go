// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.0.2"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "b8aad9afd84df33726a320ef2016cd7054b60692d6e5422e4464a38366250596",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "e8f238f55baeaa2d1a8bee7fa60feb2b0495794d48574c19ab088940a8aafb83",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "026f4c5cce884fd30852ebe1da48ef657264c32f49415db4bf60892fc9c03ef7",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "c4f5f85d61926ccc5b3eb256fe48285513306b8a31939208166a71266e0fc727",
	"code_graph-x86_64-apple-darwin.tar.gz":           "5376ff51279cf1d1e07e20877301471d623ed8c3db2abe31b17705e08c699d67",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "83b39b0d2cd81b73656a0c1f60bc988707c460f64bbeeeec82537aa52ccd4e25",
	"code_graph-x86_64-pc-windows-msvc.zip":           "cc9996dba9a7025fd1d0a0cd17ef9d93247bb6d9a2b603bf07e7aa0056a47ec5",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "a779be40892720cdb8df38ecdd91188662cce814377c1b22de0b53c46bb213d7",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "91d418ed8c13f64a66c7322fb28e1482ab6d126bea1ea6645e88ed44e71582b9",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "90931fbb33487b000173e572f78379bb5647932642a3b94e264eaa713e88cbb6",
}
