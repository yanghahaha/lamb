package configs

import (
	"testing"
)

func Test_ValiableConfig(t *testing.T) {
	ret := GetConfig().Init("/Users/woo/develop/goWorkerSpace/ams/config/config.json", "")
	if ret == false {
		t.Error("init config failed")
	}
	GetConfig().LoadConf()
	GetConfig().Init("/Users/woo/develop/goWorkerSpace/ams/config/config.json", "")
	if GetConfig().Init("/Users/woo/develop/goWorkerSpace/ams/config/config.json", "") == false {
		t.Error("123")
	}
}
func Test_UnvaildConfig_dir(t *testing.T) {
	ret := GetConfig().Init("/Users/woo/", "")
	if ret == true {
		t.Error("init config failed")
	}
}

func Test_UnvaildConfig_notexsit(t *testing.T) {
	ret := GetConfig().Init("/Users/woo/develop/goWorkerSpace/ams/config/configaaa.json", "")
	if ret == true {
		t.Error("init config failed")
	}
}
