package convert

import (
	"github.com/gin-gonic/gin"
)

const CtxUserIDKey = "userID"

func GetUserID(c *gin.Context) (int64, bool) {
	iUserID, ok := c.Get(CtxUserIDKey)
	if !ok {
		return 0, false
	}
	userID, _ := iUserID.(int64)
	return userID, true
}

func GetUserIDMust(c *gin.Context) int64 {
	userID, _ := GetUserID(c)
	return userID
}
