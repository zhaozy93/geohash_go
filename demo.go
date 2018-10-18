package geohash_go

import (
	"demo"
	"github.com/zhaozy93/geohash_go"
)

func main() {
	hashKey, _ := geohash_go.EnGeoHash(float64(39.928167), float64(116.389550), 50)
	fmt.Println(hashKey)
	lat, lng, _ := geohash_go.DeGeoHash(hashKey)
	fmt.Println(lat, lng)
}
