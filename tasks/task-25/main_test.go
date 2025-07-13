package main

import (
	"testing"
	"time"
)

func TestMySleep(t *testing.T) {
	start := time.Now()

	MySleep(500 * time.Millisecond)

	elapsed := time.Since(start)

	if elapsed < 500*time.Millisecond {
		t.Errorf("MySleep slept too short: %v", elapsed)
	}

	if elapsed > 600*time.Millisecond {
		t.Errorf("MySleep slept too long: %v", elapsed)
	}
}
