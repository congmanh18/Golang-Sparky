package main

import (
	"basic/user-service/repository"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	var userRepo = repository.NewUserRepo()
	var user = userRepo.GetUserByUID("100")
	spew.Dump(user)
}
