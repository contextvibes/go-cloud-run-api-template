# cloud-run-api/.idx/dev.nix
# Nix configuration for the Go Cloud Run API Starter template.
# This file defines the development environment for workspaces created from this template.
# To learn more about how to use Nix to configure your environment
# see: https://developers.google.com/idx/guides/customize-idx-env

{ pkgs, ... }: {
  # Which nixpkgs channel to use.
  channel = "stable-24.11"; # Or choose a specific Nixpkgs commit/tag

  # System packages available in the workspace.
  # Use https://search.nixos.org/packages to find packages.
  packages = [
    # --- Core Go Development ---
    pkgs.go # The Go compiler and runtime (matches go.mod version as closely as possible via channel)
    pkgs.gopls # Go Language Server (for editor features)
    pkgs.delve # Go Debugger (Essential for step debugging)
    pkgs.golangci-lint # Fast Go linters runner / aggregator (Enabled)
    pkgs.goimports-reviser # Tool to format and revise Go imports
    pkgs.gotools # Collection of Go analysis tools (guru, gorename, etc.)
    # pkgs.gomodifytags    # Optional: Tool for managing struct tags

    # --- Protocol Buffers & gRPC/Connect ---
    # Commented out as not currently used by this Hello World API template
    # pkgs.protobuf # Protocol Buffers compiler (protoc)
    # pkgs.protoc-gen-go # Protoc plugin for Go code generation
    # pkgs.protoc-gen-go-grpc # Protoc plugin for Go gRPC code generation
    # pkgs.protoc-gen-connect-go # Protoc plugin for Go Connect RPC

    # --- Version Control ---
    pkgs.git # Essential version control system

    # --- Utilities ---
    pkgs.patch # Standard patching utility
    pkgs.jq # Command-line JSON processor
    pkgs.tree # Directory structure viewer
    # pkgs.k6              # Optional: Load testing tool (if needed)
    pkgs.google-cloud-sdk # Optional but often useful: gcloud CLI, gsutil, etc.
  ];

  # Sets environment variables in the workspace.
  env = {
    # Example: Set GOPRIVATE for private Go modules
    # GOPRIVATE = "github.com/your-org/*";
  };

  # Enable Docker daemon service if you need to build/run containers.
  services.docker.enable = true;

  # Firebase Studio specific configurations.
  idx = {
    # VS Code extensions to install.
    # Search for extensions on https://open-vsx.org/ and use "publisher.id"
    extensions = [
      # --- Go Language Support ---
      "golang.go" # Official Go extension (debugging, testing, linting/formatting)

      # --- Version Control ---
      "GitHub.vscode-pull-request-github" # GitHub Pull Request and Issues integration

      # --- Other Useful Extensions ---
      "ms-azuretools.vscode-docker" # Docker integration
      "EditorConfig.EditorConfig"   # For maintaining consistent coding styles
      # "bierner.markdown-preview-github-styles" # For better Markdown previews
    ];

    # Workspace lifecycle hooks.
    workspace = {
      # Runs when a workspace is first created with this file.
      onCreate = {
        # Open these files by default when the workspace is created.
        # The last file in the list will be focused.
        default.openFiles = [
          "README.md"
          ".env" # Show the user the environment variables they configured (if .env is created by another hook)
          ".env.example" # Or open the example if .env is not auto-created
          "cmd/main.go"
          "internal/config/config.go"
          "internal/api/handlers.go" # Key handler logic
        ];

        # Install Go module dependencies and format/tidy code.
        installAndTidy = ''
          echo "Running go mod download and go mod tidy..."
          go mod download
          go mod tidy
          echo "Running go fmt..."
          go fmt ./...
          echo "Initial Go project setup commands complete."
        '';

        # Optional: Script to copy .env.example to .env if .env doesn't exist
        setupEnvFile = ''
          if [ ! -f ".env" ] && [ -f ".env.example" ]; then
            echo "Copying .env.example to .env..."
            cp .env.example .env
            echo "INFO: .env file created. Please review and update it with your specific values (especially GOOGLE_CLOUD_PROJECT)."
          fi
        '';
        
        # Script to install contextvibes CLI into ./bin
        installContextVibesCli = ''
          echo "Attempting to install contextvibes CLI into ./bin ..."

          if ! command -v go &> /dev/null
          then
              echo "Go command could not be found, skipping contextvibes installation."
              # Exit gracefully or 'exit 1' if critical
              # For now, we'll assume Go is present due to pkgs.go
          else
            LOCAL_BIN_DIR="$(pwd)/bin"
            mkdir -p "$LOCAL_BIN_DIR"
            echo "Target directory for contextvibes: $LOCAL_BIN_DIR"

            export GOBIN="$LOCAL_BIN_DIR"
            echo "Running: GOBIN=$GOBIN go install github.com/contextvibes/cli/cmd/contextvibes@latest"

            if go install github.com/contextvibes/cli/cmd/contextvibes@latest; then
              echo "Successfully installed contextvibes to $LOCAL_BIN_DIR/contextvibes"
              echo "You can run it using: $LOCAL_BIN_DIR/contextvibes"
              echo "Consider adding '$LOCAL_BIN_DIR' to your PATH for convenience (see README)."
              chmod +x "$LOCAL_BIN_DIR/contextvibes" || echo "Note: chmod +x on contextvibes failed."
            else
              echo "ERROR: Failed to install contextvibes."
            fi
            unset GOBIN
          fi
        '';
      };

      # Runs every time a workspace is started (or restarted).
      onStart = {
        # Example: Check tool versions or display a welcome message.
        checkGoVersion = "go version";
        # displayWelcome = "echo 'Welcome back to your Go Cloud Run API project!'";
      };
    };

    # Configure web previews if your application serves HTTP.
    previews = {
      enable = true; 
    };

    # (Optional) Default icon for workspaces created from this template.
    # Place an 'icon.png' (e.g., 128x128 or 256x256) in this .idx directory.
    # Example: .idx/icon.png
  };
}