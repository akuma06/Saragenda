package saragenda

import (
	"reflect"
	"testing"
	"github.com/spf13/viper"
	"fmt"
)

func initViperTest() {
	viper.SetConfigName("testconfig") // name of config file (without extension)
	viper.AddConfigPath(".")               // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func TestViperWrapper_UnmarshalKey(t *testing.T) {
	initViperTest()
	type args struct {
		key    string
		rawVal interface{}
	}
	tests := []struct {
		name    string
		vw      *ViperWrapper
		args    args
		wantErr bool
	}{
		{
			"ParseValid",
			NewViperWrapper(),
			args{"chambres", &Chambres{}},
			false,
		},
		{
			"ParseInvalidKey",
			NewViperWrapper(),
			args{"invalidKey", &Chambres{}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.vw.UnmarshalKey(tt.args.key, tt.args.rawVal); (err != nil) != tt.wantErr {
				t.Errorf("ViperWrapper.UnmarshalKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewViperWrapper(t *testing.T) {
	initViperTest()
	tests := []struct {
		name string
		want *ViperWrapper
	}{
		{
			"ViperWrapper",
			&ViperWrapper{
				viper.GetViper(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewViperWrapper(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewViperWrapper() = %v, want %v", got, tt.want)
			}
		})
	}
}
