package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"market/global"
	"regexp"
	"strconv"
)

// GetIntParamItem 将获取的参数进行转成int类型
func GetIntParamItem(param string, defaultInt int, c *gin.Context) (paramInt int) {
	if paramInt, _ = strconv.Atoi(c.Query(param)); paramInt < global.DEFAULT_NUM {
		paramInt = defaultInt
	}
	return
}

// TruncateString 截取指定的字符串
func TruncateString(s string, num int) string {
	runes := []rune(s)
	if len(runes) <= num {
		return s
	}
	return string(runes[:num])
}

//ClearMobileText 清除手机号
func ClearMobileText(text string) (cleanedText string) {
	// 定义手机号的正则表达式
	phoneRegex := regexp.MustCompile(`1[3456789]\d{9}`)

	// 查找所有匹配的手机号
	matches := phoneRegex.FindAllString(text, -1)

	if matches != nil {
		fmt.Println("找到的手机号：", matches)
		// 将手机号去除
		cleanedText = phoneRegex.ReplaceAllString(text, "[手机号敏感数据不予显示]")
		fmt.Println("去除手机号后的文本：", cleanedText)
	} else {
		cleanedText = text
		fmt.Println("未找到手机号")
	}

	return
}
