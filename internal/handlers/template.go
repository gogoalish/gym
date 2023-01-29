package handlers

import "html/template"

var page *template.Template

func ReadTemplate(path string) error {
	temp, err := template.ParseGlob(path)
	if err != nil {
		return err
	}
	page = temp
	return nil
}
