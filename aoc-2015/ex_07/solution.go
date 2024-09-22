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

type BitOpType string
type BitOp struct {
	inputLeft  string
	typ        BitOpType
	inputRight string
	hasSignal  bool
	signal     uint16
}

func (b *BitOp) resolve(circuit map[string]*BitOp) (val uint16) {
	if b.hasSignal {
		return b.signal
	}
	var left, right uint16

	leftInt, err := strconv.Atoi(b.inputLeft)
	if err == nil {
		left = uint16(leftInt)
	} else {
		left = circuit[b.inputLeft].resolve(circuit)
	}

	rightInt, err := strconv.Atoi(b.inputRight)
	if err == nil {
		right = uint16(rightInt)
	} else {
		right = circuit[b.inputRight].resolve(circuit)
	}

	switch b.typ {
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
	b.hasSignal = true
	b.signal = val
	return val
}

func parseBitOp(s string) (string, *BitOp) {
	vals := strings.Split(s, " -> ")
	key := vals[1]
	rest := vals[0]

	for _, typ := range []BitOpType{BitOpAnd, BitOpOr, BitOpLShift, BitOpRShift} {
		substr := string(typ)
		if strings.Contains(rest, substr) {
			leftRight := strings.Split(rest, substr)
			left := strings.TrimSpace(leftRight[0])
			right := strings.TrimSpace(leftRight[1])
			return key, &BitOp{
				inputLeft:  left,
				typ:        typ,
				inputRight: right,
				hasSignal:  false,
			}
		}
	}

	if strings.Contains(rest, BitOpNot) {
		val := rest[4:]
		return key, &BitOp{
			inputLeft:  val,
			typ:        BitOpNot,
			inputRight: val,
			hasSignal:  false,
		}
	}

	return key, &BitOp{
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
	circuit := make(map[string]*BitOp)

	for _, line := range strings.Split(data, "\n") {
		key, bitOp := parseBitOp(line)
		circuit[key] = bitOp
	}

	return circuit["a"].resolve(circuit)
}

func solveSecond(data string) uint16 {
	circuit := make(map[string]*BitOp)

	for _, line := range strings.Split(data, "\n") {
		key, bitOp := parseBitOp(line)
		circuit[key] = bitOp
	}

	circuit["b"].signal = solveFirst(data)
	circuit["b"].hasSignal = true

	return circuit["a"].resolve(circuit)
}
