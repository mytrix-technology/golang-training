package main

import (
	"context"
	"fmt"
	health "github.com/AppsFlyer/go-sundheit"
	"github.com/AppsFlyer/go-sundheit/checks"
	"github.com/pkg/errors"
	"net"
	"time"
)

func NewDNSCheck(host string, timeout time.Duration, minRequiredResults int) checks.Check {
	resolver := net.DefaultResolver

	return &checks.CustomCheck{
		CheckName: "resolve." + host,
		CheckFunc: func() (details interface{}, err error) {
			ctx, cancel := context.WithTimeout(context.TODO(), timeout)
			defer cancel()

			addrs, err := resolver.LookupHost(ctx, host)
			resolvedCount := len(addrs)
			details = fmt.Sprintf("[%d] results were resolved", resolvedCount)
			if err != nil {
				return
			}
			if resolvedCount < minRequiredResults {
				err = errors.Errorf("[%s] lookup returned %d results, but requires at least %d", host, resolvedCount, minRequiredResults)
			}

			return
		},
	}
}

func registerHealthChecks() {
	// create a new health instance
	var h = health.New()
	// Schedule a host resolution check for `example.com`, requiring at least one results, and running every 10 sec
	h.RegisterCheck(&health.Config{
		Check:           checks.NewResolveCheck("example.com", 200*time.Millisecond, 1),
		ExecutionPeriod: 10 * time.Second,
	})
	health.

	// schedule more checks...
}
