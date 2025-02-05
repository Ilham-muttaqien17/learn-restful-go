GO = go
TARGET = main

build:
	$(GO) build -o $(TARGET) cmd/main.go

run: $(TARGET)
	./$(TARGET)

clean: 
	rm -f $(TARGET)

test:
	$(GO) test ./...

