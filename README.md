# ordertocubby
A library made for the Golang @ Ocado course.

**Usage:**
Map the specified orderID to a cubbyID. `times` specifies the number of times to remap in case of collisions.
```go
cubbyID := ordertocubby.Map(orderID, times, cubbiesCount)
```
