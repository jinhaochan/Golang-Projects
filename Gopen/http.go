package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
)

func init_client() *http.Client {

	jar, err := cookiejar.New(nil)

	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{Jar: jar}

	return client

}

func bodyReader(resp *http.Response) string {
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	return string(body)
}

func getCSRF(url string, client *http.Client) string {
	resp, err := client.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	body := bodyReader(resp)

	re := regexp.MustCompile("<input type='hidden' name='user_token' value='(.*)' />")

	match := re.FindStringSubmatch(body)

	return match[1]

}

func login(site string, path string, client *http.Client) {

	final := site + path

	token := getCSRF(final, client)

	// client is updated to store information about the PHPSESSID and CSRF token
	_, _ = client.PostForm(final, url.Values{
		"password":   {"password"},
		"username":   {"admin"},
		"Login":      {"Login"},
		"user_token": {token},
	})

}
