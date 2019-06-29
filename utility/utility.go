package utility

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParamToInt(c *gin.Context, key string) (int, error) {
	param := c.Param(key)
	value, err := strconv.Atoi(param)
	if err != nil {
		return 0, err
	}
	return value, nil
}
