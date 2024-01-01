package utils

import (
	"github.com/gin-gonic/gin"
	"wechat/global"
	"strconv"
)

// GetIntParamItem 将获取的参数进行转成int类型
func GetIntParamItem(param string, defaultInt int, c *gin.Context) (paramInt int) {
	if paramInt, _ = strconv.Atoi(c.Query(param)); paramInt < global.DEFAULT_NUM {
		paramInt = defaultInt
	}
	return
}
