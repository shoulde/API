package photo

type Photo struct {
	counter int
	URL     string `json:"url"`
	UserID  int    `json:"userID"`
}

func NewPhoto() *Photo {
	return &Photo{
		counter: 1,
	}
}

func (pc *PhotoCollection) AddPhoto(photo *Photo) {
	pc.Photos[pc.counter] = photo
	pc.counter = pc.counter + 1
}

type PhotoCollection struct {
	counter int
	Photos  map[int]*Photo
}

func NewPhotoCollection() *PhotoCollection {
	return &PhotoCollection{
		Photos: map[int]*Photo{},
	}
}
