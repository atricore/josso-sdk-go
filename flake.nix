# JOSSO development environment: tools to work with JOSSO/IAM.tf

{
  description = "JOSSO/IAM.tf GO project flake";
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/release-23.05";
    flake-utils.url = "github:numtide/flake-utils";
    josso-pkgs.url = "github:sgonzalezoyuela/mynix-pkgs";
  };
  outputs = { self, nixpkgs, flake-utils, josso-pkgs }:
    flake-utils.lib.eachDefaultSystem
      (system:
        let

          pkgs = import nixpkgs {
            inherit  system;
          };
 

        in
        with pkgs;
        {
          devShells.default = mkShell {
            buildInputs = [
              go
              terraform
            ];
          };
        }
      );
}
