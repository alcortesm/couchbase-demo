package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/couchbase/gocb/v2"
)

func main() {
	fmt.Printf("gocb version: %s\n", gocb.Version())

	opts := gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: "Administrator",
			Password: "password",
		},
	}

	cluster, err := gocb.Connect("localhost", opts)
	if err != nil {
		log.Fatalf("connecting to cluster: %v", err)
	}

	defer func() {
		noOpts := (*gocb.ClusterCloseOptions)(nil)
		if err := cluster.Close(noOpts); err != nil {
			log.Fatalf("closing cluster connection: %v", err)
		}
	}()

	// wait util cluster is ready
	{
		timeout := 10 * time.Second
		// TODO: nil means desired state is online according to the
		// implementation
		opts := (*gocb.WaitUntilReadyOptions)(nil)
		if err := cluster.WaitUntilReady(timeout, opts); err != nil {
			log.Fatalf("waiting unitl cluster is ready: %v", err)
		}

		fmt.Println("cluster ready")
	}

	// ping server
	{
		opts := &gocb.PingOptions{
			Timeout: 1 * time.Second,
		}

		result, err := cluster.Ping(opts)
		if err != nil {
			log.Fatal("pinging: %v", err)
		}

		fmt.Printf("ping result: %s\n", jsonString(result))
	}

	// print diagnostics
	{
		noDiagOpts := (*gocb.DiagnosticsOptions)(nil)
		diag, err := cluster.Diagnostics(noDiagOpts)
		if err != nil {
			log.Fatalf("getting diagnostics: %v", err)
		}

		fmt.Printf("cluster diagnostics: %s\n", jsonString(diag))
	}

	const bucket = "travel-sample"

	// query 1: first 5 callsigns
	{
		query := "SELECT callsign from `travel-sample` limit 5 offset 0"
		opts := &gocb.QueryOptions{
			Timeout: 10 * time.Second,
		}

		result, err := cluster.Query(query, opts)
		if err != nil {
			log.Fatalf("query 1: %v", err)
		}

		if err := result.Close(); err != nil {
			log.Fatalf("closing result: %v", err)
		}

		// TODO: complains that the result is not closed
		metadata, err := result.MetaData()
		if err != nil {
			log.Fatalf("getting result metadata: %v", err)
		}

		fmt.Printf("query 1 metadata: %s\n", jsonString(metadata))
	}
}

func jsonString(v interface{}) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Fatalf("encoding json: %v", err)
	}

	return string(b)
}
