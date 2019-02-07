api	:
	go build

run	:
	go build
	docker run --rm eg_postgresql  --name pg-docker -e POSTGRES_PASSWORD=docker -dit -p 5432:5432  postgres 
	./go-crm

clean	:
	rm go-crm
	&& docker kill pg-docker
	&& rm *~
