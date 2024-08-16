package publications

import (
	"fmt"
	"net/url"

	"github.com/pocketbase/pocketbase/models"
)

func RecordValueToImageSrc(record *models.Record, value string) string {
	src, err := url.JoinPath("/api", "files", record.Collection().Id, record.Id, value)

	if err != nil {
		return ""
	}

	return src
}

func RecordValueToImageSrcThumbnail(record *models.Record, value string, thumbnailSize string) string {
	src := RecordValueToImageSrc(record, value)

	if value == "" {
		return value
	}

	return fmt.Sprintf("%s?thumb=%s", src, thumbnailSize)
}

func RecordPropToImageSrc(record *models.Record, key string) string {
	value := record.GetString(key)

	if value == "" {
		values := record.GetStringSlice(key)

		if len(values) == 0 {
			return ""
		}

		value = values[0]
	}

	src, err := url.JoinPath("/api", "files", record.Collection().Id, record.Id, value)

	if err != nil {
		return ""
	}

	return src
}

func RecordPropToImageSrcThumbnail(record *models.Record, key string, thumbnailSize string) string {
	value := RecordPropToImageSrc(record, key)

	if value == "" {
		return value
	}

	return fmt.Sprintf("%s?thumb=%s", value, thumbnailSize)
}
