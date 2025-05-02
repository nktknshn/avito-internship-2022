package report_revenue

import domainProduct "github.com/nktknshn/avito-internship-2022/internal/balance/domain/product"

type Out struct {
	Records []OutRecord
}

type OutRecord struct {
	ProductID    domainProduct.ProductID
	ProductTitle domainProduct.ProductTitle
	TotalRevenue int64
}
