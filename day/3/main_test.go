package main

import "testing"

func Test_gamma(t *testing.T) {
	report := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	actual := gamma(report)
	expected := "10110"

	if actual.String != expected {
		t.Errorf("gamma calcuation failed: expected: %v, actual: %v", expected, actual)
	}
}
