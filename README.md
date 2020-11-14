# time-server

commands for start: 

mkdir testapi
cd testapi
git clone https://github.com/TarlexGit/time-server
cd time-server/
go run cmd/main.go

api commands (1-url, 2-params (key - value):

1 http://localhost:8000/time/now
2 

1 http://localhost:8000/time/string
2 time - 121521.172147

1 http://localhost:8000/time/add
2 time - 111420.124046
  delta - delta
  
1 http://localhost:8000/time/correct
2 time - 010101.010101
