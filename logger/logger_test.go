package logger

import (
	"testing"
)

func Test_debug(t *testing.T) {
	ret := GetLogger().Config(LEVEL_DEBUG, "file:///Users/woo/develop/goWorkerSpace/ams/logs/ams.log|stdOut")
	if ret == false {
		t.Error("init config failed")
	}
	if !GetLogger().debug("debug") {
		t.Error("Debug not pass")
	}
	if !GetLogger().warning("warning") {
		t.Error("Warning not pass")
	}
	if !GetLogger().error("error") {
		t.Error("Error not pass")
	}
}
func Test_warning(t *testing.T) {
	ret := GetLogger().Config(LEVEL_WARNING, "file:///Users/woo/develop/goWorkerSpace/ams/logs/ams.log|stdOut")
	if ret == false {
		t.Error("init config failed")
	}
	if GetLogger().debug("debug") {
		t.Error("Debug not pass")
	}
	if !GetLogger().warning("warning") {
		t.Error("Warning not pass")
	}
	if !GetLogger().error("error") {
		t.Error("Error not pass")
	}
}
func Test_error(t *testing.T) {
	ret := GetLogger().Config(LEVEL_ERROR, "file:///Users/woo/develop/goWorkerSpace/ams/logs/ams.log|stdOut")
	if ret == false {
		t.Error("init config failed")
	}
	if GetLogger().debug("debug") {
		t.Error("Debug not pass")
	}
	if GetLogger().warning("warning") {
		t.Error("Warning not pass")
	}
	if !GetLogger().error("error") {
		t.Error("Error not pass")
	}
}

func Test_update(t *testing.T) {
	ret := GetLogger().Config(LEVEL_DEBUG, "file:///Users/woo/develop/goWorkerSpace/ams/logs/ams.log|stdOut")
	if ret == false {
		t.Error("init config failed")
	}
	if !GetLogger().debug("debug") {
		t.Error("Debug not pass")
	}
	if !GetLogger().warning("warning") {
		t.Error("Warning not pass")
	}
	if !GetLogger().error("error") {
		t.Error("Error not pass")
	}
	GetLogger().Config(LEVEL_ERROR, "file:///Users/woo/develop/goWorkerSpace/ams/logs/ams.log")
	if GetLogger().debug("debug") {
		t.Error("Debug not pass")
	}
	if GetLogger().warning("warning") {
		t.Error("Warning not pass")
	}
	if !GetLogger().error("error") {
		t.Error("Error not pass")
	}
}

func Test_timeFormat(t *testing.T) {
	println(StrFormatTime("%Y%m%d%H%i%s"))
}
