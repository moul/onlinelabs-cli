//go:build darwin || linux || windows

package object

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/scaleway/scaleway-cli/v2/core"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

type bucketConfigArgs struct {
	Region           scw.Region
	Name             string
	Tags             []string
	EnableVersioning bool `json:"enable-versioning"`
	ACL              string
}

func bucketCreateCommand() *core.Command {
	return &core.Command{
		Namespace: "object",
		Resource:  "bucket",
		Verb:      "create",
		Short:     "Create an S3 bucket",
		Long:      "Create an Object Storage Bucket with the S3 protocol. The namespace is shared between all S3 users, so its name must be unique.",
		ArgsType:  reflect.TypeOf(bucketConfigArgs{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:             "name",
				Positional:       true,
				Required:         true,
				Short:            "The unique name of the bucket",
				AutoCompleteFunc: autocompleteBucketName,
			},
			{
				Name:       "tags.{index}",
				Positional: false,
				Required:   false,
				Short:      "List of tags to set on the bucket",
			},
			{
				Name:       "enable-versioning",
				Positional: false,
				Required:   false,
				Default:    core.DefaultValueSetter("false"),
				Short:      "Whether or not objects in the bucket should have multiple versions",
			},
			{
				Name:             "acl",
				Positional:       false,
				Required:         false,
				Default:          core.DefaultValueSetter("private"),
				Short:            "The permissions given to users (grantees) to read or write objects",
				AutoCompleteFunc: autocompleteBucketACL,
			},
			core.RegionArgSpec(),
		},
		Run: func(ctx context.Context, argsI interface{}) (interface{}, error) {
			args := argsI.(*bucketConfigArgs)
			client := newS3Client(ctx, args.Region)

			if ok, possibleValues := verifyACLInput(args.ACL); !ok {
				return nil, fmt.Errorf("ACL field must be one of %v", possibleValues)
			}

			_, err := client.CreateBucket(ctx, &s3.CreateBucketInput{
				Bucket: &args.Name,
				ACL:    types.BucketCannedACL(args.ACL),
				CreateBucketConfiguration: &types.CreateBucketConfiguration{
					LocationConstraint: types.BucketLocationConstraint(args.Region),
				},
				ObjectLockEnabledForBucket: scw.BoolPtr(false),
			})
			if err != nil {
				return nil, fmt.Errorf("could not create bucket: %w", err)
			}

			err = putBucketVersioning(ctx, client, args.Name, args.EnableVersioning)
			if err != nil {
				return nil, fmt.Errorf("could not put bucket versioning: %w", err)
			}
			err = putBucketTagging(ctx, client, args.Name, args.Tags)
			if err != nil {
				return nil, fmt.Errorf("could not put bucket tags: %w", err)
			}

			bucket, err := getBucketInfo(ctx, args.Region, args.Name)
			if err != nil {
				return nil, fmt.Errorf("could not get bucket's information: %w", err)
			}

			return &BucketResponse{
				BucketInfo: bucket,
				SuccessResult: &core.SuccessResult{
					Resource: "bucket",
					Verb:     "create",
				},
			}, nil
		},
	}
}

type bucketDeleteArgs struct {
	Region scw.Region
	Name   string
}

func bucketDeleteCommand() *core.Command {
	return &core.Command{
		Namespace: "object",
		Resource:  "bucket",
		Verb:      "delete",
		Short:     "Delete an S3 bucket",
		Long:      "Delete an S3 bucket with all its content.",
		ArgsType:  reflect.TypeOf(bucketDeleteArgs{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:             "name",
				Positional:       true,
				Required:         true,
				Short:            "The unique name of the bucket",
				AutoCompleteFunc: autocompleteBucketName,
			},
			core.RegionArgSpec(),
		},
		Run: func(ctx context.Context, argsI interface{}) (interface{}, error) {
			args := argsI.(*bucketDeleteArgs)
			client := newS3Client(ctx, args.Region)

			_, err := client.DeleteBucket(ctx, &s3.DeleteBucketInput{
				Bucket: &args.Name,
			})
			if err != nil {
				return nil, err
			}

			return &core.SuccessResult{
				Resource: "bucket",
				Verb:     "delete",
			}, nil
		},
	}
}

func bucketGetCommand() *core.Command {
	return &core.Command{
		Namespace: "object",
		Resource:  "bucket",
		Verb:      "get",
		Short:     "Get information about an S3 bucket",
		Long:      "Get the properties of an S3 bucket like tags, endpoint, access control, versioning, size, etc.",
		ArgsType:  reflect.TypeOf(bucketGetArgs{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:             "name",
				Positional:       true,
				Required:         true,
				Short:            "The unique name of the bucket",
				AutoCompleteFunc: autocompleteBucketName,
			},
			{
				Name:       "with-size",
				Positional: false,
				Required:   false,
				Default:    core.DefaultValueSetter("false"),
				Short:      "Whether to return the total size of the bucket and the number of objects. This operation can take long for large buckets.",
			},
			core.RegionArgSpec(),
		},
		Run: func(ctx context.Context, argsI interface{}) (interface{}, error) {
			args := argsI.(*bucketGetArgs)
			client := newS3Client(ctx, args.Region)

			bucket, err := getBucketInfo(ctx, args.Region, args.Name)
			if err != nil {
				return nil, fmt.Errorf("could not get bucket's information: %w", err)
			}

			if !args.WithSize {
				return BucketGetResult{
					bucket,
					nil,
					nil,
					nil,
				}, nil
			}

			nbObjects, nbParts, totalSize, err := countBucketObjects(ctx, client, args.Name)
			if err != nil {
				return nil, fmt.Errorf("could not get bucket size: %w", err)
			}

			return BucketGetResult{
				bucket,
				&totalSize,
				&nbObjects,
				&nbParts,
			}, nil
		},
	}
}

