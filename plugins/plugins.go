package plugins

import "html/template"

var Plugins = map[string]func() template.HTML{
	"sidebar": Sidebar,
}

func init() {

}
func Del() {

}
