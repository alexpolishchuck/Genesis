package main

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

func load_template(writer http.ResponseWriter, path string, data any) {
	t, err := template.ParseFiles(path)

	if err != nil {
		fmt.Println("load_template. Error occurred while loading html page " + path)
		return
	}

	path_list := strings.Split(path, "/")
	filename := path_list[len(path_list)-1]

	err = t.ExecuteTemplate(writer, filename, data)
	if err != nil {
		fmt.Println("load_template. Error occurred. " + err.Error())
	}
}
