package function

import (
	"fmt"
	"io"
	"net/http"

	shell "github.com/ipfs/go-ipfs-api"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	var input []byte
	sh := shell.NewShell("https://ipfs.devopsdebug.freeddns.org")
	err := sh.Get("QmS4wvjiYFKk2hP7DumxpE53ycBfyrjuzKv34MY9eadeMT", ".")
	if r.Body != nil {
		defer r.Body.Close()

		body, _ := io.ReadAll(r.Body)

		input = body
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Body: %s", string(input))))
	w.Write([]byte(fmt.Sprintf("Body: %s", string(err))))
}
