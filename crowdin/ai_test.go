package crowdin

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/crowdin/crowdin-api-client-go/crowdin/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAIService_GenerateFineTuningDataset(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/users/1/ai/prompts/2/fine-tuning/datasets"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, path)
		testJSONBody(t, r, `{
			"projectIds": [1],
			"tmIds": [2],
			"purpose": "training",
			"dateFrom": "2024-09-23T11:26:54+00:00",
			"dateTo": "2024-09-23T11:26:54+00:00",
			"maxFileSize": 1000,
			"minExamplesCount": 1,
			"maxExamplesCount": 2
		}`)

		fmt.Fprint(w, `{
			"data": {
				"identifier": "50fb3506-4127-4ba8-8296-f97dc7e3e0c3",
				"status": "finished",
				"progress": 100,
				"attributes": {
					"projectIds": [1],
					"tmIds": [2],
					"purpose": "training",
					"dateFrom": "2024-09-23T11:26:54+00:00",
					"dateTo": "2024-09-23T11:26:54+00:00",
					"maxFileSize": 1000,
					"minExamplesCount": 1,
					"maxExamplesCount": 2
				},
				"createdAt": "2024-09-23T11:26:54+00:00",
				"updatedAt": "2024-09-23T11:26:54+00:00",
				"startedAt": "2024-09-23T11:26:54+00:00",
				"finishedAt": "2024-09-23T11:26:54+00:00"
			}
		}`)
	})

	req := &model.FineTuningDatasetAttributes{
		ProjectIDs:       []int{1},
		TMIDs:            []int{2},
		Purpose:          "training",
		DateFrom:         "2024-09-23T11:26:54+00:00",
		DateTo:           "2024-09-23T11:26:54+00:00",
		MaxFileSize:      1000,
		MinExamplesCount: 1,
		MaxExamplesCount: 2,
	}
	dataset, resp, err := client.AI.GenerateFineTuningDataset(context.Background(), 2, 1, req)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.FineTuningDataset{
		Identifier: "50fb3506-4127-4ba8-8296-f97dc7e3e0c3",
		Status:     "finished",
		Progress:   100,
		Attributes: &model.FineTuningDatasetAttributes{
			ProjectIDs:       []int{1},
			TMIDs:            []int{2},
			Purpose:          "training",
			DateFrom:         "2024-09-23T11:26:54+00:00",
			DateTo:           "2024-09-23T11:26:54+00:00",
			MaxFileSize:      1000,
			MinExamplesCount: 1,
			MaxExamplesCount: 2,
		},
		CreatedAt:  "2024-09-23T11:26:54+00:00",
		UpdatedAt:  "2024-09-23T11:26:54+00:00",
		StartedAt:  "2024-09-23T11:26:54+00:00",
		FinishedAt: "2024-09-23T11:26:54+00:00",
	}
	assert.Equal(t, expected, dataset)
}

