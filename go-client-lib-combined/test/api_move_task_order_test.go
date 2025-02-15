/*
MilMove Prime API

Testing MoveTaskOrderAPIService

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

func Test_openapi_MoveTaskOrderAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MoveTaskOrderAPIService CreateExcessWeightRecord", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var moveTaskOrderID string

		resp, httpRes, err := apiClient.MoveTaskOrderAPI.CreateExcessWeightRecord(context.Background(), moveTaskOrderID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MoveTaskOrderAPIService DownloadMoveOrder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var locator string

		resp, httpRes, err := apiClient.MoveTaskOrderAPI.DownloadMoveOrder(context.Background(), locator).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MoveTaskOrderAPIService GetMoveTaskOrder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var moveID string

		resp, httpRes, err := apiClient.MoveTaskOrderAPI.GetMoveTaskOrder(context.Background(), moveID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MoveTaskOrderAPIService ListMoves", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MoveTaskOrderAPI.ListMoves(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MoveTaskOrderAPIService UpdateMTOPostCounselingInformation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var moveTaskOrderID string

		resp, httpRes, err := apiClient.MoveTaskOrderAPI.UpdateMTOPostCounselingInformation(context.Background(), moveTaskOrderID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
