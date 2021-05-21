package teamcity

import (
	"fmt"
	"strings"
)

type BuildInfo struct {
	Project string
	Tag string
}

var Projects = map[string]string {
	"api-addons": "airslateinc/addons-api",
	"api-integrations": "airslateinc/integrations-api",
	"api-static-data": "airslateinc/static-data-api",
	"api-addons-orchestrator": "airslateinc/addons-orchestrator-api",
	"webhook-sender": "airslateinc/webhook-sender",
	"most-useful-slackbot": "srgkas/most-useful-slackbot", // Test value
}

func (info *BuildInfo) GetProjectRepo() (string, error) {
	repo, exists := Projects[info.Project]

	if !exists {
		return "", fmt.Errorf("repo for project :%s does not exist in projecs map", info.Project)
	}

	return repo, nil
}

func ParseBuildInfo(message string) (*BuildInfo, error) {
	// String example to parse
	// Succeeded - AirSlate / PROD Env / PROD: Builds / Backend: API / most-useful-slackbot / Deploy #173 | - v0.0.1 [v0.0.1]>
	parts := strings.Split(message, "|")

	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid build string: %s. Expected parts 2, given:%d", message, len(parts))
	}

	tsTreeParts := strings.Split(parts[0], " / ")

	if len(tsTreeParts) < 6 {
		return nil, fmt.Errorf("invalid build string: %s. Expected parts 2, given:%d", message, len(tsTreeParts))
	}

	project := tsTreeParts[4]

	tagParts := strings.Split(parts[1], " ")
	tag := tagParts[2]

	return &BuildInfo{Project: project, Tag: tag}, nil
}