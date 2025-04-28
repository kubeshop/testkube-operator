variable "GOCACHE"       { default = "/go/pkg" }
variable "GOMODCACHE"    { default = "/root/.cache/go-build" }
variable "ALPINE_IMAGE"  { default = "alpine:3.20.3" }

group "default" {
  targets = ["operator"]
}

group "debug" {
  targets = ["operator-debug"]
}

target "operator-meta" {}
target "operator" {
  inherits = ["operator-meta"]
  context="."
  dockerfile = "dev.Dockerfile"
  platforms = ["linux/arm64", "linux/amd64"]
  args = {
    GOCACHE = "${GOCACHE}"
    GOMODCACHE = "${GOMODCACHE}"
    ALPINE_IMAGE = "${ALPINE_IMAGE}"
  }
}

target "operator-debug" {
  inherits = ["operator"]
  target = "debug"
  args = {
    SKAFFOLD_GO_GCFLAGS = "all=-N -l"
  }
}