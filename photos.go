package go500px

import (
	"fmt"
	"net/url"
)

// PhotosService handles communication with the photos related
// methods of the 500px API.
//
// API docs: https://github.com/500px/api-documentation#photo-resources
type PhotosService struct {
	client *Client
}

// Photo is a single photo.
type Photo struct {
	Aperture          interface{} `json:"aperture"`
	Camera            interface{} `json:"camera"`
	Category          int         `json:"category"`
	CollectionsCount  int         `json:"collections_count"`
	CommentsCount     int         `json:"comments_count"`
	Converted         int         `json:"converted"`
	ConvertedBits     int         `json:"converted_bits"`
	CreatedAt         string      `json:"created_at"`
	CropVersion       int         `json:"crop_version"`
	Description       string      `json:"description"`
	FavoritesCount    int         `json:"favorites_count"`
	FocalLength       interface{} `json:"focal_length"`
	ForSale           bool        `json:"for_sale"`
	ForSaleDate       interface{} `json:"for_sale_date"`
	Height            int         `json:"height"`
	HiResUploaded     int         `json:"hi_res_uploaded"`
	HighestRating     float64     `json:"highest_rating"`
	HighestRatingDate string      `json:"highest_rating_date"`
	ID                int         `json:"id"`
	ImageFormat       string      `json:"image_format"`
	ImageURL          string      `json:"image_url"`
	Images            []struct {
		Format   string `json:"format"`
		HTTPSURL string `json:"https_url"`
		Size     int    `json:"size"`
		URL      string `json:"url"`
	} `json:"images"`
	Iso                interface{} `json:"iso"`
	Latitude           float64     `json:"latitude"`
	Lens               interface{} `json:"lens"`
	LicenseType        int         `json:"license_type"`
	LicensingRequested bool        `json:"licensing_requested"`
	Location           interface{} `json:"location"`
	Longitude          float64     `json:"longitude"`
	Name               string      `json:"name"`
	Nsfw               bool        `json:"nsfw"`
	PositiveVotesCount int         `json:"positive_votes_count"`
	Privacy            bool        `json:"privacy"`
	Profile            bool        `json:"profile"`
	Rating             float64     `json:"rating"`
	SalesCount         int         `json:"sales_count"`
	ShutterSpeed       interface{} `json:"shutter_speed"`
	Status             int         `json:"status"`
	TakenAt            interface{} `json:"taken_at"`
	TimesViewed        int         `json:"times_viewed"`
	URL                string      `json:"url"`
	User               struct {
		Affection int `json:"affection"`
		Avatars   struct {
			Default struct {
				HTTPS string `json:"https"`
			} `json:"default"`
			Large struct {
				HTTPS string `json:"https"`
			} `json:"large"`
			Small struct {
				HTTPS string `json:"https"`
			} `json:"small"`
			Tiny struct {
				HTTPS string `json:"https"`
			} `json:"tiny"`
		} `json:"avatars"`
		City            string `json:"city"`
		Country         string `json:"country"`
		CoverURL        string `json:"cover_url"`
		Firstname       string `json:"firstname"`
		Fullname        string `json:"fullname"`
		ID              int    `json:"id"`
		Lastname        string `json:"lastname"`
		StoreOn         bool   `json:"store_on"`
		UpgradeStatus   int    `json:"upgrade_status"`
		Username        string `json:"username"`
		UserpicHTTPSURL string `json:"userpic_https_url"`
		UserpicURL      string `json:"userpic_url"`
		Usertype        int    `json:"usertype"`
	} `json:"user"`
	UserID     int  `json:"user_id"`
	VotesCount int  `json:"votes_count"`
	Watermark  bool `json:"watermark"`
	Width      int  `json:"width"`
}

// PhotoStream represents a photo stream on 500px.
type PhotoStream struct {
	CurrentPage int    `json:"current_page"`
	Feature     string `json:"feature"`
	Filters     struct {
		Category interface{} `json:"category"`
		Exclude  interface{} `json:"exclude"`
	} `json:"filters"`
	Photos     []Photo `json:"photos"`
	TotalItems int     `json:"total_items"`
	TotalPages int     `json:"total_pages"`
}

// Returns a listing of twenty (up to one hundred) photos for a specified photo stream.
//
// API docs: https://github.com/500px/api-documentation/blob/master/endpoints/photo/GET_photos.md
func (s *PhotosService) GetStream(v url.Values) (*PhotoStream, error) {
	u := fmt.Sprintf("photos?%v", v.Encode())
	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, err
	}

	photos := new(PhotoStream)
	_, err = s.client.Do(req, photos)
	return photos, err
}
