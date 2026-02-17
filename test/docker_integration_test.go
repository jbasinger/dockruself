package docker_integration

import (
  "context"
  "testing"

  "github.com/docker/docker/pkg/stdcopy"
  "github.com/docker/go-sdk/container"
  "github.com/docker/go-sdk/container/wait"
)

func TestContainerLog(t *testing.T) {
  ctr, err := container.Run(
   context.Background(),
   container.WithImage("debian:latest"),
   container.WithCmd("echo", "hello world"),
   container.WithWaitStrategy(wait.ForLog("hello world"))
  )
  if err != nil { t.Fatalf("Error running container\n:%v", err); }


  logs, err := container.Logs(context.Background())
  if err != nil { t.Fatalf("Error obtaining logs\n:%v", err); }
}
