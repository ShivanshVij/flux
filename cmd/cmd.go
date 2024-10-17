package main

import (
	"github.com/loopholelabs/cmdutils/pkg/command"

	"github.com/shivanshvij/flux/cmd/api"
	"github.com/shivanshvij/flux/internal/config"
	"github.com/shivanshvij/flux/version"
)

var Cmd = command.New[*config.Config](
	"flux",
	"flux manages resin 3D printers compatible with the SDCP 3.0 protocol",
	"flux manages resin 3D printers compatible with the SDCP 3.0 protocol",
	true,
	version.V,
	config.New,
	[]command.SetupCommand[*config.Config]{api.Cmd()},
)
