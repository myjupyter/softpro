package work

import (
	"encoding/json"
	"fmt"
	"github.com/tarantool/go-tarantool"
	"net/http"
	"time"
)

type WorkerPool struct {
	Workers       []Worker
	WorkersStates []*WorkerState
}

func (wp *WorkerPool) Start(conn *tarantool.Connection, url string, sports ...SportSubs) {
	for i, sport := range sports {
		wp.Workers = append(wp.Workers, Worker{
			ID:   uint(i),
			URL:  url,
			Subs: sport,
			Conn: conn,
		})
		wp.WorkersStates = append(wp.WorkersStates, new(WorkerState))
		wp.Workers[i].Start(wp.WorkersStates[i])
	}
}

func (wp *WorkerPool) CheckWorkersSync(syncTimeout float64) bool {
	var timer float64
	for i := 0; i < len(wp.WorkersStates); i++ {
		for ; timer < syncTimeout; timer += 0.10 {
			if wp.WorkersStates[i].GetError() != nil {
				return false
			}
			if wp.WorkersStates[i].IsSync() {
				break
			}
			time.Sleep(10 * time.Millisecond)
		}
		if timer > syncTimeout {
			return false
		}
	}
	return true
}

func (wp *WorkerPool) CheckWorkerStatus() bool {
	for _, workerState := range wp.WorkersStates {
		if workerState.GetError() != nil {
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

	statusJSON, _ := json.Marshal(info)
	fmt.Fprintf(w, "%s", string(statusJSON))
}
