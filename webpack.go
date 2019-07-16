package webpack

import (
	"fmt"

	"github.com/micro/go-micro/config"
)

type entrypoints struct {
	App string `json:"app"`
}

func encoreTags(name, epyt string) (string, error) {
	config.LoadFile("./public/build/entrypoints.json")
	confMap := config.Map()

	for _, value := range confMap {
		for key, val := range value.(map[string]interface{}) {
			if key == name {
				for k, v := range val.(map[string]interface{}) {
					if k != epyt {
						continue
					}

					switch k {
					case "css":
						cssHTML := ""
						for _, file := range v.([]interface{}) {
							cssFile := file.(string)
							cssHTML += fmt.Sprintf("<link href='./public%s'>\r\n", cssFile)
						}
						return cssHTML, nil
					case "js":
						jsHTML := ""
						for _, file := range v.([]interface{}) {
							jsFile := file.(string)
							jsHTML += fmt.Sprintf("<script src='./public%s'></script>\r\n", jsFile)
						}
						return jsHTML, nil
					}
				}
			}
		}
	}

	return "", nil
}

// EncoreLinkTags EncoreLinkTags
func EncoreLinkTags(name string) (string, error) {
	html, err := encoreTags(name, "css")
	if err != nil {
		return "", err
	}
	return html, nil
}

// EncoreScriptTags EncoreScriptTags
func EncoreScriptTags(name string) (string, error) {
	html, err := encoreTags(name, "js")
	if err != nil {
		return "", err
	}
	return html, nil
}
