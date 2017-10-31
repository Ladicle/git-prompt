package main

import "testing"

func TestUpdateBranchName(t *testing.T) {
	git := Git{}
	if err := git.UpdateBranchName(); err != nil {
		t.Fatalf("can not get branch name %#v", err)
	}
}

func TestUpdateLocalStatus(t *testing.T) {
	git := Git{}
	if err := git.UpdateLocalStatus(); err != nil {
		t.Fatalf("can not get status %#v", err)
	}
}

func TestUpdateRemoteStatus(t *testing.T) {
	git := Git{
		Branch: "master",
	}
	if err := git.UpdateRemoteStatus(); err != nil {
		t.Fatalf("can not get remote status %#v", err)
	}
}
