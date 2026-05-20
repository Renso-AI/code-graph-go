# code-graph-go

`go install` launcher for [code_graph](https://cg.renso.ai).

## Install

```sh
go install github.com/renso-ai/code-graph-go/cmd/code_graph@latest
go install github.com/renso-ai/code-graph-go/cmd/code_graph-mcp@latest
```

## How it works

This module is a thin download-on-first-run launcher. `go install`
compiles a small Go binary; on first invocation it downloads the
prebuilt `code_graph` (or `code_graph-mcp`) binary for the host
platform from the
[main repo's GH Releases](https://github.com/Renso-AI/code_graph/releases),
SHA256-verifies it against the manifest baked into this module, and
execs it. Subsequent runs hit the cached binary directly under
`$XDG_CACHE_HOME/code-graph/<version>/`.

This module contains **no engine source**. The implementation lives
at https://github.com/Renso-AI/code_graph (Apache-2.0).

## How releases land here

This repo is auto-synced by the release workflow in
[Renso-AI/code_graph](https://github.com/Renso-AI/code_graph) — see
`packaging/wrappers/go/` in that repo for the templates and the
`publish-go` job in `.github/workflows/release.yml` for the
machinery.

Do not edit files in this repo by hand; the next release tag
overwrites the tree.

## License

Apache-2.0.
