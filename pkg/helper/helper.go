package helper

import "net/url"

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func StripUrlQuery(sourceUrl string) (string, error) {
	stripUrl, err := url.Parse(sourceUrl)
	if err != nil {
		return stripUrl.String(), err
	}

	query := stripUrl.Query()
	for key := range query {
		query.Del(key)
	}

	stripUrl.RawQuery = query.Encode()
	return stripUrl.String(), nil
}
