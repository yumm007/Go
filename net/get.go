package main
import "log"
import "net/http"

func main() {
	_, err := http.Get("http://www.cdsslz.net")
	if err != nil {
		log.Fatal(err)	
	}
}

