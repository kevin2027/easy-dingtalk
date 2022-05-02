package utils

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	"github.com/alibabacloud-go/tea/tea"
)

func GetOpenApiConfig() (config *openapi.Config) {
	config = &openapi.Config{}
	config.Protocol = tea.String("https")
	config.RegionId = tea.String("central")
	return
}
