package controller

import "net/http"

func Router() {
	http.ListenAndServe("localhost:8080", han)
}

func User() {

}
