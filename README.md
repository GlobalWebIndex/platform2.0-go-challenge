# GlobalWebIndex Engineering Challenge 

Project is using Go modules instead of GoPath

```bash
go mod init example.com/gwi
```

`make test` to run tests, `make run` to build and run.

Uses https://github.com/gorilla/mux

Tests rest on endpoints_test.go
Types used rest on types.go
Main logic is implemented on main.go and endpoints.go

Assets have the same model except their special data. 
As I modeled the application , these data rest on a 
	-> property Data which is a stringified string on each asset object

Last but not least, currently the whole model is on memory - however if I were to choose a persist method I would prefer a NoSql technology due to the nature of the data.
