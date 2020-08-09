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

func (w *Worker) Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

type ParserError struct {
}

func (err *ParserError) Error() string {
	return "Parser Error"
}

func ParseLinesProviderData(body []byte, sportKey string) (float64, error) {
	var dat map[string]interface{}
	if err := json.Unmarshal(body, &dat); err != nil {
		return 0, err
	}
	sport, ok := dat["lines"].(map[string]interface{})
	if !ok {
		return 0, &ParserError{}
	}
	value, err := strconv.ParseFloat(sport[sportKey].(string), 64)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func (w *Worker) Run(workerState *WorkerState) {
	URL := w.URL + w.Subs.Sport
	SPORT := strings.ToUpper(w.Subs.Sport)

	log.Info("Started working at purpose " + w.Subs.Sport)
	for {
		// Makes request
		body, err := w.Get(URL)
		if err != nil {
			log.WithFields(log.Fields{
				"what":   "GET Request",
				"worker": w.ID,
			}).Error(err)
			workerState.WriteWorkerState(err, time.Now().String(), false, w.ID)
		}

		// Json Parsing
		value, err := ParseLinesProviderData(body, SPORT)
		if err != nil {
			log.WithFields(log.Fields{
				"what":   "Parser",
				"worker": w.ID,
			}).Error(err)
			workerState.WriteWorkerState(err, time.Now().String(), false, w.ID)
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
