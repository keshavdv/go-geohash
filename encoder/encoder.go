package encoder

import (
	"errors"
)

type BoundingBox struct {
	MaxLatitude  float64
	MinLatitude  float64
	MaxLongitude float64
	MinLongitude float64
}

func EncodeWithPrecision(lat, lng float64, precision int) (point float64, err error) {
	bounds := BoundingBox{MaxLatitude: 90, MaxLongitude: 180, MinLatitude: -90, MinLongitude: -180}

	if lat < bounds.MinLatitude || lat > bounds.MaxLatitude || lng < bounds.MinLongitude || lng > bounds.MaxLongitude {
		err = errors.New("Coordinate out of bounds")
		return
	}

	var hash uint64
	var lat_bit, lng_bit uint64
	for i := 0; i < precision; i++ {
		if bounds.MaxLatitude-lat >= lat-bounds.MinLatitude {
			lat_bit = 0
			bounds.MaxLatitude = (bounds.MaxLatitude + bounds.MinLatitude) / 2
		} else {
			lat_bit = 1
			bounds.MinLatitude = (bounds.MaxLatitude + bounds.MinLatitude) / 2
		}

		if bounds.MaxLongitude-lng >= lng-bounds.MinLongitude {
			lng_bit = 0
			bounds.MaxLongitude = (bounds.MaxLongitude + bounds.MinLongitude) / 2
		} else {
			lng_bit = 1
			bounds.MinLongitude = (bounds.MaxLongitude + bounds.MinLongitude) / 2
		}
		hash <<= 1
		hash += lat_bit
		hash <<= 1
		hash += lng_bit
	}

	return float64(hash), nil
}

func Encode(lat, lng float64) (point float64, err error) {
	return EncodeWithPrecision(lat, lng, 26)
}

func DecodeWithPrecision(hash float64, precision int) (lat float64, lng float64) {
	intHash := uint64(hash)
	bounds := BoundingBox{MaxLatitude: 90, MaxLongitude: 180, MinLatitude: -90, MinLongitude: -180}
	var lat_bit, lng_bit uint64

	for i := 0; i < precision; i++ {
		lat_bit = GetBit(intHash, uint8((precision-i)*2-1))
		lng_bit = GetBit(intHash, uint8((precision-i)*2-2))

		if lat_bit == 0 {
			bounds.MaxLatitude = (bounds.MaxLatitude + bounds.MinLatitude) / 2
		} else {
			bounds.MinLatitude = (bounds.MaxLatitude + bounds.MinLatitude) / 2
		}

		if lng_bit == 0 {
			bounds.MaxLongitude = (bounds.MaxLongitude + bounds.MinLongitude) / 2
		} else {
			bounds.MinLongitude = (bounds.MaxLongitude + bounds.MinLongitude) / 2
		}
	}
	return ((bounds.MaxLatitude + bounds.MinLatitude) / 2), ((bounds.MinLongitude + bounds.MaxLongitude) / 2)
}

func Decode(point float64) (lat float64, lng float64) {
	return DecodeWithPrecision(point, 26)
}

func GetBit(bits uint64, pos uint8) uint64 {
	return (bits >> pos) & 0x01
}
