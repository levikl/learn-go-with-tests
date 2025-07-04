package acceptancetests

import (
	"fmt"
	"net/http"
	"time"
)

func SlowHandler(w http.ResponseWriter, _ *http.Request) {
	time.Sleep(2 * time.Second)
	fmt.Fprint(w, "Hello, world")
}
