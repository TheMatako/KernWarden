# KernWarden

**KernWarden** is an asynchronous infrastructure monitoring daemon written in Go.
Currently in its foundational MVP stage, it is designed to act as the active probing agent for the KernOS ecosystem, but will be deployable standalone on any Kubernetes cluster or Unix environment.

## Current Architecture

The project strictly follows the Standard Go Project Layout:

- `cmd/warden/`: The future main entry point and orchestrator of the daemon.
- `internal/probes/http/`: The isolated logic for the HTTP probing module (currently under development).

## Prerequisites

- **Go** 1.23.0
- **Make** (for build automation)

## Development Workflow

We use a `Makefile` to standardize all operations. To maintain consistency, **do not run `go` commands manually**.

```bash
# Clone the repository
git clone https://github.com/TheMatako/KernWarden.git
cd KernWarden

# Format the code (enforces gofmt standards)
make fmt

# Run static analysis (catches common errors and vulnerabilities)
make vet

# Note: The 'make build' and 'make run' commands will fail until the core entrypoint is implemented in the upcoming roadmap phases.
```

## Roadmap

This project is currently being bootstrapped.

- [x] Establish project governance, tooling, and strict repository structure.
- [x] Implement the core HTTP probe logic (`internal/probes/http`).
- [ ] Create the main orchestrator and CLI (`cmd/warden`).
- [ ] Add CI/CD pipelines via GitHub Actions for automated testing.
- [ ] Containerize the application (Docker).

## Governance & Contributing

We enforce strict software engineering standards. Before opening a Pull Request, you **must** read the [CONTRIBUTING.md](CONTRIBUTING.md) to understand our language policy, Go conventions, and Conventional Commits requirements.

## License

Distributed under the MIT License. See `LICENSE` for more information.
