package model

// Client represents a client of the organization.
type Client struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	// Client status. Enum: pending, confirmed, rejected.
	Status string `json:"status"`
	WebURL string `json:"webUrl"`
}

// ClientResponse defines the structure of a response when
// getting a client.
type ClientResponse struct {
	Data *Client `json:"data"`
}

// ClientsListResponse defines the structure of a response when
// getting a list of clients.
type ClientsListResponse struct {
	Data []*ClientResponse `json:"data"`
}
