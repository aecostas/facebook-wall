This is just a pet project to learn GoLang. The goal
is to retrieve all your posts from a Facebook account. To do so I will use the standard Go paradigms and techniques.

The same result could be achieved easily by using directly the Facebook GraphQL Web, but I would not learn Go in that case :)

## Config and running
Go to https://developers.facebook.com/tools/explorer/ and get an access token (click all the checkboxes). Copy the token in the file `main.go`:
```
func main() {
	...
	var eaccessToken string = "<your token here>"
	...
}
```

Run the application:
```
go run main.go
```

## What do I learnt
This section lists the things I learnt during my way. The goal is not to explain them, but just to have them listed here for future reviews:
* Types! After years from my last code in C/C++, types again:
```
	type Posts struct {
		Message string
		Created_time string
		Id string
	}
```

* Function return values
```
func WordCount(posts []Posts) map[string]int
```

* Explicit variable declaration
```
var posts []Posts
```
* Implicit variable declaration
```
i := 0
```

* Wildcard for unused variables (\_)
```
container, _ = request(url)
```

* Object allocation
```
p := make(PairList, len(wordsCounter))
```
* Defer (TODO)
* Channels (TODO)


## Other resources
* Short youtube video series: https://www.youtube.com/watch?v=G3PvTWRIhZA
