package openf

import "os"

type Opener interface {
	func([]string) *os.File
}
