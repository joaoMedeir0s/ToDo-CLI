package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func nextID(tasks []Task) int {
	max := 0
	for _, task := range tasks {
		if task.ID > max {
			max = task.ID
		}
	}

	return max + 1
}

func readID(args []string) (int, error) {
	if len(args) != 1 {
		return 0, errors.New("Informe o ID da tarefa")
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return 0, errors.New("ID inválido")
	}

	return id, nil
}

func cmdAdd(args []string) error {
	title := strings.TrimSpace(strings.Join(args, " "))
	if title == "" {
		return errors.New("Title can't be empty!")
	}

	tasks, err := load()
	if err != nil {
		return err //Aqui retorna err porque já foram especificados em load()
	}

	tasks = append(tasks, Task{
		ID:        nextID(tasks),
		Title:     title,
		CreatedAt: time.Now(),
	})

	if err := save(tasks); err != nil {
		return err
	}

	fmt.Printf("Task add: %s\n", title)

	return nil
}

func cmdList() error {
	tasks, err := load()
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks for now")
		return nil
	}
	// lembrando que a função range dá o index (primeiro retorno)
	// e o elemento nesse index (segunto retorno)
	for _, task := range tasks {
		check := " "
		if task.Completed {
			check = "✅"
		}
		fmt.Printf("[%s] %d - %s\n", check, task.ID, task.Title)
	}
	return nil

}
