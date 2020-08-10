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
	"github.com/arnumina/armen.core/pkg/jw"
	"github.com/arnumina/failure"
	"github.com/arnumina/logger"
)

// RunJob AFAIRE.
func (u *Utopie) RunJob(_ *jw.Job, _ *logger.Logger) *jw.Result {
	return jw.NewResult(
		failure.New(nil).Msg("this plugin has no job"), ////////////////////////////////////////////////////////////////
		0,
	)
}

/*
######################################################################################################## @(°_°)@ #######
*/
