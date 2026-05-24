// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.0.4"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "9abc1b2de5fe268f9cfccbd8462b6b58ad293bbb0240a972bb85a74204d2a20c",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "ceb41460368c9f6400cc65741dec926badd753a188c918f9b4e5a574a18787b4",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "6acf9f9960eca916cc75097963dc8319dfcc4735aa2c068d88b6d0aaefeebe2b",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "abc20ef6156faa78b39bedc56317b8179bc5d53c6dc2a1e5998aaed7823fd011",
	"code_graph-x86_64-apple-darwin.tar.gz":           "a4c6db327eeac0fefb767c88b001d76c33fc285084e3e3c0790120e42280437c",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "4fe1753a132cc2d017acace3158b49a1552b29a1448539535255c4f4c9232b6e",
	"code_graph-x86_64-pc-windows-msvc.zip":           "7a8f06705289e066501b41b74e066c1b26bc89203b1c2d43611aaf7280e1273d",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "882491e957be6786ec47a07296675d332ce0aade3889dfd31b45a683ffd93438",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "116e40bafbacce0f78f36bb060a5b8871a2db16232e1ec7969549162e9e5ee89",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "a44e565b1748cb961d09f1a10d49cc8021fd81c342aa29e9d8c06396c0197f83",
}
