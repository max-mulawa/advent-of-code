package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Inst rune

const (
	R Inst = 'R'
	L Inst = 'L'
)

func NewInstructions(instructions string) *Instructions {
	i := &Instructions{
		instructionSet: instructions,
		index:          -1,
	}
	return i
}

type Instructions struct {
	instructionSet string
	index          int
}

func (ins *Instructions) GetNext() Inst {
	ins.index = (ins.index + 1) % len(ins.instructionSet)
	for i, r := range ins.instructionSet {
		if i == ins.index {
			return Inst(r)
		}
	}
	log.Fatalf("failed to parse instruction set: %s", ins.instructionSet)
	return -1
}

type MapNode struct {
	Label string
	L     *MapNode
	R     *MapNode
}

func (n *MapNode) StopNode() bool {
	return n.Label == "ZZZ"
}

type MultiNode struct {
	nodes []*MapNode
}

func NewMultiNodeWithStart(m map[string]*MapNode) *MultiNode {
	nodes := make([]*MapNode, 0)

	for _, n := range m {
		if strings.HasSuffix(n.Label, "A") {
			nodes = append(nodes, n)
		}
	}

	return &MultiNode{
		nodes: nodes,
	}
}

func (n *MultiNode) StopNode() bool {
	for _, n := range n.nodes {
		if !strings.HasSuffix(n.Label, "Z") {
			return false
		} else {
			fmt.Println("node ", n.Label, " ends with Z")
		}
	}

	return true
}

func (n *MultiNode) String() string {
	display := "#######################\n"
	for _, n := range n.nodes {
		display += fmt.Sprintf("(%s) with L: (%s) and R: (%s) \n", n.Label, n.L.Label, n.R.Label)
	}

	return display
}

func (n *MultiNode) Instruction(ins Inst) *MultiNode {
	nodes := make([]*MapNode, 0)
	for _, n := range n.nodes {
		switch ins {
		case 'L':
			nodes = append(nodes, n.L)
		case 'R':
			nodes = append(nodes, n.R)
		}
	}

	return &MultiNode{nodes: nodes}
}

func main() {
	f, err := os.OpenFile("test2.txt", os.O_RDONLY, 0x664)
	//f, err := os.OpenFile("test.txt", os.O_RDONLY, 0x664)
	//f, err := os.OpenFile("base.txt", os.O_RDONLY, 0x664)
	//f, err := os.OpenFile("base2.txt", os.O_RDONLY, 0x664)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}
	defer f.Close()

	total := 0
	var instructionSet *Instructions
	s := bufio.NewScanner(f)

	mapNodes := make(map[string]*MapNode)
	var rootNode *MapNode

	for s.Scan() {
		l := strings.Trim(s.Text(), " ")
		if l == "" {
			continue
		}

		if instructionSet == nil {
			instructionSet = NewInstructions(l)
			continue
		} else {
			mapNodeTokens := removeEmpty(strings.Split(l, "="))
			sourceLabel := strings.Trim(mapNodeTokens[0], " ")

			destinations := strings.TrimRight(strings.TrimLeft(strings.Trim(mapNodeTokens[1], " "), "("), ")")
			lrNodes := removeEmpty(strings.Split(destinations, ","))
			leftNodeLabel := strings.Trim(lrNodes[0], " ")
			rightNodeLabel := strings.Trim(lrNodes[1], " ")

			sourceNode := getNode(mapNodes, sourceLabel)

			sourceNode.L = getNode(mapNodes, leftNodeLabel)
			sourceNode.R = getNode(mapNodes, rightNodeLabel)

			if rootNode == nil {
				rootNode = sourceNode
			}
		}

	}

	// Part 1
	// start := mapNodes["AAA"]
	// for {
	// 	inst := instructionSet.GetNext()
	// 	switch inst {
	// 	case L:
	// 		start = start.L
	// 	case R:
	// 		start = start.R
	// 	}
	// 	total++
	// 	if start.StopNode() {
	// 		break
	// 	}
	// }

	start := NewMultiNodeWithStart(mapNodes)
	for {
		//fmt.Println(start)
		inst := instructionSet.GetNext()
		fmt.Printf("move: %c\n", inst)
		start = start.Instruction(inst)
		total++
		if start.StopNode() {
			break
		}

		//time.Sleep(2 * time.Second)
	}
	fmt.Println(total)

}

func removeEmpty(a []string) []string {
	var b []string

	for _, s := range a {
		if strings.Trim(s, "") != "" {
			b = append(b, strings.Trim(s, ""))
		}
	}

	return b
}

func getNode(mapNodes map[string]*MapNode, label string) *MapNode {
	node, ok := mapNodes[label]
	if !ok {
		node = &MapNode{
			Label: label,
		}
		mapNodes[label] = node
	}

	return node
}
