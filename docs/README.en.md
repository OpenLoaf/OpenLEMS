<p align="center">
  <img src="images/lems-logo.png" alt="LEMS Logo" width="112" />
</p>

<h1 align="center">LEMS</h1>

<p align="center">
  Local Energy Management System
</p>

<p align="center">
  A lightweight local EMS runtime platform for distributed energy, microgrids, and energy storage cabinet scenarios
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.25%2B-00ADD8?logo=go&logoColor=white" alt="Go 1.25+" />
  <img src="https://img.shields.io/badge/License-AGPL--3.0--only-0E8A16" alt="License AGPL-3.0-only" />
  <img src="https://img.shields.io/badge/Platform-Linux%20%7C%20macOS%20%7C%20ARM-1F6FEB" alt="Platform Linux macOS ARM" />
</p>

<p align="center">
  <a href="https://lems.hexems.com/">Online Demo</a>
  ·
  <a href="#quick-start">Quick Start</a>
  ·
  <a href="#ui-preview">UI Preview</a>
  ·
  <a href="../CONTRIBUTING.md">Contributing</a>
</p>

`LEMS` is a local Energy Management System built for distributed energy, microgrids, energy storage cabinets, and edge energy gateway scenarios. It brings heterogeneous devices such as PCS, BMS, meters, PV, loads, and GPIO into one runtime through pluggable drivers, protocols, strategies, and data services, covering device integration, real-time monitoring, alarm handling, strategy control, automation linkage, tariff management, and external data integration.

It is not just a dashboard for charts. It is an EMS runtime platform that can actually be deployed on edge controllers, industrial gateways, industrial PCs, or on-site servers.

## Project Preview

<p align="center">
  <img src="images/energy-system-overview.png" alt="LEMS energy overview" width="88%" />
</p>

## Project Overview

In a typical distributed energy project, the site usually contains:

- PCS, BMS, meters, PV, loads, and other devices from different vendors
- Different communication methods and industrial protocols
- Different control logic, runtime constraints, and strategy goals
- Telemetry, alarms, and operational data that must be pushed to upper-level platforms

`LEMS` addresses the core challenge of fragmented devices, protocols, data, and control. It consolidates these capabilities into one unified system:

- Southbound device integration: connect on-site devices through driver and protocol plugins
- Unified runtime model: bring points, states, alarms, and strategies into a common runtime
- Northbound capabilities: expose Web, MQTT, Modbus TCP, logs, and statistical interfaces

## Reference Architecture

### Energy Storage Mode Architecture

The following diagram shows the local EMS network topology for energy storage scenarios, suitable for storage cabinets, multi-cabinet deployments, and edge data acquisition gateways.

<p align="center">
  <img src="images/energy-storage-architecture.png" alt="Energy storage mode architecture" width="88%" />
</p>

### Microgrid Mode Architecture

The following diagram shows how wind, solar, storage, load, touch panels, the cloud platform, and the mobile app are integrated in a microgrid scenario.

<p align="center">
  <img src="images/microgrid-architecture.png" alt="Microgrid mode architecture" width="88%" />
</p>

## UI Preview

### Energy Storage Cabinet Detail

The cabinet detail page shows the internal topology, runtime status, and trend charts of an energy storage cabinet.

<p align="center">
  <img src="images/energy-system-cabinet-detail.png" alt="Energy storage cabinet detail" width="80%" />
</p>

### Feature Entry and Real-Time Data

The system provides a unified feature entry page. The real-time data page shows device status, sub-device structure, and key points for operations teams.

<p align="center">
  <img src="images/programs-feature-menu.png" alt="Feature entry" width="48%" />
  <img src="images/realtime-data.png" alt="Real-time data" width="48%" />
</p>

### Alarms and Logs

The alarm page is suitable for checking active faults and historical issues, while the log page helps locate communication, push, and runtime details.

<p align="center">
  <img src="images/alarm-list.png" alt="Alarm list" width="48%" />
  <img src="images/log-viewer.png" alt="Log viewer" width="48%" />
</p>

### Statistics and Remote Management

The statistics page is used for device comparison and trend analysis. The remote management page manages external services such as MQTT and Modbus.

