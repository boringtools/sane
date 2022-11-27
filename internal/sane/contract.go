package sane

type RuleFormat string

const (
	RULES_FORMAT_GITIGNORE = "gitignore"
)

type Config struct {
	RepositoryPath string
	RulesPath      string
	RulesType      RuleFormat
	FailFast       bool
	Strict         bool
}

type RuleEngine interface {
	// Validate a single repository node with the loaded rules
	// This validation should be performed in a stateless manner
	Validate(RepositoryNode) (bool, error)

	// Finalize the validation state (if any) and return error (if any)
	Finalize() error

	// Reset validatin state
	Reset()
}

type RepositoryNode struct {
	FullPath string
}

type RepositoryNodeHandler func(RepositoryNode) error

type RepositoryWalker interface {
	Walk(RepositoryNodeHandler) error
}
