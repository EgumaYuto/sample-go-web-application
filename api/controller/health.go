package controller

import (
	"fmt"
	"log"
	"net/http"

	"cabos.io/model"
	"github.com/julienschmidt/httprouter"
)

type HealthController struct{}

func (_ HealthController) GetHealthStatus(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	_, err := model.HealthCheck()
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Error!!")
	} else {
		fmt.Fprintf(w, "OK!")
	}
}