<p align="center">
  <img src="images/statistics-analysis.png" alt="Statistics analysis" width="48%" />
  <img src="images/remote-management-mqtt.png" alt="Remote management MQTT" width="48%" />
</p>

<p align="center">
  <img src="images/remote-management-modbus.png" alt="Remote management Modbus" width="48%" />
  <img src="images/automation-control.png" alt="Automation control" width="48%" />
</p>

### Tariffs, Protocols, Drivers, and Strategies

These pages show that the system is not only for monitoring, but also for configuration, control, and management.

<p align="center">
  <img src="images/price-management.png" alt="Tariff management" width="48%" />
  <img src="images/protocol-configuration.png" alt="Protocol configuration" width="48%" />
</p>

<p align="center">
  <img src="images/driver-manager.png" alt="Driver manager" width="48%" />
  <img src="images/strategy-management.png" alt="Strategy management" width="48%" />
</p>

### System Runtime Status

The system page shows CPU, memory, disk, network traffic, and uptime information, which is useful for post-deployment operations inspection.

<p align="center">
  <img src="images/system-management.png" alt="System management" width="80%" />
</p>

## Core Capabilities

- Pluggable driver system: easy to extend for different brands, models, and field device abstractions
- Pluggable protocol system: the repository currently includes `Modbus(TCP/RTU)`, `CANBus`, `GPIO`, and more
- Real-time monitoring and alarms: supports device status, power, energy, SOC, timestamps, alarm records, and historical logs
- Strategies and automation: supports microgrid strategies, energy storage strategies, automation tasks, and local control parameter configuration
- Tariff and cost optimization: supports time-of-use tariff configuration for peak shaving, valley filling, and cost management
- External integration: supports MQTT uplink, Modbus TCP export, and remote service management
- Edge-friendly deployment: suitable for industrial gateways, industrial PCs, and edge servers
- Built-in Web management UI: accessible after build, with interactions optimized for tablet touch scenarios

## Use Cases

- Energy storage cabinet EMS: single-cabinet operation, intra-cabinet linkage, status monitoring, and power control
- Multi-cabinet EMS: one edge gateway managing multiple storage systems
- Microgrid EMS: unified access and coordinated dispatch for wind, solar, storage, and load
- Integration gateway: normalize on-site heterogeneous device data before forwarding it to third-party platforms
- Edge O&M platform: provide alarms, logs, automation tasks, tariff management, and system status visibility

## Representative Components Included

### Device Drivers

- `Pylon Tech US108` BMS
- `Pylon Checkwatt` energy storage related driver
- `Elecod MAC` PCS
- `Star Charge 100E` PCS
- `GPIO In/Out` base drivers
- `Ammeter / PV / ESS / Load` demo drivers

### Protocols and Services

- `plugins/plug_protocols/p_modbus`
- `plugins/plug_protocols/p_canbus`
- `plugins/plug_protocols/p_gpiod`
- `services/s_export_mqtt`
- `services/s_export_modbus`
- `services/s_automation`
- `services/s_price`
- `services/s_policy`

## Algorithms and Strategy Capabilities

The "algorithm" capabilities in this repository are not just a few scheduled rules. They are formed by time-slot matching, strategy parameter models, a strategy execution framework, and prediction algorithm interfaces.

- Time slot and priority matching: the tariff service filters enabled configurations first, then picks the currently effective tariff by date range, time range, and priority; cross-day ranges are handled correctly
- Local strategy parameter models: exposes typical EMS decision parameters such as SOC limits, charge/discharge efficiency, transformer capacity, safety factors, demand control, and dynamic expansion
- Energy storage strategy execution framework: supports date ranges, time ranges, and hourly target setpoints, and can dispatch controls based on device lists, device types, and current status
- Automation linkage mechanism: supports time-based and condition-based task triggering to connect strategy decisions with field actions
- Forecasting extension interfaces: `@cpp/hexlib` provides a `C++ / CGO` based `MPC + Kalman` predictor interface suitable for forecasting power, load, generation, and other time series

A more precise description is that `LEMS` already provides algorithmic foundations, a strategy framework, and extensible forecasting capabilities, with room reserved for more advanced closed-loop microgrid optimization and energy storage dispatch.

