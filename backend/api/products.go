package api

import (
	"net/http"
	db "tsn-ecommerce/db/sqlc"

	"github.com/gin-gonic/gin"
)

type ListProductsRequest struct {
	Language string `form:"language" binding:"required"`
	Limit    int32  `form:"limit" binding:"required"`
	Offset   int32  `form:"offset" binding:"required"`
}

func (server *Server) ListProducts(ctx *gin.Context) {
	var req ListProductsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create args for listing products
	args := db.ListProductWithInfoParams{
		Language: db.LanguageCode(req.Language),
		Limit:    req.Limit,
		Offset:   req.Offset,
	}

	// list products from database
	products, err := server.store.ListProductWithInfo(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}
