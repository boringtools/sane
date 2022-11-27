package sane

import (
	"strings"
	"sync/atomic"

	"github.com/go-git/go-git/v5/plumbing/format/gitignore"
)

type matchResult uint16

const (
	matchResultNoMatch = iota
	matchResultAllow
	matchResultDeny
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

func (g *gitIgnoreStyleRule) match(node RepositoryNode) matchResult {
	if g.isIgnorable(node) {
		return matchResultAllow
	}

	var ret matchResult
	mr := g.pattern.Match(strings.Split(node.FullPath, "/"), false)

	switch mr {
	case gitignore.NoMatch:
		ret = matchResultNoMatch
	case gitignore.Exclude:
		ret = matchResultAllow
	case gitignore.Include:
		ret = matchResultDeny
	default:
		ret = matchResultNoMatch
	}

	if ret == matchResultAllow {
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
