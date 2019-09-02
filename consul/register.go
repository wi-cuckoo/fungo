package main

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

func main() {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}
	agent := client.Agent()
	ss, _ := agent.Services()
	for id, s := range ss {
		fmt.Println(id, s)
	}
	// agent.ServiceRegister(&api.AgentServiceRegistration{
	// 	Kind: api.ServiceKindTypical,
	// 	ID:   "0262c827-2e75-4d65-9929-71169a5c5556",
	// 	Name: "consistentlb",
	// 	Tags: []string{"ali", "qniu"},
	// 	Port: 40041,
	// })
}
