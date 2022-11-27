package sane

import (
	"bufio"
	"fmt"
	"os"
)

const (
	MaxRules = 1000
)

type globRule struct {
	rule string
	ctr  uint32
}

type globRuleEngine struct {
	globRules []globRule
	strict    bool
}

func newRuleEngine(t RuleFormat, path string, strict bool) (Rule, error) {
	if t != RULES_FORMAT_GITIGNORE {
		return nil, fmt.Errorf("unsupported rules type: %s", t)
	}

	return newGlobRuleEngine(path, strict)
}

func newGlobRuleEngine(path string, strict bool) (Rule, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open rules: %w", err)
	}

	scanner := bufio.NewScanner(file)
	engine := globRuleEngine{
		globRules: make([]globRule, 0, MaxRules+1),
		strict:    strict,
	}

	for scanner.Scan() {
		engine.globRules = append(engine.globRules,
			globRule{rule: scanner.Text(), ctr: 0})
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read rules: %w", err)
	}

	return &engine, err
}

func (g *globRuleEngine) Validate(node RepositoryNode) (bool, error) {
	return true, nil
}

func (g *globRuleEngine) Finalize() error {
	var err error = nil

	for _, r := range g.globRules {
		if g.strict && r.ctr == 0 {
			err = fmt.Errorf("rule:%s has no match in strict mode", r.rule)
			break
		}
	}

	return err
}
