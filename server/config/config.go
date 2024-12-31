package config

type Config struct {
	Env string `yaml:"env"`
	Dryrun bool `yaml:"dryrun"`
	AllowOrigins []string `yaml:"allow_origins"`
	Db ConfigDb `yaml:"db"`
}

type ConfigDb struct {
	Ms string `yaml:"ms"`
	Name string `yaml:"name"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
	Addr string `yaml:"addr"`
	Net string `yaml:"net"`
	Collation string `yaml:"collation"`
}
