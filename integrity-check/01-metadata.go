package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	npmBase     = "https://registry.npmjs.org"
	packageName = "lodash"
)

func main() {
	// Get lodash full metadata.
	fmd, err := PackageMetadata(npmBase, packageName)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(fmd))
	// Pretty print full metadata.
	fmd, err = json.MarshalIndent(string(fmd), "", "  ")
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(fmd))
	fmt.Println("-------------------")

	// Do the same for short metadata.
	// Get lodash full metadata.
	smd, err := ShortPackageMetadata(npmBase, packageName)
	if err != nil {
		panic(err)
	}
	// Pretty print full metadata.
	smd, err = json.MarshalIndent(smd, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(smd))
}

// PackageMetadata returns the full metadata of a package.
func PackageMetadata(registry, packageName string) ([]byte, error) {
	b := make([]byte, 0)
	if packageName == "" || registry == "" {
		return b, fmt.Errorf("empty package/registry name")
	}
	resp, err := http.Get(fmt.Sprintf("%s/%s", registry, packageName))
	if err != nil {
		return b, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// ShortPackageMetadata returns the abbreviated metadata of a package.
func ShortPackageMetadata(registry, packageName string) ([]byte, error) {
	b := make([]byte, 0)
	if packageName == "" || registry == "" {
		return b, fmt.Errorf("empty package/registry name")
	}
	// Create request.
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", registry, packageName), nil)
	if err != nil {
		return b, err
	}
	// Set "Accept: application/vnd.npm.install-v1+json" header.
	req.Header.Set("Accept", "application/vnd.npm.install-v1+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return b, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
