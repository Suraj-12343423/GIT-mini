# Go Mini Git

A simplified implementation of Git built from scratch using Go to understand how version control systems work internally.

## About the Project

**Go Mini Git** is a learning project that recreates core Git concepts, including commits, object storage, commit history, and branch references.

Instead of relying on Git internally for these features, the goal is to implement the underlying mechanisms and understand how they work.

## Current Features

- [x] Repository initialization
- [x] Commit creation
- [x] Commit history using `log`
- [x] Branch reference management
- [ ] Content-addressable object storage
- [ ] File tracking with `add`
- [ ] Working tree status with `status`
- [ ] Checkout commits
- [ ] Branch creation and switching
- [ ] Merge and conflict handling

*Feature status reflects the project's current development stage and will be updated as implementation progresses.*

## Project Structure

```text
mini-git/
├── .gitignore
├── go.mod
├── main.go
└── .mini-git/        # Local Mini Git data (ignored by Git)
    ├── objects/
    └── refs/
        └── heads/
            └── main
```

The `.mini-git/` directory is used by the implementation to store its internal repository data. It is excluded from the source repository so that local test data is not uploaded to GitHub.

## Technologies

- **Language:** Go
- **Version Control:** Git and GitHub
- **Hashing:** SHA-256
- **Storage:** Local filesystem

## Getting Started

### Prerequisites

- Go installed on your system
- Git installed on your system

### Clone the Repository

```bash
git clone https://github.com/Suraj-12343423/GIT-mini.git
cd GIT-mini
```

### Run the Project

```bash
go run .
```

The available commands depend on the current implementation.

## Concepts I'm Learning

- How Git stores objects
- Content-addressable storage
- Hashing and object identification
- Commit objects and parent references
- Commit history traversal
- Branch references
- File tracking and working tree changes
- Checkout, branching, and merging

## Roadmap

1. Refine commit creation and initial commit handling.
2. Implement reliable object storage and retrieval.
3. Add file tracking and status detection.
4. Implement checkout.
5. Add branch creation and switching.
6. Implement merging and conflict handling.
7. Add tests and improve error handling.

## Project Goal

Build a working miniature version control system while developing a deeper understanding of Go, filesystem operations, data structures, hashing, and Git internals.

---

**Status:** Work in progress

**Language:** Go

**Author:** [Suraj-12343423](https://github.com/Suraj-12343423)
