/*
#######
##                                          __            _
##       ___ _______ _  ___ ___       __ __/ /____  ___  (_)__
##      / _ `/ __/  ' \/ -_) _ \  _  / // / __/ _ \/ _ \/ / -_)
##      \_,_/_/ /_/_/_/\__/_//_/ (_) \_,_/\__/\___/ .__/_/\__/
##                                               /_/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package utopie

import (
	"github.com/arnumina/armen.core/pkg/resources"

	"github.com/arnumina/armen.utopie/internal/pkg/container"
	"github.com/arnumina/armen.utopie/internal/pkg/plugin"
	"github.com/arnumina/armen.utopie/internal/pkg/web"
)

type (
	// Utopie AFAIRE.
	Utopie struct {
		*plugin.Plugin
		cr *container.Resources
	}
)

// New AFAIRE.
func New(version, builtAt string, resources resources.Container) *Utopie {
	util := resources.Util()
	plugin := plugin.New("utopie", version, util.UnixToTime(builtAt))
	logger := resources.Logger().Clone(util.LoggerPrefix(plugin.Name(), plugin.ID()))

	cr := container.New()
	cr.SetUtil(util)
	cr.SetPlugin(plugin)
	cr.SetLogger(logger)
	cr.SetApplication(resources.Application())
	cr.SetModel(resources.Model())
	cr.SetBus(resources.Bus())
	cr.SetServer(resources.Server())

	return &Utopie{
		Plugin: plugin,
		cr:     cr,
	}
}

// Build AFAIRE.
func (u *Utopie) Build() error {
	// Web
	//..................................................................................................................
	_ = web.New(u.cr).Build(u.cr)

	return nil
}

// Close AFAIRE.
func (u *Utopie) Close() {}

/*
######################################################################################################## @(°_°)@ #######
*/
