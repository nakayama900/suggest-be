package web

import (
	"chatgpt-230308/src/suggest"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func NotImplemented(c *gin.Context) {
	//c.ShouldBind()
	c.AbortWithStatus(http.StatusNotImplemented)
	//c.JSON(http.StatusForbidden, gin.H{)
}
func GuestAccess(c *gin.Context) {

	var param suggest.ParamStruct
	if err := c.Bind(&param); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	chatEntity := suggest.GenOneWord(param.Queries, suggest.PromptV1)
	response, err := chatEntity.Request()
	log.Println("q: ", c.Query("q"))
	log.Println("response: ", response)
	log.Println("MessageContent: ", chatEntity.GetRaw())
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		panic(err)
		return
	} else if response == "" {
		panic("response has no result")
		return
	}
	c.String(http.StatusOK, "%v", response)
}
