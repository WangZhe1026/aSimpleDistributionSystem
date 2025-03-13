package main

import (
	"context"
	"disysytem/log"
	"disysytem/registry"
	"disysytem/service"
	"fmt"
	stlog "log"
)

func main() {
	log.Run("./distributed.log")
	host, port := "localhost", "4000" //一般来说这里要有一个配置文件，但是这里只是一个简单demo就不搞了
	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)
	r := registry.Registration{
		ServiceName: "Log Service",
		ServiceURL:  serviceAddress,
	}

	ctx, err := service.Start(
		context.Background(),
		host,
		port,
		r,
		log.RegisterHandlers,
	)

	if err != nil {
		stlog.Fatalln(err)
	}
	<-ctx.Done()
	stlog.Println("Shutting down log service")
}
