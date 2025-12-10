APP_NAME=coven
WIN_OUT_DIR=$(CURDIR)/_bin/win-x64
LINUX_OUT_DIR=$(CURDIR)/_bin/linux-x64
GOOS ?= windows
GOARCH ?= amd64
UI_DIR := $(CURDIR)/UI/coven
TEMPLATES_DIR := $(CURDIR)/UI/card_templates
DEFAULT_CONFIG := config.json

.PHONY:	win-builder_build \
		linux-builder_build \
		win-builder_win-x64 \
		linux-builder_win-x64

win-builder_build:
	set GOOS=$(GOOS)&& set GOARCH=$(GOARCH)&& \
	go build -o "$(WIN_OUT_DIR)/$(APP_NAME).exe" "$(CURDIR)/cmd/web/main.go"

linux-builder_build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) \
	go build -o "$(WIN_OUT_DIR)/$(APP_NAME).exe" "$(CURDIR)/cmd/web/main.go"

win-builder_win-x64:
	$(MAKE) win-builder_build GOOS=windows GOARCH=amd64
	xcopy "$(UI_DIR)" "$(WIN_OUT_DIR)/ui" /E /I /Y >nul
	xcopy "$(TEMPLATES_DIR)" "$(WIN_OUT_DIR)/ui/card_templates" /E /I /Y >nul
	copy "$(CURDIR)/config\$(DEFAULT_CONFIG)" "$(WIN_OUT_DIR)\"

linux-builder_win-x64:
	$(MAKE) linux-builder_build GOOS=windows GOARCH=amd64
	cp -r "$(UI_DIR)" "$(WIN_OUT_DIR)/ui/" && \
	cp -r "$(TEMPLATES_DIR)" "$(WIN_OUT_DIR)/ui/card_templates/" && \
	cp "$(CURDIR)/config/$(DEFAULT_CONFIG)" "$(WIN_OUT_DIR)/"