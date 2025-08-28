package main

import (
	"log"

	"github.com/H4rish4nk4r/github-runner/activities"
	"github.com/H4rish4nk4r/github-runner/workflows"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}
	defer c.Close()

	w := worker.New(c, "vcluster-task-queue", worker.Options{})

	w.RegisterWorkflow(workflows.VClusterWorkflow)
	w.RegisterActivity(activities.TriggerGitHubWorkflow)

	log.Println("Starting worker...")
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
