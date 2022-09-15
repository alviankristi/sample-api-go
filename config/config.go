package config

type Config struct {
	Api struct {
		Hostname string `yaml:"hostname"`
		Port     string `yaml:"port"`
	} `yaml:"App"`

	App struct {
		Name string `yaml:"name"`
		Id   string `yaml:"id"`
	} `yaml:"Api"`

	DbConfig struct {
		Dialect          string `yaml:"dialect"`
		ConnectionString string `yaml:"connectionString"`
		MaxOpenConn      int    `yaml:"maxOpenConn"`
		MaxIdleConn      int    `yaml:"maxIdleConn"`
		MaxConnLifetime  int    `yaml:"maxConnLifetime"`
	} `yaml:"DbConfig"`
}
