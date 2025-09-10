package c_default

import "github.com/pkg/errors"

func NoData() error {
	return errors.New("no data")
}
