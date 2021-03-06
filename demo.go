package geohash_go

import (
	"fmt"
	"testing"
	"time"
)

func geotest(t *testing.T) {
	tBeg := time.Now().UnixNano()
	hashKey, _ := EnGeoHash(float64(39.928167), float64(116.389550), 10)
	// hashKey, _ := EnGeoHash(float64(53.170996), float64(3.69526), 10)
	// hashKey, _ := EnGeoHash(float64(39.990234375), float64(116.5429687), 10)
	fmt.Println("EnGeoHash Time used:", (time.Now().UnixNano()-tBeg)/1000)
	fmt.Println(hashKey)

	tBeg = time.Now().UnixNano()
	lat, lng, _ := DeGeoHash(hashKey)
	fmt.Println("DeGeoHash Time used:", (time.Now().UnixNano()-tBeg)/1000)
	fmt.Println(lat, lng)

	tBeg = time.Now().UnixNano()
	neigh, _ := GetNeighbour_back(hashKey)
	fmt.Println("GetNeighour_back Time used:", (time.Now().UnixNano()-tBeg)/1000)
	fmt.Println(neigh)

	tBeg = time.Now().UnixNano()
	neigh, _ = GetNeighbour(hashKey)
	fmt.Println("GetNeighour Time used:", (time.Now().UnixNano()-tBeg)/1000)
	fmt.Println(neigh)
}
