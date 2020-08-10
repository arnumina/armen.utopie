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

package container

import (
	"time"

	"github.com/arnumina/armen.core/pkg/resources"
	"github.com/arnumina/logger"
)

type (
	// Application AFAIRE.
	Application interface {
		resources.Application
	}

	// Bus AFAIRE.
	Bus interface {
		resources.Bus
	}

	// Model AFAIRE.
	Model interface {
		resources.Model
	}

	// Plugin AFAIRE.
	Plugin interface {
		ID() string
		Name() string
		Version() string
		BuiltAt() time.Time
	}

	// Server AFAIRE.
	Server interface {
		resources.Server
	}

	// Util AFAIRE.
	Util interface {
		resources.Util
	}

	// Container AFAIRE.
	Container interface {
		Application() Application
		Bus() Bus
		Logger() *logger.Logger
		Model() Model
		Server() Server
		Plugin() Plugin
		Util() Util
	}

	// Resources AFAIRE.
	Resources struct {
		application Application
		bus         Bus
		logger      *logger.Logger
		model       Model
		server      Server
		plugin      Plugin
		util        Util
	}
)

// New AFAIRE.
func New() *Resources {
	return &Resources{}
}

// Application AFAIRE.
func (r *Resources) Application() Application {
	return r.application
}

// SetApplication AFAIRE.
func (r *Resources) SetApplication(application Application) {
	r.application = application
}

// Bus AFAIRE.
func (r *Resources) Bus() Bus {
	return r.bus
}

// SetBus AFAIRE.
func (r *Resources) SetBus(bus Bus) {
	r.bus = bus
}

// Logger AFAIRE.
func (r *Resources) Logger() *logger.Logger {
	return r.logger
}

// SetLogger AFAIRE.
func (r *Resources) SetLogger(logger *logger.Logger) {
	r.logger = logger
}

// Model AFAIRE.
func (r *Resources) Model() Model {
	return r.model
}

// SetModel AFAIRE.
func (r *Resources) SetModel(model Model) {
	r.model = model
}

// Plugin AFAIRE.
func (r *Resources) Plugin() Plugin {
	return r.plugin
}

// SetPlugin AFAIRE.
func (r *Resources) SetPlugin(plugin Plugin) {
	r.plugin = plugin
}

// Server AFAIRE.
func (r *Resources) Server() Server {
	return r.server
}

// SetServer AFAIRE.
func (r *Resources) SetServer(server Server) {
	r.server = server
}

// Util AFAIRE.
func (r *Resources) Util() Util {
	return r.util
}

// SetUtil AFAIRE.
func (r *Resources) SetUtil(util Util) {
	r.util = util
}

/*
######################################################################################################## @(°_°)@ #######
*/
