#!/bin/bash
echo "Building for ARM"

export CGO_ENABLED=1
cd plug_protocol_gpio_sysfs || exit 1
go build -o ../arm-build/gpio export_function.go || exit 1
cd ..

#./generate_all_plugin.sh arm || exit 1
./arm-build/sync.sh zlg 192.168.0.124 /home/zlg/ems
