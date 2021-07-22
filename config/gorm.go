package config

type Mysql struct {
	Path           string    `mapstructure:"path"     json:"path"    yaml:"path"`
	Config         string    `mapstructure:"config"   json:"config"  yaml:"config"`
	Dbname         string
	Username       string
	Password       string
	MaxIdleConns   int
	MaxOpenConns   int
	LogMode        bool
	LogZap         string
}

func (m *Mysql) Dsn() string {

	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
}