package sane

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type validationResult struct {
	valid bool
	err   error
}

func TestGitIgnoreStyleRuleEngine(t *testing.T) {
	cases := []struct {
		name string

		strict bool
		rules  []string
		files  []string

		// Must have a result for each file
		results []validationResult

		finalErr error
	}{
		{
			"All files are validated",
			false,
			[]string{
				"README.md",
				"docs/*.md",
				"pkg/**/*.go",
			},
			[]string{
				"docs/a.md",
				"pkg/a.go",
				"pkg/b/c.go",
			},
			[]validationResult{
				{true, nil},
				{true, nil},
				{true, nil},
			},

			nil,
		},
		{
			"Must fail finally in strict mode",
			true,
			[]string{
				"README.md",
				"pkg/**/*.go",
			},
			[]string{
				"pkg/a.go",
			},
			[]validationResult{
				{true, nil},
			},

			errors.New("rule:README.md has no match in strict mode"),
		},
		{
			"Must fail for explicit deny",
			true,
			[]string{"!README.md", "pkg/**/*.go"},
			[]string{"README.md"},

			[]validationResult{
				{false, errors.New("README.md is explicitly denied")},
			},

			errors.New("rule:pkg/**/*.go has no match in strict mode"),
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			engine, err := newRuleEngineWithReader(RULES_FORMAT_GITIGNORE,
				strings.NewReader(strings.Join(test.rules, "\n")), test.strict)
			assert.Nil(t, err)

			for idx, file := range test.files {
				ret, err := engine.Validate(RepositoryNode{FullPath: file})
				assert.Equal(t, test.results[idx].valid, ret)
				assert.Equal(t, test.results[idx].err, err)
			}

			err = engine.Finalize()
			assert.Equal(t, test.finalErr, err)
		})
	}
}
