package hipchat

import (
	"fmt"
	"io"
)

// SendNotification sends a notification the the room
func (c *Client) SendNotification(room string, message string, color Color) error {
	result := map[string]interface{}{}
	options := map[string]interface{}{
		"message": message,
		"color":   color,
	}
	err := c.post(fmt.Sprintf("room/%v/notification", room), options, &result)
	if err == io.EOF {
		return nil
	}
	return err
}
