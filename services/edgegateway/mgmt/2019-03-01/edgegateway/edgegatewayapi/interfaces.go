package edgegatewayapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/edgegateway/mgmt/2019-03-01/edgegateway"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result edgegateway.OperationsListPage, err error)
}

var _ OperationsClientAPI = (*edgegateway.OperationsClient)(nil)

// DevicesClientAPI contains the set of methods on the DevicesClient type.
type DevicesClientAPI interface {
	CreateOrUpdate(ctx context.Context, deviceName string, dataBoxEdgeDevice edgegateway.DataBoxEdgeDevice, resourceGroupName string) (result edgegateway.DevicesCreateOrUpdateFuture, err error)
	CreateOrUpdateExtendedInfo(ctx context.Context, deviceName string, parameters edgegateway.DataBoxEdgeDeviceExtendedInfo, resourceGroupName string) (result edgegateway.DataBoxEdgeDeviceExtendedInfo, err error)
	CreateOrUpdateSecuritySettings(ctx context.Context, deviceName string, securitySettings edgegateway.SecuritySettings, resourceGroupName string) (result edgegateway.DevicesCreateOrUpdateSecuritySettingsFuture, err error)
	Delete(ctx context.Context, deviceName string, resourceGroupName string) (result edgegateway.DevicesDeleteFuture, err error)
	DownloadUpdates(ctx context.Context, deviceName string, resourceGroupName string) (result edgegateway.DevicesDownloadUpdatesFuture, err error)
	Get(ctx context.Context, deviceName string, resourceGroupName string) (result edgegateway.DataBoxEdgeDevice, err error)
	GetExtendedInformation(ctx context.Context, deviceName string, resourceGroupName string) (result edgegateway.DataBoxEdgeDeviceExtendedInfo, err error)
	GetNetworkSettings(ctx context.Context, deviceName string, resourceGroupName string) (result edgegateway.NetworkSettings, err error)
	GetUpdateSummary(ctx context.Context, deviceName string, resourceGroupName string) (result edgegateway.UpdateSummary, err error)
	InstallUpdates(ctx context.Context, deviceName string, resourceGroupName string) (result edgegateway.DevicesInstallUpdatesFuture, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, expand string) (result edgegateway.DataBoxEdgeDeviceListPage, err error)
	ListBySubscription(ctx context.Context, expand string) (result edgegateway.DataBoxEdgeDeviceListPage, err error)
	ScanForUpdates(ctx context.Context, deviceName string, resourceGroupName string) (result edgegateway.DevicesScanForUpdatesFuture, err error)
	Update(ctx context.Context, deviceName string, parameters edgegateway.DataBoxEdgeDevicePatch, resourceGroupName string) (result edgegateway.DataBoxEdgeDevice, err error)
	UploadCertificate(ctx context.Context, deviceName string, parameters edgegateway.UploadCertificateRequest, resourceGroupName string) (result edgegateway.UploadCertificateResponse, err error)
}

var _ DevicesClientAPI = (*edgegateway.DevicesClient)(nil)

// AlertsClientAPI contains the set of methods on the AlertsClient type.
type AlertsClientAPI interface {
	Get(ctx context.Context, deviceName string, name string, resourceGroupName string) (result edgegateway.Alert, err error)
	ListByDataBoxEdgeDevice(ctx context.Context, deviceName string, resourceGroupName string) (result edgegateway.AlertListPage, err error)
}

var _ AlertsClientAPI = (*edgegateway.AlertsClient)(nil)

// BandwidthSchedulesClientAPI contains the set of methods on the BandwidthSchedulesClient type.
type BandwidthSchedulesClientAPI interface {
	CreateOrUpdate(ctx context.Context, deviceName string, name string, parameters edgegateway.BandwidthSchedule, resourceGroupName string) (result edgegateway.BandwidthSchedulesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, deviceName string, name string, resourceGroupName string) (result edgegateway.BandwidthSchedulesDeleteFuture, err error)
	Get(ctx context.Context, deviceName string, name string, resourceGroupName string) (result edgegateway.BandwidthSchedule, err error)
	ListByDataBoxEdgeDevice(ctx context.Context, deviceName string, resourceGroupName string) (result edgegateway.BandwidthSchedulesListPage, err error)
}

var _ BandwidthSchedulesClientAPI = (*edgegateway.BandwidthSchedulesClient)(nil)

// JobsClientAPI contains the set of methods on the JobsClient type.
type JobsClientAPI interface {
	Get(ctx context.Context, deviceName string, name string, resourceGroupName string) (result edgegateway.Job, err error)
}

var _ JobsClientAPI = (*edgegateway.JobsClient)(nil)

// OperationsStatusClientAPI contains the set of methods on the OperationsStatusClient type.
type OperationsStatusClientAPI interface {
	Get(ctx context.Context, deviceName string, name string, resourceGroupName string) (result edgegateway.Job, err error)
}

