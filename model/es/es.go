package es

type SearchResponse struct {
	Hits struct {
		Hits []struct {
			Source struct {
				Id int64 `json:"id"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}
