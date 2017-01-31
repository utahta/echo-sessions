# Sessions middleware for Echo

[![GitHub release](https://img.shields.io/github/release/utahta/echo-sessions.svg)](https://github.com/utahta/echo-sessions/releases) 
[![Build Status](https://travis-ci.org/utahta/echo-sessions.svg?branch=master)](https://travis-ci.org/utahta/echo-sessions)

This is a package of sessions middleware for Echo.  
A thin [gorilla/sessions](https://github.com/gorilla/sessions) wrapper.

## Install

```
$ go get -u github.com/utahta/echo-sessions
```

## Usage

### Use sessions middleware
```go
import (
    "github.com/boj/redistore"
    "github.com/utahta/echo-sessions"
    "github.com/labstack/echo"
)

store, _ := redistore.NewRediStore(10, "tcp", ":6379", "", []byte("secret-key"))

e := echo.New()
e.Use(sessions.Sessions("SESSID", store))
```

### Start session
```go
s := sessions.MustStart()
```

### Set key and value
```go
s.Set("key", "value")
```

### Get value by key
```go
var v string
ok, err := s.Get("key", &v)
```
or
```go
var v string
ok := s.MustGet("key", &v)
```

### Get raw value by key
```go
v := s.GetRaw("key") // returns (interface{})
```

### Check key exists
```go
if !s.Exists("key") {
    s.Set("key", "new value")
}
```
or
```go
var v string
if ok, err := s.Get("example1", &v); !ok && err == nil {
    s.Set("example1", "new value")
}

if ok := s.MustGet("example2", &v); !ok {
    s.Set("example2", "new value")
}
```

### Delete key
```go
s.Delete("key")
```

### Save this session
```go
err := s.Save()
```

## Contributing

1. Fork it ( https://github.com/utahta/echo-sessions/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request

