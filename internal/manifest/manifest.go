// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.0.9"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "b8f07055af0a17abde8fe7d9f60be5cc1a91ab55667a96af6ddb35f877cd0174",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "64ed1ae400bf1343447def58d9d6eeea411d0c516bd047856b442d6cde96a58c",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "cf6eeef1146a1cc65e56722e8172f961497e03316668095fbac4f7aaf2f692d3",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "c8835e7e486f9ed112ba87a250994fb7c4cdb14a48015c810debd63b4e1a222e",
	"code_graph-x86_64-apple-darwin.tar.gz":           "34a9b1d4b0680cef1d9c17a75491b53c78b67a8ed6f88de379f37ea234c82b6f",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "3232450294ddca4f9d436225e1a20cb8fced9ce54a7169915c08e7a6478cccf1",
	"code_graph-x86_64-pc-windows-msvc.zip":           "01acac4abce60fb1bf41cdbe5bc54318b64ee2641798e0d6fcbd0426a7c0642d",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "dc3f0d6b8ed119aa8b08cb6fc1996ee6e4bc8d0b5e715b1bc292544b39867125",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "b9441b1d990e1c993f413f55067e70a7a43971002ec16c75b7d33e92e8624180",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "c3ca422c55c16bec13937390ba07886c5ea318caea932d2b6094a6e161f7465c",
}
