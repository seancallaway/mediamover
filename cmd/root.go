package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFile string

var rootCmd = &cobra.Command{
	Use:   "mediamover",
	Short: "mediamover is a tool for organizing your media files",
	Long: `A tool for identifying and renaming movie and television show
    files using The Movie Database (https://themoviedb.org)`,
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
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Produce verbose output")
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
}

func loadConfig() {
	viper.SetConfigFile(configFile)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to load config file: %s", viper.ConfigFileUsed())
	}
}
