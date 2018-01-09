/*
Copyright © 2016 Samsung CNCT

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var configFilename string
var ExitCode int

// init the zabra config viper instance
var zabraConfig = viper.New()

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "zabra",
	Short: "CLI for patching github repos",
	Long:  `zabra is a command line interface for cloning and patching a set of github repos`,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initZabraConfig)

	RootCmd.SetHelpCommand(helpCmd)

	RootCmd.PersistentFlags().StringVarP(
		&configFilename,
		"config",
		"c",
		"",
		"config file")
}

// Initializes zabraConfig to use flags, ENV variables and finally configuration files (in that order).
func initZabraConfig() {
	zabraConfig.BindPFlag("config", RootCmd.Flags().Lookup("config"))

	zabraConfig.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	zabraConfig.SetEnvPrefix("ZABRA") // prefix for env vars to configure cluster
	zabraConfig.AutomaticEnv()        // read in environment variables that match

	zabraConfig.SetConfigName("config")        // name of config file (without extension)
	zabraConfig.AddConfigPath("$HOME/.zabra/") // path to look for the config file in
	zabraConfig.AddConfigPath(".")             // optionally look for config in the working directory

	configFilename := zabraConfig.GetString("config")
	if configFilename != "" { // enable ability to specify config file via flag
		zabraConfig.SetConfigFile(configFilename)
	}

	// If a config file is found, read it in.
	if err := zabraConfig.ReadInConfig(); err == nil {
		fmt.Println("INFO: Using zabra config file:", zabraConfig.ConfigFileUsed())
	}

	zabraConfig.SetDefault("config", "~/.zabra/config.yaml")
}
