docker run --name retreat -e POSTGRES_PASSWORD=mysecretpassword -e POSTGRES_DB=retreat -p 5432:5432 -v ${PWD}/initial.sql:/docker-entrypoint-initdb.d/initial.sql -d postgres:11.5

go test -coverprofile=c.out && go tool cover -html=c.out