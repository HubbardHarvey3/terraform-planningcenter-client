package core

// ************ Tabs ************

type SingleTabRoot struct {
	Data Tab `json:"data,omitempty"`
}

type TabRoot struct {
	Data []Tab `json:"data,omitempty"`
}

type Tab struct {
	Type       string         `json:"type"`
	ID         string         `json:"id"`
	Attributes TabsAttributes `json:"attributes"`
}

type TabsAttributes struct {
	Name     string `json:"name"`
	Sequence int    `json:"sequence"`
	Slug     string `json:"slug"`
}
