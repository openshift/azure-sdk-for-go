// +build go1.9

// Copyright 2019 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package databoxedge

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/databoxedge/mgmt/2019-08-01/databoxedge"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AccountType = original.AccountType

const (
	BlobStorage           AccountType = original.BlobStorage
	GeneralPurposeStorage AccountType = original.GeneralPurposeStorage
)

type AlertSeverity = original.AlertSeverity

const (
	Critical      AlertSeverity = original.Critical
	Informational AlertSeverity = original.Informational
	Warning       AlertSeverity = original.Warning
)

type AuthenticationType = original.AuthenticationType

const (
	AzureActiveDirectory AuthenticationType = original.AzureActiveDirectory
	Invalid              AuthenticationType = original.Invalid
)

type AzureContainerDataFormat = original.AzureContainerDataFormat

const (
	AzureFile AzureContainerDataFormat = original.AzureFile
	BlockBlob AzureContainerDataFormat = original.BlockBlob
	PageBlob  AzureContainerDataFormat = original.PageBlob
)

type ClientPermissionType = original.ClientPermissionType

const (
	NoAccess  ClientPermissionType = original.NoAccess
	ReadOnly  ClientPermissionType = original.ReadOnly
	ReadWrite ClientPermissionType = original.ReadWrite
)

type ContainerStatus = original.ContainerStatus

const (
	NeedsAttention ContainerStatus = original.NeedsAttention
	Offline        ContainerStatus = original.Offline
	OK             ContainerStatus = original.OK
	Unknown        ContainerStatus = original.Unknown
	Updating       ContainerStatus = original.Updating
)

type DataPolicy = original.DataPolicy

const (
	Cloud DataPolicy = original.Cloud
	Local DataPolicy = original.Local
)

type DayOfWeek = original.DayOfWeek

const (
	Friday    DayOfWeek = original.Friday
	Monday    DayOfWeek = original.Monday
	Saturday  DayOfWeek = original.Saturday
	Sunday    DayOfWeek = original.Sunday
	Thursday  DayOfWeek = original.Thursday
	Tuesday   DayOfWeek = original.Tuesday
	Wednesday DayOfWeek = original.Wednesday
)

type DeviceStatus = original.DeviceStatus

const (
	DeviceStatusDisconnected          DeviceStatus = original.DeviceStatusDisconnected
	DeviceStatusMaintenance           DeviceStatus = original.DeviceStatusMaintenance
	DeviceStatusNeedsAttention        DeviceStatus = original.DeviceStatusNeedsAttention
	DeviceStatusOffline               DeviceStatus = original.DeviceStatusOffline
	DeviceStatusOnline                DeviceStatus = original.DeviceStatusOnline
	DeviceStatusPartiallyDisconnected DeviceStatus = original.DeviceStatusPartiallyDisconnected
	DeviceStatusReadyToSetup          DeviceStatus = original.DeviceStatusReadyToSetup
)

type DeviceType = original.DeviceType

const (
	DataBoxEdgeDevice DeviceType = original.DataBoxEdgeDevice
)

type DownloadPhase = original.DownloadPhase

const (
	DownloadPhaseDownloading  DownloadPhase = original.DownloadPhaseDownloading
	DownloadPhaseInitializing DownloadPhase = original.DownloadPhaseInitializing
	DownloadPhaseUnknown      DownloadPhase = original.DownloadPhaseUnknown
	DownloadPhaseVerifying    DownloadPhase = original.DownloadPhaseVerifying
)

type EncryptionAlgorithm = original.EncryptionAlgorithm

const (
	AES256        EncryptionAlgorithm = original.AES256
	None          EncryptionAlgorithm = original.None
	RSAESPKCS1V15 EncryptionAlgorithm = original.RSAESPKCS1V15
)

type InstallRebootBehavior = original.InstallRebootBehavior

const (
	NeverReboots   InstallRebootBehavior = original.NeverReboots
	RequestReboot  InstallRebootBehavior = original.RequestReboot
	RequiresReboot InstallRebootBehavior = original.RequiresReboot
)

type JobStatus = original.JobStatus

const (
	JobStatusCanceled  JobStatus = original.JobStatusCanceled
	JobStatusFailed    JobStatus = original.JobStatusFailed
	JobStatusInvalid   JobStatus = original.JobStatusInvalid
	JobStatusPaused    JobStatus = original.JobStatusPaused
	JobStatusRunning   JobStatus = original.JobStatusRunning
	JobStatusScheduled JobStatus = original.JobStatusScheduled
	JobStatusSucceeded JobStatus = original.JobStatusSucceeded
)

type JobType = original.JobType

