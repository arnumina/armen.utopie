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

package web

import (
	"net/http"

	"github.com/arnumina/armen.core/pkg/model"
	"github.com/arnumina/logger"

	"github.com/arnumina/armen.utopie/data"
	"github.com/arnumina/armen.utopie/internal/pkg/container"
)

type (
	tData struct {
		Application container.Application
		Instances   []*model.Instance
	}

	// Web AFAIRE.
	Web struct {
		application container.Application
		logger      *logger.Logger
		model       container.Model
	}
)

// New AFAIRE.
func New(ctn container.Container) *Web {
	return &Web{
		application: ctn.Application(),
		logger:      ctn.Logger(),
		model:       ctn.Model(),
	}
}

// Build AFAIRE.
func (w *Web) Build(ctn container.Container) *Web {
	router := ctn.Server().Router()

	router.HandleFunc("/", w.home).Methods("GET")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(data.Static)))

	return w
}

func (w *Web) home(rw http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.notFound(rw)
		return
	}

	w.instances(rw, r)
}

func (w *Web) instances(rw http.ResponseWriter, _ *http.Request) {
	ais, err := w.model.AllInstances()
	if err != nil {
		w.serverError(rw, err)
		return
	}

	td := &tData{
		Application: w.application,
		Instances:   ais,
	}

	t, err := parseTemplates(
		"html/page.instances.html",
		"html/layout.base.html",
	)
	if err != nil {
		w.serverError(rw, err)
		return
	}

	err = t.Execute(rw, td)
	if err != nil {
		w.serverError(rw, err)
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
