package work

import (
    "fmt"
    "net/http"
    "time"
    "encoding/json"
    "github.com/tarantool/go-tarantool"
)

type WorkerPool struct {
    Workers       []Worker
    WorkersStates []*WorkerState
}

func (wp *WorkerPool) Start(conn *tarantool.Connection, url string, sports ...SportSubs) {
    for i, sport := range sports {
        wp.Workers = append(wp.Workers, Worker{
            ID: uint(i),
            URL: url,
            Subs: sport,
            Conn: conn,
        })
        wp.WorkersStates = append(wp.WorkersStates, new(WorkerState))
        wp.Workers[i].Start(wp.WorkersStates[i])
    }
}

func (wp *WorkerPool) CheckWorkersSync() bool {
    
    for i := 0; i < len(wp.WorkersStates); i++ {
        for true {
            if wp.WorkersStates[i].GetError() != nil {
                return false
            }
            if wp.WorkersStates[i].IsSync() {
                break
            }
            time.Sleep(1 * time.Second)
        }
    }
    return true
}

func (wp *WorkerPool) CheckWorkerStatus() bool {
    for _, worker_state := range wp.WorkersStates {
        if worker_state.GetError() != nil {
            return false
        }
    }
    return true
}

func (wp *WorkerPool) ReadyHandler(w http.ResponseWriter, req *http.Request) {
    if wp.CheckWorkerStatus() {
        w.WriteHeader(http.StatusOK)
    } else {
        w.WriteHeader(http.StatusInternalServerError)
    }
   
    var states []string
    for _, state := range wp.WorkersStates {
        states = append(states, state.GetJSONInfo())
    }
    info := map[string][]string{
        "server": states,
    }

    status_json, _ := json.Marshal(info)

    fmt.Fprintf(w, string(status_json))
}


