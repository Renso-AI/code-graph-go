// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.0.18"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "eb0f1b4da429476d7d454855dd5b72b6cdf0728d942f7a8332dcdc1f520a9151",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "f4337b08d2d9c5cdbcf6dd2e5e332d8782a84baf671812f1993595d0fc889135",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "77d56cbf64167186607150e63a7878547dc7074ce9cd4b941819758610a4dce7",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "598b9a2be07a57ccfd04dd52ca2f47ac59cb1707d143202a470f8eefefe8f894",
	"code_graph-x86_64-apple-darwin.tar.gz":           "7de474aa299af0bcf06a307d3e04bca5b11d7005028ed4ed2cf6da31f5b43a89",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "f874699189688c8eea091e732d60f0ba011c056ae553f0ee85a275504238f0ce",
	"code_graph-x86_64-pc-windows-msvc.zip":           "7bf8eb4bd2d6edb928f2fd2a4787b11b2952bb25a9f6f477dbca97839277e8c1",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "e318b8fd6278eea0728babcdb8af9d64cffae64d2e92d43d5d7baf4b7db06f3b",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "401069c1f84d1acbd63221cc6817c9cd466536abe518d2f4a829ad0a22901242",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "36388d562ca88385fb5e0b0af09a08c66028e15ed9f555d7991c10d561a4be71",
}
