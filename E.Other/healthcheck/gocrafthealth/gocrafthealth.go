package main

import (
	"github.com/gocraft/health"
	"github.com/gocraft/health/sinks/bugsnag"
	"golang.org/x/net/context"
	"net/http"
	"os"
	"time"
)

// Save the stream as a global variable
var stream = health.NewStream()

// In your main func, initiailze the stream with your sinks.
func main() {
	// Log to stdout! (can also use WriterSink to write to a log file, Syslog, etc)
	stream.AddSink(&health.WriterSink{os.Stdout})

	// Log to StatsD!
	statsdSink, err := health.NewStatsDSink("127.0.0.1:8125", "myapp")
	if err != nil {
		stream.EventErr("new_statsd_sink", err)
		return
	}
	stream.AddSink(statsdSink)

	// Expose instrumentation in this app on a JSON endpoint that healthd can poll!
	sink := health.NewJsonPollingSink(time.Minute, time.Minute*5)
	stream.AddSink(sink)
	addr := ""
	sink.StartServer(addr)

	// Send errors to bugsnag!
	stream.AddSink(bugsnag.NewSink(&bugsnag.Config{APIKey: "myApiKey"}))

	// setup stream with sinks
	stream.AddSink(&health.WriterSink{os.Stdout})
	http.HandleFunc("/users", getUsers)

	// Events. Notice the camel case with dots.
	// (This is helpful when you want to use StatsD sinks)
	//job.Event("starting_server")
	//job.Event("proccess_user.by_email.gmail")
	//
	//// Event with keys and values:
	//job.EventKv("failover.started", health.Kvs{"from_ip": fmt.Sprint(currentIP)})

	// Timings:
	//startTime := time.Now()
	//// Do something...
	//job.Timing("fetch_user", time.Since(startTime).Nanoseconds()) // NOTE: Nanoseconds!
	//
	//// Timings also support keys/values:
	//job.TimingKv("fetch_user", time.Since(startTime).Nanoseconds(),
	//	health.Kvs{"user_email": userEmail})

	// Gauges:
	//job.Gauge("num_goroutines", numRunningGoroutines())
	//
	//// Timings also support keys/values:
	//job.GaugeKv("num_goroutines", numRunningGoroutines(),
	//	health.Kvs{"dispatcher": dispatcherStatus()})

	// Errors:
	//err := someFunc(user.Email)
	//if err != nil {
	//	return job.EventErr("some_func", err)
	//}
	//
	//// And with keys/Values:
	//job.EventErrKv("some_func", err, health.Kvs{"email": user.Email})

	// Make sink and add it to stream:
	sink = health.NewJsonPollingSink(time.Minute, time.Minute*5)
	stream.AddSink(sink)

	// Start the HTTP server! This will expose metrics via a JSON API.
	// NOTE: this won't interfere with your main app (if it also serves HTTP),
	// since it starts a separate net/http server.
	// In prod, addr should be a private network interface and port, like "10.2.1.4:5020"
	// In local dev, it can be something like "127.0.0.1:5020"
	sink.StartServer(addr)

	// Now that your stream is setup, start a web server or something...
}

func getUsers(rw http.ResponseWriter, r *http.Request) {
	// All logging and instrumentation should be within the context of a job!
	job := stream.NewJob("get_users")

	err := fetchUsersFromDatabase(r)
	if err != nil {
		// When in your job's context, you can log errors, events, timings, etc.
		job.EventErr("fetch_user_from_database", err)
	}

	// When done with the job, call job.Complete with a completion status.
	if err == nil {
		job.Complete(health.Success)
	} else {
		job.Complete(health.Error)
	}
}

func fetchUsersFromDatabase(r *http.Request) interface{} {
	return ""
}

type Context context.Context

//func showUser(ctx *Context) error {
//	user, err := ctx.getUser()
//	if err != nil {
//		// But we'll just log it here too!
//		return ctx.EventErr("show_user.get_user", err)
//	}
//}

//func getUser(ctx *Context) (*User, error) {
//	var u User
//	err := ctx.db.Select("SELECT * FROM users WHERE id = ?", ctx.userID).LoadStruct(&u)
//	if err != nil {
//		// Original error is here:
//		return nil, ctx.EventErr("get_user.select", err)
//	}
//	return &u, nil
//}

type Sink interface {
	EmitEvent(job string, event string, kvs map[string]string)
	EmitEventErr(job string, event string, err error, kvs map[string]string)
	EmitTiming(job string, event string, nanoseconds int64, kvs map[string]string)
	EmitGauge(job string, event string, value float64, kvs map[string]string)
	//EmitComplete(job string, status CompletionStatus, nanoseconds int64, kvs map[string]string)
}
