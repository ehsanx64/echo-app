ADDR=localhost:1323

dev:
	air

deps:
	go get github.com/labstack/echo/v4
	go get github.com/jackc/pgx/v5
	
run:
	go run cmd/echoapp/main.go
	
tidy:
	go mod tidy -v

build-docker:
	docker build --tag echoapp:latest .

build-docker-multistage:
	docker build -t echoapp:multistage -f Dockerfile.multistage .

curl-general:
	@echo "\n# /"
	curl $(ADDR)

	@echo "\n# /welcome"
	curl $(ADDR)/welcome

	@echo "\n# /welcome/cj"
	curl $(ADDR)/welcome/cj

	@echo "\n# /welcome/adam"
	curl $(ADDR)/welcome/adam

curl-user:
	@echo -n "\n# /users\n=> "
	curl $(ADDR)/users

	@echo -n "\n# /users/1\n=> "
	curl $(ADDR)/users/1

	@echo -n "\n# /users/show?team=alpha&member=subzero\n=> "
	curl "$(ADDR)/users/show?team=alpha&member=subzero"

	@echo -n "\n# /users/save (name: 'Adam Smith', email: 'adam@earth.org')\n=> "
	curl -X POST -d "name=Adam Smith" -d "email=adam@earth.org" "$(ADDR)/users/save"

sqlc-clean:
	find . -regex .*repository.*go | xargs -n 1 rm

sqlc-gen:
	cd internal/user/repository && sqlc generate
	cd internal/blog/repository && sqlc generate
