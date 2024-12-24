package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type Graph struct {
	edges map[string]map[string]struct{}
}

func newGraph() *Graph {
	return &Graph{
		edges: make(map[string]map[string]struct{}),
	}
}

func (g *Graph) addEdge(u, v string) {
	if g.edges[u] == nil {
		g.edges[u] = make(map[string]struct{})
	}
	if g.edges[v] == nil {
		g.edges[v] = make(map[string]struct{})
	}
	g.edges[u][v] = struct{}{}
	g.edges[v][u] = struct{}{}
}

func (g *Graph) getNeighbors(u string) map[string]struct{} {
	return g.edges[u]
}

func setIntersection(set1, set2 map[string]struct{}) map[string]struct{} {
	intersection := make(map[string]struct{})
	for key := range set1 {
		if _, found := set2[key]; found {
			intersection[key] = struct{}{}
		}
	}
	return intersection
}

func setToSlice(set map[string]struct{}) []string {
	slice := make([]string, 0, len(set))
	for key := range set {
		slice = append(slice, key)
	}
	return slice
}

func findNumCandidates(graph *Graph) int {
	total := 0
	for u := range graph.edges {
		for v := range graph.edges[u] {
			for w := range graph.edges[u] {
				if w == v {
					continue
				}
				if _, ok := graph.edges[w][v]; ok {
					if u[0] == 't' || v[0] == 't' || w[0] == 't' {
						total += 1
					}
				}
			}
		}
	}
	return total / 6
}

func maxClique(graph *Graph, nodes map[string]struct{}) map[string]struct{} {
	if len(nodes) == 0 {
		return make(map[string]struct{})
	}
	if len(nodes) == 1 {
		return nodes
	}

	tempNodes := make(map[string]struct{})
	for key := range nodes {
		tempNodes[key] = struct{}{}
	}

	var node string
	for key := range tempNodes {
		node = key
		delete(tempNodes, key)
		break
	}

	cliqueWithout := maxClique(graph, tempNodes)
	neighbors := graph.getNeighbors(node)
	cliqueWith := maxClique(graph, setIntersection(neighbors, tempNodes))
	cliqueWith[node] = struct{}{}

	if len(cliqueWith) > len(cliqueWithout) {
		return cliqueWith
	}
	return cliqueWithout
}

func parseInputs(bytes []uint8) *Graph {
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	graph := newGraph()
	for _, line := range lines {
		parts := strings.Split(line, "-")
		graph.addEdge(parts[0], parts[1])
	}
	return graph
}

func main() {
	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not read file: %v\n", err)
		os.Exit(1)
	}
	graph := parseInputs(bytes)
	fmt.Printf("Part I: %d\n", findNumCandidates(graph))

	nodes := make(map[string]struct{})
	for key := range graph.edges {
		nodes[key] = struct{}{}
	}
	maxClique := maxClique(graph, nodes)
	result := setToSlice(maxClique)
	sort.Strings(result)
	fmt.Printf("Part II: %s\n", strings.Join(result, ","))
}
