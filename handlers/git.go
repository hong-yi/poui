package handlers

import (
	"fmt"
	"log/slog"
	"strings"
)

func GetRepoWebUrl(gitUrl string) string {
	gitSshUserIdx := strings.Index(gitUrl, "@")
	if gitSshUserIdx > 0 {
		gitUrl = strings.ReplaceAll(gitUrl[gitSshUserIdx+1:], ":", "/")
	}
	gitUrl = strings.TrimSuffix(gitUrl, ".git")
	slog.Debug(fmt.Sprintf("repository web url: %v", gitUrl))
	return "http://" + gitUrl
}
