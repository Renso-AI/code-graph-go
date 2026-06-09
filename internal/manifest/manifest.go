// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.1.4"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "f50877c794ba799d28cabc57dd815b11332ae73ae420b0d82ace08053587194f",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "91e91e3c1208118e067c7874f1ad7505a0db8fabdba2b86f6a0616067c922000",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "b8e31616deed8eeeaa295857b0175bd9ca6671af918d336e12a85b73e8cb4ece",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "40bd913b059cb11ac626fea9bce4101f063d4a2191721f458a815876ef525872",
	"code_graph-x86_64-apple-darwin.tar.gz":           "fe6614cd5759b09c7691fbdd00ab899e691d961542e44ddb5efd1e2fa1c25ba8",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "e9a60a7cace9289935eb00af1488bd276e250f3ad7a4ce9f7fd0b9d38360b1a8",
	"code_graph-x86_64-pc-windows-msvc.zip":           "8fd65c97b1709751ded15f8fef1167af5627510d197ce2eba3165dc12117f9cb",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "1898ff645ef672a37acb6a669f4f5125dbeabcdbf634e78fd52e1c39dbfa04a8",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "d49e740f0c142613c8595a18b2cba95d0f4cb75ae2c6478c3c19dc9587178e3b",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "1304c6e5501d49c5acc861290caa32e0e5297dbbbc157f08222811698059923e",
}
