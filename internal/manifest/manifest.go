// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.0.27"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "9962a4cf772f7c327157bc1aaaa80eef1b8c7c6bf5f8a6a5ec5f460511b05b60",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "49bd256871b39961282dc657601ae5905d22a8a873d0648bcb555c346f62bd68",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "9aeb251168539637f788730a4a9159546cabd37d03831fe0ab51cd4c95166c92",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "dffa493f04ee7c26023ede7ba225dfc4e0e03cfe4c45121511b25c315942aaa4",
	"code_graph-x86_64-apple-darwin.tar.gz":           "fb1da31e0faddcc6ae1c59be4303b6b8f7071efec92f24b4b4d20150da37051d",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "9a1ea7cd7e97e7027dab2b2c16cc99308948ceb9aca748c1d232816d687c9ba0",
	"code_graph-x86_64-pc-windows-msvc.zip":           "c3fbaa415ca5bef84b709d6b2649557e7ed698c40143ffd4c01b8332d468cdf2",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "10700d13b80c8ee9a5cffe7df90fa8394ea1885e0d5fa11aea455f577ecacac6",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "cb328c9380fdc3453f883cf475a3def0d1ceaec5b8c8e98e83e2b52fe1cbef7b",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "870ff1f74bbb0aa682c98dd9f4870f52acc543d9879d8c4b27fc43635dfc62c3",
}
