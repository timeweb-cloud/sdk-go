/*
Timeweb Cloud API

Testing VPCAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_VPCAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test VPCAPIService CreateVPC", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.VPCAPI.CreateVPC(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VPCAPIService DeleteVPC", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var vpcId string

		httpRes, err := apiClient.VPCAPI.DeleteVPC(context.Background(), vpcId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VPCAPIService GetVPC", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var vpcId string

		resp, httpRes, err := apiClient.VPCAPI.GetVPC(context.Background(), vpcId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VPCAPIService GetVPCPorts", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var vpcId string

		resp, httpRes, err := apiClient.VPCAPI.GetVPCPorts(context.Background(), vpcId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VPCAPIService GetVPCServices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var vpcId string

		resp, httpRes, err := apiClient.VPCAPI.GetVPCServices(context.Background(), vpcId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VPCAPIService GetVPCs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.VPCAPI.GetVPCs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VPCAPIService UpdateVPCs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var vpcId string

		resp, httpRes, err := apiClient.VPCAPI.UpdateVPCs(context.Background(), vpcId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
