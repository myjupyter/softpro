package grpcserver

import (
    context "context"
    subs    "softpro/api/subscription"
    log     "github.com/sirupsen/logrus"
)

type GRPCServer struct {}

func (serv *GRPCServer) SubscribeOnSportsLines(ctx context.Context, req *subs.SubsRequest) (*subs.SubsResponse, error) {
    sports := req.GetSports()
    sec := req.GetSec()

    log.WithFields(log.Fields{
        "sec" : sec,
        "sports" : sports,
    }).Info("From GRPC")

    return &subs.SubsResponse{Sports: map[string]float32{"soccer" : 1.1,}}, nil
}
