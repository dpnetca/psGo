package main

import (
	"context"
	"fmt"
	stlog "log"

	"github.com/dpnetca/psGo/distributedApplications/log"
	"github.com/dpnetca/psGo/distributedApplications/registry"
	"github.com/dpnetca/psGo/distributedApplications/service"
)

func main() {
	log.Run("./app.log")
	host, port := "localhost", "4000"

	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	var r registry.Resgistration
	r.ServiceName = registry.LogService
	r.ServiceURL = serviceAddress

	ctx, err := service.Start(
		context.Background(),
		host,
		port,
		r,
		log.RegisterHandlers,
	)
	if err != nil {
		stlog.Fatal(err)
	}
	<-ctx.Done()
	fmt.Println("Shutting down log service")
}
