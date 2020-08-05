package work

import (
    "io/ioutil"
    "time"
    "net/http"
    "encoding/json"
    "strings"
    "strconv"
    "github.com/tarantool/go-tarantool"
    log "github.com/sirupsen/logrus"
)

type SportSubs struct {
    Sport string
    Seconds int 
}

type Worker struct {
    ID uint
    URL string
    Subs SportSubs
    Conn *tarantool.Connection

    Status chan interface{} 
    CheckStat chan bool
}

func (w *Worker) Run() {
    worker_log := log.WithFields(log.Fields{"worker" : w.ID})
    URL := w.URL + w.Subs.Sport
    SPORT := strings.ToUpper(w.Subs.Sport)
    worker_log.Info("Started working")
    for {
        // Makes request
        resp, err := http.Get(URL)
        if err != nil {
            worker_log.Error(err)
            w.Status <- err
            return
        }
        defer resp.Body.Close()

        // Exstracts body from response
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            worker_log.Error(err)
            w.Status <- err
            return
        }
        
        // Json Parsing
        var dat map[string]interface{}
        if err := json.Unmarshal(body, &dat); err != nil {
            worker_log.Error(err)
            w.Status <- err
            return
        }
        sport := dat["lines"].(map[string]interface{})
        value, err := strconv.ParseFloat(sport[SPORT].(string), 64)

        // DataBase Insertion
        info, err := w.Conn.Insert(w.Subs.Sport, []interface{}{nil, value})
        if err != nil {
            log.WithFields(log.Fields{
                "worker" : w.ID,
                "code" : info.Code,
                "data" : info.Data,
            }).Error(err)
            w.Status <- err
            return
        }

        //TODO: храить информацию в некотором информаторе, чтобы не ожидать
        //синхронизации линий, а давать предпоследнюю информацию
        select {
            case <- w.CheckStat:
                log.Info("Stat")
                w.Status <- nil
            default: 
        }

        time.Sleep(time.Duration(w.Subs.Seconds) * time.Second)
    }
}

func (w *Worker) Start() {
    go w.Run()
}



