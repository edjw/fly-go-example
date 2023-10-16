package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

//go:embed templates/*
var resources embed.FS
var t = loadTemplate("templates/*.tmpl")

func loadTemplate(templatePath string) *template.Template {
	return template.Must(template.ParseFS(resources, templatePath))
}

func runTemplate(t *template.Template, templateFile string, writerPointer http.ResponseWriter, data interface{}) {
	err := t.ExecuteTemplate(writerPointer, templateFile, data)
	if err != nil {
		http.Error(writerPointer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {

	port, portExists := os.LookupEnv("PORT")
	env, envExists := os.LookupEnv("GOENVIRONMENT")

	var ip string = "0.0.0.0" // Standard IP address on Fly and Render

	if envExists && env == "development" {
		ip = "127.0.0.1"
	}
	// For development, this relies on setting the GOENVIRONMENT variable
	// to "development" in your .zshrc or .bashrc file with
	// export GOENVIRONMENT=development

	if !portExists {
		port = "8080"
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"Region": os.Getenv("FLY_REGION"),
			"Name":   "Ed",
		}

		runTemplate(t, "index.html.tmpl", writer, data)

	})

	server := &http.Server{
		Addr:         ip + ":" + port,
		Handler:      nil,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Printf("Listening on http://%s:%v", ip, port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
