package services

import (
	"context"
	pagingPb "go-grpc/pb/pagination"
	productPb "go-grpc/pb/product"
)

type ProductService struct {
	productPb.UnimplementedProductServiceServer
}

func (p *ProductService) GetProducts(context.Context, *productPb.Empty) (*productPb.Products, error) {
	products := &productPb.Products{
		Pagination: &pagingPb.Pagination{
			Total:       10,
			PerPage:     5,
			CurrentPage: 1,
			LastPage:    2,
		},
		Data: []*productPb.Product{
			{
				Id:    1,
				Name:  "Metallica T-Shirt",
				Price: 10000,
				Stock: 10,
				Category: &productPb.Category{
					Id:   1,
					Name: "Shirt",
				},
			}, {
				Id:    2,
				Name:  "Wibu T-Shirt",
				Price: 20000,
				Stock: 20,
				Category: &productPb.Category{
					Id:   1,
					Name: "Shirt",
				},
			},
		},
	}
	return products, nil
}
