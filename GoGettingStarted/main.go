package main

import (
	"net/http"

	"github.com/dpnetca/ps_goGetStarted/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
