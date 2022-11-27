package sane

import "fmt"

// Primary entrypoint for validator
func Execute(config Config) error {
	Debugf("Executing sane for repository validation")
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

	Debugf("Starting repository walk")
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
		for _, path := range violatingPaths {
			Infof("Failed: %s", path)
		}

		return fmt.Errorf("%d paths failed validation", len(violatingPaths))
	}

	Debugf("Finalizing rule engine")
	err = ruleEngine.Finalize()
	if err != nil {
		return err
	}

	Debugf("Sane execution completed successfully")
	return nil
}
