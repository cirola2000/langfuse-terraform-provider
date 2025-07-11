package provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Client wraps the Langfuse API client
type Client struct {
	ApiHost   string
	SecretKey string
	PublicKey string
	client    *http.Client
}

// Project represents a Langfuse project
type Project struct {
	ID            string                 `json:"id"`
	Name          string                 `json:"name"`
	Metadata      map[string]interface{} `json:"metadata"`
	RetentionDays *int                   `json:"retentionDays"`
	CreatedAt     string                 `json:"createdAt"`
	UpdatedAt     string                 `json:"updatedAt"`
}

// ProjectsResponse represents the response from the projects list endpoint
type ProjectsResponse struct {
	Projects []Project `json:"projects"`
}

// CreateProjectRequest represents the request to create a project
type CreateProjectRequest struct {
	Name      string                 `json:"name"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	Retention *int                   `json:"retention,omitempty"`
}

// UpdateProjectRequest represents the request to update a project
type UpdateProjectRequest struct {
	Name      string                 `json:"name"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	Retention *int                   `json:"retention,omitempty"`
}

// NewClient creates a new Langfuse API client
func NewClient(apiHost, secretKey, publicKey string) *Client {
	return &Client{
		ApiHost:   apiHost,
		SecretKey: secretKey,
		PublicKey: publicKey,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// makeRequest performs an HTTP request with authentication
func (c *Client) makeRequest(method, endpoint string, body interface{}) (*http.Response, error) {
	var reqBody *bytes.Buffer
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("error marshaling request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonData)
	} else {
		reqBody = bytes.NewBuffer(nil)
	}

	url := c.ApiHost + endpoint
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.PublicKey, c.SecretKey)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return resp, nil
}

// ListProjects retrieves all projects for the organization
func (c *Client) ListProjects() ([]Project, error) {
	resp, err := c.makeRequest("GET", "/api/public/organizations/projects", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status %d", resp.StatusCode)
	}

	var projectsResp ProjectsResponse
	if err := json.NewDecoder(resp.Body).Decode(&projectsResp); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return projectsResp.Projects, nil
}

// CreateProject creates a new project
func (c *Client) CreateProject(req CreateProjectRequest) (*Project, error) {
	resp, err := c.makeRequest("POST", "/api/public/projects", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status %d", resp.StatusCode)
	}

	var project Project
	if err := json.NewDecoder(resp.Body).Decode(&project); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &project, nil
}

// GetProject retrieves a project by ID (implemented using ListProjects)
func (c *Client) GetProject(projectID string) (*Project, error) {
	projects, err := c.ListProjects()
	if err != nil {
		return nil, err
	}

	for _, project := range projects {
		if project.ID == projectID {
			return &project, nil
		}
	}

	return nil, fmt.Errorf("project with ID %s not found", projectID)
}

// UpdateProject updates an existing project
func (c *Client) UpdateProject(projectID string, req UpdateProjectRequest) (*Project, error) {
	endpoint := fmt.Sprintf("/api/public/projects/%s", projectID)
	resp, err := c.makeRequest("PUT", endpoint, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status %d", resp.StatusCode)
	}

	var project Project
	if err := json.NewDecoder(resp.Body).Decode(&project); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &project, nil
}

// DeleteProject deletes a project by ID
func (c *Client) DeleteProject(projectID string) error {
	endpoint := fmt.Sprintf("/api/public/projects/%s", projectID)
	resp, err := c.makeRequest("DELETE", endpoint, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		return fmt.Errorf("API request failed with status %d", resp.StatusCode)
	}

	return nil
} 