workflow "push" {
  on = "push"
  resolves = ["go test"]
}

action "go test" {
  uses = "docker://golang:1.12"
  runs = "go"
  args = "test -v ./..."
}
