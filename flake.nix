{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/release-20.09";
    flake-utils.url = "github:numtide/flake-utils";
  };


  outputs = inputs:
    let
      inherit (inputs.flake-utils.lib) eachDefaultSystem flattenTree mkApp;
    in
    eachDefaultSystem (system:
      let
        pkgs = inputs.nixpkgs.legacyPackages.${system};
        lib = inputs.nixpkgs.lib;
      in
      rec {
        defaultPackage = (import ./default.nix) {
          inherit pkgs lib system;
        };

        defaultApp = defaultPackage;
      }
    );
}
