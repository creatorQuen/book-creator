package config

type Config struct {
	ConfigDataBase struct {
		Host         string `envconfig:"HOST_DB"`
		Port         int    `envconfig:"PORT_DB"`
		User         string `envconfig:"USER_DB"`
		Password     string `envconfig:"PASSWORD_DB"`
		NameDataBase string `envconfig:"NAME_DB"`
		ListenPort   string `envconfig:"LISTEN_PORT"`
	}
}
