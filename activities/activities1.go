package activities

// import (
// 	"bytes"
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// )

// type VClusterInput struct {
// 	GitHubToken string
// 	RepoOwner   string
// 	RepoName    string
// 	Branch      string
// 	ClusterName string
// 	Namespace   string
// }

// func TriggerGitHubWorkflow(ctx context.Context, input VClusterInput) error {
// 	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/actions/workflows/vcluster-create.yml/dispatches", input.RepoOwner, input.RepoName)

// 	payload := map[string]interface{}{
// 		"ref": input.Branch,
// 		"inputs": map[string]string{
// 			"cluster_name": input.ClusterName,
// 			"namespace":    input.Namespace,
// 		},
// 	}

// 	body, _ := json.Marshal(payload)
// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
// 	if err != nil {
// 		return err
// 	}

// 	req.Header.Set("Authorization", "Bearer "+input.GitHubToken)
// 	req.Header.Set("Accept", "application/vnd.github+json")

// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusNoContent {
// 		return fmt.Errorf("GitHub API returned status: %s", resp.Status)
// 	}

// 	return nil
// }
