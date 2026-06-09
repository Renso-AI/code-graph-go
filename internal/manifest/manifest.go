// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.1.5"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "bc38377ab6190742c28a9cecc794a0701a43da99813a8ac595905c79ce1d24a9",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "4515abbcae37846ad1a43cbe3dd5a9c59d3ee52d3f515f46993671332472290c",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "f645336dddbbdb778674bb4dcd2a47c7f98391c3ce1a8bd198404b170671ac8f",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "057a0243695bc9858e611a4ed831c31131dd0cfb5cf47652d6fe024cd558676b",
	"code_graph-x86_64-apple-darwin.tar.gz":           "0ce9d143abbe4894e70a40b5edb9c397415b2dc20ece4c0d5cef50870e5a9a42",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "921ee24932551b416ba917bf64daac471a7dd91a4e466a6480c924a77f03fc4e",
	"code_graph-x86_64-pc-windows-msvc.zip":           "18e2deea3fb2c54e90a99ba52f60150f1232b519ef0a37acc903f33ef0ca5062",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "49df3e916712ca1a4c3d88cefdb39421302c8188c4198105ccf7c35dedab38ef",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "115c7630a3dd4aec6e0643e87b787b8f59d705448edbdfc3b97d6529a5b84323",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "6bf8d0387b1e1a7c5f1df6daa00553606b0f97430b183e953479c32111d1a415",
}