## Online Demo

- Demo URL: [https://lems.hexems.com/](https://lems.hexems.com/)
- Admin password: `888888`

For the best experience, use landscape tablet mode. The current UI layout, touch interactions, and core O&M flows have been optimized for tablet scenarios.

## Quick Start

After cloning the public repository, you can build directly in `application/`.

### Requirements

- Go `1.25+`
- Linux or macOS is recommended for development
- If CGO-related capabilities are enabled, install `GCC` or `Clang`

### 1. Clone the Repository

```bash
git clone git@github.com:OpenLoaf/OpenLEMS.git
cd OpenLEMS
```

### 2. Sync Workspace Dependencies

```bash
go work sync
```

### 3. Build the Main Program

```bash
cd application
go build -o ./bin/lems .
```

If you want to validate the entire Go workspace in one shot, run this in the repository root:

```bash
./script/build-workspace.sh
```

### 4. Start the Web Management UI

```bash
cd application
go run . --web=true --driver-path=resources/driver
```

After startup, the default address is:

- `http://localhost:8000`

### 5. Start in Demo Mode

```bash
cd application
go run . --web=true --demo=true --driver-path=resources/driver
```

## Common Startup Flags

| Flag | Short | Description |
| --- | --- | --- |
| `--web` | `-w` | Enable the Web service |
| `--demo` | - | Mark the runtime as demo mode |
| `--driver-path` | `-dp` | Driver file directory |
| `--device-name` | `-d` | Device configuration name |
| `--runtime-path` | `-rp` | Runtime database path |
| `--db-path` | `-cp` | Configuration database path |
| `--language` | `-l` | Global language |
| `--profile` | - | Configuration profile, such as `dev` or `prod` |
| `--force` | - | Force startup by ignoring PID checks |

## Repository Structure

```text
OpenLEMS/
├── application/             # Main program entry, APIs, controllers, and Web service
├── common/                  # Shared interfaces, base device/protocol types, and utilities
├── plugins/
│   ├── plug_drivers/        # Device driver plugins
│   ├── plug_protocols/      # Communication protocol plugins
│   ├── plug_policy/         # Strategy plugins
│   ├── plug_push/           # Push plugins
│   └── plug_storages/       # Storage plugins
├── services/                # Service orchestration layer, such as drivers, automation, tariffs, and remote export
├── docs/                    # Project docs and screenshot assets
└── script/                  # Workspace build and validation scripts
```

## Development and Extension

### Build the Entire Workspace

```bash
./script/build-workspace.sh
```

### Driver Development

Drivers are organized in a plugin-based pattern. It is recommended to start from the existing implementations in:

- `plugins/plug_drivers/bms/`
- `plugins/plug_drivers/pcs/`
- `plugins/plug_drivers/ess/`
- `plugins/plug_drivers/gpio/`
- `plugins/plug_drivers/demo/`

### Protocol and Service Development

If you want to extend protocols or external data integration capabilities, start with:

- `plugins/plug_protocols/p_modbus`
- `plugins/plug_protocols/p_canbus`
- `services/s_export_mqtt`
- `services/s_export_modbus`
- `services/s_automation`
- `services/s_price`

## Related Docs

- [Software positioning and capability overview](ems-plan-software-summary.md)
- [Remote API docs](../application/api/remote/README.md)
- [Automation service guide](../services/s_automation/README.md)
- [MQTT data push guide](../services/s_export_mqtt/README.md)
- [Modbus TCP export guide](../services/s_export_modbus/README.md)

## Contributing

Issues and pull requests are welcome.

Before submitting, it is recommended to read:

- [Contribution Guide](../CONTRIBUTING.md)

Local validation commands:

```bash
./script/build-workspace.sh
./script/test-workspace.sh
```

## Business Contact

- Email: `zz@hexems.com`

For project cooperation, solution consulting, private deployment, or technical exchange, contact us by email or scan the WeChat QR code below:

<p align="center">
  <img src="images/business-wechat-qr.png" alt="Business WeChat QR code" width="240" />
</p>

## License

This project is licensed under [AGPL-3.0-only](../LICENSE).
