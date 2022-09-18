package productcontroller

import (
	"net/http"

	"github.com/swildz/go-jwt-siddiq/helper"
)

func Index(w http.ResponseWriter, r *http.Request) {

	data := []map[string]interface{}{
		{
			"id":           1,
			"nama_product": "kameja",
			"stok":         1000,
		},
		{
			"id":           w,
			"nama_product": "celana",
			"stok":         1000,
		},
		{
			"id":           3,
			"nama_product": "sepatu",
			"stok":         100,
		},
	}
	helper.ResponJSON(w, http.StatusOK, data)
}
