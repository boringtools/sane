package sane

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
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

	ref, err := repo.Head()
	if err != nil {
		return fmt.Errorf("unable to find HEAD reference: %w", err)
	}

	commit, err := repo.CommitObject(ref.Hash())
	if err != nil {
		return fmt.Errorf("no commit object found with head ref hash: %s",
			ref.Hash().String())
	}

	tree, err := commit.Tree()
	if err != nil {
		return fmt.Errorf("no tree object found from head ref commit hash: %s",
			commit.Hash.String())
	}

	Infof("Using ref: %s", ref.Name())

	files := tree.Files()
	defer files.Close()

	return files.ForEach(func(f *object.File) error {
		return handler(RepositoryNode{FullPath: f.Name})
	})
}
