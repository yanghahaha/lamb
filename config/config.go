package configs

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

//AmsConfig : config of ams project
type AmsConfig struct {
	jsonConf       Config
	env            string
	path           string
	DbConf         DbConfig
	LogConf        LogConfig
	QiniuConf      QiniuConfig
	VodMonitorConf VodMonitorConfig
	LsaConf        LsaConfig
}

//DbConfig : struct of mysql dbconfg
type DbConfig struct {
	Host     string
	Port     string
	Uname    string
	Passwd   string
	Database string
	Charset  string
}

//LogConfig : config of logger
type LogConfig struct {
	Level     int
	OutputStr string
}

//QiniuConfig :config of qiniucdn
type QiniuConfig struct {
	AccessKey string
	SecretKey string
	Hub       string
}

//VodMonitorConfig : config of logger
type VodMonitorConfig struct {
	StartTime int
}

//LsaConfig : config of logger
type LsaConfig struct {
	ServerAddress string
	Pid           string
	Token         string
}

//LoadConf : load/reload config from configPath
func (cnf *AmsConfig) LoadConf() bool {
	cnf.jsonConf = NewConfig(cnf.path)
	ret, _ := cnf.jsonConf.Get(cnf.env)
	if ret == false {
		println("environment not found!")
		return false
	}
	//加载数据库配置
	if cnf.loadDbConf() == false {
		return false
	}
	//加载日志配置
	if cnf.loadLoggerConf() == false {
		return false
	}
	//加载告警配置
	if cnf.loadLsaConf() == false {
		return false
	}
	//加载七牛配置
	if cnf.loadQiniuConf() == false {
		return false
	}

	//加载点播监测配置
	if cnf.loadVodMonitorConf() == false {
		return false
	}

	//加载分析报告配置
	if cnf.loadReportConf() == false {
		return false
	}
	return true
}

func (cnf *AmsConfig) loadDbConf() bool {
	//mySqlconf.Uname + ":" + mySqlconf.Passwd + "@tcp(" + mySqlconf.Host + ":" + mySqlconf.Port + ")/" + mySqlconf.Database + "?charset=" + mySqlconf.Charset
	ret, uname := cnf.jsonConf.Get(cnf.env + ".db.uname")
	if ret == false {
		println("db.uname unset")
		return false
	}
	cnf.DbConf.Uname = uname.(string)

	ret, passwd := cnf.jsonConf.Get(cnf.env + ".db.passwd")
	if ret == false {
		println("db.password unset")
		return false
	}
	cnf.DbConf.Passwd = passwd.(string)

	ret, host := cnf.jsonConf.Get(cnf.env + ".db.host")
	if ret == false {
		println("db.host unset")
		return false
	}
	cnf.DbConf.Host = host.(string)

	ret, port := cnf.jsonConf.Get(cnf.env + ".db.port")
	if ret == false {
		println("db.port unset")
		return false
	}
	cnf.DbConf.Port = port.(string)

	ret, database := cnf.jsonConf.Get(cnf.env + ".db.database")
	if ret == false {
		println("db.database unset")
		return false
	}
	cnf.DbConf.Database = database.(string)

	ret, charset := cnf.jsonConf.Get(cnf.env + ".db.charset")
	if ret == false {
		println("db.password unset")
		return false
	}
	cnf.DbConf.Charset = charset.(string)

	return true
}

func (cnf *AmsConfig) loadLoggerConf() bool {
	ret, level := cnf.jsonConf.Get(cnf.env + ".log.level")
	if ret == false {
		println("log.level unset")
		return false
	}
	switch v := level.(type) {
	case string:
		var s string
		s = v
		cnf.LogConf.Level, _ = strconv.Atoi(s)
	case int:
		var l int
		l = v
		cnf.LogConf.Level = l
	}

	cnf.LogConf.Level, _ = strconv.Atoi(level.(string))

	ret, outputStr := cnf.jsonConf.Get(cnf.env + ".log.output")
	if ret == false {
		println("log.output unset")
		return false
	}
	cnf.LogConf.OutputStr = outputStr.(string)

	return true
}

