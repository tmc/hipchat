package hipchat

import (
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
)

type Room struct {
	ID                int         `json:"id"`
	Name              string      `json:"name"`
	Created           string      `json:"created"`
	Topic             string      `json:"topic"`
	GuestAccessURL    interface{} `json:"guest_access_url"`
	IsArchived        bool        `json:"is_archived"`
	IsGuestAccessible bool        `json:"is_guest_accessible"`
	LastActive        string      `json:"last_active"`
	Links             struct {
		Self     string `json:"self"`
		Webhooks string `json:"webhooks"`
	} `json:"links"`
	Owner struct {
		ID    int `json:"id"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
		MentionName string `json:"mention_name"`
		Name        string `json:"name"`
	} `json:"owner"`
	Participants []struct {
		ID    int `json:"id"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
		MentionName string `json:"mention_name"`
		Name        string `json:"name"`
	} `json:"participants"`
	Privacy    string `json:"privacy"`
	Statistics struct {
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"statistics"`
	XmppJid string `json:"xmpp_jid"`

	client *Client
}

func (r Room) String() string {
	return r.Name
}

func (r *Room) HistoryFrom(from time.Time) (interface{}, error) {
	result := map[string]interface{}{}
	err := r.client.get(fmt.Sprintf("room/%d/history", r.ID), &result)
	if err != nil {
		return nil, err
	}
	spew.Dump(result)
	return result, nil
}

func (r *Room) History() (interface{}, error) {
	return r.HistoryFrom(time.Now().UTC())
}
