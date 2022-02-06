PROGRAM = ed25519-login
SOURCE = *.go

build:
	CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static -s"' -o $(PROGRAM) $(SOURCE)

clean:
	rm -f $(PROGRAM)

fmt:
	gofmt -w $(SOURCE)

vet:
	go vet $(SOURCE)
