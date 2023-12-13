ADDR=localhost:1323
air:
	air
run:
	go run .
routes:
	@echo "\n# /"
	curl $(ADDR)

	@echo "\n# /welcome"
	curl $(ADDR)/welcome

	@echo "\n# /welcome/cj"
	curl $(ADDR)/welcome/cj

	@echo "\n# /welcome/adam"
	curl $(ADDR)/welcome/adam
