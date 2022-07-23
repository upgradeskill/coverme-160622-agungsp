package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type stdHandler func(w http.ResponseWriter, r *http.Request)

func myHttpTest(t *testing.T, method string, url string, stdhandler stdHandler, expect string, body io.Reader) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
    // pass 'nil' as the third parameter.
    req, err := http.NewRequest(method, url, body)
    if err != nil {
        t.Fatal(err)
    }

    // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(stdhandler)

    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method
    // directly and pass in our Request and ResponseRecorder.
    handler.ServeHTTP(rr, req)

    // Check the response body is what we expect.
    if rr.Body.String() != expect {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expect)
    } else {
		t.Logf("handler returned expected body: got %v want %v",
            rr.Body.String(), expect)
	}
}

func TestProductHandlerCreated(t *testing.T) {
	myHttpTest(
		t, "GET", "/products", ProductHandlerInit().Index, "[]", nil,
	)
}

func TestProductHandlerImplementController(t * testing.T)  {
	result := true
	st := reflect.TypeOf(ProductHandlerInit())
	_, ok := st.MethodByName("Route")
	if result {
		result = ok
	}

	_, ok = st.MethodByName("Index")
	if result {
		result = ok
	}

	_, ok = st.MethodByName("Create")
	if result {
		result = ok
	}
	_, ok = st.MethodByName("Show")
	if result {
		result = ok
	}
	_, ok = st.MethodByName("Update")
	if result {
		result = ok
	}
	_, ok = st.MethodByName("Delete")
	if result {
		result = ok
	}

	if !result {
		t.Error("Handler not implement controller interface")
	}
}
