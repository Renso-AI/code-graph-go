// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.6.1"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "15787bd1476d1a4a53238fb8c66dda32f36582adb4c2dec349eeb38a7a371691",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "ad50f8aba9b70ac8968dc4c3845825371bebd0d5998ef9683eb2811db56e7849",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "59230339fc826de9c0cce995ad768dd66cd9f51e814aaea4d670c9bfc0d172d9",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "9e5e875a847aeab6694d07839375ec243c01274690939ba4c050b6b44a673325",
	"code_graph-x86_64-apple-darwin.tar.gz":           "c7e46d5d6a23847d2d205dc52915cb52bc9de2505c5c7c1fe61b348207321e05",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "5b80178a24441d0712d00ef8feb4bad068a715eaedaf784931263211862ad336",
	"code_graph-x86_64-pc-windows-msvc.zip":           "59330af029f659631c6e65784310e7927ab0328dacb434c5e266cedba89dc497",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "f4e05cc1418c17ce95e10485fd17d0f59be8c420e972266d72d5d76565a3cfd0",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "695c96be57476e53aceaae82b3e0e4c1d45fd535f424780f68b1d33d2c1eb447",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "de6f075b7a898afa9f7abcb61d34a424001408c81c372f0336cd73258c7d0ac6",
}
