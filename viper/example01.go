package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func Example01() {
	viper.SetDefault("ContentDir", "content")
	viper.SetDefault("LayoutDir", "layouts")
	viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})

	fmt.Println("ContentDir: ", viper.GetString("ContentDir"))
	fmt.Println("LayoutDir: ", viper.GetString("LayoutDir"))
	fmt.Println("Taxonomies: ", viper.GetStringMapString("Taxonomies"))
}
