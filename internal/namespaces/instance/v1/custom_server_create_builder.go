package instance

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/scaleway/scaleway-cli/v2/core"
	block "github.com/scaleway/scaleway-sdk-go/api/block/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/scaleway/scaleway-sdk-go/api/marketplace/v2"
	"github.com/scaleway/scaleway-sdk-go/logger"
	"github.com/scaleway/scaleway-sdk-go/scw"
	"github.com/scaleway/scaleway-sdk-go/validation"
)

type ServerBuilder struct {
	// createdReq is the request being built
	createReq *instance.CreateServerRequest
	// createIPReqs is filled with requests if one or more IP are needed
	createIPReqs []*instance.CreateIPRequest

	// volumes is the list of requested volumes
	volumes []*VolumeBuilder

	// rootVolume is the builder for the root volume
	rootVolume *VolumeBuilder

	// All needed APIs
	apiMarketplace *marketplace.API
	apiInstance    *instance.API
	apiBlock       *block.API

	// serverType is filled with the ServerType if CommercialType is found in the API.
	serverType *instance.ServerType
	// serverImage is filled with the Image if one is provided
	serverImage *instance.Image
}

// NewServerBuilder creates a new builder for a server with requested commercialType in given zone.
// commercialType will be used to validate that added components are supported.
func NewServerBuilder(client *scw.Client, name string, zone scw.Zone, commercialType string) *ServerBuilder {
	sb := &ServerBuilder{
		createReq: &instance.CreateServerRequest{
			Name:           name,
			Zone:           zone,
			CommercialType: commercialType,
		},
		apiMarketplace: marketplace.NewAPI(client),
		apiInstance:    instance.NewAPI(client),
		apiBlock:       block.NewAPI(client),
	}

	sb.serverType = getServerType(sb.apiInstance, sb.createReq.Zone, sb.createReq.CommercialType)

	return sb
}

func (sb *ServerBuilder) AddOrganizationID(orgID *string) *ServerBuilder {
	if orgID != nil {
		sb.createReq.Organization = orgID
	}

	return sb
}

func (sb *ServerBuilder) AddProjectID(projectID *string) *ServerBuilder {
	if projectID != nil {
		sb.createReq.Project = projectID
	}

	return sb
}

func (sb *ServerBuilder) AddEnableIPv6(enableIPv6 *bool) *ServerBuilder {
	sb.createReq.EnableIPv6 = enableIPv6 //nolint: staticcheck

	return sb
}

func (sb *ServerBuilder) AddTags(tags []string) *ServerBuilder {
	if len(tags) > 0 {
		sb.createReq.Tags = tags
	}

	return sb
}

func (sb *ServerBuilder) AddRoutedIPEnabled(routedIPEnabled *bool) *ServerBuilder {
	if routedIPEnabled != nil {
		sb.createReq.RoutedIPEnabled = routedIPEnabled //nolint: staticcheck // Field is deprecated but still supported
	}

	return sb
}

func (sb *ServerBuilder) AddDynamicIPRequired(dynamicIPRequired *bool) *ServerBuilder {
	if dynamicIPRequired != nil {
		sb.createReq.DynamicIPRequired = dynamicIPRequired
	}

	return sb
}

func (sb *ServerBuilder) AddAdminPasswordEncryptionSSHKeyID(adminPasswordEncryptionSSHKeyID *string) *ServerBuilder {
	if adminPasswordEncryptionSSHKeyID != nil {
		sb.createReq.AdminPasswordEncryptionSSHKeyID = adminPasswordEncryptionSSHKeyID
	}

	return sb
}

func (sb *ServerBuilder) isWindows() bool {
	return commercialTypeIsWindowsServer(sb.createReq.CommercialType)
}

func (sb *ServerBuilder) rootVolumeTemplate() *instance.VolumeServerTemplate {
	rootVolume, exists := sb.createReq.Volumes["0"]
	if !exists {
		return nil
	}

	return rootVolume
}

func (sb *ServerBuilder) rootVolumeIsSBS() bool {
	rootVolume := sb.rootVolumeTemplate()
	if rootVolume == nil {
		return false
	}

	return rootVolume.VolumeType == instance.VolumeVolumeTypeSbsVolume
}

