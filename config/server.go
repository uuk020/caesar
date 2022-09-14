package config

type Server struct {
	Name        string `mapstructure:"name"`
	Description string `mapstructure:"description"`
	Port        int    `mapstructure:"port"`
	Debug       bool   `mapstructure:"debug"`
	Timezone    string `mapstructure:"timezone"`
	SecretKey   string `mapstructure:"secret_key"`
	Mysql       Mysql  `mapstructure:"mysql"`
	Redis       Redis  `mapstructure:"redis"`
	LogAddr     string `mapstructure:"logAddr"`
	Email       Email  `mapstructure:"email"`
}

type Mysql struct {
	Name         string `mapstructure:"name"`
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"db_name"`
	MaxOpenConns int    `mapstucture:"max_open_conns"`
	MaxIdleConns int    `mapstucture:"max_idle_conns"`
}

type Redis struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
}

type Email struct {
	Name     string `mapstructure:"name"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
}