const (
	JobTypeDownloadUpdates  JobType = original.JobTypeDownloadUpdates
	JobTypeInstallUpdates   JobType = original.JobTypeInstallUpdates
	JobTypeInvalid          JobType = original.JobTypeInvalid
	JobTypeRefreshContainer JobType = original.JobTypeRefreshContainer
	JobTypeRefreshShare     JobType = original.JobTypeRefreshShare
	JobTypeScanForUpdates   JobType = original.JobTypeScanForUpdates
)

type Kind = original.Kind

const (
	KindIOT  Kind = original.KindIOT
	KindRole Kind = original.KindRole
)

type KindBasicTrigger = original.KindBasicTrigger

const (
	KindFileEvent          KindBasicTrigger = original.KindFileEvent
	KindPeriodicTimerEvent KindBasicTrigger = original.KindPeriodicTimerEvent
	KindTrigger            KindBasicTrigger = original.KindTrigger
)

type MetricAggregationType = original.MetricAggregationType

const (
	MetricAggregationTypeAverage      MetricAggregationType = original.MetricAggregationTypeAverage
	MetricAggregationTypeCount        MetricAggregationType = original.MetricAggregationTypeCount
	MetricAggregationTypeMaximum      MetricAggregationType = original.MetricAggregationTypeMaximum
	MetricAggregationTypeMinimum      MetricAggregationType = original.MetricAggregationTypeMinimum
	MetricAggregationTypeNone         MetricAggregationType = original.MetricAggregationTypeNone
	MetricAggregationTypeNotSpecified MetricAggregationType = original.MetricAggregationTypeNotSpecified
	MetricAggregationTypeTotal        MetricAggregationType = original.MetricAggregationTypeTotal
)

type MetricCategory = original.MetricCategory

const (
	Capacity    MetricCategory = original.Capacity
	Transaction MetricCategory = original.Transaction
)

type MetricUnit = original.MetricUnit

const (
	Bytes          MetricUnit = original.Bytes
	BytesPerSecond MetricUnit = original.BytesPerSecond
	Count          MetricUnit = original.Count
	CountPerSecond MetricUnit = original.CountPerSecond
	Milliseconds   MetricUnit = original.Milliseconds
	NotSpecified   MetricUnit = original.NotSpecified
	Percent        MetricUnit = original.Percent
	Seconds        MetricUnit = original.Seconds
)

type MonitoringStatus = original.MonitoringStatus

const (
	Disabled MonitoringStatus = original.Disabled
	Enabled  MonitoringStatus = original.Enabled
)

type NetworkAdapterDHCPStatus = original.NetworkAdapterDHCPStatus

const (
	NetworkAdapterDHCPStatusDisabled NetworkAdapterDHCPStatus = original.NetworkAdapterDHCPStatusDisabled
	NetworkAdapterDHCPStatusEnabled  NetworkAdapterDHCPStatus = original.NetworkAdapterDHCPStatusEnabled
)

type NetworkAdapterRDMAStatus = original.NetworkAdapterRDMAStatus

const (
	Capable   NetworkAdapterRDMAStatus = original.Capable
	Incapable NetworkAdapterRDMAStatus = original.Incapable
)

type NetworkAdapterStatus = original.NetworkAdapterStatus

const (
	Active   NetworkAdapterStatus = original.Active
	Inactive NetworkAdapterStatus = original.Inactive
)

type NetworkGroup = original.NetworkGroup

const (
	NetworkGroupNone    NetworkGroup = original.NetworkGroupNone
	NetworkGroupNonRDMA NetworkGroup = original.NetworkGroupNonRDMA
	NetworkGroupRDMA    NetworkGroup = original.NetworkGroupRDMA
)

type NodeStatus = original.NodeStatus

const (
	NodeStatusDown         NodeStatus = original.NodeStatusDown
	NodeStatusRebooting    NodeStatus = original.NodeStatusRebooting
	NodeStatusShuttingDown NodeStatus = original.NodeStatusShuttingDown
	NodeStatusUnknown      NodeStatus = original.NodeStatusUnknown
	NodeStatusUp           NodeStatus = original.NodeStatusUp
)

type OrderState = original.OrderState

const (
	Arriving               OrderState = original.Arriving
	AwaitingFulfilment     OrderState = original.AwaitingFulfilment
	AwaitingPreparation    OrderState = original.AwaitingPreparation
	AwaitingReturnShipment OrderState = original.AwaitingReturnShipment
	AwaitingShipment       OrderState = original.AwaitingShipment
	CollectedAtMicrosoft   OrderState = original.CollectedAtMicrosoft
	Declined               OrderState = original.Declined
	Delivered              OrderState = original.Delivered
	LostDevice             OrderState = original.LostDevice
	ReplacementRequested   OrderState = original.ReplacementRequested
	ReturnInitiated        OrderState = original.ReturnInitiated
	Shipped                OrderState = original.Shipped
	ShippedBack            OrderState = original.ShippedBack
	Untracked              OrderState = original.Untracked
)

