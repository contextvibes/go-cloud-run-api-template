# Root idx-template.nix for multi-variant Go templates
{ pkgs, environment ? "cloud-run-api", ... }: {
  packages = [ pkgs.bash ];

  bootstrap = ''
    set -ex

    echo "Bootstrapping selected project type: ${environment}"
    echo "Source directory to copy: ${./.}/${environment}"
    echo "Target workspace name (from env): $WS_NAME"
    echo "Final output directory (from Nix): $out"

    if [ ! -d "${./.}/${environment}" ]; then
      echo "CRITICAL ERROR: Source directory '${./.}/${environment}' does not exist for the selected project type."
      exit 1
    fi

    cp -rf "${./.}/${environment}" "$WS_NAME"
    chmod -R +w "$WS_NAME"
    mv "$WS_NAME" "$out"

    echo "Bootstrapping complete for project type: ${environment}."
    echo "Workspace content is now in: $out"
  '';
}
