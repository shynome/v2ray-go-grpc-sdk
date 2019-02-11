package core

import (
	"reflect"
	"sync"

	"v2ray.com/core/features"
)

type resolution struct {
	deps     []reflect.Type
	callback interface{}
}

// Instance combines all functionalities in V2Ray.
type Instance struct {
	access             sync.Mutex
	features           []features.Feature
	featureResolutions []resolution
	running            bool
}
