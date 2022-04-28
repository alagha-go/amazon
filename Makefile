restart:
	go build main.go && sudo systemctl restart amazon

reload:
	sudo systemctl restart amazon

build:
	go build main.go

add:
	git remote add origin https://alagha-go:[token]@github.com/alagha-go/amazon.git