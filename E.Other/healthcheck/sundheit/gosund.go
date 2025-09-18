package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"math/rand/v2"
	"net/http"
	"time"

	"github.com/AppsFlyer/go-sundheit"
	"github.com/AppsFlyer/go-sundheit/checks"
	healthhttp "github.com/AppsFlyer/go-sundheit/http"
)

func main() {
	// create a new health instance
	h := gosundheit.New()

	// define an HTTP dependency check
	httpCheckConf := checks.HTTPCheckConfig{
		CheckName: "httpbin.url.check",
		Timeout:   1 * time.Second,
		// dependency you're checking - use your own URL here...
		// this URL will fail 50% of the times
		URL: "http://httpbin.org/status/200,300",
	}
	// create the HTTP check for the dependency
	// fail fast when you misconfigured the URL. Don't ignore errors!!!
	httpCheck, err := checks.NewHTTPCheck(httpCheckConf)
	if err != nil {
		fmt.Println(err)
		return // your call...
	}

	// Alternatively panic when creating a check fails
	httpCheck = checks.Must(checks.NewHTTPCheck(httpCheckConf))

	err = h.RegisterCheck(
		httpCheck,
		gosundheit.InitialDelay(time.Second),       // the check will run once after 1 sec
		gosundheit.ExecutionPeriod(10*time.Second), // the check will be executed every 10 sec
	)

	if err != nil {
		fmt.Println("Failed to register check: ", err)
		return // or whatever
	}

	// define more checks...

	// register a health endpoint
	http.Handle("/admin/health.json", healthhttp.HandleHealthJSON(h))

	// serve HTTP
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Check is the API for defining health checks.
// A valid check has a non empty Name() and a check (Execute()) function.
type Check interface {
	// Name is the name of the check.
	// Check names must be metric compatible.
	Name() string
	// Execute runs a single time check, and returns an error when the check fails, and an optional details object.
	Execute() (details interface{}, err error)
}

//func lotteryCheck() (details interface{}, err error) {
//	lottery := rand.Float32()
//	details = fmt.Sprintf("lottery=%f", lottery)
//	if lottery < 0.5 {
//		err = errors.New("Sorry, I failed")
//	}
//	return
//}

type Lottery struct {
	myname      string
	probability float32
}

func (l Lottery) Execute() (details interface{}, err error) {
	lottery := rand.Float32()
	details = fmt.Sprintf("lottery=%f", lottery)
	if lottery < l.probability {
		err = errors.New("Sorry, I failed")
	}
	return
}

func (l Lottery) Name() string {
	return l.myname
}

//type checkEventsLogger struct{}
//
//func (l checkEventsLogger) OnCheckRegistered(name string, res gosundheit.Result) {
//	log.Printf("Check %q registered with initial result: %v\n", name, res)
//}
//
//func (l checkEventsLogger) OnCheckStarted(name string) {
//	log.Printf("Check %q started...\n", name)
//}
//
//func (l checkEventsLogger) OnCheckCompleted(name string, res gosundheit.Result) {
//	log.Printf("Check %q completed with result: %v\n", name, res)
//}

type healthLogger struct{}

//func (l healthLogger) OnResultsUpdated(results map[string]Result) {
//	log.Printf("There are %d results, general health is %t\n", len(results), allHealthy(results))
//}

// This listener can act both as check and health listener for reporting metrics
//oc := opencensus.NewMetricsListener()
//h := gosundheit.New(gosundheit.WithCheckListeners(oc), gosundheit.WithHealthListeners(oc))
// ...
//view.Register(opencensus.DefaultHealthViews...)
// or register individual views. For example:
//view.Register(opencensus.ViewCheckExecutionTime, opencensus.ViewCheckStatusByName, ...)
