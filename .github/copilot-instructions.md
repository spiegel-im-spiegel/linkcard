# Copilot Instructions for linkcard

## Project Overview

- This repository provides a small CLI for Hugo link card workflows.
- Main focus areas are link metadata fetching, image handling, and JSON output/data persistence.
- Keep changes practical, incremental, and easy to review.

## Coding Guidelines

- Language: Go.
- Prefer small, focused changes over broad refactors.
- Preserve existing public behavior unless explicitly asked to change it.
- Keep source-code comments in English.
- Follow existing package boundaries and naming style in this repository.

## Validation

- Use Taskfile tasks for local validation.
- Primary check command:

```sh
task test
```

- If needed, run additional checks:

```sh
task govulncheck
```

- For routine code changes, prefer this full check:

```sh
task
```

## Dependencies

- Do not add new external dependencies unless there is a clear benefit.
- If a dependency is added, explain why in the change summary.

## Documentation and Maintenance

- Update related documentation when behavior, options, or workflows change.
- Keep README and package comments consistent with the current implementation.
- Prefer explicit notes about assumptions and constraints when they are not obvious from code.

## Current Maintenance Context

- During this phase, prioritize output stability and risk reduction over feature expansion.
- Treat URL fetch behavior, generated JSON format, and image output paths as compatibility-sensitive.

## Pull Request Workflow

- Start from up-to-date main and create a dedicated working branch.
- Keep commits focused and easy to review (split docs-only and behavior changes when practical).
- Run local checks before pushing:

```sh
task
```

- Push the branch and open a PR against main.
- Confirm all GitHub Actions checks are green before merging.
- After merge, delete both remote and local working branches.

## Release Tag Workflow

- Ensure local main is clean and synchronized with origin/main.
- Create an annotated release tag:

```sh
git tag -a vX.Y.Z -m "Release vX.Y.Z"
```

- Push the tag:

```sh
git push origin vX.Y.Z
```

- Confirm build workflow is triggered by the tag push.
- After build completion, verify GitHub Release publication:
  - release notes are generated
  - binary assets are uploaded
  - checksum file is uploaded
