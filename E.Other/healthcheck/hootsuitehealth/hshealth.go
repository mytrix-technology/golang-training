package main

import (
	"database/sql"
	"fmt"
	"github.com/hootsuite/healthchecks"
	"github.com/hootsuite/healthchecks/checks/httpsc"
	"github.com/hootsuite/healthchecks/checks/sqlsc"
	"io/ioutil"
	"net/http"
	"strings"
)

type HttpStatusChecker struct {
	BaseUrl string
	Name    string
}

func (h HttpStatusChecker) Traverse(traversalPath []string, action string) (string, error) {
	dependencies := ""
	if len(traversalPath) > 0 {
		dependencies = fmt.Sprintf("&dependencies=%s", strings.Join(traversalPath, ","))
	}

	// Build our HTTP request
	url := fmt.Sprintf("%s/status/traverse?action=%s%s", h.BaseUrl, action, dependencies)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error creating request: %s \n", err.Error())
		return "", err
	}

	// Execute HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error executing request: %s \n", err.Error())
		return "", err
	}

	// Defer the closing of the body
	defer resp.Body.Close()

	// Read our response
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s", err.Error())
		return "", err
	}

	// Return our response
	return string(responseBody), nil
}

func main() {
	// Define a StatusEndpoint at '/status/db' for a database dependency
	myDB := sql.DB{}
	db := healthchecks.StatusEndpoint{
		Name:          "The DB",
		Slug:          "db",
		Type:          "internal",
		IsTraversable: false,
		StatusCheck: sqlsc.SQLDBStatusChecker{
			DB: myDB,
		},
		TraverseCheck: nil,
	}

	// Define a StatusEndpoint at '/status/service-organization' for the Organization service
	org := healthchecks.StatusEndpoint{
		Name:          "Organization Service",
		Slug:          "service-organization",
		Type:          "http",
		IsTraversable: true,
		StatusCheck: httpsc.HttpStatusChecker{
			BaseUrl: "[Read value from config]",
		},
		TraverseCheck: httpsc.HttpStatusChecker{
			BaseUrl: "[Read value from config]",
		},
	}

	// Define the list of StatusEndpoints for your service
	statusEndpoints := []healthchecks.StatusEndpoint{db, org}

	// Set the path for the about and version files
	aboutFilePath := "conf/about.json"
	versionFilePath := "conf/version.txt"

	// Set up any service injected customData for /status/about response.
	// Values can be any valid JSON conversion and will override values set in about.json.
	customData := make(map[string]interface{})
	// Examples:
	//
	// String value
	// customData["a-string"] = "some-value"
	//
	// Number value
	// customData["a-number"] = 123
	//
	// Boolean value
	// customData["a-bool"] = true
	//
	// Array
	// customData["an-array"] = []string{"val1", "val2"}
	//
	// Custom object
	// customObject := make(map[string]interface{})
	// customObject["key1"] = 1
	// customObject["key2"] = "some-value"
	// customData["an-object"] = customObject

	// Register all the "/status/..." requests to use our health checking framework
	http.Handle("/status/", healthchecks.Handler(statusEndpoints, aboutFilePath, versionFilePath, customData))
}

type RedisClient

func (c RedisClient) Ping() (interface{}, interface{}) {
	
}

type RedisStatusChecker struct {
	client RedisClient
}

func (r RedisStatusChecker) CheckStatus(name string) healthchecks.StatusList {
	pong, err := r.client.Ping()

	// Set a default response
	s := healthchecks.Status{
		Description:  name,
		Result: healthchecks.OK,
		Details: "",
	}

	// Handle any errors that Ping() function returned
	if err != nil {
		s = healthchecks.Status{
			Description:  name,
			Result: healthchecks.CRITICAL,
			Details: err.Error(),
		}
	}

	// Make sure the pong response is what we expected
	if pong != "PONG" {
		s = healthchecks.Status{
			Description:  name,
			Result: healthchecks.CRITICAL,
			Details: fmt.Sprintf("Expecting `PONG` response, got `%s`", pong),
		}
	}

	// Return our response
	return healthchecks.StatusList{ StatusList: []healthchecks.Status{ s }}
}
