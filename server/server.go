package main

import (
	"net/http"
	"fmt"
	"github.com/labstack/echo/v4"
)



// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
  	// User ID from path `users/:id`
  	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func main() {
	// assinging forms and their webpages
	e := echo.New()
	e.File("/", "public/index.html")
	e.File("/form", "public/form.html")
	e.POST("/submit", handleSubmit)
	e.Logger.Fatal(e.Start(":1323"))
}

// handleSubmit handles the form submission
func handleSubmit(c echo.Context) error {
    // Retrieve form data
    name := c.FormValue("name")
    email := c.FormValue("email")
    password := c.FormValue("password")

    // Print form data to the console (or you can save it to a database, etc.)
    fmt.Printf("Name: %s\n", name)
    fmt.Printf("Email: %s\n", email)
    fmt.Printf("Password: %s\n", password)

    // Send a response back to the client
    return c.String(http.StatusOK, "Form submitted successfully!")
}
