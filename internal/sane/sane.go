package sane

import (
	"fmt"
	"strings"
)

const (
	REPOSITORY_PATH_PLACEHOLDER  = "$REPOSITORY"
	REPOSITORY_SANE_DEFAULT_FILE = ".sane"
)

// Primary entrypoint for validator
func Execute(config Config) error {
	if strings.HasPrefix(config.RulesPath, REPOSITORY_PATH_PLACEHOLDER) {
		config.RulesPath = strings.Replace(config.RulesPath,
			REPOSITORY_PATH_PLACEHOLDER, config.RepositoryPath, 1)
	}

	Infof("Using Repository: %s", config.RepositoryPath)
	Infof("Using Rule: %s", config.RulesPath)

	repository, err := newGitRepositoryWalker(config.RepositoryPath)
	if err != nil {
		return err
	}

	ruleEngine, err := newRuleEngine(config.RulesType,
		config.RulesPath, config.Strict)
	if err != nil {
		return err
	}

	violatingPaths := []string{}
	err = repository.Walk(func(node RepositoryNode) error {
		ok, err := ruleEngine.Validate(node)
		if err != nil {
			return err
		}

		if !ok {
			if config.FailFast {
				return fmt.Errorf("fail-fast: no rules match for: %s", node.FullPath)
			} else {
				violatingPaths = append(violatingPaths, node.FullPath)
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	if len(violatingPaths) > 0 {
		for _, path := range violatingPaths {
			Infof("Failed: %s", path)
		}

		return fmt.Errorf("%d paths failed validation", len(violatingPaths))
	}

	err = ruleEngine.Finalize()
	if err != nil {
		return err
	}

	Infof("Finished validation")
	return nil
}
