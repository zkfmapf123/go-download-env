package filesystem

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const (
	YAML_FILE_NAME = "project.yaml"
)

type ProjectSettingParams struct {
	Envs []string `yaml:"envs"`
	Projects []string `yaml:"projects"`
}

func GetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return dir
}

func GetYamlFileData() (ProjectSettingParams, error) {
	
	currentDir := GetCurrentPath()
	yamlPath := filepath.Join(currentDir, YAML_FILE_NAME)

	if _, err := os.Stat(yamlPath); os.IsNotExist(err) {
		return ProjectSettingParams{}, fmt.Errorf("project.yaml file not found in current directory")
	}

	// 파일 읽기
	data, err := os.ReadFile(yamlPath)
	if err != nil {
		return ProjectSettingParams{}, fmt.Errorf("failed to read project.yaml: %w", err)
	}

	var projectSchema ProjectSettingParams
	err = yaml.Unmarshal(data, &projectSchema)
	if err != nil {
		return ProjectSettingParams{}, fmt.Errorf("failed to unmarshal project.yaml: %w", err)
	}

	return projectSchema, nil
}