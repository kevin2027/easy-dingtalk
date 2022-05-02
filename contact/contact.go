package contact

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/kevin2027/easy-dingtalk/oauth2"
	"github.com/kevin2027/easy-dingtalk/utils"
	"golang.org/x/xerrors"
)

type Contact interface {
	GetUserInfo(userid string) (res *GetUserInfoResponseResult, err error)
}

func NewContact(oauth2 oauth2.Oauth2) (r Contact) {
	return &inner{
		oauth2: oauth2,
	}
}

type inner struct {
	oauth2 oauth2.Oauth2
}

func (d *inner) GetUserInfo(userid string) (res *GetUserInfoResponseResult, err error) {
	accessToken, err := d.oauth2.GetAccessToken()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	body := make(map[string]interface{})
	body["userid"] = userid
	body["language"] = "zh_CN"
	query := map[string]*string{
		"access_token": tea.String(accessToken),
	}
	resp, err := utils.DoRquest(http.MethodPost, "/topapi/v2/user/get", query, body)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	defer resp.Body.Close()
	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	var result GetUserInfoResponse
	err = json.Unmarshal(bs, &result)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	if result.Errcode != 0 {
		err = xerrors.Errorf("%s", result.Errmsg)
		return
	}
	res = result.Result
	return
}

type GetUserInfoResponse struct {
	utils.DintalkResponse
	Result *GetUserInfoResponseResult `json:"result"`
}

type GetUserInfoResponseResult struct {
	Userid           string                                   `json:"userid"`
	Unionid          string                                   `json:"unionid"`
	Name             string                                   `json:"name"`
	Avatar           string                                   `json:"avatar"`
	StateCode        string                                   `json:"state_code"`
	ManagerUserid    string                                   `json:"manager_userid"`
	Mobile           string                                   `json:"mobile"`
	HideMobile       bool                                     `json:"hide_mobile"`
	Telephone        string                                   `json:"telephone"`
	JobNumber        string                                   `json:"job_number"`
	Title            string                                   `json:"title"`
	Email            string                                   `json:"email"`
	WorkPlace        string                                   `json:"work_place"`
	Remark           string                                   `json:"remark"`
	ExclusiveAccount string                                   `json:"exclusive_account"`
	OrgEmail         string                                   `json:"org_email"`
	DeptIdList       []int                                    `json:"dept_id_list"`
	DeptOrderList    []*GetUserInfoResponseResultDeptOrder    `json:"dept_order_list"`
	Extension        string                                   `json:"extension"`
	HiredDate        string                                   `json:"hired_date"`
	Active           bool                                     `json:"active"`
	RealAuthed       bool                                     `json:"real_authed"`
	OrgEmailType     string                                   `json:"org_email_type"`
	Senior           bool                                     `json:"senior"`
	Admin            bool                                     `json:"admin"`
	Boss             bool                                     `json:"boss"`
	LeaderInDept     []*GetUserInfoResponseResultLeaderInDept `json:"leader_in_dept"`
	RoleList         []*GetUserInfoResponseResultRoleList     `json:"role_list"`
	UnionEmpExt      []*GetUserInfoResponseResultUnionEmpExt  `json:"union_emp_ext"`
}

type GetUserInfoResponseResultDeptOrder struct {
	DeptId int `json:"dept_id"`
	Order  int `json:"order"`
}

type GetUserInfoResponseResultLeaderInDept struct {
	DeptId int  `json:"dept_id"`
	Leader bool `json:"leader"`
}

type GetUserInfoResponseResultRoleList struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	GroupName string `json:"group_name"`
}

type GetUserInfoResponseResultUnionEmpExt struct {
	Userid          string                                  `json:"userid"`
	UnionEmpMapList []*GetUserInfoResponseResultUnionEmpMap `json:"union_emp_map_list"`
	CorpId          string                                  `json:"corp_id"`
}

type GetUserInfoResponseResultUnionEmpMap struct {
	Userid string `json:"userid"`
	CorpId string `json:"corp_id"`
}
