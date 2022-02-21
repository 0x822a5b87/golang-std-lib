package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func Example02() {
	viper.SetConfigName("config")         // name of config file (without extension)
	viper.SetConfigType("yaml")           // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/appname/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	viper.AddConfigPath(".")              // optionally look for config in the working directory
	err := viper.ReadInConfig()           // Find and read the config file
	if err != nil {                       // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
		} else {
			// Config file was found but another error was produced
		}
	}

	// Config file found and successfully parsed
	fmt.Println("kind = ", viper.GetString("kind"))
	fmt.Println("apiVersion = ", viper.GetString("apiVersion"))
	fmt.Println("name = ", viper.GetString("name"))
	fmt.Println("metadata.labels.location = ", viper.GetString("metadata.labels.location"))

	var dest []map[string]interface{}
	err = viper.UnmarshalKey("spec.sources", &dest)
	if err != nil {
		return
	}
	for i, contents := range dest {
		fmt.Printf("[%d]\n", i)
		for k, v := range contents {
			fmt.Println("k = ", k, ", value = ", v)
		}
	}

	keys := viper.AllKeys()
	for k, v := range keys {
		fmt.Println("keys: k = ", k, ", value = ", v)
	}

	specSources := viper.Get("spec.sources")
	fmt.Println("spec.sources: ", specSources)
	map0 := specSources.([]interface{})
	for k, v := range map0 {
		v0 := v.(map[interface{}]interface{})
		fmt.Println("map0: k0 = ", k, ", v0 = ", v0)
	}

	viper.Set("metadata.labels.location", "sz")
	viper.WriteConfig()
}
