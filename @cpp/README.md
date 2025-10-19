# @cpp: MPC Kalman Filter C++ Library + Go Bindings

This module contains a C++ implementation of MPC Kalman Filter for microgrid prediction with Go bindings generated using `c-for-go`.

## Features

- **MPC Kalman Filter**: Optimized for microgrid load/PV prediction
- **State Space Model**: [value, velocity, acceleration] with trend decay
- **Uncertainty Quantification**: Provides prediction confidence intervals
- **Online Learning**: Continuous adaptation to new measurements
- **C++ Performance**: High-speed computation with Go integration

## Structure

- `hexlib/src/include/hexlib.h`: C API header with MPC functions
- `hexlib/src/mpc_kalman.cpp`: MPC Kalman Filter implementation
- `hexlib/src/hexlib.cpp`: Basic utility functions
- `hexlib/hexlib.go`: Go bindings for C++ library
- `hexlib/hexlib_test.go`: Go test file with comprehensive tests
- `CMakeLists.txt` and `Makefile`: Build scripts

## API Functions

### Core MPC Functions

```c
// Create predictor from historical data
void* mpc_create_predictor(const double* historical_data, int data_count);

// Generate multi-step predictions with uncertainty
PredictionArray* mpc_predict(void* predictor, int prediction_steps);

// Update with new measurement
void mpc_update_measurement(void* predictor, double measurement);

// Get current filtered state
PredictionResult mpc_get_current_state(void* predictor);

// Cleanup
void mpc_free_predictor(void* predictor);
void mpc_free_prediction_array(PredictionArray* array);
```

## Usage

### Build the library:
```bash
cd @cpp
make build
```

### Run tests:
```bash
cd @cpp
./run_example.sh
```

### Go Integration Example:
```go
// Create predictor with historical data
historicalData := []float64{50.0, 50.0, 45.0, 40.0, 35.0}
predictor := hexlib.MpcCreatePredictor(&historicalData[0], len(historicalData))
defer hexlib.MpcFreePredictor(predictor)

// Get predictions
predictions := hexlib.MpcPredict(predictor, 10)
defer hexlib.MpcFreePredictionArray(predictions)

// Update with new measurement
hexlib.MpcUpdateMeasurement(predictor, 38.0)
```

## Integration with @internal/

The generated Go bindings can be imported and used within the `@internal/` directory for EMS prediction services. The MPC Kalman Filter provides:

1. **Real-time prediction**: Suitable for MPC control loops
2. **Uncertainty quantification**: Risk-aware decision making  
3. **Adaptive learning**: Continuous model improvement
4. **High performance**: C++ computation with Go convenience

Perfect for microgrid energy management systems requiring reliable short-term forecasting.


