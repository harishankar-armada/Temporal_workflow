package workflows

// import (
// 	"time"

// 	"github.com/H4rish4nk4r/github-runner/activities"
// 	"go.temporal.io/sdk/workflow"
// )

// func VClusterWorkflow(ctx workflow.Context, input activities.VClusterInput) error {
// 	opts := workflow.ActivityOptions{
// 		StartToCloseTimeout: time.Minute * 5,
// 	}
// 	ctx = workflow.WithActivityOptions(ctx, opts)

// 	return workflow.ExecuteActivity(ctx, activities.TriggerGitHubWorkflow, input).Get(ctx, nil)
// }
