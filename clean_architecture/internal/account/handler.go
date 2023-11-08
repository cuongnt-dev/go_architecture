package account

import "github.com/gin-gonic/gin"

type Handler struct {
	UseCase UseCaseI
}

func NewHandler(uc *UserCase) *Handler {
	return &Handler{
		UseCase: uc,
	}
}

func (h *Handler) CreateTransaction(ctx *gin.Context) {

}

func (h *Handler) GetTransaction(ctx *gin.Context) {

}
