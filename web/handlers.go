package main

import (
	"fmt"
	"net/http"
)

var htmlContent = `
<!DOCTYPE html>
<html>
<head><title>%s</title></head>
<body>
%s
</body>
</html>`

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Printf("%s %s", r.Method, r.URL.Path)
	app.render(w, "index.html", nil)
}

func (app *application) about(w http.ResponseWriter, _ *http.Request) {
	aboutContent := fmt.Sprintf(htmlContent, "About",
		"<h2>About</h2><div>Big things start with small steps</div>")
	_, _ = w.Write([]byte(aboutContent))
}

func (app *application) contact(w http.ResponseWriter, _ *http.Request) {
	contactContent := `
<h2>Contact</h2>
<div>Dive in and sample, leave your reviews</div>`
	contactContent = fmt.Sprintf(htmlContent, "Contact", contactContent)
	_, _ = w.Write([]byte(contactContent))
}