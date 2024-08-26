#!/bin/bash


./generate_all_plugin.sh arm || exit 1
./arm-build/sync.sh zlg 192.168.0.124 /home/zlg/ems
