package ex_07

import (
	"aoc_2015/utils"
	"fmt"
	"strconv"
	"strings"
)

const (
	BitOpAnd    = "AND"
	BitOpOr     = "OR"
	BitOpLShift = "LSHIFT"
	BitOpRShift = "RSHIFT"
	BitOpNot    = "NOT"
	BitOpNoOp   = "NOOP"
)

type BitOp string
type Wire struct {
	inputLeft  string
	typ        BitOp
	inputRight string
	hasSignal  bool
	signal     uint16
}

func (w *Wire) resolve(circuit map[string]*Wire) (val uint16) {
	if w.hasSignal {
		return w.signal
	}
	var left, right uint16

	leftInt, err := strconv.Atoi(w.inputLeft)
	if err == nil {
		left = uint16(leftInt)
	} else {
		left = circuit[w.inputLeft].resolve(circuit)
	}

	rightInt, err := strconv.Atoi(w.inputRight)
	if err == nil {
		right = uint16(rightInt)
	} else {
		right = circuit[w.inputRight].resolve(circuit)
	}

	switch w.typ {
	case BitOpAnd:
		val = left & right
	case BitOpOr:
		val = left | right
	case BitOpLShift:
		val = left << right
	case BitOpRShift:
		val = left >> right
	case BitOpNot:
		val = ^left
	case BitOpNoOp:
		val = left
	default:
		panic("what the fuck")
	}
	w.hasSignal = true
	w.signal = val
	return val
}

func parseBitOp(s string) (string, *Wire) {
	vals := strings.Split(s, " -> ")
	key := vals[1]
	rest := vals[0]

	for _, typ := range []BitOp{BitOpAnd, BitOpOr, BitOpLShift, BitOpRShift} {
		substr := string(typ)
		if strings.Contains(rest, substr) {
			leftRight := strings.Split(rest, substr)
			left := strings.TrimSpace(leftRight[0])
			right := strings.TrimSpace(leftRight[1])
			return key, &Wire{
				inputLeft:  left,
				typ:        typ,
				inputRight: right,
				hasSignal:  false,
			}
		}
	}

	if strings.Contains(rest, BitOpNot) {
		val := rest[4:]
		return key, &Wire{
			inputLeft:  val,
			typ:        BitOpNot,
			inputRight: val,
			hasSignal:  false,
		}
	}

	return key, &Wire{
		inputLeft:  rest,
		typ:        BitOpNoOp,
		inputRight: rest,
		hasSignal:  false,
	}
}

func Solve() {
	data := utils.LoadExercise("07")
	fmt.Println(solveFirst(data))
	fmt.Println(solveSecond(data))
}

func solveFirst(data string) uint16 {
	circuit := make(map[string]*Wire)

	for _, line := range strings.Split(data, "\n") {
		key, bitOp := parseBitOp(line)
		circuit[key] = bitOp
	}

	return circuit["a"].resolve(circuit)
}

func solveSecond(data string) uint16 {
	circuit := make(map[string]*Wire)

	for _, line := range strings.Split(data, "\n") {
		key, bitOp := parseBitOp(line)
		circuit[key] = bitOp
	}

	circuit["b"].signal = solveFirst(data)
	circuit["b"].hasSignal = true

	return circuit["a"].resolve(circuit)
}
