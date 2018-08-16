package main

import (
	"fmt"

	"Saragenda/agenda"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("testconfig") // name of config file (without extension)
	viper.AddConfigPath(".")               // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	agenda.SetManager(agenda.NewViperWrapper(), agenda.NewMemoryStore())
	agenda.LoadChambres()
	fmt.Printf("\n%v\n", agenda.GetChambres())
}
