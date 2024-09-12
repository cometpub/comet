package publications

import (
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

func FindPublicationAuthor(app core.App, publication *models.Record, username string) *models.Record {
	var author *models.Record = nil

	for _, record := range publication.ExpandedAll("authors") {
		if record.Username() == username {
			author = record
			break
		}
	}

	return author
}
