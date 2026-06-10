// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.1.9"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "56efcca2b081b4ed7be267e4f8d8962e5ec3fc3dd92bba9ce040486723741bb4",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "8a31f92333e5c53ad3b8b9d713ee1ac020bc66b8a4ec8cd659224836783f5139",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "03e38ee66d4a462ec8354d0d2c9678b999aa9edc1b40ab753f01fb6044553edb",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "7c5af8983478618c2dc2be07371ac378e80d8456482a570c348a8ec1c4d2b33c",
	"code_graph-x86_64-apple-darwin.tar.gz":           "1d23db45aebc3749ad9f2c1424eea52d295ebfd278f393c499f7c0660b13c1c9",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "30e22cec89229ae87875a29f7143737c85331e95be647033c99bcec8154b4007",
	"code_graph-x86_64-pc-windows-msvc.zip":           "25b2fcc222c94918772772be4a8f1da49e49262f1454611bce6370ea29c2dda5",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "75a2f487bda791ddc93a0185e4ef6efad52086366aa622752dc5290290698a4c",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "7d0a2b76b0e0abfa59dd3ed900bfdac5b3541d2bfb875c873e5be3c03402556f",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "1a3dae7fc32422354f44061aaeb83b6375b5908bc2a45e82556b73d7ace42221",
}
