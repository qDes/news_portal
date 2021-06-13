.PHONY: db-up
db-up:
	docker-compose -f docker-compose.db.yml up -d --build

.PHONY: db-down
db-down:
	docker-compose -f docker-compose.db.yml down -v

.PHONY: kafka-up
kafka-up:
	docker-compose -f docker-compose.kafka.yml up -d --build

.PHONY: kafka-down
kafka-down:
	docker-compose -f docker-compose.kafka.yml down -v

.PHONY: infr-up
infr-up:
	docker-compose -f docker-compose.db.yml up -d --build && docker-compose -f docker-compose.kafka.yml up -d --build

.PHONY: infr-down
infr-down:
	docker-compose -f docker-compose.db.yml down -v && docker-compose -f docker-compose.kafka.yml down -v