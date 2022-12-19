//go:build mage

package main

import (
  "github.com/magefile/mage/mg"
  "github.com/magefile/mage/sh"
  "fmt"
)

type (
  Run mg.Namespace
  Migrate mg.Namespace
  Generate mg.Namespace
)

func (Run) CreditDB() error {
   if err := sh.RunV("docker", "build", "-t", "carbx/creditdb", "-f", "./db/credit/Dockerfile", "./db/credit/"); err != nil {
    return err
  }
  
  sh.RunV("docker", "container", "stop", "carbx_credit_db")
  sh.RunV("docker", "container", "rm", "carbx_credit_db")

  return sh.RunV("docker", "run", "-p", "5432:5432", "-d", "--name=carbx_credit_db", "carbx/creditdb")
}

func (Migrate) CreditDB() error {
  return sh.RunV("migrate", "-source", "file://db/credit/migrations", "-database", "postgres://postgres:asdfasdf@localhost:5432/postgres?sslmode=disable", "up")
}

func (Generate) Api(name string) error {
  return sh.RunV("protoc", "--go_out=.", "--go-grpc_out=.", fmt.Sprintf("./app/api/proto/%s.proto", name))
}
