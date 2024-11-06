# Makefile for building pati-interpreter and pati-linter for multiple platforms

# Output directory
DIST_DIR = dist

# Go source files
INTERPRETER_MAIN = pati-interpreter/main.go
LINTER_MAIN = pati-linter/main.go

# Platforms and architectures
PLATFORMS = linux/amd64 darwin/amd64 windows/amd64 linux/arm64 darwin/arm64

# Create the output directory if it doesn't exist
$(DIST_DIR):
	mkdir -p $(DIST_DIR)

# Build for each platform and architecture
define BUILD_TARGET
build-$(1):
	GOOS=$(word 1, $(subst /, ,$(1))) GOARCH=$(word 2, $(subst /, ,$(1))) go build -o $(DIST_DIR)/pati-interpreter-$(word 1, $(subst /, ,$(1)))-$(word 2, $(subst /, ,$(1)))$(if $(findstring windows,$(word 1, $(subst /, ,$(1)))),.exe) $(INTERPRETER_MAIN)
	GOOS=$(word 1, $(subst /, ,$(1))) GOARCH=$(word 2, $(subst /, ,$(1))) go build -o $(DIST_DIR)/pati-linter-$(word 1, $(subst /, ,$(1)))-$(word 2, $(subst /, ,$(1)))$(if $(findstring windows,$(word 1, $(subst /, ,$(1)))),.exe) $(LINTER_MAIN)
	@echo "Built binaries for $(word 1, $(subst /, ,$(1)))-$(word 2, $(subst /, ,$(1)))"
endef

$(foreach platform,$(PLATFORMS),$(eval $(call BUILD_TARGET,$(platform))))

# Build all binaries
all: $(foreach platform,$(PLATFORMS),build-$(platform))
	@echo "Built all binaries for pati-interpreter and pati-linter"

# Clean the dist directory
clean:
	rm -rf $(DIST_DIR)/*
	@echo "Cleaned up $(DIST_DIR)"
