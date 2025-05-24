package main

/*
Necessary data for registring service:
1. ID - Unique Identifier
2. Name
3. Tags []string
4. Address
5. Port
6. Metadata map[string]string
7. Health Check Config
*/
type Service struct {
	ID      string   `json:",omitempty"`
	Name    string   `json:",omitempty"`
	Tags    []string `json:",omitempty"`
	Port    int      `json:",omitempty"`
	Address string   `json:",omitempty"`
}