func (sb *ServerBuilder) marketplaceImageType() marketplace.LocalImageType {
	if sb.rootVolumeIsSBS() {
		return marketplace.LocalImageTypeInstanceSbs
	}
	return marketplace.LocalImageTypeInstanceLocal
}

// AddImage handle a custom image argument.
// image could be:
//   - A local image UUID.
//   - An image label.
func (sb *ServerBuilder) AddImage(image string) (*ServerBuilder, error) {
	switch {
	case image == "none":
		return sb, nil
	case !validation.IsUUID(image):
		imageLabel := strings.ReplaceAll(image, "-", "_")

		localImage, err := sb.apiMarketplace.GetLocalImageByLabel(&marketplace.GetLocalImageByLabelRequest{
			ImageLabel:     imageLabel,
			Zone:           sb.createReq.Zone,
			CommercialType: sb.createReq.CommercialType,
			Type:           sb.marketplaceImageType(),
		})
		if err != nil {
			return sb, err
		}

		image = localImage.ID
	}

	sb.createReq.Image = &image

	getImageResponse, err := sb.apiInstance.GetImage(&instance.GetImageRequest{
		Zone:    sb.createReq.Zone,
		ImageID: *(sb.createReq.Image),
	})
	if err != nil {
		logger.Warningf("cannot get image %s: %s", sb.createReq.Image, err)
	} else {
		sb.serverImage = getImageResponse.Image
	}

	return sb, nil
}

// AddIP takes an ip argument and change requests accordingly.
// ip could be:
//   - "new"
//   - A flexible IP UUID
//   - A flexible IP address
//   - "dynamic"
//   - "none"
func (sb *ServerBuilder) AddIP(ip string) (*ServerBuilder, error) {
	switch {
	case ip == "" || ip == "new" || ip == "ipv4":
		sb.createIPReqs = []*instance.CreateIPRequest{{
			Zone:    sb.createReq.Zone,
			Project: sb.createReq.Project,
			Type:    instance.IPTypeRoutedIPv4,
		}}
	case ip == "ipv6":
		sb.createIPReqs = []*instance.CreateIPRequest{{
			Zone:    sb.createReq.Zone,
			Project: sb.createReq.Project,
			Type:    instance.IPTypeRoutedIPv6,
		}}
	case ip == "both":
		sb.createIPReqs = []*instance.CreateIPRequest{{
			Zone:    sb.createReq.Zone,
			Project: sb.createReq.Project,
			Type:    instance.IPTypeRoutedIPv4,
		}, {
			Zone:    sb.createReq.Zone,
			Project: sb.createReq.Project,
			Type:    instance.IPTypeRoutedIPv6,
		}}
	case validation.IsUUID(ip):
		sb.createReq.PublicIP = scw.StringPtr(ip)
	case net.ParseIP(ip) != nil:
		logger.Debugf("finding public IP UUID from address: %s", ip)
		res, err := sb.apiInstance.GetIP(&instance.GetIPRequest{
			Zone: sb.createReq.Zone,
			IP:   ip,
		})
		if err != nil { // FIXME: isNotFoundError
			return sb, fmt.Errorf("%s does not belong to you", ip)
		}
		sb.createReq.PublicIP = scw.StringPtr(res.IP.ID)

	case ip == "dynamic":
		sb.createReq.DynamicIPRequired = scw.BoolPtr(true)
	case ip == "none":
		sb.createReq.DynamicIPRequired = scw.BoolPtr(false)
	default:
		return sb, fmt.Errorf(`invalid IP "%s", should be either 'new', 'ipv4', 'ipv6', 'both', 'dynamic', 'none', an IP address ID or a reserved flexible IP address`, ip)
	}

	return sb, nil
}

