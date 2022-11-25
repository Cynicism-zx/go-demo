package i18n

import (
	"io/ioutil"
	"testing"

	"go-demo/utils/i18n/locales"
)

func TestLoadYaml(t *testing.T) {
	LoadYaml()
}

func TestLoadJson(t *testing.T) {
	LoadJson()
}

func TestLoadJsonFile(t *testing.T) {
	filePath := locales.Path("zh-CN.json")
	f, _ := ioutil.ReadFile(filePath)
	LoadJsonFile(f)
}
