package sane

import (
	"strings"
	"sync/atomic"

	"github.com/go-git/go-git/v5/plumbing/format/gitignore"
)

type gitIgnoreStyleRule struct {
	rule    string
	ctr     uint32
	pattern gitignore.Pattern
}

func newGitIgnoreStyleRule(pattern string) *gitIgnoreStyleRule {
	p := gitignore.ParsePattern(pattern, nil)
	return &gitIgnoreStyleRule{
		rule:    pattern,
		ctr:     0,
		pattern: p,
	}
}

func (g *gitIgnoreStyleRule) match(node RepositoryNode) bool {
	if g.isIgnorable(node) {
		return true
	}

	var ret bool
	mr := g.pattern.Match(strings.Split(node.FullPath, "/"), false)

	switch mr {
	case gitignore.NoMatch:
		ret = false
	case gitignore.Exclude:
		ret = true
	case gitignore.Include:
		ret = false
	default:
		ret = false
	}

	if ret {
		atomic.AddUint32(&g.ctr, 1)
	}

	return ret
}

func (g *gitIgnoreStyleRule) anyMatch() bool {
	return (g.ctr > 0)
}

func (g *gitIgnoreStyleRule) isIgnorable(node RepositoryNode) bool {
	ignoredFiles := []string{
		".gitignore",
		".sane",
	}

	for _, ignore := range ignoredFiles {
		if strings.HasSuffix(node.FullPath, ignore) {
			return true
		}
	}

	return false
}
