package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"ruopen/models"
	"strings"
)

func GetTfeWorkspaceUrls(tfPath string, maxDepth int) []string {
	tfstatePaths := scanTfeDir(tfPath, maxDepth)
	tfeWorkspaceUrls := []string{}
	for _, path := range tfstatePaths {
		tfState, err := parseTfState(path)
		workspaceUrl, err := getWorkspaceUrl(tfState)
		if err != nil {
			slog.Error(fmt.Sprintf("unable to parse tfstate: %v", err))
			return nil
		}
		tfeWorkspaceUrls = append(tfeWorkspaceUrls, *workspaceUrl)
	}
	return tfeWorkspaceUrls
}

func scanTfeDir(tfPath string, maxDepth int) []string {
	var tfstatePaths []string
	err := filepath.WalkDir(tfPath, func(path string, d fs.DirEntry, err error) error {
		srcPathDepth := strings.Count(path, string(os.PathSeparator)) - strings.Count(tfPath, string(os.PathSeparator))
		if d.IsDir() && srcPathDepth > maxDepth {
			return fs.SkipDir
		}
		slog.Debug(d.Name())
		if filepath.Ext(path) == ".tfstate" {
			tfstatePaths = append(tfstatePaths, path)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	return tfstatePaths
}

func getWorkspaceUrl(tfState *models.TfState) (*string, error) {
	if tfState.Backend.Type != "cloud" {
		return nil, errors.New("ERR_NOT_TFE")
	}

	url := fmt.Sprintf("https://app.terraform.io/app/%s/workspaces/%s", tfState.Backend.Config.Organization, tfState.Backend.Config.Workspaces.Name)
	return &url, nil
}

func parseTfState(path string) (*models.TfState, error) {
	var tfState models.TfState
	tfs, err := os.ReadFile(path)
	err = json.Unmarshal(tfs, &tfState)
	if err != nil {
		return nil, err
	}
	slog.Debug(fmt.Sprintf("%v", tfState))
	return &tfState, nil
}
