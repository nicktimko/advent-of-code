// Here are some example programs that use these features:
package intcode_test

import (
	"strings"
	"testing"

	"github.com/nicktimko/aoc-2019-golang/intcode"
)

func TestCrashNiceBadOp(t *testing.T) {
	badOpTape := []int{88}

	proc := intcode.New(badOpTape, []int{})

	for proc.Running() {
		proc.ProcessInstruction()
	}
	crashed, reason := proc.Crashed()
	if !crashed {
		t.Error("processor didn't crash")
	}
	if !strings.HasPrefix(reason, "unknown op") {
		t.Errorf("processor crashed in an unexpected way: '%s'", reason)
	}

}

func TestCrashNicePointerLost(t *testing.T) {
	pointerLostTape := []int{1, 0, 0, 3}

	proc := intcode.New(pointerLostTape, []int{})

	for proc.Running() {
		proc.ProcessInstruction()
	}
	crashed, reason := proc.Crashed()
	if !crashed {
		t.Error("processor didn't crash")
	}
	if !strings.HasPrefix(reason, "pointer left tape") {
		t.Errorf("processor crashed in an unexpected way: '%s'", reason)
	}
}

func TestCrashNiceBadAddressing(t *testing.T) {
	badAddressingTape := []int{9901, 0, 0, 3, 99}

	proc := intcode.New(badAddressingTape, []int{})

	for proc.Running() {
		proc.ProcessInstruction()
	}
	crashed, reason := proc.Crashed()
	if !crashed {
		t.Error("processor didn't crash")
	}
	if !strings.HasPrefix(reason, "error in parameter") {
		t.Errorf("processor crashed in an unexpected way: '%s'", reason)
	}
}
