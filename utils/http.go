package utils

import (
	"time"

	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"golang.org/x/xerrors"
)

func DoRquest(method string, path string, query map[string]*string, body interface{}) (response *tea.Response, err error) {
	req := tea.NewRequest()
	req.Protocol = tea.String("https")
	req.Pathname = tea.String(path)
	req.Domain = tea.String("oapi.dingtalk.com")
	req.Method = tea.String(method)
	req.Headers = map[string]*string{
		"accept":       tea.String("application/json"),
		"content-type": tea.String("application/json; charset=utf-8"),
		"host":         tea.String("oapi.dingtalk.com"),
	}
	req.Query = query
	req.Body = tea.ToReader(util.ToJSONString(body))
	runtimeOptions := &tea.RuntimeObject{
		ConnectTimeout: tea.Int(int(time.Second) * 2),
	}
	response, err = tea.DoRequest(req, tea.ToMap(runtimeOptions))
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}
