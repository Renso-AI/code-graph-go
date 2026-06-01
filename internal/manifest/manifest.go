// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.0.21"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "7e3407351d8e4b61103a1dd216344aed8ad92ebdc18dba31a9cecee4a76eb0e3",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "1bd62be11f78b328f63df586f425bc2ec8042645c47cb78aab86e5c55b905e82",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "46e7169760eecf908fdd252e1523108bc817076ea643db6499ac297e6d50191a",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "03b28f5b6f93fbdb1664f394b8d6147a9ce08ab39c38158d27c3d59e2fcac942",
	"code_graph-x86_64-apple-darwin.tar.gz":           "98e74fa6186e6101aa623018b9efbe3683407c2cb24e868933887d629c5c6335",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "5180ddb8beaa449b226832db2df174a61ab25a26c85397cf883618b50fadd5ff",
	"code_graph-x86_64-pc-windows-msvc.zip":           "140169595b604ec32d0c6c64f116dab8547d5f2c49349c16988d68c0dabff734",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "8bf37ac8959bf49632eb70addf05ea81ee2c3dc1b5611678573d1845321256b6",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "606e5919f75a2d81714f56a607997e07720d7857353d25edc2e11ff2d60e5e1e",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "429c67840cefb11501d06c60ade47bb2c23409a987473d852f6a84bb9eac79ba",
}
