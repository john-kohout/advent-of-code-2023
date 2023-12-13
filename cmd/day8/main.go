package main

import (
	"bufio"
	"fmt"
	"strings"

	"advent-of-code/pkg/file"
)

const (
	Left  Instruction = 'L'
	Right Instruction = 'R'
)

type Instruction rune

type Node struct {
	left  NodeID
	right NodeID
}

func newNode(s string) Node {
	var node Node
	s = strings.ReplaceAll(s, "(", "")
	s = strings.ReplaceAll(s, ")", "")
	node.left = NodeID(strings.Split(s, ", ")[0])
	node.right = NodeID(strings.Split(s, ", ")[1])
	return node
}

type NodeID string

func main() {
	scanner, f := file.NewScanner("tests/input.txt")
	defer f.Close()
	result := GetResult(scanner)
	fmt.Println(result)
}

func GetResult(scanner *bufio.Scanner) int {
	result := 0
	var results []int
	var nodeIDs []NodeID

	var instructions []Instruction
	network := make(map[NodeID]Node)
	firstLine := true
	for scanner.Scan() {
		if firstLine {
			instructions = getInstructions(scanner.Text())
			firstLine = false
			continue
		}
		if len(scanner.Text()) == 0 {
			continue
		}
		nodeIDs = append(nodeIDs, addToMap(scanner.Text(), &network)...)
	}

	for _, nodeID := range nodeIDs {
		results = append(results, TraverseMap(instructions, network, nodeID))
	}

	result = lcmSlice(results)

	return result
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func lcmSlice(numbers []int) int {
	result := numbers[0]
	for _, num := range numbers[1:] {
		result = lcm(result, num)
	}
	return result
}

func TraverseMap(instructions []Instruction, network map[NodeID]Node, nodeID NodeID) int {
	result := 0
	endReached := false
	for !endReached {
		for _, step := range instructions {
			node, _ := network[nodeID]
			switch step {
			case Left:
				nodeID = node.left
			case Right:
				nodeID = node.right
			}
			result++
			if nodeID[2] == 'Z' {
				endReached = true
				break
			}
		}
	}
	return result
}

func addToMap(line string, network *map[NodeID]Node) []NodeID {
	nodeID := NodeID(strings.Split(line, " = ")[0])
	node := newNode(strings.Split(line, " = ")[1])
	(*network)[nodeID] = node
	var startingNodes []NodeID
	if string(nodeID)[2] == 'A' {
		startingNodes = append(startingNodes, nodeID)
	}

	return startingNodes
}

func getInstructions(line string) []Instruction {
	var instructions []Instruction
	for _, r := range line {
		instructions = append(instructions, Instruction(r))
	}
	return instructions
}
