APP_NAME=coven
WIN_OUT_DIR=$(CURDIR)/_bin/win-x64
UI_DIR := $(CURDIR)/UI/coven
TEMPLATES_DIR := $(CURDIR)/UI/card_templates
DEFAULT_CONFIG := config.json

.PHONY: coven-win-x64

coven-win-x64:
	set GOOS=windows&& set GOARCH=amd64&& \
	go build -o "$(WIN_OUT_DIR)/$(APP_NAME).exe" "$(CURDIR)/cmd/web/main.go" && \
	xcopy "$(UI_DIR)" "$(WIN_OUT_DIR)/ui" /E /I /Y >nul
	xcopy "$(TEMPLATES_DIR)" "$(WIN_OUT_DIR)/templs" /E /I /Y >nul
	copy "$(CURDIR)/config\$(DEFAULT_CONFIG)" "$(WIN_OUT_DIR)\" >nul 2>&1 || echo "Config not found"