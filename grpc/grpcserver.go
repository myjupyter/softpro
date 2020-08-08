package grpcserver

import (
	subs "github.com/myjupyter/softpro/api/subscription"
	log "github.com/sirupsen/logrus"
	ttool "github.com/tarantool/go-tarantool"
	"io"
	"time"
)

type GRPCError struct{}

func (err *GRPCError) Error() string {
	return "Wrong request"
}

type GRPCServer struct {
	Conn *ttool.Connection
}

func (serv *GRPCServer) GetActualDataFromStorage(sports []string) (map[string]float64, error) {
	info := make(map[string]float64)

	for _, sport := range sports {
		resp, err := serv.Conn.Call("box.space."+sport+".index.id:max", []interface{}{})
		if err != nil {
			log.Warn(err)
			return nil, err
		}

		index := resp.Data[0].([]interface{})[0].(uint64)
		value := resp.Data[0].([]interface{})[1].(float64)

		log.WithFields(log.Fields{
			"sport": sport,
			"id":    index,
			"value": value,
		}).Debug("DEBUG")

		info[sport] = value
	}
	return info, nil
}

func (serv *GRPCServer) SubscribeOnSportsLines(stream subs.Subscribtion_SubscribeOnSportsLinesServer) error {
	errChan := make(chan error)
	cmdChan := make(chan bool)
	reqChan := make(chan *subs.SubsRequest)

	go func(errChan chan error, cmdChan chan bool, reqChan chan *subs.SubsRequest) {
		var req *subs.SubsRequest
		var sec int64 = 1

		for {
			select {
			case <-cmdChan:
				return
			case req = <-reqChan:
				sec = req.Sec
			case <-time.After(time.Duration(sec) * time.Second):
			}
			if req != nil {
				sports, err := serv.GetActualDataFromStorage(req.Sports)
				if err != nil {
					errChan <- err
					return
				}
				resp := &subs.SubsResponse{Sports: sports}
				if err = stream.Send(resp); err != nil {
					errChan <- err
					return
				}
			}
		}
	}(errChan, cmdChan, reqChan)

	for {
		req, err := stream.Recv()
		if err != nil {
			cmdChan <- true
			return err
		}

		if err == io.EOF {
			cmdChan <- true
			return nil
		}

		select {
		case err = <-errChan:
			return err
		default:
		}
		if req.Sec > 0 && len(req.Sports) != 0 {
			reqChan <- req
		} else {
			cmdChan <- true
			return &GRPCError{}
		}
	}
}
