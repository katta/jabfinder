package cmd

import (
  "fmt"
  "github.com/spf13/cobra/doc"
  "log"
  "os"
  "github.com/spf13/cobra"
  "strings"

  homedir "github.com/mitchellh/go-homedir"
  "github.com/spf13/viper"

)


var cfgFile string
var genDoc bool


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
  Use:   "jabfinder",
  Short: "Finds the COVID vaccines available",
  Long: `Uses the CoWIN APIs to find the vaccines available
for the given coordinates.
`,

  //	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  if genDoc {
    err := doc.GenMarkdownTree(rootCmd, "./docs/")
    if err != nil {
      log.Fatal(err)
    }
  }

}

func init() {
  cobra.OnInitialize(initConfig)

  // Here you will define your flags and configuration settings.
  // Cobra supports persistent flags, which, if defined here,
  // will be global for your application.

  rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.jabfinder.yaml)")
  rootCmd.PersistentFlags().BoolVar(&genDoc, "generateDoc", false, "Set to true to generate Documents (Must be run from SCM cloned location")

  // Cobra also supports local flags, which will only run
  // when this action is called directly.
  rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


// initConfig reads in config file and ENV variables if set.
func initConfig() {
  if cfgFile != "" {
    // Use config file from the flag.
    viper.SetConfigFile(cfgFile)
  } else {
    // Find home directory.
    home, err := homedir.Dir()
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    // Search config in home directory with name ".jabfinder" (without extension).
    viper.AddConfigPath(home)
    viper.AddConfigPath(".")

    viper.SetEnvPrefix("JABF")
    replacer := strings.NewReplacer(".", "_")
    viper.SetEnvKeyReplacer(replacer)
    viper.AutomaticEnv()
    viper.SetConfigName(".jabfinder")
    viper.SetConfigType("yaml")
  }

  viper.AutomaticEnv() // read in environment variables that match

  // If a config file is found, read it in.
  if err := viper.ReadInConfig(); err == nil {
    fmt.Println("Using config file:", viper.ConfigFileUsed())
  }
}

