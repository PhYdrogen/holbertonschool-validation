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
      responseCode: 404,
      body:         "404 page not found\n",
    },
    {
      name:         "Hello page",
      URI:          "/hello?name=Holberton",
      responseCode: 200,
      body:         "Hello Holberton!",
    },
		{
      name:         "Health",
      URI:  "/health",
      responseCode: 200,
      body:         "ALIVE",
    },
		{
      name:         "Mr Jones",
      URI:  "/hello?name=Mr Jones",
      responseCode: 400,
      body:         "400 Bad Request",
    },
		{
      name:         "Hello nothing",
      URI:          "/hello?name=",
      responseCode: 400,
      body:         "",
    },
		{
      name:         "Hello weird",
      URI:          "/hello?name=%%",
      responseCode: 200,
      body:         "Hello Gab!",
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