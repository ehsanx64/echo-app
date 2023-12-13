ADDR=localhost:1323

air:
	air

run:
	go run .

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
	@echo "\n# /users"
	curl $(ADDR)/users

	@echo "\n# /users/1"
	curl $(ADDR)/users/1

