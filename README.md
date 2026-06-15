# topica

A command line for topica.

`topica` is a single pure-Go binary. It reads public topica data
over plain HTTPS, shapes it into clean records, and prints output that pipes
into the rest of your tools. No API key, nothing to run alongside it.

The same package is also a [resource-URI driver](#use-it-as-a-resource-uri-driver),
so a host program like [ant](https://github.com/tamnd/ant) can address
topica as `topica://` URIs.

## Install

```bash
go install github.com/tamnd/topica-cli/cmd/topica@latest
```

Or grab a prebuilt binary from the [releases](https://github.com/tamnd/topica-cli/releases), or run
the container image:

```bash
docker run --rm ghcr.io/tamnd/topica:latest --help
```

## Usage

```bash
topica page <path>                      # fetch one page as a record
topica page <path> -o json              # as JSON, ready for jq
topica page <path> --template '{{.Body}}'  # just the readable body text
topica links <path>                     # the pages it links to, one per line
topica --help                           # the whole command tree
```

Every command shares one output contract: `-o table|json|jsonl|csv|tsv|url|raw`,
`--fields` to pick columns, `--template` for a custom line, and `-n` to limit.
The default adapts to where output goes (a table on a terminal, JSONL in a
pipe), so the same command reads well by hand and parses cleanly downstream.

This is a fresh scaffold. It ships one example resource type, `page`, wired end
to end. Model the real topica records in `topica/` and declare their
operations in `topica/domain.go`; each one becomes a command, an HTTP
route, and an MCP tool at once.

## Serve it

The same operations are available over HTTP and as an MCP tool set for agents,
with no extra code:

```bash
topica serve --addr :7777    # GET /v1/page/<path>  returns NDJSON
topica mcp                   # speak MCP over stdio
```

## Use it as a resource-URI driver

`topica` registers a `topica` domain the way a program registers a
database driver with `database/sql`. A host enables it with one blank import:

```go
import _ "github.com/tamnd/topica-cli/topica"
```

Then [ant](https://github.com/tamnd/ant) (or any program that links the package)
dereferences `topica://` URIs without knowing anything about topica:

```bash
ant get topica://page/<path>   # fetch the record
ant cat topica://page/<path>   # just the body text
ant ls  topica://page/<path>   # the pages it links to, each addressable
ant url topica://page/<path>   # the live https URL
```

## Development

```
cmd/topica/   thin main: hands cli.NewApp to kit.Run
cli/                 assembles the kit App from the topica domain
topica/                the library: HTTP client, data models, and domain.go (the driver)
docs/                tago documentation site
```

```bash
make build      # ./bin/topica
make test       # go test ./...
make vet        # go vet ./...
```

## Releasing

Push a version tag and GitHub Actions runs GoReleaser, which builds the
archives, Linux packages, the multi-arch GHCR image, checksums, SBOMs, and a
cosign signature:

```bash
git tag v0.1.0
git push --tags
```

The Homebrew and Scoop steps self-disable until their tokens exist, so the first
release works with no extra secrets.

## License

Apache-2.0. See [LICENSE](LICENSE).
