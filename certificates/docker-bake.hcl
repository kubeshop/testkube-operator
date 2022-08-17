// docker-bake.hcl
target "docker-metadata-action" {}

target "build" {
  inherits = ["docker-metadata-action"]
  context = "./"
  dockerfile = "./certificates/Dockerfile"
  platforms = [
    "linux/amd64",
    "linux/arm64"
  ]
}
