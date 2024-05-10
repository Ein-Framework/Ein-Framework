package entity

type Alert struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Scope       string `scope:"scope"`
}