type PlatformType = original.PlatformType

const (
	Linux   PlatformType = original.Linux
	Windows PlatformType = original.Windows
)

type RoleStatus = original.RoleStatus

const (
	RoleStatusDisabled RoleStatus = original.RoleStatusDisabled
	RoleStatusEnabled  RoleStatus = original.RoleStatusEnabled
)

type RoleTypes = original.RoleTypes

const (
	ASA       RoleTypes = original.ASA
	Cognitive RoleTypes = original.Cognitive
	Functions RoleTypes = original.Functions
	IOT       RoleTypes = original.IOT
)

type SSLStatus = original.SSLStatus

const (
	SSLStatusDisabled SSLStatus = original.SSLStatusDisabled
	SSLStatusEnabled  SSLStatus = original.SSLStatusEnabled
)

type ShareAccessProtocol = original.ShareAccessProtocol

const (
	NFS ShareAccessProtocol = original.NFS
	SMB ShareAccessProtocol = original.SMB
)

type ShareAccessType = original.ShareAccessType

const (
	Change ShareAccessType = original.Change
	Custom ShareAccessType = original.Custom
	Read   ShareAccessType = original.Read
)

type ShareStatus = original.ShareStatus

const (
	ShareStatusNeedsAttention ShareStatus = original.ShareStatusNeedsAttention
	ShareStatusOffline        ShareStatus = original.ShareStatusOffline
	ShareStatusOK             ShareStatus = original.ShareStatusOK
	ShareStatusUnknown        ShareStatus = original.ShareStatusUnknown
	ShareStatusUpdating       ShareStatus = original.ShareStatusUpdating
)

type SkuName = original.SkuName

const (
	Edge              SkuName = original.Edge
	Gateway           SkuName = original.Gateway
	TEA1Node          SkuName = original.TEA1Node
	TEA1NodeHeater    SkuName = original.TEA1NodeHeater
	TEA1NodeUPS       SkuName = original.TEA1NodeUPS
	TEA1NodeUPSHeater SkuName = original.TEA1NodeUPSHeater
	TEA4NodeHeater    SkuName = original.TEA4NodeHeater
	TEA4NodeUPSHeater SkuName = original.TEA4NodeUPSHeater
	TMA               SkuName = original.TMA
)

type SkuRestrictionReasonCode = original.SkuRestrictionReasonCode

const (
	NotAvailableForSubscription SkuRestrictionReasonCode = original.NotAvailableForSubscription
	QuotaID                     SkuRestrictionReasonCode = original.QuotaID
)

type SkuTier = original.SkuTier

const (
	Standard SkuTier = original.Standard
)

type StorageAccountStatus = original.StorageAccountStatus

const (
	StorageAccountStatusNeedsAttention StorageAccountStatus = original.StorageAccountStatusNeedsAttention
	StorageAccountStatusOffline        StorageAccountStatus = original.StorageAccountStatusOffline
	StorageAccountStatusOK             StorageAccountStatus = original.StorageAccountStatusOK
	StorageAccountStatusUnknown        StorageAccountStatus = original.StorageAccountStatusUnknown
	StorageAccountStatusUpdating       StorageAccountStatus = original.StorageAccountStatusUpdating
)

type TimeGrain = original.TimeGrain

const (
	PT12H TimeGrain = original.PT12H
	PT15M TimeGrain = original.PT15M
	PT1D  TimeGrain = original.PT1D
	PT1H  TimeGrain = original.PT1H
	PT1M  TimeGrain = original.PT1M
	PT30M TimeGrain = original.PT30M
	PT5M  TimeGrain = original.PT5M
	PT6H  TimeGrain = original.PT6H
)

type UpdateOperation = original.UpdateOperation

const (
	UpdateOperationDownload UpdateOperation = original.UpdateOperationDownload
	UpdateOperationInstall  UpdateOperation = original.UpdateOperationInstall
	UpdateOperationNone     UpdateOperation = original.UpdateOperationNone
	UpdateOperationScan     UpdateOperation = original.UpdateOperationScan
)

type UpdateOperationStage = original.UpdateOperationStage

