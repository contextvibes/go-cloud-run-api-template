# Changelog for Go Starter Templates for Firebase Studio

All notable changes to this project (the collection of templates and the overall template system) will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.0.1] - 2025-05-08

### Added
- Initial project setup for Go Starter Templates for Firebase Studio.
- Root `idx-template.json` and `idx-template.nix` for template selection.
- First template: `cloud-run-api` (Go Hello World API for Cloud Run).
  - Includes Go application code, Dockerfile, and Nix environment (`cloud-run-api/.idx/dev.nix`).
  - Generic module path (`your-module-name`) and dependency on `dui-go` established for the template.
- Root `README.md`, `CONTRIBUTING.md`, `LICENSE`, `ROADMAP.md`, and this `CHANGELOG.md`.

---
<!--
Template for future releases:

## [X.Y.Z] - YYYY-MM-DD

### Added
- New template: `template-name` ([Link to template PR or subdir](./template-name/))
- Feature X for `existing-template`.

### Changed
- Improved Y in `existing-template`.
- Updated root `idx-template.json` for new options.

### Fixed
- Bug Z in `existing-template`.

### Removed
- Deprecated template `old-template-name`.

-->