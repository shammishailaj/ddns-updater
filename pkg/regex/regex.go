package regex

import (
	"regexp"
)

// Regex FindString functions
var (
	FindIP = regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`).FindString
)

// Regex MatchString functions
var (
	IP = regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`).MatchString
	Domain = regexp.MustCompile(`^(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9]))\.([a-zA-Z]{2,6}|[a-zA-Z0-9-]{2,30}\.[a-zA-Z]{2,3})$`).MatchString
	GodaddyKeySecret = regexp.MustCompile(`^[A-Za-z0-9]{12}\_[A-Za-z0-9]{22}\:[A-Za-z0-9]{22}$`).MatchString
	DuckDNSToken = regexp.MustCompile(`^[a-f0-9]{8}\-[a-f0-9]{4}\-[a-f0-9]{4}\-[a-f0-9]{4}\-[a-f0-9]{12}$`).MatchString
	NamecheapPassword = regexp.MustCompile(`^[a-f0-9]{32}$`).MatchString
	DreamhostKey = regexp.MustCompile(`^[A-Z0-9]{16}$`).MatchString
)