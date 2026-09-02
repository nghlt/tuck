.PHONY: build install clean

BINARY := tuck

build:
	@echo "🔨 Building $(BINARY)..."
	@if go build -o $(BINARY) .; then \
		echo "✅ Build succeeded: ./$(BINARY)"; \
		./$(BINARY) version; \
	else \
		echo "❌ Build failed"; \
		exit 1; \
	fi

install:
	@echo "🔨 Installing $(BINARY) to $$(go env GOPATH)/bin..."
	@if go install .; then \
		echo "✅ Installed: $$(go env GOPATH)/bin/$(BINARY)"; \
		echo "   (make sure that directory is in your PATH)"; \
	else \
		echo "❌ Install failed"; \
		exit 1; \
	fi

clean:
	@rm -f $(BINARY)
	@echo "🧹 Removed ./$(BINARY)"
