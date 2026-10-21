# \AIGatewayInferenceAPI

All URIs are relative to *https://api.timeweb.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelAiGatewayBatch**](AIGatewayInferenceAPI.md#CancelAiGatewayBatch) | **Post** /v1/batches/{batch_id}/cancel | Отмена батча
[**CancelAiGatewayResponse**](AIGatewayInferenceAPI.md#CancelAiGatewayResponse) | **Post** /v1/responses/{response_id}/cancel | Отмена фонового ответа (Responses API)
[**CreateAiGatewayBatch**](AIGatewayInferenceAPI.md#CreateAiGatewayBatch) | **Post** /v1/batches | Создание батча
[**CreateAiGatewayChatCompletion**](AIGatewayInferenceAPI.md#CreateAiGatewayChatCompletion) | **Post** /v1/chat/completions | Генерация ответа в диалоге (Chat Completions)
[**CreateAiGatewayCompletion**](AIGatewayInferenceAPI.md#CreateAiGatewayCompletion) | **Post** /v1/completions | Продолжение текста (Completions)
[**CreateAiGatewayEmbedding**](AIGatewayInferenceAPI.md#CreateAiGatewayEmbedding) | **Post** /v1/embeddings | Получение эмбеддингов
[**CreateAiGatewayImage**](AIGatewayInferenceAPI.md#CreateAiGatewayImage) | **Post** /v1/images/generations | Генерация изображений
[**CreateAiGatewayImageEdit**](AIGatewayInferenceAPI.md#CreateAiGatewayImageEdit) | **Post** /v1/images/edits | Редактирование изображений
[**CreateAiGatewayResponse**](AIGatewayInferenceAPI.md#CreateAiGatewayResponse) | **Post** /v1/responses | Создание ответа (Responses API)
[**CreateAiGatewaySpeech**](AIGatewayInferenceAPI.md#CreateAiGatewaySpeech) | **Post** /v1/audio/speech | Синтез речи
[**CreateAiGatewayTranscription**](AIGatewayInferenceAPI.md#CreateAiGatewayTranscription) | **Post** /v1/audio/transcriptions | Распознавание речи
[**DeleteAiGatewayFile**](AIGatewayInferenceAPI.md#DeleteAiGatewayFile) | **Delete** /v1/files/{file_id} | Удаление файла
[**DeleteAiGatewayResponse**](AIGatewayInferenceAPI.md#DeleteAiGatewayResponse) | **Delete** /v1/responses/{response_id} | Удаление ответа (Responses API)
[**DownloadAiGatewayFileContent**](AIGatewayInferenceAPI.md#DownloadAiGatewayFileContent) | **Get** /v1/files/{file_id}/content | Скачивание файла
[**GetAiGatewayBatch**](AIGatewayInferenceAPI.md#GetAiGatewayBatch) | **Get** /v1/batches/{batch_id} | Получение батча
[**GetAiGatewayBatches**](AIGatewayInferenceAPI.md#GetAiGatewayBatches) | **Get** /v1/batches | Получение списка батчей
[**GetAiGatewayFile**](AIGatewayInferenceAPI.md#GetAiGatewayFile) | **Get** /v1/files/{file_id} | Получение информации о файле
[**GetAiGatewayModels**](AIGatewayInferenceAPI.md#GetAiGatewayModels) | **Get** /v1/models | Получение списка моделей AI Gateway
[**GetAiGatewayResponse**](AIGatewayInferenceAPI.md#GetAiGatewayResponse) | **Get** /v1/responses/{response_id} | Получение ответа (Responses API)
[**UploadAiGatewayFile**](AIGatewayInferenceAPI.md#UploadAiGatewayFile) | **Post** /v1/files | Загрузка файла



## CancelAiGatewayBatch

> AiGatewayBatch CancelAiGatewayBatch(ctx, batchId).Execute()

Отмена батча



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    batchId := "batch_abc123" // string | ID батча

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIGatewayInferenceAPI.CancelAiGatewayBatch(context.Background(), batchId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayInferenceAPI.CancelAiGatewayBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelAiGatewayBatch`: AiGatewayBatch
    fmt.Fprintf(os.Stdout, "Response from `AIGatewayInferenceAPI.CancelAiGatewayBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchId** | **string** | ID батча | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelAiGatewayBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AiGatewayBatch**](AiGatewayBatch.md)

### Authorization

[AiGatewayKey](../README.md#AiGatewayKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelAiGatewayResponse

> AiGatewayResponse CancelAiGatewayResponse(ctx, responseId).Execute()

Отмена фонового ответа (Responses API)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    responseId := "resp_abc123" // string | ID ответа Responses API

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIGatewayInferenceAPI.CancelAiGatewayResponse(context.Background(), responseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayInferenceAPI.CancelAiGatewayResponse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelAiGatewayResponse`: AiGatewayResponse
    fmt.Fprintf(os.Stdout, "Response from `AIGatewayInferenceAPI.CancelAiGatewayResponse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**responseId** | **string** | ID ответа Responses API | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelAiGatewayResponseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AiGatewayResponse**](AiGatewayResponse.md)

### Authorization

[AiGatewayKey](../README.md#AiGatewayKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAiGatewayBatch

> AiGatewayBatch CreateAiGatewayBatch(ctx).AiGatewayBatchCreateRequest(aiGatewayBatchCreateRequest).Execute()

Создание батча



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    aiGatewayBatchCreateRequest := *openapiclient.NewAiGatewayBatchCreateRequest("file-abc123", "/v1/chat/completions", "24h") // AiGatewayBatchCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIGatewayInferenceAPI.CreateAiGatewayBatch(context.Background()).AiGatewayBatchCreateRequest(aiGatewayBatchCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayInferenceAPI.CreateAiGatewayBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAiGatewayBatch`: AiGatewayBatch
    fmt.Fprintf(os.Stdout, "Response from `AIGatewayInferenceAPI.CreateAiGatewayBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAiGatewayBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiGatewayBatchCreateRequest** | [**AiGatewayBatchCreateRequest**](AiGatewayBatchCreateRequest.md) |  | 

### Return type

[**AiGatewayBatch**](AiGatewayBatch.md)

### Authorization

[AiGatewayKey](../README.md#AiGatewayKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAiGatewayChatCompletion

> AiGatewayChatCompletion CreateAiGatewayChatCompletion(ctx).RequestBody(requestBody).Execute()

Генерация ответа в диалоге (Chat Completions)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIGatewayInferenceAPI.CreateAiGatewayChatCompletion(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayInferenceAPI.CreateAiGatewayChatCompletion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAiGatewayChatCompletion`: AiGatewayChatCompletion
    fmt.Fprintf(os.Stdout, "Response from `AIGatewayInferenceAPI.CreateAiGatewayChatCompletion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAiGatewayChatCompletionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]interface{}** |  | 

### Return type

[**AiGatewayChatCompletion**](AiGatewayChatCompletion.md)

### Authorization

[AiGatewayKey](../README.md#AiGatewayKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAiGatewayCompletion

> AiGatewayCompletion CreateAiGatewayCompletion(ctx).RequestBody(requestBody).Execute()

Продолжение текста (Completions)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIGatewayInferenceAPI.CreateAiGatewayCompletion(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayInferenceAPI.CreateAiGatewayCompletion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAiGatewayCompletion`: AiGatewayCompletion
    fmt.Fprintf(os.Stdout, "Response from `AIGatewayInferenceAPI.CreateAiGatewayCompletion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAiGatewayCompletionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]interface{}** |  | 

### Return type

[**AiGatewayCompletion**](AiGatewayCompletion.md)

### Authorization

[AiGatewayKey](../README.md#AiGatewayKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAiGatewayEmbedding

> AiGatewayEmbeddingList CreateAiGatewayEmbedding(ctx).RequestBody(requestBody).Execute()

Получение эмбеддингов



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIGatewayInferenceAPI.CreateAiGatewayEmbedding(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayInferenceAPI.CreateAiGatewayEmbedding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAiGatewayEmbedding`: AiGatewayEmbeddingList
    fmt.Fprintf(os.Stdout, "Response from `AIGatewayInferenceAPI.CreateAiGatewayEmbedding`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAiGatewayEmbeddingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]interface{}** |  | 

### Return type

[**AiGatewayEmbeddingList**](AiGatewayEmbeddingList.md)

### Authorization

[AiGatewayKey](../README.md#AiGatewayKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAiGatewayImage

> AiGatewayImagesResponse CreateAiGatewayImage(ctx).RequestBody(requestBody).Execute()

Генерация изображений



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIGatewayInferenceAPI.CreateAiGatewayImage(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayInferenceAPI.CreateAiGatewayImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAiGatewayImage`: AiGatewayImagesResponse
    fmt.Fprintf(os.Stdout, "Response from `AIGatewayInferenceAPI.CreateAiGatewayImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAiGatewayImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]interface{}** |  | 

### Return type

[**AiGatewayImagesResponse**](AiGatewayImagesResponse.md)

### Authorization

[AiGatewayKey](../README.md#AiGatewayKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAiGatewayImageEdit

> AiGatewayImagesResponse CreateAiGatewayImageEdit(ctx).Model(model).Image(image).Prompt(prompt).Mask(mask).N(n).Size(size).Quality(quality).Execute()

Редактирование изображений



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    model := "model_example" // string | ID модели: GPT Image (`openai/gpt-image-*`) или FLUX.2 (`black_forest_labs/flux-2-*`)
    image := os.NewFile(1234, "some_file") // *os.File | Исходное изображение или несколько файлов в поле `image[]`
    prompt := "prompt_example" // string | Что изменить
    mask := os.NewFile(1234, "some_file") // *os.File | Маска: прозрачные области будут изменены (optional)
    n := int32(56) // int32 | Количество изображений (optional) (default to 1)
    size := "size_example" // string | Размер изображения (optional)
    quality := "quality_example" // string | Качество изображения (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIGatewayInferenceAPI.CreateAiGatewayImageEdit(context.Background()).Model(model).Image(image).Prompt(prompt).Mask(mask).N(n).Size(size).Quality(quality).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayInferenceAPI.CreateAiGatewayImageEdit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAiGatewayImageEdit`: AiGatewayImagesResponse
    fmt.Fprintf(os.Stdout, "Response from `AIGatewayInferenceAPI.CreateAiGatewayImageEdit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAiGatewayImageEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **model** | **string** | ID модели: GPT Image (&#x60;openai/gpt-image-*&#x60;) или FLUX.2 (&#x60;black_forest_labs/flux-2-*&#x60;) | 
 **image** | ***os.File** | Исходное изображение или несколько файлов в поле &#x60;image[]&#x60; | 
 **prompt** | **string** | Что изменить | 
 **mask** | ***os.File** | Маска: прозрачные области будут изменены | 
 **n** | **int32** | Количество изображений | [default to 1]
 **size** | **string** | Размер изображения | 
 **quality** | **string** | Качество изображения | 

### Return type

[**AiGatewayImagesResponse**](AiGatewayImagesResponse.md)

### Authorization

[AiGatewayKey](../README.md#AiGatewayKey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAiGatewayResponse

> AiGatewayResponse CreateAiGatewayResponse(ctx).RequestBody(requestBody).Execute()

Создание ответа (Responses API)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIGatewayInferenceAPI.CreateAiGatewayResponse(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayInferenceAPI.CreateAiGatewayResponse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAiGatewayResponse`: AiGatewayResponse
    fmt.Fprintf(os.Stdout, "Response from `AIGatewayInferenceAPI.CreateAiGatewayResponse`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAiGatewayResponseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]interface{}** |  | 

### Return type

[**AiGatewayResponse**](AiGatewayResponse.md)

### Authorization

[AiGatewayKey](../README.md#AiGatewayKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAiGatewaySpeech

> *os.File CreateAiGatewaySpeech(ctx).RequestBody(requestBody).Execute()

Синтез речи



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIGatewayInferenceAPI.CreateAiGatewaySpeech(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayInferenceAPI.CreateAiGatewaySpeech``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAiGatewaySpeech`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `AIGatewayInferenceAPI.CreateAiGatewaySpeech`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAiGatewaySpeechRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]interface{}** |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[AiGatewayKey](../README.md#AiGatewayKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: audio/mpeg, audio/wav, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAiGatewayTranscription

> AiGatewayTranscription CreateAiGatewayTranscription(ctx).File(file).Model(model).Language(language).Prompt(prompt).ResponseFormat(responseFormat).Temperature(temperature).Execute()

Распознавание речи



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    file := os.NewFile(1234, "some_file") // *os.File | Аудиофайл
    model := "model_example" // string | ID модели распознавания речи (`*-transcribe`)
    language := "language_example" // string | Язык аудио в формате ISO-639-1 (optional)
    prompt := "prompt_example" // string | Подсказка для распознавания: термины, имена, стиль (optional)
    responseFormat := "responseFormat_example" // string | Формат ответа (optional)
    temperature := float32(8.14) // float32 | Температура выборки (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIGatewayInferenceAPI.CreateAiGatewayTranscription(context.Background()).File(file).Model(model).Language(language).Prompt(prompt).ResponseFormat(responseFormat).Temperature(temperature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayInferenceAPI.CreateAiGatewayTranscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAiGatewayTranscription`: AiGatewayTranscription
    fmt.Fprintf(os.Stdout, "Response from `AIGatewayInferenceAPI.CreateAiGatewayTranscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAiGatewayTranscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | Аудиофайл | 
 **model** | **string** | ID модели распознавания речи (&#x60;*-transcribe&#x60;) | 
 **language** | **string** | Язык аудио в формате ISO-639-1 | 
 **prompt** | **string** | Подсказка для распознавания: термины, имена, стиль | 
 **responseFormat** | **string** | Формат ответа | 
 **temperature** | **float32** | Температура выборки | 

### Return type

[**AiGatewayTranscription**](AiGatewayTranscription.md)

### Authorization

[AiGatewayKey](../README.md#AiGatewayKey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAiGatewayFile

> AiGatewayFileDeleted DeleteAiGatewayFile(ctx, fileId).Execute()

Удаление файла



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    fileId := "file-abc123" // string | ID файла

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIGatewayInferenceAPI.DeleteAiGatewayFile(context.Background(), fileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayInferenceAPI.DeleteAiGatewayFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAiGatewayFile`: AiGatewayFileDeleted
    fmt.Fprintf(os.Stdout, "Response from `AIGatewayInferenceAPI.DeleteAiGatewayFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | ID файла | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAiGatewayFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AiGatewayFileDeleted**](AiGatewayFileDeleted.md)

### Authorization

[AiGatewayKey](../README.md#AiGatewayKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAiGatewayResponse

> AiGatewayResponseDeleted DeleteAiGatewayResponse(ctx, responseId).Execute()

Удаление ответа (Responses API)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    responseId := "resp_abc123" // string | ID ответа Responses API

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIGatewayInferenceAPI.DeleteAiGatewayResponse(context.Background(), responseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayInferenceAPI.DeleteAiGatewayResponse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAiGatewayResponse`: AiGatewayResponseDeleted
    fmt.Fprintf(os.Stdout, "Response from `AIGatewayInferenceAPI.DeleteAiGatewayResponse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**responseId** | **string** | ID ответа Responses API | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAiGatewayResponseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AiGatewayResponseDeleted**](AiGatewayResponseDeleted.md)

### Authorization

[AiGatewayKey](../README.md#AiGatewayKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadAiGatewayFileContent

> *os.File DownloadAiGatewayFileContent(ctx, fileId).Execute()

Скачивание файла



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    fileId := "file-abc123" // string | ID файла

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIGatewayInferenceAPI.DownloadAiGatewayFileContent(context.Background(), fileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayInferenceAPI.DownloadAiGatewayFileContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadAiGatewayFileContent`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `AIGatewayInferenceAPI.DownloadAiGatewayFileContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | ID файла | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadAiGatewayFileContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[AiGatewayKey](../README.md#AiGatewayKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiGatewayBatch

> AiGatewayBatch GetAiGatewayBatch(ctx, batchId).Execute()

Получение батча



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    batchId := "batch_abc123" // string | ID батча

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIGatewayInferenceAPI.GetAiGatewayBatch(context.Background(), batchId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayInferenceAPI.GetAiGatewayBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAiGatewayBatch`: AiGatewayBatch
    fmt.Fprintf(os.Stdout, "Response from `AIGatewayInferenceAPI.GetAiGatewayBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchId** | **string** | ID батча | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiGatewayBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AiGatewayBatch**](AiGatewayBatch.md)

### Authorization

[AiGatewayKey](../README.md#AiGatewayKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiGatewayBatches

> AiGatewayBatchList GetAiGatewayBatches(ctx).Limit(limit).After(after).Execute()

Получение списка батчей



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    limit := int32(20) // int32 | Количество батчей на странице (optional)
    after := "batch_abc123" // string | Вернуть батчи после батча с этим ID (значение `last_id` предыдущей страницы) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIGatewayInferenceAPI.GetAiGatewayBatches(context.Background()).Limit(limit).After(after).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayInferenceAPI.GetAiGatewayBatches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAiGatewayBatches`: AiGatewayBatchList
    fmt.Fprintf(os.Stdout, "Response from `AIGatewayInferenceAPI.GetAiGatewayBatches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAiGatewayBatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Количество батчей на странице | 
 **after** | **string** | Вернуть батчи после батча с этим ID (значение &#x60;last_id&#x60; предыдущей страницы) | 

### Return type

[**AiGatewayBatchList**](AiGatewayBatchList.md)

### Authorization

[AiGatewayKey](../README.md#AiGatewayKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiGatewayFile

> AiGatewayFile GetAiGatewayFile(ctx, fileId).Execute()

Получение информации о файле



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    fileId := "file-abc123" // string | ID файла

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIGatewayInferenceAPI.GetAiGatewayFile(context.Background(), fileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayInferenceAPI.GetAiGatewayFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAiGatewayFile`: AiGatewayFile
    fmt.Fprintf(os.Stdout, "Response from `AIGatewayInferenceAPI.GetAiGatewayFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | ID файла | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiGatewayFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AiGatewayFile**](AiGatewayFile.md)

### Authorization

[AiGatewayKey](../README.md#AiGatewayKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiGatewayModels

> AiGatewayModelList GetAiGatewayModels(ctx).Execute()

Получение списка моделей AI Gateway



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIGatewayInferenceAPI.GetAiGatewayModels(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayInferenceAPI.GetAiGatewayModels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAiGatewayModels`: AiGatewayModelList
    fmt.Fprintf(os.Stdout, "Response from `AIGatewayInferenceAPI.GetAiGatewayModels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiGatewayModelsRequest struct via the builder pattern


### Return type

[**AiGatewayModelList**](AiGatewayModelList.md)

### Authorization

[AiGatewayKey](../README.md#AiGatewayKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiGatewayResponse

> AiGatewayResponse GetAiGatewayResponse(ctx, responseId).Execute()

Получение ответа (Responses API)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    responseId := "resp_abc123" // string | ID ответа Responses API

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIGatewayInferenceAPI.GetAiGatewayResponse(context.Background(), responseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayInferenceAPI.GetAiGatewayResponse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAiGatewayResponse`: AiGatewayResponse
    fmt.Fprintf(os.Stdout, "Response from `AIGatewayInferenceAPI.GetAiGatewayResponse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**responseId** | **string** | ID ответа Responses API | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiGatewayResponseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AiGatewayResponse**](AiGatewayResponse.md)

### Authorization

[AiGatewayKey](../README.md#AiGatewayKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadAiGatewayFile

> AiGatewayFile UploadAiGatewayFile(ctx).File(file).Purpose(purpose).Model(model).TargetModelNames(targetModelNames).Execute()

Загрузка файла



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    file := os.NewFile(1234, "some_file") // *os.File | Файл. Для Batch API — JSONL, по одному запросу в строке
    purpose := "purpose_example" // string | Назначение: `batch` — входной файл Batch API, `user_data` и другие назначения OpenAI — файлы для передачи модели
    model := "model_example" // string | Модель, для которой загружается файл. Обязателен, если не передан `target_model_names` (optional)
    targetModelNames := "targetModelNames_example" // string | Список моделей через запятую — вместо `model` (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIGatewayInferenceAPI.UploadAiGatewayFile(context.Background()).File(file).Purpose(purpose).Model(model).TargetModelNames(targetModelNames).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayInferenceAPI.UploadAiGatewayFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadAiGatewayFile`: AiGatewayFile
    fmt.Fprintf(os.Stdout, "Response from `AIGatewayInferenceAPI.UploadAiGatewayFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadAiGatewayFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | Файл. Для Batch API — JSONL, по одному запросу в строке | 
 **purpose** | **string** | Назначение: &#x60;batch&#x60; — входной файл Batch API, &#x60;user_data&#x60; и другие назначения OpenAI — файлы для передачи модели | 
 **model** | **string** | Модель, для которой загружается файл. Обязателен, если не передан &#x60;target_model_names&#x60; | 
 **targetModelNames** | **string** | Список моделей через запятую — вместо &#x60;model&#x60; | 

### Return type

[**AiGatewayFile**](AiGatewayFile.md)

### Authorization

[AiGatewayKey](../README.md#AiGatewayKey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

