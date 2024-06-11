package functv

import "fmt"

func TopoSortExtend(m map[string][]string) ([]string, error) {
	var order []string
	seen := map[string]bool{}
	stack := map[string]bool{}
	var visitAll func([]string) error

	visitAll = func(items []string) error {
		for _, v := range items {
			if stack[v] {
				return fmt.Errorf("cyclic: %s", v)
			}
			if !seen[v] {
				seen[v] = true
				stack[v] = true
				err := visitAll(m[v])

				if err != nil {
					return err
				}
				stack[v] = false
				order = append(order, v)
			}
		}
		return nil
	}

	for k := range m {
		err := visitAll([]string{k})
		if err != nil {
			return nil, err
		}
	}
	return order, nil
}
