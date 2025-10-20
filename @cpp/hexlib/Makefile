.PHONY: build clean help info shared prepare

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

# 创建build目录（作为依赖，只做一次）
$(BUILD_DIR):
	@mkdir -p $@

# 准备构建环境 - 自动创建必要的目录
prepare: | $(BUILD_DIR)
	@echo "✓ Build directory prepared: $(BUILD_DIR)/"

# 编译C++对象文件
# 必须先创建 build 目录
$(BUILD_DIR)/%.o: $(SRC_DIR)/%.cpp $(INCLUDE_DIR)/hexlib.h
	@mkdir -p $(BUILD_DIR)
	@echo "Compiling $<..."
	$(CXX) $(CXXFLAGS) $(CFLAGS) -c $< -o $@
	@echo "✓ Compiled: $@"

# 生成静态库
# 不需要显式依赖 $(BUILD_DIR)，因为 $(CPP_OBJECTS) 已经隐含了
$(CPP_LIB): $(CPP_OBJECTS)
	@echo "Creating static library..."
	ar rcs $@ $^
	@echo "✓ Static library built: $(CPP_LIB)"
	@echo ""
	@echo "=========================================="
	@echo "hexlib C++ library build complete"
	@echo "=========================================="
	@echo "Location: $(CPP_LIB)"
	@echo "Ready to use with CGO"
	@echo ""
	@ls -lh $(CPP_LIB)

# 生成动态库（可选）
$(SHARED_LIB): $(CPP_OBJECTS)
	@echo "Creating shared library..."
	$(CXX) $(LDFLAGS) -o $(SHARED_LIB) $(CPP_OBJECTS)
	@echo "✓ Shared library built: $(SHARED_LIB)"
	@echo ""
	@echo "=========================================="
	@echo "hexlib C++ shared library build complete"
	@echo "=========================================="
	@echo "Location: $(SHARED_LIB)"
	@ls -lh $(SHARED_LIB)

# 默认目标（静态库）
all: $(CPP_LIB)

# 构建目标（静态库，等同于 all）
build: $(CPP_LIB)

# 构建动态库
shared: $(SHARED_LIB)

# 清理目标
clean:
	@echo "Cleaning hexlib C++ artifacts..."
	@rm -rf $(BUILD_DIR)
	@echo "✓ Clean complete"

# 帮助信息
help:
	@echo ""
	@echo "╔══════════════════════════════════════════════════════════════════╗"
	@echo "║              hexlib Makefile - Available targets                 ║"
	@echo "╚══════════════════════════════════════════════════════════════════╝"
	@echo ""
	@echo "📦 Build Targets:"
	@echo "  make build        - Build C++ static library (default, recommended)"
	@echo "  make shared       - Build C++ shared library (.so/.dylib/.dll)"
	@echo "  make prepare      - Prepare build directory"
	@echo ""
	@echo "🧹 Maintenance:"
	@echo "  make clean        - Remove all C++ build artifacts"
	@echo "  make help         - Show this help message"
	@echo "  make info         - Show build configuration"
	@echo ""
	@echo "⚙️  Configuration:"
	@echo "  CXX               = $(CXX)"
	@echo "  CXXFLAGS          = $(CXXFLAGS)"
	@echo "  Platform          = $(UNAME_S)"
	@echo ""
	@echo "📁 Directories:"
	@echo "  Build artifacts   → $(BUILD_DIR)/"
	@echo "  Source files      → $(SRC_DIR)/"
	@echo "  Include files     → $(INCLUDE_DIR)/"
	@echo ""
	@echo "📝 Usage Examples:"
	@echo "  make              # Build static library (same as 'make build')"
	@echo "  make build        # Compile C++ and create libhexlib.a"
	@echo "  make shared       # Create shared library"
	@echo "  make clean        # Clean all artifacts and rebuild fresh"
	@echo ""

# 显示编译信息
info:
	@echo ""
	@echo "╔══════════════════════════════════════════════════════════════════╗"
	@echo "║                  hexlib Build Information                        ║"
	@echo "╚══════════════════════════════════════════════════════════════════╝"
	@echo ""
	@echo "🔧 Compiler Configuration:"
	@echo "  C++ Compiler:     $(CXX)"
	@echo "  C++ Version:      $$($(CXX) --version 2>/dev/null | head -1)"
	@echo "  C++ Flags:        $(CXXFLAGS)"
	@echo "  C Flags:          $(CFLAGS)"
	@echo ""
	@echo "📂 Paths:"
	@echo "  Source Directory: $(SRC_DIR)/"
	@echo "  Include Path:     $(INCLUDE_DIR)/"
	@echo "  Build Directory:  $(BUILD_DIR)/"
	@echo ""
	@echo "📄 Source Files:"
	@echo "  $(CPP_SOURCES)"
	@echo ""
	@echo "🎯 Build Targets:"
	@echo "  Static Library:   $(CPP_LIB)"
	@echo "  Shared Library:   $(SHARED_LIB)"
	@echo ""
	@echo "🖥️  Platform: $(UNAME_S)"
	@echo ""

# 验证构建环境
verify:
	@echo "Verifying hexlib build environment..."
	@echo ""
	@echo "Checking compiler..."
	@which $(CXX) > /dev/null || (echo "❌ Compiler not found: $(CXX)" && exit 1)
	@echo "✓ C++ Compiler: $$(which $(CXX))"
	@echo ""
	@echo "Checking source files..."
	@for src in $(CPP_SOURCES); do \
		if [ -f "$$src" ]; then \
			echo "✓ Found: $$src"; \
		else \
			echo "❌ Missing: $$src" && exit 1; \
		fi; \
	done
	@echo ""
	@echo "Checking include directory..."
	@if [ -d "$(INCLUDE_DIR)" ]; then \
		echo "✓ Include directory exists: $(INCLUDE_DIR)"; \
		ls -la $(INCLUDE_DIR)/*.h 2>/dev/null || echo "  (No header files found)"; \
	else \
		echo "❌ Include directory missing: $(INCLUDE_DIR)"; \
		exit 1; \
	fi
	@echo ""
	@echo "✓ Build environment verification complete!"

# 验证编译结果
verify-build:
	@echo "Verifying build artifacts..."
	@echo ""
	@if [ -f "$(CPP_LIB)" ]; then \
		echo "✓ Static library exists:"; \
		file $(CPP_LIB); \
		ls -lh $(CPP_LIB); \
		ar -t $(CPP_LIB) || true; \
	else \
		echo "❌ Static library not found: $(CPP_LIB)"; \
		exit 1; \
	fi
	@echo ""
	@echo "✓ Build artifacts verification complete!"

# 完整构建流程（从验证到编译）
full-build: verify prepare build verify-build
	@echo ""
	@echo "✅ Full build process completed successfully!"
