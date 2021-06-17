// Package ordertocubby provides routines for mapping an order to a cubby
package ordertocubby

import (
	"hash/fnv"
	"strconv"
)

// Map an order by its ID to the correct cubby for that order by hashing it the specified number of times
func Map(orderID string, times, cubbiesCount uint32) string {
	result := hash(orderID)
	for times > 1 {
		result = hash(strconv.Itoa(int(result)))
		times--
	}

	return strconv.Itoa(int(result % cubbiesCount))
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
