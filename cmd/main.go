package main 

import (
	"html/template"
	"io"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
)

/*
This struct holds the parsed templates. It has a method Render that executes a specified 
template with the provided data.
*/
type Templates struct {
	templates *template.Template
}

/*
This method takes a writer, the name of the template, the data to be passed to the template, 
and the Echo context. It executes the template and writes the output to the writer. 
*/
func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

/*
This function initializes the Templates struct by parsing all HTML files in the views directory. 
It uses template.Must to ensure that any error during parsing will cause a panic, making it 
easier to catch issues at startup.
*/
func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

type Count struct {
	Count int
}

/*
1. Creating an Echo Instance: e := echo.New() initializes a new Echo instance.
2. Middleware: e.Use(middleware.Logger()) adds logging middleware to log HTTP requests.
3. Count Initialization: count := Count{Count: 0} initializes the counter to zero.
4. Template Renderer: e.Renderer = newTemplate() sets the custom template renderer.
5. GET Route: The route e.GET("/") renders the "index" template with the current count 
when the root URL is accessed.
6. POST Route: The route e.POST("/count") increments the count and renders the "count" 
template with the updated count when a POST request is made to /count.
7. Starting the Server: e.Logger.Fatal(e.Start(":42069")) starts the server on port 
42069 and logs any fatal errors.
*/

func main() {
	// instantiating new echo thing
	e := echo.New()
	// adding middleware to it
	e.Use(middleware.Logger())

	// the counter for count ?
	count := Count{Count: 0}
	e.Renderer = newTemplate()
	
	// each get request to the index page that updates count by one
	e.GET("/", func(c echo.Context) error {
		count.Count++
		return c.Render(200, "index", count) 
	})

	// each post request sent with button increase count too	
	e.POST("/count", func(c echo.Context) error {
		count.Count++
		return c.Render(200, "count", count) 
	})

	e.Logger.Fatal(e.Start(":42069"))
}