// AddVolumes build volume templates from arguments.
//
// More format details in buildVolumeTemplate function.
//
// Also add default volumes to server, ex: scratch storage for GPU servers
func (sb *ServerBuilder) AddVolumes(rootVolume string, additionalVolumes []string) (*ServerBuilder, error) {
	var err error

	if len(additionalVolumes) > 0 || rootVolume != "" {
		if rootVolume != "" {
			rootVolumeBuilder, err := NewVolumeBuilder(sb.createReq.Zone, rootVolume)
			if err != nil {
				return sb, fmt.Errorf("failed to create root volume builder: %w", err)
			}
			sb.rootVolume = rootVolumeBuilder
		}
		for _, additionalVolume := range additionalVolumes {
			additionalVolumeBuilder, err := NewVolumeBuilder(sb.createReq.Zone, additionalVolume)
			if err != nil {
				return sb, fmt.Errorf("failed to create additional volume builder: %w", err)
			}
			sb.volumes = append(sb.volumes, additionalVolumeBuilder)
		}

		volumes := make(map[string]*instance.VolumeServerTemplate, len(sb.volumes)+1)
		if sb.rootVolume != nil {
			volumes["0"], err = sb.rootVolume.BuildVolumeServerTemplate(sb.apiInstance, sb.apiBlock)
			if err != nil {
				return sb, fmt.Errorf("failed to build root volume: %w", err)
			}
		}
		for i, volume := range sb.volumes {
			volumeTemplate, err := volume.BuildVolumeServerTemplate(sb.apiInstance, sb.apiBlock)
			if err != nil {
				return sb, fmt.Errorf("failed to build volume template: %w", err)
			}
			index := strconv.Itoa(i + 1)
			volumeTemplate.Name = scw.StringPtr(sb.createReq.Name + "-" + index)
			volumes[index] = volumeTemplate
		}
		// Sanitize the volume map to respect API schemas
		sb.createReq.Volumes = volumes
	}

	if sb.serverType != nil {
		sb.createReq.Volumes = addDefaultVolumes(sb.serverType, sb.createReq.Volumes)
	}

	return sb, nil
}

// ValidateVolumes validates that the volumes are valid and sanitize the prepared template.
// Server creation should fail if ValidateVolumes is not ran before.
func (sb *ServerBuilder) ValidateVolumes() error {
	volumes := sb.createReq.Volumes
	if volumes != nil {
		// Validate root volume type and size.
		if _, hasRootVolume := volumes["0"]; sb.serverImage != nil && hasRootVolume {
			if err := validateRootVolume(sb.serverImage.RootVolume.Size, volumes["0"]); err != nil {
				return err
			}
		} else {
			logger.Warningf("skipping root volume validation")
		}

		// Validate total local volume sizes.
		if sb.serverType != nil && sb.serverImage != nil {
			if err := validateLocalVolumeSizes(volumes, sb.serverType, sb.createReq.CommercialType, sb.serverImage.RootVolume.Size); err != nil {
				return err
			}
		} else {
			logger.Warningf("skip local volume size validation")
		}

		// Sanitize the volume map to respect API schemas
		sb.createReq.Volumes = sanitizeVolumeMap(sb.createReq.Name, volumes)
	}

	return nil
}

func (sb *ServerBuilder) AddBootType(bootType string) *ServerBuilder {
	instanceBootType := instance.BootType(bootType)
	sb.createReq.BootType = &instanceBootType

	return sb
}

func (sb *ServerBuilder) AddSecurityGroup(securityGroupID string) *ServerBuilder {
	if securityGroupID != "" {
		sb.createReq.SecurityGroup = &securityGroupID
	}

	return sb
}

func (sb *ServerBuilder) AddPlacementGroup(placementGroupID string) *ServerBuilder {
	if placementGroupID != "" {
		sb.createReq.PlacementGroup = &placementGroupID
	}

	return sb
}

func (sb *ServerBuilder) Validate() error {
	if sb.isWindows() && sb.createReq.AdminPasswordEncryptionSSHKeyID == nil {
		return &core.CliError{
			Err:     core.MissingRequiredArgumentError("admin-password-encryption-ssh-key-id").Err,
			Details: "Expected a SSH Key ID to encrypt Admin RDP password. If not provided, no password will be generated. Key must be RSA Public Key.",
			Hint:    "Use completion or get your ssh key id using 'scw iam ssh-key list',",
			Code:    1,
			Empty:   false,
		}
	}

	if sb.serverType != nil && sb.serverImage != nil {
		if err := validateImageServerTypeCompatibility(sb.serverImage, sb.serverType, sb.createReq.CommercialType); err != nil {
			return err
		}
	} else {
		logger.Warningf("skipping image server-type compatibility validation")
	}

	return sb.ValidateVolumes()
}

func (sb *ServerBuilder) Build() (*instance.CreateServerRequest, []*instance.CreateIPRequest) {
	return sb.createReq, sb.createIPReqs
}

