SHELL=/usr/local/bin/fish

BUILD_DIR := build
BINARY_NAME := server
SERVER_FILE_PATH := server/main.go
SSH_KEY_PATH := $(MIKANMUSIC_SSH_KEY_PATH)
DST := $(MIKANMUSIC_DNS)
DST_DIR_PATH := ~
MIGRATE_SOURCE := $(MIKANMUSIC_MIGRATE_SRC)
MIGRATE_DB_LOCAL := $(MIKANMUSIC_DB_LOCAL)
MIGRATE_DB_REMOTE := $(MIKANMUSIC_DB_REMOTE)

deploy:
	mkdir $(BUILD_DIR)
	env GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME) $(SERVER_FILE_PATH)
	zip $(BUILD_DIR)/$(BINARY_NAME) $(BUILD_DIR)/$(BINARY_NAME)
	scp -i $(SSH_KEY_PATH) $(BUILD_DIR)/$(BINARY_NAME).zip ec2-user@$(DST):$(DST_DIR_PATH)
	rm -r $(BUILD_DIR)

# マイグレーションUP
migu:
	migrate -source $(MIGRATE_SOURCE) -database "$(MIGRATE_DB_LOCAL)" up

# マイグレーションDOWN
migd:
	migrate -source $(MIGRATE_SOURCE) -database "$(MIGRATE_DB_LOCAL)" down

# リモート: マイグレーションUP
migur:
	migrate -source $(MIGRATE_SOURCE) -database "$(MIGRATE_DB_REMOTE)" up

# リモート: マイグレーションDOWN
migdr:
	migrate -source $(MIGRATE_SOURCE) -database "$(MIGRATE_DB_REMOTE)" down 
