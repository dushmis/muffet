package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"
)

const (
	rootURL             = "http://localhost:8080"
	existentURL         = "http://localhost:8080/foo"
	nonExistentURL      = "http://localhost:8080/bar"
	erroneousURL        = "http://localhost:8080/erroneous"
	fragmentURL         = "http://localhost:8080/fragment"
	nonExistentIDURL    = "http://localhost:8080/#non-existent-id"
	baseURL             = "http://localhost:8080/base"
	invalidBaseURL      = "http://localhost:8080/invalid-base"
	redirectURL         = "http://localhost:8080/redirect"
	invalidRedirectURL  = "http://localhost:8080/invalid-redirect"
	missingMetadataURL  = "http://localhost:8081"
	invalidRobotsTxtURL = "http://localhost:8082"
	noResponseURL       = "http://localhost:8083"
)

type handler struct{}

func (handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "", "/":
		w.Write([]byte(htmlWithBody(`<a href="/foo" />`)))
	case "/foo":
		w.Write([]byte(htmlWithBody(`<a href="/" />`)))
	case "/erroneous":
		w.Write([]byte(htmlWithBody(`
			<a href=":" />
			<a href="mailto:me@right.here" />
			<a href="/bar" />
			<a href="#foo" />
		`)))
	case "/fragment":
		w.Write([]byte(htmlWithBody(`<a id="foo" href="#foo" />`)))
	case "/base":
		w.Write([]byte(`
			<html>
				<head>
					<base href="/parent/" />
				</head>
				<body>
					<a href="child" />
				</body>
			</html>
		`))
	case "/invalid-base":
		w.Write([]byte(`
			<html>
				<head>
					<base href=":" />
				</head>
				<body>
					<a href="child" />
				</body>
			</html>
		`))
	case "/parent/child":
	case "/redirect":
		w.Header().Add("Location", "/")
		w.WriteHeader(300)
	case "/invalid-redirect":
		w.WriteHeader(300)
	case "/robots.txt":
		u, err := url.Parse(erroneousURL)

		if err != nil {
			panic(err)
		}

		v, err := url.Parse(fragmentURL)

		if err != nil {
			panic(err)
		}

		w.Write([]byte(fmt.Sprintf(`
			User-agent: *
			Disallow: %v
			Disallow: %v
		`, u.Path, v.Path)))
	case "/sitemap.xml":
		w.Write([]byte(fmt.Sprintf(`
			<?xml version="1.0" encoding="UTF-8"?>
			<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
				<url>
					<loc>%v</loc>
					<lastmod>896-5-4</lastmod>
					<changefreq>monthly</changefreq>
					<priority>0.8</priority>
				</url>
				<url>
					<loc>%v</loc>
				</url>
			</urlset>
		`, rootURL, existentURL)))
	default:
		w.WriteHeader(404)
	}
}

type noMetadataHandler struct{}

func (noMetadataHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "", "/":
	default:
		w.WriteHeader(404)
	}
}

type invalidRobotsTxtHandler struct{}

func (invalidRobotsTxtHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "", "/":
	case "/robots.txt":
		w.Write([]byte(`
			Disallow: /
		`))
	default:
		w.WriteHeader(404)
	}
}

func TestMain(m *testing.M) {
	go http.ListenAndServe(":8080", handler{})
	go http.ListenAndServe(":8081", noMetadataHandler{})
	go http.ListenAndServe(":8082", invalidRobotsTxtHandler{})

	time.Sleep(time.Millisecond)

	os.Exit(m.Run())
}

func htmlWithBody(b string) string {
	return fmt.Sprintf(`<html><body>%v</body></html>`, b)
}
