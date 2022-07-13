package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DisplayMode int

const (
	Home DisplayMode = iota
	Register
	Download
)

func main() {
	m := make(map[string]int)
	m["1005.com"] = int(Home)
	m["1202.com"] = int(Home)
	m["1208.com"] = int(Home)
	m["1203.com"] = int(Home)
	m["1204.com"] = int(Home)
	m["1302.com"] = int(Home)
	m["1303.com"] = int(Register)
	m["1310.com"] = int(Register)
	m["1307.com"] = int(Register)
	m["1305.com"] = int(Register)
	m["1308.com"] = int(Register)
	m["1304.com"] = int(Register)
	m["1306.com"] = int(Download)
	m["1406.com"] = int(Download)
	m["1403.com"] = int(Download)
	m["1422.com"] = int(Download)
	m["1430.com"] = int(Download)
	m["1405.com"] = int(Download)
	m["1432.com"] = int(Download)
	router := gin.Default()

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
	//download,site
	router.GET("/displaymode", func(c *gin.Context) {
		site := c.DefaultQuery("site", "567bet.in")
		displaymode := "Home"
		v, exist := m[site]
		if exist {
			switch v {
			case int(Home):
				displaymode = "Home"
			case int(Register):
				displaymode = "Register"
			case int(Download):
				displaymode = "Download"
			}
		}
		c.String(http.StatusOK, "Hello %s", displaymode)
	})
	router.Run(":8080")
}
