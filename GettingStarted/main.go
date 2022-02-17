package main

import (
	"net/http"

	"github.com/dpnetca/psGo/GettingStarted/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
