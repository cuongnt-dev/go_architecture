package main

import (
	"clean_architecture/internal/account"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	accountRepository := account.NewRepository()
	accountUsecase := account.NewUserCase(accountRepository)
	accountHandler := account.NewHandler(accountUsecase)

	r.POST("/account/transaction/list", accountHandler.GetTransaction)
	r.POST("/account/transaction/add", accountHandler.CreateTransaction)
	r.Run()
}
