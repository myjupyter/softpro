package main

import (
    "net"
    "net/http"
    "softpro/work"
    "flag"
    subs     "softpro/api/subscription"
    grpcserv "softpro/grpc"
    grpc     "google.golang.org/grpc"
    log      "github.com/sirupsen/logrus"
    ttool    "github.com/tarantool/go-tarantool"
    viper    "github.com/spf13/viper"    
)

var (
    Url      = "http://localhost:8000/api/v1/lines/"
    DBAddr   = "127.0.0.1:3301"
    HttpAddr = "127.0.0.1:8001"
    GrpcAddr = "127.0.0.1:8002"
    
    LogLevel = map[string]log.Level {
        "debug"   : log.DebugLevel ,
        "info"    : log.InfoLevel,
        "warning" : log.WarnLevel,
        "error"   : log.ErrorLevel,
        "fatal"   : log.FatalLevel,
    }
)

func main() {
    // Reads flags 
    name := flag.String("name", "config", "Configuration file name without extension")
    dir := flag.String("dir", ".", "Configuration file directory")
    
    flag.Parse()

    // Loads configs
    viper.SetConfigName(*name)
    viper.AddConfigPath(*dir)
    if err := viper.ReadInConfig(); err != nil {
        log.Fatal(err)
    }

    // Sets Logger Level
    level, ok := LogLevel[viper.GetString("loglevel")]
    if ok {
        log.SetLevel(level)
    } else {
        log.Warning("Wrong logger level. Debug level is setted by default")
        log.SetLevel(log.DebugLevel)
    }

    var sports []work.SportSubs
    for sport, sec := range viper.GetStringMap("lineprovider.sports") {
        sports = append(sports, work.SportSubs{Sport: sport, Seconds: int(sec.(float64))})
    }

    // Connection to the Data Storage
    DBAddr = viper.GetString("datastorage.address") + ":" + viper.GetString("datastorage.port")
    conn, err := ttool.Connect(DBAddr, ttool.Opts{
        User: viper.GetString("datastorage.login"),
        Pass: viper.GetString("datastorage.password"),
    })
    if err != nil {
        log.WithFields(log.Fields{"what" : "Data Storage",}).Fatal(err)
        return
    }
    defer conn.Close()

    var pool work.WorkerPool
    pool.Start(conn, Url, sports...)
    
    // Sets handler for /ready
    http.HandleFunc("/ready", pool.ReadyHandler)
    
    HttpAddr = viper.GetString("http.address") + ":" + viper.GetString("http.port")
    go http.ListenAndServe(HttpAddr, nil) 


    // Waits for the first line synchronization
    if pool.CheckWorkersSync() {
        s := grpc.NewServer()
        srv := &grpcserv.GRPCServer{Conn: conn}

        subs.RegisterSubscribtionServer(s, srv)
        GrpcAddr = viper.GetString("grpc.address") + ":" + viper.GetString("grpc.port")
        l, err := net.Listen("tcp", GrpcAddr)
        if err != nil {
            log.Fatal(err)
        }

        if err := s.Serve(l); err != nil {
            log.Fatal(err)
        }
    }
}
