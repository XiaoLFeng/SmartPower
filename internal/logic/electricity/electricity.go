package electricity

import (
	"SmartPower/internal/dao"
	"SmartPower/internal/model/do"
	"SmartPower/internal/model/entity"
	"SmartPower/internal/service"
	"context"
	"github.com/bamboo-services/bamboo-utils/bcode"
	"github.com/bamboo-services/bamboo-utils/berror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

type sElectric struct {
}

func init() {
	service.RegisterElectric(New())
}

func New() *sElectric {
	return &sElectric{}
}

// GetCompanyByHeader
//
// # 通过请求头获取企业
//
// 通过请求头获取企业, 用于通过请求头获取企业. 用户在这里可以通过请求头获取企业，获取企业的时候需要用户提供请求头，获取成功后会返回企业信息.
// 将会从请求头获取授权信息，授权信息将会直接查询指定用户信息，然后查询企业信息.
//
// # 参数:
//   - ctx: context.Context, 上下文
//
// # 返回:
//   - company: *entity.XfCompanies, 企业信息
//   - err: error, 错误信息
func (s *sElectric) GetCompanyByHeader(ctx context.Context) (company *entity.XfCompanies, err error) {
	glog.Noticef(ctx, "[LOGIC] 执行 GetCompanyByHeader | 通过请求头获取企业")
	// 获取请求头
	getRequest := g.RequestFromCtx(ctx)
	getToken := getRequest.Header.Get("Authorization")
	userInfo, err := service.Token().GetUserByToken(ctx, getToken)
	if err != nil {
		return nil, berror.NewErrorHasError(bcode.ServerInternalError, err)
	}
	if userInfo == nil {
		return nil, berror.NewError(bcode.OperationFailed, "用户不存在")
	}
	// 根据用户查询企业
	var getCompany *entity.XfCompanies
	err = dao.XfCompanies.Ctx(ctx).Where(do.XfCompanies{Uuid: userInfo.Uuid}).Scan(&getCompany)
	if err != nil {
		return nil, berror.NewErrorHasError(bcode.ServerInternalError, err)
	}
	if getCompany == nil {
		return nil, berror.NewError(bcode.OperationFailed, "企业不存在")
	}
	return getCompany, nil
}
