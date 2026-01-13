package main

import (
	"net/url"
	"time"

	"github.com/google/uuid"
	"golang.org/x/exp/slog"
)

type Agreement struct {
	ID       int
	finishAt *time.Time
	startAt  time.Time
}

func main() {
	gremio := ""

	tenantID, err := uuid.Parse(gremio)
	if err == nil {
		slog.Info("vamos!", "tenantID", tenantID)
		return
	}
	slog.Info("ai ai não!", "err", err)

	//		ID       int
	//		finishAt *time.Time
	//		startAt  time.Time
	//	}
	//
	// finishAt := time.Date(2025, 03, 01, 0, 0, 0, 0, time.UTC)
	//
	//	agreement1 := Agreement{
	//		ID:       1,
	//		startAt:  time.Date(2025, 01, 01, 0, 0, 0, 0, time.UTC),
	//		finishAt: &finishAt,
	//	}
	//
	//	agreement2 := Agreement{
	//		ID:       2,
	//		startAt:  time.Date(2025, 03, 01, 0, 0, 0, 0, time.UTC),
	//		finishAt: nil,
	//	}
	//
	//	accrualDates := []time.Time{
	//		time.Date(2025, 01, 01, 0, 0, 0, 0, time.UTC), // resultado esperado -> agreement_id = 1
	//		time.Date(2025, 02, 01, 0, 0, 0, 0, time.UTC), // resultado esperado -> agreement_id = 1
	//		time.Date(2025, 03, 01, 0, 0, 0, 0, time.UTC), // resultado esperado -> agreement_id = 1
	//		time.Date(2025, 04, 01, 0, 0, 0, 0, time.UTC), // resultado esperado -> agreement_id = 2
	//		time.Date(2025, 05, 01, 0, 0, 0, 0, time.UTC), // resultado esperado -> agreement_id = 2
	//		time.Date(2025, 06, 01, 0, 0, 0, 0, time.UTC), // resultado esperado -> agreement_id = 2
	//		time.Date(2025, 07, 01, 0, 0, 0, 0, time.UTC), // resultado esperado -> agreement_id = 2
	//		time.Date(2025, 8, 01, 0, 0, 0, 0, time.UTC),  // resultado esperado -> agreement_id = 2
	//		time.Date(2025, 9, 01, 0, 0, 0, 0, time.UTC),  // resultado esperado -> agreement_id = 2
	//		time.Date(2025, 10, 01, 0, 0, 0, 0, time.UTC), // resultado esperado -> agreement_id = 2
	//		time.Date(2025, 11, 01, 0, 0, 0, 0, time.UTC), // resultado esperado -> agreement_id = 2
	//		time.Date(2025, 12, 01, 0, 0, 0, 0, time.UTC), // resultado esperado -> agreement_id = 2
	//	}
}

type ParsedURL struct {
	protocol string
	username string
	password string
	hostname string
	port     string
	pathname string
	search   string
	hash     string
}

func newParsedURL(urlString string) ParsedURL {
	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		return ParsedURL{}
	}

	password, _ := parsedUrl.User.Password()

	return ParsedURL{
		protocol: parsedUrl.Scheme,
		username: parsedUrl.User.Username(),
		password: password,
		hostname: parsedUrl.Hostname(),
		port:     parsedUrl.Port(),
		pathname: parsedUrl.Path,
		search:   parsedUrl.RawQuery,
		hash:     parsedUrl.Fragment,
	}
}
