package timelines

import (
	"net/url"
	"strconv"

	"yatter-backend-go/app/domain/object"
)

func parseCondition(query url.Values) (*object.FindStatusCondition, error) {
	cond := new(object.FindStatusCondition)

	if s := query.Get("max_id"); s != "" {
		n, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, err
		}
		cond.MaxID = n
	}

	if s := query.Get("since_id"); s != "" {
		n, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, err
		}
		cond.SinceID = n
	}

	if s := query.Get("limit"); s != "" {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		cond.Limit = n
	}

	return cond, nil
}
