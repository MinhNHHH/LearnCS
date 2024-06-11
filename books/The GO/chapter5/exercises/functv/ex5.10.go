package functv

func TopoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	visitAll(keys)
	return order
}

func TopoSortNew(m map[string][]string) []string {
	inDegrees := map[string]int{}
	order := []string{}
	for node := range m {
		inDegrees[node] = 0
	}

	// Calculate in-degrees
	for _, neighbors := range m {
		for _, neighbor := range neighbors {
			inDegrees[neighbor]++
		}
	}

	queue := []string{}
	for node, degree := range inDegrees {
		if degree == 0 {
			queue = append(queue, node)
		}
	}

	// Perform topological sort
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		order = append(order, current)
		for _, neighbor := range m[current] {
			inDegrees[neighbor]--
			if inDegrees[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// If order contains all nodes, return order, otherwise there is a cycle
	if len(order) == len(inDegrees) {
		return order
	}
	return []string{} // Return an empty slice if there is a cycle
}
