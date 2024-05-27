package dcompany

// DCompany
//
// # 公司
//
// 公司信息.
//
// # 方法
//   - Cods: string, 公司代码
//   - Name: string, 公司名称
//   - Address: string, 公司地址
//   - Representative: string, 公司代表
type DCompany struct {
	Cods           string `json:"cods"`
	Name           string `json:"name"`
	Address        string `json:"address"`
	Representative string `json:"representative"`
}