type PostServerCreationSetupFunc func(ctx context.Context, server *instance.Server) error

func (sb *ServerBuilder) BuildPostCreationSetup() PostServerCreationSetupFunc {
	rootVolume := sb.rootVolume
	volumes := sb.volumes

	return func(ctx context.Context, server *instance.Server) error {
		if rootVolume != nil {
			serverRootVolume, hasRootVolume := server.Volumes["0"]
			if !hasRootVolume {
				return errors.New("root volume not found")
			}
			rootVolume.ExecutePostCreationSetup(ctx, sb.apiBlock, serverRootVolume.ID)
		}
		for i, volume := range volumes {
			serverVolume, serverHasVolume := server.Volumes[strconv.Itoa(i+1)]
			if !serverHasVolume {
				return fmt.Errorf("volume %d not found in server", i+1)
			}
			volume.ExecutePostCreationSetup(ctx, sb.apiBlock, serverVolume.ID)
		}

		return nil
	}
}

type VolumeBuilder struct {
	Zone       scw.Zone
	VolumeType instance.VolumeVolumeType

	// SnapshotID is the ID of the snapshot the volume should be created from.
	SnapshotID *string
	// VolumeID is the ID of the volume if one should be imported.
	VolumeID *string
	// Size is the size of the created Volume. If used, the volume should be created from scratch.
	Size *scw.Size
	// IOPS is the io per second to be configured for a created volume.
	IOPS *uint32
}

// NewVolumeBuilder creates a volume builder from a 'volumes' argument item.
//
// Volumes definition must be through multiple arguments (eg: volumes.0="l:20GB" volumes.1="b:100GB" volumes.2="sbs:50GB:15000)
//
// A valid volume format is either
// - a "creation" format: ^((local|l|block|b|scratch|s|sbs):)?\d+GB?(:\d+)?$ (size is handled by go-humanize, so other sizes are supported)
// - a "creation" format with a snapshot id: l:<uuid> b:<uuid>
// - a UUID format
func NewVolumeBuilder(zone scw.Zone, flagV string) (*VolumeBuilder, error) {
	parts := strings.Split(strings.TrimSpace(flagV), ":")
	vb := &VolumeBuilder{
		Zone: zone,
	}

	if len(parts) == 3 {
		iops, err := strconv.ParseUint(parts[2], 10, 32)
		if err != nil {
			return nil, fmt.Errorf("invalid volume iops %s in %s volume", parts[2], flagV)
		}
		vb.IOPS = scw.Uint32Ptr(uint32(iops))
		parts = parts[0:2]
	}

	if len(parts) == 2 {
		switch parts[0] {
		case "l", "local":
			vb.VolumeType = instance.VolumeVolumeTypeLSSD
		case "b", "block":
			vb.VolumeType = instance.VolumeVolumeTypeBSSD
		case "s", "scratch":
			vb.VolumeType = instance.VolumeVolumeTypeScratch
		case "sbs":
			vb.VolumeType = instance.VolumeVolumeTypeSbsVolume
		default:
			return nil, fmt.Errorf("invalid volume type %s in %s volume", parts[0], flagV)
		}

		if validation.IsUUID(parts[1]) {
			vb.SnapshotID = scw.StringPtr(parts[1])
		} else {
			size, err := humanize.ParseBytes(parts[1])
			if err != nil {
				return nil, fmt.Errorf("invalid size format %s in %s volume", parts[1], flagV)
			}
			vb.Size = scw.SizePtr(scw.Size(size))
		}

		return vb, nil
	}

	// UUID format.
	if len(parts) == 1 && validation.IsUUID(parts[0]) {
		vb.VolumeID = scw.StringPtr(parts[0])

		return vb, nil
	}

	return nil, &core.CliError{
		Err:     fmt.Errorf("invalid volume format '%s'", flagV),
		Details: "",
		Hint:    `You must provide either a UUID ("11111111-1111-1111-1111-111111111111"), a local volume size ("local:100G" or "l:100G") or a block volume size ("block:100G" or "b:100G").`,
	}
}

