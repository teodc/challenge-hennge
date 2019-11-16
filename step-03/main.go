package main

import (
	"bytes"
	"crypto/sha512"
	"encoding/base32"
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/xlzd/gotp"
)

const (
	URL = "https://hdechallenge-solve.appspot.com/challenge/003/endpoint"
	USERID = "user@example.com"
	GISTURL = "https://gist.github.com/teodc/abcd1234"
	SECRETSUFFIX = "HDECHALLENGE003"
	DIGITS = 10
	INTERVAL = 30
)

func main() {
	payload := bytes.NewBuffer([]byte(`{"contact_email": "` + USERID + `", "github_url": "` + GISTURL + `"}`))

	req, err := http.NewRequest("POST", URL, payload)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Content-Type", "application/json")

	req.SetBasicAuth(USERID, makeTotpToken(USERID+SECRETSUFFIX, DIGITS, INTERVAL))

	client := &http.Client{}

	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		panic(err)
	}

	reqDump, _ := httputil.DumpRequestOut(req, true)
	resDump, _ := httputil.DumpResponse(res, true)

	fmt.Printf("%s", reqDump)
	fmt.Printf("\n\n------------\n\n")
	fmt.Printf("%s", resDump)
}

func makeTotpToken(secret string, digits int, interval int) string {
	secret = base32.StdEncoding.EncodeToString([]byte(secret))

	hasher := gotp.Hasher{
		HashName: "sha512",
		Digest:   sha512.New,
	}

	return gotp.NewTOTP(secret, digits, interval, &hasher).Now()
}
