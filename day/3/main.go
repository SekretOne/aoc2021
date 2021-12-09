package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

type rating struct {
	String  string
	BitSize int
	Value   int
}

func readInput(fileName string) []string {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	content := string(b)
	return strings.Split(content, "\r\n")
}

// calculate gamma rating from finding the most common bit in each entry
func gamma(report []string) rating {
	g := ""

	bitSize := len(report[0])
	for i := 0; i < bitSize; i++ {
		ones := 0
		zeroes := 0

		for _, e := range report {
			if e[i:i+1] == "1" {
				ones += 1
			} else {
				zeroes += 1
			}
		}

		var sigBit string
		if ones > zeroes {
			sigBit = "1"
		} else {
			sigBit = "0"
		}

		g += sigBit
	}

	val, err := strconv.ParseInt(g, 2, 0)

	if err != nil {
		log.Fatal(err)
	}

	return rating{
		String:  g,
		BitSize: len(g),
		Value:   int(val),
	}
}

// calculate the epsilon rating by inverting the gamma
func epsilon(g rating) rating {
	mask := int(math.Pow(2, float64(len(g.String)))) - 1

	val := g.Value ^ mask

	epBits := strconv.FormatInt(int64(val), 2)
	padding := fmt.Sprintf("%d", g.BitSize)

	return rating{
		String:  fmt.Sprintf("%0"+padding+"v", epBits),
		BitSize: g.BitSize,
		Value:   val,
	}
}

func o2Gen(report []string) rating {
	return reduce(report, 0, true)
}

func co2Scrub(report []string) rating {
	return reduce(report, 0, false)
}

// recursive reducer that filters down the most or least popular
func reduce(r []string, bitPos int, useCommon bool) rating {
	bitSize := len(r[0])

	if bitPos == bitSize {
		val, err := strconv.ParseInt(r[0], 2, 0)

		if err != nil {
			log.Fatalf("unable to parse %v into an int, %v", r[0], err)
		}

		return rating{
			String:  r[0],
			BitSize: bitSize,
			Value:   int(val),
		}
	}

	ones := make([]string, 0, len(r))
	zeroes := make([]string, 0, len(r))

	for _, e := range r {
		if e[bitPos:bitPos+1] == "1" {
			ones = append(ones, e)
		} else {
			zeroes = append(zeroes, e)
		}
	}

	if len(ones)-len(zeroes) >= 0 == useCommon && len(ones) > 0 || len(zeroes) == 0 {
		return reduce(ones, bitPos+1, useCommon)
	} else {
		return reduce(zeroes, bitPos+1, useCommon)
	}
}

func main() {
	report := readInput("day/3/input.txt")

	g := gamma(report)
	e := epsilon(g)
	fmt.Printf("gamma: %+v\n", g)
	fmt.Printf("epsilon: %+v\n", e)

	fmt.Println(g.Value * e.Value)
	// -- day 2
	o2 := o2Gen(report)
	co2 := co2Scrub(report)

	fmt.Printf("o2: %+v\n", o2)
	fmt.Printf("co2: %+v\n", co2)

	fmt.Println(o2.Value * co2.Value)
}
