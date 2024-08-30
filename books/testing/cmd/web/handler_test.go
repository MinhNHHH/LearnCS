package main

import (
	"context"
	"crypto/tls"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func Test_application_handler(t *testing.T) {
	var theTests = []struct {
		name                    string
		url                     string
		expectedStatusCode      int
		expectedURL             string
		expectedFirstStatusCode int
	}{
		{"home", "/", http.StatusOK, "/", http.StatusOK},
		{"404", "/404", http.StatusNotFound, "/fish", http.StatusNotFound},
		{"profile", "/user/profile", http.StatusOK, "/", http.StatusTemporaryRedirect},
	}

	routes := app.routes()

	// create a test server
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := &http.Client{
		Transport: tr,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	// range through test data
	for _, e := range theTests {
		resp, err := ts.Client().Get(ts.URL + e.url)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}
		if resp.StatusCode != e.expectedStatusCode {
			t.Errorf("for %s: expectedStatusCode %d, but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
		}

		if resp.Request.URL.Path != e.expectedURL {
			t.Errorf("%s: expected final url of %s but got %s", e.name, e.expectedURL, resp.Request.URL.Path)
		}

		resp2, _ := client.Get(ts.URL + e.url)
		if resp2.StatusCode != e.expectedFirstStatusCode {
			t.Errorf("%s: expected first returned status code to be %d but got %d", e.name, e.expectedFirstStatusCode, resp2.StatusCode)
		}
	}
}

// func TestApp_renderWithBadTemplate(t *testing.T) {
// 	// set templatepath to a location with a bad template
// 	pathToTemplates = "/home/minh/Desktop/learn_cs/books/testing/cmd/testdata/"
// 	req, _ := http.NewRequest("GET", "/", nil)
// 	req = addContextAndSessionToRequest(req, app)
// 	rr := httptest.NewRecorder()
//
// 	err := app.render(rr, req, "bad.page.gohtml", &TemplateData{})
// 	if err == nil {
// 		t.Error("expected error from bad template, but did not get")
// 	}
// pathToTemplates = "/home/minh/Desktop/learn_cs/books/testing/cmd/template/"

// }

func getCtx(req *http.Request) context.Context {
	ctx := context.WithValue(req.Context(), contextUserKey, "unknown")
	return ctx
}

func addContextAndSessionToRequest(req *http.Request, app application) *http.Request {
	req = req.WithContext(getCtx(req))

	ctx, _ := app.Session.Load(req.Context(), req.Header.Get("X-Session"))
	return req.WithContext(ctx)
}

func TestAppHone(t *testing.T) {
	var tests = []struct {
		name         string
		putInSession string
		expectedHTML string
	}{
		{"first visit", "", "<small>From Session:"},
		{"second visit", "hello word", "<small>From Session: hello word"},
	}

	for _, e := range tests {
		req, _ := http.NewRequest("GET", "/", nil)
		req = addContextAndSessionToRequest(req, app)

		if e.putInSession != "" {
			app.Session.Put(req.Context(), "test", e.putInSession)
		}

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(app.Home)

		handler.ServeHTTP(rr, req)

		// check status code
		if rr.Code != http.StatusOK {
			t.Errorf("TestAppHone returned wrong status code; expected 200 but got %d", rr.Code)
		}

		body, _ := io.ReadAll(rr.Body)
		if !strings.Contains(string(body), e.expectedHTML) {
			t.Errorf("%s: did not find correct %s in body", e.name, e.expectedHTML)
		}

	}
}

func Test_app_Login(t *testing.T) {
	var tests = []struct {
		name               string
		postedData         url.Values
		expectedStatusCode int
		expectedLoc        string
	}{
		{
			name: "valid login",
			postedData: url.Values{
				"email":    {"admin@example.com"},
				"password": {"secret"},
			},
			expectedStatusCode: http.StatusSeeOther,
			expectedLoc:        "/user/profile",
		},
		{
			name: "missing form data",
			postedData: url.Values{
				"email":    {""},
				"password": {""},
			},
			expectedStatusCode: http.StatusSeeOther,
			expectedLoc:        "/",
		},
		{
			name: "bad credentials",
			postedData: url.Values{
				"email":    {"you@abcd.com"},
				"password": {"password"},
			},
			expectedStatusCode: http.StatusSeeOther,
			expectedLoc:        "/",
		},
		{
			name: "user not found",
			postedData: url.Values{
				"email":    {"asdbc@gmail.com"},
				"password": {"password"},
			},
			expectedStatusCode: http.StatusSeeOther,
			expectedLoc:        "/",
		},
	}

	for _, e := range tests {
		req, _ := http.NewRequest("POST", "/login", strings.NewReader(e.postedData.Encode()))
		req = addContextAndSessionToRequest(req, app)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		rr := httptest.NewRecorder()
		hander := http.HandlerFunc(app.Login)
		hander.ServeHTTP(rr, req)

		if rr.Code != e.expectedStatusCode {
			t.Errorf("%s: returned wrong status code; expected %d, but got %d", e.name, e.expectedStatusCode, rr.Code)
		}
		actualLoc, err := rr.Result().Location()
		if err != nil {
			t.Errorf("%s: expected location %s but got %s", e.name, e.expectedLoc, actualLoc.String())
		} else {

			if actualLoc.String() != e.expectedLoc {
				t.Errorf("%s: redirected wrpmg Loc; expected %s, but got %s", e.name, e.expectedLoc, actualLoc)
			}

		}
	}
}
