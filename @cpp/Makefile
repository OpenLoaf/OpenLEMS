.PHONY: all build clean install generate

BUILD_DIR=build

all: build

build:
	cmake -S . -B $(BUILD_DIR) -DCMAKE_BUILD_TYPE=Release
	cmake --build $(BUILD_DIR) --config Release

install:
	cmake --install $(BUILD_DIR)

clean:
	rm -rf $(BUILD_DIR)

generate:
	bash bindings/gen.sh


