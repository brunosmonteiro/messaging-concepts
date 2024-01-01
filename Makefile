# Target to spin up Docker Compose
.PHONY: channels/selective-consumers
channels/selective-consumers:
	@echo "Starting Docker Compose..."
	docker-compose -f ./resources/docker-compose.yml up -d base-activemq
	@echo "Building Go project..."
	go run ./api/channels/selective-consumers

.PHONY: channels/persistence
channels/persistence:
	@echo "Starting Docker Compose..."
	docker-compose -f ./resources/docker-compose.yml up -d persistent-activemq-kaha-db persistent-activemq-mysql mysql
	@echo "Building Go project..."
	go run ./api/channels/persistence

# Target to stop Docker Compose
.PHONY: docker-down
docker-down:
	@echo "Stopping Docker Compose..."
	docker-compose -f ./resources/docker-compose.yml down
