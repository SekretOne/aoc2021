package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

type rate struct {
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

func gamma(report []string) rate {
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

	return rate{
		String:  g,
		BitSize: len(g),
		Value:   int(val),
	}
}

func epsilon(g rate) rate {
	mask := int(math.Pow(2, float64(len(g.String)))) - 1

	val := g.Value ^ mask

	epBits := strconv.FormatInt(int64(val), 2)
	padding := fmt.Sprintf("%d", g.BitSize)

	return rate{
		String:  fmt.Sprintf("%0"+padding+"v", epBits),
		BitSize: g.BitSize,
		Value:   val,
	}
}

func main() {
	report := readInput("day/3/input.txt")

	g := gamma(report)
	e := epsilon(g)
	fmt.Printf("gamma: %+v\n", g)
	fmt.Printf("epsilon: %+v\n", e)

	fmt.Println(g.Value * e.Value)
}
