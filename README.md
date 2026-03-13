# Go Repository Template

This repository is a template for Golang projects. It includes

- workflows
- pull request template
- dependabot config
- mockery installed via `go tool`
- golangci-lint install via `go tool`
- ignore file
- Makefile

## Getting Started

- Initialize your Go module (e.g., `go mod init github.com/<org>/<repo>`)
- Update `internal/` and `test/` packages as needed. They currently contain
placeholder files to demonstrate how linting, mocking, and testing can be done.
- Update `.github/workflows/on-prs.yml` and `.github/workflows/on-schedule.yml`
as needed. For example, you may need to update the lint version, modify the
schedule, etc.
- Update `mockery` or `golangci-lint`. You may want to upgrade or downgrade the
`mockery` or `golanci-lint` tools. You can do so by:

```bash
# Usage
go get -tool -modfile=modfiles/<tool>/go.mod github.com/<org>/<repo>/cmd/<tool>@<version>

# Example
go get -tool -modfile=modfiles/golangci-lint/go.mod github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.10.1
go get -tool -modfile=modfiles/mockery/go.mod github.com/vektra/mockery/v3@v3.2.5

# You can also uninstall a tool with:
go get -tool -modfile=modfiles/mockery.mod github.com/vektra/mockery/v3@none
```
