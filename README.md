GEOHASH in GO
=======================

An implementation of GEOHASH in GOLANG.
The desciption of GEOHASH can find at [here](http://geohash.org/) 


Getting the library
-------------------

` go get github.com/zhaozy93/geohash_go`

Usage
-------------------
``` golang
 import "github.com/zhaozy93/geohash_go"
 hashKey, err := geohash_go.EnGeoHash(float64(39.928167), float64(116.389550), 10)
 lat, lng, err := geohash_go.DeGeoHash(hashKey)
 neigh, err = geohash_go.GetNeighbour(hashKey)
)
```

Bugs
-------------------
If you have any questions or find some unexpectedly discovered, please comment or contribute.  

Refer
------------------
[GeoHash核心原理解析](https://www.cnblogs.com/LBSer/p/3310455.html)

[地理围栏算法解析（Geo-fencing）](https://www.cnblogs.com/LBSer/p/4471742.html) 

[离我最近之geohash算法](https://blog.csdn.net/sunrise_2013/article/details/42395261) 

[用打表的方式解决求Geohash当前区域周围8个区域编码](https://blog.csdn.net/dokd229933/article/details/49202981) 

[Geohash求当前区域周围8个区域编码的一种思路](https://blog.csdn.net/dokd229933/article/details/47021515)



-------------------
This code has been placed under the Apache License 2.0.
