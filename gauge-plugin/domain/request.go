package domain

import (
	"crypto/md5"
	"fmt"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

type Request struct {
	Method  string
	Url     Url
	Headers RequestHeaders
	Body    string
}

type RequestHeaders map[string]HeaderMatcher

type HeaderMatcher map[string]string

func (r *Request) toRequestStep(consumerName string) Step {
	var request string
	if len(r.Headers) > 0 {
		if r.IsBodyAvailable() {
			request = fmt.Sprintf(`URL"%s"にボディ<file:%s>、ヘッダー"%s"で、%sリクエストを送る`, r.Url, r.bodyFilePath(consumerName), r.Headers, r.Method)
		} else {
			request = fmt.Sprintf(`URL"%s"にヘッダー"%s"で、%sリクエストを送る`, r.Url, r.Headers, r.Method)
		}
	} else {
		if r.IsBodyAvailable() {
			request = fmt.Sprintf(`URL"%s"にボディ<file:%s>で、%sリクエストを送る`, r.Url, r.bodyFilePath(consumerName), r.Method)
		} else {
			request = fmt.Sprintf(`URL"%s"に%sリクエストを送る`, r.Url, r.Method)
		}
	}
	return Step(request)
}

func (r *Request) IsBodyAvailable() bool {
	return len(r.Body) > 0 && (r.Method == "POST" || r.Method == "PUT")
}

func (r *Request) bodyFilePath(consumerName string) string {
	return filepath.Join(RequestBodiesRelativePath(consumerName), r.ToBodyFileName())
}

func (r *Request) ToBodyFileName() string {
	re := regexp.MustCompile("[/|?|=|&]")
	replaced := re.ReplaceAllString(strings.TrimPrefix(r.Url.String(), "/"), "_")
	return fmt.Sprintf("%s_%s_%x.json", strings.ToLower(r.Method), replaced, r.toBodyHash())
}

func (r *Request) toBodyHeading() string {
	if r.IsBodyAvailable() {
		return fmt.Sprintf(" (body: %x)", r.toBodyHash())
	} else {
		return ""
	}
}

func (r *Request) toBodyHash() [16]byte {
	return md5.Sum([]byte(r.Body))
}

func (r *Request) toHeaderHeading() string {
	if len(r.Headers) > 0 {
		return fmt.Sprintf(" (%s)", r.Headers)
	} else {
		return ""
	}
}

func (h RequestHeaders) String() string {
	var headers []string
	for k, v := range h {
		headers = append(headers, fmt.Sprintf(`%s: %s`, k, v))
	}

	sort.Slice(headers, func(i, j int) bool { return headers[i] < headers[j] })

	return strings.Join(headers, ` \n `)
}

func (m HeaderMatcher) String() string {
	return m["equalTo"]
}
