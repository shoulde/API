package photo

type Photo struct {
	ID     int    `json:"ID,omitempty"`
	URL    string `json:"url,omitempty"`
	UserID int    `json:"userID,omitempty"`
}

func NewPhoto() *Photo {
	return &Photo{
		ID: 1,
	}
}

func (pc *PhotoCollection) AddPhoto(photo *Photo) {
	pc.Photos = append(pc.Photos, photo)
}

type PhotoCollection struct {
	counter int
	Photos  []*Photo
}

func NewPhotoCollection() *PhotoCollection {
	return &PhotoCollection{
		Photos: []*Photo{},
	}
}
