// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.0.26"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "51273d534afe268692bc85bebb81288afbca501b1cf08eda74e7af9de255965f",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "b1ea99271ac1f44ed461691dc647af3ad1db26f9b6674fae61758b17f582ad33",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "027fddb5781818e45b307e7656f81676b367e24b3a7e98ce4f5b0d01738432c3",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "2dbcfcf219112450805bd71bfae46a4b94ced08d9fe056c8ea6639e1edb65f38",
	"code_graph-x86_64-apple-darwin.tar.gz":           "1a065ae724807b50d137d69bd50c5fb672caf1a8efc328ec77460d3e230ba16a",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "ea93e55941bc551c6a3f49b19a30775502c1c8205c291211b5d9072aa51c0aa1",
	"code_graph-x86_64-pc-windows-msvc.zip":           "907a38e7347af7d50b4719aac2454943462d5302740f7d770d0fe2a37a11c656",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "0c6f909465696f317da92b03f080f00749f59b625e5af6806e53746f807bb3a8",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "c98eeed0215dce17de5c982a51fbd601615cc372c3a3c5394211aa9ca9d7a5c1",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "32e197218a3d1c3c09a6e180128d3bad0b6734e8665b13d56dc3d43a4237d767",
}
