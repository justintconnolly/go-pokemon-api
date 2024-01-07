package main

import (
	"time"

	"github.com/juju/ratelimit"
)

var bucket = ratelimit.NewBucket(time.Second, 20)
