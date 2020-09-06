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
	"html/template"
	"io/ioutil"
	"path/filepath"

	"github.com/arnumina/armen.utopie/data"
)

func readTemplate(file string) ([]byte, error) {
	f, err := data.Static.Open(file)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	return ioutil.ReadAll(f)
}

func parseTemplates(files ...string) (*template.Template, error) {
	var t *template.Template

	for _, file := range files {
		b, err := readTemplate(file)
		if err != nil {
			return nil, err
		}

		s := string(b)
		name := filepath.Base(file)

		var tmpl *template.Template

		if t == nil {
			t = template.New(name)
		}

		if name == t.Name() {
			tmpl = t
		} else {
			tmpl = t.New(name)
		}

		_, err = tmpl.Parse(s)
		if err != nil {
			return nil, err
		}
	}

	return t, nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
