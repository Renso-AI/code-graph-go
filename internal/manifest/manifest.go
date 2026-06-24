// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.4.2"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "03522b72451f4b4ab2a3567ac722261e53775bf0c09dd44fdd833d0f12d10e9d",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "53a7076c17a3c4d791798e85af620698f7ddfb22cd3a47039f33a049c550eedf",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "940080abff48b7a3e953f2c711fa7b1ab1161acfeac05a1b99ddbf56d4daea34",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "3113df83f1e8512ef3d8d39d9b64a926d432aff7ba3e3f42b4a94149d0e71ac9",
	"code_graph-x86_64-apple-darwin.tar.gz":           "987960050f1d723baf110c05609aeea6f3412206447bc7d96cb2a07de797b5e7",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "48625fdf3df50aa3433bab9eea8ab71d51f4485a563f13b65aa5ff9e26a14867",
	"code_graph-x86_64-pc-windows-msvc.zip":           "b6bba938e88ef22f0ec32b023555a6e52bf1370055d5f667b3e8ae18f327024f",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "fa7d63c1c4c15eee2d4177346fbc8148174591f7188732405adf498733708dd0",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "f428d1c9d5267fd7d8e9b9a54525312f7ef979c4f3db9d25ae3b12e080b24661",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "3f4c6ff30b8c7d93e2ad8ba4d0a334ffaa35ea902244ebdd35dc15602e170e62",
}
