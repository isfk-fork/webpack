package webpack

import (
	"fmt"

	"github.com/micro/go-micro/config"
)

func encoreTags(name, epyt string) (string, error) {
	// Load json config file
	config.LoadFile("./template/build/entrypoints.json")
	conf := config.Map()
	entrypoints := conf["entrypoints"]
	fmt.Println(entrypoints)
	return "", nil
}

// EncoreLinkTags EncoreLinkTags
func EncoreLinkTags(name string) {
	encoreTags(name, "css")
}

// EncoreScriptTags EncoreScriptTags
func EncoreScriptTags(name string) {
	encoreTags(name, "js")
}
