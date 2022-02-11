build:
	go build -ldflags="-s -w" -o bin/

release:
ifdef tag
		sh build.sh github.com/reaper47/183-day-rule $(tag)
		gh release create $(tag) ./release/$(tag)/*
else
		@echo 'Add the tag argument, i.e. `make release tag=v1.0.0`'
endif
	
%:
	@: 

.PHONY: build release
