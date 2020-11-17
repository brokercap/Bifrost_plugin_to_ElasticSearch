package src

import (
	"github.com/brokercap/Bifrost/sdk/plugin/http_manager"
	"testing"
)

func TestHtmlView(t *testing.T)  {
	http_manager.SetToServer("elasticsearchTest","elasticsearch","uri")
	p := &http_manager.Param {
		//HtmlDir: "/mnt/hgfs/project/gopath/src/github.com/brokercap/Bifrost_plugin_to_ElasticSearch/",
		HtmlDir: "E:/project/gopath/src/github.com/brokercap/Bifrost_plugin_to_ElasticSearch/www/",
	}
	http_manager.Start(p)
}