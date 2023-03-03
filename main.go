package jsonRestConsumer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func PostRequest(uri string, body io.Reader, client http.Client) (string, error) {
	fmt.Print("Make post request on: ")
	fmt.Println(uri)
	response, err := client.Post(uri, "application/json", body)

	if err != nil {
		fmt.Println("failed to handle post request")
		return "", err
	}

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Failed to parse response")
		return "", err
	}

	result := string(content)

	fmt.Print("Printing response from post request")
	fmt.Println(result)

	return result, nil
}

func PatchRequest(uri string, body any) (string, error) {
	fmt.Print("Make post request on: ")
	fmt.Println(uri)

	request, _ := json.Marshal(body)
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPatch, uri, bytes.NewBuffer(request))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println("failed to handle post request")
		return "", err
	}

	response, err := client.Do(req)

	if err != nil {
		fmt.Println("failed to handle post request")
		return "", err
	}

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Failed to parse response")
		return "", err
	}

	result := string(content)
	return result, nil
}

func PutRequest(uri string, body any) (string, error) {
	fmt.Print("Make post request on: ")
	fmt.Println(uri)

	request, _ := json.Marshal(body)
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, uri, bytes.NewBuffer(request))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println("failed to handle post request")
		return "", err
	}

	response, err := client.Do(req)

	if err != nil {
		fmt.Println("failed to handle post request")
		return "", err
	}

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Failed to parse response")
		return "", err
	}

	result := string(content)
	return result, nil
}

func GetRequest(uri string) (string, error) {
	fmt.Print("Make post request on: ")
	fmt.Println(uri)
	response, err := http.Get(uri)

	if err != nil {
		fmt.Println("failed to handle post request")
		return "", err
	}

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Failed to parse response")
		return "", err
	}

	result := string(content)

	fmt.Print("Printing response from post request")
	fmt.Println(result)

	return result, nil
}
