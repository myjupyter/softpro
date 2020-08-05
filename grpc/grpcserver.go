package grpcserver

import (
    context "context"
    subs    "softpro/api/subscription"
    log     "github.com/sirupsen/logrus"
    ttool   "github.com/tarantool/go-tarantool"
)

type SportLatestIndex map[string]uint64

type GRPCServer struct {
    Conn *ttool.Connection
    SportInd SportLatestIndex
}

func (serv *GRPCServer) GetDataFromStorage() map[string]float32 {
    info := make(map[string]float32)

    for sport, i := range serv.SportInd {
        var resp *ttool.Response
        var err  error
        if i != 0 {
            resp, err = serv.Conn.Select(sport, nil, 0, 1, ttool.IterGt, []interface{}{uint(i)})
        } else {
            resp, err = serv.Conn.Call(sport + ".index.id:max", []interface{}{})
        }
        if err != nil {
            log.Fatal(err)
        }
        
        new_index := resp.Data[0].([]interface{})[0].(uint64)
        value := resp.Data[0].([]interface{})[1].(float64)
        
        log.WithFields(log.Fields{
            "sport" : sport,
            "id"    : new_index,
            "value" : value,
        }).Info("DEBUG")

        serv.SportInd[sport] = new_index
        info[sport] = float32(value) 
    }
    return info
}

func (serv *GRPCServer) SubscribeOnSportsLines(ctx context.Context, req *subs.SubsRequest) (*subs.SubsResponse, error) {
    sports := req.GetSports()

    var info map[string]float32
    if len(sports) == 0 {
        info = serv.GetDataFromStorage()
    } else {
        tempSportInd := SportLatestIndex{}
        for _, sport := range sports {
            if index, ok := serv.SportInd[sport]; ok {
                tempSportInd[sport] = index
            } else {
                tempSportInd[sport] = 0
            }
        }
        serv.SportInd = tempSportInd
        info = serv.GetDataFromStorage()
    }

    return &subs.SubsResponse{Sports: info}, nil
}
