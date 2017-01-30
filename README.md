# Sessions middleware for Echo

[![Build Status](https://travis-ci.org/utahta/echo-sessions.svg?branch=master)](https://travis-ci.org/utahta/echo-sessions)

This is a package of sessions middleware for Echo.  
A thin [gorilla/sessions](https://github.com/gorilla/sessions) wrapper.

## Install

```
$ go get -u github.com/utahta/echo-sessions
```

## Usage

Use middleware
```go
import (
    gsessions "github.com/gorilla/sessions"
    "github.com/utahta/echo-sessions"
    "github.com/labstack/echo"
)

e := echo.New()
e.Use(sessions.Sessions("SESSID", gsessions.NewCookieStore()))
```

Start sessions
```go
s := sessions.MustStart()
```

Set key and value
```go
s.Set("key", "value")
```

Get value by key
```go
var v string
err := s.Get("key", &v)
```

Check key exists
```go
var v string
if err := s.Get("key", &v); err == sessions.ErrorNoSuchKey {
    s.Set("key", "new value")
}
```

Delete key
```go
s.Delete("key")
```

Save this session
```go
err := s.Save()
```

## Contributing

1. Fork it ( https://github.com/utahta/echo-sessions/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request

