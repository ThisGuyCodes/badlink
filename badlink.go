package badlink

import "gopkg.in/yaml.v2"

func Thing(this yaml.Marshaler) yaml.Marshaler {
	return this
}
