// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.0.16"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "90db5f78f784d22dea80de73ce7e143c75f6cc6efac33708cd9f133f5f4b0939",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "3b170f4d4490030c4ed31b6ffe937d979510ddee4d8eeedc1a4a3b532a6c2932",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "c5b47394442141a8620b7ebb327985c7f6ad66b7fa231d49e972fc9fdf0590de",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "7ce1cdaf38d45a36bb0c4855e04f47a01e2558793998556e30eeef5cb93c19be",
	"code_graph-x86_64-apple-darwin.tar.gz":           "e1bdfb362e7f62a608469b104f759d8c7c138431a83c0e4cd3431f041ef7caf1",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "eefb4b58ce8eb80ac253b8f81eb8ff4ec17723dc61971145a13352ea6b960bf3",
	"code_graph-x86_64-pc-windows-msvc.zip":           "0248f3bac233bdd5ee3fde7ad535c5a26496821635772e540bac1911689813fc",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "5bd016a9626c81f94fec0c1484b2b3b359e20959a549e32dee3e25440ed2f5fd",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "4ed35d6097a5bb8e20f1b861cd148ba7ab41329b6d9c6dbadc8a24507526b60f",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "290a0c0502fb615e6b945da73a43e6fad777d2a36d3a5d4a46acf885a7c9c935",
}
