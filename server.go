package poker_api

import (
	"fmt"
	"net/http"
)

func PlayerServer(req *http.Request, res http.ResponseWriter) {
	fmt.Fprintf(res, "20")
}
