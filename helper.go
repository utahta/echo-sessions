package sessions

import "github.com/labstack/echo"

func Set(c echo.Context, key interface{}, v interface{}) {
	MustStart(c).Set(key, v)
}

func GetRaw(c echo.Context, key interface{}) interface{} {
	return MustStart(c).GetRaw(key)
}

func Get(c echo.Context, key interface{}, dst interface{}) (bool, error) {
	return MustStart(c).Get(key, dst)
}

func MustGet(c echo.Context, key interface{}, dst interface{}) bool {
	return MustStart(c).MustGet(key, dst)
}

func Delete(c echo.Context, key interface{}) {
	MustStart(c).Delete(key)
}

func Exists(c echo.Context, key interface{}) bool {
	return MustStart(c).Exists(key)
}

func Clear(c echo.Context) {
	MustStart(c).Clear()
}

func Flashes(c echo.Context, vars ...string) []interface{} {
	return MustStart(c).Flashes(vars...)
}

func AddFlash(c echo.Context, value interface{}, vars ...string) {
	MustStart(c).AddFlash(value, vars...)
}

func Save(c echo.Context) error {
	return MustStart(c).Save()
}
