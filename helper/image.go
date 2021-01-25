package helper

import (
	"github.com/spf13/viper"
	"strings"
)

func GetAssetPath(path string) string {
	if path != "" {
		path = strings.TrimSuffix(viper.GetString("CDN_ENDPOINT"), "/") + "/" +
			strings.TrimSuffix(viper.GetString("MINIO_BUCKET"), "/") + "/" +
			strings.TrimPrefix(path, "/")
	}
	return path
}
