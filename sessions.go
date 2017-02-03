package sessions

import (
	"reflect"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

const contextKey = "github.com/utahta/echo-sessions"

var (
	ErrSessionNotFound = errors.New("Session not found")
)

type session struct {
	ctx     echo.Context
	name    string
	store   sessions.Store
	Session *sessions.Session
}

// Get the session handler
func Start(c echo.Context) (*session, error) {
	s, ok := c.Get(contextKey).(*session)
	if !ok {
		return nil, ErrSessionNotFound
	}

	if s.Session != nil {
		return s, nil
	}

	ss, err := s.store.New(c.Request(), s.name)
	if err != nil {
		return nil, err
	}
	s.Session = ss
	return s, nil
}

// Get the session handler
// if get error, cause panic
func MustStart(c echo.Context) *session {
	s, err := Start(c)
	if err != nil {
		panic(err)
	}
	return s
}

func (s *session) Set(key interface{}, v interface{}) {
	s.Session.Values[key] = v
}

func (s *session) GetRaw(key interface{}) interface{} {
	return s.Session.Values[key]
}

func (s *session) Get(key interface{}, dst interface{}) (bool, error) {
	v := s.GetRaw(key)
	if v == nil {
		return false, nil
	}
	set := reflect.ValueOf(dst)
	if set.Kind() != reflect.Ptr {
		return false, errors.Errorf("Expected pointer to a struct, got %#v", dst)
	}
	if !set.CanSet() {
		set = set.Elem()
	}

	vv := reflect.ValueOf(v)
	if vv.Kind() == reflect.Ptr {
		vv = vv.Elem()
	}

	if set.Type() != vv.Type() {
		return false, errors.Errorf("Expected same type, got %v and %v", set.Type(), vv.Type())
	}
	set.Set(vv)
	return true, nil
}

func (s *session) MustGet(key interface{}, dst interface{}) bool {
	ok, err := s.Get(key, dst)
	if err != nil {
		panic(err)
	}
	return ok
}

func (s *session) Delete(key interface{}) {
	delete(s.Session.Values, key)
}

func (s *session) Exists(key interface{}) bool {
	return s.GetRaw(key) != nil
}

func (s *session) Clear() {
	for k := range s.Session.Values {
		s.Delete(k)
	}
}

// gorilla/sessions Flashes wrap
func (s *session) Flashes(vars ...string) []interface{} {
	return s.Session.Flashes(vars...)
}

// gorilla/sessions AddFlash wrap
func (s *session) AddFlash(value interface{}, vars ...string) {
	s.Session.AddFlash(value, vars...)
}

func (s *session) Save() error {
	return s.Session.Save(s.ctx.Request(), s.ctx.Response().Writer())
}
