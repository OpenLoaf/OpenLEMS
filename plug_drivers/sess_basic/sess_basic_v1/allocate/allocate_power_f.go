package allocate

import (
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/optimize/convex/lp"
)

type SessBasic struct {
	Id                string
	Name              string
	Soc               int
	CurrentPower      float64   // 当前功率（充电负值，放电正值）
	RatedPower        int       // 额定功率
	MaxDischargePower float64   // 最小功率（充电最大功率）
	MaxChargePower    float64   // 最大功率（放电最大功率）
	CycleCount        int       // 循环次数
	EfficiencyCurve   []float64 // 效率曲线 从10%、20%一直到100% 共10条
}

/*
AllocatePower
tol 容差用来确定算法在寻找最优解时的精度。具体来说，当最大约简成本（reduced cost）低于 tol 时，算法认为已经找到了最优解，并终止计算。
*/
func AllocatePower(totalPower, tol float64, efficiencySegment int, showLog bool, cabinets []*SessBasic) ([]float64, error) {
	if len(cabinets) == 0 {
		return nil, gerror.New("cabinets is empty")
	}

	isCharge := totalPower < 0
	if isCharge {
		// 功率为负数，表示充电,变成正数
		totalPower = -totalPower
	}
	if len(cabinets) == 1 { // 处理只有单个柜子的情况
		if isCharge {
			maxChargePower := cabinets[0].MaxChargePower
			if totalPower < maxChargePower {
				return []float64{totalPower}, nil
			}
			return []float64{maxChargePower}, nil
		} else {
			maxDischargePower := cabinets[0].MaxDischargePower
			if totalPower < maxDischargePower {
				return []float64{totalPower}, nil
			}
			return []float64{maxDischargePower}, nil
		}
	}

	totalCycleCount := 0 // 总循环次数
	for _, ess := range cabinets {
		totalCycleCount += ess.CycleCount
	}
	numCabinets := len(cabinets)
	numVars := numCabinets
	if efficiencySegment != 0 {
		numVars = numCabinets * efficiencySegment
	}

	c := make([]float64, numVars)
	// 不等式约束,将会有总的功率约束和每个柜子的功率约束
	g := mat.NewDense(numCabinets+1, numVars, nil)
	h := make([]float64, numCabinets+1)

	for i, cabinet := range cabinets {
		// soc 权重
		socWeight := 0.0
		if isCharge {
			// 如果是充电，soc越低，权重越高
			socWeight = float64(100.0-float64(cabinet.Soc)) * 2
		} else {
			socWeight = float64(cabinet.Soc) * 2
		}

		// 循环权重
		cycleWeight := 100.0 - float64(cabinet.CycleCount)/float64(totalCycleCount)*100
		// 功率权重
		powerWeight := 0.0
		if isCharge {
			powerWeight = cabinet.MaxChargePower / float64(cabinet.RatedPower) * 100
		} else {
			powerWeight = cabinet.MaxDischargePower / float64(cabinet.RatedPower) * 100
		}

		usageWeight := 0.0
		if cabinet.CurrentPower != 0 {
			usageWeight = 10
		}
		if showLog {
			fmt.Printf("柜子%d: socWeight: %.2f, cycleWeight: %.2f, powerWeight: %.2f, usageWeight: %.2f\n", i, socWeight, cycleWeight, powerWeight, usageWeight)
		}

		if efficiencySegment == 0 {
			c[i] = -totalPower - powerWeight - socWeight - cycleWeight - usageWeight
			g.Set(i, i, 1)
		} else {
			tempWeight := -totalPower - powerWeight - socWeight - cycleWeight - usageWeight
			for j := 0; j < efficiencySegment; j++ {
				// 能效权重
				c[i*efficiencySegment+j] = tempWeight - cabinet.EfficiencyCurve[j]
				// 设置每个柜子的功率约束
				g.Set(i, i*efficiencySegment+j, 1)
			}
		}

		if isCharge {
			h[i] = cabinet.MaxChargePower // 如果充电，使用最大功率
		} else {
			h[i] = cabinet.MaxDischargePower //
		}

	}
	for i := 0; i < numVars; i++ {

		// 设置总的功率约束
		g.Set(numCabinets, i, 1)
	}
	h[numCabinets] = totalPower

	if showLog {
		fmt.Printf("C目标函数: %v\n", c)
		fmt.Println("G不等式:")
		matPrint(g)
		fmt.Printf("h不等式右侧: %v\n", h)
	}

	cNew, aNew, bNew := lp.Convert(c, g, h, nil, nil)

	_, x, err := lp.Simplex(cNew, aNew, bNew, tol, nil)

	if err != nil {
		return nil, err
	}

	if showLog {
		fmt.Println("resultX: ", x)
	}

	/*
		[1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12]
		[[ 1,  2,  3,  4], [ 5,  6,  7,  8], [ 9, 10, 11, 12]]
		[10, 26, 42]
	*/
	resultX := make([]float64, numCabinets)
	sumResultX := 0.0
	// 把结果转换为柜子的数组
	for i := 0; i < numVars; i++ {
		var j = i
		if efficiencySegment != 0 {
			j = i / efficiencySegment
		}
		if j == 0 {
			resultX[0] += x[i]
			sumResultX += x[i]
		} else {
			resultX[j] += x[i]
			sumResultX += x[i]
		}
	}
	if sumResultX > totalPower {
		return nil, gerror.Newf("功率分配结果: %v异常！功率分配总和:%v超过需求总功率:%v", resultX, sumResultX, totalPower)
	}
	if showLog {
		fmt.Println("功率分配结果: ", resultX, "请求功率：", totalPower, "累计功率：", sumResultX)
	}
	return resultX, nil
}

// 辅助函数：打印矩阵
func matPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Excerpt(0))
	fmt.Printf("%v\n", fa)
}
