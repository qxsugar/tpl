dest ?= app

.PHONY: build
generate-api:
	rm -rf dest
	go-archetype transform --transformations=transformation-api.yml --source=src/api --destination=$(desc)