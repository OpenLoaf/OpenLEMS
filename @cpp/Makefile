.PHONY: all build clean install generate

# 检测操作系统
ifeq ($(OS),Windows_NT)
    DETECTED_OS := Windows
    RM_CMD := rmdir /s /q
    MKDIR_CMD := mkdir
    SEP := \\
    LIB_EXT := .dll
    LIB_PREFIX := 
    SHELL := cmd
else
    DETECTED_OS := $(shell uname -s)
    RM_CMD := rm -rf
    MKDIR_CMD := mkdir -p
    SEP := /
    LIB_EXT := .so
    LIB_PREFIX := lib
    SHELL := /bin/bash
endif

BUILD_DIR := build
LIB_NAME := $(LIB_PREFIX)hexlib$(LIB_EXT)

all: build

build:
	@echo "Building for $(DETECTED_OS)..."
	cmake -S . -B $(BUILD_DIR) -DCMAKE_BUILD_TYPE=Release
	cmake --build $(BUILD_DIR) --config Release
	@echo "Build completed. Library: $(BUILD_DIR)$(SEP)$(LIB_NAME)"

install:
	cmake --install $(BUILD_DIR)

clean:
	@echo "Cleaning build directory for $(DETECTED_OS)..."
ifeq ($(OS),Windows_NT)
	@if exist $(BUILD_DIR) $(RM_CMD) $(BUILD_DIR)
else
	@$(RM_CMD) $(BUILD_DIR)
endif

generate:
ifeq ($(OS),Windows_NT)
	@echo "Generating bindings for Windows..."
	@if exist bindings\gen.sh bash bindings/gen.sh
	@if exist bindings\gen.bat bindings\gen.bat
else
	@echo "Generating bindings for $(DETECTED_OS)..."
	@bash bindings/gen.sh
endif

# 显示当前配置信息
info:
	@echo "Detected OS: $(DETECTED_OS)"
	@echo "Build Directory: $(BUILD_DIR)"
	@echo "Library Name: $(LIB_NAME)"
	@echo "Library Path: $(BUILD_DIR)$(SEP)$(LIB_NAME)"
	@echo "Remove Command: $(RM_CMD)"
	@echo "Make Directory Command: $(MKDIR_CMD)"


