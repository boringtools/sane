# Sane: Validate Repository Structure
A Git repository structure validation tool that follows a `gitignore` like syntax to define a desired repository structure. 

## TL;DR

Create a `.sane` file declaring the acceptable repository structure

```bash
README.md
docs/*.md
docs/**/*.png
main.go
cmd/**/*.go
pkg/**/*.go
```

Run `sane` to validate the current repository structure based on `.sane` file

```bash
sane validate
```

On successful validation, exit code is `0`.

### Things to note

1. Current directory is picked for validation by default
2. Target directory for validation must be a git repository
3. Validation rules are picked from `.sane` in current directory by default
4. Only git objects are validated, ignoring unstaged or git ignored paths

## Installation

```bash
go install https://github.com/abhisek/sane@latest
```

## Usage

### Normal Mode

Run `sane` in normal mode where it will validate any file observed in the repository against the rules in `.sane` file found in the root of the directory

```bash
./sane -p /path/to/dir
```

### Strict Mode

Run `sane` in strict mode to ensure that every rule in `.sane` file must have at least one match. This mode helps ensure that required files are created and available in the repository. For example, in strict mode, `sane` will fail if `README.md` is declared in `.sane` file but does not exist in the repository root.

```bash
./sane --strict -p /path/to/dir
```

### Other options

Optionally provide a rules path instead of the default `.sane`

```bash
./sane -p /path/to/dir -r /path/to/dir/rules/sane.rules
```

## FAQ

### Why couple it with a Git repository?

The rule engine is capable of validating any file path against he rules. As
such the rules are not coupled with a git repository. However, this tool seems
to be useful within the scope of a Git (or equivalent) repository while
ignoring the unstaged or ignored files.

### Why use `.gitignore` style rules?

To reduce cognitive load of learning another format and re-use conventional
syntax of `gitignore` but in an inverse context i.e. `.sane` rules define what
is acceptable instead of what should be ignored
