package main

import "testing"

var (
	report = []string{
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
)

func Test_gamma(t *testing.T) {
	actual := gamma(report)
	expected := "10110"

	if actual.String != expected {
		t.Errorf("gamma calcuation failed: expected: %v, actual: %v", expected, actual)
	}
}

func Test_o2Gen(t *testing.T) {
	actual := o2Gen(report)
	expected := "10111"

	if actual.String != expected {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func Test_co2Scrub(t *testing.T) {
	actual := co2Scrub(report)
	expected := "01010"

	if actual.String != expected {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}
