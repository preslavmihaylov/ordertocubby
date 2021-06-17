package ordertocubby

import (
	"hash/fnv"
	"strconv"
)

// Map an order by its ID to the correct cubby for that order
func Map(orderID string, cubbiesCount uint32) string {
	return strconv.Itoa(int(hash(orderID) % cubbiesCount))
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
