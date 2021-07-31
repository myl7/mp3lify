package mp3lify

import "net/url"

func CleanUrl(s string) (string, error) {
	u, err := url.Parse(s)
	if err != nil {
		return "", err
	}

	if u.Scheme == "" {
		u.Scheme = "https"
	}
	s = u.String()

	u, err = url.ParseRequestURI(s)
	if err != nil {
		return "", err
	}

	return u.String(), nil
}
