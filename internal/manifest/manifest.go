// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.0.3"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "7f20623febffcba7f13e4d460b2284ae82b1fdc1d4903b68ac04421e2667aebb",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "127889941a7de3f7fbf70df2743e1f1846f627d072aa8a46a7663e33c5aeade2",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "458fc8f7cb74d4e6fe95aa2c8879b0be6e8ed775c624664ebcb983405ee0f86e",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "82d15ef04cb46f66e37100459c6fb455d03159d62952cd65bc84ce8a11cc22cc",
	"code_graph-x86_64-apple-darwin.tar.gz":           "6c8ff29ef1a1c083fbf85db3e3dd38c3a7266031f0722de9bed403b8b5361b23",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "ae938adb2cb014de8e94bf4909f6e093471c957df8da1089179706d19bb3ef89",
	"code_graph-x86_64-pc-windows-msvc.zip":           "d9c8198a3109eacab6d706fb079f22eff2b22dbd1e2cd4a9ad7089c025689002",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "543376e889bcce84449ebbf7ed24cd0d6be316fc6aa5049b8bfa5b13c69a93d3",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "33b9bef0c7e297446e480199223b1a4d584a08af1a026a2ac784642c97d54980",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "66d3a321b1a27251f22f46e87060c3e3febcadfa468da1b9bc715b633c4999a1",
}
