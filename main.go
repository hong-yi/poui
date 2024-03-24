package main

import (
	"flag"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/pkg/errors"
	"log/slog"
	"os"
	"path/filepath"
	"ruopen/handlers"
	"ruopen/tui"
)

var (
	selectedTfeWorkspaces []string
	selectedGitRepoUrls   []string
)

func main() {
	openTfc := flag.Bool("tfc", false, "whether to open Terraform Cloud")
	openGit := flag.Bool("git", false, "whether to open your VCS WebUI")
	openAll := flag.Bool("a", false, "whether to open both TFE and VCS WebUI")
	maxDepth := flag.Int("maxScanDepth", 3, "maximum depth to scan for .tfstate files")
	flag.Parse()

	if os.Getenv("POUI_DEBUG") == "1" {
		h := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug})
		slog.SetDefault(slog.New(h))
	}

	path, err := os.Getwd()
	if err != nil {
		slog.Error(fmt.Sprintf("unable to get working dir - %v", err))
	}

	if *openTfc || *openAll {
		tfeUrls, erroredUrls := handlers.GetTfeWorkspaceUrls(path, *maxDepth)
		slog.Debug(fmt.Sprintf("%v", tfeUrls))

		if len(erroredUrls) > 0 {
			fmt.Println(tui.WarningStyle.Render("[WARN] the following tfstate files were unable to be parsed: \n"))
			for _, url := range erroredUrls {
				fmt.Println(fmt.Sprintf("\t- %s", url))
			}
			fmt.Println()
			fmt.Println("Trying any remaining workspaces...")
			fmt.Println()
		}

		if len(tfeUrls) < 1 {
			fmt.Println(tui.WarningStyle.Render("Unable to find any tfstate.\n\nPlease run `terraform init` and try again."))
			return
		}
		if len(tfeUrls) > 1 {
			err := tui.CreateMultiSelectForm("workspaces found. Select which workspace(s) you wish to open.", tfeUrls, &selectedTfeWorkspaces).Run()
			if err != nil {
				return
			}
		} else {
			selectedTfeWorkspaces = append(selectedTfeWorkspaces, tfeUrls...)
		}

		for _, url := range selectedTfeWorkspaces {
			handlers.Open(url)
		}
	}

	// for git
	if *openGit || *openAll {
		tld, err := findTopLevelGitDir(path)
		slog.Debug(tld)
		if err != nil {
			slog.Error(fmt.Sprintf("%v", err))
			fmt.Println(tui.WarningStyle.Render("Unable to find any repositories.\n\nPlease run `git init` and try again."))
			return
		}
		r, err := git.PlainOpen(tld)
		rr, err := r.Remotes()
		if err != nil {
			slog.Error(fmt.Sprintf("%v", err))
			return
		}
		var repoUrls []string
		if len(rr) > 1 {
			for _, remote := range rr {
				gitUrl := handlers.GetRepoWebUrl(remote.Config().URLs[0])
				repoUrls = append(repoUrls, gitUrl)
			}
			err := tui.CreateMultiSelectForm("remote repositories found. Select which repositories you wish to open.", repoUrls, &selectedGitRepoUrls).Run()
			if err != nil {
				return
			}
		}
		if len(rr) == 1 {
			selectedGitRepoUrls = append(selectedGitRepoUrls, handlers.GetRepoWebUrl(rr[0].Config().URLs[0]))
		}
		for _, url := range selectedGitRepoUrls {
			err = handlers.Open(url)
			if err != nil {
				return
			}
		}
	}
}

func findTopLevelGitDir(workingDir string) (string, error) {
	dir, err := filepath.Abs(workingDir)
	if err != nil {
		return "", errors.Wrap(err, "invalid working dir")
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, ".git")); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", errors.New("no git repository found")
		}
		dir = parent
	}
}
