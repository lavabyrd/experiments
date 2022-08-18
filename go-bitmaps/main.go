package go_bitmaps

import "math/rand"

const restaurants = 65536
const bitmapLength = restaurants / 8

var (
	nearMetro      = make([]byte, bitmapLength)
	privateParking = make([]byte, bitmapLength)
	terrace        = make([]byte, bitmapLength)
	reservations   = make([]byte, bitmapLength)
	veganFriendly  = make([]byte, bitmapLength)
	expensive      = make([]byte, bitmapLength)
)

func fill(r *rand.Rand, b []byte, probability float32)

func indexes(a []uint64) []int

func andNoBoundsCheck() {

}