const (
	UpdateOperationStageDownloadComplete UpdateOperationStage = original.UpdateOperationStageDownloadComplete
	UpdateOperationStageDownloadFailed   UpdateOperationStage = original.UpdateOperationStageDownloadFailed
	UpdateOperationStageDownloadStarted  UpdateOperationStage = original.UpdateOperationStageDownloadStarted
	UpdateOperationStageFailure          UpdateOperationStage = original.UpdateOperationStageFailure
	UpdateOperationStageInitial          UpdateOperationStage = original.UpdateOperationStageInitial
	UpdateOperationStageInstallComplete  UpdateOperationStage = original.UpdateOperationStageInstallComplete
	UpdateOperationStageInstallFailed    UpdateOperationStage = original.UpdateOperationStageInstallFailed
	UpdateOperationStageInstallStarted   UpdateOperationStage = original.UpdateOperationStageInstallStarted
	UpdateOperationStageRebootInitiated  UpdateOperationStage = original.UpdateOperationStageRebootInitiated
	UpdateOperationStageRescanComplete   UpdateOperationStage = original.UpdateOperationStageRescanComplete
	UpdateOperationStageRescanFailed     UpdateOperationStage = original.UpdateOperationStageRescanFailed
	UpdateOperationStageRescanStarted    UpdateOperationStage = original.UpdateOperationStageRescanStarted
	UpdateOperationStageScanComplete     UpdateOperationStage = original.UpdateOperationStageScanComplete
	UpdateOperationStageScanFailed       UpdateOperationStage = original.UpdateOperationStageScanFailed
	UpdateOperationStageScanStarted      UpdateOperationStage = original.UpdateOperationStageScanStarted
	UpdateOperationStageSuccess          UpdateOperationStage = original.UpdateOperationStageSuccess
	UpdateOperationStageUnknown          UpdateOperationStage = original.UpdateOperationStageUnknown
)

type UserType = original.UserType

const (
	UserTypeARM             UserType = original.UserTypeARM
	UserTypeLocalManagement UserType = original.UserTypeLocalManagement
	UserTypeShare           UserType = original.UserTypeShare
)

