package sane

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	MaxRules = 1000
)

type gitIgnoreStyleRuleEngine struct {
	gitIgnoreStyleRules []*gitIgnoreStyleRule
	strict              bool
}

func newRuleEngine(t RuleFormat, path string, strict bool) (RuleEngine, error) {
	if t != RULES_FORMAT_GITIGNORE {
		return nil, fmt.Errorf("unsupported rules type: %s", t)
	}

	return newGitIgnoreStyleRuleEngine(path, strict)
}

func newGitIgnoreStyleRuleEngine(path string, strict bool) (RuleEngine, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open rules: %w", err)
	}

	scanner := bufio.NewScanner(file)
	engine := gitIgnoreStyleRuleEngine{
		gitIgnoreStyleRules: make([]*gitIgnoreStyleRule, 0, MaxRules+1),
		strict:              strict,
	}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}

		rule := newGitIgnoreStyleRule(line)
		engine.gitIgnoreStyleRules = append(engine.gitIgnoreStyleRules, rule)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read rules: %w", err)
	}

	return &engine, err
}

func (g *gitIgnoreStyleRuleEngine) Validate(node RepositoryNode) (bool, error) {
	var ret bool
	var err error

	for _, p := range g.gitIgnoreStyleRules {
		result := p.match(node)
		switch result {
		case matchResultAllow:
			ret = true
			break
		case matchResultNoMatch:
			ret = false
			break
		case matchResultDeny:
			err = fmt.Errorf("%s is explicitly denied", node.FullPath)
			break
		}

		// First match break out
		if ret || (err != nil) {
			break
		}
	}

	return ret, err
}

func (g *gitIgnoreStyleRuleEngine) Finalize() error {
	var err error = nil
	for _, r := range g.gitIgnoreStyleRules {
		if g.strict && !r.anyMatch() {
			err = fmt.Errorf("rule:%s has no match in strict mode", r.rule)
			break
		}
	}

	return err
}

func (g *gitIgnoreStyleRuleEngine) Reset() {
	// Nothing
}
