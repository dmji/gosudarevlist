package parser

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

type Credentials struct {
	Login    string
	Password string
}

func HttpClientWithAuth(host string, cred Credentials) (*http.Client, error) {
	jarc, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Jar: jarc,
	}

	hostUrl, err := url.Parse(host)
	if err != nil {
		return nil, err
	}
	loginUrl, _ := hostUrl.Parse("/auth/login/")
	login := loginUrl.String()
	_, err = client.PostForm(
		login,
		url.Values{
			"login":    {cred.Login},
			"password": {cred.Password},
		},
	)

	if len(client.Jar.Cookies(hostUrl)) < 3 {
		return nil, fmt.Errorf("error on login to '%s'", hostUrl.Host)
	}

	if err != nil {
		return nil, err
	}

	return client, nil
}
