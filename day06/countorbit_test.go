package day06_test

import (
	"testing"

	"github.com/nicktimko/aoc-2019-golang/aocio"
	"github.com/nicktimko/aoc-2019-golang/day06"
)

type countOrbitTrial struct {
	input    map[string]string
	expected int
}

var demoGraph = map[string]string{
	/*
		|          G - H       J - K - L
		|         /           /
		|  COM - B - C - D - E - F
		|                 \
		|                  I
	*/
	"B": "COM",
	"C": "B",
	"D": "C",
	"E": "D",
	"F": "E",
	"G": "B",
	"H": "G",
	"I": "D",
	"J": "E",
	"K": "J",
	"L": "K",
}

func TestCounter(t *testing.T) {
	trials := []countOrbitTrial{
		{map[string]string{"B": "COM"}, 1},
		{map[string]string{"B": "COM", "C": "B"}, 3},
		{map[string]string{"B": "COM", "C": "COM"}, 2},
		{demoGraph, 42},
	}
	for n, trial := range trials {
		result := day06.CountOrbits(trial.input)
		if result != trial.expected {
			t.Errorf("countOrbits(trials[%d].input) -> %d != %d", n, result, trial.expected)
		}
	}
}

type getParentsTrial struct {
	inputGraph  map[string]string
	inputOrigin string
	expected    []string
}

func TestGetParents(t *testing.T) {
	trials := []getParentsTrial{
		{map[string]string{"B": "COM", "C": "B"}, "B", []string{"COM", "B"}},
		{map[string]string{"B": "COM", "C": "B"}, "C", []string{"COM", "B", "C"}},
		{demoGraph, "COM", []string{"COM"}},
		{demoGraph, "D", []string{"COM", "B", "C", "D"}},
		{demoGraph, "I", []string{"COM", "B", "C", "D", "I"}},
		{demoGraph, "K", []string{"COM", "B", "C", "D", "E", "J", "K"}},
	}
	for n, trial := range trials {
		result := day06.AllParents(trial.inputGraph, trial.inputOrigin)
		if !aocio.EqStringSlice(result, trial.expected) {
			t.Errorf("allParents(trials[%d]) -> %v != %v", n, result, trial.expected)
		}
	}
}

type getLatestParentTrial struct {
	inputGraph map[string]string
	inputN1    string
	inputN2    string
	expected   string
}

func TestLatestCommonParent(t *testing.T) {
	trials := []getLatestParentTrial{
		{map[string]string{"B": "COM", "C": "COM"}, "B", "C", "COM"},
		{map[string]string{"B": "COM", "C": "B"}, "B", "C", "B"},
		{demoGraph, "H", "C", "B"},
		{demoGraph, "C", "H", "B"},
		{demoGraph, "L", "I", "D"},
	}
	for n, trial := range trials {
		result := day06.LatestCommonParent(trial.inputGraph, trial.inputN1, trial.inputN2)
		if result != trial.expected {
			t.Errorf("LatestCommonParent(trials[%d]) -> %v != %v", n, result, trial.expected)
		}
	}
}

type getDistanceTrial struct {
	inputGraph map[string]string
	inputN1    string
	inputN2    string
	expected   int
}

func TestDistance(t *testing.T) {
	demoGraphMod := map[string]string{
		"YOU": "K",
		"SAN": "I",
	}
	for k, v := range demoGraph {
		demoGraphMod[k] = v
	}

	trials := []getDistanceTrial{
		{map[string]string{"B": "COM", "C": "COM"}, "B", "C", 2},
		{map[string]string{"B": "COM", "C": "B"}, "B", "C", 1},
		{demoGraph, "H", "C", 3},
		{demoGraph, "C", "H", 3},
		{demoGraph, "L", "I", 5},
		{demoGraphMod, demoGraphMod["YOU"], demoGraphMod["SAN"], 4},
	}
	for n, trial := range trials {
		result := day06.Distance(trial.inputGraph, trial.inputN1, trial.inputN2)
		if result != trial.expected {
			t.Errorf("Distance(trials[%d]) -> %v != %v", n, result, trial.expected)
		}
	}
}