type bucketListArgs struct {
	Region scw.Region
}

func bucketListCommand() *core.Command {
	return &core.Command{
		Namespace: "object",
		Resource:  "bucket",
		Verb:      "list",
		Short:     "List S3 buckets",
		Long:      "List all existing S3 buckets in the specified region",
		ArgsType:  reflect.TypeOf(bucketListArgs{}),
		ArgSpecs: core.ArgSpecs{
			core.RegionArgSpec(),
		},
		Run: func(ctx context.Context, argsI interface{}) (interface{}, error) {
			args := argsI.(*bucketListArgs)
			client := newS3Client(ctx, args.Region)

			buckets, err := client.ListBuckets(ctx, &s3.ListBucketsInput{})
			if err != nil {
				return nil, fmt.Errorf("could not list buckets: %w", err)
			}
			return buckets.Buckets, nil
		},
	}
}

func bucketUpdateCommand() *core.Command {
	return &core.Command{
		Namespace: "object",
		Resource:  "bucket",
		Verb:      "update",
		Short:     "Update an S3 bucket",
		Long:      "Update an S3 bucket's properties like tags, access control and versioning.",
		ArgsType:  reflect.TypeOf(bucketConfigArgs{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:             "name",
				Positional:       true,
				Required:         true,
				Short:            "The unique name of the bucket",
				AutoCompleteFunc: autocompleteBucketName,
			},
			{
				Name:       "tags.{index}",
				Positional: false,
				Required:   false,
				Short:      "List of new tags to set on the bucket",
			},
			{
				Name:       "enable-versioning",
				Positional: false,
				Required:   false,
				Default:    core.DefaultValueSetter("false"),
				Short:      "Whether or not objects in the bucket should have multiple versions",
			},
			{
				Name:             "acl",
				Positional:       false,
				Required:         false,
				Default:          core.DefaultValueSetter("private"),
				Short:            "The permissions given to users (grantees) to read or write objects",
				AutoCompleteFunc: autocompleteBucketACL,
			},
			core.RegionArgSpec(),
		},
		Run: func(ctx context.Context, argsI interface{}) (interface{}, error) {
			args := argsI.(*bucketConfigArgs)
			client := newS3Client(ctx, args.Region)

			err := putBucketVersioning(ctx, client, args.Name, args.EnableVersioning)
			if err != nil {
				return nil, fmt.Errorf("could not update bucket versioning: %w", err)
			}

			err = putBucketTagging(ctx, client, args.Name, args.Tags)
			if err != nil {
				return nil, fmt.Errorf("could not update bucket tags: %w", err)
			}

			err = putBucketACL(ctx, client, args.Name, args.ACL)
			if err != nil {
				return nil, fmt.Errorf("could not update bucket ACL: %w", err)
			}

			bucketResponse, err := getBucketInfo(ctx, args.Region, args.Name)
			if err != nil {
				return nil, fmt.Errorf("could not get bucket's information: %w", err)
			}

			return &BucketResponse{
				BucketInfo: bucketResponse,
				SuccessResult: &core.SuccessResult{
					Resource: "bucket",
					Verb:     "update",
				},
			}, nil
		},
	}
}

func getBucketInfo(ctx context.Context, region scw.Region, name string) (*bucketInfo, error) {
	client := newS3Client(ctx, region)
	bucket := &bucketInfo{
		ID:     name,
		Region: region,
	}

	// get versioning
	versioningOutput, err := client.GetBucketVersioning(ctx, &s3.GetBucketVersioningInput{
		Bucket: &name,
	})
	if err != nil {
		return nil, fmt.Errorf("could not get bucket versioning: %w", err)
	}
	switch versioningOutput.Status {
	case types.BucketVersioningStatusSuspended, "":
		bucket.EnableVersioning = false
	case types.BucketVersioningStatusEnabled:
		bucket.EnableVersioning = true
	}

	// get tagging
	tagging, err := client.GetBucketTagging(ctx, &s3.GetBucketTaggingInput{
		Bucket: &name,
	})
	if err != nil && !strings.Contains(err.Error(), "NoSuchTagSet") {
		return nil, fmt.Errorf("could not get bucket tagging: %w", err)
	} else if tagging != nil {
		bucket.Tags = tagging.TagSet
	}

	// get ACL
	acl, err := client.GetBucketAcl(ctx, &s3.GetBucketAclInput{
		Bucket: &name,
	})
	if err != nil {
		return nil, fmt.Errorf("could not get bucket ACL: %w", err)
	}
	bucket.Owner = normalizeOwnerID(acl.Owner.ID)
	bucket.ACL = awsACLToCustomGrants(acl)

	// get endpoints
	bucket.APIEndpoint = getAPIEndpoint(region.String())
	bucket.BucketEndpoint, err = getBucketEndpoint(name, region.String())
	if err != nil {
		return nil, err
	}

	return bucket, nil
}

