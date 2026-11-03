package nats

import (
	"testing"
	"time"
)

func TestUpdateRun(t *testing.T) {
	var none *UpdateRun
	if none.Running() {
		t.Error("a nil run must not be running")
	}
	if u, d, f := none.Counts(); u+d+f != 0 {
		t.Error("a nil run has no results")
	}

	run := &UpdateRun{Started: time.Now(), Results: []UpdateResult{{}, {UpToDate: true}, {Failed: true}, {}}}
	if !run.Running() {
		t.Error("a started run without finish time is running")
	}
	if u, d, f := run.Counts(); u != 2 || d != 1 || f != 1 {
		t.Errorf("unexpected counts %d updated, %d up to date, %d failed", u, d, f)
	}

	run.Finished = time.Now()
	if run.Running() {
		t.Error("a finished run is not running")
	}
}
