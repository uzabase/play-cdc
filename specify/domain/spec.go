package domain

import (
	"fmt"
	"strings"
)

func GenerateSpec(contract Contract) string {
	request := contract.Request
	response := contract.Response

	var spec strings.Builder

	spec.WriteString(fmt.Sprintf("## %s %s\n", request.Method, request.UrlPath))
	spec.WriteString(fmt.Sprintf("* URL\"%s\"に%sリクエストを送る\n", request.UrlPath, request.Method))
	spec.WriteString(fmt.Sprintf("* レスポンスステータスコードが\"%d\"である\n", response.Status))

	return spec.String()
}
