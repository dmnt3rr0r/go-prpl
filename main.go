package go_prpl

import (
	"github.com/gin-gonic/gin"
)



func main() {
	// use the bcapMap to lookup the capabilities of a browser
	// TODO: test that the bcapMap is set up correctly
	// TODO: unit test the normalize and the uaCaps
	_, _ = initCapMap("browscap.ini", "capmap.yaml")

	r := gin.Default()

	r.Run()
}
