package geohash_go

import (
	"errors"
)

const (
	latMin = float64(-90.0)
	latMax = float64(90.0)
	lngMin = float64(-180.0)
	lngMax = float64(180.0)
)

var direction = map[string][]int{
	// dir: []int{上，下，左，右}
	"top":         []int{1, 0, 0, 0},
	"lefttop":     []int{1, 0, 1, 0},
	"righttop":    []int{1, 0, 0, 1},
	"left":        []int{0, 0, 1, 0},
	"right":       []int{0, 0, 0, 1},
	"bottom":      []int{0, 1, 0, 0},
	"bottomleft":  []int{0, 1, 1, 0},
	"bottomright": []int{0, 1, 0, 1},
}

var bitMap = map[string]string{
	"00000": "0",
	"00001": "1",
	"00010": "2",
	"00011": "3",
	"00100": "4",
	"00101": "5",
	"00110": "6",
	"00111": "7",
	"01000": "8",
	"01001": "9",
	"01010": "b",
	"01011": "c",
	"01100": "d",
	"01101": "e",
	"01110": "f",
	"01111": "g",
	"10000": "h",
	"10001": "j",
	"10010": "k",
	"10011": "m",
	"10100": "n",
	"10101": "p",
	"10110": "q",
	"10111": "r",
	"11000": "s",
	"11001": "t",
	"11010": "u",
	"11011": "v",
	"11100": "w",
	"11101": "x",
	"11110": "y",
	"11111": "z",
}

var strMap = map[string]string{
	"0": "00000",
	"1": "00001",
	"2": "00010",
	"3": "00011",
	"4": "00100",
	"5": "00101",
	"6": "00110",
	"7": "00111",
	"8": "01000",
	"9": "01001",
	"b": "01010",
	"c": "01011",
	"d": "01100",
	"e": "01101",
	"f": "01110",
	"g": "01111",
	"h": "10000",
	"j": "10001",
	"k": "10010",
	"m": "10011",
	"n": "10100",
	"p": "10101",
	"q": "10110",
	"r": "10111",
	"s": "11000",
	"t": "11001",
	"u": "11010",
	"v": "11011",
	"w": "11100",
	"x": "11101",
	"y": "11110",
	"z": "11111",
}

func EnGeoHash(lat, lng float64, accury int) (string, error) {
	if accury%5 != 0 {
		return "", errors.New("accury error")
	}
	latIndex := EnIndex(latMin, latMax, lat, accury)
	lngIndex := EnIndex(lngMin, lngMax, lng, accury)
	hashIndex := make([]byte, accury*2)
	for i := 0; i < accury; i++ {
		hashIndex[2*i] = lngIndex[i]
		hashIndex[2*i+1] = latIndex[i]
	}
	hashKey := ""
	end := accury * 2
	for {
		start := end - 5
		if start > 0 {
			hashKey = bitMap[string(hashIndex[start:end])] + hashKey
		} else {
			hashKey = bitMap[string(hashIndex[0:end])] + hashKey
			break
		}
		end = end - 5
	}
	return hashKey, nil
}

func EnIndex(min, max, num float64, accury int) string {
	index := ""
	for i := 0; i < accury; i++ {
		mid := (min + max) / 2
		if num < mid {
			index += "0"
			max = mid
		} else {
			index += "1"
			min = mid
		}
	}
	return index
}

func DeIndex(min, max float64, index []byte) float64 {
	mid := (min + max) / 2
	for i := 0; i < len(index); i++ {
		if index[i] == '0' {
			max = mid
		} else {
			min = mid
		}
		mid = (min + max) / 2
	}
	return mid
}

func DeGeoHash(hashKey string) (float64, float64, error) {
	hashStr := ""
	if len(hashKey)%2 != 0 {
		return float64(0), float64(0), errors.New("HashKey error")
	}
	for _, k := range hashKey {
		hashStr += strMap[string(k)]
	}
	latIndex := make([]byte, len(hashStr)/2)
	lngIndex := make([]byte, len(hashStr)/2)
	for i := 0; i < len(hashStr)/2; i++ {
		lngIndex[i] = hashStr[i*2]
		latIndex[i] = hashStr[i*2+1]
	}
	lat := DeIndex(latMin, latMax, latIndex)
	lng := DeIndex(lngMin, lngMax, lngIndex)
	return lat, lng, nil
}

func GetNeighour(hashKey string) (map[string]string, error) {
	hashStr := ""
	if len(hashKey)%2 != 0 {
		return nil, errors.New("HashKey error")
	}
	neighboour := make(map[string]string)
	for _, k := range hashKey {
		hashStr += strMap[string(k)]
	}
	latIndex := make([]byte, len(hashStr)/2)
	lngIndex := make([]byte, len(hashStr)/2)
	for i := 0; i < len(hashStr)/2; i++ {
		lngIndex[i] = hashStr[i*2]
		latIndex[i] = hashStr[i*2+1]
	}
	for k, _ := range direction {
		neighboour[k] = GetNeirghbourDir(latIndex, lngIndex, k)
	}
	return neighboour, nil
}

func GetNeirghbourDir(latIndex, lngIndex []byte, dir string) string {
	dirOper := direction[dir]
	if dirOper[0] == 1 {
		latIndex = CalculateBinary(latIndex, true)
	} else if dirOper[1] == 1 {
		latIndex = CalculateBinary(latIndex, false)
	}

	if dirOper[2] == 1 {
		lngIndex = CalculateBinary(lngIndex, false)
	} else if dirOper[3] == 1 {
		lngIndex = CalculateBinary(lngIndex, true)
	}
	hashIndex := make([]byte, len(lngIndex)*2)
	for i := 0; i < len(lngIndex); i++ {
		hashIndex[2*i] = lngIndex[i]
		hashIndex[2*i+1] = latIndex[i]
	}
	hashKey := ""
	end := len(hashIndex)
	for {
		start := end - 5
		if start > 0 {
			hashKey = bitMap[string(hashIndex[start:end])] + hashKey
		} else {
			hashKey = bitMap[string(hashIndex[0:end])] + hashKey
			break
		}
		end = end - 5
	}
	return hashKey
}

func CalculateBinary(index []byte, add bool) []byte {
	old := make([]byte, len(index))
	for i, v := range index {
		old[i] = v
	}
	last := old[len(old)-1]
	if last == '0' && add {
		old[len(old)-1] = '1'
		return old
	}
	if last == '1' && !add {
		old[len(old)-1] = '0'
		return old
	}
	if last == '1' && add {
		isEnd := true
		old[len(old)-1] = '0'
		for i := len(old) - 2; i >= 0; i-- {
			if old[i] == '1' {
				old[i] = '0'
			} else if old[i] == '0' {
				old[i] = '1'
				isEnd = false
				break
			}
		}
		if isEnd {
			for i, _ := range old {
				old[i] = '0'
			}
		}
		return old
	}

	if last == '0' && !add {
		isStart := true
		old[len(old)-1] = '1'
		for i := len(old) - 2; i >= 0; i-- {
			if old[i] == '0' {
				old[i] = '1'
			} else if old[i] == '1' {
				old[i] = '0'
				isStart = false
				break
			}
		}
		if isStart {
			for i, _ := range old {
				old[i] = '1'
			}
		}
		return old
	}
	return old
}
