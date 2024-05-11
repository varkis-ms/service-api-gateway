package config

import "github.com/spf13/viper"

type Config struct {
	Env                   string `mapstructure:"ENV"`
	PortHttp              string `mapstructure:"HTTP_PORT"`
	AuthClientAddr        string `mapstructure:"AUTH_CLIENT_ADDR"`
	UserInfoClientAddr    string `mapstructure:"USER_INFO_CLIENT_ADDR"`
	CompetitionClientAddr string `mapstructure:"COMPETITION_CLIENT_ADDR"`
	SolutionClientAddr    string `mapstructure:"SOLUTION_CLIENT_ADDR"`
}

// LoadConfig Конструктор для создания Config, который содержит считанные из .env файла данные.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
