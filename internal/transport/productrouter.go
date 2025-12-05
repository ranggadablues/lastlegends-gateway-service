package transport

import (
	"net/http"

	"lastlegends-gateway-service/middleware"

	"lastlegends-gateway-service/internal/service"

	"lastlegends-gateway-service/internal/endpoint"

	"lastlegends-gateway-service/helper"

	"github.com/gorilla/mux"
	"github.com/ranggadablues/gosok/common"
	pb "github.com/ranggadablues/lastlegends-proto-library/product-proto/pb"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func RegisterProductRoutes(mux *mux.Router, eps endpoint.ProductEndpoints) {
	const productsId = "/products/{id}"

	publicRouter := mux.PathPrefix("/").Subrouter()

	// GET /products (list products)
	publicRouter.Methods("GET").Path("/products").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		getProductsHandler(w, r, eps.Service)
	})

	// GET /products/{id} (get product by id)
	publicRouter.Methods("GET").Path(productsId).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		getProductByIdHandler(w, r, eps.Service)
	})

	privateRouter := mux.PathPrefix("/").Subrouter()

	// POST /products (create product)
	privateRouter.Methods("POST").Path("/products").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		createProductHandler(w, r, eps.Service)
	})

	// PUT /users/{id} (update user)
	privateRouter.Methods("PUT").Path(productsId).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		updateProductByIdHandler(w, r, eps.Service)
	})

	// DELETE /users/{id} (delete user)
	privateRouter.Methods("DELETE").Path(productsId).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		deleteProductByIdHandler(w, r, eps.Service)
	})
	privateRouter.Use(middleware.AuthMiddleware)
}

func getProductsHandler(w http.ResponseWriter, r *http.Request, productService service.IProductService) {
	req := &pb.ListProductsRequest{}
	resp, err := productService.GetProducts(r.Context(), req)
	if err != nil {
		helper.WriteErrorHeader(w, err)
		return
	}

	helper.WriteSuccessHeader(w, resp)
}

func getProductByIdHandler(w http.ResponseWriter, r *http.Request, productService service.IProductService) {
	vars := mux.Vars(r)
	req := &pb.GetProductByIdRequest{
		Id: vars["id"],
	}

	resp, err := productService.GetProductById(r.Context(), req)
	if err != nil {
		helper.WriteErrorHeader(w, err)
		return
	}

	helper.WriteSuccessHeader(w, resp)
}

func createProductHandler(w http.ResponseWriter, r *http.Request, productService service.IProductService) {
	var body pb.Product
	if err := common.Payload(&body, r); err != nil {
		helper.WriteErrorHeader(w, err)
		return
	}

	req := &pb.CreateProductRequest{
		Product: &body,
	}
	resp, err := productService.CreateProduct(r.Context(), req)
	if err != nil {
		helper.WriteErrorHeader(w, err)
		return
	}
	helper.WriteSuccessHeader(w, resp)
}

func updateProductByIdHandler(w http.ResponseWriter, r *http.Request, productService service.IProductService) {
	vars := mux.Vars(r)
	var payload bson.M
	if err := common.Payload(&payload, r); err != nil {
		helper.WriteErrorHeader(w, err)
		return
	}

	body := pb.Product{}
	bsonBytes, err := bson.Marshal(payload)
	if err != nil {
		helper.WriteErrorHeader(w, err)
		return
	}
	bson.Unmarshal(bsonBytes, &body)
	req := &pb.UpdateProductByIdRequest{
		Id:      vars["id"],
		Product: &body,
	}

	_, err = productService.UpdateProductById(r.Context(), req)
	if err != nil {
		helper.WriteErrorHeader(w, err)
		return
	}

	helper.WriteSuccessHeader(w, nil)
}

func deleteProductByIdHandler(w http.ResponseWriter, r *http.Request, productService service.IProductService) {
	vars := mux.Vars(r)
	req := &pb.DeleteProductByIdRequest{
		Id: vars["id"],
	}

	resp, err := productService.DeleteProductById(r.Context(), req)
	if err != nil {
		helper.WriteErrorHeader(w, err)
		return
	}

	helper.WriteSuccessHeader(w, resp)
}
