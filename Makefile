# Target to spin up Docker Compose
.PHONY: channels/selective-consumers
channels/selective-consumers:
	@echo "Starting Docker Compose..."
	docker-compose -f ./resources/docker-compose.yml up -d
	@echo "Building Go project..."
	go run ./api/channels/selective-consumers

# Target to stop Docker Compose
.PHONY: docker-down
docker-down:
	@echo "Stopping Docker Compose..."
	docker-compose -f /../../resources/docker-compose.yml down
