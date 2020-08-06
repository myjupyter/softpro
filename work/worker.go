package work

import (
    "io/ioutil"
    "time"
    "net/http"
    "encoding/json"
    "strings"
    "strconv"
//    "sync"
    "github.com/tarantool/go-tarantool"
    log "github.com/sirupsen/logrus"
    //"fmt"
)

type WorkerState struct {
    latest_sync string  
    status bool        
    worker_id uint         

    err error
//    mutex sync.Mutex
}

func (ws *WorkerState) GetError() error {
    return ws.err
}

func (ws *WorkerState) IsSync() bool {
    return ws.status
}

func (ws *WorkerState) WriteWorkerState(err error, latest_sync string, status bool, id uint) {
   // ws.mutex.Lock()1
    *ws = WorkerState{err: err, latest_sync: latest_sync, status: status, worker_id: id}
    //  ws.mutex.Unlock()
}

func (ws *WorkerState) GetJSONInfo() string {
   // ws.mutex.Lock()
   info := map[string]interface{}{
        "latest_sync" : ws.latest_sync,
        "status"      : ws.status,
        "worker_id"   : ws.worker_id,
    }

    ws_json, _ := json.Marshal(info)
    // ws.mutex.Unlock()
    return string(ws_json)
}

type SportSubs struct {
    Sport string
    Seconds int 
}

type Worker struct {
    ID uint
    URL string
    Subs SportSubs
    Conn *tarantool.Connection
}

func (w *Worker) Run(worker_state *WorkerState) {
    worker_log := log.WithFields(log.Fields{"worker" : w.ID})
    URL := w.URL + w.Subs.Sport
    SPORT := strings.ToUpper(w.Subs.Sport)
    worker_log.Info("Started working")
    
    for {
        // Makes request
        resp, err := http.Get(URL)
        if err != nil {
            worker_log.Error(err)
            worker_state.WriteWorkerState(err, time.Now().String(), false, w.ID)
            return
        }
        defer resp.Body.Close()

        // Exstracts body from response
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            worker_log.Error(err)
            worker_state.WriteWorkerState(err, time.Now().String(), false, w.ID)
            return
        }
        
        // Json Parsing
        var dat map[string]interface{}
        if err := json.Unmarshal(body, &dat); err != nil {
            worker_log.Error(err)
            worker_state.WriteWorkerState(err, time.Now().String(), false, w.ID)
            return
        }
        sport := dat["lines"].(map[string]interface{})
        value, err := strconv.ParseFloat(sport[SPORT].(string), 64)

        // DataBase Insertion
        info, err := w.Conn.Insert(w.Subs.Sport, []interface{}{nil, value, uint64(time.Now().Unix())})
        if err != nil {
            log.WithFields(log.Fields{
                "worker" : w.ID,
                "code" : info.Code,
                "data" : info.Data,
            }).Error(err)
            worker_state.WriteWorkerState(err, time.Now().String(), false, w.ID)
            return
        }
        worker_state.WriteWorkerState(nil, time.Now().String(), true, w.ID)
        
        time.Sleep(time.Duration(w.Subs.Seconds) * time.Second)
    }
}

func (w *Worker) Start(worker_state *WorkerState) {
    go w.Run(worker_state)
}



