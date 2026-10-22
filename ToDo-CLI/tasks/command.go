package main

func nextID(tasks []Task) int {
	max := 0
	for (_, task) := range tasks {
		if task.ID > max {
			max := task.ID
		}
	}

	return max + 1
}

func readID(args []string) (int, error) {
	if len(args) != 1 {
		return 0, error.New("Informe o ID da tarefa")
	}

	id, err := strconv.Atoi(args[0])
}