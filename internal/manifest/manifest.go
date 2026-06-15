// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.3.1"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "d1a1afaa6013b7c7d5d69ce905137ed2a162847e47b7941198cd7bdb31615170",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "efc9e4585595c6f557238942e525fc5ffa338b481663a30f0ea1d1d6ff420696",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "4e619fee8a6294149195e5016a43c7ba8e6aad126ad433492658406a1474c013",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "8916fd790aaf107c61e678e0e79506f529ac4006d6fe997308b91f53071751af",
	"code_graph-x86_64-apple-darwin.tar.gz":           "bea51965d5a93eb3d588fd916a77fbbea3e3903caef91ddba2b64519049c77e1",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "4bec7bdf6773b860ad07cdec28a46e97e708ebdc85b4932a2e8b1fe002efb5ba",
	"code_graph-x86_64-pc-windows-msvc.zip":           "1f62aa84ef47772bd3603f862a9f5e5e3f53f1690e224c0a67041d35d8dbdc87",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "189fea5768da206fab17d782472451ad822bbd8fecc4c12ed697ded058d12e58",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "8e7e3a6b6350fadb82548684e3ebd9bec040863acbf39712b955e5ee1302d1ca",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "2f5db3edd4f7ad7b2ad28dc4ea5dd35cf9b927bfa338b58cbb39d42fe4690789",
}
