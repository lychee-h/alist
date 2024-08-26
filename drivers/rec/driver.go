package _rec

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/alist-org/alist/v3/drivers/base"
	"github.com/alist-org/alist/v3/internal/driver"
	"github.com/alist-org/alist/v3/internal/errs"
	"github.com/alist-org/alist/v3/internal/model"
	"github.com/alist-org/alist/v3/pkg/utils"
	"github.com/go-resty/resty/v2"
)

type RecCloud struct {
	model.Storage
	Addition
	client *resty.Client
}

func (d *RecCloud) Config() driver.Config {
	return config
}

func (d *RecCloud) GetAddition() driver.Additional {
	return &d.Addition
}

func (d *RecCloud) Init(ctx context.Context) error {
	// TODO login / refresh token
	//op.MustSaveDriverStorage(d)
	body := map[string]string{
		"username":    d.Username,
		"password":    d.Password,
		"resultInput": d.ResultInput,
	}
	res, err := d.client.R().SetBody(body).Post("http://127.0.0.1/token")
	if err != nil {
		return err
	}
	respFromPy := RespFromPy{}
	json.Unmarshal(res.Body(), &respFromPy)
	d.client = base.NewRestyClient().
		SetHeader("x-auth-token", respFromPy.Token)
	return nil
}

func (d *RecCloud) Drop(ctx context.Context) error {
	return nil
}

func (d *RecCloud) List(ctx context.Context, dir model.Obj, args model.ListArgs) ([]model.Obj, error) {
	// TODO return the files list, required
	res := make([]model.Obj, 0)
	var listResp ListResponse
	resp, err := d.client.R().SetQueryParam("disk_type", "cloud").
		SetQueryParam("is_rec", "false").
		SetQueryParam("category", "all").
		SetQueryParam("group_number", d.GroupNumber).
		Get("https://recapi.ustc.edu.cn/api/v2/folder/content/" + dir.GetID())
	if err != nil {
		return nil, errs.NotSupport
	}
	err = json.Unmarshal(resp.Body(), &listResp)
	if err != nil {
		return nil, errs.NotSupport
	}
	for _, file := range listResp.Entity.Datas {
		if file.Type == "folder" {
			lastOpTime := utils.MustParseCNTime(file.LastUpdateDate)
			res = append(res, &model.Object{
				ID:       file.Number,
				Name:     file.Name,
				Modified: lastOpTime,
				IsFolder: true,
			})

		} else if file.Type == "file" {
			lastOpTime := utils.MustParseCNTime(file.LastUpdateDate)
			res = append(res, &model.ObjThumb{
				Object: model.Object{
					ID:       file.Number,
					Name:     file.Name,
					Modified: lastOpTime,
					Size:     file.Bytes,
				},
				Thumbnail: model.Thumbnail{},
			})
		}
	}
	return res, nil
}

func (d *RecCloud) Link(ctx context.Context, file model.Obj, args model.LinkArgs) (*model.Link, error) {
	// TODO return link of file, required
	// 拿到下载链接 post请求
	resp, err := d.client.R().SetBody(DownLoadBody{
		FilesList:   []string{file.GetID()},
		GroupNumber: d.GroupNumber,
	}).Post("https://recapi.ustc.edu.cn/api/v2/download")
	if err != nil {
		return nil, errs.ObjectNotFound
	}

	var downloadResp downloadResponse
	err = json.Unmarshal(resp.Body(), &downloadResp)
	if err != nil {
		return nil, errs.NotSupport
	}
	// 去除转义字符
	unquotedURL := strings.ReplaceAll(downloadResp.Entity[file.GetID()], `\/`, `/`)

	// 返回下载链接
	res := &model.Link{URL: unquotedURL}
	return res, nil
	// return nil, errs.NotImplement
}

func (d *RecCloud) MakeDir(ctx context.Context, parentDir model.Obj, dirName string) (model.Obj, error) {
	// TODO create folder, optional
	return nil, errs.NotImplement
}

func (d *RecCloud) Move(ctx context.Context, srcObj, dstDir model.Obj) (model.Obj, error) {
	// TODO move obj, optional
	return nil, errs.NotImplement
}

func (d *RecCloud) Rename(ctx context.Context, srcObj model.Obj, newName string) (model.Obj, error) {
	// TODO rename obj, optional
	return nil, errs.NotImplement
}

func (d *RecCloud) Copy(ctx context.Context, srcObj, dstDir model.Obj) (model.Obj, error) {
	// TODO copy obj, optional
	return nil, errs.NotImplement
}

func (d *RecCloud) Remove(ctx context.Context, obj model.Obj) error {
	// TODO remove obj, optional
	return errs.NotImplement
}

func (d *RecCloud) Put(ctx context.Context, dstDir model.Obj, stream model.FileStreamer, up driver.UpdateProgress) (model.Obj, error) {
	// TODO upload file, optional
	return nil, errs.NotImplement
}

//func (d *Template) Other(ctx context.Context, args model.OtherArgs) (interface{}, error) {
//	return nil, errs.NotSupport
//}

var _ driver.Driver = (*RecCloud)(nil)
