package main

import (
	"bytes"
	"fmt"
	"net/http"
)

// set url, data, referer and cookie

func main() {

	for i := 0; i < 10; i++ {
		url := "https://docs.google.com/forms/d/e/1FAIpQLSelQnG8G6XrDdgR6C0LGKHwA8fbl8hyMGA9j4IDkNoDrz30Dg/formResponse"

		var data = []byte(`entry.2005620554=test&entry.1045781291=arshsharma461%40gmail.com&entry.1065046570=tet&fvv=1&draftResponse=%5Bnull%2Cnull%2C%221152171705221382118%22%5D%0D%0A&pageHistory=0&fbzx=1152171705221382118`)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
		req.Header.Set("Host", "docs.google.com")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:76.0) Gecko/20100101 Firefox/76.0")
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Origin", "https://docs.google.com")
		req.Header.Set("Referer", "https://docs.google.com/forms/d/e/1FAIpQLSelQnG8G6XrDdgR6C0LGKHwA8fbl8hyMGA9j4IDkNoDrz30Dg/viewform?fbzx=1152171705221382118")
		req.Header.Set("Cookie", "S=spreadsheet_forms=biRwGN0s_44HedTSeV0qKkkYNCjTODbnpIOLLfD12tE; GOOGLE_ABUSE_EXEMPTION=ID=9bd13419a97709a0:TM=1622033287:C=r:IP=223.226.250.36-:S=kdp4BL0qw81zMo8f5F_O6oo; NID=216=p-gByHGAtEeRzcqw44tF2hUEMExX_HJjvZpe26cGUK07tvTEZKB1dQOuVighijuYmRmCzIOpglPuYmpIP92icfJMdgTgeAwpQLwKt7gnV181WKK8hJp34yfCWEgX-SXgnFOdX9M2iMWy4OlKohwLD80GeTnjuzuTKXmcPFa5-kY")
		req.Header.Set("Upgrade-Insecure-Requests", "1")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		fmt.Println("response Status:", resp.Status)
	}
}
