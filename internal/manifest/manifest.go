// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.0.19"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "93c4eefb933bbb8f851f3cfe7d8b1ebd3ad6f05c0e1e42f65b4f6173f2c9cd63",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "02dbcaefe674011147e3ef9f28b3e0485722a3fa39e237da78e6377ffb8df73f",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "da5de8974b8bc01cb7007e3015a11cdf1a3a452e4877cfd572100dc141baf9d6",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "680d30bbfe834a317feb7de1efaf5b1cb4dbf77a22b90ef41296b3ad15842e3c",
	"code_graph-x86_64-apple-darwin.tar.gz":           "4520ff4da46501c58048dea4d0feae67785bd2721b33dae5fd7f9c46f529d87e",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "3042bde451547269595ceb9bfcd72931cab793c3f09279919bd5946e176c3e28",
	"code_graph-x86_64-pc-windows-msvc.zip":           "818e152bb4eb000fdaf42cfa44767dbacabc908fefd699663659354ea1068d06",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "842ad09e68525255794778b036aa81fcbf86f5b20c76937ecfe14c608d869ffc",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "54c153c5659fa7dc468b2ed7cf41bf1942c98e9335253978821beca538775fbd",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "a8a6e7854d61e058d0edf60b7f790384107d0db5aabdef14216a3091d64ad485",
}
