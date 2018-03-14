package main

import ("fmt"
		"io/ioutil"
		"net/http"
		"encoding/json")

type Posts struct {
	Message string
	Created_time string
	Id string
}

type Navigation struct {
	Previous string
	Next string
}

type Wall []Posts

type Container struct {
	Data []Posts
	Paging Navigation
}

func request(url string) (*Container, error){
	resp, err := http.Get(url)

	if err != nil || (resp.Status != "200 OK") {
		fmt.Println("Error getting URL")
		fmt.Println("HTTP Status: ", resp.Status)

		return  nil,err
	}

	fmt.Println("HTTP Status: ", resp.Status)
	fmt.Println("ContentLength: ", resp.ContentLength)
	fmt.Println("Close: ", resp.Close)
	fmt.Println("TransferEncoding: ", resp.TransferEncoding)

	if resp.ContentLength < 0 {
		fmt.Println("No Data returned")
	}

	data, err := ioutil.ReadAll(resp.Body)

	if err == nil {
		fmt.Println(err)
	}

	var m Container
	err = json.Unmarshal(data, &m)

	if err != nil {
		fmt.Println("json.Unmarshal -> %v", err)
	}

	return &m, nil
}

func main() {
	var container *Container
	fmt.Println("starting...")
	var eaccessToken string = "EAACEdEose0cBAEVKkfn3GjCBBWVjBBvwuDAoN4uwWeuxqyE1bJJd9FSDkgpGJ4dDzpJVSg4kBkNH4XG64TpsetX6r5igQoiTzEiDCsAHztLX5xpSwd4TZAxffRlsKle2vvGEsMwOuRGVZBRNdThB2jQTQjKMwnf5ZAsN9aOMmmrnkFB1mldZCMQkQi4RqDWv5u11vGD4gAZDZD"
	var url string = "https://graph.facebook.com/v2.12/10153862771857156/posts?limit=25&until=1517841695&__paging_token=enc_AdBLO4Ynl6ZA7DY4O68GKcR2SGufSw4fJOK3AeWwfZAKZCOxKDhNmTnBshouGFiAZBcj8Kyelu24hvr8y7o5aUiLIIKh&access_token=" + eaccessToken
	container, _ = request(url)

	fmt.Println(container.Paging.Next)
}
