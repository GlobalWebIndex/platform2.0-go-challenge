VERSION = 0.0.1
REGISTRY ?= docker.io
REPO ?= nagabe
EXE_NAME = gwichallenge

.PHONY: build
build:
	go build -o $(EXE_NAME) *.go

.PHONE: run
run: build
	./gwichallenge

.PHONE: test
test:
	go test

.PHONY: builddocker
builddocker: build
	docker build -t $(REGISTRY)/$(REPO)/$(EXE_NAME):$(VERSION) .

.PHONY: pushdocker 
pushdocker:
	docker tag $(REGISTRY)/$(REPO)/$(EXE_NAME):$(VERSION) $(REGISTRY)/$(REPO)/$(EXE_NAME):latest
	docker push $(REGISTRY)/$(REPO)/$(EXE_NAME):$(VERSION)
	docker push $(REGISTRY)/$(REPO)/$(EXE_NAME):latest

.PHONY: clean	
clean:
	rm -f $(EXE_NAME)