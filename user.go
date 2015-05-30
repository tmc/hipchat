package hipchat

type User struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	MentionName string      `json:"mention_name"`
	Created     string      `json:"created"`
	AtlassianID interface{} `json:"atlassian_id"`
	Email       string      `json:"email"`
	Group       struct {
		ID    int `json:"id"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
		Name string `json:"name"`
	} `json:"group"`
	IsDeleted    bool   `json:"is_deleted"`
	IsGroupAdmin bool   `json:"is_group_admin"`
	IsGuest      bool   `json:"is_guest"`
	LastActive   string `json:"last_active"`
	Links        struct {
		Self string `json:"self"`
	} `json:"links"`
	PhotoURL string `json:"photo_url"`
	Presence struct {
		Client struct {
			Type    string `json:"type"`
			Version string `json:"version"`
		} `json:"client"`
		IsOnline bool `json:"is_online"`
	} `json:"presence"`
	Timezone string `json:"timezone"`
	Title    string `json:"title"`
	XmppJid  string `json:"xmpp_jid"`
}

type userResponse struct {
	*User
	*Error `json:"error"`
}

func (u User) String() string {
	return u.Name
}
