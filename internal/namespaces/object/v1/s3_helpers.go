//go:build darwin || linux || windows

package object

import (
	"context"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/scaleway/scaleway-cli/v2/core"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func newS3Client(ctx context.Context, region scw.Region) *s3.Client {
	httpClient := core.ExtractHTTPClient(ctx)
	scwClient := core.ExtractClient(ctx)
	accessKey, ok := scwClient.GetAccessKey()
	if !ok {
		return nil
	}
	secretKey, ok := scwClient.GetSecretKey()
	if !ok {
		return nil
	}

	var customEndpoint string
	if ep := os.Getenv("SCW_S3_ENDPOINT"); ep != "" {
		customEndpoint = ep
	} else {
		customEndpoint = "https://s3." + region.String() + ".scw.cloud"
	}

	return s3.New(s3.Options{
		APIOptions:    nil,
		ClientLogMode: 0,
		Credentials: aws.CredentialsProviderFunc(func(_ context.Context) (aws.Credentials, error) {
			return aws.Credentials{
				AccessKeyID:     accessKey,
				SecretAccessKey: secretKey,
			}, nil
		}),
		BaseEndpoint: scw.StringPtr(customEndpoint),
		Region:       region.String(),
		HTTPClient:   httpClient,
	})
}

// Caching BucketCannedACL values for shell completion
var completeBucketACLCache []types.BucketCannedACL

func autocompleteBucketACL(_ context.Context, prefix string, _ any) core.AutocompleteSuggestions {
	suggestions := core.AutocompleteSuggestions(nil)

	if len(completeBucketACLCache) == 0 {
		var awsCannedACL types.BucketCannedACL
		completeBucketACLCache = awsCannedACL.Values()
	}

	for _, acl := range completeBucketACLCache {
		if strings.HasPrefix(string(acl), prefix) {
			suggestions = append(suggestions, string(acl))
		}
	}

	return suggestions
}

func verifyACLInput(aclInput string) (bool, []types.BucketCannedACL) {
	var awsCannedACL types.BucketCannedACL
	possibleValues := awsCannedACL.Values()

	for _, possibleValue := range possibleValues {
		if string(possibleValue) == aclInput {
			return true, nil
		}
	}
	return false, possibleValues
}

func awsACLToCustomGrants(output *s3.GetBucketAclOutput) []CustomS3ACLGrant {
	customGrants := []CustomS3ACLGrant(nil)
	for _, grant := range output.Grants {
		var grantee *string
		switch grant.Grantee.Type {
		case types.TypeCanonicalUser:
			grantee = scw.StringPtr(normalizeOwnerID(grant.Grantee.ID))
		case types.TypeGroup:
			split := strings.Split(*grant.Grantee.URI, "/")
			grantee = scw.StringPtr(split[len(split)-1])
		}
		customGrants = append(customGrants, CustomS3ACLGrant{
			Grantee:    grantee,
			Permission: grant.Permission,
		})
	}
	return customGrants
}

func normalizeOwnerID(id *string) string {
	if id == nil {
		return ""
	}
	tab := strings.Split(*id, ":")
	if len(tab) != 2 {
		return ""
	}
	return tab[0]
}
