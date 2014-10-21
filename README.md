# go-geohash

A Go implementation of integer geohashing and also a Redis backed proximity search system

## Usage
### Encoder:

    hash := encoder.Encode(float64(1.0), float64(2.0))   # returns a uint64 geohash
    lat,lng := encoder.Decode(hash)   # returns float64 values representing latitude and longitude of original point
### Proximity:
Coming soon

## References
Code is based off 

* https://github.com/yinqiwen/geohash-int
* https://github.com/arjunmehta/node-geo-proximity