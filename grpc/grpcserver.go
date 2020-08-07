package grpcserver

import (
	subs "github.com/myjupyter/softpro/api/subscription"
	log "github.com/sirupsen/logrus"
	ttool "github.com/tarantool/go-tarantool"
	"io"
	"time"
)

type GRPCServer struct {
	Conn *ttool.Connection
}

func (serv *GRPCServer) GetActualDataFromStorage(sports []string) map[string]float64 {
	info := make(map[string]float64)

	for _, sport := range sports {
		resp, err := serv.Conn.Call("box.space."+sport+".index.id:max", []interface{}{})
		if err != nil {
			log.Fatal(err)
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
	return info
}

func (serv *GRPCServer) SubscribeOnSportsLines(stream subs.Subscribtion_SubscribeOnSportsLinesServer) error {
	errChan := make(chan error)
	cmdChan := make(chan bool)
	reqChan := make(chan *subs.SubsRequest)

	go func(errChan chan error, cmdChan chan bool, reqChan chan *subs.SubsRequest) {
		var req *subs.SubsRequest
		for {
			select {
			case <-cmdChan:
				return
			case req = <-reqChan:
			default:
				if req != nil {
					resp := &subs.SubsResponse{Sports: serv.GetActualDataFromStorage(req.Sports)}
					if err := stream.Send(resp); err != nil {
						errChan <- err
						return
					}
					time.Sleep(time.Duration(req.Sec) * time.Second)
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
		if req.Sec != 0 && len(req.Sports) != 0 {
			reqChan <- req
		}
	}
}
