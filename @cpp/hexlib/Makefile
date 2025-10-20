.PHONY: build clean help

# 目录配置
SRC_DIR=src
INCLUDE_DIR=$(SRC_DIR)/include
BUILD_DIR=build
CPP_SOURCES=$(SRC_DIR)/hexlib.cpp $(SRC_DIR)/mpc_kalman.cpp
CPP_OBJECTS=$(BUILD_DIR)/hexlib.o $(BUILD_DIR)/mpc_kalman.o
CPP_LIB=$(BUILD_DIR)/libhexlib.a

# 编译器选项
CXX?=clang++
CXXFLAGS=-std=c++11 -fPIC -Wall -O2
CFLAGS=-I$(INCLUDE_DIR)

# 检查操作系统
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Darwin)
	# macOS 特定配置
	LDFLAGS=-dynamiclib
	SHARED_LIB=$(BUILD_DIR)/libhexlib.dylib
else ifeq ($(UNAME_S),Linux)
	# Linux 特定配置
	LDFLAGS=-shared
	SHARED_LIB=$(BUILD_DIR)/libhexlib.so
else
	# Windows 特定配置（如需要）
	LDFLAGS=-shared
	SHARED_LIB=$(BUILD_DIR)/hexlib.dll
endif

# 默认目标
all: $(CPP_LIB)

# 创建build目录
$(BUILD_DIR):
	@mkdir -p $(BUILD_DIR)

# 编译C++对象文件
$(BUILD_DIR)/%.o: $(SRC_DIR)/%.cpp $(INCLUDE_DIR)/hexlib.h | $(BUILD_DIR)
	@echo "Compiling $<..."
	$(CXX) $(CXXFLAGS) $(CFLAGS) -c $< -o $@
	@echo "✓ Compiled: $@"

# 生成静态库
$(CPP_LIB): $(CPP_OBJECTS)
	@echo "Creating static library..."
	ar rcs $(CPP_LIB) $(CPP_OBJECTS)
	@echo "✓ Static library built: $(CPP_LIB)"

# 生成动态库（可选，需要时取消注释下面的targets）
shared: $(SHARED_LIB)

$(SHARED_LIB): $(CPP_OBJECTS)
	@echo "Creating shared library..."
	$(CXX) $(LDFLAGS) -o $(SHARED_LIB) $(CPP_OBJECTS)
	@echo "✓ Shared library built: $(SHARED_LIB)"

# 构建目标
build: $(CPP_LIB)
	@echo ""
	@echo "=========================================="
	@echo "hexlib C++ library build complete"
	@echo "=========================================="
	@echo "Location: $(CPP_LIB)"
	@echo "Ready to use with CGO"

# 清理目标
clean:
	@echo "Cleaning hexlib C++ artifacts..."
	@rm -rf $(BUILD_DIR)
	@echo "✓ Clean complete"

# 帮助信息
help:
	@echo "hexlib Makefile - Available targets:"
	@echo ""
	@echo "  make build        - Build C++ static library (default)"
	@echo "  make shared       - Build C++ shared library"
	@echo "  make clean        - Remove C++ build artifacts"
	@echo "  make help         - Show this help message"
	@echo ""
	@echo "Configuration:"
	@echo "  CXX=$(CXX)"
	@echo "  CXXFLAGS=$(CXXFLAGS)"
	@echo "  Platform: $(UNAME_S)"
	@echo ""
	@echo "Build artifacts are stored in: $(BUILD_DIR)/"
	@echo "Source files in: $(SRC_DIR)/"

# 显示编译信息
info:
	@echo "hexlib Build Information"
	@echo "========================="
	@echo "C++ Compiler: $(CXX)"
	@echo "C++ Flags: $(CXXFLAGS)"
	@echo "Include Directory: $(INCLUDE_DIR)"
	@echo "Build Directory: $(BUILD_DIR)"
	@echo "Source Files: $(CPP_SOURCES)"
	@echo "Platform: $(UNAME_S)"
	@echo "Static Library: $(CPP_LIB)"
	@echo "Shared Library: $(SHARED_LIB)"
