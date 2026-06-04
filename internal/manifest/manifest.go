// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.0.25"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "196c4a1966337d1c45508eba42da8068e2bf25a000379059ff52b91403038d25",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "05e4e37a5dbfbb7e2b9e933b0808bcee84ff5b001d230165a184032dfe320624",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "1c71efc2a7f336a5c47981d502ae71a0f022f3eb7c58db48068e33dd3332a78f",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "c031c6587c9a7dcc0e28a1f4972f345c2c6e4811c4dc1910bcf7d551a07b9104",
	"code_graph-x86_64-apple-darwin.tar.gz":           "9b56f6e18e6872cb854b396e05c4d357568d955a6819bbc64bc1fb323727d4a0",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "9c55c3d9867d29228e97ba982ff740691aeaca4fd61339a725b7a1385c95c782",
	"code_graph-x86_64-pc-windows-msvc.zip":           "ee4de870035b255cc61179b977852306768394cdd1cad9eb80e1065d81c14d68",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "c64345f0857e0e46c40b77728746cb27a5c387b00f67dcb05ac4848493f173ef",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "3668a4d704896f2c5b81a5bdfc48581b179806e7b56f8839584e3b6447db551b",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "1ad7743bb8982dc8d2928e1d25d4e621599de607f4a6b339ef1d7e99544abcf5",
}
