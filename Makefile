
api	        :
	        go build

run-people	:
	        go build ./people/main -o people
	        docker run --rm eg_postgresql  --name pg-docker -e POSTGRES_PASSWORD=docker -dit -p 5432:5432  postgres 
	        ./go-crm

clean	        :
	        rm go-crm
	        && docker kill pg-docker
	        && rm *~

run             :
	        go run ./people/main
