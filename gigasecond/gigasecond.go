package gigasecond

import "time"

func AddGigasecond(t time.Time) time.Time {
	gigasecond := time.Duration(1e9) * time.Second
	return t.Add(gigasecond)
}
