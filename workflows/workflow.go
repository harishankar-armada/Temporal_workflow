package workflows

import (
	"time"

	"github.com/H4rish4nk4r/github-runner/activities"
	"go.temporal.io/sdk/workflow"
)

// Make sure this struct is defined somewhere accessible
type VClusterParams struct {
	GitHubToken  string
	RepoOwner    string
	RepoName     string
	Branch       string
	ClusterName  string
	Namespace    string
	RunnerName   string
	RunnerLabels string
	Memory       string
	Storage      string
	Cpu          string
}

func VClusterWorkflow(ctx workflow.Context, params VClusterParams) error {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: 5 * time.Minute,
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	// Reference the activity function directly
	return workflow.ExecuteActivity(ctx, activities.TriggerGitHubWorkflow, params).Get(ctx, nil)
}
