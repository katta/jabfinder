package notifiers

import (
	"github.com/katta/jabfinder/pkg/models"
)

type Notifier interface {
	Notify(flatSessions []models.FlatSession)
}
