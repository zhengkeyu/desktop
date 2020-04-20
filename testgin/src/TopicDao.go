package src

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTopicDetail(c *gin.Context) {
	c.String(http.StatusOK, "%v", c.Param("flag"))
}

func GetTopicList(c *gin.Context) {
	qs := TopicQuery{} //获取到query信息
	err := c.BindQuery(&qs) //绑定query信息到结构体上
	if err != nil{
		//panic(err)
		c.String(http.StatusBadRequest,err.Error())
		return
	}
	fmt.Println("query = ",qs)
	q := c.Query("username") //获取查询(query)
	if q == "" {
		c.String(http.StatusOK, "获取所有的列表")
	} else {
		c.String(http.StatusOK, "获取%v的列表", q)
	}
}

//判断登录
func MustLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, ok := c.GetQuery("token"); !ok {
			c.String(http.StatusUnauthorized, "缺少token参数")
			c.Abort() //不会再往下走
		}else{
			//验证token
		}
	}
}
func NewTopic(c *gin.Context) {
	c.String(http.StatusOK, "新增帖子")
}

func DelTopic(c *gin.Context) {
	c.String(http.StatusOK, "删除帖子")
}
