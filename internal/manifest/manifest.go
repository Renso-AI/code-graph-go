// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.6.0"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "4995c5f50e17b1ae6f2cd207d2a7a3192d10f4fce56238ec9928f21fc7c4b2ce",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "3175cd4a05fabb1a00af7cb3136f97361dbf5bb7823897fbe4d56920a31a14c2",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "56c80584daf1ae5a72fd06d79708bd30ea96c06323c8983ce214a322247a1f2b",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "e40ee04a64bbfa28d195dbf8def54d166def80444c7f708b1133c1ccc18683f3",
	"code_graph-x86_64-apple-darwin.tar.gz":           "a1e824df9c39d75181ac4c9e8146e11e47500ad7638d2f50f3c247baf9c32520",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "426e32774225d349585857c64dce62a6860333dfa234737ddedfdc4c9740e0cc",
	"code_graph-x86_64-pc-windows-msvc.zip":           "6a2eb74f886186d6f411c4888cba671047c6a8c1bb9d5c6790d21cf359b96d6a",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "14c0dd7af920296bde88abae05540e3d4d3d579e07b96237813057b7116b25d5",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "336374ce2ac158f5e383a6f9d8a710cd2ba7367448586f3ec0d83cae6ba54512",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "d4e89b298a70c5a1d18d717b476edbcc8acd5d925addd9a1e4e1f659a73d7b61",
}
