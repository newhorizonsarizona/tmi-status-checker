SHELL := /bin/bash -o pipefail

APP_NAME = tmi-status-checker
PACKAGE_NAME = $(APP_NAME)
SUBSCRIPTION_ID = eb792c5c-94c2-48d5-b355-c807ecdbe88e
WEB_SCRAPER=web_scraper
CAPTURE_SCREENSHOT=capture_screenshot

.PHONY: test test-* format build

format:
	gofmt -w .
	pushd $(CAPTURE_SCREENSHOT) && prettier --write capture.js && popd

lint:
	gofmt -l .

install-tools:
	sudo add-apt-repository -y ppa:longsleep/golang-backports
	sudo apt update -y
	sudo apt install -y golang-go
	curl -fsSL https://deb.nodesource.com/setup_20.x | sudo bash -
	sudo apt-get install -y nodejs
	npm install --save-dev puppeteer
	npm install -g prettier

generate-report:
	pushd $(WEB_SCRAPER) && go run main.go && popd

generate-screenshot:
	pushd $(CAPTURE_SCREENSHOT) && node capture.js && popd

generate-all: generate-report generate-screenshot
	