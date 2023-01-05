package gcp_client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	cloudresourcemanager "google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/option"
)

type Client struct {
	projects []string

	GcpServices *GcpServices

	ProjectId string
}

const (
	defaultProjectIdName = "<CHANGE_THIS_TO_YOUR_PROJECT_ID>"
	serviceAccountEnvKey = "SELEFRA_SERVICE_ACCOUNT_KEY_JSON"
)

func NewGcpClient(projects []string, services *GcpServices) *Client {
	c := &Client{
		projects:    projects,
		GcpServices: services,
	}
	if len(projects) == 1 {
		c.ProjectId = projects[0]
	}
	return c
}

func (c Client) withProject(project string) *Client {
	c.ProjectId = project
	return &c
}

func isValidJson(content []byte) error {
	var v map[string]interface{}
	err := json.Unmarshal(content, &v)
	if err != nil {
		var syntaxError *json.SyntaxError
		if errors.As(err, &syntaxError) {
			return fmt.Errorf("the environment variable %s should contain valid JSON object. %w", serviceAccountEnvKey, err)
		}
		return err
	}
	return nil
}

func NewClients(config Config) ([]*Client, error) {
	client, err := newClient(config)
	if err != nil {
		return nil, err
	}
	return []*Client{client}, nil
}

func newClient(providerConfig Config) (*Client, error) {

	projects := providerConfig.ProjectIDs
	if providerConfig.FolderMaxDepth == 0 {
		providerConfig.FolderMaxDepth = 5
	}

	serviceAccountKeyJSON := []byte(providerConfig.ServiceAccountKeyJSON)
	if len(serviceAccountKeyJSON) == 0 && providerConfig.ServiceAccountKeyJSONFile != "" {
		data, err := os.ReadFile(providerConfig.ServiceAccountKeyJSONFile)
		if err == nil {
			serviceAccountKeyJSON = data
		}
	}
	if len(serviceAccountKeyJSON) == 0 {
		if providerConfig.ServiceAccountEnvKey == "" {
			providerConfig.ServiceAccountEnvKey = serviceAccountEnvKey
		}
		serviceAccountKeyJSON = []byte(os.Getenv(providerConfig.ServiceAccountEnvKey))
	}

	options := []option.ClientOption{option.WithRequestReason("selefra resource fetch")}
	if len(serviceAccountKeyJSON) != 0 {
		if err := isValidJson(serviceAccountKeyJSON); err != nil {
			return nil, err
		}
		options = append(options, option.WithCredentialsJSON(serviceAccountKeyJSON))
	}

	services, err := initServices(context.Background(), options)
	if err != nil {
		return nil, fmt.Errorf("err:%w projects %v name:%s", err, projects, providerConfig.AccountName)
	}

	if len(projects) == 0 {
		var err error
		projects, err = getProjectsV1(options...)
		if err != nil {
			return nil, fmt.Errorf("get projects(v1) failed: %w, projects:%v name:%s", err, projects, providerConfig.AccountName)
		}
	}

	err = validateProjects(projects)
	if err != nil {
		return nil, err
	}

	client := NewGcpClient(projects, services)

	return client, nil
}

func validateProjects(projects []string) error {
	for _, project := range projects {
		if project == defaultProjectIdName {
			return errors.New("please specify a valid project_id in config.yml instead of <CHANGE_THIS_TO_YOUR_PROJECT_ID>")
		}
		if project == "" {
			return errors.New("please specify a valid project_id in config.yml instead of empty string")
		}
	}
	return nil
}

func getProjectsV1(options ...option.ClientOption) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	service, err := cloudresourcemanager.NewService(ctx, options...)
	if err != nil {
		return nil, err
	}
	var (
		projects []string
	)

	call := service.Projects.List().Context(ctx).Filter("lifecycleState=ACTIVE")

	for {
		output, err := call.Do()
		if err != nil {
			return nil, err
		}
		for _, project := range output.Projects {
			projects = append(projects, project.ProjectId)
		}
		if output.NextPageToken == "" {
			break
		}
		call.PageToken(output.NextPageToken)
	}

	if len(projects) == 0 {
		return nil, fmt.Errorf("no active projects")
	}

	return projects, nil
}
