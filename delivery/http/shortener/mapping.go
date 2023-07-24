package shortener

import (
	"url-shortener/common/helper"
	"url-shortener/delivery/http/shortener/model"
	domain "url-shortener/domain/shortener"
)

func mapRequestToLink(req *model.LinkRequest, l *domain.Link) {
	l.OriginalURL = req.URL
	l.ShortURL = helper.RandStringRunes()

	if req.ExpiredAt != nil {
		l.ExpireAt = req.ExpiredAt
	} else {
		threeDays := helper.ThreeDaysFromToday()
		l.ExpireAt = &threeDays
	}
}
