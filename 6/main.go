package main

import (
	"fmt"
	"strings"

	"github.com/jmcshane/advent-of-code/pkg/input"
)

func main() {
	content, err := input.ReadInputFile(6)
	if err != nil {
		panic(err)
	}
	orbits := strings.Split(content, "\n")
	CountNodes(orbits)
}

func CountNodes(orbits []string) int {
	orbitMap := map[string]Node{}
	for _, orbit := range orbits {
		data := strings.Split(orbit, ")")
		if len(data) < 2 {
			continue
		}
		parentNode := data[0]
		childNode := data[1]
		val, ok := orbitMap[parentNode]
		if !ok {
			val = Node{ParentNodeName: "", Name: parentNode}
			orbitMap[parentNode] = val
		}
		child, childOK := orbitMap[childNode]
		if !childOK {
			child = Node{ParentNodeName: parentNode, Name: childNode}
			orbitMap[childNode] = child
		} else {
			child.ParentNodeName = parentNode
			orbitMap[childNode] = child
		}
	}
	nodes := make([]Node, 0, len(orbitMap))

	for _, v := range orbitMap {
		nodes = append(nodes, v)
	}

	// fmt.Printf("nodes are %+v\n", nodes)
	// value := countParents(nodes, orbitMap)
	// fmt.Printf("Total parent nodes %d\n", value)
	you, _ := orbitMap["YOU"]
	countParentForNode(you, orbitMap)
	san, _ := orbitMap["SAN"]
	countParentForNode(san, orbitMap)
	return 0
}

func countParents(nodes []Node, nodeMap map[string]Node) int {
	sum := 0
	for _, node := range nodes {
		sum += countParentForNode(node, nodeMap)
	}
	return sum
}

func countParentForNode(node Node, nodeMap map[string]Node) int {
	if node.ParentNodeName == "" {
		fmt.Printf("Terminating on node %s\n", node.Name)
		return 0
	}
	fmt.Printf("Node %s orbits node %s, adding orbit\n", node.Name, node.ParentNodeName)
	ParentNode, _ := nodeMap[node.ParentNodeName]
	return countParentForNode(ParentNode, nodeMap) + 1
}

type Node struct {
	ParentNodeName string
	Name           string
}
