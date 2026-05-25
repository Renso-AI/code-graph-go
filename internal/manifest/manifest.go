// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.0.8"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "3091e131bdfd80e66d36c7dc67779adcdd9d510c74770f6fabb4bcd50a6ba731",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "b97474e5a4254cc851280e8646e97b3e168e09655927a34054d42bdce81e4174",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "cdca0ee78ee4b528da4140750914a482a574165e339110635c406e63afc8cdb9",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "999a67663492450286b48de03cabac96acb8435a1acbc32bc6b1076e97f5a89e",
	"code_graph-x86_64-apple-darwin.tar.gz":           "a87bec9a28b74432fc9da57bd40d95efb5dfa1d5df48b3374158e76316e5d0d0",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "9dfa60f9f7f331eeadb6bf64a18b24a24958a31de0afaf7786fd705852341ac5",
	"code_graph-x86_64-pc-windows-msvc.zip":           "82036299f30c79d4da6a6e4b8d49876dd6dbdac6e9babf791188be98db22549a",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "80dd1007a31e99068882e66fa86c6f3856d53c7d23af4bdd36b908b6c9c4ad93",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "db3aa415e314ee16cc6160533d786304f97c312ea658a91a1cc4fc7e2d78b930",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "6259f86777154cad16718e74800ca007bdce6113336dd139829c104b9f991a28",
}
