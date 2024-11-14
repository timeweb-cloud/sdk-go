/*
Timeweb Cloud API

Testing ServersAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	openapiclient "github.com/timeweb-cloud/sdk-go"
)

func Test_openapi_ServersAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ServersAPIService AddServerIP", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serverId int32

		resp, httpRes, err := apiClient.ServersAPI.AddServerIP(context.Background(), serverId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService CloneServer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serverId int32

		resp, httpRes, err := apiClient.ServersAPI.CloneServer(context.Background(), serverId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService CreateServer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ServersAPI.CreateServer(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService CreateServerDisk", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serverId int32

		resp, httpRes, err := apiClient.ServersAPI.CreateServerDisk(context.Background(), serverId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService CreateServerDiskBackup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serverId int32
		var diskId int32

		resp, httpRes, err := apiClient.ServersAPI.CreateServerDiskBackup(context.Background(), serverId, diskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService DeleteServer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serverId int32

		resp, httpRes, err := apiClient.ServersAPI.DeleteServer(context.Background(), serverId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService DeleteServerDisk", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serverId int32
		var diskId int32

		httpRes, err := apiClient.ServersAPI.DeleteServerDisk(context.Background(), serverId, diskId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService DeleteServerDiskBackup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serverId int32
		var diskId int32
		var backupId int32

		httpRes, err := apiClient.ServersAPI.DeleteServerDiskBackup(context.Background(), serverId, diskId, backupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService DeleteServerIP", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serverId int32

		httpRes, err := apiClient.ServersAPI.DeleteServerIP(context.Background(), serverId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService GetConfigurators", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ServersAPI.GetConfigurators(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService GetOsList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ServersAPI.GetOsList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService GetServer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serverId int32

		resp, httpRes, err := apiClient.ServersAPI.GetServer(context.Background(), serverId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService GetServerDisk", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serverId int32
		var diskId int32

		resp, httpRes, err := apiClient.ServersAPI.GetServerDisk(context.Background(), serverId, diskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService GetServerDiskAutoBackupSettings", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serverId int32
		var diskId int32

		resp, httpRes, err := apiClient.ServersAPI.GetServerDiskAutoBackupSettings(context.Background(), serverId, diskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService GetServerDiskBackup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serverId int32
		var diskId int32
		var backupId int32

		resp, httpRes, err := apiClient.ServersAPI.GetServerDiskBackup(context.Background(), serverId, diskId, backupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService GetServerDiskBackups", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serverId int32
		var diskId int32

		resp, httpRes, err := apiClient.ServersAPI.GetServerDiskBackups(context.Background(), serverId, diskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService GetServerDisks", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serverId int32

		resp, httpRes, err := apiClient.ServersAPI.GetServerDisks(context.Background(), serverId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService GetServerIPs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serverId int32

		resp, httpRes, err := apiClient.ServersAPI.GetServerIPs(context.Background(), serverId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService GetServerLogs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serverId int32

		resp, httpRes, err := apiClient.ServersAPI.GetServerLogs(context.Background(), serverId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService GetServerStatistics", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serverId int32

		resp, httpRes, err := apiClient.ServersAPI.GetServerStatistics(context.Background(), serverId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService GetServers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ServersAPI.GetServers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService GetServersPresets", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ServersAPI.GetServersPresets(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService GetSoftware", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ServersAPI.GetSoftware(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService HardShutdownServer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serverId int32

		httpRes, err := apiClient.ServersAPI.HardShutdownServer(context.Background(), serverId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService ImageUnmountAndServerReload", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serverId int32

		httpRes, err := apiClient.ServersAPI.ImageUnmountAndServerReload(context.Background(), serverId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService PerformActionOnBackup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serverId int32
		var diskId int32
		var backupId int32

		httpRes, err := apiClient.ServersAPI.PerformActionOnBackup(context.Background(), serverId, diskId, backupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService PerformActionOnServer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serverId int32

		httpRes, err := apiClient.ServersAPI.PerformActionOnServer(context.Background(), serverId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService RebootServer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serverId int32

		httpRes, err := apiClient.ServersAPI.RebootServer(context.Background(), serverId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService ResetServerPassword", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serverId int32

		httpRes, err := apiClient.ServersAPI.ResetServerPassword(context.Background(), serverId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService ShutdownServer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serverId int32

		httpRes, err := apiClient.ServersAPI.ShutdownServer(context.Background(), serverId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService StartServer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serverId int32

		httpRes, err := apiClient.ServersAPI.StartServer(context.Background(), serverId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService UpdateServer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serverId int32

		resp, httpRes, err := apiClient.ServersAPI.UpdateServer(context.Background(), serverId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService UpdateServerDisk", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serverId int32
		var diskId int32

		resp, httpRes, err := apiClient.ServersAPI.UpdateServerDisk(context.Background(), serverId, diskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService UpdateServerDiskAutoBackupSettings", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serverId int32
		var diskId int32

		resp, httpRes, err := apiClient.ServersAPI.UpdateServerDiskAutoBackupSettings(context.Background(), serverId, diskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService UpdateServerDiskBackup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serverId int32
		var diskId int32
		var backupId int32

		resp, httpRes, err := apiClient.ServersAPI.UpdateServerDiskBackup(context.Background(), serverId, diskId, backupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService UpdateServerIP", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serverId int32

		resp, httpRes, err := apiClient.ServersAPI.UpdateServerIP(context.Background(), serverId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService UpdateServerNAT", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serverId int32

		httpRes, err := apiClient.ServersAPI.UpdateServerNAT(context.Background(), serverId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService UpdateServerOSBootMode", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serverId int32

		httpRes, err := apiClient.ServersAPI.UpdateServerOSBootMode(context.Background(), serverId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
