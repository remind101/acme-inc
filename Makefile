.PHONY: build build-envoy

build:
	docker build --no-cache -t remind101/acme-inc .

build-envoy:
	docker build -t remind101/acme-inc:envoy .