func TestAIService_GetFineTuningDatasetGenerationStatus(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/ai/prompts/2/fine-tuning/datasets/50fb3506-4127-4ba8-8296-f97dc7e3e0c3"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		fmt.Fprint(w, `{
			"data": {
				"identifier": "50fb3506-4127-4ba8-8296-f97dc7e3e0c3",
				"status": "finished",
				"progress": 100,
				"attributes": {
					"projectIds": [1],
					"tmIds": [2],
					"purpose": "training",
					"dateFrom": "2024-09-23T11:26:54+00:00",
					"dateTo": "2024-09-23T11:26:54+00:00",
					"maxFileSize": 1000,
					"minExamplesCount": 1,
					"maxExamplesCount": 2
				},
				"createdAt": "2024-09-23T11:26:54+00:00",
				"updatedAt": "2024-09-23T11:26:54+00:00",
				"startedAt": "2024-09-23T11:26:54+00:00",
				"finishedAt": "2024-09-23T11:26:54+00:00"
			}
		}`)
	})

	status, resp, err := client.AI.GetFineTuningDatasetGenerationStatus(context.Background(), 2, "50fb3506-4127-4ba8-8296-f97dc7e3e0c3", 0)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.FineTuningDataset{
		Identifier: "50fb3506-4127-4ba8-8296-f97dc7e3e0c3",
		Status:     "finished",
		Progress:   100,
		Attributes: &model.FineTuningDatasetAttributes{
			ProjectIDs:       []int{1},
			TMIDs:            []int{2},
			Purpose:          "training",
			DateFrom:         "2024-09-23T11:26:54+00:00",
			DateTo:           "2024-09-23T11:26:54+00:00",
			MaxFileSize:      1000,
			MinExamplesCount: 1,
			MaxExamplesCount: 2,
		},
		CreatedAt:  "2024-09-23T11:26:54+00:00",
		UpdatedAt:  "2024-09-23T11:26:54+00:00",
		StartedAt:  "2024-09-23T11:26:54+00:00",
		FinishedAt: "2024-09-23T11:26:54+00:00",
	}
	assert.Equal(t, expected, status)
}

func TestAIService_DownloadFineTuningDataset(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/ai/prompts/2/fine-tuning/datasets/50fb3506-4127-4ba8-8296-f97dc7e3e0c3/download"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		fmt.Fprint(w, `{
			"data": {
				"url": "https://production-enterprise-importer.downloads.crowdin.com/992000002/2/14.xliff",
				"expireIn": "2024-09-20T10:31:21+00:00"
			}
		}`)
	})

	downloadLink, resp, err := client.AI.DownloadFineTuningDataset(context.Background(), 2, "50fb3506-4127-4ba8-8296-f97dc7e3e0c3", 0)
	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "https://production-enterprise-importer.downloads.crowdin.com/992000002/2/14.xliff", downloadLink.URL)
	assert.Equal(t, "2024-09-20T10:31:21+00:00", downloadLink.ExpireIn)
}

func TestAIService_ListFineTuningJobs(t *testing.T) {
	tests := []struct {
		name          string
		opts          *model.FineTuningJobsListOptions
		expectedQuery string
	}{
		{
			name: "nil options",
			opts: nil,
		},
		{
			name: "empty options",
			opts: &model.FineTuningJobsListOptions{},
		},
		{
			name: "with options",
			opts: &model.FineTuningJobsListOptions{
				Statuses:    []string{"finished,in_progress"},
				OrderBy:     "createdAt",
				ListOptions: model.ListOptions{Offset: 1, Limit: 25},
			},
			expectedQuery: "?limit=25&offset=1&orderBy=createdAt&statuses=finished%2Cin_progress",
		},
	}

	client, mux, teardown := setupClient()
	defer teardown()

	for userID, tt := range tests {
		userID++
		path := fmt.Sprintf("/api/v2/users/%d/ai/prompts/fine-tuning/jobs", userID)
		mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			testURL(t, r, path+tt.expectedQuery)

			fmt.Fprint(w, `{
				"data": [
					{
						"data": {
							"identifier": "50fb3506-4127-4ba8-8296-f97dc7e3e0c3",
							"status": "finished"
						}
					},
					{
						"data": {
							"identifier": "50fb3506-4127-4ba8-8296-f97dc7e3e0c4",
							"status": "in_progress"
						}			
					}
				],
				"pagination": {
					"offset": 2,
					"limit": 25
				}
			}`)
		})

		jobs, resp, err := client.AI.ListFineTuningJobs(context.Background(), userID, tt.opts)
		require.NoError(t, err)

		expected := []*model.FineTuningJob{
			{Identifier: "50fb3506-4127-4ba8-8296-f97dc7e3e0c3", Status: "finished"},
			{Identifier: "50fb3506-4127-4ba8-8296-f97dc7e3e0c4", Status: "in_progress"},
		}
		assert.Equal(t, expected, jobs)

		assert.Equal(t, 2, resp.Pagination.Offset)
		assert.Equal(t, 25, resp.Pagination.Limit)
	}
}

func TestAIService_ListFineTuningJobs_invalidJSON(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/ai/prompts/fine-tuning/jobs", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `invalid json`)
	})

	prompts, _, err := client.AI.ListFineTuningJobs(context.Background(), 0, nil)
	require.Error(t, err)
	assert.Nil(t, prompts)
}

func TestAIService_ListFineTuningEvents(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	path := "/api/v2/ai/prompts/1/fine-tuning/jobs/50fb3506-4127-4ba8-8296-f97dc7e3e0c3/events"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		fmt.Fprint(w, `{
				"data": [
					{
						"data": {
							"id": "ftevent-0HOW0hqwQ7G1b70qOC8S7UsT"
						}
					},
					{
						"data": {
							"id": "ftevent-0HOW0hqwQ7G1b70qOC8S7UsQ"
						}
					}
				],
				"pagination": {
					"offset": 2,
					"limit": 25
				}
			}`)
	})

	jobs, resp, err := client.AI.ListFineTuningEvents(context.Background(), 1, "50fb3506-4127-4ba8-8296-f97dc7e3e0c3", 0)
	require.NoError(t, err)

	expected := []*model.FineTuningEvent{
		{ID: "ftevent-0HOW0hqwQ7G1b70qOC8S7UsT"},
		{ID: "ftevent-0HOW0hqwQ7G1b70qOC8S7UsQ"},
	}
	assert.Equal(t, expected, jobs)

	assert.Equal(t, 2, resp.Pagination.Offset)
	assert.Equal(t, 25, resp.Pagination.Limit)
}

func TestAIService_ListFineTuningEvents_invalidJSON(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/users/1/ai/prompts/1/fine-tuning/jobs/50fb3506-4127-4ba8-8296-f97dc7e3e0c3/events", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `invalid json`)
	})

	prompts, _, err := client.AI.ListFineTuningEvents(context.Background(), 1, "50fb3506-4127-4ba8-8296-f97dc7e3e0c3", 1)
	require.Error(t, err)
	assert.Nil(t, prompts)
}

func TestAIService_CreateFineTuningJob(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/users/1/ai/prompts/2/fine-tuning/jobs"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, path)
		testJSONBody(t, r, `{
			"dryRun": false,
			"hyperparameters": {
				"batchSize": 100,
				"learningRateMultiplier": 1,
				"nEpochs": 10
			},
			"trainingOptions": {
				"projectIds": [1],
				"tmIds": [2],
				"dateFrom": "2024-09-23T11:26:54+00:00",
				"dateTo": "2024-09-23T11:26:54+00:00",
				"maxFileSize": 100,
				"minExamplesCount": 1,
				"maxExamplesCount": 10
			},
			"validationOptions": {
				"projectIds": [2],
				"tmIds": [3],
				"dateFrom": "2024-09-23T11:26:54+00:00",
				"dateTo": "2024-09-23T11:26:54+00:00",
				"maxFileSize": 2000,
				"minExamplesCount": 2,
				"maxExamplesCount": 3
			}
		}`)
		fmt.Fprint(w, `{
			"data": {
				"identifier": "50fb3506-4127-4ba8-8296-f97dc7e3e0c3",
				"status": "finished",
				"progress": 100,
				"attributes": {
					"dryRun": true,
					"aiPromptId": 0,
					"hyperparameters": {
						"batchSize": 100,
						"learningRateMultiplier": 1,
						"nEpochs": 10
					},
					"trainingOptions": {
						"projectIds": [1],
						"tmIds": [2],
						"dateFrom": "2024-09-23T11:26:54+00:00",
						"dateTo": "2024-09-23T11:26:54+00:00",
						"maxFileSize": 100,
						"minExamplesCount": 1,
						"maxExamplesCount": 10
					},
					"validationOptions": {
						"projectIds": [2],
						"tmIds": [3],
						"dateFrom": "2024-09-23T11:26:54+00:00",
						"dateTo": "2024-09-23T11:26:54+00:00",
						"maxFileSize": 2000,
						"minExamplesCount": 2,
						"maxExamplesCount": 3
					},
					"baseModel": "gpt-4o-mini-2024-07-18",
					"fineTunedModel": "ft:gpt-4o-mini-2024-07-18:2024-08-12:9vQgFOfp",
					"trainedTokensCount": 0,
					"trainingDatasetUrl": "https://crowdin-tmp.s3.eu-central-1.amazonaws.com/72057531/ai_fine_tuning/0fc00e9e-3a80-49d1-a8cb-39c9a81c0ae2/training-dataset.jsonl",
					"validationDatasetUrl": "https://crowdin-tmp.s3.eu-central-1.amazonaws.com/72057531/ai_fine_tuning/0fc00e9e-3a80-49d1-a8cb-39c9a81c0ae2/validation-dataset.jsonl",
					"metadata": {
						"cost": 0.1,
						"costCurrency": "USD"
					}
				},
				"createdAt": "2024-09-23T11:26:54+00:00",
				"updatedAt": "2024-09-23T11:26:54+00:00",
				"startedAt": "2024-09-23T11:26:54+00:00",
				"finishedAt": "2024-09-23T11:26:54+00:00"
			}
		}`)
	})

	req := &model.FineTuningJobCreateRequest{
		DryRun: ToPtr(false),
		TrainingOptions: &model.FineTuningJobOptions{
			ProjectIDs:       []int{1},
			TMIDs:            []int{2},
			DateFrom:         "2024-09-23T11:26:54+00:00",
			DateTo:           "2024-09-23T11:26:54+00:00",
			MaxFileSize:      100,
			MinExamplesCount: 1,
			MaxExamplesCount: 10,
		},
		ValidationOptions: &model.FineTuningJobOptions{
			ProjectIDs:       []int{2},
			TMIDs:            []int{3},
			DateFrom:         "2024-09-23T11:26:54+00:00",
			DateTo:           "2024-09-23T11:26:54+00:00",
			MaxFileSize:      2000,
			MinExamplesCount: 2,
			MaxExamplesCount: 3,
		},
		Hyperparameters: &model.FineTuningJobHyperparameters{
			BatchSize:              100,
			LearningRateMultiplier: 1,
			NEpochs:                10,
		},
	}
	job, resp, err := client.AI.client.AI.CreateFineTuningJob(context.Background(), 2, 1, req)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.FineTuningJob{
		Identifier: "50fb3506-4127-4ba8-8296-f97dc7e3e0c3",
		Status:     "finished",
		Progress:   100,
		Attributes: &model.FineTuningJobAttributes{
			DryRun:     true,
			AIPromptID: 0,
			Hyperparameters: &model.FineTuningJobHyperparameters{
				BatchSize:              100,
				LearningRateMultiplier: 1,
				NEpochs:                10,
			},
			TrainingOptions: &model.FineTuningJobOptions{
				ProjectIDs:       []int{1},
				TMIDs:            []int{2},
				DateFrom:         "2024-09-23T11:26:54+00:00",
				DateTo:           "2024-09-23T11:26:54+00:00",
				MaxFileSize:      100,
				MinExamplesCount: 1,
				MaxExamplesCount: 10,
			},
			ValidationOptions: &model.FineTuningJobOptions{
				ProjectIDs:       []int{2},
				TMIDs:            []int{3},
				DateFrom:         "2024-09-23T11:26:54+00:00",
				DateTo:           "2024-09-23T11:26:54+00:00",
				MaxFileSize:      2000,
				MinExamplesCount: 2,
				MaxExamplesCount: 3,
			},
			BaseModel:            "gpt-4o-mini-2024-07-18",
			FineTunedModel:       "ft:gpt-4o-mini-2024-07-18:2024-08-12:9vQgFOfp",
			TrainedTokensCount:   0,
			TrainingDatasetURL:   "https://crowdin-tmp.s3.eu-central-1.amazonaws.com/72057531/ai_fine_tuning/0fc00e9e-3a80-49d1-a8cb-39c9a81c0ae2/training-dataset.jsonl",
			ValidationDatasetURL: "https://crowdin-tmp.s3.eu-central-1.amazonaws.com/72057531/ai_fine_tuning/0fc00e9e-3a80-49d1-a8cb-39c9a81c0ae2/validation-dataset.jsonl",
			Metadata: &model.FineTuningJobMetadata{
				Cost:         0.1,
				CostCurrency: "USD",
			},
		},
		CreatedAt:  "2024-09-23T11:26:54+00:00",
		UpdatedAt:  "2024-09-23T11:26:54+00:00",
		StartedAt:  "2024-09-23T11:26:54+00:00",
		FinishedAt: "2024-09-23T11:26:54+00:00",
	}
	assert.Equal(t, expected, job)
}

func TestAIService_GetFineTuningJobStatus(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/ai/prompts/2/fine-tuning/jobs/50fb3506-4127-4ba8-8296-f97dc7e3e0c3"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		fmt.Fprint(w, `{
			"data": {
				"identifier": "50fb3506-4127-4ba8-8296-f97dc7e3e0c3",
				"status": "finished",
				"progress": 100,
				"attributes": {
					"dryRun": true,
					"aiPromptId": 0,
					"hyperparameters": {
						"batchSize": 100,
						"learningRateMultiplier": 1,
						"nEpochs": 10
					},
					"trainingOptions": {
						"projectIds": [1],
						"tmIds": [2],
						"dateFrom": "2024-09-23T11:26:54+00:00",
						"dateTo": "2024-09-23T11:26:54+00:00",
						"maxFileSize": 100,
						"minExamplesCount": 1,
						"maxExamplesCount": 10
					},
					"validationOptions": {
						"projectIds": [2],
						"tmIds": [3],
						"dateFrom": "2024-09-23T11:26:54+00:00",
						"dateTo": "2024-09-23T11:26:54+00:00",
						"maxFileSize": 2000,
						"minExamplesCount": 2,
						"maxExamplesCount": 3
					},
					"baseModel": "gpt-4o-mini-2024-07-18",
					"fineTunedModel": "ft:gpt-4o-mini-2024-07-18:2024-08-12:9vQgFOfp",
					"trainedTokensCount": 0,
					"trainingDatasetUrl": "https://crowdin-tmp.s3.eu-central-1.amazonaws.com/72057531/ai_fine_tuning/0fc00e9e-3a80-49d1-a8cb-39c9a81c0ae2/training-dataset.jsonl",
					"validationDatasetUrl": "https://crowdin-tmp.s3.eu-central-1.amazonaws.com/72057531/ai_fine_tuning/0fc00e9e-3a80-49d1-a8cb-39c9a81c0ae2/validation-dataset.jsonl",
					"metadata": {
						"cost": 0.1,
						"costCurrency": "USD"
					}
				},
				"createdAt": "2024-09-23T11:26:54+00:00",
				"updatedAt": "2024-09-23T11:26:54+00:00",
				"startedAt": "2024-09-23T11:26:54+00:00",
				"finishedAt": "2024-09-23T11:26:54+00:00"
			}
		}`)
	})

	status, resp, err := client.AI.GetFineTuningJobStatus(context.Background(), 2, "50fb3506-4127-4ba8-8296-f97dc7e3e0c3", 0)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.FineTuningJob{
		Identifier: "50fb3506-4127-4ba8-8296-f97dc7e3e0c3",
		Status:     "finished",
		Progress:   100,
		Attributes: &model.FineTuningJobAttributes{
			DryRun:     true,
			AIPromptID: 0,
			Hyperparameters: &model.FineTuningJobHyperparameters{
				BatchSize:              100,
				LearningRateMultiplier: 1,
				NEpochs:                10,
			},
			TrainingOptions: &model.FineTuningJobOptions{
				ProjectIDs:       []int{1},
				TMIDs:            []int{2},
				DateFrom:         "2024-09-23T11:26:54+00:00",
				DateTo:           "2024-09-23T11:26:54+00:00",
				MaxFileSize:      100,
				MinExamplesCount: 1,
				MaxExamplesCount: 10,
			},
			ValidationOptions: &model.FineTuningJobOptions{
				ProjectIDs:       []int{2},
				TMIDs:            []int{3},
				DateFrom:         "2024-09-23T11:26:54+00:00",
				DateTo:           "2024-09-23T11:26:54+00:00",
				MaxFileSize:      2000,
				MinExamplesCount: 2,
				MaxExamplesCount: 3,
			},
			BaseModel:            "gpt-4o-mini-2024-07-18",
			FineTunedModel:       "ft:gpt-4o-mini-2024-07-18:2024-08-12:9vQgFOfp",
			TrainedTokensCount:   0,
			TrainingDatasetURL:   "https://crowdin-tmp.s3.eu-central-1.amazonaws.com/72057531/ai_fine_tuning/0fc00e9e-3a80-49d1-a8cb-39c9a81c0ae2/training-dataset.jsonl",
			ValidationDatasetURL: "https://crowdin-tmp.s3.eu-central-1.amazonaws.com/72057531/ai_fine_tuning/0fc00e9e-3a80-49d1-a8cb-39c9a81c0ae2/validation-dataset.jsonl",
			Metadata: &model.FineTuningJobMetadata{
				Cost:         0.1,
				CostCurrency: "USD",
			},
		},
		CreatedAt:  "2024-09-23T11:26:54+00:00",
		UpdatedAt:  "2024-09-23T11:26:54+00:00",
		StartedAt:  "2024-09-23T11:26:54+00:00",
		FinishedAt: "2024-09-23T11:26:54+00:00",
	}
	assert.Equal(t, expected, status)
}

func TestAIService_GetPrompt(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/ai/prompts/2"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		fmt.Fprint(w, `{
			"data": {
				"id": 2,
				"name": "Pre-translate prompt",
				"action": "pre_translate",
				"aiProviderId": 2,
				"aiModelId": "gpt-3.5-turbo-instruct",
				"isEnabled": true,
				"enabledProjectIds": [1],
				"config": {
					"mode": "basic",
					"snippets": ["\u0025custom:companyDescription\u0025"],
					"otherLanguageTranslations": {
						"isEnabled": true,
						"languageIds": ["uk"]
					},
					"glossaryTerms": true,
					"tmSuggestions": true,
					"fileContext": true,
					"generateFileSummary": true,
					"screenshots": true,
					"projectContext": true,
					"organizationContext": true,
					"siblingsStrings": true,
					"retryOnQaIssues": true
				},
				"promptPreview": "Translate the text",
				"isFineTuningAvailable": true,
				"createdBy": 1,
				"updatedBy": 2,
				"lastUsedBy": null,
				"lastUsedAt": null,
				"usageCount": 5,
				"createdAt": "2023-09-20T11:11:05+00:00",
				"updatedAt": "2023-09-20T12:22:20+00:00"
			}
		}`)
	})

	prompt, resp, err := client.AI.GetPrompt(context.Background(), 2, 0)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.Prompt{
		ID:                2,
		Name:              "Pre-translate prompt",
		Action:            "pre_translate",
		AIProviderID:      2,
		AIModelID:         "gpt-3.5-turbo-instruct",
		IsEnabled:         true,
		EnabledProjectIDs: []int{1},
		Config: model.PromptConfig{
			Mode:     model.ModeBasic,
			Snippets: []string{"%custom:companyDescription%"},
			OtherLanguageTranslations: &model.OtherLanguageTranslations{
				IsEnabled:   ToPtr(true),
				LanguageIDs: []string{"uk"},
			},
			GlossaryTerms:       ToPtr(true),
			TMSuggestions:       ToPtr(true),
			FileContext:         ToPtr(true),
			GenerateFileSummary: ToPtr(true),
			Screenshots:         ToPtr(true),
			ProjectContext:      ToPtr(true),
			OrganizationContext: ToPtr(true),
			SiblingsStrings:     ToPtr(true),
			RetryOnQAIssues:     ToPtr(true),
		},
		PromptPreview:         ToPtr("Translate the text"),
		IsFineTuningAvailable: true,
		CreatedBy:             ToPtr(1),
		UpdatedBy:             ToPtr(2),
		UsageCount:            5,
		CreatedAt:             "2023-09-20T11:11:05+00:00",
		UpdatedAt:             "2023-09-20T12:22:20+00:00",
	}
	assert.Equal(t, expected, prompt)
}

func TestAIService_ListPrompts(t *testing.T) {
	tests := []struct {
		name          string
		opts          *model.AIPromtsListOptions
		expectedQuery string
	}{
		{
			name: "nil options",
			opts: nil,
		},
		{
			name: "empty options",
			opts: &model.AIPromtsListOptions{},
		},
		{
			name: "with options",
			opts: &model.AIPromtsListOptions{
				ProjectID:   1,
				Action:      model.ActionQACheck,
				ListOptions: model.ListOptions{Offset: 1, Limit: 25},
			},
			expectedQuery: "?action=qa_check&limit=25&offset=1&projectId=1",
		},
	}

	client, mux, teardown := setupClient()
	defer teardown()

	for userID, tt := range tests {
		userID++
		path := fmt.Sprintf("/api/v2/users/%d/ai/prompts", userID)
		mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			testURL(t, r, path+tt.expectedQuery)

			fmt.Fprint(w, `{
				"data": [
					{
						"data": {
							"id": 2
						}
					},
					{
						"data": {
							"id": 4
						}
					}
				],
				"pagination": {
					"offset": 10,
					"limit": 25
				}
			}`)
		})

		prompts, resp, err := client.AI.ListPrompts(context.Background(), userID, tt.opts)
		require.NoError(t, err)

		expected := []*model.Prompt{{ID: 2}, {ID: 4}}
		assert.Equal(t, expected, prompts)

		assert.Equal(t, 10, resp.Pagination.Offset)
		assert.Equal(t, 25, resp.Pagination.Limit)
	}
}

func TestAIService_ListPrompts_invalidJSON(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/users/1/ai/prompts", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `invalid json`)
	})

	prompts, _, err := client.AI.ListPrompts(context.Background(), 1, nil)
	require.Error(t, err)
	assert.Nil(t, prompts)
}

func TestAIService_AddPrompt(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/users/1/ai/prompts"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, path)
		testJSONBody(t, r, `{
			"name":"Pre-translate prompt",
			"action":"pre_translate",
			"aiProviderId":1,
			"aiModelId":"gpt-3.5-turbo-instruct",
			"config":{
				"mode":"basic"
			}
		}`)

		fmt.Fprint(w, `{
			"data": {
				"id": 2,
				"name": "Pre-translate prompt",
				"action": "pre_translate"
			}
		}`)
	})

	req := &model.PromptAddRequest{
		Name:         "Pre-translate prompt",
		Action:       "pre_translate",
		AIProviderID: 1,
		AIModelID:    "gpt-3.5-turbo-instruct",
		Config: model.PromptConfig{
			Mode: model.ModeBasic,
		},
	}
	prompt, resp, err := client.AI.AddPrompt(context.Background(), 1, req)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.Prompt{
		ID:     2,
		Name:   "Pre-translate prompt",
		Action: "pre_translate",
	}
	assert.Equal(t, expected, prompt)
}

func TestAIService_EditPrompt(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/users/1/ai/prompts/2"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		testURL(t, r, path)
		testBody(t, r, `[{"op":"replace","path":"/name","value":"Pre-translate prompt"}]`+"\n")

		fmt.Fprint(w, `{
			"data": {
				"id": 2,
				"name": "Pre-translate prompt",
				"action": "pre_translate"
			}
		}`)
	})

	req := []*model.UpdateRequest{
		{
			Op:    "replace",
			Path:  "/name",
			Value: "Pre-translate prompt",
		},
	}
	prompt, resp, err := client.AI.EditPrompt(context.Background(), 2, 1, req)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.Prompt{
		ID:     2,
		Name:   "Pre-translate prompt",
		Action: "pre_translate",
	}
	assert.Equal(t, expected, prompt)
}

func TestAIService_DeletePrompt(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	var userID = 1
	t.Run("path with user id", func(t *testing.T) {
		const path = "/api/v2/users/1/ai/prompts/2"
		mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodDelete)
			testURL(t, r, path)
			w.WriteHeader(http.StatusNoContent)
		})

		resp, err := client.AI.DeletePrompt(context.Background(), 2, userID)
		require.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, resp.StatusCode)
	})

	userID = 0
	t.Run("path without user id", func(t *testing.T) {
		const path = "/api/v2/ai/prompts/2"
		mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodDelete)
			testURL(t, r, path)
			w.WriteHeader(http.StatusNoContent)
		})

		resp, err := client.AI.DeletePrompt(context.Background(), 2, userID)
		require.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, resp.StatusCode)
	})
}

func TestAIService_GetProvider(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/ai/providers/2"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		fmt.Fprint(w, `{
			"data": {
				"id": 2,
				"name": "OpenAI",
				"type": "open_ai",
				"credentials": {
					"apiKey": "string",
					"baseUrl": "https://proxy.example.com",
					"headers": {
						"X-Custom": "value"
					},
					"sendCustomHeadersOnly": false
				},
				"config": {
					"actionRules": [
						{
							"action": "pre_translate",
							"availableAiModelIds": [
								"gpt-3.5-turbo-instruct"
							]
						}
					]
				},
				"isEnabled": true,
				"useSystemCredentials": false,
				"createdAt": "2023-09-20T11:11:05+00:00",
				"updatedAt": "2023-09-20T12:22:20+00:00",
				"promptsCount": 3
			}
		}`)
	})

	provider, resp, err := client.AI.GetProvider(context.Background(), 2, 0)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.Provider{
		ID:   2,
		Name: "OpenAI",
		Type: model.OpenAI,
		Credentials: map[string]string{
			"apiKey":  "string",
			"baseUrl": "https://proxy.example.com",
		},
		RawCredentials: map[string]any{
			"apiKey":                "string",
			"baseUrl":               "https://proxy.example.com",
			"headers":               map[string]any{"X-Custom": "value"},
			"sendCustomHeadersOnly": false,
		},
		Config: model.ProviderConfig{
			ActionRules: []model.ActionRule{
				{
					Action:              model.ActionPreTranslate,
					AvailableAIModelIDs: []string{"gpt-3.5-turbo-instruct"},
				},
			},
		},
		IsEnabled:            true,
		UseSystemCredentials: false,
		CreatedAt:            "2023-09-20T11:11:05+00:00",
		UpdatedAt:            "2023-09-20T12:22:20+00:00",
		PromptsCount:         3,
	}
	assert.Equal(t, expected, provider)
}

func TestAIService_ListProviders(t *testing.T) {
	tests := []struct {
		name          string
		opts          *model.ListOptions
		expectedQuery string
	}{
		{
			name: "nil options",
			opts: nil,
		},
		{
			name: "empty options",
			opts: &model.ListOptions{},
		},
		{
			name:          "with options",
			opts:          &model.ListOptions{Offset: 1, Limit: 25},
			expectedQuery: "?limit=25&offset=1",
		},
	}

	client, mux, teardown := setupClient()
	defer teardown()

	for userID, tt := range tests {
		userID++
		path := fmt.Sprintf("/api/v2/users/%d/ai/providers", userID)
		mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			testURL(t, r, path+tt.expectedQuery)

			fmt.Fprint(w, `{
				"data": [
					{
						"data": {
							"id": 2
						}
					},
					{
						"data": {
							"id": 4
						}
					}
				],
				"pagination": {
					"offset": 10,
					"limit": 25
				}
			}`)
		})

		providers, resp, err := client.AI.ListProviders(context.Background(), userID, tt.opts)
		require.NoError(t, err)

		expected := []*model.Provider{{ID: 2}, {ID: 4}}
		assert.Equal(t, expected, providers)

		assert.Equal(t, 10, resp.Pagination.Offset)
		assert.Equal(t, 25, resp.Pagination.Limit)
	}
}

func TestAIService_ListProviders_invalidJSON(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/users/1/ai/providers", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `invalid json`)
	})

	providers, _, err := client.AI.ListProviders(context.Background(), 1, nil)
	require.Error(t, err)
	assert.Nil(t, providers)
}

func TestAIService_AddProvider(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/users/1/ai/providers"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, path)
		testJSONBody(t, r, `{
			"name":"OpenAI",
			"type":"open_ai",
			"credentials":{
				"resourceName":"resourceName",
				"apiKey":"apiKey123",
				"deploymentName":"deploymentName",
				"apiVersion":"v.1.0"
			},
			"config":{
				"actionRules":null
			},
			"isEnabled":true,
			"useSystemCredentials":false
		}`)

		fmt.Fprint(w, `{
			"data": {
				"id": 2,
				"name": "OpenAI",
				"type": "open_ai",
				"credentials": {
					"resourceName": "resourceName",
					"apiKey": "apiKey123",
					"deploymentName": "deploymentName",
					"apiVersion": "v.1.0"
				},
				"config": {
					"actionRules": [
						{
							"action": "pre_translate",
							"availableAiModelIds": [
								"gpt-3.5-turbo-instruct"
							]
						}
					]
				},
				"isEnabled": true,
				"useSystemCredentials": false,
				"createdAt": "2023-09-20T11:11:05+00:00",
				"updatedAt": "2023-09-20T12:22:20+00:00"
			}
		}`)
	})

	req := &model.ProviderAddRequest{
		Name: "OpenAI",
		Type: model.OpenAI,
		Credentials: map[string]string{
			"resourceName":   "resourceName",
			"apiKey":         "apiKey123",
			"deploymentName": "deploymentName",
			"apiVersion":     "v.1.0",
		},
		IsEnabled:            ToPtr(true),
		UseSystemCredentials: ToPtr(false),
	}
	provider, resp, err := client.AI.AddProvider(context.Background(), 1, req)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.Provider{
		ID:   2,
		Name: "OpenAI",
		Type: model.OpenAI,
		Credentials: map[string]string{
			"resourceName":   "resourceName",
			"apiKey":         "apiKey123",
			"deploymentName": "deploymentName",
			"apiVersion":     "v.1.0",
		},
		RawCredentials: map[string]any{
			"resourceName":   "resourceName",
			"apiKey":         "apiKey123",
			"deploymentName": "deploymentName",
			"apiVersion":     "v.1.0",
		},
		Config: model.ProviderConfig{
			ActionRules: []model.ActionRule{
				{
					Action:              model.ActionPreTranslate,
					AvailableAIModelIDs: []string{"gpt-3.5-turbo-instruct"},
				},
			},
		},
		IsEnabled:            true,
		UseSystemCredentials: false,
		CreatedAt:            "2023-09-20T11:11:05+00:00",
		UpdatedAt:            "2023-09-20T12:22:20+00:00",
	}
	assert.Equal(t, expected, provider)
}

func TestAIService_EditProvider(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/users/12345/ai/providers/2"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		testURL(t, r, path)
		testBody(t, r, `[{"op":"replace","path":"/name","value":"OpenAI"}]`+"\n")

		fmt.Fprint(w, `{
			"data": {
				"id": 2,
				"name": "OpenAI",
				"type": "open_ai"
			}
		}`)
	})

	req := []*model.UpdateRequest{
		{
			Op:    "replace",
			Path:  "/name",
			Value: "OpenAI",
		},
	}
	provider, resp, err := client.AI.EditProvider(context.Background(), 2, 12345, req)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.Provider{
		ID:   2,
		Name: "OpenAI",
		Type: model.OpenAI,
	}
	assert.Equal(t, expected, provider)
}

func TestAIService_DeleteProvider(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	t.Run("path with user id", func(t *testing.T) {
		const path = "/api/v2/users/12345/ai/providers/2"
		mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodDelete)
			testURL(t, r, path)
			w.WriteHeader(http.StatusNoContent)
		})

		resp, err := client.AI.DeleteProvider(context.Background(), 2, 12345)
		require.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, resp.StatusCode)
	})

	t.Run("path without user id (enterprise client)", func(t *testing.T) {
		const path = "/api/v2/ai/providers/2"
		mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodDelete)
			testURL(t, r, path)
			w.WriteHeader(http.StatusNoContent)
		})

		resp, err := client.AI.DeleteProvider(context.Background(), 2, 0)
		require.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, resp.StatusCode)
	})
}

func TestAIService_ListProviderModels(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/ai/providers/2/models"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		fmt.Fprint(w, `{
			"data": [
				{
					"data": {
						"id": "gpt-3.5-turbo-instruct"
					}
				}
			]
		}`)
	})

	models, resp, err := client.AI.ListProviderModels(context.Background(), 2, 0)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := []*model.ProviderModel{
		{
			ID: "gpt-3.5-turbo-instruct",
		},
	}
	assert.Equal(t, expected, models)
}

func TestAIService_ListProviderModels_invalidJSON(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/ai/providers/2/models", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `invalid json`)
	})

	models, _, err := client.AI.ListProviderModels(context.Background(), 2, 0)
	require.Error(t, err)
	assert.Nil(t, models)
}

func TestAIService_CreateProxyChatCompletion(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/users/1/ai/providers/2/chat/completions"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, path)
		testBody(t, r, `{"model":"gemini-1.5-pro","stream":false}`+"\n")

		fmt.Fprint(w, `{
			"data": {
				"modelId": "gpt-3.5-turbo-instruct"
			}
		}`)
	})

	req := &model.CreateProxyChatCompletionRequest{
		Model:  "gemini-1.5-pro",
		Stream: ToPtr(false),
	}
	completion, resp, err := client.AI.CreateProxyChatCompletion(context.Background(), 2, 1, req)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.ProxyChatCompletion{}
	assert.Equal(t, expected, completion)
}

func TestAIService_ClonePrompt(t *testing.T) {
	tests := []struct {
		name         string
		userID       int
		req          *model.PromptCloneRequest
		path         string
		expectedBody string
	}{
		{
			name:         "crowdin with name",
			userID:       1,
			req:          &model.PromptCloneRequest{Name: "Cloned prompt"},
			path:         "/api/v2/users/1/ai/prompts/2/clones",
			expectedBody: `{"name":"Cloned prompt"}` + "\n",
		},
		{
			name:         "enterprise with nil request",
			userID:       0,
			req:          nil,
			path:         "/api/v2/ai/prompts/2/clones",
			expectedBody: `{}` + "\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, mux, teardown := setupClient()
			defer teardown()

			mux.HandleFunc(tt.path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, http.MethodPost)
				testURL(t, r, tt.path)
				testBody(t, r, tt.expectedBody)

				fmt.Fprint(w, `{
					"data": {
						"id": 3,
						"name": "Cloned prompt",
						"action": "pre_translate",
						"aiProviderId": null,
						"aiModelId": null,
						"isEnabled": true,
						"enabledProjectIds": [],
						"config": {"mode": "basic"},
						"usageCount": 0,
						"createdAt": "2023-09-20T11:11:05+00:00",
						"updatedAt": "2023-09-20T12:22:20+00:00"
					}
				}`)
			})

			prompt, resp, err := client.AI.ClonePrompt(context.Background(), 2, tt.userID, tt.req)
			require.NoError(t, err)
			assert.NotNil(t, resp)

			expected := &model.Prompt{
				ID:                3,
				Name:              "Cloned prompt",
				Action:            "pre_translate",
				IsEnabled:         true,
				EnabledProjectIDs: []int{},
				Config:            model.PromptConfig{Mode: model.ModeBasic},
				CreatedAt:         "2023-09-20T11:11:05+00:00",
				UpdatedAt:         "2023-09-20T12:22:20+00:00",
			}
			assert.Equal(t, expected, prompt)
		})
	}
}

func TestAIService_GeneratePromptCompletion(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/users/1/ai/prompts/2/completions"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, path)
		testJSONBody(t, r, `{
			"resources": {
				"projectId": 1,
				"sourceLanguageId": "en",
				"targetLanguageId": "uk",
				"stringIds": [1, 2],
				"overridePromptValues": {
					"projectName": "Project",
					"projectDescription": "Project description",
					"organizationName": "Org"
				}
			},
			"tools": [
				{
					"tool": {
						"type": "function",
						"function": {
							"description": "Get the weather",
							"name": "get_weather",
							"parameters": {"type": "object"}
						}
					}
				}
			],
			"tool_choice": "auto"
		}`)

		fmt.Fprint(w, `{
			"data": {
				"identifier": "50fb3506-4127-4ba8-8296-f97dc7e3e0c3",
				"status": "created",
				"progress": 0,
				"attributes": {
					"aiPromptId": 2
				},
				"createdAt": "2023-09-23T11:26:54+00:00",
				"updatedAt": "2023-09-23T11:26:54+00:00",
				"startedAt": "2023-09-23T11:26:54+00:00",
				"finishedAt": "2023-09-23T11:26:54+00:00"
			}
		}`)
	})

	req := &model.PromptCompletionRequest{
		Resources: &model.PromptCompletionResources{
			ProjectID:        1,
			SourceLanguageID: "en",
			TargetLanguageID: "uk",
			StringIDs:        []int{1, 2},
			OverridePromptValues: &model.PromptCompletionOverridePromptValues{
				ProjectName:        "Project",
				ProjectDescription: "Project description",
				OrganizationName:   "Org",
			},
		},
		Tools: []*model.PromptCompletionTool{
			{
				Tool: &model.AITool{
					Type: "function",
					Function: &model.AIToolFunction{
						Description: "Get the weather",
						Name:        "get_weather",
						Parameters:  map[string]any{"type": "object"},
					},
				},
			},
		},
		ToolChoice: "auto",
	}
	completion, resp, err := client.AI.GeneratePromptCompletion(context.Background(), 2, 1, req)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.PromptCompletion{
		Identifier: "50fb3506-4127-4ba8-8296-f97dc7e3e0c3",
		Status:     "created",
		Progress:   0,
		Attributes: &model.PromptCompletionAttributes{AIPromptID: 2},
		CreatedAt:  "2023-09-23T11:26:54+00:00",
		UpdatedAt:  "2023-09-23T11:26:54+00:00",
		StartedAt:  "2023-09-23T11:26:54+00:00",
		FinishedAt: "2023-09-23T11:26:54+00:00",
	}
	assert.Equal(t, expected, completion)
}

func TestAIService_GeneratePromptCompletion_invalidRequest(t *testing.T) {
	client, _, teardown := setupClient()
	defer teardown()

	_, _, err := client.AI.GeneratePromptCompletion(context.Background(), 2, 1, &model.PromptCompletionRequest{})
	require.EqualError(t, err, "resources is required")
}

func TestAIService_GetPromptCompletionStatus(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/ai/prompts/2/completions/50fb3506-4127-4ba8-8296-f97dc7e3e0c3"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		fmt.Fprint(w, `{
			"data": {
				"identifier": "50fb3506-4127-4ba8-8296-f97dc7e3e0c3",
				"status": "finished",
				"progress": 100,
				"attributes": {
					"aiPromptId": 2
				},
				"createdAt": "2023-09-23T11:26:54+00:00",
				"updatedAt": "2023-09-23T11:26:54+00:00",
				"startedAt": "2023-09-23T11:26:54+00:00",
				"finishedAt": "2023-09-23T11:26:54+00:00"
			}
		}`)
	})

	completion, resp, err := client.AI.GetPromptCompletionStatus(context.Background(), 2, "50fb3506-4127-4ba8-8296-f97dc7e3e0c3", 0)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.PromptCompletion{
		Identifier: "50fb3506-4127-4ba8-8296-f97dc7e3e0c3",
		Status:     "finished",
		Progress:   100,
		Attributes: &model.PromptCompletionAttributes{AIPromptID: 2},
		CreatedAt:  "2023-09-23T11:26:54+00:00",
		UpdatedAt:  "2023-09-23T11:26:54+00:00",
		StartedAt:  "2023-09-23T11:26:54+00:00",
		FinishedAt: "2023-09-23T11:26:54+00:00",
	}
	assert.Equal(t, expected, completion)
}

func TestAIService_CancelPromptCompletion(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/users/1/ai/prompts/2/completions/50fb3506-4127-4ba8-8296-f97dc7e3e0c3"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		testURL(t, r, path)

		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := client.AI.CancelPromptCompletion(context.Background(), 2, "50fb3506-4127-4ba8-8296-f97dc7e3e0c3", 1)
	require.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

func TestAIService_DownloadPromptCompletion(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/ai/prompts/2/completions/50fb3506-4127-4ba8-8296-f97dc7e3e0c3/download"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		fmt.Fprint(w, `{
			"data": {
				"url": "https://test.com",
				"expireIn": "2023-09-20T10:31:21+00:00"
			}
		}`)
	})

	link, resp, err := client.AI.DownloadPromptCompletion(context.Background(), 2, "50fb3506-4127-4ba8-8296-f97dc7e3e0c3", 0)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.DownloadLink{URL: "https://test.com", ExpireIn: "2023-09-20T10:31:21+00:00"}
	assert.Equal(t, expected, link)
}

func TestAIService_TranslateFile(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/users/1/ai/file-translations"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, path)
		testJSONBody(t, r, `{
			"storageId": 123,
			"sourceLanguageId": "en",
			"targetLanguageId": "uk",
			"type": "json",
			"parserVersion": 1,
			"tmIds": [1],
			"glossaryIds": [2],
			"styleGuideIds": [3],
			"aiPromptId": 4,
			"instructions": ["Keep a formal tone"],
			"attachmentIds": [5]
		}`)

		fmt.Fprint(w, `{
			"data": {
				"identifier": "50fb3506-4127-4ba8-8296-f97dc7e3e0c3",
				"status": "created",
				"progress": 0,
				"attributes": {
					"stage": "start",
					"error": null,
					"downloadName": null,
					"sourceLanguageId": "en",
					"targetLanguageId": "uk",
					"originalFileName": "Sample_Chrome.json",
					"detectedType": null,
					"parserVersion": null
				},
				"createdAt": "2026-01-23T11:26:54+00:00",
				"updatedAt": null,
				"startedAt": null,
				"finishedAt": null
			}
		}`)
	})

	req := &model.AIFileTranslationRequest{
		StorageID:        123,
		SourceLanguageID: "en",
		TargetLanguageID: "uk",
		Type:             "json",
		ParserVersion:    1,
		TMIDs:            []int{1},
		GlossaryIDs:      []int{2},
		StyleGuideIDs:    []int{3},
		AIPromptID:       4,
		Instructions:     []string{"Keep a formal tone"},
		AttachmentIDs:    []int{5},
	}
	job, resp, err := client.AI.TranslateFile(context.Background(), 1, req)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.AIFileTranslation{
		Identifier: "50fb3506-4127-4ba8-8296-f97dc7e3e0c3",
		Status:     "created",
		Progress:   0,
		Attributes: &model.AIFileTranslationAttributes{
			Stage:            "start",
			SourceLanguageID: ToPtr("en"),
			TargetLanguageID: "uk",
			OriginalFileName: "Sample_Chrome.json",
		},
		CreatedAt: "2026-01-23T11:26:54+00:00",
	}
	assert.Equal(t, expected, job)
}

func TestAIService_TranslateFile_invalidRequest(t *testing.T) {
	client, _, teardown := setupClient()
	defer teardown()

	_, _, err := client.AI.TranslateFile(context.Background(), 1, &model.AIFileTranslationRequest{})
	require.EqualError(t, err, "storageId is required")
}

func TestAIService_GetFileTranslationStatus(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/ai/file-translations/50fb3506-4127-4ba8-8296-f97dc7e3e0c3"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		fmt.Fprint(w, `{
			"data": {
				"identifier": "50fb3506-4127-4ba8-8296-f97dc7e3e0c3",
				"status": "failed",
				"progress": 100,
				"attributes": {
					"stage": "done",
					"error": {
						"stage": "import",
						"message": "Failed to parse file"
					},
					"downloadName": "file.json",
					"sourceLanguageId": "en",
					"targetLanguageId": "uk",
					"originalFileName": "Sample_Chrome.json",
					"detectedType": "chrome",
					"parserVersion": 2
				},
				"createdAt": "2026-01-23T11:26:54+00:00",
				"updatedAt": "2026-01-23T11:26:55+00:00",
				"startedAt": "2026-01-23T11:26:56+00:00",
				"finishedAt": "2026-01-23T11:26:57+00:00"
			}
		}`)
	})

	job, resp, err := client.AI.GetFileTranslationStatus(context.Background(), "50fb3506-4127-4ba8-8296-f97dc7e3e0c3", 0)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.AIFileTranslation{
		Identifier: "50fb3506-4127-4ba8-8296-f97dc7e3e0c3",
		Status:     "failed",
		Progress:   100,
		Attributes: &model.AIFileTranslationAttributes{
			Stage:            "done",
			Error:            &model.AIFileTranslationError{Stage: "import", Message: "Failed to parse file"},
			DownloadName:     ToPtr("file.json"),
			SourceLanguageID: ToPtr("en"),
			TargetLanguageID: "uk",
			OriginalFileName: "Sample_Chrome.json",
			DetectedType:     ToPtr("chrome"),
			ParserVersion:    ToPtr(2),
		},
		CreatedAt:  "2026-01-23T11:26:54+00:00",
		UpdatedAt:  ToPtr("2026-01-23T11:26:55+00:00"),
		StartedAt:  ToPtr("2026-01-23T11:26:56+00:00"),
		FinishedAt: ToPtr("2026-01-23T11:26:57+00:00"),
	}
	assert.Equal(t, expected, job)
}

func TestAIService_CancelFileTranslation(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/users/1/ai/file-translations/50fb3506-4127-4ba8-8296-f97dc7e3e0c3"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		testURL(t, r, path)

		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := client.AI.CancelFileTranslation(context.Background(), "50fb3506-4127-4ba8-8296-f97dc7e3e0c3", 1)
	require.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

func TestAIService_DownloadTranslatedFile(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/users/1/ai/file-translations/50fb3506-4127-4ba8-8296-f97dc7e3e0c3/download"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		fmt.Fprint(w, `{
			"data": {
				"url": "https://test.com/file",
				"expireIn": "2023-09-20T10:31:21+00:00"
			}
		}`)
	})

	link, resp, err := client.AI.DownloadTranslatedFile(context.Background(), "50fb3506-4127-4ba8-8296-f97dc7e3e0c3", 1)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.DownloadLink{URL: "https://test.com/file", ExpireIn: "2023-09-20T10:31:21+00:00"}
	assert.Equal(t, expected, link)
}

func TestAIService_DownloadTranslatedFileStrings(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/ai/file-translations/50fb3506-4127-4ba8-8296-f97dc7e3e0c3/translations"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		fmt.Fprint(w, `{
			"data": {
				"url": "https://test.com/strings",
				"expireIn": "2023-09-20T10:31:21+00:00"
			}
		}`)
	})

	link, resp, err := client.AI.DownloadTranslatedFileStrings(context.Background(), "50fb3506-4127-4ba8-8296-f97dc7e3e0c3", 0)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.DownloadLink{URL: "https://test.com/strings", ExpireIn: "2023-09-20T10:31:21+00:00"}
	assert.Equal(t, expected, link)
}

func TestAIService_TranslateStrings(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/users/1/ai/translate"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, path)
		testJSONBody(t, r, `{
			"strings": ["Text 1", "Text 2"],
			"sourceLanguageId": "en",
			"targetLanguageId": "uk",
			"tmIds": [1],
			"glossaryIds": [2],
			"styleGuideIds": [3],
			"instructions": ["Keep a formal tone"],
			"attachmentIds": [4],
			"aiProviderId": 5,
			"aiModelId": "gpt-4.1"
		}`)

		fmt.Fprint(w, `{
			"data": {
				"sourceLanguageId": "en",
				"targetLanguageId": "uk",
				"translations": ["Текст 1", "Текст 2"]
			}
		}`)
	})

	req := &model.AITranslateStringsRequest{
		Strings:          []string{"Text 1", "Text 2"},
		SourceLanguageID: "en",
		TargetLanguageID: "uk",
		TMIDs:            []int{1},
		GlossaryIDs:      []int{2},
		StyleGuideIDs:    []int{3},
		Instructions:     []string{"Keep a formal tone"},
		AttachmentIDs:    []int{4},
		AIProviderID:     5,
		AIModelID:        "gpt-4.1",
	}
	result, resp, err := client.AI.TranslateStrings(context.Background(), 1, req)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.AITranslateStrings{
		SourceLanguageID: "en",
		TargetLanguageID: "uk",
		Translations:     []string{"Текст 1", "Текст 2"},
	}
	assert.Equal(t, expected, result)
}

func TestAIService_TranslateStrings_enterprise(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/ai/translate"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, path)
		testBody(t, r, `{"strings":["Text"],"targetLanguageId":"uk","aiPromptId":7}`+"\n")

		fmt.Fprint(w, `{"data": {"sourceLanguageId": "en", "targetLanguageId": "uk", "translations": ["Текст"]}}`)
	})

	req := &model.AITranslateStringsRequest{
		Strings:          []string{"Text"},
		TargetLanguageID: "uk",
		AIPromptID:       7,
	}
	result, _, err := client.AI.TranslateStrings(context.Background(), 0, req)
	require.NoError(t, err)
	assert.Equal(t, []string{"Текст"}, result.Translations)
}

func TestAIService_TranslateStrings_invalidRequest(t *testing.T) {
	client, _, teardown := setupClient()
	defer teardown()

	req := &model.AITranslateStringsRequest{Strings: []string{"Text"}, TargetLanguageID: "uk"}
	_, _, err := client.AI.TranslateStrings(context.Background(), 1, req)
	require.EqualError(t, err, "aiPromptId or aiProviderId with aiModelId is required")
}

func TestAIService_ListAllProviderModels(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/users/1/ai/providers/models"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		fmt.Fprint(w, `{
			"data": [
				{
					"data": {
						"id": "gpt-4.1",
						"provider": "open_ai",
						"providerName": "OpenAI",
						"providerId": 2,
						"contextWindow": 1047576,
						"maxOutputTokens": 32768,
						"supportsStreaming": true,
						"supportsFunctionCalling": true,
						"supportsJsonMode": true,
						"supportsJsonSchema": true,
						"supportsVision": true,
						"isCompatibleWithAiLimit": true
					}
				}
			]
		}`)
	})

	models, resp, err := client.AI.ListAllProviderModels(context.Background(), 1)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := []*model.ProviderModel{
		{
			ID:                      "gpt-4.1",
			Provider:                ToPtr("open_ai"),
			ProviderName:            ToPtr("OpenAI"),
			ProviderID:              ToPtr(2),
			ContextWindow:           ToPtr(1047576),
			MaxOutputTokens:         ToPtr(32768),
			SupportsStreaming:       ToPtr(true),
			SupportsFunctionCalling: ToPtr(true),
			SupportsJSONMode:        ToPtr(true),
			SupportsJSONSchema:      ToPtr(true),
			SupportsVision:          ToPtr(true),
			IsCompatibleWithAILimit: true,
		},
	}
	assert.Equal(t, expected, models)
}

func TestAIService_ListAllProviderModels_invalidJSON(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/ai/providers/models", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `invalid json`)
	})

	models, _, err := client.AI.ListAllProviderModels(context.Background(), 0)
	require.Error(t, err)
	assert.Nil(t, models)
}

func TestAIService_ListSupportedProviderModels(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/ai/providers/supported-models"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path+"?enabled=true&limit=10&orderBy=id+desc&providerType=open_ai")

		fmt.Fprint(w, `{
			"data": [
				{
					"data": {
						"providerId": 2,
						"providerType": "open_ai",
						"providerName": "OpenAI",
						"id": "gpt-5.2",
						"displayName": "GPT-5.2",
						"supportReasoning": true,
						"intelligence": 5,
						"speed": 3,
						"price": {
							"input": 1.25,
							"output": 10
						},
						"modalities": {
							"input": {"text": true, "image": true, "audio": false},
							"output": {"text": true, "image": false, "audio": false}
						},
						"contextWindow": 400000,
						"maxOutputTokens": 128000,
						"knowledgeCutoff": "2024-09-30T00:00:00+00:00",
						"releaseDate": null,
						"features": {
							"streaming": true,
							"structuredOutput": true,
							"functionCalling": true
						}
					}
				}
			],
			"pagination": {
				"offset": 0,
				"limit": 10
			}
		}`)
	})

	opts := &model.SupportedModelsListOptions{
		ProviderType: model.OpenAI,
		Enabled:      ToPtr(true),
		OrderBy:      "id desc",
		ListOptions:  model.ListOptions{Limit: 10},
	}
	models, resp, err := client.AI.ListSupportedProviderModels(context.Background(), 0, opts)
	require.NoError(t, err)
	assert.Equal(t, 10, resp.Pagination.Limit)

	expected := []*model.SupportedModel{
		{
			ProviderID:       ToPtr(2),
			ProviderType:     "open_ai",
			ProviderName:     "OpenAI",
			ID:               "gpt-5.2",
			DisplayName:      "GPT-5.2",
			SupportReasoning: true,
			Intelligence:     5,
			Speed:            3,
			Price:            &model.SupportedModelPrice{Input: 1.25, Output: 10},
			Modalities: &model.SupportedModelModalities{
				Input:  &model.SupportedModelModality{Text: true, Image: true},
				Output: &model.SupportedModelModality{Text: true},
			},
			ContextWindow:   400000,
			MaxOutputTokens: 128000,
			KnowledgeCutoff: ToPtr("2024-09-30T00:00:00+00:00"),
			Features: &model.SupportedModelFeatures{
				Streaming:        true,
				StructuredOutput: true,
				FunctionCalling:  true,
			},
		},
	}
	assert.Equal(t, expected, models)
}

func TestAIService_ListSupportedProviderModels_invalidJSON(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/users/1/ai/providers/supported-models", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `invalid json`)
	})

	models, _, err := client.AI.ListSupportedProviderModels(context.Background(), 1, nil)
	require.Error(t, err)
	assert.Nil(t, models)
}

func TestAIService_Gateway(t *testing.T) {
	tests := []struct {
		name         string
		method       string
		path         string
		expectedBody string
		call         func(c *Client) (map[string]any, *Response, error)
	}{
		{
			name:   "GET crowdin",
			method: http.MethodGet,
			path:   "/api/v2/users/1/ai/providers/2/gateway/models",
			call: func(c *Client) (map[string]any, *Response, error) {
				return c.AI.GatewayGet(context.Background(), 2, 1, "models")
			},
		},
		{
			name:         "POST enterprise",
			method:       http.MethodPost,
			path:         "/api/v2/ai/providers/2/gateway/chat/completions",
			expectedBody: `{"model":"gpt-4.1"}` + "\n",
			call: func(c *Client) (map[string]any, *Response, error) {
				return c.AI.GatewayPost(context.Background(), 2, 0, "/chat/completions", map[string]any{"model": "gpt-4.1"})
			},
		},
		{
			name:         "PUT crowdin",
			method:       http.MethodPut,
			path:         "/api/v2/users/1/ai/providers/2/gateway/files/1",
			expectedBody: `{"name":"file"}` + "\n",
			call: func(c *Client) (map[string]any, *Response, error) {
				return c.AI.GatewayPut(context.Background(), 2, 1, "files/1", map[string]any{"name": "file"})
			},
		},
		{
			name:         "PATCH enterprise",
			method:       http.MethodPatch,
			path:         "/api/v2/ai/providers/2/gateway/files/1",
			expectedBody: `{"name":"file"}` + "\n",
			call: func(c *Client) (map[string]any, *Response, error) {
				return c.AI.GatewayPatch(context.Background(), 2, 0, "files/1", map[string]any{"name": "file"})
			},
		},
		{
			name:   "DELETE crowdin",
			method: http.MethodDelete,
			path:   "/api/v2/users/1/ai/providers/2/gateway/files/1",
			call: func(c *Client) (map[string]any, *Response, error) {
				return c.AI.GatewayDelete(context.Background(), 2, 1, "files/1")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, mux, teardown := setupClient()
			defer teardown()

			mux.HandleFunc(tt.path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, tt.method)
				testURL(t, r, tt.path)
				testBody(t, r, tt.expectedBody)

				fmt.Fprint(w, `{"id": "resp-1", "object": "response"}`)
			})

			res, resp, err := tt.call(client)
			require.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, map[string]any{"id": "resp-1", "object": "response"}, res)
		})
	}
}

func TestAIService_GenerateReport(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/users/1/ai/reports"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, path)
		testJSONBody(t, r, `{
			"type": "tokens-usage-raw-data",
			"schema": {
				"dateFrom": "2024-01-23T07:00:14+00:00",
				"dateTo": "2024-09-27T07:00:14+00:00",
				"format": "csv",
				"projectIds": [1],
				"promptIds": [2],
				"userIds": [3]
			}
		}`)

		fmt.Fprint(w, `{
			"data": {
				"identifier": "50fb3506-4127-4ba8-8296-f97dc7e3e0c3",
				"status": "created",
				"progress": 0,
				"attributes": {
					"format": "csv",
					"reportType": "tokens-usage-raw-data",
					"schema": {"dateFrom": "2024-01-23T07:00:14+00:00"}
				},
				"createdAt": "2023-09-23T11:26:54+00:00",
				"updatedAt": "2023-09-23T11:26:54+00:00",
				"startedAt": "2023-09-23T11:26:54+00:00",
				"finishedAt": "2023-09-23T11:26:54+00:00",
				"eta": "1 second"
			}
		}`)
	})

	req := &model.AIReportGenerateRequest{
		Type: model.AIReportUsageRawData,
		Schema: &model.AIReportSchema{
			DateFrom:   "2024-01-23T07:00:14+00:00",
			DateTo:     "2024-09-27T07:00:14+00:00",
			Format:     "csv",
			ProjectIDs: []int{1},
			PromptIDs:  []int{2},
			UserIDs:    []int{3},
		},
	}
	report, resp, err := client.AI.GenerateReport(context.Background(), 1, req)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.AIReport{
		Identifier: "50fb3506-4127-4ba8-8296-f97dc7e3e0c3",
		Status:     "created",
		Progress:   0,
		Attributes: &model.AIReportAttributes{
			Format:     "csv",
			ReportType: "tokens-usage-raw-data",
			Schema:     map[string]any{"dateFrom": "2024-01-23T07:00:14+00:00"},
		},
		CreatedAt:  "2023-09-23T11:26:54+00:00",
		UpdatedAt:  "2023-09-23T11:26:54+00:00",
		StartedAt:  "2023-09-23T11:26:54+00:00",
		FinishedAt: "2023-09-23T11:26:54+00:00",
		ETA:        "1 second",
	}
	assert.Equal(t, expected, report)
}

func TestAIService_GenerateReport_invalidRequest(t *testing.T) {
	client, _, teardown := setupClient()
	defer teardown()

	_, _, err := client.AI.GenerateReport(context.Background(), 1, &model.AIReportGenerateRequest{Type: model.AIReportCostsByUsers})
	require.EqualError(t, err, "schema is required")
}

func TestAIService_CheckReportStatus(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/ai/reports/50fb3506-4127-4ba8-8296-f97dc7e3e0c3"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		fmt.Fprint(w, `{
			"data": {
				"identifier": "50fb3506-4127-4ba8-8296-f97dc7e3e0c3",
				"status": "finished",
				"progress": 100,
				"attributes": {
					"format": "json",
					"reportType": "costs-by-users",
					"schema": {}
				},
				"createdAt": "2023-09-23T11:26:54+00:00",
				"updatedAt": "2023-09-23T11:26:54+00:00",
				"startedAt": "2023-09-23T11:26:54+00:00",
				"finishedAt": "2023-09-23T11:26:54+00:00",
				"eta": "0 seconds"
			}
		}`)
	})

	report, resp, err := client.AI.CheckReportStatus(context.Background(), "50fb3506-4127-4ba8-8296-f97dc7e3e0c3", 0)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	assert.Equal(t, "finished", report.Status)
	assert.Equal(t, 100, report.Progress)
	assert.Equal(t, "costs-by-users", report.Attributes.ReportType)
	assert.Equal(t, "json", report.Attributes.Format)
}

func TestAIService_DownloadReport(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/users/1/ai/reports/50fb3506-4127-4ba8-8296-f97dc7e3e0c3/download"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		fmt.Fprint(w, `{
			"data": {
				"url": "https://test.com/report",
				"expireIn": "2023-09-20T10:31:21+00:00"
			}
		}`)
	})

	link, resp, err := client.AI.DownloadReport(context.Background(), "50fb3506-4127-4ba8-8296-f97dc7e3e0c3", 1)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.DownloadLink{URL: "https://test.com/report", ExpireIn: "2023-09-20T10:31:21+00:00"}
	assert.Equal(t, expected, link)
}

func TestAIService_ListRequestLogs(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/users/1/ai/request-logs"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path+"?limit=25&projectId=2&sourceAction=ai_gateway&statuses=success%2Cerror&systemCredentials=false")

		fmt.Fprint(w, `{
			"data": [
				{
					"data": {
						"id": 1,
						"requestId": "req-1",
						"createdAt": "2024-01-23T07:00:14+00:00",
						"status": "success",
						"httpStatus": 200,
						"model": "gpt-4.1",
						"sourceAction": "ai_gateway",
						"promptAction": null,
						"systemCredentials": false,
						"isAutoTriggered": true,
						"durationMs": 1500,
						"inputTokens": 100,
						"outputTokens": 50,
						"totalCost": 0.0025,
						"userId": 3,
						"projectId": 2,
						"promptId": null,
						"aiProviderId": 4,
						"tokenName": "token",
						"oauthClientId": null,
						"oauthClientName": null,
						"ip": "127.0.0.1",
						"userAgent": "agent",
						"error": null
					}
				}
			],
			"pagination": {
				"offset": 0,
				"limit": 25
			}
		}`)
	})

	opts := &model.AIRequestLogsListOptions{
		ProjectID:         2,
		SourceAction:      "ai_gateway",
		Statuses:          []string{"success", "error"},
		SystemCredentials: ToPtr(false),
		ListOptions:       model.ListOptions{Limit: 25},
	}
	logs, resp, err := client.AI.ListRequestLogs(context.Background(), 1, opts)
	require.NoError(t, err)
	assert.Equal(t, 25, resp.Pagination.Limit)

	expected := []*model.AIRequestLog{
		{
			ID:                1,
			RequestID:         "req-1",
			CreatedAt:         "2024-01-23T07:00:14+00:00",
			Status:            "success",
			HTTPStatus:        ToPtr(200),
			Model:             "gpt-4.1",
			SourceAction:      "ai_gateway",
			SystemCredentials: false,
			IsAutoTriggered:   true,
			DurationMs:        ToPtr(1500),
			InputTokens:       ToPtr(100),
			OutputTokens:      ToPtr(50),
			TotalCost:         ToPtr(0.0025),
			UserID:            ToPtr(3),
			ProjectID:         ToPtr(2),
			AIProviderID:      4,
			TokenName:         ToPtr("token"),
			IP:                ToPtr("127.0.0.1"),
			UserAgent:         ToPtr("agent"),
		},
	}
	assert.Equal(t, expected, logs)
}

func TestAIService_ListRequestLogs_invalidJSON(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/ai/request-logs", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `invalid json`)
	})

	logs, _, err := client.AI.ListRequestLogs(context.Background(), 0, nil)
	require.Error(t, err)
	assert.Nil(t, logs)
}

func TestAIService_ExportRequestLogs(t *testing.T) {
	tests := []struct {
		name         string
		userID       int
		req          *model.AIRequestLogsExportRequest
		path         string
		expectedBody string
	}{
		{
			name:   "crowdin with filters",
			userID: 1,
			req: &model.AIRequestLogsExportRequest{
				Format:          "csv",
				ProjectID:       2,
				Statuses:        []string{"error"},
				IsAutoTriggered: ToPtr(false),
				CreatedAfter:    "2024-01-23T07:00:14+00:00",
			},
			path:         "/api/v2/users/1/ai/request-logs/exports",
			expectedBody: `{"format":"csv","projectId":2,"statuses":["error"],"isAutoTriggered":false,"createdAfter":"2024-01-23T07:00:14+00:00"}` + "\n",
		},
		{
			name:         "enterprise with nil request",
			req:          nil,
			path:         "/api/v2/ai/request-logs/exports",
			expectedBody: `{}` + "\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, mux, teardown := setupClient()
			defer teardown()

			mux.HandleFunc(tt.path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, http.MethodPost)
				testURL(t, r, tt.path)
				testBody(t, r, tt.expectedBody)

				fmt.Fprint(w, `{
					"data": {
						"identifier": "50fb3506-4127-4ba8-8296-f97dc7e3e0c3",
						"status": "created",
						"progress": 0,
						"attributes": {
							"format": "csv",
							"filters": {"projectId": 2}
						},
						"createdAt": "2024-01-23T07:00:14+00:00",
						"updatedAt": "2024-01-23T07:00:14+00:00",
						"startedAt": null,
						"finishedAt": null,
						"eta": null
					}
				}`)
			})

			export, resp, err := client.AI.ExportRequestLogs(context.Background(), tt.userID, tt.req)
			require.NoError(t, err)
			assert.NotNil(t, resp)

			expected := &model.AIRequestLogsExport{
				Identifier: "50fb3506-4127-4ba8-8296-f97dc7e3e0c3",
				Status:     "created",
				Progress:   0,
				Attributes: &model.AIRequestLogsExportAttributes{
					Format:  "csv",
					Filters: map[string]any{"projectId": float64(2)},
				},
				CreatedAt: "2024-01-23T07:00:14+00:00",
				UpdatedAt: "2024-01-23T07:00:14+00:00",
			}
			assert.Equal(t, expected, export)
		})
	}
}

func TestAIService_CheckRequestLogsExportStatus(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/users/1/ai/request-logs/exports/50fb3506-4127-4ba8-8296-f97dc7e3e0c3"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		fmt.Fprint(w, `{
			"data": {
				"identifier": "50fb3506-4127-4ba8-8296-f97dc7e3e0c3",
				"status": "finished",
				"progress": 100,
				"attributes": {
					"format": "csv",
					"filters": {}
				},
				"createdAt": "2024-01-23T07:00:14+00:00",
				"updatedAt": "2024-01-23T07:00:15+00:00",
				"startedAt": "2024-01-23T07:00:16+00:00",
				"finishedAt": "2024-01-23T07:00:17+00:00",
				"eta": "3 seconds"
			}
		}`)
	})

	export, resp, err := client.AI.CheckRequestLogsExportStatus(context.Background(), "50fb3506-4127-4ba8-8296-f97dc7e3e0c3", 1)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.AIRequestLogsExport{
		Identifier: "50fb3506-4127-4ba8-8296-f97dc7e3e0c3",
		Status:     "finished",
		Progress:   100,
		Attributes: &model.AIRequestLogsExportAttributes{
			Format:  "csv",
			Filters: map[string]any{},
		},
		CreatedAt:  "2024-01-23T07:00:14+00:00",
		UpdatedAt:  "2024-01-23T07:00:15+00:00",
		StartedAt:  ToPtr("2024-01-23T07:00:16+00:00"),
		FinishedAt: ToPtr("2024-01-23T07:00:17+00:00"),
		ETA:        ToPtr("3 seconds"),
	}
	assert.Equal(t, expected, export)
}

func TestAIService_DownloadRequestLogsExport(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/ai/request-logs/exports/50fb3506-4127-4ba8-8296-f97dc7e3e0c3/download"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		fmt.Fprint(w, `{
			"data": {
				"url": "https://test.com/logs.csv",
				"expireIn": "2023-09-20T10:31:21+00:00"
			}
		}`)
	})

	link, resp, err := client.AI.DownloadRequestLogsExport(context.Background(), "50fb3506-4127-4ba8-8296-f97dc7e3e0c3", 0)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.DownloadLink{URL: "https://test.com/logs.csv", ExpireIn: "2023-09-20T10:31:21+00:00"}
	assert.Equal(t, expected, link)
}

// aiSettingsJSON is the AI settings fixture used in the settings tests.
const aiSettingsJSON = `{
	"data": {
		"preTranslationAiPromptId": 1,
		"editorSuggestionAiPromptId": 2,
		"qaCheckActionAiPromptId": 3,
		"contextReviewAiPromptId": null,
		"alignmentActionAiPromptId": 4,
		"dailyCostLimit": 10.5,
		"monthlyCostLimit": null,
		"userDailyCostLimit": 1,
		"userMonthlyCostLimit": null,
		"isLimitingActive": true,
		"perUserOverrides": [
			{
				"userId": 5,
				"costLimitMode": "custom",
				"dailyCostLimit": 2,
				"monthlyCostLimit": 20,
				"createdAt": "2024-01-23T07:00:14+00:00",
				"updatedAt": "2024-01-23T07:00:14+00:00"
			}
		]
	}
}`

func expectedAISettings() *model.AISettings {
	return &model.AISettings{
		PreTranslationAIPromptID:   1,
		EditorSuggestionAIPromptID: 2,
		QACheckActionAIPromptID:    ToPtr(3),
		AlignmentActionAIPromptID:  ToPtr(4),
		DailyCostLimit:             ToPtr(10.5),
		UserDailyCostLimit:         ToPtr(1.0),
		IsLimitingActive:           true,
		PerUserOverrides: []*model.AIUserCostLimitOverride{
			{
				UserID:           5,
				CostLimitMode:    "custom",
				DailyCostLimit:   ToPtr(2.0),
				MonthlyCostLimit: ToPtr(20.0),
				CreatedAt:        "2024-01-23T07:00:14+00:00",
				UpdatedAt:        "2024-01-23T07:00:14+00:00",
			},
		},
	}
}

func TestAIService_GetSettings(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/ai/settings"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		fmt.Fprint(w, aiSettingsJSON)
	})

	settings, resp, err := client.AI.GetSettings(context.Background(), 0)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	assert.Equal(t, expectedAISettings(), settings)
}

func TestAIService_EditSettings(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/users/1/ai/settings"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		testURL(t, r, path)
		testBody(t, r, `[{"op":"replace","path":"/dailyCostLimit","value":10.5},{"op":"remove","path":"/monthlyCostLimit"}]`+"\n")

		fmt.Fprint(w, aiSettingsJSON)
	})

	req := []*model.UpdateRequest{
		{Op: model.OpReplace, Path: "/dailyCostLimit", Value: 10.5},
		{Op: model.OpRemove, Path: "/monthlyCostLimit"},
	}
	settings, resp, err := client.AI.EditSettings(context.Background(), 1, req)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	assert.Equal(t, expectedAISettings(), settings)
}

func TestAIService_GetProjectAISettings(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/projects/1/ai/settings"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		fmt.Fprint(w, `{
			"data": {
				"preTranslationAiPromptId": 1,
				"editorSuggestionAiPromptId": 2,
				"alignmentActionAiPromptId": null,
				"qaCheckActionAiPromptId": 3,
				"contextReviewAiPromptId": 4
			}
		}`)
	})

	settings, resp, err := client.AI.GetProjectAISettings(context.Background(), 1)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.ProjectAISettings{
		EditorSuggestionAIPromptID: 2,
		QACheckActionAIPromptID:    ToPtr(3),
		ContextReviewAIPromptID:    ToPtr(4),
	}
	assert.Equal(t, expected, settings)
}

func TestAIService_ListSnippets(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/users/1/ai/settings/snippets"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path+"?limit=10&offset=5")

		fmt.Fprint(w, `{
			"data": [
				{
					"data": {
						"id": 1,
						"description": "Company description",
						"placeholder": "\u0025custom:companyDescription\u0025",
						"value": "Crowdin",
						"createdAt": "2024-01-23T07:00:14+00:00",
						"updatedAt": "2024-01-23T07:00:14+00:00"
					}
				}
			],
			"pagination": {
				"offset": 5,
				"limit": 10
			}
		}`)
	})

	snippets, resp, err := client.AI.ListSnippets(context.Background(), 1, &model.ListOptions{Limit: 10, Offset: 5})
	require.NoError(t, err)
	assert.Equal(t, 5, resp.Pagination.Offset)

	expected := []*model.AISnippet{
		{
			ID:          1,
			Description: "Company description",
			Placeholder: "%custom:companyDescription%",
			Value:       "Crowdin",
			CreatedAt:   "2024-01-23T07:00:14+00:00",
			UpdatedAt:   "2024-01-23T07:00:14+00:00",
		},
	}
	assert.Equal(t, expected, snippets)
}

func TestAIService_ListSnippets_invalidJSON(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/ai/settings/snippets", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `invalid json`)
	})

	snippets, _, err := client.AI.ListSnippets(context.Background(), 0, nil)
	require.Error(t, err)
	assert.Nil(t, snippets)
}

func TestAIService_GetSnippet(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/ai/settings/snippets/1"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		fmt.Fprint(w, `{
			"data": {
				"id": 1,
				"description": "Company description",
				"placeholder": "\u0025custom:companyDescription\u0025",
				"value": "Crowdin",
				"createdAt": "2024-01-23T07:00:14+00:00",
				"updatedAt": "2024-01-23T07:00:14+00:00"
			}
		}`)
	})

	snippet, resp, err := client.AI.GetSnippet(context.Background(), 1, 0)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.AISnippet{
		ID:          1,
		Description: "Company description",
		Placeholder: "%custom:companyDescription%",
		Value:       "Crowdin",
		CreatedAt:   "2024-01-23T07:00:14+00:00",
		UpdatedAt:   "2024-01-23T07:00:14+00:00",
	}
	assert.Equal(t, expected, snippet)
}

func TestAIService_AddSnippet(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/users/1/ai/settings/snippets"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, path)
		testJSONBody(t, r, `{
			"description": "Company description",
			"placeholder": "\u0025custom:companyDescription\u0025",
			"value": "Crowdin"
		}`)

		fmt.Fprint(w, `{
			"data": {
				"id": 1,
				"description": "Company description",
				"placeholder": "\u0025custom:companyDescription\u0025",
				"value": "Crowdin",
				"createdAt": "2024-01-23T07:00:14+00:00",
				"updatedAt": "2024-01-23T07:00:14+00:00"
			}
		}`)
	})

	req := &model.AISnippetAddRequest{
		Description: "Company description",
		Placeholder: "%custom:companyDescription%",
		Value:       "Crowdin",
	}
	snippet, resp, err := client.AI.AddSnippet(context.Background(), 1, req)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	assert.Equal(t, 1, snippet.ID)
	assert.Equal(t, "%custom:companyDescription%", snippet.Placeholder)
}

func TestAIService_AddSnippet_invalidRequest(t *testing.T) {
	client, _, teardown := setupClient()
	defer teardown()

	_, _, err := client.AI.AddSnippet(context.Background(), 1, &model.AISnippetAddRequest{})
	require.EqualError(t, err, "description is required")
}

func TestAIService_EditSnippet(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/ai/settings/snippets/1"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		testURL(t, r, path)
		testBody(t, r, `[{"op":"replace","path":"/value","value":"Crowdin Enterprise"}]`+"\n")

		fmt.Fprint(w, `{
			"data": {
				"id": 1,
				"description": "Company description",
				"placeholder": "\u0025custom:companyDescription\u0025",
				"value": "Crowdin Enterprise",
				"createdAt": "2024-01-23T07:00:14+00:00",
				"updatedAt": "2024-01-23T07:00:15+00:00"
			}
		}`)
	})

	req := []*model.UpdateRequest{{Op: model.OpReplace, Path: "/value", Value: "Crowdin Enterprise"}}
	snippet, resp, err := client.AI.EditSnippet(context.Background(), 1, 0, req)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	assert.Equal(t, "Crowdin Enterprise", snippet.Value)
}

func TestAIService_DeleteSnippet(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/users/1/ai/settings/snippets/1"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		testURL(t, r, path)

		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := client.AI.DeleteSnippet(context.Background(), 1, 1)
	require.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

func TestAIService_ListUsageMembers(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/ai/usage/members"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path+"?orderBy=id&userIds=1%2C2")

		fmt.Fprint(w, `{
			"data": [
				{
					"data": {
						"user": {
							"id": 1,
							"username": "john",
							"fullName": "John Smith",
							"avatarUrl": "https://test.com/avatar"
						},
						"dailyCostLimit": null,
						"dailyCostSpent": 0.5,
						"dailyResetAt": "2024-01-24T00:00:00+00:00",
						"monthlyCostLimit": 100,
						"monthlyCostSpent": 12.25,
						"monthlyResetAt": "2024-02-01T00:00:00+00:00"
					}
				}
			],
			"pagination": {
				"offset": 0,
				"limit": 25
			}
		}`)
	})

	opts := &model.AIUsageMembersListOptions{UserIDs: []int{1, 2}, OrderBy: "id"}
	members, resp, err := client.AI.ListUsageMembers(context.Background(), 0, opts)
	require.NoError(t, err)
	assert.Equal(t, 25, resp.Pagination.Limit)

	expected := []*model.AIUsageMember{
		{
			User: &model.ShortUser{
				ID:        1,
				Username:  "john",
				FullName:  "John Smith",
				AvatarURL: "https://test.com/avatar",
			},
			DailyCostSpent:   0.5,
			DailyResetAt:     "2024-01-24T00:00:00+00:00",
			MonthlyCostLimit: ToPtr(100.0),
			MonthlyCostSpent: 12.25,
			MonthlyResetAt:   "2024-02-01T00:00:00+00:00",
		},
	}
	assert.Equal(t, expected, members)
}

func TestAIService_ListUsageMembers_invalidJSON(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/users/1/ai/usage/members", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `invalid json`)
	})

	members, _, err := client.AI.ListUsageMembers(context.Background(), 1, nil)
	require.Error(t, err)
	assert.Nil(t, members)
}

func TestAIService_GetUsageMember(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/users/1/ai/usage/members/5"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		fmt.Fprint(w, `{
			"data": {
				"user": {
					"id": 5,
					"username": "jane",
					"fullName": "Jane Doe",
					"avatarUrl": ""
				},
				"dailyCostLimit": 0,
				"dailyCostSpent": 0,
				"dailyResetAt": "2024-01-24T00:00:00+00:00",
				"monthlyCostLimit": null,
				"monthlyCostSpent": 0,
				"monthlyResetAt": "2024-02-01T00:00:00+00:00"
			}
		}`)
	})

	member, resp, err := client.AI.GetUsageMember(context.Background(), 5, 1)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.AIUsageMember{
		User:           &model.ShortUser{ID: 5, Username: "jane", FullName: "Jane Doe"},
		DailyCostLimit: ToPtr(0.0),
		DailyResetAt:   "2024-01-24T00:00:00+00:00",
		MonthlyResetAt: "2024-02-01T00:00:00+00:00",
	}
	assert.Equal(t, expected, member)
}