// buildSnapshotVolume builds the requested volume template to create a new volume from a snapshot
func (vb *VolumeBuilder) buildSnapshotVolume(api *instance.API) (*instance.VolumeServerTemplate, error) {
	if vb.SnapshotID == nil {
		return nil, errors.New("tried to build a volume from snapshot with an empty ID")
	}
	res, err := api.GetSnapshot(&instance.GetSnapshotRequest{
		Zone:       vb.Zone,
		SnapshotID: *vb.SnapshotID,
	})
	if err != nil {
		if core.IsNotFoundError(err) {
			return nil, fmt.Errorf("snapshot %s does not exist", *vb.SnapshotID)
		}
	}

	snapshotType := res.Snapshot.VolumeType

	if snapshotType != instance.VolumeVolumeTypeUnified && snapshotType != vb.VolumeType {
		return nil, fmt.Errorf("snapshot of type %s not compatible with requested volume type %s", snapshotType, vb.VolumeType)
	}

	return &instance.VolumeServerTemplate{
		Name:         &res.Snapshot.Name,
		VolumeType:   vb.VolumeType,
		BaseSnapshot: &res.Snapshot.ID,
		Size:         &res.Snapshot.Size,
	}, nil
}

// buildImportedVolume builds the requested volume template to import an existing volume
func (vb *VolumeBuilder) buildImportedVolume(api *instance.API, blockAPI *block.API) (*instance.VolumeServerTemplate, error) {
	if vb.VolumeID == nil {
		return nil, errors.New("tried to import a volume with an empty ID")
	}

	res, err := api.GetVolume(&instance.GetVolumeRequest{
		Zone:     vb.Zone,
		VolumeID: *vb.VolumeID,
	})
	if err != nil && !core.IsNotFoundError(err) {
		return nil, err
	}

	if res != nil {
		// Check that volume is not already attached to a server.
		if res.Volume.Server != nil {
			return nil, fmt.Errorf("volume %s is already attached to %s server", res.Volume.ID, res.Volume.Server.ID)
		}

		return &instance.VolumeServerTemplate{
			ID:         &res.Volume.ID,
			VolumeType: res.Volume.VolumeType,
			Size:       &res.Volume.Size,
		}, nil
	}

	blockRes, err := blockAPI.GetVolume(&block.GetVolumeRequest{
		Zone:     vb.Zone,
		VolumeID: *vb.VolumeID,
	})
	if err != nil {
		if core.IsNotFoundError(err) {
			return nil, fmt.Errorf("volume %s does not exist", *vb.VolumeID)
		}
		return nil, err
	}

	if len(blockRes.References) > 0 {
		return nil, fmt.Errorf("volume %s is already attached to %s %s", blockRes.ID, blockRes.References[0].ProductResourceID, blockRes.References[0].ProductResourceType)
	}

	return &instance.VolumeServerTemplate{
		ID:         &blockRes.ID,
		VolumeType: instance.VolumeVolumeTypeSbsVolume, // TODO: support snapshot
	}, nil
}

// buildNewVolume builds the requested volume template to create a new volume with requested size
func (vb *VolumeBuilder) buildNewVolume() (*instance.VolumeServerTemplate, error) {
	return &instance.VolumeServerTemplate{
		VolumeType: vb.VolumeType,
		Size:       vb.Size,
	}, nil
}

// BuildVolumeServerTemplate builds the requested volume template to be used in a CreateServerRequest
func (vb *VolumeBuilder) BuildVolumeServerTemplate(apiInstance *instance.API, apiBlock *block.API) (*instance.VolumeServerTemplate, error) {
	if vb.SnapshotID != nil {
		return vb.buildSnapshotVolume(apiInstance)
	}

	if vb.VolumeID != nil {
		return vb.buildImportedVolume(apiInstance, apiBlock)
	}

	return vb.buildNewVolume()
}

// ExecutePostCreationSetup executes requests that are required after volume creation.
func (vb *VolumeBuilder) ExecutePostCreationSetup(ctx context.Context, apiBlock *block.API, volumeID string) {
	if vb.IOPS != nil {
		_, err := apiBlock.UpdateVolume(&block.UpdateVolumeRequest{
			VolumeID: volumeID,
			PerfIops: vb.IOPS,
		},
			scw.WithContext(ctx),
		)
		if err != nil {
			core.ExtractLogger(ctx).Warning(fmt.Sprintf("Failed to update volume %s IOPS: %s", volumeID, err.Error()))
		}
	}
}
