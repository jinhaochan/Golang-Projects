GoPen
---

A Go binary to exploit the Damn Vulnerable Web Application (https://dvwa.co.uk/)

Built with concurrency as a feature, it runs all the given exploits simultaneously against the different pages

Currently performs:

1. Command Injection
2. SQL Injection
3. Blind SQL Injection
4. Javascript Manipulation
5. CSRF
6. Local File Inclusion
7. DOM XSS

Cools things about GoLang learnt here:
1. Cookiejar to store http client session information
2. waitGroup for concurrency management 

