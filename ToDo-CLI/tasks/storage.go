package main

import (
	"encoding/json"
	"errors"
	"os"
)

const tasksFile = "tasks.json"

// func funcName(parameters) (returnType)
func load() ([]Task, error) {
	data, err := os.ReadFile(tasksFile)
	// Ve se lançou erro pois o slice ainda não existe
	if errors.Is(err, os.ErrNotExist) {
		return []Task{}, nil
	}

	// Verifica se aconteceu algum erro
	if err != nil {
		return nil, err
	}

	var tasks []Task
	// Colocado assim para que o err só seja assim dentro do if
	// Verifica se ocorreu erro no processo de transformar JSON -> Struct
	// tira de data e coloca no slice tasks
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil

}

func save(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")

	if err != nil {
		return err
	}

	return os.WriteFile(tasksFile, data, 0644)

}
