package main

import (
	"day29/user"
	"day29/utils"
)

func main() {

	user.NewUser().Check()

	utils.NewFamilyAccount().MainMenu()

}
