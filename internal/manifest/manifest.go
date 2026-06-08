// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.1.2"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "dede2aab18ce48298e852267bb9f965500b18f275a861b53c4bb5066f529ac89",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "b41dca4f5e554151472c6385eb8dbc11cfdf75fc37602a9e1eb37eaf5c21bed3",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "ced829e31864b3bd1a34d31503683b559b8d211bbb7496d8522f6315c8590c98",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "c9d5cc01d37a41f068864d327ca632b941c6485fb3797923cf0cc589c9e2cb9f",
	"code_graph-x86_64-apple-darwin.tar.gz":           "e99c3a7a45d41db75031e805784737bb092b3530fd4cf8d80afcd9a18cf6b7d8",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "43d9c4b1c335e8da26bb9c15f05aed91e5619a042be1d5c7eef9f8856ceadfc2",
	"code_graph-x86_64-pc-windows-msvc.zip":           "b0e69d967561911da33c2c980beb4bbf4b65fbe1063573dc57b9f19c53fee90a",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "5e6eaaf6ba0440bbb22a74db63b684bd223e0451acbad54895b3c020c283f1ad",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "ec164b47674b0677d0b271b92b25f45b9dbddac9939688da31af5fb003124cdb",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "a5f5f1fa31eafd2dc854acdb19b02fdaf5a372a846be738451793513d56b30c9",
}
