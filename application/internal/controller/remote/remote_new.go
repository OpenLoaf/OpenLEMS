package remote

import (
	"application/api/remote"
)

type ControllerV1 struct{}

func NewV1() remote.IRemoteV1 {
	return &ControllerV1{}
}
