package conf

import "github.com/hdget/hdsdk"

type Config struct {
	hdsdk.Config `mapstructure:",squash"`
}
