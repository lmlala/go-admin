package service

import (
	"errors"
	"go-admin/app/ops/models"
	"go-admin/app/ops/service/dto"
	"go-admin/common/actions"

	cDto "go-admin/common/dto"

	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk/service"
)

type OpsOnPremise struct {
	service.Service
}

// GetPage 获取SysUser列表
func (e *OpsOnPremise) GetPage(c *dto.OpsOnPremiseGetPageReq, p *actions.DataPermission, list *[]models.OpsOnPremise, count *int64) error {
	var err error
	var data models.OpsOnPremise

	// err = e.Orm.Debug().Preload("Dept").
	err = e.Orm.Debug().
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

// // Get 获取SysUser对象
// func (e *SysUser) Get(d *dto.SysUserById, p *actions.DataPermission, model *models.SysUser) error {
// 	var data models.SysUser

// 	err := e.Orm.Model(&data).Debug().
// 		Scopes(
// 			actions.Permission(data.TableName(), p),
// 		).
// 		First(model, d.GetId()).Error
// 	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
// 		err = errors.New("查看对象不存在或无权查看")
// 		e.Log.Errorf("db error: %s", err)
// 		return err
// 	}
// 	if err != nil {
// 		e.Log.Errorf("db error: %s", err)
// 		return err
// 	}
// 	return nil
// }

// Insert 创建OnPremise对象
func (e *OpsOnPremise) Insert(c *dto.OnPremiseInsertReq) error {
	var err error
	var data models.OpsOnPremise
	var i int64
	err = e.Orm.Model(&data).Where("name = ?", c.Name).Count(&i).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}

	if i > 0 {
		err := errors.New("私有化项目名称已存在！")
		e.Log.Errorf("db error: %s", err)
		return err
	}
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

// Update 修改OpsOnPremise对象
func (e *OpsOnPremise) Update(c *dto.OnPremiseUpdateReq, p *actions.DataPermission) error {
	var err error
	var model models.OpsOnPremise
	db := e.Orm.Scopes(
		actions.Permission(model.TableName(), p),
	).First(&model, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("Service UpdateOpsOnPremise error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")

	}
	c.Generate(&model)
	update := e.Orm.Model(&model).Where("on_premise_id = ?", &model.OnPremiseId).Updates(&model)
	if err = update.Error; err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if update.RowsAffected == 0 {
		err = errors.New("update OnPremise error")
		log.Warnf("db update error")
		return err
	}
	return nil
}

// // UpdateAvatar 更新用户头像
// func (e *SysUser) UpdateAvatar(c *dto.UpdateSysUserAvatarReq, p *actions.DataPermission) error {
// 	var err error
// 	var model models.SysUser
// 	db := e.Orm.Scopes(
// 		actions.Permission(model.TableName(), p),
// 	).First(&model, c.GetId())
// 	if err = db.Error; err != nil {
// 		e.Log.Errorf("Service UpdateSysUser error: %s", err)
// 		return err
// 	}
// 	if db.RowsAffected == 0 {
// 		return errors.New("无权更新该数据")

// 	}
// 	err = e.Orm.Table(model.TableName()).Where("user_id =? ", c.UserId).Updates(c).Error
// 	if err != nil {
// 		e.Log.Errorf("Service UpdateSysUser error: %s", err)
// 		return err
// 	}
// 	return nil
// }

// // UpdateStatus 更新用户状态
// func (e *SysUser) UpdateStatus(c *dto.UpdateSysUserStatusReq, p *actions.DataPermission) error {
// 	var err error
// 	var model models.SysUser
// 	db := e.Orm.Scopes(
// 		actions.Permission(model.TableName(), p),
// 	).First(&model, c.GetId())
// 	if err = db.Error; err != nil {
// 		e.Log.Errorf("Service UpdateSysUser error: %s", err)
// 		return err
// 	}
// 	if db.RowsAffected == 0 {
// 		return errors.New("无权更新该数据")

// 	}
// 	err = e.Orm.Table(model.TableName()).Where("user_id =? ", c.UserId).Updates(c).Error
// 	if err != nil {
// 		e.Log.Errorf("Service UpdateSysUser error: %s", err)
// 		return err
// 	}
// 	return nil
// }

// // ResetPwd 重置用户密码
// func (e *SysUser) ResetPwd(c *dto.ResetSysUserPwdReq, p *actions.DataPermission) error {
// 	var err error
// 	var model models.SysUser
// 	db := e.Orm.Scopes(
// 		actions.Permission(model.TableName(), p),
// 	).First(&model, c.GetId())
// 	if err = db.Error; err != nil {
// 		e.Log.Errorf("At Service ResetSysUserPwd error: %s", err)
// 		return err
// 	}
// 	if db.RowsAffected == 0 {
// 		return errors.New("无权更新该数据")
// 	}
// 	c.Generate(&model)
// 	err = e.Orm.Omit("username", "nick_name", "phone", "role_id", "avatar", "sex").Save(&model).Error
// 	if err != nil {
// 		e.Log.Errorf("At Service ResetSysUserPwd error: %s", err)
// 		return err
// 	}
// 	return nil
// }

// Remove 删除OpsOnPremise
func (e *OpsOnPremise) Remove(c *dto.OnPremiseById, p *actions.DataPermission) error {
	var err error
	var data models.OpsOnPremise

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("Error found in  RemoveOpsOnPremise : %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

// // UpdatePwd 修改SysUser对象密码
// func (e *SysUser) UpdatePwd(id int, oldPassword, newPassword string, p *actions.DataPermission) error {
// 	var err error

// 	if newPassword == "" {
// 		return nil
// 	}
// 	c := &models.SysUser{}

// 	err = e.Orm.Model(c).
// 		Scopes(
// 			actions.Permission(c.TableName(), p),
// 		).Select("UserId", "Password", "Salt").
// 		First(c, id).Error
// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return errors.New("无权更新该数据")
// 		}
// 		e.Log.Errorf("db error: %s", err)
// 		return err
// 	}
// 	var ok bool
// 	ok, err = pkg.CompareHashAndPassword(c.Password, oldPassword)
// 	if err != nil {
// 		e.Log.Errorf("CompareHashAndPassword error, %s", err.Error())
// 		return err
// 	}
// 	if !ok {
// 		err = errors.New("incorrect Password")
// 		e.Log.Warnf("user[%d] %s", id, err.Error())
// 		return err
// 	}
// 	c.Password = newPassword
// 	db := e.Orm.Model(c).Where("user_id = ?", id).
// 		Select("Password", "Salt").
// 		Updates(c)
// 	if err = db.Error; err != nil {
// 		e.Log.Errorf("db error: %s", err)
// 		return err
// 	}
// 	if db.RowsAffected == 0 {
// 		err = errors.New("set password error")
// 		log.Warnf("db update error")
// 		return err
// 	}
// 	return nil
// }

// func (e *SysUser) GetProfile(c *dto.SysUserById, user *models.SysUser, roles *[]models.SysRole, posts *[]models.SysPost) error {
// 	err := e.Orm.Preload("Dept").First(user, c.GetId()).Error
// 	if err != nil {
// 		return err
// 	}
// 	err = e.Orm.Find(roles, user.RoleId).Error
// 	if err != nil {
// 		return err
// 	}
// 	err = e.Orm.Find(posts, user.PostIds).Error
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
