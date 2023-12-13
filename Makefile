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
	@echo -n "\n# /users\n=> "
	curl $(ADDR)/users

	@echo -n "\n# /users/1\n=> "
	curl $(ADDR)/users/1

	@echo -n "\n# /users/show?team=alpha&member=subzero\n=> "
	curl "$(ADDR)/users/show?team=alpha&member=subzero"
