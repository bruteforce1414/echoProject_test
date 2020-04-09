package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/api/user/:id", getUserPath)

	//http://localhost:1323/api/user?name=name1&password=password&age=17
	e.GET("/api/user", getUserQuery)

	e.POST("/api/user", postUserJson)

	e.Logger.Fatal(e.Start(":1323"))
}

// e.GET("/users/:id", getUser)
func getUserPath(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

//e.GET("/show", show)
func getUserQuery(c echo.Context) error {
	// Get team and member from the query string
	name := c.QueryParam("name")
	password := c.QueryParam("password")
	age := c.QueryParam("age")
	return c.String(http.StatusOK, "name:"+name+"\n"+"password:"+password+"\n"+"age"+age)
}

//curl -F "name=Joe Smith" -F "password=password"  -F "age=17" http://localhost:1323/api/user
// => name:Joe Smith, password:password, age=17
func postUserJson(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	password := c.FormValue("password")
	age := c.FormValue("age")

	return c.String(http.StatusOK, "name:"+name+"\n"+"password:"+password+"\n"+"age"+age)
}
