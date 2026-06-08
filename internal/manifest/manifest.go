// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.1.1"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "c40ed6a1c7f125f55c71b3f6688bd2056ff7a44b089b777cf02af7e67e5c148f",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "7d90463a4026892156e0cc22516cb4a90d34f416e6613e94be7baaab9d356df4",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "c2ba5f26309248d4255743d2fc9f21f9b6c2a6ee9d0380e88f71ea848e4ac7be",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "9ff22849082ecac8c663ea63d2679650fae0391e441a1c47c7b0c21cbeba71a7",
	"code_graph-x86_64-apple-darwin.tar.gz":           "c56306b52999af1534f3813640e346b3b904a53081556e226b1eeebee3945d1c",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "4816f22b931f6de59c16f2055773024fa756fcc7aa65cd5a173da599b405ccde",
	"code_graph-x86_64-pc-windows-msvc.zip":           "7226e64c62b0ba9a834cb7e18d7fb95e6e7efff90638f194ab3f2501dd579cb2",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "dc89f6f830f741d6f8c5f6b81aaff87ebbd12e45900f09ae6331b72744ad7725",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "ee12c92a4504ebfe4d57b52f885014e11d10f0f8bc6914f25b66dce66e211f95",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "87a7e67d40b63824e3fa972bf905ebc01ba743f4c88513c23d44945543bbd834",
}
