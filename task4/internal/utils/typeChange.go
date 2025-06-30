package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetUintID(c *gin.Context, key string) (uint, error) {
	val, exists := c.Get(key)
	if !exists {
		return 0, fmt.Errorf("key %s not found", key)
	}

	switch v := val.(type) {
	case uint:
		return v, nil
	case int:
		if v < 0 {
			return 0, fmt.Errorf("negative value for uint")
		}
		return uint(v), nil
	case float64:
		if v < 0 || v > float64(^uint(0)) {
			return 0, fmt.Errorf("value out of uint range")
		}
		return uint(v), nil
	case string:
		id, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("invalid uint string")
		}
		return uint(id), nil
	default:
		return 0, fmt.Errorf("unsupported type %T", v)
	}
}
