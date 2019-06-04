package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"sectracker/gpsservice/handler"
	gpsservice "sectracker/gpsservice/proto/gpsservice"
	"sectracker/gpsservice/subscriber"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.gpsservice"),
		micro.Version("latest"),
	)
	// Initialise service
	service.Init()

	// Register Handler
	_ = gpsservice.RegisterGpsserviceHandler(service.Server(), new(handler.Gpsservice))

	// Register Struct as Subscriber
	_ = micro.RegisterSubscriber("go.micro.srv.gpsservice", service.Server(), new(subscriber.Gpsservice))

	// Register Function as Subscriber
	_ = micro.RegisterSubscriber("go.micro.srv.gpsservice", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
