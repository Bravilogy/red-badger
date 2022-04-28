.DEFAULT_GOAL := docker-run

CURRENT_DIR = $(shell pwd)

docker-run:
	@if [[ "$(docker image inspect martians > /dev/null 2>&1)" != "" ]]; then \
	  echo "Building the container first"; \
	  docker build -t martians .; \
	fi
	@docker run \
		--mount type=bind,source=$(CURRENT_DIR)/input.txt,target=/tmp/input.txt \
		martians app --input /tmp/input.txt
