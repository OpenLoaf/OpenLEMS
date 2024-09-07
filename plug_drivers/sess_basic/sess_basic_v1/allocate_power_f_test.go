package sess_basic_v1

import (
	"fmt"
	"testing"
)

func TestAllocatePower(t *testing.T) {
	//EfficiencyCurve := []float64{100, 100, 100, 100, 100, 100, 100, 100, 100, 100}
	essList := []*sSessBasic{
		{Name: "1", Soc: 55, RatedPower: 100, MaxDischargePower: 100, MaxChargePower: 100, CycleCount: 15, EfficiencyCurve: nil},
		{Name: "2", Soc: 80, RatedPower: 100, MaxDischargePower: 100, MaxChargePower: 100, CycleCount: 100, EfficiencyCurve: nil},
	}

	power, err := AllocatePower(50, 200, 0, true, essList)
	if err != nil {
		t.Error(err)
	}
	for i, p := range power {
		ess := essList[i]
		fmt.Printf("Id: %s, Power: %f\n", ess.Name, p)
	}

}
