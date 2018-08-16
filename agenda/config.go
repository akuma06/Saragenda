package agenda

import (
	"github.com/spf13/viper"
)

type ViperWrapper struct {
	v *viper.Viper
}

func (vw *ViperWrapper) UnmarshalKey(key string, rawVal interface{}) error  {
	return vw.v.UnmarshalKey(key, rawVal)
}

func NewViperWrapper() *ViperWrapper {
	return &ViperWrapper{
		v: viper.GetViper(),
	}
}