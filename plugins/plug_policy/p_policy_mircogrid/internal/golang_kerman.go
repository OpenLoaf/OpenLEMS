package internal

import (
	"math"

	"github.com/konimarti/kalman"
	"github.com/konimarti/lti"
	"gonum.org/v1/gonum/mat"
)

// PredictionResult 与 C++ 的 PredictionResult 对齐
type PredictionResult struct {
	Value       float64
	Uncertainty float64
}

// Predictor 复刻 C++ 的三状态模型 [值, 速度, 加速度]
type Predictor struct {
	// 系统与噪声
	sys   lti.Discrete
	noise kalman.Noise

	// 滤波器与上下文（持久状态）
	filter kalman.Filter
	ctx    kalman.Context

	// 参数
	dt          float64
	trendDecay  float64
	nonNegative bool
}

// NewPredictorFromHistory 使用历史数据初始化（等价于 C++ 初始化+在线学习）
func NewPredictorFromHistory(historical []float64) *Predictor {
	p := &Predictor{
		dt:          1.0,
		trendDecay:  0.25,
		nonNegative: true,
	}

	// 状态转移矩阵 F (Ad)
	dt := p.dt
	td := p.trendDecay
	Ad := mat.NewDense(3, 3, []float64{
		1, dt, 0.5 * dt * dt,
		0, 1 - td*dt, dt,
		0, 0, 1 - 2*td*dt,
	})

	// 无控制输入（B, D 全零）
	Bd := mat.NewDense(3, 1, nil)
	C := mat.NewDense(1, 3, []float64{1, 0, 0})
	D := mat.NewDense(1, 1, nil)

	// 过程噪声 Q，测量噪声 R（与 C++ 保持一致）
	Q := mat.NewDense(3, 3, []float64{
		0.05, 0, 0,
		0, 0.3, 0,
		0, 0, 0.1,
	})
	R := mat.NewDense(1, 1, []float64{1.0})

	p.sys = lti.Discrete{
		Ad: Ad,
		Bd: Bd,
		C:  C,
		D:  D,
	}
	p.noise = kalman.Noise{
		Q: Q,
		R: R,
	}
	p.filter = kalman.NewFilter(p.sys, p.noise)

	// 初始状态 X0 与协方差 P0（与 C++ 保持一致）
	X0Dense := estimateInitialX(historical, p.dt)
	X0 := mat.NewVecDense(3, []float64{X0Dense.At(0, 0), X0Dense.At(1, 0), X0Dense.At(2, 0)})
	P0 := mat.NewDense(3, 3, []float64{
		5, 0, 0,
		0, 5, 0,
		0, 0, 5,
	})
	p.ctx = kalman.Context{
		X: X0,
		P: P0,
	}

	// 在线学习：对每个历史点执行一次测量更新（与 C++ 的 create_predictor 循环等价）
	u := mat.NewVecDense(1, nil) // 无控制输入
	for _, z := range historical {
		zv := mat.NewVecDense(1, []float64{z})
		p.filter.Apply(&p.ctx, zv, u) // 预测+更新
		p.clampNonNegative()
	}

	return p
}

// Update 注入新测量（等价于 C API mpc_update_measurement 的 predict+update）
func (p *Predictor) Update(measurement float64) {
	u := mat.NewVecDense(1, nil)
	z := mat.NewVecDense(1, []float64{measurement})
	p.filter.Apply(&p.ctx, z, u)
	p.clampNonNegative()
}

// Predict 预测未来 steps 步（与 C++ predictFuture 对齐），不改变内部状态
func (p *Predictor) Predict(steps int) []PredictionResult {
	if steps <= 0 {
		return nil
	}
	// 临时副本（不影响内部 ctx）
	X := mat.DenseCopyOf(p.ctx.X)
	P := mat.DenseCopyOf(p.ctx.P)

	Ad := p.sys.Ad
	Q := p.noise.Q

	results := make([]PredictionResult, 0, steps)
	for i := 0; i < steps; i++ {
		// x = Ad * x
		var nextX mat.Dense
		nextX.Mul(Ad, X)

		// P = Ad * P * Ad^T + Q
		var tmp mat.Dense
		tmp.Mul(Ad, P)
		var AdT mat.Dense
		AdT.CloneFrom(Ad.T())
		var nextP mat.Dense
		nextP.Mul(&tmp, &AdT)
		nextP.Add(&nextP, Q)

		X = &nextX
		P = &nextP

		val := X.At(0, 0)
		if p.nonNegative && val < 0 {
			val = 0
		}
		unc := math.Sqrt(math.Max(0, P.At(0, 0)))

		results = append(results, PredictionResult{
			Value:       val,
			Uncertainty: unc,
		})
	}
	return results
}

// CurrentState 与 C++ 的 getCurrentState 对齐
func (p *Predictor) CurrentState() PredictionResult {
	val := p.ctx.X.AtVec(0)
	if p.nonNegative && val < 0 {
		val = 0
	}
	unc := math.Sqrt(math.Max(0, p.ctx.P.At(0, 0)))
	return PredictionResult{
		Value:       val,
		Uncertainty: unc,
	}
}

// 可选：打开/关闭非负约束
func (p *Predictor) SetNonNegativeConstraint(enable bool) {
	p.nonNegative = enable
}

func (p *Predictor) clampNonNegative() {
	if !p.nonNegative {
		return
	}
	if p.ctx.X.AtVec(0) < 0 {
		p.ctx.X.SetVec(0, 0)
	}
}

// 从历史数据估计 X0 = [value, velocity, acceleration]^T
func estimateInitialX(data []float64, dt float64) *mat.Dense {
	n := len(data)
	if n == 0 {
		return mat.NewDense(3, 1, nil)
	}
	value := data[n-1]

	vel := 0.0
	if n >= 2 {
		vel = (data[n-1] - data[n-2]) / dt
	}
	acc := 0.0
	if n >= 3 {
		velPrev := (data[n-2] - data[n-3]) / dt
		acc = (vel - velPrev) / dt
	}
	return mat.NewDense(3, 1, []float64{value, vel, acc})
}
