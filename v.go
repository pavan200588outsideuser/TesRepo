package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// fetchUserProfileImage simulates fetching a user profile image from a URL provided by the user.
// This function is vulnerable to SSRF.
func fetchUserProfileImage(imageURL string) ([]byte, error) {
	resp, err := http.Get(imageURL) // Direct use of user-provided URL
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-OK status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil

}

func functest() {
	// Example of how an attacker could exploit this:
	// An attacker provides an internal IP address or a service endpoint.
	internalServiceURL := "http://127.0.0.1:8080/admin" // Example of an internal service
	_, err := fetchUserProfileImage(internalServiceURL)
	if err != nil {
		fmt.Printf("Error fetching image: %v\n", err)
	} else {
		fmt.Println("Image fetched successfully (potentially from an internal service!)")
	}

	// Another example: targeting a different internal resource
	metadataServiceURL := "http://169.254.169.254/latest/meta-data/" // AWS EC2 metadata service
	_, err = fetchUserProfileImage(metadataServiceURL)
	if err != nil {
		fmt.Printf("Error fetching metadata: %v\n", err)
	} else {
		fmt.Println("Metadata fetched successfully (potentially from a cloud metadata service!)")
	}
}
