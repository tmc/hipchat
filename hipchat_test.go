package hipchat

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func mkclient(t *testing.T, trace bool) *Client {
	authToken := ""
	if envToken := os.Getenv("HIPCHAT_AUTH_TOKEN"); envToken != "" {
		authToken = envToken
	}
	client, err := NewClient(authToken)
	if trace {
		client.TraceOn(log.New(os.Stderr, "[hipchat test] ", log.LstdFlags))
	}
	if err != nil {
		t.Fatal("error creating client:", err)
	}
	return client
}

func TestNewClient(t *testing.T) {
	mkclient(t, false)
}

func TestRoomFuncs(t *testing.T) {
	client := mkclient(t, false)
	rooms, err := client.Rooms()
	if err != nil {
		t.Fatal("error fetching rooms:", err)
	}

	for _, room := range rooms {
		if room.ID != 48436 {
			continue
		}
		fmt.Println("ROOM:", room, room.ID)
		r, err := client.GetRoom(room.Name)
		if err != nil {
			t.Fatal(err)
		}
		spew.Dump(r)
		h, err := r.History()
		if err != nil {
			t.Fatal(err)
		}
		spew.Dump(h)
	}
}

func TestUserFuncs(t *testing.T) {
	client := mkclient(t, true)
	users, err := client.Users()
	if err != nil {
		t.Fatal("error fetching users:", err)
	}
	for _, u := range users {
		if u.Name != "Travis Cline" {
			continue
		}
		log.Println(users)
		user, err := client.GetUser(fmt.Sprint(u.ID))
		if err != nil {
			t.Fatal(err)
		}
		spew.Dump(user)
	}
}
