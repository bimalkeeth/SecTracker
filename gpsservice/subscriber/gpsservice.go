package subscriber

import (
	"context"
	"github.com/micro/go-log"
	gpsservice "sectracker/gpsservice/proto/gpsservice"
)

type Gpsservice struct{}

func (e *Gpsservice) Handle(ctx context.Context, msg *gpsservice.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *gpsservice.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
