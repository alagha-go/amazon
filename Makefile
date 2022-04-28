restart:
	go build main.go && sudo systemctl restart amazon

reload:
	sudo systemctl restart amazon

build:
	go build main.go

add:
	git remote set-url origin https://[token]@github.com/alagha-go/amazon.git