func getAPIEndpoint(region string) string {
	if customEndpoint := os.Getenv("SCW_S3_ENDPOINT"); customEndpoint != "" {
		return customEndpoint
	}
	return fmt.Sprintf("https://s3.%s.scw.cloud", region)
}

func getBucketEndpoint(name, region string) (string, error) {
	if customEndpoint := os.Getenv("SCW_S3_ENDPOINT"); customEndpoint != "" {
		u, err := url.Parse(customEndpoint)
		if err != nil {
			return "", fmt.Errorf("could not parse custom endpoint %s: %w", customEndpoint, err)
		}
		u = u.JoinPath(name, u.Path)
		return u.String(), nil
	}
	return fmt.Sprintf("https://%s.s3.%s.scw.cloud", name, region), nil
}

func countBucketObjects(ctx context.Context, client *s3.Client, name string) (nbObjects, nbParts int64, totalSize scw.Size, err error) {
	var size int64

	// count full objects
	objectsList, err := client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: &name,
	})
	if err != nil {
		return nbObjects, nbParts, totalSize, fmt.Errorf("could not list objects: %w", err)
	}
	for _, object := range objectsList.Contents {
		nbObjects++
		size += *object.Size
	}

	// count parts
	multipartUploads, err := client.ListMultipartUploads(ctx, &s3.ListMultipartUploadsInput{
		Bucket: &name,
	})
	if err != nil {
		return nbObjects, nbParts, totalSize, fmt.Errorf("could not list multipart uploads: %w", err)
	}
	for _, upload := range multipartUploads.Uploads {
		partsList, err := client.ListParts(ctx, &s3.ListPartsInput{
			Bucket:   &name,
			Key:      upload.Key,
			UploadId: upload.UploadId,
		})
		if err != nil {
			return nbObjects, nbParts, totalSize, fmt.Errorf("could not list parts: %w", err)
		}
		for _, part := range partsList.Parts {
			nbParts++
			size += *part.Size
		}
	}

	return nbObjects, nbParts, scw.Size(size), nil
}

func putBucketVersioning(ctx context.Context, client *s3.Client, name string, enabled bool) error {
	request := &s3.PutBucketVersioningInput{
		Bucket: &name,
	}
	if enabled {
		request.VersioningConfiguration = &types.VersioningConfiguration{
			Status: types.BucketVersioningStatusEnabled,
		}
	} else {
		request.VersioningConfiguration = &types.VersioningConfiguration{
			Status: types.BucketVersioningStatusSuspended,
		}
	}
	_, err := client.PutBucketVersioning(ctx, request)
	if err != nil {
		return err
	}
	return nil
}

func putBucketTagging(ctx context.Context, client *s3.Client, name string, tags []string) error {
	if len(tags) == 0 {
		return nil
	}
	newTags := []types.Tag(nil)
	for _, tag := range tags {
		trim := strings.Trim(tag, "\"")
		split := strings.Split(trim, "=")
		newTags = append(newTags, types.Tag{
			Key:   &split[0],
			Value: &split[1],
		})
	}
	_, err := client.PutBucketTagging(ctx, &s3.PutBucketTaggingInput{
		Bucket:  &name,
		Tagging: &types.Tagging{TagSet: newTags},
	})
	if err != nil {
		return err
	}
	return nil
}

func putBucketACL(ctx context.Context, client *s3.Client, name string, acl string) error {
	_, err := client.PutBucketAcl(ctx, &s3.PutBucketAclInput{
		Bucket: &name,
		ACL:    types.BucketCannedACL(acl),
	})
	if err != nil {
		return err
	}
	return nil
}

// Caching ListBuckets response for shell completion
var completeListBucketsCache []types.Bucket

func autocompleteBucketName(ctx context.Context, prefix string, request any) core.AutocompleteSuggestions {
	var region scw.Region
	switch t := request.(type) {
	case bucketConfigArgs:
		region = t.Region
	case bucketDeleteArgs:
		region = t.Region
	case bucketGetArgs:
		region = t.Region
	}

	suggestions := core.AutocompleteSuggestions(nil)
	client := newS3Client(ctx, region)

	if completeListBucketsCache == nil {
		buckets, err := client.ListBuckets(ctx, &s3.ListBucketsInput{})
		if err != nil {
			return nil
		}
		completeListBucketsCache = buckets.Buckets
	}

	for _, bucket := range completeListBucketsCache {
		if strings.HasPrefix(*bucket.Name, prefix) {
			suggestions = append(suggestions, *bucket.Name)
		}
	}

	return suggestions
}
