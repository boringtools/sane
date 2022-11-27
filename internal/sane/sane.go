package sane

import "fmt"

// Primary entrypoint for validator
func Execute(config SaneConfig) error {
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
				return fmt.Errorf("no rules match for: %s", node.FullPath)
			} else {
				violatingPaths = append(violatingPaths, node.FullPath)
			}
		}

		return nil
	})

	if err != nil {
		return nil
	}

	if len(violatingPaths) > 0 {
		return fmt.Errorf("%d paths failed validation", len(violatingPaths))
	}

	if config.Strict {
		err := ruleEngine.Finalize()
		if err != nil {
			return err
		}
	}

	return nil
}
