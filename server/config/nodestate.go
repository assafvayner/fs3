package config

import (
	"fmt"
  "os"
)

var PRIMARY_PORT = 5000
var BACKUP_PORT = 50000

func IsPrimary() (bool) {
  if len(os.Args) < 2 {
    fmt.Fprintln(os.Stderr, "Not enough params")
    os.Exit(1)
  }

  pbArg := os.Args[1]

  if pbArg == "primary" {
    return true
  }

  if pbArg != "backup" {
    fmt.Fprintln(os.Stderr, "Bad primary/backup param")
    os.Exit(1)
  }

  return false
}

func IsDev() (bool) {
  // no arg provided assume prod
  if len(os.Args) < 3 {
    return false
  }
  stageArg := os.Args[2]
  if stageArg == "dev" {
    return true
  }
  if stageArg != "prod" {
    fmt.Fprintln(os.Stderr, "Bad dev/prod param")
    os.Exit(1)
  }
  return false
}
