# Contributing to Go Starter Templates for Firebase Studio

Thank you for considering contributing to the Go Starter Templates for Firebase Studio! We welcome improvements, bug fixes, new template ideas, and enhancements that help developers get started quickly with Go projects in Firebase Studio.

## Community Guidelines

We strive to maintain a positive and welcoming environment. All participants are expected to act professionally and respectfully toward others. Please see the Code of Conduct section in the main `README.md`.

## Getting Started

1.  **Prerequisites:**
    *   Ensure you have Go (see the `go.mod` in the specific template you're working on, e.g., `cloud-run-api/go.mod`, or the latest stable Go version for new templates).
    *   Git installed and configured.
    *   Access to [Firebase Studio](https://idx.dev/) is highly recommended for testing template behavior and Nix environments.
    *   Familiarity with Nix is helpful for modifying `.idx/dev.nix` files.

2.  **Fork & Clone:**
    *   Fork the `go-cloud-run-api-template` repository on GitHub: `https://github.com/contextvibes/go-cloud-run-api-template`
    *   Clone your fork locally:
        ```bash
        git clone https://github.com/YOUR_USERNAME/go-cloud-run-api-template.git
        cd go-cloud-run-api-template
        ```

## Making Changes

1.  **Create a Branch:** Before making changes, create a new branch from the `main` branch:
    ```bash
    git checkout main
    git pull origin main # Ensure your main is up-to-date
    git checkout -b feature/your-template-feature # Example: feature/add-grpc-template
    # or
    git checkout -b fix/improve-cloud-run-readme # Example: fix/update-cloud-run-readme
    ```

2.  **Implement Your Changes:**
    *   **Improving Existing Templates (e.g., `cloud-run-api`):**
        *   Make your code changes within the template's subdirectory (e.g., `cloud-run-api/`).
        *   Update documentation (`README.md`, `CHANGELOG.md` if it's a user-facing template changelog) within that subdirectory.
        *   If modifying the Nix environment, update the relevant `.idx/dev.nix` file.
        *   Test your changes thoroughly, ideally within Firebase Studio.
    *   **Adding New Templates:**
        *   Create a new subdirectory for your template (e.g., `my-new-template/`).
        *   Include all necessary files: Go code, `Dockerfile` (if applicable), `.idx/dev.nix`, `README.md`, a user-template `CHANGELOG.md`, etc.
        *   Update the root `idx-template.json` to include your new template as an option in the `params.options` list.
        *   Update the root `idx-template.nix` if any special bootstrapping logic is needed for your template type (though often, just copying the directory is fine).
        *   Add a mention of your new template in the root `README.md` and `ROADMAP.md`.
    *   **General Project Improvements:**
        *   Changes to root files like `README.md`, `ROADMAP.md`, or the root Nix/JSON template files.

3.  **Follow Style:**
    *   **Go Code:** Adhere to standard Go formatting (`go fmt ./...`) and linting (`go vet ./...`, `golangci-lint run ./...` if available in the template's dev environment).
    *   **Documentation:** Keep Markdown clear and well-formatted.

4.  **Test:**
    *   **For Template Code:** Run any tests included with the template (e.g., `(cd cloud-run-api && go test ./...)`).
    *   **Firebase Studio Testing:** The most crucial step is to test the template in Firebase Studio:
        1.  Push your branch to your fork.
        2.  In Firebase Studio, create a new workspace from a Git repository, using the URL of *your forked repository and specific branch*.
        3.  Select the template (if multiple) and ensure it initializes correctly.
        4.  Test the development environment, build process, and any specific features of the template.

5.  **Commit:** Commit your changes using clear and descriptive commit messages. Consider following the [Conventional Commits](https://www.conventionalcommits.org/) specification (e.g., `feat(template): Add gRPC service template`, `fix(cloud-run): Correct Dockerfile instruction`, `docs(readme): Clarify setup steps`).
    ```bash
    git add .
    git commit -m "feat(cloud-run): Add graceful shutdown to API server"
    ```

## Submitting a Pull Request

1.  **Push:** Push your feature or fix branch to your fork on GitHub:
    ```bash
    git push origin feature/your-template-feature
    ```
2.  **Open PR:** Go to the original `contextvibes/go-cloud-run-api-template` repository on GitHub. GitHub should automatically detect your pushed branch and prompt you to create a Pull Request against the `main` branch.
3.  **Describe:** Fill out the Pull Request template, clearly describing the problem you're solving or the template you're adding/improving. Link to any relevant issues from the `ROADMAP.md` or issue tracker.
4.  **Review:** Respond to any feedback or code review comments.

## Finding Ways to Contribute

For a detailed list of planned new templates, enhancements to existing ones, or other ideas, please see our [ROADMAP.md](ROADMAP.md). We welcome contributions to items listed there or your own suggestions that align with the project's goals!

Thank you for contributing to the Go Starter Templates for Firebase Studio!