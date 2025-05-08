{ pkgs, ... }: {
  # Which nixpkgs channel to use. (https://status.nixos.org/)
  channel = "stable-24.11"; # Or choose a specific Nixpkgs commit/tag

  # Use https://search.nixos.org/packages to find packages for Go development
  packages = [
    # --- Core Go Development ---
    pkgs.go # The Go compiler and runtime

    # --- Version Control ---
    pkgs.git # Essential version control system
  ];

  # Sets environment variables in the workspace
  env = { };

  # Enable Docker daemon service if you need to build/run containers
  services.docker.enable = true;

  idx = {
    # Search for extensions on https://open-vsx.org/ and use "publisher.id"
    extensions = [
      # --- Go Language Support ---
      "golang.go" # Official Go extension (debugging, testing, linting/formatting)

      # --- Version Control ---
      "GitHub.vscode-pull-request-github" # GitHub Pull Request and Issues integration
    ];

    workspace = {
      # Runs when a workspace is first created with this `dev.nix` file
      onCreate = {
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
      # Runs every time a workspace is started
      onStart = { };
    };

    # Enable previews and customize configuration if you're running web services
    previews = {
      enable = false;
    };
  };
}
