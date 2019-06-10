package configs

import (
	"gopkg.in/gomail.v2"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"mirrors_status/internal/log"
)

type InfluxDBConf struct {
	Host     string `yaml:"host"`
	Port     int `yaml:"port"`
	DBName   string `yaml:"dbName"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type HttpConf struct {
	Port int `yaml:"port"`
	Host string `yaml:"host"`
	AllowOrigin string `yaml:"allow-origin"`
	AdminMail string `yaml:"admin-mail"`
}

type MySQLConf struct {
	Host     string `yaml:"host"`
	Port     int `yaml:"port"`
	DBName   string `yaml:"dbName"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type CdnCheckerConf struct {
	DefaultCdn string `yaml:"default-cdn"`
	UserAgent string `yaml:"user-agent"`
	ApiSite string `yaml:"api-site"`
	ApiPath string `yaml:"api-path"`
	Target string `yaml:"target"`
	SourceUrl string `yaml:"source-url"`
	SourcePath string `yaml:"source-path"`
}

type LdapConf struct {
	Server string `yml:"server"`
	Port int `yml:"port"`
	BindDn string `yml:"bind_dn"`
	BindPasswd string `yml:"bind_passwd"`
	UserSearch string `yml:"user_search"`
	GroupSearch string `yml:"group_search"`
}

type MailConf struct {
	Host string `yml:"host"`
	Port int `yml:"port"`
	Username string `yml:"username"`
	Password string `yml:"password"`
}

type JenkinsConf struct {
	Addr string `yml:"addr"`
	Trigger string `yml:"trigger"`
	Delay int `yml:"delay"`
	Retry int `yml:"retry"`
}

type RedisConf struct {
	Host string `yml:"host"`
	Port int `yml:"port"`
	Username string `yml:"username"`
	Password string `yml:"password"`
	DBName int `yml:"client"`
}

type ServerConf struct {
	InfluxDB *InfluxDBConf `yaml:"influxdb"`
	MySQLDB *MySQLConf `yaml:"mysql"`
	Http     *HttpConf     `yaml:"http"`
	CdnChecker *CdnCheckerConf `yaml:"cdn-checker"`
	Ldap *LdapConf `yml:"ldap"`
	Mail *MailConf `yml:"mail"`
	Jenkins *JenkinsConf `yml:"jenkins"`
	Redis *RedisConf `yml:"redis"`
}

func ErrHandler(op string, err error) {
	if err != nil {
		log.Fatalf("%s found error: %v", op, err)
	}
}

func NewServerConfig() *ServerConf {
	var serverConf ServerConf
	ymlFile, err := ioutil.ReadFile("configs/config.yml")
	ErrHandler("opening file", err)

	err = yaml.Unmarshal(ymlFile, &serverConf)
	ErrHandler("unmarshal yaml", err)

	return &serverConf
}

var (
	MailDialer *gomail.Dialer
)

func GetMailDialer() *gomail.Dialer {
	return MailDialer
}

func InitMailClient(conf *MailConf) {
	log.Infof("Trying connecting mail server:%s:%d %s", conf.Host, conf.Port, conf.Username)
	MailDialer = &gomail.Dialer{
		Host: conf.Host,
		Port: conf.Port,
		Username: conf.Username,
		Password: conf.Password,
		SSL: true,
	}
}