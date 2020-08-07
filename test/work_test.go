package test

import (
	"github.com/myjupyter/softpro/work"
	"testing"
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
