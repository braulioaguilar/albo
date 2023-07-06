package config

import (
	"log"

	"github.com/spf13/viper"
)

type config struct {
	MARVEL_API string
	MONGO_URL  string
	API_KEY    string
	HASH       string
	TS         string
	DURATION   int
}

var Config config

func init() {
	viper.AutomaticEnv()
	viper.BindEnv("MARVEL_API")
	viper.BindEnv("MONGO_URL")
	viper.BindEnv("API_KEY")
	viper.BindEnv("HASH")
	viper.BindEnv("TS")
	viper.BindEnv("DURATION")

	viper.SetDefault("MARVEL_API", "http://gateway.marvel.com")
	viper.SetDefault("MONGO_URL", "mongodb://localhost:27017")
	viper.SetDefault("API_KEY", "9b2f073ee571a47da65a782a644cfffa")
	viper.SetDefault("HASH", "53c019e897a0467afba1614fdd342df9")
	viper.SetDefault("TS", "1688497908")
	viper.SetDefault("DURATION", 60)

	if err := viper.Unmarshal(&Config); err != nil {
		log.Panicf("Error unmarshalling configuration: %s", err)
	}

	log.Println("Parameters loaded are:")
	for _, k := range viper.AllKeys() {
		log.Printf("\t%s=%v\n", k, viper.Get(k))
	}
}
