package main

import (
	"fmt"
	gremlingo "github.com/apache/tinkerpop/gremlin-go/driver"
)

func main() {
	// Creating the connection to the server with default settings.
	driverRemoteConnection, err := gremlingo.NewDriverRemoteConnection("ws://localhost:8182/gremlin")
	// Handle error
	if err != nil {
		fmt.Println(err)
		return
	}
	// Cleanup
	defer driverRemoteConnection.Close()

	// Create an anonymous traversal source with remote
	g := gremlingo.Traversal_().WithRemote(driverRemoteConnection)

	// Add a vertex with properties to the graph with the terminal step Iterate()
	promise := g.AddV("gremlin").Property("language", "go").Iterate()

	// The returned promised is a go channel to wait for all submitted steps to finish execution and return error.
	err = <-promise
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get the value of the property
	result, err := g.V().HasLabel("gremlin").Values("language").ToList()
	if err != nil {
		fmt.Println(err)
		return
	}
	// Print the result
	for _, r := range result {
		fmt.Println(r.GetString())
	}
}
