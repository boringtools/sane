package sane

type RuleFormat string

const (
	RULES_FORMAT_GITIGNORE = "gitignore"
)

type SaneConfig struct {
	RepositoryPath string
	RulesPath      string
	RulesType      RuleFormat
	FailFast       bool
	Strict         bool
}

type Rule interface {
	Validate(RepositoryNode) (bool, error)
	Finalize() error
}

type RepositoryNode struct {
	FullPath string
}

type RepositoryNodeHandler func(RepositoryNode) error

type RepositoryWalker interface {
	Walk(RepositoryNodeHandler) error
}
