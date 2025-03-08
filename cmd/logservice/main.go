package main

import (
	"context"
	"disysytem/log"
	"disysytem/service"
	stlog "log"
)

func main() {
	log.Run("./distributed.log")
	host, port := "localhost", "4000" //一般来说这里要有一个配置文件，但是这里只是一个简单demo就不搞了
	ctx, err := service.Start(
		context.Background(),
		"Log Service",
		host,
		port,
		log.RegisterHandlers,
	)
	if err != nil {
		stlog.Fatalln(err)
	}
	<-ctx.Done()
	stlog.Println("Shutting down log service")
}
