package main

import "sync"

var wg sync.WaitGroup

func main() {

	mainSite := "http://192.168.71.129/"

	client := init_client()

	login(mainSite, "login.php", client)

	wg.Add(7)

	go cmdInject(mainSite, "vulnerabilities/exec/", client)

	go SQLInject(mainSite, "vulnerabilities/sqli/", client)

	go blindSQLInject(mainSite, "vulnerabilities/sqli_blind/", client)

	go javascript(mainSite, "vulnerabilities/javascript/", client)

	go csrf(mainSite, "vulnerabilities/csrf/", client)

	go rfi(mainSite, "vulnerabilities/fi/", client)

	go xss_dom(mainSite, "vulnerabilities/xss_d/", client)

	wg.Wait()

}
