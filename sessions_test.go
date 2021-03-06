package sessions

import (
	"net/http"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
)

func testContext() echo.Context {
	e := echo.New()
	return e.NewContext(&http.Request{}, nil)
}

func testSession(c echo.Context) *session {
	return &session{ctx: c, name: "TESTSESSID", store: sessions.NewCookieStore()}
}

func TestStart(t *testing.T) {
	c := testContext()

	if _, err := Start(c); err != ErrSessionNotFound {
		t.Errorf("Expected get error session not found, got %v", err)
	}

	c.Set(contextKey, testSession(c))
	s, err := Start(c)
	if err != nil {
		t.Error(err)
	}

	if s == nil {
		t.Error("Expected get session, got nil")
	}
}

func TestMustStart(t *testing.T) {
	c := testContext()
	c.Set(contextKey, testSession(c))

	if s := MustStart(c); s == nil {
		t.Error("Expected get session, got nil")
	}
}

func TestMustStartPanic(t *testing.T) {
	c := testContext()
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected get panic")
		}
	}()

	MustStart(c)
}

func TestSession_Set(t *testing.T) {
	c := testContext()
	c.Set(contextKey, testSession(c))
	s := MustStart(c)

	tests := []struct {
		key      string
		expected int
	}{
		{"aaa", 100},
		{"bbb", 200},
	}

	for _, test := range tests {
		s.Set(test.key, test.expected)

		v, ok := s.Session.Values[test.key]
		if !ok {
			t.Error("Expected value exists, got empty")
		}

		if v != test.expected {
			t.Errorf("Expected value %v, got %v", test.expected, v)
		}
	}
}

func TestSession_Get(t *testing.T) {
	c := testContext()
	c.Set(contextKey, testSession(c))
	s := MustStart(c)
	const key = "key"

	// test basic type
	s.Set(key, 100)
	var dstInt int
	if ok, err := s.Get(key, &dstInt); !ok && err == nil {
		t.Error(err)
	}
	if dstInt != 100 {
		t.Errorf("Expected get 100, got %v", dstInt)
	}

	// test pointer
	dstInt = 0
	srcInt := 100
	s.Set(key, &srcInt)
	if ok, err := s.Get(key, &dstInt); !ok && err == nil {
		t.Error(err)
	}
	if dstInt != 100 {
		t.Errorf("Expected get %v, got %v", srcInt, dstInt)
	}

	// test struct
	type testStruct struct {
		Num     int
		Name    string
		Request *http.Request
	}
	srcStruct := testStruct{Num: 1, Name: "sess", Request: &http.Request{Method: "PUT"}}
	s.Set(key, &srcStruct)
	var dstStruct testStruct
	if ok, err := s.Get(key, &dstStruct); !ok && err == nil {
		t.Errorf("Expected get struct, got %v", dstStruct)
	}
	if srcStruct.Num != dstStruct.Num {
		t.Errorf("Expected get struct Num, got %v", dstStruct.Num)
	}
	if srcStruct.Name != dstStruct.Name {
		t.Errorf("Expected get struct Name, got %v", dstStruct.Name)
	}
	if srcStruct.Request.Method != dstStruct.Request.Method {
		t.Errorf("Expected get struct Request.Method, got %v", dstStruct.Request.Method)
	}

	// test not pointer
	dstInt = 0
	if _, err := s.Get(key, dstInt); err == nil {
		t.Error("Expected get error, got nil")
	}

	// test not compared type
	srcFloat := 1.0
	dstInt = 0
	s.Set(key, &srcFloat)
	if _, err := s.Get(key, &dstInt); err == nil {
		t.Error("Expected get error, got nil")
	}

	// test no such key
	if ok, err := s.Get("no_such_key", nil); ok {
		t.Error(err)
	}
}

func TestSession_MustGet(t *testing.T) {
	c := testContext()
	c.Set(contextKey, testSession(c))
	s := MustStart(c)

	var dst int
	if ok := s.MustGet("key", &dst); ok {
		t.Error("Expected get false, got true")
	}

	s.Set("key", 100)
	if ok := s.MustGet("key", &dst); !ok {
		t.Error("Expected get true, got false")
	}
	if dst != 100 {
		t.Errorf("Expected get 100, got %d", dst)
	}
}

func TestSession_MustGetPanic(t *testing.T) {
	c := testContext()
	c.Set(contextKey, testSession(c))
	s := MustStart(c)
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected get panic")
		}
	}()

	var dst int
	s.Set("key", 100)
	s.MustGet("key", dst)
}

func TestSession_GetRaw(t *testing.T) {
	c := testContext()
	c.Set(contextKey, testSession(c))
	s := MustStart(c)
	var v interface{}

	v = s.GetRaw("key")
	if v != nil {
		t.Error("Expected get false, got true")
	}

	s.Set("key", "value")
	v = s.GetRaw("key")
	if v == nil {
		t.Error("Expected get true, got false")
	}

	vv, ok := v.(string)
	if !ok {
		t.Error("Expected get true, got false")
	}
	if vv != "value" {
		t.Errorf("Expected get value, got %v", vv)
	}
}

func TestSession_Delete(t *testing.T) {
	c := testContext()
	c.Set(contextKey, testSession(c))
	s := MustStart(c)

	s.Set("key1", 100)
	s.Set("key2", 200)

	if _, ok := s.Session.Values["key1"]; !ok {
		t.Error("Expected key1 exists, got empty")
	}
	if _, ok := s.Session.Values["key2"]; !ok {
		t.Error("Expected key2 exists, got empty")
	}

	s.Delete("key1")
	if _, ok := s.Session.Values["key1"]; ok {
		t.Error("Expected key1 empty, got key exists")
	}
	if _, ok := s.Session.Values["key2"]; !ok {
		t.Error("Expected key2 exists, got empty")
	}

	s.Delete("key2")
	if _, ok := s.Session.Values["key1"]; ok {
		t.Error("Expected key1 empty, got key1 exists")
	}
	if _, ok := s.Session.Values["key2"]; ok {
		t.Error("Expected key2 empty, got key2 exists")
	}
}

func TestSession_Exists(t *testing.T) {
	c := testContext()
	c.Set(contextKey, testSession(c))
	s := MustStart(c)

	if s.Exists("key") {
		t.Error("Expected get false, got true")
	}

	s.Set("key", "value")
	if !s.Exists("key") {
		t.Error("Expected get true, got false")
	}
}
