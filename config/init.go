package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
)

var AppConfiguration AppConfigurations
var InfraConfiguration InfraConfigurations
var ProfileConfiguration ProfileConfigurations
var DB *gorm.DB

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	viper.AddConfigPath("./config")
	viper.SetConfigType("yml")

	viper.SetConfigName("config-profile.yml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf(fmt.Sprintf("Error reading config file: %v", err.Error()))
	}

	err := viper.Unmarshal(&ProfileConfiguration)

	if err != nil {
		log.Println(err.Error())
	}

	if ProfileConfiguration.Profile.Active != "" {
		activeProfile := ProfileConfiguration.Profile.Active

		infraYaml := fmt.Sprintf("config-infra-%v.yml", activeProfile)
		appYaml := fmt.Sprintf("config-app-%v.yml", activeProfile)

		viper.SetConfigName(infraYaml)

		if err := viper.ReadInConfig(); err != nil {
			log.Println(err.Error())
		}

		err = viper.Unmarshal(&InfraConfiguration)

		if err != nil {
			log.Println(err.Error())
		}

		viper.SetConfigName(appYaml)

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf(fmt.Sprintf("Error reading config file: %v", err.Error()))
		}

		err = viper.Unmarshal(&AppConfiguration)

		if err != nil {
			log.Println(err.Error())
		}
	}

	DB, err = DBConn()

	if err != nil {
		log.Fatal(err)
	}
}
