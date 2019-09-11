package v1

import (
	"bytes"
	"net/http/httptest"
	"testing"
)

func TestApiMessage(t *testing.T) {
	reader := bytes.NewReader([]byte(`
{
  "type": 1,
  "object": {
    "text": "Hello"
  }
}
`))
	req := httptest.NewRequest("GET", "/", reader)
	rw := httptest.NewRecorder()
	api := NewApi()

	api.OnEvent = func(e Event) {
		t.Errorf(`this message but not event`)
	}
	api.OnMessage = func(m Message) {
		if m.Text != "Hello" {
			t.Errorf(`wrong message text: %s expected Hello`, m.Text)
		}
	}

	api.ServeHTTP(rw, req)
}

func TestApiEvent(t *testing.T) {
	reader := bytes.NewReader([]byte(`
{
  "type": 2,
  "object": {
    "alert": "oops"
  }
}
`))
	req := httptest.NewRequest("GET", "/", reader)
	rw := httptest.NewRecorder()
	api := NewApi()

	api.OnEvent = func(e Event) {
		if e.Alert != "oops" {
			t.Errorf(`unexpected alert %s expect oops`, e.Alert)
		}
	}
	api.OnMessage = func(m Message) {
		t.Errorf(`this event but not message`)
	}

	api.ServeHTTP(rw, req)
}

func TestApiUnknown(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("unknown did not panic")
		}
	}()
	reader := bytes.NewReader([]byte(`
{
  "object": {
    "text": "Hello"
  }
}
`))
	req := httptest.NewRequest("GET", "/", reader)
	rw := httptest.NewRecorder()
	api := NewApi()

	api.OnEvent = func(e Event) {
		t.Errorf(`this unknown but not event`)
	}
	api.OnMessage = func(m Message) {
		t.Errorf(`this unknown but not message`)
	}

	api.ServeHTTP(rw, req)
}

func TestApiBrokenJSON(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("unknown did not panic")
		}
	}()
	reader := bytes.NewReader([]byte(`
  "object": {
    "text": "Hello"
  }
}
`))
	req := httptest.NewRequest("GET", "/", reader)
	rw := httptest.NewRecorder()
	api := NewApi()

	api.OnEvent = func(e Event) {
		t.Errorf(`this broken json but not event`)
	}
	api.OnMessage = func(m Message) {
		t.Errorf(`this broken json but not message`)
	}

	api.ServeHTTP(rw, req)
}