package config

import (
	"fmt"
)

// SQLDeployTargetConfig represents configuration for the SQL client.
type SQLDeployTargetConfig struct {
	Version string `json:"version"` /// Version of the SQL client.
	DBType string `json:"dbType"` // postgres, mysql, etc.
}

// SQLApplicationSpec represents app configuration for the SQL plugin.
type SQLApplicationSpec struct {
	// Input for SQL deployment (e.g., script path).
	Input SQLDeploymentInput `json:"input"`
	// Configuration for quick sync.
	QuickSync SQLDeployStageOptions `json:"quickSync"`
}

func (s *SQLApplicationSpec) Validate() error {
	if s.Input.ScriptPath == "" {
		return fmt.Errorf("scriptPath must not be empty")
	}
	if s.Input.DBType == "" {
		return fmt.Errorf("dbType must not be empty")
	}
	return nil
}

// SQLDeploymentInput defines inputs for SQL stages.
type SQLDeploymentInput struct {
	// Path to the SQL script file.
	ScriptPath string `json:"scriptPath"`
	// Database type (e.g., "postgres", "mysql").
	DBType string `json:"dbType"`
}

// SQLDeployStageOptions contains configurable fields for SQL_APPLY stage.
type SQLDeployStageOptions struct {
	// Additional flags for the SQL client.
	Flags []string `json:"flags"`
}

// SQLDiffStageOptions contains configurable fields for SQL_DIFF stage.
type SQLDiffStageOptions struct {
	// Exit pipeline if no changes are detected.
	ExitOnNoChanges bool `json:"exitOnNoChanges"`
}