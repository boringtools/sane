package sane

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGitIgnoreStyleRuleMatching(t *testing.T) {
	cases := []struct {
		name    string
		pattern string
		path    string
		result  matchResult
	}{
		{
			"Match file",
			"README.md",
			"README.md",
			matchResultAllow,
		},
		{
			"Does not match a file",
			"/README.md",
			"a.md",
			matchResultNoMatch,
		},
		{
			"Explicit deny a file",
			"!README.md",
			"README.md",
			matchResultDeny,
		},
		{
			"Does not match file in sub-directory",
			"/README.md",
			"a/README.md",
			matchResultNoMatch,
		},
		{
			"Does not deny file in sub-directory",
			"!/README.md",
			"a/README.md",
			matchResultNoMatch,
		},
		{
			"Match a glob",
			"/docs/*.md",
			"docs/hello.md",
			matchResultAllow,
		},
		{
			"Match a super glob",
			"/docs/**/*.png",
			"docs/a/b/c/d/e.png",
			matchResultAllow,
		},
		{
			"Deny with super glob",
			"!/docs/**/*.png",
			"docs/a/b/c/d/e.png",
			matchResultDeny,
		},
		{
			"Deny any file matching pattern",
			"!*.exe",
			"docs/a/b/c.exe",
			matchResultDeny,
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			g := newGitIgnoreStyleRule(test.pattern)
			r := g.match(RepositoryNode{FullPath: test.path})

			assert.Equal(t, test.result, r)
		})
	}
}
