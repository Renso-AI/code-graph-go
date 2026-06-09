// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.1.7"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "f52414021ee272f3cbf3bba5d78104503594bc685e66ae1e925129ced94f2064",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "521f84cad1d5f39fd46c5620fcb729711a57a82048400ebf75698880d187ce28",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "e3612780b2eb6fb8999aad19026ccdee277644e5ae99d0df18a77dada3a41603",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "64c58e270bf6307c862a84b733ef26a94a7bead44c7f6bc95d5d218885980667",
	"code_graph-x86_64-apple-darwin.tar.gz":           "262c93ed169159ce138cfca8e19147b5fe97474c98c9d441feff3b69b8c59572",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "2975586f659fe1a281c619630c911fb42e19e801f6b32b844538091c65c58048",
	"code_graph-x86_64-pc-windows-msvc.zip":           "8ea4b4bea03b720936f4bb20d60d0d713263eef6035c539175587abd1f70771f",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "c7178d2e4f0c5cdfe5c5d81e24af6f92cd0234522a7189e74b600fe6933e39b2",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "dfdd17b5567ea6a2e92690351e22eb7195dcefc86dcb7601173c2272ecb6cfea",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "a1ad84b16f53f1227bceab322b37cb4467c6a09edceb78dce9c4f662e5d906fb",
}
