package mypackage_test

import (
	"testing"

	"github.com/Tenjin0/go_crash_course/21_testing/mypackage"
)

func TestCount(t *testing.T) {

	if mypackage.Count() != 1 {
		t.Error("expected 1")
	}
	if mypackage.Count() != 2 {
		t.Error("expected 2")
	}
	if mypackage.Count() != 3 {
		t.Error("expected 3")
	}
}
