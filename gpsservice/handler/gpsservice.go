package handler

import (
	"context"
	"github.com/micro/go-log"
	gpsservice "sectracker/gpsservice/proto/gpsservice"
)

type Gpsservice struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Gpsservice) Call(ctx context.Context, req *gpsservice.Request, rsp *gpsservice.Response) error {
	log.Log("Received Gpsservice.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Gpsservice) Stream(ctx context.Context, req *gpsservice.StreamingRequest, stream gpsservice.Gpsservice_StreamStream) error {
	log.Logf("Received Gpsservice.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&gpsservice.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Gpsservice) PingPong(ctx context.Context, stream gpsservice.Gpsservice_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&gpsservice.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
