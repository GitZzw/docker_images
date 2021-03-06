package main

import "github.com/hashicorp/consul/api"
import "fmt"

func main() {
	// Get a new client
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	// Get a handle to the KV API
	kv := client.KV()

	// PUT a new KV pair
	p := &api.KVPair{Key: "test_consul", Value: []byte("\n2021.03.06 test in Shanghai")}
	_, err = kv.Put(p, nil)
	if err != nil {
		panic(err)
	}

	// Lookup the pair
	pair, _, err := kv.Get("test_consul", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Key-Value: %v %s\n", pair.Key, pair.Value)
}
