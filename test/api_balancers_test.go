/*
Timeweb Cloud API

Testing BalancersAPIService

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

func Test_openapi_BalancersAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BalancersAPIService AddIPsToBalancer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var balancerId int32

		httpRes, err := apiClient.BalancersAPI.AddIPsToBalancer(context.Background(), balancerId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BalancersAPIService CreateBalancer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.BalancersAPI.CreateBalancer(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BalancersAPIService CreateBalancerRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var balancerId int32

		resp, httpRes, err := apiClient.BalancersAPI.CreateBalancerRule(context.Background(), balancerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BalancersAPIService DeleteBalancer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var balancerId int32

		resp, httpRes, err := apiClient.BalancersAPI.DeleteBalancer(context.Background(), balancerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BalancersAPIService DeleteBalancerRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var balancerId int32
		var ruleId int32

		httpRes, err := apiClient.BalancersAPI.DeleteBalancerRule(context.Background(), balancerId, ruleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BalancersAPIService DeleteIPsFromBalancer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var balancerId int32

		httpRes, err := apiClient.BalancersAPI.DeleteIPsFromBalancer(context.Background(), balancerId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BalancersAPIService GetBalancer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var balancerId int32

		resp, httpRes, err := apiClient.BalancersAPI.GetBalancer(context.Background(), balancerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BalancersAPIService GetBalancerIPs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var balancerId int32

		resp, httpRes, err := apiClient.BalancersAPI.GetBalancerIPs(context.Background(), balancerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BalancersAPIService GetBalancerRules", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var balancerId int32

		resp, httpRes, err := apiClient.BalancersAPI.GetBalancerRules(context.Background(), balancerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BalancersAPIService GetBalancers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.BalancersAPI.GetBalancers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BalancersAPIService GetBalancersPresets", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.BalancersAPI.GetBalancersPresets(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BalancersAPIService UpdateBalancer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var balancerId int32

		resp, httpRes, err := apiClient.BalancersAPI.UpdateBalancer(context.Background(), balancerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BalancersAPIService UpdateBalancerRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var balancerId int32
		var ruleId int32

		resp, httpRes, err := apiClient.BalancersAPI.UpdateBalancerRule(context.Background(), balancerId, ruleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
