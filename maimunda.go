package main

type Torrent struct {
	Link     string `json:"link"`
	Size     string `json:"size"`
	BG_AUDIO bool   `json:"bg_audio"`
	BG_SUBS  bool   `json:"bg_subs"`
}

type Movie struct {
	Title       string    `json:"title"`
	Genres      []string  `json:"genres"`
	Directors   []string  `json:"directors"`
	Actors      []string  `json:"actors"`
	Countries   []string  `json:"countries"`
	Year        int       `json:"year"`
	Description string    `json:"description"`
	PreviewLink string    `json:"preiewLink"`
	Rating      float64   `json:"rating"`
	Torrents    []Torrent `json:"torrets"`
}

type MoviesResponse struct {
	Value []Movie `json:"value"`
	Count int     `json:"count"`
}
