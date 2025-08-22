OutDir?=/Users/zhao/Documents/01.Code/Hex/ems-plan/application/out/drivers

# 检查 TARGET 是否为空
ifeq ($(Project),)
    $(error Project is not set. Please specify Project, e.g., 'make Project=gpio_basic')
endif

RunDir?=$(Project)/bin
BuildBin?=1# 编译可执行文件
Main?=$(Project)/main.go
VERSION?=$(shell go run $(Main) version)
FileName?=$(Project)_$(VERSION)
BUILD_TIME=$(shell date "+%Y-%m-%d %H:%M:%S")
COMMIT_HASH=$(shell git rev-parse HEAD)


build:
	@echo "Building version $(VERSION)"
	@echo "Build time: $(BUILD_TIME)"
	@echo "Commit hash: $(COMMIT)"
	go build -buildmode=plugin \
		-ldflags "-s -w -X 'main.buildTime=\"$(BUILD_TIME)\"' -X 'main.commitHash=$(COMMIT_HASH)'" \
		-o ${RunDir}/${FileName}.driver \
		$(Main)
ifneq ($(strip $(RunDir)), $(strip $(OutDir)))
	cp ${RunDir}/${FileName}.driver ${OutDir}/
endif

	@echo "===> Build driver done"

ifeq ($(strip $(BuildBin)), 1)
	@echo "Building exe..."
	go build \
			-ldflags "-s -w -X 'main.buildTime=\"$(BUILD_TIME)\"' -X 'main.commitHash=$(COMMIT_HASH)'" \
			-o ${RunDir}/${FileName} \
			$(Main)
	@echo "Build exe done"
endif



clean:
	@echo "Cleaning..."
	rm -f ${OutDir}/$(Project)*.driver
	rm -f ${RunDir}/*
	@echo "Clean done"