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

