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

func request(url string) (*Container, error) {
	var m Container

	resp, err := http.Get(url)

	if err != nil || (resp.Status != "200 OK") {
		fmt.Println("Error getting URL")
		fmt.Println("HTTP Status: ", resp.Status)
		return  nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = json.Unmarshal(data, &m)

	if err != nil {
		fmt.Println("json.Unmarshal -> %v", err)
		return nil, err
	}
	return &m, nil
}

func main() {
	var container *Container
	var posts []Posts

	var eaccessToken string = "EAACEdEose0cBAJamrYHMywKXwI3uxc8GJeZC6beW12lazKcSMF1ktpLjZAqeFT4lO3rRZBSuylTFC62vrBZAG5j4amH9QpcPv3OHACk99goSaJ4ls1ZAInkPyh0pHLYiXxCV3eSAEIZCNK6SGRIZAvWrcK0lt3Tywt1ZB5BMWzCKEdmwTDlcamNogAKbUYcdVTGSb540B9m3kgZDZD"
	var url string = "https://graph.facebook.com/v2.12/10153862771857156/posts?limit=25&until=1517841695&__paging_token=enc_AdBLO4Ynl6ZA7DY4O68GKcR2SGufSw4fJOK3AeWwfZAKZCOxKDhNmTnBshouGFiAZBcj8Kyelu24hvr8y7o5aUiLIIKh&access_token=" + eaccessToken

	for {
		container, _ = request(url)
		posts = append(posts, container.Data...)
		if container.Paging.Next == "" {
			break;
		}

		url = container.Paging.Next
	}

	fmt.Println(len(posts))
}