var _ OperationsStatusClientAPI = (*edgegateway.OperationsStatusClient)(nil)

// OrdersClientAPI contains the set of methods on the OrdersClient type.
type OrdersClientAPI interface {
	CreateOrUpdate(ctx context.Context, deviceName string, order edgegateway.Order, resourceGroupName string) (result edgegateway.OrdersCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, deviceName string, resourceGroupName string) (result edgegateway.OrdersDeleteFuture, err error)
	Get(ctx context.Context, deviceName string, resourceGroupName string) (result edgegateway.Order, err error)
	ListByDataBoxEdgeDevice(ctx context.Context, deviceName string, resourceGroupName string) (result edgegateway.OrderListPage, err error)
}

var _ OrdersClientAPI = (*edgegateway.OrdersClient)(nil)

// RolesClientAPI contains the set of methods on the RolesClient type.
type RolesClientAPI interface {
	CreateOrUpdate(ctx context.Context, deviceName string, name string, role edgegateway.BasicRole, resourceGroupName string) (result edgegateway.RolesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, deviceName string, name string, resourceGroupName string) (result edgegateway.RolesDeleteFuture, err error)
	Get(ctx context.Context, deviceName string, name string, resourceGroupName string) (result edgegateway.RoleModel, err error)
	ListByDataBoxEdgeDevice(ctx context.Context, deviceName string, resourceGroupName string) (result edgegateway.RoleListPage, err error)
}

var _ RolesClientAPI = (*edgegateway.RolesClient)(nil)

// SharesClientAPI contains the set of methods on the SharesClient type.
type SharesClientAPI interface {
	CreateOrUpdate(ctx context.Context, deviceName string, name string, share edgegateway.Share, resourceGroupName string) (result edgegateway.SharesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, deviceName string, name string, resourceGroupName string) (result edgegateway.SharesDeleteFuture, err error)
	Get(ctx context.Context, deviceName string, name string, resourceGroupName string) (result edgegateway.Share, err error)
	ListByDataBoxEdgeDevice(ctx context.Context, deviceName string, resourceGroupName string) (result edgegateway.ShareListPage, err error)
	Refresh(ctx context.Context, deviceName string, name string, resourceGroupName string) (result edgegateway.SharesRefreshFuture, err error)
}

var _ SharesClientAPI = (*edgegateway.SharesClient)(nil)

// StorageAccountCredentialsClientAPI contains the set of methods on the StorageAccountCredentialsClient type.
type StorageAccountCredentialsClientAPI interface {
	CreateOrUpdate(ctx context.Context, deviceName string, name string, storageAccountCredential edgegateway.StorageAccountCredential, resourceGroupName string) (result edgegateway.StorageAccountCredentialsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, deviceName string, name string, resourceGroupName string) (result edgegateway.StorageAccountCredentialsDeleteFuture, err error)
	Get(ctx context.Context, deviceName string, name string, resourceGroupName string) (result edgegateway.StorageAccountCredential, err error)
	ListByDataBoxEdgeDevice(ctx context.Context, deviceName string, resourceGroupName string) (result edgegateway.StorageAccountCredentialListPage, err error)
}

var _ StorageAccountCredentialsClientAPI = (*edgegateway.StorageAccountCredentialsClient)(nil)

// TriggersClientAPI contains the set of methods on the TriggersClient type.
type TriggersClientAPI interface {
	CreateOrUpdate(ctx context.Context, deviceName string, name string, trigger edgegateway.BasicTrigger, resourceGroupName string) (result edgegateway.TriggersCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, deviceName string, name string, resourceGroupName string) (result edgegateway.TriggersDeleteFuture, err error)
	Get(ctx context.Context, deviceName string, name string, resourceGroupName string) (result edgegateway.TriggerModel, err error)
	ListByDataBoxEdgeDevice(ctx context.Context, deviceName string, resourceGroupName string) (result edgegateway.TriggerListPage, err error)
}

var _ TriggersClientAPI = (*edgegateway.TriggersClient)(nil)

// UsersClientAPI contains the set of methods on the UsersClient type.
type UsersClientAPI interface {
	CreateOrUpdate(ctx context.Context, deviceName string, name string, userParameter edgegateway.User, resourceGroupName string) (result edgegateway.UsersCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, deviceName string, name string, resourceGroupName string) (result edgegateway.UsersDeleteFuture, err error)
	Get(ctx context.Context, deviceName string, name string, resourceGroupName string) (result edgegateway.User, err error)
	ListByDataBoxEdgeDevice(ctx context.Context, deviceName string, resourceGroupName string) (result edgegateway.UserListPage, err error)
}

var _ UsersClientAPI = (*edgegateway.UsersClient)(nil)