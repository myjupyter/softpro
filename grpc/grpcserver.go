package grpcserver

import (
    "io"
    "time"
    subs    "softpro/api/subscription"
    log     "github.com/sirupsen/logrus"
    ttool   "github.com/tarantool/go-tarantool"
)

type GRPCServer struct {
    Conn *ttool.Connection
}

func (serv *GRPCServer) GetActualDataFromStorage(sports []string) map[string]float64 {
    info := make(map[string]float64)

    for _, sport := range sports {
        resp, err := serv.Conn.Call("box.space." + sport + ".index.id:max", []interface{}{})
        if err != nil {
            log.Fatal(err)
        }
        
        index := resp.Data[0].([]interface{})[0].(uint64)
        value := resp.Data[0].([]interface{})[1].(float64)
        
        log.WithFields(log.Fields{
            "sport" : sport,
            "id"    : index,
            "value" : value,
        }).Debug("DEBUG")

        info[sport] = value
    }
    return info
}

func (serv *GRPCServer) SubscribeOnSportsLines(stream subs.Subscribtion_SubscribeOnSportsLinesServer) error {
    err_chan := make(chan error)
    cmd_chan := make(chan bool)
    req_chan := make(chan *subs.SubsRequest)

    go func(err_chan chan error, cmd_chan chan bool, req_chan chan *subs.SubsRequest) {
        var req *subs.SubsRequest
        for {
            select {
            case <- cmd_chan:
                return
            case req = <-req_chan:
            default:
                if req != nil {
                    resp := &subs.SubsResponse{Sports: serv.GetActualDataFromStorage(req.Sports)}
                    if err := stream.Send(resp); err != nil {
                        err_chan <- err
                        return
                    }
                    time.Sleep(time.Duration(req.Sec) * time.Second)
                }
            }
        }
    }(err_chan, cmd_chan, req_chan)

    for {
        req, err := stream.Recv()
        if err != nil {
            cmd_chan <- true
            return err
        }

        if err == io.EOF {
            cmd_chan <- true
            return nil
        }
        
        select {
            case err =<-err_chan:
                return err
            default:
        }
        if (req.Sec != 0 && len(req.Sports) != 0) {
            req_chan <-req
        }
    }
}
