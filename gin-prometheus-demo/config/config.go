package config

var conf *Config

// Config Apollo读取配置
type Config struct {
	Etcd string
	Log  string
	Db   string
}

func InitConfig() {

}

func GetConf() *Config {
	return conf
}

// Extension
//func (c *Config)()  {
//
//}
