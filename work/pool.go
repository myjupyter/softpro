package work

import (
    "fmt"
    "net/http"
    "github.com/tarantool/go-tarantool"
)

type WorkerPool struct {
    Workers []Worker
}

func (wp *WorkerPool) Start(conn *tarantool.Connection, url string, sports ...SportSubs) {
    for i, sport := range sports {
        wp.Workers = append(wp.Workers, Worker{
            ID: uint(i),
            URL: url,
            Subs: sport,
            Conn: conn,
            Status: make(chan interface{}),
            CheckStat: make(chan bool),
        })
        wp.Workers[i].Start()
    }
}

// TODO: собирать информацию о воркерах
func (wp *WorkerPool) CheckWorkersStatus() bool {
    for _, worker := range wp.Workers {
        worker.CheckStat <- true
        if err := <-worker.Status; err != nil {
            return false
        } 
    }
    return true
}

func (wp *WorkerPool) ReadyHandler(w http.ResponseWriter, req *http.Request) {
    if wp.CheckWorkersStatus() {
        w.WriteHeader(http.StatusOK)
        //TODO: Более детальная информация
        fmt.Fprintf(w, "OK")
        return
    }
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, "Not OK")
}


