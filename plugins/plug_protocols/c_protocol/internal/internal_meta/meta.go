package internal_meta

import (
	"common/c_base"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"math/big"
	"reflect"
	"strconv"
)

func MetaAddrDecString(p *c_base.Meta) string {
	return strconv.FormatUint(uint64(p.Addr), 10)
}

func MetaValueToString(p *c_base.Meta, value *gvar.Var) string {
	switch SystemTypeGetReflectKind(p.SystemType, p.ReadType, p.BitLength) {
	case reflect.Bool:
		b := value.Bool()
		if b {
			return "true"
		} else {
			return "false"
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.String:
		return fmt.Sprintf("%v", value.String())
	case reflect.Float32, reflect.Float64:
		return big.NewFloat(value.Float64()).Text('f', p.Precise)
	default:
		return fmt.Sprintf("%v", value.String())
	}
}
