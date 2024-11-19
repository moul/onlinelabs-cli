package object

import (
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/scaleway/scaleway-cli/v2/core"
	"github.com/scaleway/scaleway-cli/v2/internal/human"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

type CustomS3ACLGrant struct {
	Grantee    *string
	Permission types.Permission
}

type BucketResponse struct {
	SuccessResult *core.SuccessResult
	BucketInfo    *bucketInfo
}

func bucketResponseMarshalerFunc(i interface{}, opt *human.MarshalOpt) (string, error) {
	resp := i.(BucketResponse)

	messageStr, err := resp.SuccessResult.MarshalHuman()
	if err != nil {
		return "", err
	}
	bucketStr, err := bucketInfoMarshalerFunc(*resp.BucketInfo, opt)
	if err != nil {
		return "", err
	}

	return strings.Join([]string{
		messageStr,
		bucketStr,
	}, "\n"), nil
}

type bucketInfo struct {
	ID               string
	Region           scw.Region
	APIEndpoint      string
	BucketEndpoint   string
	EnableVersioning bool
	Tags             []types.Tag
	ACL              []CustomS3ACLGrant
	Owner            string
}

func bucketInfoMarshalerFunc(i interface{}, opt *human.MarshalOpt) (string, error) {
	// To avoid recursion of human.Marshal we create a dummy type
	type tmp bucketInfo
	info := tmp(i.(bucketInfo))

	opt.Sections = []*human.MarshalSection{
		{
			FieldName:   "Tags",
			HideIfEmpty: true,
		},
		{
			FieldName:   "ACL",
			HideIfEmpty: true,
		},
	}
	str, err := human.Marshal(info, opt)
	if err != nil {
		return "", err
	}
	return str, nil
}

type BucketGetResult struct {
	*bucketInfo
	Size      *scw.Size
	NbObjects *int64
	NbParts   *int64
}

func bucketGetResultMarshalerFunc(i interface{}, opt *human.MarshalOpt) (string, error) {
	type tmp BucketGetResult
	result := tmp(i.(BucketGetResult))
	opt.Sections = []*human.MarshalSection{
		{
			FieldName:   "Tags",
			HideIfEmpty: true,
		},
		{
			FieldName:   "ACL",
			HideIfEmpty: true,
		},
	}
	str, err := human.Marshal(result, opt)
	if err != nil {
		return "", err
	}
	return str, nil
}

type bucketGetArgs struct {
	Region   scw.Region
	Name     string
	WithSize bool `json:"with-size"`
}

func bucketMarshalerFunc(i interface{}, opt *human.MarshalOpt) (string, error) {
	type tmp []types.Bucket
	result := tmp(i.([]types.Bucket))
	opt.Fields = []*human.MarshalFieldOpt{
		{
			FieldName: "Name",
			Label:     "Name",
		},
		{
			FieldName: "CreationDate",
			Label:     "Creation Date",
		},
	}
	str, err := human.Marshal(result, opt)
	if err != nil {
		return "", err
	}
	return str, nil
}
