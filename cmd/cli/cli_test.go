package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	main()
	defer func(t *testing.T) {
		if r := recover(); r != nil {
			t.Error("game run failed, panic")
		}
	}(t)

	if game.Winner() == "" {
		t.Error("no winner game failed")
	}
}
