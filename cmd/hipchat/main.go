package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/tmc/hipchat"
)

type Options struct {
	Verbose   []bool `short:"v" long:"verbose" description:"Be verbose"`
	AuthToken string `short:"a" long:"authToken" description:"Hipchat auth token" env:"HIPCHAT_AUTH_TOKEN"`
}

func (o *Options) client() *hipchat.Client {
	client, err := hipchat.NewClient(o.AuthToken)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating hipchat client:", err)
		os.Exit(1)
	}
	if len(o.Verbose) > 0 {
		client.TraceOn(log.New(os.Stderr, "", log.LstdFlags))
	}
	return client
}

var globalOptions Options

var optionsParser = flags.NewNamedParser("hipchat", flags.Default)

func init() {
	optionsParser.AddGroup("Global options", "", &globalOptions)
}

func main() {
	if _, err := optionsParser.Parse(); err != nil {
		os.Exit(1)
	}
}
