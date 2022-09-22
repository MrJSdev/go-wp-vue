package entity

type Category struct {
	TermID    int    `json:"term_id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	TermGroup int    `json:"term_group"`
}
