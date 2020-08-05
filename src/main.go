package main

import (
    "net"
    "net/http"
    "softpro/work"
    subs     "softpro/api/subscription"
    grpcserv "softpro/grpc"
    grpc     "google.golang.org/grpc"
    log      "github.com/sirupsen/logrus"
    ttool    "github.com/tarantool/go-tarantool"
    
)

const (
    Url = "http://localhost:8000/api/v1/lines/"
    DBAddress = "127.0.0.1:3301"
    Address = ":8081"
)

//TODO: передача аргументов в приложение:
// 1. Адрес HTTP
// 2. Адресс GRPC
// 3. Секунды на каждый вид спорта
// 4. Адрес хранилища
// 5. Уровень логирования
func main() {
    log.SetLevel(log.DebugLevel)

    // TODO: настроить передачу времени
    sports := []work.SportSubs{
        work.SportSubs{"soccer",   3},
        work.SportSubs{"football", 2},
        work.SportSubs{"baseball", 4},
    }

    conn, err := ttool.Connect(DBAddress, ttool.Opts{})
    if err != nil {
        log.WithFields(log.Fields{
            "who" : "Data Storage",
        }).Fatal(err)
        return
    }
    defer conn.Close()
    log.Debug("Connected to Data Storage")

    var pool work.WorkerPool
    pool.Start(conn, Url, sports...)
    log.Debug("Pool has been run")

    
    log.Debug("HTTP server listens")
    http.HandleFunc("/ready", pool.ReadyHandler)
    go http.ListenAndServe(Address, nil) 

    
    log.Debug("GRPC server listens")
    s := grpc.NewServer()
    srv := &grpcserv.GRPCServer{Conn: conn}

    subs.RegisterSubscribtionServer(s, srv)
    l, err := net.Listen("tcp", ":8080")
    if err != nil {
        log.Fatal(err)
    }

    if err := s.Serve(l); err != nil {
        log.Fatal(err)
    }
}
