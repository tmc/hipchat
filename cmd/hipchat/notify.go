package main

import (
	"log"
	"os"

	"io/ioutil"

	"github.com/tmc/hipchat"
)

type Notify struct {
	Room    string `short:"r" long:"room" description:"Room name or ID" required:"true" env:"HIPCHAT_ROOM"`
	Message string `short:"m" long:"message" description:"Message to send" default:"-"`
	Color   string `short:"c" long:"color" description:"Color of message (yellow, green, red, purple, gray, random)" default:"yellow"`
}

var notifyOptions Notify

func init() {
	if _, err := optionsParser.AddCommand("notify", "send room notification", "", &notifyOptions); err != nil {
		log.Fatal(err)
	}
}

func (options *Notify) Execute(args []string) error {
	client := globalOptions.client()
	/*
		room, err := client.GetRoom(options.Room)
		if err != nil {
			return err
		}
	*/
	msg := options.Message
	if options.Message == "-" {
		if buf, err := ioutil.ReadAll(os.Stdin); err != nil {
			return err
		} else {
			msg = string(buf)
		}
	}
	//return room.SendNotification(msg, hipchat.Color(options.Color))
	return client.SendNotification(options.Room, msg, hipchat.Color(options.Color))
}
