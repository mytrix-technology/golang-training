package main

import (
	"database/sql"
	"github.com/dotse/go-health"
	//client "github.com/dotse/go-health"
	"os"
)

type MyApplication struct {
	db sql.DB
}

func (app *MyApplication) CheckHealth() []health.Check {
	c := health.Check{}
	if err := app.db.Ping(); err != nil {
		c.Status = health.StatusFail
		c.Output = err.Error()
	}

	//resp, err := client.CheckHealth(c.config)
	//if err == nil {
	//	fmt.Printf("Status: %s\n", resp.Status)
	//} else {
	//	fmt.Printf("ERROR: %v\n", err)
	//}

	return []health.Check{c}
}

func main() {
	if os.Args[1] == "healthcheck" {
		//client.CheckHealth()
		//client.CheckHealthCommand()
	}

	//app := NewMyApplication()
	//health.Register(true, "my-application", app)
	//server.Start()

	// Your other code...
}
