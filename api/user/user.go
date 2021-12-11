package user

import (
	myjwt "firstGin/middleware/jwt"
	"firstGin/models/request"
	"firstGin/models/response"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func userLogin(c *gin.Context) {

	json := request.User{}
	c.Bind(&json)
	createToken(c, json)
}

func createToken(c *gin.Context, user request.User) {
	j := &myjwt.JWT{
		[]byte("GPKYYDS"),
	}
	claims := myjwt.CustomClaims{
		user.Name,
		user.Role,
		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "GPK",                           //签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"msg":    err.Error(),
		})
		return
	}
	log.Println(token)
	data := response.LoginResult{
		User:  user,
		Token: token,
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "登陆成功!",
		"data":   data,
	})
	return
}

func GetDataByTime(c *gin.Context) {
	claims := c.MustGet("claims").(*myjwt.CustomClaims)
	if claims != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "token有效",
			"data":   claims,
		})
	}
}
