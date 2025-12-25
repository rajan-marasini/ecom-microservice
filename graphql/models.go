package main

type Account struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Order []Order `json:"orders"`
}

func (p *PaginationInput) Bound() (uint64, uint64) {
	skip := uint64(0)
	take := uint64(0)
	if p.Skip != nil {
		skip = uint64(*p.Skip)
	}
	if p.Take != nil {
		take = uint64(*p.Take)
	}
	return skip, take
}
