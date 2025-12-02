package common

func BFS[T comparable](start T, neighbors func(T) []T, isGoal func(T) bool) ([]T, int, bool) {
	queue := NewQueue[T]()
	queue.Enqueue(start)

	visited := make(map[T]bool)
	visited[start] = true

	parent := make(map[T]T)
	distance := make(map[T]int)
	distance[start] = 0

	for !queue.IsEmpty() {
		current, _ := queue.Dequeue()

		if isGoal(current) {
			// Reconstruct path
			path := []T{current}
			for node := current; node != start; node = parent[node] {
				path = append([]T{parent[node]}, path...)
			}
			return path, distance[current], true
		}

		for _, neighbor := range neighbors(current) {
			if !visited[neighbor] {
				visited[neighbor] = true
				parent[neighbor] = current
				distance[neighbor] = distance[current] + 1
				queue.Enqueue(neighbor)
			}
		}
	}

	return nil, 0, false
}
