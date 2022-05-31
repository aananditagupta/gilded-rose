# GO Starter

- Run :

```shell
go run texttest_fixture.go gilded-rose.go
```

- Run tests :

```shell
go test
```

- Run tests and coverage :

```shell
go test -coverprofile=coverage.out

go tool cover -html=coverage.out
```


Test Report: 
- Assumptions?

For the following project - there was one assumption made. This is with regards to the conjured items and their quality is affected. In the test it was mentioned that "Conjured items degrade in Quality twice as fast as normal items" - I have made the assumption that this happens both at the end of each day and also once the sell in date has crossed

There was however one ambuigity in the code - it was mentioned that "Once the sell by date has passed, Quality degrades twice as fast" but as seen in the code - the quality is still being decreased by 1. Keeping this in mind - I have deducted teh quality of conjured items by 2 (on both occasions -> day end and past the sell in date)

I have also broken down the increase and decrease in both sell in and quality into 4 different functions just to make sure that our code is optimised and refactored for future changes. 

Overall it was an interesting experience and I really enjoyed doing the test. 