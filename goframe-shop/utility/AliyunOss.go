package utility

import (
	"context"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func UploadOssFile(ctx context.Context, file *ghttp.UploadFile) (name string, url string, err error) {
	dirPath := "/tmp/"
	fileName, err := file.Save(dirPath, true)
	if err != nil {
		return "", "", nil
	}

	localFile := dirPath + fileName
	bucketName := g.Cfg().MustGet(ctx, "aliyunoss.bucket").String()
	endPoint := g.Cfg().MustGet(ctx, "aliyunoss.endpoint").String()
	accessKeyId := g.Cfg().MustGet(ctx, "aliyunoss.accessKeyId").String()
	accessKeySecret := g.Cfg().MustGet(ctx, "aliyunoss.accessKeySecret").String()
	ossUrl := g.Cfg().MustGet(ctx, "aliyunoss.ossUrl").String()
	client, err := oss.New(endPoint, accessKeyId, accessKeySecret)
	if err != nil {
		return "", "", err
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return "", "", err
	}

	ossDirPath := "goframe_demo/"
	err = bucket.PutObjectFromFile(ossDirPath+fileName, localFile)
	if err != nil {
		return "", "", err
	}

	err = os.Remove(localFile)
	if err != nil {
		return "", "", err
	}

	return file.Filename, ossUrl + ossDirPath + fileName, nil
}
