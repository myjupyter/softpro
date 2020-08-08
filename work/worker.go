package work

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/tarantool/go-tarantool"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type WorkerState struct {
	LatestSync string
	Status     bool
	WorkerID   uint

	err error
}

func (ws *WorkerState) GetError() error {
	return ws.err
}

func (ws *WorkerState) IsSync() bool {
	return ws.Status
}

func (ws *WorkerState) WriteWorkerState(err error, LatestSync string, Status bool, id uint) {
	*ws = WorkerState{err: err, LatestSync: LatestSync, Status: Status, WorkerID: id}
}

func (ws *WorkerState) GetJSONInfo() string {
	info := map[string]interface{}{
		"latest_sync": ws.LatestSync,
		"is_working":  ws.Status,
		"worker_id":   ws.WorkerID,
	}

	wsJSON, _ := json.Marshal(info)
	return string(wsJSON)
}

type SportSubs struct {
	Sport   string
	Seconds int
}

type Worker struct {
	ID   uint
	URL  string
	Subs SportSubs
	Conn *tarantool.Connection
}

func (w *Worker) Run(workerState *WorkerState) {
	workerLog := log.WithFields(log.Fields{"worker": w.ID})
	URL := w.URL + w.Subs.Sport
	SPORT := strings.ToUpper(w.Subs.Sport)
	workerLog.Info("Started working at purpose " + w.Subs.Sport)

	for {
		// Makes request
		resp, err := http.Get(URL)
		if err != nil {
			workerLog.Error(err)
			workerState.WriteWorkerState(err, time.Now().String(), false, w.ID)
			return
		}

		if resp.StatusCode != 200 {
			workerLog.Errorf("HTTP %s", resp.Status)
			workerState.WriteWorkerState(err, time.Now().String(), false, w.ID)
			return
		}
		defer resp.Body.Close()

		// Exstracts body from response
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			workerLog.Error(err)
			workerState.WriteWorkerState(err, time.Now().String(), false, w.ID)
			return
		}

		// Json Parsing
		var dat map[string]interface{}
		if err := json.Unmarshal(body, &dat); err != nil {
			workerLog.Error(err)
			workerState.WriteWorkerState(err, time.Now().String(), false, w.ID)
			return
		}

		sport := dat["lines"].(map[string]interface{})
		value, err := strconv.ParseFloat(sport[SPORT].(string), 64)
		if err != nil {
			workerLog.Error(err)
			workerState.WriteWorkerState(err, time.Now().String(), false, w.ID)
			return
		}

		// DataBase Insertion
		info, err := w.Conn.Insert(w.Subs.Sport, []interface{}{nil, value, uint64(time.Now().Unix())})
		if err != nil {
			log.WithFields(log.Fields{
				"what":   "insertion to Data Storage",
				"worker": w.ID,
				"code":   info.Code,
				"data":   info.Data,
			}).Error(err)
			workerState.WriteWorkerState(err, time.Now().String(), false, w.ID)
			return
		}
		workerState.WriteWorkerState(nil, time.Now().String(), true, w.ID)

		time.Sleep(time.Duration(w.Subs.Seconds) * time.Second)
	}
}

func (w *Worker) Start(workerState *WorkerState) {
	go w.Run(workerState)
}
