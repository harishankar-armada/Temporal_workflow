package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/H4rish4nk4r/github-runner/workflows"
	"github.com/joho/godotenv"
	"go.temporal.io/sdk/client"
)

func main() {

	http.HandleFunc("/", serveForm)
	http.HandleFunc("/submit", handleSubmit)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
func serveForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, nil)
}

func handleSubmit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	name := r.FormValue("vclusterName")
	cpu := r.FormValue("cpu")
	namespace := name + "ns"
	memory := r.FormValue("memory")
	storage := r.FormValue("storage")

	_ = godotenv.Load()

	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatalln("GITHUB_TOKEN not set in environment")
	}

	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("cannot create Temporal client", err)
	}
	defer c.Close()

	workflowID := "vcluster-workflow-" + time.Now().Format("20060102150405")

	params := workflows.VClusterParams{
		GitHubToken:  token,
		RepoOwner:    "H4rish4nk4r",
		RepoName:     "github_runner2",
		Branch:       "main",
		RunnerName:   "hari-runner",
		RunnerLabels: "self-hosted",
		ClusterName:  name,
		Namespace:    namespace,
		Memory:       memory,
		Storage:      storage,
		Cpu:          cpu,
	}

	we, err := c.ExecuteWorkflow(context.Background(), client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: "vcluster-task-queue",
	}, workflows.VClusterWorkflow, params)
	if err != nil {
		log.Fatalln("unable to execute workflow", err)
	}

	log.Println("Started workflow:", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
}