func (cnf *AmsConfig) loadQiniuConf() bool {
	ret, ak := cnf.jsonConf.Get(cnf.env + ".qiniu.accessKey")
	if ret == false {
		println("qiniu.accessKey unset")
		return false
	}
	switch v := ak.(type) {
	case string:
		var s string
		s = v
		cnf.QiniuConf.AccessKey = s
	default:
		println("qiniu.accessKey should by string")
		return false
	}

	ret, sk := cnf.jsonConf.Get(cnf.env + ".qiniu.secretKey")
	if ret == false {
		println("qiniu.secretKey unset")
		return false
	}
	switch v := sk.(type) {
	case string:
		var s string
		s = v
		cnf.QiniuConf.SecretKey = s
	default:
		println("qiniu.secretKey should by string")
		return false
	}

	ret, hub := cnf.jsonConf.Get(cnf.env + ".qiniu.hub")
	if ret == false {
		println("qiniu.hub unset")
		return false
	}
	switch v := hub.(type) {
	case string:
		var s string
		s = v
		cnf.QiniuConf.Hub = s
	default:
		println("qiniu.hub should by string")
		return false
	}

	return true
}

func (cnf *AmsConfig) loadVodMonitorConf() bool {
	ret, startTime := cnf.jsonConf.Get(cnf.env + ".vodMonitor.startTime")
	if ret == false {
		println("vodMonitor.startTime unset")
		return false
	}
	switch v := startTime.(type) {
	case int:
		var s int
		s = v
		cnf.VodMonitorConf.StartTime = s
	case float64:
		var s int
		s = int(v)
		cnf.VodMonitorConf.StartTime = s
	default:
		fmt.Println("vodMonitor.startTime should by numbers but ", reflect.TypeOf(startTime))
		cnf.VodMonitorConf.StartTime = 0
		return false
	}
	return true
}

func (cnf *AmsConfig) loadLsaConf() bool {
	ret, serverAddress := cnf.jsonConf.Get(cnf.env + ".alarm.serverAddress")
	if ret == false {
		println("alarm.serverAddress unset")
		return false
	}
	switch v := serverAddress.(type) {
	case string:
		var s string
		s = v
		cnf.LsaConf.ServerAddress = s
	default:
		fmt.Println("alarm.serverAddress should by string but ", reflect.TypeOf(serverAddress))
		return false
	}

	ret, pid := cnf.jsonConf.Get(cnf.env + ".alarm.pid")
	if ret == false {
		println("alarm.pid unset")
		return false
	}
	switch v := pid.(type) {
	case string:
		var s string
		s = v
		cnf.LsaConf.Pid = s
	default:
		fmt.Println("alarm.pid should by string but ", reflect.TypeOf(pid))
		return false
	}

	ret, token := cnf.jsonConf.Get(cnf.env + ".alarm.token")
	if ret == false {
		println("alarm.token unset")
		return false
	}
	switch v := token.(type) {
	case string:
		var s string
		s = v
		cnf.LsaConf.Token = s
	default:
		fmt.Println("alarm.token should by string but ", reflect.TypeOf(pid))
		return false
	}
	return true
}

func (cnf *AmsConfig) loadReportConf() bool {
	return true
}

//Init :init config with path and environment
func (cnf *AmsConfig) Init(path string, env string) bool {
	if checkConfigPath(path) == false {
		return false
	}
	if env == "" {
		env = "pro"
	}
	cnf.env = env
	cnf.path = path
	if cnf.LoadConf() == false {
		return false
	}
	return true
}

//单例
var _instance *AmsConfig

//GetConfig 获取单例对象
func GetConfig() *AmsConfig {
	if _instance == nil {
		_instance = new(AmsConfig)
	}
	return _instance
}

func checkConfigPath(_path string) bool {
	fileInfo, err := os.Stat(_path)
	if err != nil && os.IsNotExist(err) {
		println("config file not exist")
		return false
	}
	if fileInfo.IsDir() {
		println("config path can't be directory")
		return false
	}
	return true
}
