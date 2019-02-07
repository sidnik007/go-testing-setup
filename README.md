
In Go, you can run the unit test using `go test` command, and for acceptance testing (with the help of godog library) you can execute the feature files using `godog *.feature` But in case you want to bind and execute both unit and acceptance test
using just `go test`, then you can use a `TestMain` method and configure it as shown.