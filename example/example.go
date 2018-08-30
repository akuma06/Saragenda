package example

import (
	"fmt"

	"Saragenda"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("testconfig") // name of config file (without extension)
	viper.AddConfigPath(".")               // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	err = saragenda.SetManager(saragenda.NewViperWrapper(), saragenda.NewMemoryStore())
	if err != nil {
		fmt.Println(err)
	}
	err = saragenda.LoadChambres()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\n%v\n", saragenda.GetChambres())
}
