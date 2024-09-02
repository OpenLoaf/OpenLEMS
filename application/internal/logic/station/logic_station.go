package station

import "application/internal/service"

type sStation struct {
}

func init() {
	service.RegisterStation(&sStation{})
}
