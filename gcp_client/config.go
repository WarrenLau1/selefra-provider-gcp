package gcp_client

type Config struct {
	ProjectIDs                []string `yaml:"project_ids,omitempty" mapstructure:"project_ids"`
	FolderIDs                 []string `yaml:"folder_ids,omitempty" mapstructure:"folder_ids"`
	FolderMaxDepth            uint     `yaml:"folders_max_depth,omitempty" mapstructure:"folders_max_depth"`
	ServiceAccountKeyJSON     string   `yaml:"service_account_key_json,omitempty" mapstructure:"service_account_key_json"`
	AccountName               string   `yaml:"account_name,omitempty" mapstructure:"account_name"`
	ServiceAccountEnvKey      string   `yaml:"service_account_env_key,omitempty" mapstructure:"service_account_env_key"`
	ServiceAccountKeyJSONFile string   `yaml:"service_account_env_key,omitempty" mapstructure:"service_account_env_key"`
}
