package backend

type ShowSample struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Bpm             int    `json:"bpm"`
	Instrument      string `json:"instrument"`
	Key             string `json:"key"`
	Length          string `json:"length"`
	Sample_file_url string `json:"sample_file_url"`
	Cover_url       string `json:"cover_url"`
	Price           string `json:"price"`
	Sample_type     string `json:"sample_type"`
	Quantity        int    `json:"quantity"`
}


func (h *handler)