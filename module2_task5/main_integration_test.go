package main

import (
  //nolint:staticcheck
  "io/ioutil"
  "net/http"
  "net/http/httptest"
  "testing"
)

func Test_server(t *testing.T) {
  if testing.Short() {
    t.Skip("Flag `-short` provided: skipping Integration Tests.")
  }

  tests := []struct {
    name         string
    URI          string
    responseCode int
    body         string
  }{
    {
      name:         "Home page",
      URI:          "",
      responseCode: 200,
      body:         "<!DOCTYPE html>\n<html lang=\"en-us\">\n  <head>\n    <meta charset=\"utf-8\">\n    <meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge,chrome=1\">\n    \n    <title>Awesome Inc.</title>\n    <meta name=\"viewport\" content=\"width=device-width,minimum-scale=1\">\n    <meta name=\"description\" content=\"\">\n    <meta name=\"generator\" content=\"Hugo 0.84.0\" />\n    \n    \n    \n    \n      <meta name=\"robots\" content=\"noindex, nofollow\">\n    \n\n    \n<link rel=\"stylesheet\" href=\"/ananke/css/main.min.css\" >\n\n\n\n    \n    \n    \n      \n\n    \n\n    \n    \n      <link href=\"/index.xml\" rel=\"alternate\" type=\"application/rss+xml\" title=\"Awesome Inc.\" />\n      <link href=\"/index.xml\" rel=\"feed\" type=\"application/rss+xml\" title=\"Awesome Inc.\" />\n      \n    \n    \n    <meta property=\"og:title\" content=\"Awesome Inc.\" />\n<meta property=\"og:description\" content=\"\" />\n<meta property=\"og:type\" content=\"website\" />\n<meta property=\"og:url\" content=\"http://example.org/\" />\n\n<meta itemprop=\"name\" content=\"Awesome Inc.\">\n<meta itemprop=\"description\" content=\"\"><meta name=\"twitter:card\" content=\"summary\"/>\n<meta name=\"twitter:title\" content=\"Awesome Inc.\"/>\n<meta name=\"twitter:description\" content=\"\"/>\n\n\t\n  </head>\n\n  <body class=\"ma0 avenir bg-near-white\">\n\n    \n\n  <header>\n    <div class=\"pb3-m pb6-l bg-black\">\n      <nav class=\"pv3 ph3 ph4-ns\" role=\"navigation\">\n  <div class=\"flex-l justify-between items-center center\">\n    <a href=\"/\" class=\"f3 fw2 hover-white no-underline white-90 dib\">\n      \n        Awesome Inc.\n      \n    </a>\n    <div class=\"flex-l items-center\">\n      \n\n      \n      \n<div class=\"ananke-socials\">\n  \n</div>\n\n    </div>\n  </div>\n</nav>\n\n      <div class=\"tc-l pv3 ph3 ph4-ns\">\n        <h1 class=\"f2 f-subheadline-l fw2 light-silver mb0 lh-title\">\n          Awesome Inc.\n        </h1>\n        \n      </div>\n    </div>\n  </header>\n\n\n    <main class=\"pb7\" role=\"main\">\n      \n <article class=\"cf ph3 ph5-l pv3 pv4-l f4 tc-l center measure-wide lh-copy mid-gray\">\n    \n  </article>\n  \n  \n  \n  \n  \n  \n  \n    \n    \n\n    <div class=\"pa3 pa4-ns w-100 w-70-ns center\">\n      \n       \n          <h1 class=\"flex-none\">\n            Recent Posts\n          </h1>\n        \n\n      \n\n      <section class=\"w-100 mw8\">\n        \n        \n          <div class=\"relative w-100 mb4\">\n            \n<article class=\"bb b--black-10\">\n  <div class=\"db pv4 ph3 ph0-l no-underline dark-gray\">\n    <div class=\"flex flex-column flex-row-ns\">\n      \n      <div class=\"blah w-100\">\n        <h1 class=\"f3 fw1 athelas mt0 lh-title\">\n          <a href=\"/posts/welcome/\" class=\"color-inherit dim link\">\n            Welcome to Awesome Inc.\n            </a>\n        </h1>\n        <div class=\"f6 f5-l lh-copy nested-copy-line-height nested-links\">\n          Enter your name below and click on the button &ldquo;Say Hello&rdquo;:\nSay Hello   function loadXMLDoc() { var xhttp = new XMLHttpRequest(); xhttp.onreadystatechange = function() { if (this.readyState == 4 && this.status == 200) { document.getElementById(\"helloResponse\").innerHTML = this.responseText; } }; xhttp.open(\"GET\", \"/hello?name=\" + document.getElementById(\"name\").value, true); xhttp.send(); }   \n        </div>\n          <a href=\"/posts/welcome/\" class=\"ba b--moon-gray bg-light-gray br2 color-inherit dib f7 hover-bg-moon-gray link mt2 ph2 pv1\">read more</a>\n        \n      </div>\n    </div>\n  </div>\n</article>\n\n          </div>\n        \n      </section>\n\n      \n\n      </div>\n  \n\n    </main>\n    <footer class=\"bg-black bottom-0 w-100 pa3\" role=\"contentinfo\">\n  <div class=\"flex justify-between\">\n  <a class=\"f4 fw4 hover-white no-underline white-70 dn dib-ns pv2 ph3\" href=\"http://example.org/\" >\n    &copy;  Awesome Inc. 2023 \n  </a>\n    <div>\n<div class=\"ananke-socials\">\n  \n</div>\n</div>\n  </div>\n</footer>\n\n  </body>\n</html>\n",
    },
  }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      ts := httptest.NewServer(setupRouter())
      defer ts.Close()

      res, err := http.Get(ts.URL + tt.URI)
      if err != nil {
        t.Fatal(err)
      }

      // Check that the status code is what you expect.
      expectedCode := tt.responseCode
      gotCode := res.StatusCode
      if gotCode != expectedCode {
        t.Errorf("handler returned wrong status code: got %q want %q", gotCode, expectedCode)
      }

      // Check that the response body is what you expect.
      expectedBody := tt.body
      bodyBytes, err := ioutil.ReadAll(res.Body)
      res.Body.Close()
      if err != nil {
        t.Fatal(err)
      }
      gotBody := string(bodyBytes)
      if gotBody != expectedBody {
        t.Errorf("handler returned unexpected body: got %q want %q", gotBody, expectedBody)
      }
    })
  }
}