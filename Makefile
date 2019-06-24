.PHONY: all deps deps-down

deps:
	@docker-compose -f ./deps/docker-compose.yml --project-name=go-pgbouncer-example up -d

deps-down:
	@docker-compose -f ./deps/docker-compose.yml --project-name=go-pgbouncer-example down

run:
	@go run *.go
