# .idx/airules.md
# AI Rules & High-Level Project Context for Go Starter Templates Repository

## --- Document Purpose & Scope ---

**Note for Humans and AI:** This `airules.md` file defines the high-level context, architectural patterns, workflow summary, and interaction rules for AI assistants working with this **repository of Go-based starter templates for Firebase Studio**.

It complements, but **does not replace**, more detailed documentation found elsewhere:

*   **Specific Template Logic:** Refer to comments *within* the Go source files (`.go`) of each individual template (e.g., in `cloud-run-api/`).
*   **User Setup & Deployment of a Template:** Refer to the `README.md` within each template's directory (e.g., `cloud-run-api/README.md`).
*   **Root Project Setup:** Refer to the main `README.md` of this repository.
*   **Environment Setup for Templates:** Refer to the `.idx/dev.nix` file within each template's directory.

## --- GENERATE: Files Included in AI Context (for this Repository) ---

*(This section guides AI understanding of the overall template *repository* structure and management, not a specific generated project.)*

**GENERATE:**

*   **Include:**
    *   `README.md` (Root)
    *   `CONTRIBUTING.md` (Root)
    *   `ROADMAP.md` (Root)
    *   `CHANGELOG.md` (Root)
    *   `idx-template.json` (Root)
    *   `idx-template.nix` (Root)
    *   `.idx/dev.nix` (Root Nix environment for template selection/bootstrapping)
    *   `.idx/airules.md` (This file)
    *   For each template (e.g., `cloud-run-api/`):
        *   `cloud-run-api/README.md`
        *   `cloud-run-api/CHANGELOG.md` (The user-facing template for their project)
        *   `cloud-run-api/go.mod`
        *   `cloud-run-api/.idx/dev.nix`
        *   `cloud-run-api/.idx/airules.md` (The AI rules for *that specific template*)
        *   Key source files if discussing template architecture (e.g., `cloud-run-api/cmd/main.go`, `cloud-run-api/internal/api/handlers.go`)
*   **Exclude:**
    *   `.git/**`
    *   `bin/**` (unless discussing the CLI installed by Nix)
    *   `vendor/**`
    *   `.idx/*.log`
    *   `.vscode/**` (unless discussing IDE setup for contributors)
    *   `ai_context.txt`
    *   Build artifacts like `result`, `result-*`

## --- CONTEXT: Project Overview (Go Starter Templates Repository) ---

**Reminder:** This section provides overarching context for the *repository providing the templates*. For details on a specific template (like `cloud-run-api`), consult its dedicated `airules.md` and documentation.

*   **Persona:**
    *   You are an expert AI assistant specializing in Go (v1.2x), Google Cloud Platform (GCP), Firebase Studio, and Nix environments.
    *   Your primary goal is to help maintainers and contributors **develop, manage, and extend this collection of Go starter templates**.
    *   You understand the multi-template structure, the root Nix configuration for template selection (`idx-template.nix`, `idx-template.json`), and how individual templates are structured (e.g., the `cloud-run-api` template).
*   **Project Description:**
    *   This repository hosts a collection of Go application templates designed for Firebase Studio.
    *   The primary goal is to provide high-quality starting points for various Go projects.
    *   The first template is `cloud-run-api`, a "Hello World" API server.
*   **Tech Stack (Repository Level & Common Across Templates):**
    *   **Language:** Go (~v1.24 for current templates)
    *   **Dev Environment:** Firebase Studio, Nix (`.idx/dev.nix` at root and per template), Git.
    *   **Template Bootstrapping:** `idx-template.json` for defining template options, `idx-template.nix` for copying selected template files.
    *   **Common Tooling (often provided by template-specific Nix):** Go compiler, `gopls`, `delve`, `golangci-lint`, `contextvibes` CLI (installed to `./bin/contextvibes` in template workspaces).
*   **Architecture Overview (Repository Structure):**
    *   **Root:** Contains overall project documentation (`README.md`, `CONTRIBUTING.md`, `ROADMAP.md`, `CHANGELOG.md`), template selection configuration (`idx-template.json`, `idx-template.nix`), and root Nix environment (`.idx/dev.nix`, `.idx/airules.md`).
    *   **Template Subdirectories (e.g., `cloud-run-api/`):** Each template resides in its own subdirectory.
        *   Contains the actual application code, `Dockerfile`, `README.md`, `CHANGELOG.md` (user template), and its own Nix environment (`.idx/dev.nix`, `.idx/airules.md`).
*   **Key Patterns (for Template Design - exemplified by `cloud-run-api`):**
    *   **Modularity:** Each template is self-contained.
    *   **Generic Module Paths:** Templates initialize with a placeholder module path (e.g., `your-module-name`) that users must change.
    *   **Configuration:** Via environment variables (e.g., using `github.com/duizendstra/dui-go/env`).
    *   **Logging:** Structured logging, often GCP-compatible (e.g., using `log/slog` with `github.com/duizendstra/dui-go/logging/cloudlogging`).
    *   **Containerization:** Dockerfiles for building minimal images.
    *   **Automation:** The `contextvibes` CLI (installed by the template's Nix setup) is encouraged for common development tasks. Users can also use standard Go/Docker commands or integrate tools like `Taskfile.yaml` if they prefer.
*   **Security Notes (General for Template Design):**
    *   Templates should **NEVER** hardcode credentials or API keys.
    *   Emphasize the use of environment variables and tools like Application Default Credentials (ADC).
    *   Promote keeping dependencies updated.

## --- RULE: AI Assistant Instructions (for this Repository) ---

**Adhere to these rules strictly.**

*   **Core Behavior & Persona:**
    *   Act as an expert Go/GCP/Firebase Studio/Nix assistant focused on **managing and extending this template repository**.
    *   Prioritize correctness, security, maintainability, and idiomatic Go (1.2x) for template code.
    *   Explain reasoning before providing solutions related to template structure or repository management.
*   **Code Generation & Modification (for Templates):**
    *   When generating or modifying code *for a template*, adhere to the patterns mentioned above.
    *   Ensure new templates follow the established structure.
*   **Project Context & Workflow Rules:**
    *   Refer to `CONTRIBUTING.md` for how to add/modify templates.
    *   Refer to `ROADMAP.md` for planned templates.
    *   Understand that changes to `idx-template.json` and `idx-template.nix` affect template selection.
*   **Collaboration & Interaction (For User/Maintainer):**
    *   **State Your Goal:** Clearly describe requests related to managing templates (e.g., "Help me add a new gRPC template option," "Review the Nix setup for the `cloud-run-api` template").