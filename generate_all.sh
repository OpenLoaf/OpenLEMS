#!/bin/bash



# 获取当前脚本所在目录
SOURCE_DIR="$(dirname "$(realpath "$0")")"
ARM_BUILD_DIR=$SOURCE_DIR/arm-build
echo "SOURCE_DIR: $SOURCE_DIR"
cd $SOURCE_DIR || exit 1




if [ "$1" == "arm" ]; then
  echo "Building for ARM"
  # 编译ems
  cd application || exit 1
  gf build || exit 1
  # 拷贝到arm-build目录
  cp bin/ems $ARM_BUILD_DIR || exit 1
  DRIVER_DIR=$SOURCE_DIR/arm-build/drivers
else
  DRIVER_DIR=$SOURCE_DIR/application/resources/drivers
fi



cd ${SOURCE_DIR}/plug_drivers || exit 1
for dir in `ls $SOURCE_DIR/plug_drivers`; do
#    cd "$currentPath/$(dirname $dir)" || exit 1
    if [ -f $dir/main.go ]; then
      echo "===> Running go generate in $dir "
      make Project=$dir OutDir=$DRIVER_DIR RunDir=$DRIVER_DIR BuildBin=0 || exit 1
    fi
done
