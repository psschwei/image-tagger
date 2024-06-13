package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type Catalog struct {
	Repositories []string
}

type ImageTags struct {
	Name string
	Tags []string
}

func main() {

	var c Catalog

	registry := os.Getenv("REGISTRY")
	token := os.Getenv("TOKEN")
	project := os.Getenv("PROJECT")

	req, err := http.NewRequest(http.MethodGet, registry+"/v2/_catalog", nil)
	if err != nil {
		fmt.Printf("Error creating images request: %s\n", err)
		os.Exit(1)
	}
	req.Header.Set("Authorization", "Bearer "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error getting images: %s\n", err)
		os.Exit(1)
	}

	err = json.NewDecoder(res.Body).Decode(&c)
	if err != nil {
		fmt.Println("Unable to decode response")
		os.Exit(1)
	}

	for _, image := range c.Repositories {
		if strings.HasPrefix(image, project) {
			fmt.Printf("image: %s\n", image)

			var tags ImageTags
			req, err := http.NewRequest(http.MethodGet, registry+"/v2/"+image+"/tags/list", nil)
			if err != nil {
				fmt.Printf("Error creating tags requeset: %s\n", err)
				os.Exit(1)
			}
			req.Header.Set("Authorization", "Bearer "+token)

			res, err := http.DefaultClient.Do(req)
			if err != nil {
				fmt.Printf("Error getting tags: %s\n", err)
				os.Exit(1)
			}

			err = json.NewDecoder(res.Body).Decode(&tags)
			for _, t := range tags.Tags {
				fmt.Printf("tag: %s\n", t)
			}
		}
	}
}
