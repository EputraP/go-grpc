package services

import (
	productPb "go-grpc/pb/product"
)

type ProductService struct {
	productPb.UnimplementedProductServiceServer
}
