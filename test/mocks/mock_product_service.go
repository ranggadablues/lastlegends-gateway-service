package mocks

import (
	"context"

	pb "github.com/ranggadablues/lastlegends-proto-library/product-proto/pb"
)

type MockProductService struct {
	GetProductsFn       func(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error)
	GetProductByIdFn    func(ctx context.Context, req *pb.GetProductByIdRequest) (*pb.GetProductByIdResponse, error)
	CreateProductFn     func(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error)
	UpdateProductByIdFn func(ctx context.Context, req *pb.UpdateProductByIdRequest) (*pb.UpdateProductByIdResponse, error)
	DeleteProductByIdFn func(ctx context.Context, req *pb.DeleteProductByIdRequest) (*pb.DeleteProductByIdResponse, error)
}

func (m *MockProductService) GetProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	return m.GetProductsFn(ctx, req)
}
func (m *MockProductService) GetProductById(ctx context.Context, req *pb.GetProductByIdRequest) (*pb.GetProductByIdResponse, error) {
	return m.GetProductByIdFn(ctx, req)
}
func (m *MockProductService) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	return m.CreateProductFn(ctx, req)
}
func (m *MockProductService) UpdateProductById(ctx context.Context, req *pb.UpdateProductByIdRequest) (*pb.UpdateProductByIdResponse, error) {
	return m.UpdateProductByIdFn(ctx, req)
}
func (m *MockProductService) DeleteProductById(ctx context.Context, req *pb.DeleteProductByIdRequest) (*pb.DeleteProductByIdResponse, error) {
	return m.DeleteProductByIdFn(ctx, req)
}
