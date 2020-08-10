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

package plugin

import (
	"time"

	"github.com/arnumina/uuid"
)

type (
	// Plugin AFAIRE.
	Plugin struct {
		id      string
		name    string
		version string
		builtAt time.Time
	}
)

// New AFAIRE.
func New(name, version string, builtAt time.Time) *Plugin {
	return &Plugin{
		id:      uuid.New(),
		name:    name,
		version: version,
		builtAt: builtAt,
	}
}

// ID AFAIRE.
func (p *Plugin) ID() string {
	return p.id
}

// Name AFAIRE.
func (p *Plugin) Name() string {
	return p.name
}

// Version AFAIRE.
func (p *Plugin) Version() string {
	return p.version
}

// BuiltAt AFAIRE.
func (p *Plugin) BuiltAt() time.Time {
	return p.builtAt
}

/*
######################################################################################################## @(°_°)@ #######
*/
