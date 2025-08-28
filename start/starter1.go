package main

// import (
// 	"context"
// 	"log"
// 	"os"
// 	"time"

// 	"github.com/H4rish4nk4r/github-runner/activities"
// 	"github.com/H4rish4nk4r/github-runner/workflows"
// 	"github.com/joho/godotenv"
// 	"go.temporal.io/sdk/client"
// )

// func main() {
// 	_ = godotenv.Load()

// 	token := os.Getenv("GITHUB_TOKEN")
// 	if token == "" {
// 		log.Fatalln("GITHUB_TOKEN not set in environment")
// 	}

// 	c, err := client.NewClient(client.Options{})
// 	if err != nil {
// 		log.Fatalln("cannot create Temporal client", err)
// 	}
// 	defer c.Close()

// 	workflowID := "vcluster-workflow-" + time.Now().Format("20060102150405")

// 	input := activities.VClusterInput{
// 		GitHubToken: token,
// 		RepoOwner:   "H4rish4nk4r",
// 		RepoName:    "github_runner",
// 		Branch:      "main",
// 		ClusterName: "vcluster",
// 		Namespace:   "vcluster",
// 	}

// 	we, err := c.ExecuteWorkflow(context.Background(), client.StartWorkflowOptions{
// 		ID:        workflowID,
// 		TaskQueue: "vcluster-task-queue",
// 	}, workflows.VClusterWorkflow, input)
// 	if err != nil {
// 		log.Fatalln("unable to execute workflow", err)
// 	}

// 	log.Println("Started workflow:", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
// }
