package day10

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/MahammadAgayev/advent-of-code2025/common"
)

type Machine struct {
	lights []bool
	wires []wire
	joltage []int
}

type wire []int

func (m *Machine) newWire() {
	wire := make([]int, 0, 1)
	m.wires = append(m.wires, wire)
}

func (m *Machine) addWire(num int) {
	wire := m.wires[len(m.wires)-1]
	wire = append(wire, num)
	m.wires[len(m.wires) -1] = wire
}

func (m *Machine) print() {
	for _, sg := range m.lights {
		fmt.Printf("%v,", sg)
	}

	fmt.Println("\nwires...")
	for _, bt := range m.wires {
		fmt.Println(bt)
	}
}

type state struct {
	lights []bool
	wireId int
	depth int
}

type q struct {
	arr []state
}

func newq() q {
	return q {
		arr:  make([]state, 0, 32),
	}
}

func (q *q) enq(a state)  {
	q.arr = append(q.arr, a)
}

func (q *q) deq() (bool, state) {
	if len(q.arr) == 0 {
		return true, state{}
	}
	a := q.arr[0]
	q.arr = q.arr[1:]

	return false, a
}


const (
	lightStart = '['
	lightStop = ']'

	wireStart = '('
	wireStop = ')'

	joltageStart = '{'
	joltageStop = '}'
)

func Main() error {
	machines, err := readInput()

	if err != nil {
		return err
	}

	sum := 0
	for _, m := range machines {
		sum += howManyWires(m)
	}
	fmt.Printf("Part 1: %d\n", sum)

	return nil
}

func howManyWires(m *Machine) int {
	min_ := 99999999
	for i := range m.wires {
		ln := bfs(m, i)
		if ln < min_ {
			min_ = ln
		}
	}

	return min_
}

func readInput() ([]*Machine, error) {
	scanner := common.ReadScanner()
	machines := make([]*Machine, 0)

	for scanner.Scan() {
		line := scanner.Text()
		m, err := parseLine(line)

		if err != nil {
			return nil, err
		}

		machines = append(machines, m)
	}

	return machines, nil
}


func parseLine(line string) (*Machine, error) {
	machine := &Machine{
		lights: make([]bool, 0, 1),
		wires: make([]wire, 0),
		joltage: make([]int, 0),
	}

	lightParse := false
	wireParse := false
	joltageParse := false
	numBuffer := ""

	for _, ch := range line {
		switch ch {
		case lightStart:
			lightParse = true
			continue
		case lightStop:
			lightParse = false
			continue
		case wireStart:
			machine.newWire()
			wireParse = true
			continue
		case wireStop:
			wireParse = false
			continue
		case joltageStart:
			joltageParse = true
			continue
		case joltageStop:
			if numBuffer != "" {
				num, err := strconv.Atoi(numBuffer)
				if err != nil {
					return nil, err
				}
				machine.joltage = append(machine.joltage, num)
				numBuffer = ""
			}
			joltageParse = false
			continue
		case ',':
			if joltageParse && numBuffer != "" {
				num, err := strconv.Atoi(numBuffer)
				if err != nil {
					return nil, err
				}
				machine.joltage = append(machine.joltage, num)
				numBuffer = ""
			}
		    continue
		case ' ':
			continue
		}

		if lightParse {
		   lightOn := false
		   if ch == '#' {
		       lightOn = true
		   }
		   machine.lights = append(machine.lights, lightOn)
		}

		if wireParse {
			num, err := strconv.Atoi(string(ch))

			if err != nil {
				return nil, err
			}

			machine.addWire(num)
		}

		if joltageParse {
			numBuffer += string(ch)
		}
	}

	return machine, nil
}


func bfs(m *Machine, wireId int) int {
	q := newq()
	startLights := make([]bool, len(m.lights))
	q.enq(state{wireId: wireId, lights: startLights, depth: 1})

	for true {
		isEmpty, qVal := q.deq();

		if isEmpty {
			break
		}

		nextWireId := qVal.wireId
		lights := qVal.lights


		// fmt.Println("before pressing wire", m.wires[nextWireId], lights)
		lights = press(lights, m.wires[nextWireId])
		// fmt.Println("pressed wire", m.wires[nextWireId], lights)

		if lightEqual(lights, m.lights) {
			return qVal.depth
		}

		for i := range m.wires {
			if i == nextWireId {
				continue
			}
			newLights := make([]bool, len(m.lights))
			copy(newLights, lights)
			q.enq(state{wireId: i, lights: newLights, depth: qVal.depth+1})
		}

	}

	return 9999999
}


// oke, we do a BFS, but since it's kinda
func press(state []bool, wire wire) []bool {
	for _, nm := range wire {
		state[nm] = !state[nm]
	}

	return state
}

func lightEqual(light1 []bool, light2 []bool) bool {
	return reflect.DeepEqual(light1, light2)
}
