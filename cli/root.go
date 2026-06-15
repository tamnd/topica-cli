// Package cli assembles the topica command tree from the topica
// domain on top of the any-cli/kit framework.
package cli

import (
	"github.com/tamnd/any-cli/kit"
	"github.com/tamnd/topica-cli/topica"
)

// Build metadata, set via -ldflags at release time.
var (
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
)

// NewApp assembles the kit application from the topica domain. The
// domain's Register installs the client factory and every operation, so the
// binary and a host (ant, which blank-imports the package) share one source of
// truth. kit.Run turns the App into the CLI, plus the serve and mcp surfaces and
// the typed-error-to-exit-code mapping.
//
// To add a command, declare it in topica/domain.go with kit.Handle and it
// appears here automatically. Reach for app.AddCommand only for a verb that does
// not fit the emit-records shape, the way version does below.
func NewApp() *kit.App {
	id := topica.Domain{}.Info().Identity
	id.Version = Version

	app := kit.New(id)
	(topica.Domain{}).Register(app)
	app.AddCommand(newVersionCmd())
	return app
}
