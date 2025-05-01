package report_revenue

type Out struct {
	Records []OutRecord `json:"records"`
}

type OutRecord struct {
	ProductTitle string `json:"product_title"`
	TotalRevenue int64  `json:"total_revenue"`
}
