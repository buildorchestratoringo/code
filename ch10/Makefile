all:
	cat Makefile

start:
	CUBE_WORKER_HOST=localhost \
	CUBE_WORKER_PORT=5555 \
	CUBE_MANAGER_HOST=localhost \
	CUBE_MANAGER_PORT=5556 \
	go run main.go

list:
	@curl -s localhost:5556/tasks | jq

send/%:
	@curl -s --request POST \
	--header 'Content-Type: application/json' \
	--data @task$*.json \
	localhost:5556/tasks | jq

delete:
	@echo 'curl -s --request DELETE "localhost:5556/tasks/xxxxx"'

