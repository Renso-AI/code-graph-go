// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.5.0"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "718c1a597a3750000eae4938a101cf7d2561acf69e3f55a23996e5ccff8c1b8d",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "eb1cbfdccb872758a781528169e20348d402be486e94375fe612da7925086fd4",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "eaa9f80abd389029a87eb1a1c33509cd1cb4ab66bced52f6f4c276603abc485f",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "7c77fa96664a6f2234cbcac489e0b89903bff36626bf813c312354f7bd98c682",
	"code_graph-x86_64-apple-darwin.tar.gz":           "5f3b589269a46d7d7e0440ce5015b281cd50ed314f409d743204df68d7b53b1a",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "910d6a9eb640b75a48bd09373c855c574628cf4e97241deed5efcad8666976f0",
	"code_graph-x86_64-pc-windows-msvc.zip":           "70d8a9e2c51f66b18e52df4e83eeff8909c7e759080960300fc81a42882d5178",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "59a08339ebcdad94920488c79c4ee8413974335a252d72687d04d06e4a1922b8",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "dfab38b3d0e7d536798728f15dc1a50de88e5d81de6a31191e4022850266f0e2",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "19004c4fa5ff6bb13323dfd83aaebe3d0c264e70ab1babb035216c8bdc7ff78d",
}
