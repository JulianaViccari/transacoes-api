unitTest:
	go test ./internal/core/domain/... --coverprofile=coverage.txt

coverage: unitTest
	go tool cover --html=coverage.txt