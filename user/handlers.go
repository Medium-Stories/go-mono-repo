package user

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

func CreateAccountHandler(client pbUser.AccountManagementClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req *CreateAccount
		if err := c.ShouldBindJSON(&req); err != nil {
			logrus.Error(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		account, err := client.AddAccount(c.Request.Context(), &pbUser.AccountRequest{})
		if err != nil {
			logrus.Error(err)
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		c.JSON(http.StatusOK, account)
	}
}
