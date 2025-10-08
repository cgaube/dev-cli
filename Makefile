# Set the name for your binary
BINARY_NAME=dev

# The default command to run when you just type "make"
.DEFAULT_GOAL := build

# Build the application binary
build:
	@echo "Building ${BINARY_NAME}..."
	go build -o ${BINARY_NAME} .

# Build and run the application
run:
	./${BINARY_NAME}

# Remove the built binary
clean:
	@echo "Cleaning up..."
	rm -f ${BINARY_NAME}

# Phony targets are not files. This prevents conflicts with files of the same name.
.PHONY: build run clean
