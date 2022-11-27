package sane

import (
	"fmt"

	"github.com/go-git/go-git/v5"
)

type gitRepositoryWalker struct {
	path string
}

func newGitRepositoryWalker(path string) (RepositoryWalker, error) {
	return &gitRepositoryWalker{path: path}, nil
}

func (g *gitRepositoryWalker) Walk(handler RepositoryNodeHandler) error {
	repo, err := git.PlainOpen(g.path)
	if err != nil {
		return err
	}

	fmt.Printf("Git repo opened: %v\n", repo)
	return nil
}
