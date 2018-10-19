package geohash_go

import (
	"fmt"
	"github.com/zhaozy93/geohash_go"
)

func main() {
	hashKey, _ := geohash_go.EnGeoHash(float64(39.928167), float64(116.389550), 50)
	fmt.Println(hashKey)
	lat, lng, _ := geohash_go.DeGeoHash(hashKey)
	fmt.Println(lat, lng)
	neig, _ := geohash_go.GetNeighour(hashKey)
	fmt.Println(neig)
}
