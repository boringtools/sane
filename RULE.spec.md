# Sane Rules Format

The `.sane` rules format (pattern) is syntactically similar to [.gitignore](https://git-scm.com/docs/gitignore)
but semantically opposite.

## Source

* Only a single `.sane` file is supported

This behaviour is unlike `.gitignore` file which supports per directory config
scoped within the directory and its sub-directories.

The rationale for choosing a single `.sane` file is to maintain a single source
of truth for an acceptable directory structure.

Following files are ignored while matching

```text
/.gitignore
/.sane
```

## Pattern

The patterns in `.sane` file is syntactically similar to `.gitignore` however
it is semantically opposite. While `.gitignore` is used to ignore certain files
from being staged into a git repository, `.sane` patterns are used to allow or
explicitly declare the required files or directory structure in a git
repository.

## Matching

* A rule pattern match indicates a file is allowed
* A negated rule pattern match indicates a file is not allowed
* First match criteria is used while matching multiple rules

To allow all markdown files in `docs/*.md` except `docs/README.md`

```text
!/docs/README.md
/docs/*.md
```

The order is required to explicitly disallow `docs/README.md` within first
match constraints.

* Glob rules are supported e.g. `/docs/**/*.md`
* Only blob objects (files) in a git repository are matched
* Tree objects (directories) are not matched

To prevent creation of directories

```text
!/docs/**/*
```
