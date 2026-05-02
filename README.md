# toolbox-nix

A codefly toolbox plugin for Nix flake introspection and (read-only)
expression evaluation. Canonical owner of the `nix` binary — the
`codefly-dev/toolbox-bash` plugin refuses every `nix` invocation and
routes callers here.

## Tools (read-only)

- `nix.flake_metadata(flake?)` — wraps `nix flake metadata --json`.
  Returns description, lastModified, narHash, original ref. Path or
  URL accepted; defaults to `.`.
- `nix.flake_show(flake?)` — wraps `nix flake show --json`. Surfaces
  the flake's output structure: packages, devShells, apps.
- `nix.eval(expr, timeout_ms?)` — wraps `nix eval --json --read-only --expr <expr>`.
  Read-only mode forbids store mutations. 30s default timeout, 5min cap.
  Output capped at 4 MiB.

## Configuration

| Env var                     | Default       | Purpose                                                |
| --------------------------- | ------------- | ------------------------------------------------------ |
| `CODEFLY_TOOLBOX_VERSION`   | `0.0.0-dev`   | Identity version surfaced via `Identity()`             |
| `CODEFLY_TOOLBOX_NIX_BIN`   | _PATH lookup_ | Override the `nix` binary path (mostly for tests)      |

The toolbox unconditionally appends `--extra-experimental-features
nix-command flakes` so it works on stock nix installs without
requiring `nix.conf` tweaks.

## Build & test

```bash
go build ./...
go test ./...
```

## Contract

This plugin implements the codefly Toolbox gRPC contract defined in
[`codefly-dev/core`](https://github.com/codefly-dev/core) at
`proto/codefly/services/toolbox/v0/toolbox.proto`.
