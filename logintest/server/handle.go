package server

import (
	"desktop/logintest/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserLogin(c *gin.Context) {
	reqData := sql.Register{}
	err := c.ShouldBind(&reqData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": "参数错误或用户名和密码不规范",
		})
		return
	}

	_, err = sql.CheckLogin(reqData)
	if err != nil {
		fmt.Println("3,err:",err.Error())
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
		return
	}

	t1, t2, err := sql.NewToken(reqData.UserName)
	if err != nil {
		fmt.Println("4,err:",err.Error())
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token1": t1,
		"token2": t2,
	})
}

func RefreshToken(c *gin.Context) {
	token := sql.ReqToken{}
	err := c.ShouldBind(&token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": "err",
		})
		return
	}

	b,err := sql.CheckRefreshTimeByToken6(token.Token)
	if err != nil || !b{
		c.JSON(http.StatusOK,gin.H{
			"err":"err",
		})
		return
	}

	err = sql.RefreshToken2(token.Token)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"err":"err",
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"ok":"ok",
	})
}

func UserRegister(c *gin.Context) {
	reqData := sql.Register{}
	err := c.ShouldBind(&reqData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": "参数错误或用户名和密码不规范",
		})
		return
	}

	res, err := sql.CheckRegister(reqData.UserName)
	if err != nil {
		fmt.Println("1,err:",err)
		c.JSON(http.StatusOK, gin.H{
			"err": "err",
		})
		return
	} else if !res {
		c.JSON(http.StatusOK, gin.H{
			"err": "用户名已被注册",
		})
		return
	}

	err = sql.InsertUser(reqData)
	if err != nil {
		fmt.Println("2,err:",err)
		c.JSON(http.StatusOK, gin.H{
			"err": "err",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok": "ok",
	})
}
