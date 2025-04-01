SHELL := /bin/bash -o pipefail

CLUB_NUMBER=6350
APP_NAME = tmi-status-checker
PACKAGE_NAME = $(APP_NAME)
SUBSCRIPTION_ID = eb792c5c-94c2-48d5-b355-c807ecdbe88e
WEB_SCRAPER=web_scraper
CAPTURE_SCREENSHOT=capture_screenshot
SEND_NOTIFICATION=send_notification

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
	sudo apt-get install -y nodejs libatk1.0-0 libc6 libcairo2 \
							libcups2 libdbus-1-3 libexpat1 libfontconfig1 libgbm1 libgcc1 \
							libgdk-pixbuf2.0-0 libglib2.0-0 libgtk-3-0 libnspr4 libpango-1.0-0 \
							libpangocairo-1.0-0 libstdc++6 libx11-6 libx11-xcb1 libxcb1 libxcomposite1 \
							libxcursor1 libxdamage1 libxext6 libxfixes3 libxi6 libxrandr2 libxrender1 \
							libxss1 libxtst6 ca-certificates fonts-liberation libnss3 lsb-release \
    							xdg-utils wget ca-certificates
	sudo update-ca-certificates
	npm install --save-dev puppeteer
	npm install -g prettier

generate-report:
	export CLUB_NUMBER=$(CLUB_NUMBER) && pushd $(WEB_SCRAPER) && go run main.go && popd

generate-screenshot:
	export CLUB_NUMBER=$(CLUB_NUMBER) && pushd $(CAPTURE_SCREENSHOT) && node capture.js && popd

generate-all: generate-report generate-screenshot

test: generate-all
	go run main.go

send-notification:
	export CLUB_NUMBER=$(CLUB_NUMBER) && pushd $(SEND_NOTIFICATION) && go run main.go && popd
