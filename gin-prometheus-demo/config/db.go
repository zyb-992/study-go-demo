package config

type Db interface {
}

type Mysql struct {
}

func (c *Config) DbConf() Db {

}
