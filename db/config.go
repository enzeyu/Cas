package db

import "github.com/spf13/viper"

func LoadConfig(fpath string) (c *Config, err error) {
	if fpath == "" {
		fpath = "E:\\GOPATH\\src\\Cas\\db\\config.yaml"
	}
	v := viper.New()
	v.SetConfigFile(fpath)
	v.SetConfigType("yaml")
	if err1 := v.ReadInConfig(); err1 != nil {
		err = err1
		return
	}
	c = &Config{}
	c.Redis.Path=v.GetString("redis.path")
	c.Redis.Host=v.GetString("redis.host")
	return
}

type Config struct {
	Redis Rediscon
}

type Rediscon struct {
	Path string `json:"path"`
	Host string `json:"host"`
}