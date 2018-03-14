package main

import ("fmt"
		"io/ioutil"
		"net/http"
		"encoding/json"
		"strings"
		"sort"
)

type Pair struct {
	Key   string
	Value int
}

// A slice of pairs that implements sort.Interface to sort by values
type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

type Posts struct {
	Message string
	Created_time string
	Id string
}

type Navigation struct {
	Previous string
	Next string
}

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

func WordCount(posts []Posts) map[string]int {
	m := make(map[string]int)

	for i:=0; i<len(posts); i++ {
		for _, f := range strings.Fields(posts[i].Message) {
			m[f] += 1
		}
	}

	return m
}

func main() {
	var container *Container
	var posts []Posts

	var eaccessToken string = "EAACEdEose0cBAJnsrzGhIo2h42ZCFhzg9xWkrSnI8Y1yNZCTWhLzXWq12regA538bvDfzmub1GJSS2lDDr8ZCLNrCcZAb7QTUV3mbVIfKSh3BQoWoM0p9fEyM5QEyxyLDHQ0GPMZB0xaIUY2ZCxjqf3pZCMvvp82vW8xfIBlZAqaesxGBJ7g09ojmZBVAQIRSMlguRH81uuNZBygZDZD"
	var url string = "https://graph.facebook.com/v2.12/10153862771857156/posts?__paging_token=enc_AdBLO4Ynl6ZA7DY4O68GKcR2SGufSw4fJOK3AeWwfZAKZCOxKDhNmTnBshouGFiAZBcj8Kyelu24hvr8y7o5aUiLIIKh&access_token=" + eaccessToken

	for {
		container, _ = request(url)
		posts = append(posts, container.Data...)
		if container.Paging.Next == "" {
			break;
		}

		url = container.Paging.Next
	}

	wordsCounter := WordCount(posts)

	p := make(PairList, len(wordsCounter))

	i := 0
	for k, v := range wordsCounter {
		p[i] = Pair{k, v}
		i++
	}

	sort.Sort(p)
	for _, k := range p {
		fmt.Println("%s: %s", k.Key, k.Value)
	}

}
