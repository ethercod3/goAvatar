#!/usr/bin/env nu

const app = "goavatar"
const package = "./cmd/goavatar"

let targets = [
  { os: "linux", arch: "amd64", ext: "" },
  { os: "linux", arch: "arm64", ext: "" },
  { os: "darwin", arch: "amd64", ext: "" },
  { os: "darwin", arch: "arm64", ext: "" },
  { os: "windows", arch: "amd64", ext: ".exe" },
]

mkdir dist

for target in $targets {
  let output = $"dist/($app)-($target.os)-($target.arch)($target.ext)"
  with-env { GOOS: $target.os, GOARCH: $target.arch, CGO_ENABLED: "0" } {
    go build -trimpath -ldflags="-s -w" -o $output $package
  }
}
