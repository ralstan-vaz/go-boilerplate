package rating

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/ralstan-vaz/go-boilerplate/config"
// )

// type MockRating struct {
// }

// func (r *MockRating) Get(req GetRequest) GetResponse {
// 	return GetResponse{ID: req.ID, Stars: "10"}
// }

// func NewMockRating() *MockRating {
// 	return &MockRating{}
// }

// func TestGetRating(t *testing.T) {
// 	// mockRating := NewMockRating()

// 	conf := &config.Config{}
// 	conf.User.RatingsUrl = "http://www.mocky.io/v2/5edbe434320000b5ad5d282f"
// 	mockRating := NewRating(conf)

// 	pkg := NewPkg(mockRating)
// 	req := GetRequest{ID: "Ashley"}
// 	res := pkg.Get(req)
// 	fmt.Println(res)
// }
