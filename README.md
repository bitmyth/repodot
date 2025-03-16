# Repodot
To initialize a Git project with commonly used configuration files such as:

- `.gitignore`
- `.dockerignore`
- `Makefile`
- `Dockerfile`

## Install

```bash
go install github.com/bitmyth/repodot@latest
```

---

### Explanation

1. **Configuration Files**:
   - `.gitignore`: Specifies files and directories to be ignored by Git.
   - `.dockerignore`: Specifies files and directories to be ignored by Docker during image builds.
   - `Makefile`: Provides a set of commands for building, testing, and deploying the project.
   - `Dockerfile`: Defines the steps to build a Docker image for the project.

2. **Installation Command**:
   - The `go install` command installs the `repodot` tool, which helps automate the creation of these configuration files.

---

### Usage

After installing `repodot`, you can use it to generate the necessary configuration files for your project, ensuring a consistent and efficient setup.
