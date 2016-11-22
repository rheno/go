package util

import (

"net/http"
"strings"
"io/ioutil"

)


/* Success Response */
type Success struct {
	Result string
}

/* Error Response */
type Error struct {
	Result string
}

/* Request Method */
func Request(success func(Success),
	     failed func(Error),
	     urls string,
	     methods string,
	     headers map[string]string,
	     payload string) {

	/* Set New Request */
	req, err :=  http.NewRequest(methods, urls, strings.NewReader(payload))


	/* Error Response */
	if err != nil {
		failed(Error{err.Error()})
		return
	}

	/* Set all headers */
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	/* Request to Server */
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		failed(Error{err.Error()})
		return
	}

	defer res.Body.Close()

	body, _ :=  ioutil.ReadAll(res.Body)

	/* Success Response */
	success(Success{string(body)})

}
