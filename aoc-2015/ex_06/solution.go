package ex_06

import (
	"aoc_2015/utils"
	"fmt"
	"strconv"
	"strings"
)

const (
	height             = 1000
	width              = 1000
	MessageTypeTurnOn  = "turn on"
	MessageTypeTurnOff = "turn off"
	MessageTypeToggle  = "toggle"
)

type Pos struct {
	x int
	y int
}

func parsePos(s string) Pos {
	xy := strings.Split(s, ",")
	x, _ := strconv.Atoi(xy[0])
	y, _ := strconv.Atoi(xy[1])
	return Pos{x, y}
}

type MessageType string
type Message struct {
	typ  MessageType
	from Pos
	to   Pos
}

func parseMessage(s string) Message {
	var typ MessageType
	var rest string

	for _, t := range []MessageType{MessageTypeToggle, MessageTypeTurnOn, MessageTypeTurnOff} {
		if s[:len(t)] == string(t) {
			typ = t
			rest = s[len(t)+1:]
		}
	}

	fromTo := strings.Split(rest, " through ")
	return Message{typ, parsePos(fromTo[0]), parsePos(fromTo[1])}
}

type Light struct {
	x          int
	y          int
	brightness int
}

type Lights struct {
	arr [][]*Light
}

func NewLights() *Lights {
	lights := make([][]*Light, width)
	for index := range lights {
		lights[index] = make([]*Light, height)
	}

	for i := range height {
		for j := range width {
			lights[i][j] = &Light{i, j, 0}
		}
	}

	return &Lights{lights}
}

func (l *Lights) turnOn(from Pos, to Pos) {
	for i := from.x; i <= to.x; i++ {
		for j := from.y; j <= to.y; j++ {
			l.arr[i][j].brightness += 1
		}
	}
}

func (l *Lights) turnOff(from Pos, to Pos) {
	for i := from.x; i <= to.x; i++ {
		for j := from.y; j <= to.y; j++ {
			brightness := l.arr[i][j].brightness
			l.arr[i][j].brightness = max(brightness-1, 0)
		}
	}
}

func (l *Lights) toggle(from Pos, to Pos) {
	for i := from.x; i <= to.x; i++ {
		for j := from.y; j <= to.y; j++ {
			l.arr[i][j].brightness += 2
		}
	}
}

func (l *Lights) countBrightness() int {
	total := 0

	for _, row := range l.arr {
		for _, light := range row {
			total += light.brightness
		}
	}

	return total
}

func Solve() {
	data := utils.LoadExercise("06")
	fmt.Println(solve06(data))
}

func solve06(data string) int {
	lights := NewLights()

	for _, line := range strings.Split(data, "\n") {
		message := parseMessage(line)
		switch message.typ {
		case MessageTypeTurnOn:
			lights.turnOn(message.from, message.to)
		case MessageTypeTurnOff:
			lights.turnOff(message.from, message.to)
		case MessageTypeToggle:
			lights.toggle(message.from, message.to)
		}
	}

	return lights.countBrightness()
}
