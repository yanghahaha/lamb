package logger

import (
	"os"
	"time"

	"strings"

	"errors"

	"fmt"

	logger "github.com/apsdehal/go-logger"
)

const (
	LEVEL_ERROR   = 1
	LEVEL_WARNING = 2
	LEVEL_INFO    = 3
	LEVEL_DEBUG   = 4
)

//ELogger 日志类
type ELogger struct {
	inConfig   bool
	level      int
	logger     logger.Logger
	loggerList []logger.Logger
}

func (log *ELogger) debug(msg string) bool {
	if log.inConfig {
		time.Sleep(10000000)
		log.debug(msg)
	}
	if log.level >= LEVEL_DEBUG {
		len := len(log.loggerList)
		for i := 0; i < len; i++ {
			log.loggerList[i].Debug(msg)
		}
		return true
	}
	return false
}
func (log *ELogger) info(msg string) bool {
	if log.inConfig {
		time.Sleep(10000000)
		log.info(msg)
	}
	if log.level >= LEVEL_INFO {
		len := len(log.loggerList)
		for i := 0; i < len; i++ {
			log.loggerList[i].Info(msg)
		}
		return true
	}
	return false
}
func (log *ELogger) error(msg string) bool {
	if log.inConfig {
		time.Sleep(10000000)
		log.error(msg)
	}
	if log.level >= LEVEL_ERROR {
		len := len(log.loggerList)
		for i := 0; i < len; i++ {
			log.loggerList[i].Error(msg)
		}
		return true
	}
	return false
}
func (log *ELogger) warning(msg string) bool {
	if log.inConfig {
		time.Sleep(10000000)
		log.warning(msg)
	}
	if log.level >= LEVEL_WARNING {
		len := len(log.loggerList)
		for i := 0; i < len; i++ {
			log.loggerList[i].Warning(msg)
		}
		return true
	}
	return false
}

func (log *ELogger) initOutput(outConfigStr string) (err error) {
	//解析outconfigstr
	configList := strings.Split(outConfigStr, "|")
	configLen := len(configList)
	if configLen == 0 {
		return errors.New("没有日志配置")
	}
	for i := 0; i < configLen; i++ {
		outStr := configList[i]
		//判断输出类型
		if outStr == "stdOut" { //控制台输出
			logwriter, err := logger.New("AMS", 1, os.Stdout)
			if err != nil {
				println("日志初始化出错: " + err.Error())
			}
			log.loggerList = append(log.loggerList, *logwriter)
		} else if strings.Contains(outStr, "file://") { //日志文件输出
			logfilePath := StrFormatTime(strings.Replace(outStr, "file://", "", 1))
			logfile, err := os.OpenFile(logfilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
			if err != nil {
				println("日志初始化出错: " + err.Error())
			}
			logwriter, err := logger.New("AMS", 1, logfile)
			if err != nil {
				println("日志初始化出错: " + err.Error())
			}
			log.loggerList = append(log.loggerList, *logwriter)
		} else { //不支持的类型
			println("不支持的日志配置:" + outStr)
		}
	}
	return nil
}

//Config 配置logger
func (log *ELogger) Config(level int, outstr string) bool {
	if level != LEVEL_DEBUG && level != LEVEL_ERROR && level != LEVEL_WARNING && level != LEVEL_INFO {
		level = LEVEL_ERROR
	}
	log.level = level
	log.loggerList = []logger.Logger{}
	log.inConfig = true
	err := log.initOutput(outstr)
	log.inConfig = false
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

var _instance *ELogger

//GetLogger 获取logger单例
func GetLogger() *ELogger {
	if _instance == nil {
		_instance = new(ELogger)
	}
	return _instance
}

//Debug 写Debug日志
func Debug(msg string) {
	try(
		func() { GetLogger().debug(msg) },
		func(e interface{}) { fmt.Println(e) })
}

//Warning 写Warning日志
func Warning(msg string) {
	try(
		func() { GetLogger().warning(msg) },
		func(e interface{}) { fmt.Println(e) })
}

//Info 写Info日志
func Info(msg string) {
	try(
		func() { GetLogger().info(msg) },
		func(e interface{}) { fmt.Println(e) })
}

//Error 写Error日志
func Error(msg string) {
	try(
		func() { GetLogger().error(msg) },
		func(e interface{}) { fmt.Println(e) })
}

func try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}

//StrFormatTime 字符串格式化为当前时间
//%Y %m %d %H %i %s : year month day hour minutes second
func StrFormatTime(instr string) (outstr string) {
	//2006-01-02 15:04:05
	nowtime := time.Now()
	year := nowtime.Format("2006")
	month := nowtime.Format("01")
	day := nowtime.Format("02")
	hour := nowtime.Format("15")
	minutes := nowtime.Format("04")
	sencond := nowtime.Format("05")
	outstr = strings.Replace(instr, "%Y", year, -1)
	outstr = strings.Replace(outstr, "%m", month, -1)
	outstr = strings.Replace(outstr, "%d", day, -1)
	outstr = strings.Replace(outstr, "%H", hour, -1)
	outstr = strings.Replace(outstr, "%i", minutes, -1)
	outstr = strings.Replace(outstr, "%s", sencond, -1)
	return
}
