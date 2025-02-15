package filesystem

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	YAML_FILE_NAME = "project.yaml"
)

type ProjectSettingParams struct {
	Envs []string `yaml:"envs"`
	Projects map[string][]string `yaml:"projects"`
	IsUseCommonEnvironments bool `yaml:"is_use_common_environment"`
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

func GetEnvFilesCurrentDir() ([]string, error) {
	path := GetCurrentPath()

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	// .env로 시작하는 파일들 필터링
	var envFiles []string
	for _, entry := range entries {
		if !entry.IsDir() && strings.Contains(entry.Name(), ".env") {
			envFiles = append(envFiles, entry.Name())
		}
	}

	return envFiles, nil
}

func EnvFileToMap(envFile string) (map[string]string, error) {
	path := GetCurrentPath()

	data, err := os.ReadFile(filepath.Join(path, envFile))
	if err != nil {
		return nil, fmt.Errorf("failed to read env file: %w", err)
	}

	envMap := make(map[string]string)
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.Contains(line, "=") {
			parts := strings.SplitN(line, "=", 2)
			envMap[parts[0]] = parts[1]
		}
	}

	return envMap, nil
}
