package pluginproxy

import (
	"bytes"
	"fmt"
	"net/url"
	"text/template"

	m "github.com/grafana/grafana/pkg/models"
	"github.com/grafana/grafana/pkg/plugins"
)

// InterpolateString accepts template data and return a string with substitutions
func InterpolateString(text string, data templateData) (string, error) {
	t, err := template.New("content").Parse(text)
	if err != nil {
		return "", fmt.Errorf("无法解析模板 %s", text)
	}

	var contentBuf bytes.Buffer
	err = t.Execute(&contentBuf, data)
	if err != nil {
		return "", fmt.Errorf("无法执行模板 %s", text)
	}

	return contentBuf.String(), nil
}

// InterpolateURL accepts template data and return a string with substitutions
func InterpolateURL(anURL *url.URL, route *plugins.AppPluginRoute, orgID int64, appID string) (*url.URL, error) {
	query := m.GetPluginSettingByIdQuery{OrgId: orgID, PluginId: appID}
	result, err := url.Parse(anURL.String())
	if query.Result != nil {
		if len(query.Result.JsonData) > 0 {
			data := templateData{
				JsonData: query.Result.JsonData,
			}
			interpolatedResult, err := InterpolateString(anURL.String(), data)
			if err == nil {
				result, err = url.Parse(interpolatedResult)
				if err != nil {
					return nil, fmt.Errorf("解析插件路由地址时出错 %v", err)
				}
			}
		}
	}

	return result, err
}
