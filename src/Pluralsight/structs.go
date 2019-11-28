// structs are custom data types

package main

import (
	"fmt"
)

func main() {
	// Custom type, can have methods associated with them
	type courseMeta struct {
		Author string
		Level string
		Rating float64
	}

	// Declare and initialize with zero values
	// var DockerDeepDive courseMeta
	// Equivalent to above but gives us a pointer
	// DockerDeepDive := new(courseMeta)

	// Composite Literal way
	DockerDeepDive := courseMeta{
		Author: "Nigel Poulton",
		Level: "Intermediate",
		Rating: 5,
	}

	fmt.Println("\n",DockerDeepDive)
	fmt.Println("\n",DockerDeepDive.Author)

	DockerDeepDive.Author = "Brian Gallagher"
	fmt.Println("\n",DockerDeepDive.Author)
}

