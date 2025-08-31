package c_device

import "common/c_base"

type SBasePolicyParamImpl[T c_base.IDriver] struct {
	SBasePolicyImpl[T]
}

func (s *SBasePolicyParamImpl[T]) name() {

}
