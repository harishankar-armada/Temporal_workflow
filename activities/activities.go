package activities

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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

func TriggerGitHubWorkflow(ctx context.Context, params VClusterParams) error {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/actions/workflows/cluster-provision.yaml/dispatches", params.RepoOwner, params.RepoName)

	payload := map[string]interface{}{
		"ref": params.Branch,
		"inputs": map[string]string{
			"cluster_name": params.ClusterName,
			"namespace_name":    params.Namespace,
		},
	}

	body, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "token "+params.GitHubToken) // <-- here
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Content-Type", "application/json") // <-- here

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("GitHub API returned status: %s, body: %s", resp.Status, string(bodyBytes))
	}

	return nil
}
