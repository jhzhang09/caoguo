package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _WxUserinfoMgr struct {
	*_BaseMgr
}

// WxUserinfoMgr open func
func WxUserinfoMgr(db *gorm.DB) *_WxUserinfoMgr {
	if db == nil {
		panic(fmt.Errorf("WxUserinfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WxUserinfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("wx_userinfo"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WxUserinfoMgr) GetTableName() string {
	return "wx_userinfo"
}

// Get 获取
func (obj *_WxUserinfoMgr) Get() (result WxUserinfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WxUserinfoMgr) Gets() (results []*WxUserinfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_WxUserinfoMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOpenid openid获取 微信用户唯一标识符
//
//微信用户唯一标识符
func (obj *_WxUserinfoMgr) WithOpenid(openid string) Option {
	return optionFunc(func(o *options) { o.query["openid"] = openid })
}

// WithNickName nick_name获取 用户昵称
func (obj *_WxUserinfoMgr) WithNickName(nickName string) Option {
	return optionFunc(func(o *options) { o.query["nick_name"] = nickName })
}

// WithAvatarURL avatar_url获取 用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。
func (obj *_WxUserinfoMgr) WithAvatarURL(avatarURL string) Option {
	return optionFunc(func(o *options) { o.query["avatar_url"] = avatarURL })
}

// WithGender gender获取 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
func (obj *_WxUserinfoMgr) WithGender(gender int) Option {
	return optionFunc(func(o *options) { o.query["gender"] = gender })
}

// WithCity city获取 用户所在城市
func (obj *_WxUserinfoMgr) WithCity(city string) Option {
	return optionFunc(func(o *options) { o.query["city"] = city })
}

// WithProvince province获取 用户所在省份
func (obj *_WxUserinfoMgr) WithProvince(province string) Option {
	return optionFunc(func(o *options) { o.query["province"] = province })
}

// WithCountry country获取 用户所在国家
func (obj *_WxUserinfoMgr) WithCountry(country string) Option {
	return optionFunc(func(o *options) { o.query["country"] = country })
}

// WithLanguage language获取 用户的语言，简体中文为zh_CN
func (obj *_WxUserinfoMgr) WithLanguage(language []byte) Option {
	return optionFunc(func(o *options) { o.query["language"] = language })
}

// WithPhone phone获取 用户绑定的手机号
func (obj *_WxUserinfoMgr) WithPhone(phone string) Option {
	return optionFunc(func(o *options) { o.query["phone"] = phone })
}

// GetByOption 功能选项模式获取
func (obj *_WxUserinfoMgr) GetByOption(opts ...Option) (result WxUserinfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WxUserinfoMgr) GetByOptions(opts ...Option) (results []*WxUserinfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_WxUserinfoMgr) GetFromID(id int) (result WxUserinfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_WxUserinfoMgr) GetBatchFromID(ids []int) (results []*WxUserinfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromOpenid 通过openid获取内容 微信用户唯一标识符
//
//微信用户唯一标识符
func (obj *_WxUserinfoMgr) GetFromOpenid(openid string) (result WxUserinfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("openid = ?", openid).Find(&result).Error

	return
}

// GetBatchFromOpenid 批量唯一主键查找 微信用户唯一标识符
//
//微信用户唯一标识符
func (obj *_WxUserinfoMgr) GetBatchFromOpenid(openids []string) (results []*WxUserinfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("openid IN (?)", openids).Find(&results).Error

	return
}

// GetFromNickName 通过nick_name获取内容 用户昵称
func (obj *_WxUserinfoMgr) GetFromNickName(nickName string) (results []*WxUserinfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nick_name = ?", nickName).Find(&results).Error

	return
}

// GetBatchFromNickName 批量唯一主键查找 用户昵称
func (obj *_WxUserinfoMgr) GetBatchFromNickName(nickNames []string) (results []*WxUserinfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nick_name IN (?)", nickNames).Find(&results).Error

	return
}

// GetFromAvatarURL 通过avatar_url获取内容 用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。
func (obj *_WxUserinfoMgr) GetFromAvatarURL(avatarURL string) (results []*WxUserinfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("avatar_url = ?", avatarURL).Find(&results).Error

	return
}

// GetBatchFromAvatarURL 批量唯一主键查找 用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。
func (obj *_WxUserinfoMgr) GetBatchFromAvatarURL(avatarURLs []string) (results []*WxUserinfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("avatar_url IN (?)", avatarURLs).Find(&results).Error

	return
}

// GetFromGender 通过gender获取内容 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
func (obj *_WxUserinfoMgr) GetFromGender(gender int) (results []*WxUserinfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gender = ?", gender).Find(&results).Error

	return
}

// GetBatchFromGender 批量唯一主键查找 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
func (obj *_WxUserinfoMgr) GetBatchFromGender(genders []int) (results []*WxUserinfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gender IN (?)", genders).Find(&results).Error

	return
}

// GetFromCity 通过city获取内容 用户所在城市
func (obj *_WxUserinfoMgr) GetFromCity(city string) (results []*WxUserinfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("city = ?", city).Find(&results).Error

	return
}

// GetBatchFromCity 批量唯一主键查找 用户所在城市
func (obj *_WxUserinfoMgr) GetBatchFromCity(citys []string) (results []*WxUserinfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("city IN (?)", citys).Find(&results).Error

	return
}

// GetFromProvince 通过province获取内容 用户所在省份
func (obj *_WxUserinfoMgr) GetFromProvince(province string) (results []*WxUserinfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("province = ?", province).Find(&results).Error

	return
}

// GetBatchFromProvince 批量唯一主键查找 用户所在省份
func (obj *_WxUserinfoMgr) GetBatchFromProvince(provinces []string) (results []*WxUserinfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("province IN (?)", provinces).Find(&results).Error

	return
}

// GetFromCountry 通过country获取内容 用户所在国家
func (obj *_WxUserinfoMgr) GetFromCountry(country string) (results []*WxUserinfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("country = ?", country).Find(&results).Error

	return
}

// GetBatchFromCountry 批量唯一主键查找 用户所在国家
func (obj *_WxUserinfoMgr) GetBatchFromCountry(countrys []string) (results []*WxUserinfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("country IN (?)", countrys).Find(&results).Error

	return
}

// GetFromLanguage 通过language获取内容 用户的语言，简体中文为zh_CN
func (obj *_WxUserinfoMgr) GetFromLanguage(language []byte) (results []*WxUserinfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("language = ?", language).Find(&results).Error

	return
}

// GetBatchFromLanguage 批量唯一主键查找 用户的语言，简体中文为zh_CN
func (obj *_WxUserinfoMgr) GetBatchFromLanguage(languages [][]byte) (results []*WxUserinfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("language IN (?)", languages).Find(&results).Error

	return
}

// GetFromPhone 通过phone获取内容 用户绑定的手机号
func (obj *_WxUserinfoMgr) GetFromPhone(phone string) (results []*WxUserinfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("phone = ?", phone).Find(&results).Error

	return
}

// GetBatchFromPhone 批量唯一主键查找 用户绑定的手机号
func (obj *_WxUserinfoMgr) GetBatchFromPhone(phones []string) (results []*WxUserinfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("phone IN (?)", phones).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_WxUserinfoMgr) FetchByPrimaryKey(id int) (result WxUserinfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueByOpenid primay or index 获取唯一内容
func (obj *_WxUserinfoMgr) FetchUniqueByOpenid(openid string) (result WxUserinfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("openid = ?", openid).Find(&result).Error

	return
}
