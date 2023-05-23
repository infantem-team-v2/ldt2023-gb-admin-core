package main

import (
	"fmt"
	"gb-auth-gate/internal/pkg/server"
)

// @title Core backend app for Leaders of Digital Transformation
// @description Main service that works with each other and summarizing data
// @version 1.0.0
// @contact.name Docs developer
// @contact.url https://t.me/KlenoviySirop
// @contact.email KlenoviySir@yandex.ru

// @host gate.gb.ldt2023.infantem.tech
// @schemes https

// @securityDefinitions AuthJWT
// @in header
// @name Authorization
// @description JWT token in authorization bearer

func main() {
	if err := server.
		NewServer().
		MapHandlers().
		Run(); err != nil {
		panic(fmt.Sprintf("can't start server %+v", err))
	}
}
