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

func getEnvWithDefault(name, defaultValue string) string {

	// LookupEnv returns a string value for the given key if the key exists in the
	// environment, else the second return value is false.
	val, exists := os.LookupEnv(name)
	if !exists {
		return defaultValue
	}
	return val
}

func runTemplate(t *template.Template, templateFile string, writerPointer http.ResponseWriter, data interface{}) {
	err := t.ExecuteTemplate(writerPointer, templateFile, data)
	if err != nil {
		http.Error(writerPointer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {

	port := getEnvWithDefault("PORT", "8080")
	env := getEnvWithDefault("GOENVIRONMENT", "development") // This relies on setting the GOENVIRONMENT variable to "development" in your .zshrc or.bashrc file

	var ip string = "0.0.0.0"

	if env == "development" {
		ip = "127.0.0.1"
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
