
gocraft-up:
	@echo "# run this in another terminal"
	@echo cd ./snippets/gocraft/
	@echo go run . -db new.db &
	@echo cd ../../
	@echo "# then in this terminal"
	@echo make createblock
	@echo make deleteblock

createblock:
	curl localhost:9999/blocks/13x13x0 -d '{ "X": 13, "Y": 13, "Z": 0,"Kind": 1 }' -H "Content-Type: application/json"

deleteblock:
	curl localhost:9999/blocks/13x13x0 -X DELETE


