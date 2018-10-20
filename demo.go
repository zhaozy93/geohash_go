package main

import (
	"fmt"
	"github.com/zhaozy93/geohash_go"
)

func main() {
	tBeg := time.Now().UnixNano()
	hashKey, _ := geohash_go.EnGeoHash(float64(39.928167), float64(116.389550), 10)
	// hashKey, _ := geohash_go.EnGeoHash(float64(53.170996), float64(3.69526), 10)
	// hashKey, _ := geohash_go.EnGeoHash(float64(39.990234375), float64(116.5429687), 10)
	fmt.Println("EnGeoHash Time used:", (time.Now().UnixNano()-tBeg)/1000)
	fmt.Println(hashKey)

	tBeg = time.Now().UnixNano()
	lat, lng, _ := geohash_go.DeGeoHash(hashKey)
	fmt.Println("DeGeoHash Time used:", (time.Now().UnixNano()-tBeg)/1000)
	fmt.Println(lat, lng)

	tBeg = time.Now().UnixNano()
	neigh, _ := geohash_go.GetNeighbour_back(hashKey)
	fmt.Println("GetNeighour_back Time used:", (time.Now().UnixNano()-tBeg)/1000)
	fmt.Println(neigh)

	tBeg = time.Now().UnixNano()
	neigh, _ = geohash_go.GetNeighbour(hashKey)
	fmt.Println("GetNeighour Time used:", (time.Now().UnixNano()-tBeg)/1000)
	fmt.Println(neigh)
}
