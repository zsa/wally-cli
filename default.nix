{ pkgs ? (import <nixpkgs> { }), lib ? (import <nixpkgs/lib>), system ? builtins.currentSystem }:

assert lib.versionAtLeast pkgs.go.version "1.14";

pkgs.buildGoModule rec {
  name = "wally-cli";
  version = "v2.0.0";

  src = ./.;

  vendorSha256 = "m2QuNd0/cfAdFdVzctG+E7t/OsslcufXyh6HX2i1KKg=";

  subPackages = [ "." ];

  buildInputs = with pkgs; [ libusb1 ];
  nativeBuildInputs = with pkgs; [ pkg-config ];

  meta = with lib; {
    description = "Flash your ZSA Keyboard the EZ way.";
    homepage = "https://github.com/zsa/wally-cli";
    license = licenses.mit;
    maintainers = [ johnrichardrinehart ];
    platforms = platforms.linux ++ platforms.darwin;
  };
}
