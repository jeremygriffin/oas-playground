/*
MilMove Prime API

Testing MtoShipmentAPIService

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

func Test_openapi_MtoShipmentAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MtoShipmentAPIService CreateMTOAgent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mtoShipmentID string

		resp, httpRes, err := apiClient.MtoShipmentAPI.CreateMTOAgent(context.Background(), mtoShipmentID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MtoShipmentAPIService CreateMTOShipment", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MtoShipmentAPI.CreateMTOShipment(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MtoShipmentAPIService CreateSITExtension", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mtoShipmentID string

		resp, httpRes, err := apiClient.MtoShipmentAPI.CreateSITExtension(context.Background(), mtoShipmentID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MtoShipmentAPIService DeleteMTOShipment", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mtoShipmentID string

		httpRes, err := apiClient.MtoShipmentAPI.DeleteMTOShipment(context.Background(), mtoShipmentID).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MtoShipmentAPIService UpdateMTOAgent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mtoShipmentID string
		var agentID string

		resp, httpRes, err := apiClient.MtoShipmentAPI.UpdateMTOAgent(context.Background(), mtoShipmentID, agentID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MtoShipmentAPIService UpdateMTOShipment", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mtoShipmentID string

		resp, httpRes, err := apiClient.MtoShipmentAPI.UpdateMTOShipment(context.Background(), mtoShipmentID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MtoShipmentAPIService UpdateMTOShipmentAddress", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mtoShipmentID string
		var addressID string

		resp, httpRes, err := apiClient.MtoShipmentAPI.UpdateMTOShipmentAddress(context.Background(), mtoShipmentID, addressID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MtoShipmentAPIService UpdateMTOShipmentStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mtoShipmentID string

		resp, httpRes, err := apiClient.MtoShipmentAPI.UpdateMTOShipmentStatus(context.Background(), mtoShipmentID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MtoShipmentAPIService UpdateReweigh", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mtoShipmentID string
		var reweighID string

		resp, httpRes, err := apiClient.MtoShipmentAPI.UpdateReweigh(context.Background(), mtoShipmentID, reweighID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MtoShipmentAPIService UpdateShipmentDestinationAddress", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mtoShipmentID string

		resp, httpRes, err := apiClient.MtoShipmentAPI.UpdateShipmentDestinationAddress(context.Background(), mtoShipmentID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
