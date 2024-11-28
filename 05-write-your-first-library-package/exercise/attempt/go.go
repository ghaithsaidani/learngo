package attempt

import "runtime"


func version() string {
	return runtime.Version()
}