package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/sync/errgroup"
)

func NewCurrentDirGit() (Git, error) {
	eg := errgroup.Group{}
	git := Git{}

	eg.Go(func() error { return git.UpdateLocalStatus() })
	eg.Go(func() error { return git.UpdateRemoteStatus() })

	err := eg.Wait()
	return git, err
}

type Git struct {
	Branch       string
	StagedNum    int // Un-commit change number
	ChangedNum   int // Un-staged change number
	UntrackedNum int // Un-tracked change number
	AheadNum     int // The number of ahead changes
	BehindNum    int // The number of behind changes
	ConflictNum  int
	NoRemote     bool
}

func (g *Git) Print() {
	var icon int
	if g.ConflictNum != 0 {
		icon = 0
	} else if g.BehindNum != 0 {
		icon = 1
	} else if g.AheadNum != 0 || g.NoRemote {
		icon = 2
	} else {
		icon = 3
	}

	var status string
	if g.StagedNum != 0 || g.ChangedNum != 0 || g.UntrackedNum != 0 {
		status = fmt.Sprintf(" %d,%d,%d", g.StagedNum, g.ChangedNum, g.UntrackedNum)
	} else if icon == 3 {
		icon = 4
	}

	fmt.Printf("%d %s%s", icon, g.Branch, status)
}

// UpdateRemoteStatus is function to update status of remote repository changes
func (g *Git) UpdateRemoteStatus() error {
	branch, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		return err
	}
	g.Branch = strings.TrimSpace(string(branch))

	ref := fmt.Sprintf("origin/%s...%s", g.Branch, g.Branch)
	revgit, err := exec.Command("git", "rev-list", "--left-right", ref).Output()
	if err != nil {
		g.NoRemote = true
		return nil
	}

	difflines := strings.Split(string(revgit), "\n")
	for _, r := range difflines {
		if strings.TrimSpace(r) == "" {
			continue
		}
		if strings.HasPrefix(r, "<") {
			g.BehindNum++
		} else {
			g.AheadNum++
		}
	}
	return nil
}

func (g *Git) UpdateLocalStatus() error {
	s, err := exec.Command("git", "status", "--porcelain").Output()
	if err != nil {
		return err
	}
	states := strings.Split(string(s), "\n")
	if len(states[0]) == 0 {
		return nil
	}
	for _, s := range states {
		if len(s) < 1 {
			continue
		}
		switch s[0] {
		case '?':
			g.UntrackedNum++
		case 'A', 'M', 'D', 'R':
			g.StagedNum++
		}
		switch s[1] {
		case 'M', 'D':
			g.ChangedNum++
		case 'U':
			g.ConflictNum++
		}
	}
	return nil
}

func main() {
	git, err := NewCurrentDirGit()
	if err != nil {
		os.Exit(0)
	}
	git.Print()
}
