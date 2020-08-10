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
)

func (w *Web) serverError(rw http.ResponseWriter, err error) {
	w.logger.Error( //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
		http.StatusText(http.StatusInternalServerError),
		"reason", err.Error(),
	)

	http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (w *Web) clientError(rw http.ResponseWriter, status int) {
	http.Error(rw, http.StatusText(status), status)
}

func (w *Web) notFound(rw http.ResponseWriter) {
	w.clientError(rw, http.StatusNotFound)
}

/*
######################################################################################################## @(°_°)@ #######
*/
