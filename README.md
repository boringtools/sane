# Sane: Validate Repository Structure
A Git repository structure validation tool that follows a `gitignore` like syntax to define a desired repository structure. 

## TL;DR

Create a `.sane` file declaring the accepting repository structure

```
README.md
docs/*.md
docs/**/*.png
main.go
cmd/**/*.go
pkg/**/*.go
```

Run `sane` to validate the current repository structure based on `.sane` file

```bash
./sane
```

## Installation

```bash
go install https://github.com/abhisek/sane@latest
```

## Usage

Run `sane` in normal mode where it will validate any file observed in the repository against the rules in `.sane` file found in the root of the directory

```
./sane -d /path/to/dir
```

Run `sane` in strict mode to ensure that every rule in `.sane` file must have at least one match. This mode helps ensure that required files are created and available in the repository. For example, in strict mode, `sane` will fail if `README.md` is declared in `.sane` file but does not exist in the repository root.

```
./sane --strict -d /path/to/dir
```
