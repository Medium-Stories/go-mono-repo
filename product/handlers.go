package product

import (
	"github.com/gin-gonic/gin"
	pbProduct "github.com/medium-stories/go-mono-repo/product/proto"
	"github.com/sirupsen/logrus"
	"net/http"
)

func GetProductsByFilterHandler(client pbProduct.ProductClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		products, err := client.GetProductsByFilter(c.Request.Context(), &pbProduct.GetProductsByFilterRequest{})
		if err != nil {
			logrus.Error(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c.JSON(
			http.StatusOK,
			products,
		)
	}
}
