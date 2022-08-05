package gateway

import (
	"github.com/gin-gonic/gin"
	pbUser "github.com/medium-stories/go-mono-repo/user/proto"
	"github.com/sirupsen/logrus"
	"net/http"
)

type CreateAccount struct {
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Nickname  string `json:"nickname"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Country   string `json:"country"`
}

func (api *api) CreateAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req *CreateAccount
		if err := c.ShouldBindJSON(&req); err != nil {
			logrus.Error(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		account, err := api.rpcClient.AddAccount(c.Request.Context(), &pbUser.AccountRequest{
			FirstName: req.Firstname,
			LastName:  req.Lastname,
			Nickname:  req.Nickname,
			Password:  req.Password,
			Email:     req.Email,
			Country:   req.Country,
		})
		if err != nil {
			logrus.Error(err)
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		c.JSON(http.StatusOK, account)
	}
}

func (api *api) DeleteAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")

		resp, err := api.rpcClient.DeleteAccount(c.Request.Context(), &pbUser.DeleteAccountRequest{
			Id: idParam,
		})
		if err != nil {
			logrus.Error(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusOK, resp)
	}
}
