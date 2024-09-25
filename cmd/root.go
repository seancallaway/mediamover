package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	Default struct {
		ApiKey    string `mapstructure:"api_key"`
		TvRoot    string `mapstructure:"tv_root"`
		MovieRoot string `mapstructure:"movie_root"`
	}
}

var configFile string

var rootCmd = &cobra.Command{
	Use:   "mediamover",
	Short: "mediamover is a tool for organizing your media files",
	Long: `A tool for identifying and renaming movie and television show
    files using The Movie Database (https://themoviedb.org)`,
	Run: func(cmd *cobra.Command, args []string) {
		// do stuff here.
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(loadConfig)

	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "config.ini", "The configuration file")
}

func loadConfig() {
	viper.SetConfigFile(configFile)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Unable to load config file:", viper.ConfigFileUsed())
	}
}
