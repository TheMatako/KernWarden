# Contributing to KernWarden

We strictly enforce software engineering standards to maintain the security, traceability, and readability of the codebase. This is a Go-based infrastructure tooling project. Therefore, performance and code hygiene are non-negotiable.

## 1. Language Policy

**ALL code, variable names, documentation, and comments MUST be written in English.** Pull requests or commits containing any other language will be automatically rejected by the reviewers.

## 2. Development Workflow & Rules

- **No Direct Push**: Pushing directly to the `main` branch is strictly forbidden and blocked by repository rules.
- **Mandatory Pull Requests**: All changes must be submitted via a Pull Request (PR) from a dedicated branch.
- **CI Validation**: No PR will be merged if the Continuous Integration (CI) pipeline (build, lint, test) fails.
- **Squash and Merge**: We use the "Squash and Merge" strategy to keep the `main` history clean, linear, and easily auditable.

## 3. Tooling & Go Standards

Use the provided `Makefile` for all development tasks. Manual execution of Go tools is discouraged to ensure consistency across environments.

- `make fmt`: Format the code according to Go standards. Code that is not `gofmt` compliant will fail the CI.
- `make vet`: Run static analysis to catch common errors.
- `make test`: Execute the test suite. **All new features must be covered by unit tests.**
- `make build`: Verify that the project compiles into a standalone binary.

## 4. Commit Convention

We follow the strict [Conventional Commits](https://www.conventionalcommits.org/) standard. This allows us to automate versioning and changelog generation.

**Format:** `type(scope): short description`

**Allowed Types:**

- `feat`: A new feature (e.g., a new probe)
- `fix`: A bug fix
- `docs`: Documentation only changes
- `chore`: Maintenance tasks, dependency updates, CI/CD changes
- `refactor`: A code change that neither fixes a bug nor adds a feature
- `test`: Adding missing tests or correcting existing tests

**Examples:**

- ✅ `feat(http): implement concurrent probing logic` *(Correct)*
- ✅ `fix(core): resolve memory leak in client timeout` *(Correct)*
- ❌ `feat/http: add new probe` *(Wrong: do not use slashes)*
- ❌ `added http probe` *(Wrong: missing type and scope)*

## 5. Pull Request Convention

Since we use the *Squash and Merge* strategy, your PR title will become the final commit message on the `main` branch. It must follow the exact same Conventional Commits format as above.

### PR Title Example

`feat(dns): add latency tracking to DNS resolution`

### PR Description Template

Your PR description must clearly explain the intent and the testing methodology. Use the following structure:

```text
### Problem / Goal
[Explain briefly what issue this PR solves or what feature it adds. Link to any existing issue, e.g., "Resolves #12"]

### Solution
[Explain the technical approach you took. e.g., "Implemented a custom http.Transport to force a 5-second timeout on all requests."]

### Testing
- [ ] Code passes `make vet` and `make fmt`
- [ ] Unit tests were added/updated
- [ ] Code compiles successfully via `make build`
