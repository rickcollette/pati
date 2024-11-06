# Makefile for building pati-interpreter and pati-linter

# Output directory
DIST_DIR = dist

# Go source files
INTERPRETER_MAIN = pati-interpreter/main.go
LINTER_MAIN = pati-linter/main.go

# Executable names
INTERPRETER_BIN = $(DIST_DIR)/pati
LINTER_BIN = $(DIST_DIR)/pati-linter

# Create the output directory if it doesn't exist
$(DIST_DIR):
	mkdir -p $(DIST_DIR)

# Build the pati-interpreter
pati-interpreter: $(DIST_DIR)
	go build -o $(INTERPRETER_BIN) $(INTERPRETER_MAIN)
	@echo "Built pati-interpreter -> $(INTERPRETER_BIN)"

# Build the pati-linter
pati-linter: $(DIST_DIR)
	go build -o $(LINTER_BIN) $(LINTER_MAIN)
	@echo "Built pati-linter -> $(LINTER_BIN)"

# Build both pati-interpreter and pati-linter
all: pati-interpreter pati-linter
	@echo "Built both pati-interpreter and pati-linter"

# Clean the dist directory
clean:
	rm -rf $(DIST_DIR)/*
	@echo "Cleaned up $(DIST_DIR)"
