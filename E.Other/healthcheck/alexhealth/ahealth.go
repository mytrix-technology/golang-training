package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/alexliesenfeld/health"
	httpCheck "github.com/hellofresh/health-go/v4/checks/http"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc/balancer/grpclb/state"
	"log"
	"net/http"
	"time"
)

type CheckState state.State

type CheckerState state.State

// This is a very simple example that shows the basic features of this library.
func main() {
	db, _ := sql.Open("sqlite3", "simple.sqlite")
	defer db.Close()

	// Create a new Checker.
	checker := health.NewChecker(

		// Set the time-to-live for our cache to 1 second (default).
		health.WithCacheDuration(1*time.Second),

		// Configure a global timeout that will be applied to all checks.
		health.WithTimeout(10*time.Second),

		// A check configuration to see if our database connection is up.
		// The check function will be executed for each HTTP request.
		health.WithCheck(health.Check{
			Name:    "database",      // A unique check name.
			Timeout: 2 * time.Second, // A check specific timeout.
			Check:   db.PingContext,
		}),

		// The following check will be executed periodically every 15 seconds
		// started with an initial delay of 3 seconds. The check function will NOT
		// be executed for each HTTP request.
		health.WithPeriodicCheck(15*time.Second, 3*time.Second, health.Check{
			Name: "search",
			// The check function checks the health of a component. If an error is
			// returned, the component is considered unavailable (or "down").
			// The context contains a deadline according to the configured timeouts.
			Check: func(ctx context.Context) error {
				return fmt.Errorf("this makes the check fail")
			},
		}),

		// Set a status listener that will be invoked when the health status changes.
		// More powerful hooks are also available (see docs).
		health.WithStatusListener(func(ctx context.Context, state health.CheckerState) {
			log.Println(fmt.Sprintf("health status changed to %s", state.Status))
		}),
	)

	// Create a new health check http.Handler that returns the health status
	// serialized as a JSON string. You can pass pass further configuration
	// options to NewHandler to modify default configuration.
	http.Handle("/health", health.NewHandler(checker))
	log.Fatalln(http.ListenAndServe(":3000", nil))

	http.Handle("/health", health.NewHandler(
		health.NewChecker(
			health.WithCheck(health.Check{
				Name: "google",
				Check: httpCheck.New(httpCheck.Config{
					URL: "https://www.google.com",
				}),
			}),
		)))
	log.Fatalln(http.ListenAndServe(":3000", nil))

	//health.WithPeriodicCheck(5*time.Second, 0, health.Check{
	//	Name:   "search",
	//	Check:  myCheckFunc,
	//	StatusListener: func (ctx context.Context, name string, state CheckState) ) {
	//	log.Printf("status of component '%s' changed to %s", name, state.Status)
	//},
	//}),

	//health.WithStatusListener(func (ctx context.Context, state CheckerState)) {
	//	log.Printf("overall system health status changed to %s", state.Status)
	//}),

}

func myCheckFunc(ctx context.Context) error {
	return nil
}
