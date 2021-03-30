package main

func main() {

	mainSite := "http://192.168.71.129/"

	client := init_client()

	login(mainSite, "login.php", client)

	cmdInject(mainSite, "vulnerabilities/exec/", client)
}
