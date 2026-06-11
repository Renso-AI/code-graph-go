// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.1.12"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "346865df46025f93139fe622eb968e27b7d0f561320ddece67cffc5a54d65b90",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "dbf2afdcfe594f7a6ee17ea457fe5093d09283b7944aa612bca9a773bdf06f8c",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "8e30b3652fe461f76625465549457628f32fed97b64eaf1223d0a93b14767796",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "c24ca28afeaa28aa3cd27438846a87cde56eeae423e399e69a380cdc87cb93b0",
	"code_graph-x86_64-apple-darwin.tar.gz":           "b0174d47e998165d48b75dd7e7339dc53703839984a212e29e693e9ae28e58fa",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "f46599d74ea4c43d9b0616a154f375a7b5e6f3bc8748fc9a09ef690c15a48c57",
	"code_graph-x86_64-pc-windows-msvc.zip":           "770082d54dad1b996d8254d4a66477db9b851516041eb7b0b5d96a3d04d39bdd",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "bc916218e73e0025ec3f017bd0f9b71cbbdc708926166b90456e0863ea0f411d",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "cefc78fc6574eaf5680c276d240c2cc020d37fed50486055a8540dee8a65b895",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "e1778874e4a604bd50c4a70b8997dce57f90b40d17848060e1fa7a7b605ef3e2",
}
