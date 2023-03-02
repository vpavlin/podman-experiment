package main

import (
	"context"
	"fmt"
	"log"
	"podman-demo/pkg/containermanager"
	"time"

	"github.com/google/uuid"
	"gopkg.in/square/go-jose.v2/json"
)

type Data struct {
	Id     string
	Name   string
	Number int
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := run(ctx)
	if err != nil {
		log.Fatal(err)
	}

}

func run(ctx context.Context) error {
	m, err := containermanager.NewManager("unix://run/user/1000/podman/podman.sock")
	if err != nil {
		return err
	}
	defer m.Close(true)

	input := Data{
		Id:     "id",
		Name:   "test",
		Number: 10,
	}
	data, err := json.MarshalIndent(input, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println("Input: \n", string(data))

	name := uuid.NewString()

	task := containermanager.Task{
		Timeout: 15 * time.Second,
		Image:   "transformer",
		Name:    name,
		Input:   input,
	}

	mt, cancel := m.WithTimeout(15 * time.Second)
	defer cancel()
	err = mt.PullIfNotPresent(nil, task.Image)
	if err != nil {
		return err
	}

	out, err := m.Task(task)
	if err != nil {
		return err
	}

	task.Name = uuid.NewString()

	out, err = m.Task(task)
	if err != nil {
		return err
	}

	data, err = json.MarshalIndent(out, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println("STDOUT: \n", string(data))

	return nil
}
