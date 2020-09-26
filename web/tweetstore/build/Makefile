



all: webserver

clean:
	rm -f webapp/app/webserver

webserver:
	CGO_ENABLED=0 go build -o webapp/app/webserver src/webserver.go 


