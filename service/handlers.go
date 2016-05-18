package service

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// getAllCatalogItemsHandler returns a fake list of catalog items
func getAllCatalogItemsHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		catalog := make([]catalogItem, 2)
		catalog[0] = fakeItem("ABC1234")
		catalog[1] = fakeItem("STAPLER99")
		formatter.JSON(w, http.StatusOK, catalog)
	}
}

// getCatalogItemDetailsHandler returns a fake catalog item. The key takeaway here
// is that we're using a backing service to get fulfillment status for the individual
// item.
func getCatalogItemDetailsHandler(formatter *render.Render, serviceClient fulfillmentClient) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		sku := vars["sku"]
		status, err := serviceClient.getFulfillmentStatus(sku)
		if err == nil {
			formatter.JSON(w, http.StatusOK, catalogItem{
				ProductID:       1,
				SKU:             sku,
				Description:     "This is a fake product",
				Price:           1599, // $15.99
				ShipsWithin:     status.ShipsWithin,
				QuantityInStock: status.QuantityInStock,
			})
		} else {
			formatter.JSON(w, http.StatusInternalServerError, fmt.Sprintf("Fulfillment Client error: %s", err.Error()))
		}
	}
}

func rootHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.Text(w, http.StatusOK, "Catalog Service, see http://github.com/cloudnativego/backing-catalog for API.")
	}
}

func fakeItem(sku string) (item catalogItem) {
	item.SKU = sku
	item.Description = "This is a fake product"
	item.Price = 1599
	item.QuantityInStock = 75
	item.ShipsWithin = 14
	return
}
