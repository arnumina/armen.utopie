// +build devel

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

package data

//go:generate vfsgendev -tag devel -source="github.com/arnumina/armen.utopie/data".Static

import "net/http"

// Static AFAIRE.
var Static http.FileSystem = http.Dir("static")

/*
######################################################################################################## @(°_°)@ #######
*/
