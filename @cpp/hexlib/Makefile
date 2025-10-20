.PHONY: build clean help info shared prepare verify verify-build full-build

# 目录配置
SRC_DIR=src
INCLUDE_DIR=$(SRC_DIR)/include
BUILD_DIR=build
CPP_SOURCES=$(SRC_DIR)/hexlib.cpp $(SRC_DIR)/mpc_kalman.cpp

# 检查操作系统
ifeq ($(OS),Windows_NT)
	# ========== Windows 配置 ==========
	UNAME_S := Windows

	# Windows 编译器和工具
	CXX?=clang++
	CXXFLAGS=-std=c++11 -fPIC -Wall -O2
	CFLAGS=-I$(INCLUDE_DIR)

    # 库文件配置（使用 GNU ar 生成 libhexlib.a，便于与 CGO -lhexlib 匹配）
    CPP_OBJECTS=$(BUILD_DIR)/hexlib.o $(BUILD_DIR)/mpc_kalman.o
    CPP_LIB=$(BUILD_DIR)/libhexlib.a

    # Windows 特定工具
    MKDIR_CMD=mkdir
    RM_CMD=rmdir /s /q
    LS_CMD=dir
    AR=ar
    AR_OPTS=rcs

	# 目录就绪哨兵，避免与同名 phony 目标冲突
	BUILD_SENTINEL=$(BUILD_DIR)/.dir_ready

	# 共享库配置
	LDFLAGS=-shared
	SHARED_LIB=$(BUILD_DIR)/hexlib.dll

	# 禁用 Unix 特定的目标
	SKIP_VERIFY=1
else
	# ========== Unix/Linux/macOS 配置 ==========
	UNAME_S := $(shell uname -s)

	# Unix 编译器和工具
	CXX?=clang++
	CXXFLAGS=-std=c++11 -fPIC -Wall -O2
	CFLAGS=-I$(INCLUDE_DIR)

	# 库文件配置
	CPP_OBJECTS=$(BUILD_DIR)/hexlib.o $(BUILD_DIR)/mpc_kalman.o

	# Unix 特定工具
	MKDIR_CMD=mkdir -p
	RM_CMD=rm -rf
	LS_CMD=ls -lh
	AR=ar
	AR_OPTS=rcs

	# 平台特定配置
	ifeq ($(UNAME_S),Darwin)
		# macOS 特定配置
		CPP_LIB=$(BUILD_DIR)/libhexlib.a
		LDFLAGS=-dynamiclib
		SHARED_LIB=$(BUILD_DIR)/libhexlib.dylib
	else ifeq ($(UNAME_S),Linux)
		# Linux 特定配置
		CPP_LIB=$(BUILD_DIR)/libhexlib.a
		LDFLAGS=-shared
		SHARED_LIB=$(BUILD_DIR)/libhexlib.so
	else
		# 其他 Unix 系统（默认配置）
		CPP_LIB=$(BUILD_DIR)/libhexlib.a
		LDFLAGS=-shared
		SHARED_LIB=$(BUILD_DIR)/libhexlib.so
	endif

	# 启用 verify 目标
	SKIP_VERIFY=0
endif

# 创建build目录（作为依赖，只做一次）
ifeq ($(OS),Windows_NT)
$(BUILD_DIR):
	@if not exist $(BUILD_DIR) $(MKDIR_CMD) $(BUILD_DIR)
else
$(BUILD_DIR):
	@$(MKDIR_CMD) $@
endif

# Windows：通过哨兵文件确保目录存在（防止与同名目标冲突）
ifeq ($(OS),Windows_NT)
$(BUILD_SENTINEL):
	@if not exist $(BUILD_DIR) $(MKDIR_CMD) $(BUILD_DIR)
	@echo. > $(BUILD_SENTINEL)
endif

# 准备构建环境 - 自动创建必要的目录
prepare: | $(BUILD_DIR)
	@echo "✓ Build directory prepared: $(BUILD_DIR)/"

# 编译C++对象文件（按平台区分，避免 Windows 上重复创建目录导致错误）
ifeq ($(OS),Windows_NT)
# Windows: 通过哨兵文件确保 build 目录存在
$(BUILD_DIR)/%.o: $(SRC_DIR)/%.cpp $(INCLUDE_DIR)/hexlib.h | $(BUILD_SENTINEL)
	@echo "Compiling $<..."
	$(CXX) $(CXXFLAGS) $(CFLAGS) -c $< -o $@
	@echo "✓ Compiled: $@"
else
# Unix/macOS: 保持原有行为
$(BUILD_DIR)/%.o: $(SRC_DIR)/%.cpp $(INCLUDE_DIR)/hexlib.h
	@$(MKDIR_CMD) $(BUILD_DIR)
	@echo "Compiling $<..."
	$(CXX) $(CXXFLAGS) $(CFLAGS) -c $< -o $@
	@echo "✓ Compiled: $@"
endif

# 生成静态库
# 不需要显式依赖 $(BUILD_DIR)，因为 $(CPP_OBJECTS) 已经隐含了
$(CPP_LIB): $(CPP_OBJECTS)
	@echo "Creating static library..."
ifeq ($(OS),Windows_NT)
	@$(AR) $(AR_OPTS) $@ $(CPP_OBJECTS)
else
	@$(AR) $(AR_OPTS) $@ $^
endif
	@echo "✓ Static library built: $(CPP_LIB)"
	@echo ""
	@echo "=========================================="
	@echo "hexlib C++ library build complete"
	@echo "=========================================="
	@echo "Location: $(CPP_LIB)"
	@echo "Ready to use with CGO"
	@echo ""
ifeq ($(OS),Windows_NT)
	@echo Build artifacts: $(CPP_LIB)
else
	@$(LS_CMD) $(CPP_LIB)
endif

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
ifeq ($(OS),Windows_NT)
	@echo Build artifacts: $(SHARED_LIB)
else
	@$(LS_CMD) $(SHARED_LIB)
endif

# 默认目标（静态库）
all: $(CPP_LIB)

# 构建目标（静态库，等同于 all）
build: $(CPP_LIB)

# 构建动态库
shared: $(SHARED_LIB)

# 清理目标
clean:
	@echo "Cleaning hexlib C++ artifacts..."
ifeq ($(OS),Windows_NT)
	@if exist $(BUILD_DIR) $(RM_CMD) $(BUILD_DIR)
else
	@$(RM_CMD) $(BUILD_DIR)
endif
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

# 验证构建环境（仅在 Unix/Linux/macOS 上可用）
verify:
ifeq ($(OS),Windows_NT)
	@echo "Note: Verification targets are not available on Windows."
	@echo "Please ensure compiler and source files are present manually."
else
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
endif

# 验证编译结果（仅在 Unix/Linux/macOS 上可用）
verify-build:
ifeq ($(OS),Windows_NT)
	@echo "Note: Verification targets are not available on Windows."
	@echo "Check $(CPP_LIB) manually to verify build succeeded."
else
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
endif

# 完整构建流程（从验证到编译）
ifeq ($(OS),Windows_NT)
full-build: prepare build
	@echo ""
	@echo "✅ Full build process completed successfully!"
else
full-build: verify prepare build verify-build
	@echo ""
	@echo "✅ Full build process completed successfully!"
endif
