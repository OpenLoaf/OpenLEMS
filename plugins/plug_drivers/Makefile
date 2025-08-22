# 输出目录配置
OutDir?=/Users/zhao/Documents/01.Code/Hex/ems-plan/application/out/drivers

# 自定义输出目录参数
ifneq ($(outdir),)
    OutDir := $(outdir)
endif

# 颜色定义
RED=\033[0;31m
GREEN=\033[0;32m
YELLOW=\033[1;33m
BLUE=\033[0;34m
PURPLE=\033[0;35m
CYAN=\033[0;36m
WHITE=\033[1;37m
NC=\033[0m # No Color

# 获取所有驱动项目列表
DRIVER_PROJECTS := $(shell find . -maxdepth 1 -type d -name "*" | grep -v "^\.$$" | grep -v "^\./bin$$" | sed 's|^\./||' | sort)

# 检查 project 参数（仅在构建时检查）
ifneq ($(MAKECMDGOALS),help)
ifneq ($(MAKECMDGOALS),list)
ifneq ($(MAKECMDGOALS),status)
ifneq ($(MAKECMDGOALS),clean-all)
ifneq ($(MAKECMDGOALS),build-all)
ifneq ($(MAKECMDGOALS),all)
ifneq ($(MAKECMDGOALS),clean)
    ifeq ($(project),)
        $(error project is not set. Please specify project, e.g., 'make project=gpio_basic' or 'make all' for all drivers)
    else
        PROJECTS_TO_BUILD := $(project)
    endif
endif
endif
endif
endif
endif
endif
endif

# 构建配置
BuildBin?=1
BUILD_TIME=$(shell date "+%Y-%m-%d %H:%M:%S")
COMMIT_HASH=$(shell git rev-parse HEAD 2>/dev/null || echo "unknown")

# 日志函数
define log_info
	@echo "$(GREEN)[INFO]$(NC) $(1)"
endef

define log_warn
	@echo "$(YELLOW)[WARN]$(NC) $(1)"
endef

define log_error
	@echo "$(RED)[ERROR]$(NC) $(1)"
endef

define log_success
	@echo "$(GREEN)[SUCCESS]$(NC) $(1)"
endef

define log_build_start
	@echo "$(BLUE)[BUILD]$(NC) 开始构建驱动: $(CYAN)$(1)$(NC)"
endef

define log_build_end
	@echo "$(GREEN)[BUILD]$(NC) 驱动构建完成: $(CYAN)$(1)$(NC)"
endef

# 默认目标
.PHONY: help
help:
	@echo "$(CYAN)=== EMS Plan 驱动构建工具 ===$(NC)"
	@echo ""
	@echo "$(YELLOW)使用方法:$(NC)"
	@echo "  make project=<驱动名称>     # 构建单个驱动"
	@echo "  make all                   # 构建所有驱动"
	@echo "  make clean project=<驱动名称> # 清理单个驱动"
	@echo "  make clean-all             # 清理所有驱动"
	@echo ""
	@echo "$(YELLOW)可选参数:$(NC)"
	@echo "  outdir=<路径>              # 指定输出目录"
	@echo ""
	@echo "$(YELLOW)可用驱动列表:$(NC)"
	@for project in $(DRIVER_PROJECTS); do \
		echo "  - $$project"; \
	done
	@echo ""
	@echo "$(YELLOW)示例:$(NC)"
	@echo "  make project=gpio_basic"
	@echo "  make all"
	@echo "  make clean project=gpio_basic"
	@echo "  make project=gpio_basic outdir=./custom_output"

# 构建目标
.PHONY: build build-single build-all all
build: build-single

# all目标 - 构建所有驱动
all: build-all

# 默认目标
.DEFAULT_GOAL := help

# 当指定project参数时，自动执行build目标
ifneq ($(project),)
    .DEFAULT_GOAL := build
endif

build-single:
	@if [ "$(project)" = "" ]; then \
		$(call log_error,"project 参数未设置"); \
		exit 1; \
	fi
	@if [ ! -d "$(project)" ]; then \
		echo "$(RED)[ERROR]$(NC) 驱动目录不存在: $(project)"; \
		exit 1; \
	fi
	@if [ ! -f "$(project)/main.go" ]; then \
		echo "$(RED)[ERROR]$(NC) main.go 文件不存在: $(project)/main.go"; \
		exit 1; \
	fi
	@echo "$(BLUE)[BUILD]$(NC) 开始构建驱动: $(CYAN)$(project)$(NC)"
	@echo "$(BLUE)构建信息:$(NC)"
	@echo "  项目: $(CYAN)$(project)$(NC)"
	@echo "  构建时间: $(CYAN)$(BUILD_TIME)$(NC)"
	@echo "  提交哈希: $(CYAN)$(COMMIT_HASH)$(NC)"
	@echo "  输出目录: $(CYAN)$(OutDir)$(NC)"
	@echo ""
	
	# 创建输出目录
	@mkdir -p $(project)/bin
	@mkdir -p $(OutDir)
	@echo "  创建输出目录: $(CYAN)$(OutDir)$(NC)"
	
	# 获取版本信息
	@echo "$(YELLOW)获取版本信息...$(NC)"
	@cd $(project) && \
	VERSION=$$(go run main.go version 2>/dev/null || echo "v1.0.0"); \
	FileName="$(project)_$$VERSION"; \
	echo "  版本: $(CYAN)$$VERSION$(NC)"; \
	echo "  文件名: $(CYAN)$$FileName$(NC)"; \
	echo ""
	
	# 构建插件
	@echo "$(YELLOW)构建插件文件...$(NC)"
	@cd $(project) && \
	VERSION=$$(go run main.go version 2>/dev/null || echo "v1.0.0"); \
	FileName="$(project)_$$VERSION"; \
	if go build -buildmode=plugin \
		-ldflags "-s -w -X 'main.buildTime=\"$(BUILD_TIME)\"' -X 'main.commitHash=$(COMMIT_HASH)'" \
		-o bin/$$FileName.driver \
		main.go 2>&1; then \
		echo "$(GREEN)[SUCCESS]$(NC) 插件构建成功"; \
		FILE_SIZE=$$(ls -lh bin/$$FileName.driver | awk '{print $$5}'); \
		echo "  文件: $(CYAN)$(OutDir)/$$FileName.driver$(NC) ($$FILE_SIZE)"; \
	else \
		ERROR_MSG=$$(cd $(project) && go build -buildmode=plugin \
			-ldflags "-s -w -X 'main.buildTime=\"$(BUILD_TIME)\"' -X 'main.commitHash=$(COMMIT_HASH)'" \
			-o bin/$$FileName.driver \
			main.go 2>&1 | head -5); \
		echo "$(RED)[ERROR]$(NC) 插件构建失败"; \
		echo "  原因: $(RED)$$ERROR_MSG$(NC)"; \
		exit 1; \
	fi
	
	# 复制到输出目录
	@echo "$(YELLOW)复制到输出目录...$(NC)"
	@cd $(project) && \
	VERSION=$$(go run main.go version 2>/dev/null || echo "v1.0.0"); \
	FileName="$(project)_$$VERSION"; \
	if echo "$(OutDir)" | grep -q "^/"; then \
		cp bin/$$FileName.driver $(OutDir)/; \
	else \
		cp bin/$$FileName.driver ../$(OutDir)/; \
	fi
	
	# 构建可执行文件（如果需要）
	@if [ "$(BuildBin)" = "1" ]; then \
		echo "$(YELLOW)构建可执行文件...$(NC)"; \
		cd $(project) && \
		VERSION=$$(go run main.go version 2>/dev/null || echo "v1.0.0"); \
		FileName="$(project)_$$VERSION"; \
		if go build \
			-ldflags "-s -w -X 'main.buildTime=\"$(BUILD_TIME)\"' -X 'main.commitHash=$(COMMIT_HASH)'" \
			-o bin/$$FileName \
			main.go 2>/dev/null; then \
			echo "$(GREEN)[SUCCESS]$(NC) 可执行文件构建成功"; \
			EXE_SIZE=$$(ls -lh bin/$$FileName | awk '{print $$5}'); \
			echo "  文件: $(CYAN)$(project)/bin/$$FileName$(NC) ($$EXE_SIZE)"; \
		else \
			EXE_ERROR=$$(cd $(project) && go build \
				-ldflags "-s -w -X 'main.buildTime=\"$(BUILD_TIME)\"' -X 'main.commitHash=$(COMMIT_HASH)'" \
				-o bin/$$FileName \
				main.go 2>&1 | head -3); \
			echo "$(YELLOW)[WARN]$(NC) 可执行文件构建失败（插件文件已成功构建）"; \
			echo "  原因: $(YELLOW)$$EXE_ERROR$(NC)"; \
		fi; \
	fi
	
	@echo "$(GREEN)[BUILD]$(NC) 驱动构建完成: $(CYAN)$(project)$(NC)"
	@echo "$(GREEN)=== 构建完成 ===$(NC)"

build-all:
	$(call log_info,"开始批量构建所有驱动...")
	@echo "$(BLUE)构建信息:$(NC)"
	@echo "  构建时间: $(CYAN)$(BUILD_TIME)$(NC)"
	@echo "  提交哈希: $(CYAN)$(COMMIT_HASH)$(NC)"
	@echo "  驱动数量: $(CYAN)$(words $(DRIVER_PROJECTS))$(NC)"
	@echo ""
	
	@success_count=0; \
	fail_count=0; \
	for project in $(DRIVER_PROJECTS); do \
		echo "$(BLUE)----------------------------------------$(NC)"; \
		BUILD_OUTPUT=$$(make project=$$project BuildBin=$(BuildBin) 2>&1); \
		if [ $$? -eq 0 ]; then \
			echo "$(GREEN)[SUCCESS]$(NC) $$project 构建成功"; \
			DRIVER_FILE=$$(find $(OutDir) -name "$$project*.driver" 2>/dev/null | head -1); \
			if [ -n "$$DRIVER_FILE" ]; then \
				FILE_SIZE=$$(ls -lh "$$DRIVER_FILE" | awk '{print $$5}'); \
				echo "  文件: $(CYAN)$$DRIVER_FILE$(NC) ($$FILE_SIZE)"; \
			fi; \
			success_count=$$((success_count + 1)); \
		else \
			echo "$(RED)[ERROR]$(NC) $$project 构建失败"; \
			ERROR_MSG=$$(echo "$$BUILD_OUTPUT" | grep -E "\[ERROR\]|原因:" | head -2); \
			if [ -n "$$ERROR_MSG" ]; then \
				echo "  原因: $(RED)$$ERROR_MSG$(NC)"; \
			fi; \
			fail_count=$$((fail_count + 1)); \
		fi; \
	done; \
	echo "$(BLUE)----------------------------------------$(NC)"; \
	echo "$(GREEN)批量构建完成!$(NC)"; \
	echo "  成功: $(GREEN)$$success_count$(NC)"; \
	echo "  失败: $(RED)$$fail_count$(NC)"; \
	echo "  总计: $(CYAN)$(words $(DRIVER_PROJECTS))$(NC)"

# 清理目标
.PHONY: clean clean-all
clean:
	@if [ "$(project)" = "" ]; then \
		echo "$(RED)[ERROR]$(NC) project 参数未设置"; \
		echo "$(YELLOW)使用方法:$(NC) make clean project=<驱动名称>"; \
		echo "$(YELLOW)或者使用:$(NC) make clean-all 清理所有驱动"; \
		exit 1; \
	fi
	@echo "$(GREEN)[INFO]$(NC) 清理驱动: $(project)"
	@rm -f $(OutDir)/$(project)*.driver
	@rm -f $(project)/bin/*
	@echo "$(GREEN)[SUCCESS]$(NC) 清理完成"

clean-all:
	$(call log_info,"清理所有驱动...")
	@for project in $(DRIVER_PROJECTS); do \
		echo "清理 $$project..."; \
		rm -f $(OutDir)/$$project*.driver; \
		rm -f $$project/bin/*; \
	done
	@$(call log_success,"所有驱动清理完成")

# 列出所有驱动
.PHONY: list
list:
	@echo "$(CYAN)=== 可用驱动列表 ===$(NC)"
	@for project in $(DRIVER_PROJECTS); do \
		if [ -f "$$project/main.go" ]; then \
			echo "$(GREEN)✓$(NC) $$project"; \
		else \
			echo "$(RED)✗$(NC) $$project (缺少main.go)"; \
		fi; \
	done
	@echo ""
	@echo "总计: $(words $(DRIVER_PROJECTS)) 个驱动"

# 检查驱动状态
.PHONY: status
status:
	@echo "$(CYAN)=== 驱动构建状态 ===$(NC)"
	@for project in $(DRIVER_PROJECTS); do \
		if [ -f "$$project/main.go" ]; then \
			driver_files=$$(find $(OutDir) -name "$$project*.driver" 2>/dev/null | wc -l); \
			if [ $$driver_files -gt 0 ]; then \
				echo "$(GREEN)✓$(NC) $$project (已构建)"; \
			else \
				echo "$(YELLOW)○$(NC) $$project (未构建)"; \
			fi; \
		else \
			echo "$(RED)✗$(NC) $$project (无效)"; \
		fi; \
	done