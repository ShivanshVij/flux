package version

import (
	"github.com/loopholelabs/cmdutils/pkg/version"

	"github.com/shivanshvij/flux/internal/config"
)

var (
	// GitCommit is filled in at build time and contains the last git commit hash when this application was built
	GitCommit = ""

	// GoVersion is filled in at build time and contains the golang version upon which this application was built
	GoVersion = ""

	// Platform is filled in at build time and contains the platform upon which this application was built
	Platform = ""

	// Version is filled in at build time and contains the official release version of this application
	Version = ""

	// BuildDate is filled in at build time and contains the date when this application was build
	BuildDate = ""
)

var V = version.New[*config.Config](GitCommit, GoVersion, Platform, Version, BuildDate)
