package gigasecond

import "time"

// AddGigasecond add a gigasecond to the provided time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1e9)
}
