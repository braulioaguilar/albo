package config

import (
	"log"

	"github.com/spf13/viper"
)

type config struct {
	MONGO_URI string
	DB_NAME   string
	DURATION  int
}

var Config config

func init() {
	viper.AutomaticEnv()
	viper.BindEnv("MONGO_URI")
	viper.BindEnv("DURATION")
	viper.BindEnv("DB_NAME")

	viper.SetDefault("MONGO_URI", "mongodb://mongo_albo:27017/admin?authSource=admin")
	viper.SetDefault("DURATION", 60)
	viper.SetDefault("DB_NAME", "albo")

	if err := viper.Unmarshal(&Config); err != nil {
		log.Panicf("Error unmarshalling configuration: %s", err)
	}

	log.Println("Parameters loaded are:")
	for _, k := range viper.AllKeys() {
		log.Printf("\t%s=%v\n", k, viper.Get(k))
	}
}
