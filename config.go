package mp3lify

import (
	"errors"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func env(key string, val *string) string {
	v := os.Getenv(key)
	if v != "" {
		return v
	} else if val != nil {
		return *val
	} else {
		panic(errors.New(fmt.Sprintf("Env %s not found", key)))
	}
}

// func strPtr(s string) *string {
// 	return &s
// }

// func envInt(key string, val *int) int {
// 	v := os.Getenv(key)
// 	if v != "" {
// 		d, err := strconv.Atoi(v)
// 		if err != nil {
// 			panic(errors.New(fmt.Sprintf("Env %s is not int", key)))
// 		}
// 		return d
// 	} else if val != nil {
// 		return *val
// 	} else {
// 		panic(errors.New(fmt.Sprintf("Env %s not found", key)))
// 	}
// }

// func intPtr(d int) *int {
// 	return &d
// }

// authToken Token set at header X-Auth-Token to authenticate users
var authToken = env("AUTH_TOKEN", nil)

// cacheDir Directory to cache raw audio files
var cacheDir = env("CACHE_DIR", nil)
