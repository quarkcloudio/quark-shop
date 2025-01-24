package response

// 商品分类
type ItemCategoryResp struct {
	Id       int                `json:"id"`
	Pid      int                `json:"pid"`
	Title    string             `json:"title"`
	CoverId  string             `json:"cover_id,omitempty"`
	Children []ItemCategoryResp `json:"children,omitempty" gorm:"-"`
}

// 商品列表
type ItemIndexResp struct {
	Id         int     `json:"id"`          // 商品id
	Name       string  `json:"name"`        // 商品名称
	Image      string  `json:"image"`       // 商品图片
	Price      float64 `json:"price"`       // 商品价格
	FictiSales int     `json:"ficti_sales"` // 商品虚拟销量
}

// 商品详情
type ItemDetailResp struct {
	Id          int                 `json:"id"`           // 商品id
	Name        string              `json:"name"`         // 商品名称
	Price       float64             `json:"price"`        // 商品价格
	SliderImage []string            `json:"slider_image"` // 商品轮播图
	OtPrice     float64             `json:"ot_price"`     // 市场价（划线价）
	Stock       int                 `json:"stock"`        // 库存
	FictiSales  int                 `json:"ficti_sales"`  // 商品虚拟销量
	Content     string              `json:"content"`      // 商品详情
	Status      int                 `json:"status"`       // 商品状态（0-未上架；1-上架）
	SpecType    int                 `json:"spec_type"`    // 规格（0-单；1-多）
	AttrValues  []ItemAttrValueResp `json:"attr_values"`  // 规格列表
}

// 商品属性
type ItemAttrValueResp struct {
	AttrValueId        int     `json:"attr_value_id"`
	AttrValueStock     int     `json:"attr_value_stock"`
	AttrValueSales     int     `json:"attr_value_sales"`
	AttrValuePrice     float64 `json:"attr_value_price"`
	AttrValueOtPrice   float64 `json:"attr_value_ot_price"`
	AttrValueImage     string  `json:"attr_value_image"`
	AttrValueIsDefault bool    `json:"attr_value_is_default"`
}
