package alarm

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

//Alarm 用来发送告警
type Alarm struct {
	serverAddress string
	projectid     string
	token         string
}

//SendWarning 发送普通警告信息
func (al *Alarm) SendWarning(msg string, etime string) (result string, err error) {
	nowtime := time.Now().Unix()
	md5ctx := md5.New()
	md5ctx.Write([]byte(al.token + strconv.FormatInt(nowtime, 10)))
	key := hex.EncodeToString(md5ctx.Sum(nil))
	requrl := al.serverAddress + "?p=%s&t=%d&k=%s&l=1&e=%d&m=%s"
	requrl = fmt.Sprintf(requrl, al.projectid, nowtime, key, etime, msg)
	_, err = http.Get(requrl)
	if err != nil {
		return "http Get Failed", err
	}
	return "ok", nil
}

//SendAlarm 发送严重告警
func (al *Alarm) SendAlarm(msg string, etime string) (result string, err error) {
	//msg预处理, 中文需要转码
	msg = url.QueryEscape(msg)
	nowtime := time.Now().Unix()
	md5ctx := md5.New()
	md5ctx.Write([]byte(al.token + strconv.FormatInt(nowtime, 10)))
	key := hex.EncodeToString(md5ctx.Sum(nil))
	requrl := al.serverAddress + "?p=%s&t=%d&k=%s&l=2&e=%d&m=%s"
	requrl = fmt.Sprintf(requrl, al.projectid, nowtime, key, etime, msg)
	_, err = http.Get(requrl)
	if err != nil {
		return "http Get Failed", err
	}
	return "ok", nil
}

//Config 配置告警功能
func (al *Alarm) Config(server string, pid string, token string) {
	al.serverAddress = server
	al.projectid = pid
	al.token = token
}

var _instance *Alarm

//GetAlarm 获取Alarm单例
func GetAlarm() *Alarm {
	if _instance == nil {
		_instance = new(Alarm)
	}
	return _instance
}
