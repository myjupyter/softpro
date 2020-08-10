package test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/myjupyter/softpro/work"
	"net/http"
	"testing"
	"time"
)

func TestWriteWorkerState(t *testing.T) {
	expected := work.WorkerState{LatestSync: "1", Status: true, WorkerID: 0}
	endpoint := work.WorkerState{}

	endpoint.WriteWorkerState(nil, "1", true, 0)
	if expected != endpoint {
		t.Fatalf("Wrong working WriteWorkerState, endpoint %v should be equal to %v", endpoint, expected)
	}
}

func TestIsSync(t *testing.T) {
	endpoint := work.WorkerState{}

	endpoint.WriteWorkerState(nil, "1", true, 0)
	if !endpoint.IsSync() {
		t.Fatal("Failed IsSyns test, state should be synchronized")
	}
}

func TestGet(t *testing.T) {
	const (
		addr = "localhost:8000"
		text = "Some test text ..."
	)

	handler := http.DefaultServeMux
	handler.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(writer, "%s", text)
	})

	serv := http.Server{
		Addr:    addr,
		Handler: handler,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go func() {
		<-ctx.Done()
		_ = serv.Shutdown(ctx)
	}()

	go func() {
		err := serv.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			t.Error(err)
		}
	}()

	worker := work.Worker{}
	endpoint, err := worker.Get("http://" + addr + "/test")
	if err != nil {
		t.Fatal(err)
	}
	if string(endpoint) != text {
		t.Fatalf("TestGet has not been passed: wrong returned value; expected(%s) and received(%s)", text, string(endpoint))
	}

}

func TestParseLinesProviderData(t *testing.T) {

	type Req struct {
		Lines map[string]string `json:"lines"`
	}

	value := 1.34
	jsonTest, _ := json.Marshal(Req{
		Lines: map[string]string{
			"soccer": fmt.Sprintf("%f", value),
		},
	})

	endpoint, err := work.ParseLinesProviderData(jsonTest, "soccer")
	if err != nil {
		t.Fatal(err)
	}
	if endpoint != value {
		t.Fatalf("TestParseLinesProviderData has not been passed: wrong returned value; expected(%f) and received(%f)", value, endpoint)
	}
}
