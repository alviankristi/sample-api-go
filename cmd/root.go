package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/alviankristi/catalyst-backend-task/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Conf *config.Config
)

var RootCmd = &cobra.Command{
	Use:   "go-base",
	Short: "Catalyst Backend Task Boilerplate",
	Long:  `Catalyst Backend Task Boilerplate with passwordless authentication.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(LoadConfig)
}

// initConfig reads in config file .
func LoadConfig() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	err := viper.ReadInConfig()
	//return error if config unable to read
	if err != nil {
		log.Fatalf("error config file: %p \n", err)
	}

	//return error if config unable to cast to config struct
	err = viper.Unmarshal(&Conf)
	if err != nil {
		log.Fatalf("unable to decode into config struct, %v", err)
	}

}
