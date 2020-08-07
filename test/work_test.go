package test

import (
	"github.com/myjupyter/softpro/work"
	"testing"
)

func TestWriteWorkerState(t *testing.T) {
	not_expected := work.WorkerState{}
	endpoint := work.WorkerState{}

	endpoint.WriteWorkerState(nil, "1", true, 0)
	if not_expected == endpoint {
		t.Fatalf("Wrong working WriteWorkerState, endpoint %v shouldn't be equal to %v", endpoint, not_expected)
	}
}

func TestIsSync(t *testing.T) {
	endpoint := work.WorkerState{}

	endpoint.WriteWorkerState(nil, "1", true, 0)
	if !endpoint.IsSync() {
		t.Fatal("Failed IsSyns test, state should be synchronized")
	}
}