type ARMBaseModel = original.ARMBaseModel
type Address = original.Address
type Alert = original.Alert
type AlertErrorDetails = original.AlertErrorDetails
type AlertList = original.AlertList
type AlertListIterator = original.AlertListIterator
type AlertListPage = original.AlertListPage
type AlertProperties = original.AlertProperties
type AlertsClient = original.AlertsClient
type AsymmetricEncryptedSecret = original.AsymmetricEncryptedSecret
type Authentication = original.Authentication
type AzureContainerInfo = original.AzureContainerInfo
type BandwidthSchedule = original.BandwidthSchedule
type BandwidthScheduleProperties = original.BandwidthScheduleProperties
type BandwidthSchedulesClient = original.BandwidthSchedulesClient
type BandwidthSchedulesCreateOrUpdateFuture = original.BandwidthSchedulesCreateOrUpdateFuture
type BandwidthSchedulesDeleteFuture = original.BandwidthSchedulesDeleteFuture
type BandwidthSchedulesList = original.BandwidthSchedulesList
type BandwidthSchedulesListIterator = original.BandwidthSchedulesListIterator
type BandwidthSchedulesListPage = original.BandwidthSchedulesListPage
type BaseClient = original.BaseClient
type BasicRole = original.BasicRole
type BasicTrigger = original.BasicTrigger
type ClientAccessRight = original.ClientAccessRight
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type ContactDetails = original.ContactDetails
type Container = original.Container
type ContainerList = original.ContainerList
type ContainerListIterator = original.ContainerListIterator
type ContainerListPage = original.ContainerListPage
type ContainerProperties = original.ContainerProperties
type ContainersClient = original.ContainersClient
type ContainersCreateOrUpdateFuture = original.ContainersCreateOrUpdateFuture
type ContainersDeleteFuture = original.ContainersDeleteFuture
type ContainersRefreshFuture = original.ContainersRefreshFuture
type Device = original.Device
type DeviceExtendedInfo = original.DeviceExtendedInfo
type DeviceExtendedInfoProperties = original.DeviceExtendedInfoProperties
type DeviceList = original.DeviceList
type DeviceListIterator = original.DeviceListIterator
type DeviceListPage = original.DeviceListPage
type DevicePatch = original.DevicePatch
type DeviceProperties = original.DeviceProperties
type DevicesClient = original.DevicesClient
type DevicesCreateOrUpdateFuture = original.DevicesCreateOrUpdateFuture
type DevicesCreateOrUpdateSecuritySettingsFuture = original.DevicesCreateOrUpdateSecuritySettingsFuture
type DevicesDeleteFuture = original.DevicesDeleteFuture
type DevicesDownloadUpdatesFuture = original.DevicesDownloadUpdatesFuture
type DevicesInstallUpdatesFuture = original.DevicesInstallUpdatesFuture
type DevicesScanForUpdatesFuture = original.DevicesScanForUpdatesFuture
type FileEventTrigger = original.FileEventTrigger
type FileSourceInfo = original.FileSourceInfo
type FileTriggerProperties = original.FileTriggerProperties
type IoTDeviceInfo = original.IoTDeviceInfo
type IoTRole = original.IoTRole
type IoTRoleProperties = original.IoTRoleProperties
type Ipv4Config = original.Ipv4Config
type Ipv6Config = original.Ipv6Config
type Job = original.Job
type JobErrorDetails = original.JobErrorDetails
type JobErrorItem = original.JobErrorItem
type JobProperties = original.JobProperties
type JobsClient = original.JobsClient
type MetricDimensionV1 = original.MetricDimensionV1
type MetricSpecificationV1 = original.MetricSpecificationV1
type MountPointMap = original.MountPointMap
type NetworkAdapter = original.NetworkAdapter
type NetworkAdapterPosition = original.NetworkAdapterPosition
type NetworkSettings = original.NetworkSettings
type NetworkSettingsProperties = original.NetworkSettingsProperties
type Node = original.Node
type NodeList = original.NodeList
type NodeProperties = original.NodeProperties
type NodesClient = original.NodesClient
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationProperties = original.OperationProperties
type OperationsClient = original.OperationsClient
type OperationsList = original.OperationsList
type OperationsListIterator = original.OperationsListIterator
type OperationsListPage = original.OperationsListPage
type OperationsStatusClient = original.OperationsStatusClient
type Order = original.Order
type OrderList = original.OrderList
type OrderListIterator = original.OrderListIterator
type OrderListPage = original.OrderListPage
type OrderProperties = original.OrderProperties
type OrderStatus = original.OrderStatus
type OrdersClient = original.OrdersClient
type OrdersCreateOrUpdateFuture = original.OrdersCreateOrUpdateFuture
type OrdersDeleteFuture = original.OrdersDeleteFuture
type PeriodicTimerEventTrigger = original.PeriodicTimerEventTrigger
type PeriodicTimerProperties = original.PeriodicTimerProperties
type PeriodicTimerSourceInfo = original.PeriodicTimerSourceInfo
type RawCertificateData = original.RawCertificateData
type RefreshDetails = original.RefreshDetails
type ResourceTypeSku = original.ResourceTypeSku
type Role = original.Role
type RoleList = original.RoleList
type RoleListIterator = original.RoleListIterator
type RoleListPage = original.RoleListPage
type RoleModel = original.RoleModel
type RoleSinkInfo = original.RoleSinkInfo
type RolesClient = original.RolesClient
type RolesCreateOrUpdateFuture = original.RolesCreateOrUpdateFuture
type RolesDeleteFuture = original.RolesDeleteFuture
type SecuritySettings = original.SecuritySettings
type SecuritySettingsProperties = original.SecuritySettingsProperties
type ServiceSpecification = original.ServiceSpecification
type Share = original.Share
type ShareAccessRight = original.ShareAccessRight
type ShareList = original.ShareList
type ShareListIterator = original.ShareListIterator
type ShareListPage = original.ShareListPage
type ShareProperties = original.ShareProperties
type SharesClient = original.SharesClient
type SharesCreateOrUpdateFuture = original.SharesCreateOrUpdateFuture
type SharesDeleteFuture = original.SharesDeleteFuture
type SharesRefreshFuture = original.SharesRefreshFuture
type Sku = original.Sku
type SkuCost = original.SkuCost
type SkuInformationList = original.SkuInformationList
type SkuLocationInfo = original.SkuLocationInfo
type SkuRestriction = original.SkuRestriction
type SkuRestrictionInfo = original.SkuRestrictionInfo
type SkusClient = original.SkusClient
type StorageAccount = original.StorageAccount
type StorageAccountCredential = original.StorageAccountCredential
type StorageAccountCredentialList = original.StorageAccountCredentialList
type StorageAccountCredentialListIterator = original.StorageAccountCredentialListIterator
type StorageAccountCredentialListPage = original.StorageAccountCredentialListPage
type StorageAccountCredentialProperties = original.StorageAccountCredentialProperties
type StorageAccountCredentialsClient = original.StorageAccountCredentialsClient
type StorageAccountCredentialsCreateOrUpdateFuture = original.StorageAccountCredentialsCreateOrUpdateFuture
type StorageAccountCredentialsDeleteFuture = original.StorageAccountCredentialsDeleteFuture
type StorageAccountList = original.StorageAccountList
type StorageAccountListIterator = original.StorageAccountListIterator
type StorageAccountListPage = original.StorageAccountListPage
type StorageAccountProperties = original.StorageAccountProperties
type StorageAccountsClient = original.StorageAccountsClient
type StorageAccountsCreateOrUpdateFuture = original.StorageAccountsCreateOrUpdateFuture
type StorageAccountsDeleteFuture = original.StorageAccountsDeleteFuture
type SymmetricKey = original.SymmetricKey
type TrackingInfo = original.TrackingInfo
type Trigger = original.Trigger
type TriggerList = original.TriggerList
type TriggerListIterator = original.TriggerListIterator
type TriggerListPage = original.TriggerListPage
type TriggerModel = original.TriggerModel
type TriggersClient = original.TriggersClient
type TriggersCreateOrUpdateFuture = original.TriggersCreateOrUpdateFuture
type TriggersDeleteFuture = original.TriggersDeleteFuture
type UpdateDownloadProgress = original.UpdateDownloadProgress
type UpdateInstallProgress = original.UpdateInstallProgress
type UpdateSummary = original.UpdateSummary
type UpdateSummaryProperties = original.UpdateSummaryProperties
type UploadCertificateRequest = original.UploadCertificateRequest
type UploadCertificateResponse = original.UploadCertificateResponse
type User = original.User
type UserAccessRight = original.UserAccessRight
type UserList = original.UserList
type UserListIterator = original.UserListIterator
type UserListPage = original.UserListPage
type UserProperties = original.UserProperties
type UsersClient = original.UsersClient
type UsersCreateOrUpdateFuture = original.UsersCreateOrUpdateFuture
type UsersDeleteFuture = original.UsersDeleteFuture

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAlertListIterator(page AlertListPage) AlertListIterator {
	return original.NewAlertListIterator(page)
}
func NewAlertListPage(getNextPage func(context.Context, AlertList) (AlertList, error)) AlertListPage {
	return original.NewAlertListPage(getNextPage)
}
func NewAlertsClient(subscriptionID string) AlertsClient {
	return original.NewAlertsClient(subscriptionID)
}
func NewAlertsClientWithBaseURI(baseURI string, subscriptionID string) AlertsClient {
	return original.NewAlertsClientWithBaseURI(baseURI, subscriptionID)
}
func NewBandwidthSchedulesClient(subscriptionID string) BandwidthSchedulesClient {
	return original.NewBandwidthSchedulesClient(subscriptionID)
}
func NewBandwidthSchedulesClientWithBaseURI(baseURI string, subscriptionID string) BandwidthSchedulesClient {
	return original.NewBandwidthSchedulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewBandwidthSchedulesListIterator(page BandwidthSchedulesListPage) BandwidthSchedulesListIterator {
	return original.NewBandwidthSchedulesListIterator(page)
}
func NewBandwidthSchedulesListPage(getNextPage func(context.Context, BandwidthSchedulesList) (BandwidthSchedulesList, error)) BandwidthSchedulesListPage {
	return original.NewBandwidthSchedulesListPage(getNextPage)
}
func NewContainerListIterator(page ContainerListPage) ContainerListIterator {
	return original.NewContainerListIterator(page)
}
func NewContainerListPage(getNextPage func(context.Context, ContainerList) (ContainerList, error)) ContainerListPage {
	return original.NewContainerListPage(getNextPage)
}
func NewContainersClient(subscriptionID string) ContainersClient {
	return original.NewContainersClient(subscriptionID)
}
func NewContainersClientWithBaseURI(baseURI string, subscriptionID string) ContainersClient {
	return original.NewContainersClientWithBaseURI(baseURI, subscriptionID)
}
func NewDeviceListIterator(page DeviceListPage) DeviceListIterator {
	return original.NewDeviceListIterator(page)
}
func NewDeviceListPage(getNextPage func(context.Context, DeviceList) (DeviceList, error)) DeviceListPage {
	return original.NewDeviceListPage(getNextPage)
}
func NewDevicesClient(subscriptionID string) DevicesClient {
	return original.NewDevicesClient(subscriptionID)
}
func NewDevicesClientWithBaseURI(baseURI string, subscriptionID string) DevicesClient {
	return original.NewDevicesClientWithBaseURI(baseURI, subscriptionID)
}
func NewJobsClient(subscriptionID string) JobsClient {
	return original.NewJobsClient(subscriptionID)
}
func NewJobsClientWithBaseURI(baseURI string, subscriptionID string) JobsClient {
	return original.NewJobsClientWithBaseURI(baseURI, subscriptionID)
}
func NewNodesClient(subscriptionID string) NodesClient {
	return original.NewNodesClient(subscriptionID)
}
func NewNodesClientWithBaseURI(baseURI string, subscriptionID string) NodesClient {
	return original.NewNodesClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsListIterator(page OperationsListPage) OperationsListIterator {
	return original.NewOperationsListIterator(page)
}
func NewOperationsListPage(getNextPage func(context.Context, OperationsList) (OperationsList, error)) OperationsListPage {
	return original.NewOperationsListPage(getNextPage)
}
func NewOperationsStatusClient(subscriptionID string) OperationsStatusClient {
	return original.NewOperationsStatusClient(subscriptionID)
}
func NewOperationsStatusClientWithBaseURI(baseURI string, subscriptionID string) OperationsStatusClient {
	return original.NewOperationsStatusClientWithBaseURI(baseURI, subscriptionID)
}
func NewOrderListIterator(page OrderListPage) OrderListIterator {
	return original.NewOrderListIterator(page)
}
func NewOrderListPage(getNextPage func(context.Context, OrderList) (OrderList, error)) OrderListPage {
	return original.NewOrderListPage(getNextPage)
}
func NewOrdersClient(subscriptionID string) OrdersClient {
	return original.NewOrdersClient(subscriptionID)
}
func NewOrdersClientWithBaseURI(baseURI string, subscriptionID string) OrdersClient {
	return original.NewOrdersClientWithBaseURI(baseURI, subscriptionID)
}
func NewRoleListIterator(page RoleListPage) RoleListIterator {
	return original.NewRoleListIterator(page)
}
func NewRoleListPage(getNextPage func(context.Context, RoleList) (RoleList, error)) RoleListPage {
	return original.NewRoleListPage(getNextPage)
}
func NewRolesClient(subscriptionID string) RolesClient {
	return original.NewRolesClient(subscriptionID)
}
func NewRolesClientWithBaseURI(baseURI string, subscriptionID string) RolesClient {
	return original.NewRolesClientWithBaseURI(baseURI, subscriptionID)
}
func NewShareListIterator(page ShareListPage) ShareListIterator {
	return original.NewShareListIterator(page)
}
func NewShareListPage(getNextPage func(context.Context, ShareList) (ShareList, error)) ShareListPage {
	return original.NewShareListPage(getNextPage)
}
func NewSharesClient(subscriptionID string) SharesClient {
	return original.NewSharesClient(subscriptionID)
}
func NewSharesClientWithBaseURI(baseURI string, subscriptionID string) SharesClient {
	return original.NewSharesClientWithBaseURI(baseURI, subscriptionID)
}
func NewSkusClient(subscriptionID string) SkusClient {
	return original.NewSkusClient(subscriptionID)
}
func NewSkusClientWithBaseURI(baseURI string, subscriptionID string) SkusClient {
	return original.NewSkusClientWithBaseURI(baseURI, subscriptionID)
}
func NewStorageAccountCredentialListIterator(page StorageAccountCredentialListPage) StorageAccountCredentialListIterator {
	return original.NewStorageAccountCredentialListIterator(page)
}
func NewStorageAccountCredentialListPage(getNextPage func(context.Context, StorageAccountCredentialList) (StorageAccountCredentialList, error)) StorageAccountCredentialListPage {
	return original.NewStorageAccountCredentialListPage(getNextPage)
}
func NewStorageAccountCredentialsClient(subscriptionID string) StorageAccountCredentialsClient {
	return original.NewStorageAccountCredentialsClient(subscriptionID)
}
func NewStorageAccountCredentialsClientWithBaseURI(baseURI string, subscriptionID string) StorageAccountCredentialsClient {
	return original.NewStorageAccountCredentialsClientWithBaseURI(baseURI, subscriptionID)
}
func NewStorageAccountListIterator(page StorageAccountListPage) StorageAccountListIterator {
	return original.NewStorageAccountListIterator(page)
}
func NewStorageAccountListPage(getNextPage func(context.Context, StorageAccountList) (StorageAccountList, error)) StorageAccountListPage {
	return original.NewStorageAccountListPage(getNextPage)
}
func NewStorageAccountsClient(subscriptionID string) StorageAccountsClient {
	return original.NewStorageAccountsClient(subscriptionID)
}
func NewStorageAccountsClientWithBaseURI(baseURI string, subscriptionID string) StorageAccountsClient {
	return original.NewStorageAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func NewTriggerListIterator(page TriggerListPage) TriggerListIterator {
	return original.NewTriggerListIterator(page)
}
func NewTriggerListPage(getNextPage func(context.Context, TriggerList) (TriggerList, error)) TriggerListPage {
	return original.NewTriggerListPage(getNextPage)
}
func NewTriggersClient(subscriptionID string) TriggersClient {
	return original.NewTriggersClient(subscriptionID)
}
func NewTriggersClientWithBaseURI(baseURI string, subscriptionID string) TriggersClient {
	return original.NewTriggersClientWithBaseURI(baseURI, subscriptionID)
}
func NewUserListIterator(page UserListPage) UserListIterator {
	return original.NewUserListIterator(page)
}
func NewUserListPage(getNextPage func(context.Context, UserList) (UserList, error)) UserListPage {
	return original.NewUserListPage(getNextPage)
}
func NewUsersClient(subscriptionID string) UsersClient {
	return original.NewUsersClient(subscriptionID)
}
func NewUsersClientWithBaseURI(baseURI string, subscriptionID string) UsersClient {
	return original.NewUsersClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAccountTypeValues() []AccountType {
	return original.PossibleAccountTypeValues()
}
func PossibleAlertSeverityValues() []AlertSeverity {
	return original.PossibleAlertSeverityValues()
}
func PossibleAuthenticationTypeValues() []AuthenticationType {
	return original.PossibleAuthenticationTypeValues()
}
func PossibleAzureContainerDataFormatValues() []AzureContainerDataFormat {
	return original.PossibleAzureContainerDataFormatValues()
}
func PossibleClientPermissionTypeValues() []ClientPermissionType {
	return original.PossibleClientPermissionTypeValues()
}
func PossibleContainerStatusValues() []ContainerStatus {
	return original.PossibleContainerStatusValues()
}
func PossibleDataPolicyValues() []DataPolicy {
	return original.PossibleDataPolicyValues()
}
func PossibleDayOfWeekValues() []DayOfWeek {
	return original.PossibleDayOfWeekValues()
}
func PossibleDeviceStatusValues() []DeviceStatus {
	return original.PossibleDeviceStatusValues()
}
func PossibleDeviceTypeValues() []DeviceType {
	return original.PossibleDeviceTypeValues()
}
func PossibleDownloadPhaseValues() []DownloadPhase {
	return original.PossibleDownloadPhaseValues()
}
func PossibleEncryptionAlgorithmValues() []EncryptionAlgorithm {
	return original.PossibleEncryptionAlgorithmValues()
}
func PossibleInstallRebootBehaviorValues() []InstallRebootBehavior {
	return original.PossibleInstallRebootBehaviorValues()
}
func PossibleJobStatusValues() []JobStatus {
	return original.PossibleJobStatusValues()
}
func PossibleJobTypeValues() []JobType {
	return original.PossibleJobTypeValues()
}
func PossibleKindBasicTriggerValues() []KindBasicTrigger {
	return original.PossibleKindBasicTriggerValues()
}
func PossibleKindValues() []Kind {
	return original.PossibleKindValues()
}
func PossibleMetricAggregationTypeValues() []MetricAggregationType {
	return original.PossibleMetricAggregationTypeValues()
}
func PossibleMetricCategoryValues() []MetricCategory {
	return original.PossibleMetricCategoryValues()
}
func PossibleMetricUnitValues() []MetricUnit {
	return original.PossibleMetricUnitValues()
}
func PossibleMonitoringStatusValues() []MonitoringStatus {
	return original.PossibleMonitoringStatusValues()
}
func PossibleNetworkAdapterDHCPStatusValues() []NetworkAdapterDHCPStatus {
	return original.PossibleNetworkAdapterDHCPStatusValues()
}
func PossibleNetworkAdapterRDMAStatusValues() []NetworkAdapterRDMAStatus {
	return original.PossibleNetworkAdapterRDMAStatusValues()
}
func PossibleNetworkAdapterStatusValues() []NetworkAdapterStatus {
	return original.PossibleNetworkAdapterStatusValues()
}
func PossibleNetworkGroupValues() []NetworkGroup {
	return original.PossibleNetworkGroupValues()
}
func PossibleNodeStatusValues() []NodeStatus {
	return original.PossibleNodeStatusValues()
}
func PossibleOrderStateValues() []OrderState {
	return original.PossibleOrderStateValues()
}
func PossiblePlatformTypeValues() []PlatformType {
	return original.PossiblePlatformTypeValues()
}
func PossibleRoleStatusValues() []RoleStatus {
	return original.PossibleRoleStatusValues()
}
func PossibleRoleTypesValues() []RoleTypes {
	return original.PossibleRoleTypesValues()
}
func PossibleSSLStatusValues() []SSLStatus {
	return original.PossibleSSLStatusValues()
}
func PossibleShareAccessProtocolValues() []ShareAccessProtocol {
	return original.PossibleShareAccessProtocolValues()
}
func PossibleShareAccessTypeValues() []ShareAccessType {
	return original.PossibleShareAccessTypeValues()
}
func PossibleShareStatusValues() []ShareStatus {
	return original.PossibleShareStatusValues()
}
func PossibleSkuNameValues() []SkuName {
	return original.PossibleSkuNameValues()
}
func PossibleSkuRestrictionReasonCodeValues() []SkuRestrictionReasonCode {
	return original.PossibleSkuRestrictionReasonCodeValues()
}
func PossibleSkuTierValues() []SkuTier {
	return original.PossibleSkuTierValues()
}
func PossibleStorageAccountStatusValues() []StorageAccountStatus {
	return original.PossibleStorageAccountStatusValues()
}
func PossibleTimeGrainValues() []TimeGrain {
	return original.PossibleTimeGrainValues()
}
func PossibleUpdateOperationStageValues() []UpdateOperationStage {
	return original.PossibleUpdateOperationStageValues()
}
func PossibleUpdateOperationValues() []UpdateOperation {
	return original.PossibleUpdateOperationValues()
}
func PossibleUserTypeValues() []UserType {
	return original.PossibleUserTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
