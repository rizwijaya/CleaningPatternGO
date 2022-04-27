.PHONY : format install build

#Website/Restfull API
run:
	go run main.go

initialize depedency Go:
	go mod init
	go mod tidy
	
install:
	go mod download

build:
	go build main.go

start:
	./main

# live reload using nodemon: npm -g i nodemon
run-nodemon:
	nodemon --exec go run main.go

#MQTT
run-mqtt:
	python /mqtt/service.py
	
initialize library MQTT:
	pip3 install -r /mqtt/requirements.txt