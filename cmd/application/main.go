package main

import (
	"flag"
	subs "github.com/myjupyter/softpro/api/subscription"
	grpcserv "github.com/myjupyter/softpro/grpc"
	"github.com/myjupyter/softpro/work"
	log "github.com/sirupsen/logrus"
	viper "github.com/spf13/viper"
	ttool "github.com/tarantool/go-tarantool"
	grpc "google.golang.org/grpc"
	"net"
	"net/http"
)

var LogLevel = map[string]log.Level{
	"debug":   log.DebugLevel,
	"info":    log.InfoLevel,
	"warning": log.WarnLevel,
	"error":   log.ErrorLevel,
	"fatal":   log.FatalLevel,
}

func main() {
	// Reads flags
	name := flag.String("name", "config", "Configuration file name without extension")
	dir := flag.String("dir", "configs", "Configuration file directory")

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
	dbAddr := viper.GetString("datastorage.address") + ":" + viper.GetString("datastorage.port")
	conn, err := ttool.Connect(dbAddr, ttool.Opts{
		User: viper.GetString("datastorage.login"),
		Pass: viper.GetString("datastorage.password"),
	})
	if err != nil {
		log.WithFields(log.Fields{"what": "Data Storage"}).Fatal(err)
		return
	}
	defer conn.Close()

	url := viper.GetString("lineprovider.url")
	var pool work.WorkerPool
	pool.Start(conn, url, sports...)

	// Sets handler for /ready
	http.HandleFunc("/ready", pool.ReadyHandler)

	httpAddr := viper.GetString("http.address") + ":" + viper.GetString("http.port")
	go func() {
		err := http.ListenAndServe(httpAddr, nil)
		if err != nil {
			log.WithFields(log.Fields{"what": "HTTP Server"}).Fatal(err)
		}
	}()

	// Waits for the first line synchronization
	timeOut := viper.GetDuration("datastorage.timeout")
	if pool.CheckWorkersSync(timeOut) {
		log.Debug("GRPC Started")
		s := grpc.NewServer()
		srv := &grpcserv.GRPCServer{Conn: conn}

		subs.RegisterSubscribtionServer(s, srv)
		grpcAddr := viper.GetString("grpc.address") + ":" + viper.GetString("grpc.port")
		l, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			log.WithFields(log.Fields{"what": "GRPC Server"}).Fatal(err)
		}

		if err := s.Serve(l); err != nil {
			log.WithFields(log.Fields{"what": "GRPC Server"}).Fatal(err)
		}
	} else {
		log.WithFields(log.Fields{"what": "Data Storage access time out"}).Fatal(err)
	}
}
