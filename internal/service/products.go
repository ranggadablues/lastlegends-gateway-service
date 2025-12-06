package service

import (
	"context"

	pb "github.com/ranggadablues/lastlegends-proto-library/product-proto/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IProductService interface {
	GetProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error)
	GetProductById(ctx context.Context, req *pb.GetProductByIdRequest) (*pb.GetProductByIdResponse, error)
	CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error)
	UpdateProductById(ctx context.Context, req *pb.UpdateProductByIdRequest) (*pb.UpdateProductByIdResponse, error)
	DeleteProductById(ctx context.Context, req *pb.DeleteProductByIdRequest) (*pb.DeleteProductByIdResponse, error)
}

type productService struct {
	client pb.ProductServiceClient
}

func NewProductService(addr string) IProductService {
	conn, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	return &productService{client: pb.NewProductServiceClient(conn)}
}

func (s *productService) GetProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	return s.client.ListProducts(ctx, req)
}

func (s *productService) GetProductById(ctx context.Context, req *pb.GetProductByIdRequest) (*pb.GetProductByIdResponse, error) {
	return s.client.GetProductById(ctx, req)
}

func (s *productService) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	return s.client.CreateProduct(ctx, req)
}

func (s *productService) UpdateProductById(ctx context.Context, req *pb.UpdateProductByIdRequest) (*pb.UpdateProductByIdResponse, error) {
	return s.client.UpdateProductById(ctx, req)
}

func (s *productService) DeleteProductById(ctx context.Context, req *pb.DeleteProductByIdRequest) (*pb.DeleteProductByIdResponse, error) {
	return s.client.DeleteProductById(ctx, req)
}
