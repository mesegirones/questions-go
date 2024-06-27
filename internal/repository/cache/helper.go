package cache

import "fmt"

type CacheKey string

func (ck CacheKey) Format(args ...interface{}) string {
	return fmt.Sprintf(string(ck), args...)
}
