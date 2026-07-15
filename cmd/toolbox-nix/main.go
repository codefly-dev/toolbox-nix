// Command toolbox-nix is the standalone binary form of the codefly
// nix toolbox plugin. Loaded via the standard agent loader
// (core/agents/manager.Load); registers a Toolbox server through
// agents.Serve.
//
// Configuration:
//
//	CODEFLY_TOOLBOX_VERSION   — Identity version. Default "0.0.0-dev".
//	CODEFLY_TOOLBOX_NIX_BIN   — override the nix binary path. Mostly
//	                             useful for tests; production callers
//	                             leave it unset and rely on PATH.
package main

import (
	"github.com/codefly-dev/core/agents"
	coretoolbox "github.com/codefly-dev/core/toolbox"
	nix "github.com/codefly-dev/toolbox-nix"
)

func main() {
	server := nix.New(coretoolbox.Version())
	if bin := coretoolbox.Environment("CODEFLY_TOOLBOX_NIX_BIN", ""); bin != "" {
		server = server.WithBinary(bin)
	}
	agents.ServeToolbox(server)
}
