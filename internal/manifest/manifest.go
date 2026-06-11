// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.1.13"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "b55e72ea5cdf15627d6999fca5ffef0f88465781396cff68bf1253370ebd78ea",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "71020ae58ff491de599e1146c0298640fa8f6374fd6d79a1e990da143889f6dc",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "3bc7369d74bd923327b2baa3fe72ccbae502495d3a0b737ce51df04efcfee6cb",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "77a8d8144b506e48f117465cdd6a959a2122918812f4d14648ebc198c4ce207b",
	"code_graph-x86_64-apple-darwin.tar.gz":           "cee5414c5a9f2b983a2dde4e27d5325d702c6911be1da746ed43b062c3380e96",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "f409ce52624f69865703cef555fb1b6058ae421f9bf2cdab20d2b19c36f67076",
	"code_graph-x86_64-pc-windows-msvc.zip":           "eb63a83212004d90e7fb1ad83f1e710a03e57d0e4c7105a058aa7805012e3e1e",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "d58f659767516febfbff8090deb9102b997f5a8969588e92f11838b530d13f2c",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "13d2ee7a7e4a87992afb18f6b1e21461d40e94e904b6174bfeac2e51b8ee8fd8",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "3a7c4876c01e69a71830ec2047fe7fb49270c75abd072534d3c0e07a69c48b4e",
}
