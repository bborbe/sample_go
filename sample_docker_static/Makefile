default: builddocker

setup:
	go get github.com/bborbe/sample_go/sample_webserver

buildgo:
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o main ./go/src/github.com/bborbe/sample_go/sample_webserver

builddocker:
	docker build -t bborbe/static-go-webserver-build -f ./Dockerfile.build .
	docker run -t bborbe/static-go-webserver-build /bin/true
	docker cp `docker ps -q -n=1 -f ancestor=bborbe/static-go-webserver-build`:/main .
	docker build --rm=true --tag=bborbe/static-go-webserver -f Dockerfile.static .
	rm main

run: builddocker
	docker run -p 8080:8080 bborbe/static-go-webserver
