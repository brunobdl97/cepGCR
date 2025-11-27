package internal

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func newTestClient(t *testing.T, fn roundTripFunc) *http.Client {
	t.Helper()
	return &http.Client{Transport: roundTripFunc(fn)}
}

func response(body string, status int) *http.Response {
	return &http.Response{
		StatusCode: status,
		Header:     make(http.Header),
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}

func TestHandleSuccess(t *testing.T) {
	h := &GetWeatherHandler{client: newTestClient(t, func(r *http.Request) (*http.Response, error) {
		switch {
		case strings.Contains(r.URL.Host, "viacep.com.br"):
			return response(`{"localidade":"Sao Paulo"}`, http.StatusOK), nil
		case strings.Contains(r.URL.Host, "api.weatherapi.com"):
			return response(`{"current":{"temp_c":20}}`, http.StatusOK), nil
		default:
			return nil, errors.New("unexpected host")
		}
	})}

	req := httptest.NewRequest(http.MethodGet, "/cep/12345678", nil)
	req.SetPathValue("cep", "12345678")
	rec := httptest.NewRecorder()

	h.Handle(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	body := rec.Body.String()
	if !strings.Contains(body, `"temp_C":20`) {
		t.Fatalf("expected celsius value in response, got %s", body)
	}
	if !strings.Contains(body, `"temp_F":68`) {
		t.Fatalf("expected fahrenheit conversion, got %s", body)
	}
	if !strings.Contains(body, `"temp_K":293.15`) {
		t.Fatalf("expected kelvin conversion, got %s", body)
	}
}

func TestHandleInvalidCep(t *testing.T) {
	h := NewGetWeatherHandler()
	req := httptest.NewRequest(http.MethodGet, "/cep/123", nil)
	rec := httptest.NewRecorder()

	h.Handle(rec, req)

	if rec.Code != http.StatusUnprocessableEntity {
		t.Fatalf("expected status 422, got %d", rec.Code)
	}
}

func TestHandleViaCepError(t *testing.T) {
	h := &GetWeatherHandler{client: newTestClient(t, func(r *http.Request) (*http.Response, error) {
		return nil, errors.New("viacep down")
	})}

	req := httptest.NewRequest(http.MethodGet, "/cep/12345678", nil)
	req.SetPathValue("cep", "12345678")
	rec := httptest.NewRecorder()

	h.Handle(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", rec.Code)
	}

	if body := strings.TrimSpace(rec.Body.String()); body != "cannot find zipcode" {
		t.Fatalf("unexpected body: %s", body)
	}
}

func TestHandleWeatherApiError(t *testing.T) {
	h := &GetWeatherHandler{client: newTestClient(t, func(r *http.Request) (*http.Response, error) {
		switch {
		case strings.Contains(r.URL.Host, "viacep.com.br"):
			return response(`{"localidade":"Sao Paulo"}`, http.StatusOK), nil
		case strings.Contains(r.URL.Host, "api.weatherapi.com"):
			return response(`{"error":"not found"}`, http.StatusNotFound), nil
		default:
			return nil, errors.New("unexpected host")
		}
	})}

	req := httptest.NewRequest(http.MethodGet, "/cep/12345678", nil)
	req.SetPathValue("cep", "12345678")
	rec := httptest.NewRecorder()

	h.Handle(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", rec.Code)
	}

	if !strings.Contains(rec.Body.String(), "error fetching weather data") {
		t.Fatalf("unexpected body: %s", rec.Body.String())
	}
}
