# \ImagesAPI

All URIs are relative to *https://api.timeweb.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateImage**](ImagesAPI.md#CreateImage) | **Post** /api/v1/images | Создание образа
[**CreateImageDownloadUrl**](ImagesAPI.md#CreateImageDownloadUrl) | **Post** /api/v1/images/{image_id}/download-url | Создание ссылки на скачивание образа
[**DeleteImage**](ImagesAPI.md#DeleteImage) | **Delete** /api/v1/images/{image_id} | Удаление образа
[**DeleteImageDownloadURL**](ImagesAPI.md#DeleteImageDownloadURL) | **Delete** /api/v1/images/{image_id}/download-url/{image_url_id} | Удаление ссылки на образ
[**GetImage**](ImagesAPI.md#GetImage) | **Get** /api/v1/images/{image_id} | Получение информации о образе
[**GetImageDownloadURL**](ImagesAPI.md#GetImageDownloadURL) | **Get** /api/v1/images/{image_id}/download-url/{image_url_id} | Получение информации о ссылке на скачивание образа
[**GetImageDownloadURLs**](ImagesAPI.md#GetImageDownloadURLs) | **Get** /api/v1/images/{image_id}/download-url | Получение информации о ссылках на скачивание образов
[**GetImages**](ImagesAPI.md#GetImages) | **Get** /api/v1/images | Получение списка образов
[**UpdateImage**](ImagesAPI.md#UpdateImage) | **Patch** /api/v1/images/{image_id} | Обновление информации о образе
[**UploadImage**](ImagesAPI.md#UploadImage) | **Post** /api/v1/images/{image_id} | Загрузка образа



## CreateImage

> ImageOutResponse CreateImage(ctx).ImageInAPI(imageInAPI).Execute()

Создание образа



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
    imageInAPI := *openapiclient.NewImageInAPI(openapiclient.location("ru-1"), openapiclient.OS("centos")) // ImageInAPI | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesAPI.CreateImage(context.Background()).ImageInAPI(imageInAPI).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.CreateImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateImage`: ImageOutResponse
    fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.CreateImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imageInAPI** | [**ImageInAPI**](ImageInAPI.md) |  | 

### Return type

[**ImageOutResponse**](ImageOutResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateImageDownloadUrl

> ImageDownloadResponse CreateImageDownloadUrl(ctx, imageId).ImageUrlIn(imageUrlIn).Execute()

Создание ссылки на скачивание образа



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
    imageId := "imageId_example" // string | ID образа.
    imageUrlIn := *openapiclient.NewImageUrlIn() // ImageUrlIn | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesAPI.CreateImageDownloadUrl(context.Background(), imageId).ImageUrlIn(imageUrlIn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.CreateImageDownloadUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateImageDownloadUrl`: ImageDownloadResponse
    fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.CreateImageDownloadUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string** | ID образа. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateImageDownloadUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **imageUrlIn** | [**ImageUrlIn**](ImageUrlIn.md) |  | 

### Return type

[**ImageDownloadResponse**](ImageDownloadResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteImage

> DeleteImage(ctx, imageId).Execute()

Удаление образа



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
    imageId := "imageId_example" // string | ID образа.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ImagesAPI.DeleteImage(context.Background(), imageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.DeleteImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string** | ID образа. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteImageDownloadURL

> DeleteImageDownloadURL(ctx, imageId, imageUrlId).Execute()

Удаление ссылки на образ



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
    imageId := "imageId_example" // string | ID образа.
    imageUrlId := "imageUrlId_example" // string | ID ссылки.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ImagesAPI.DeleteImageDownloadURL(context.Background(), imageId, imageUrlId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.DeleteImageDownloadURL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string** | ID образа. | 
**imageUrlId** | **string** | ID ссылки. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteImageDownloadURLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImage

> ImageOutResponse GetImage(ctx, imageId).Execute()

Получение информации о образе



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
    imageId := "imageId_example" // string | ID образа.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesAPI.GetImage(context.Background(), imageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.GetImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetImage`: ImageOutResponse
    fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.GetImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string** | ID образа. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ImageOutResponse**](ImageOutResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageDownloadURL

> ImageDownloadResponse GetImageDownloadURL(ctx, imageId, imageUrlId).Execute()

Получение информации о ссылке на скачивание образа



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
    imageId := "imageId_example" // string | ID образа.
    imageUrlId := "imageUrlId_example" // string | ID ссылки.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesAPI.GetImageDownloadURL(context.Background(), imageId, imageUrlId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.GetImageDownloadURL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetImageDownloadURL`: ImageDownloadResponse
    fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.GetImageDownloadURL`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string** | ID образа. | 
**imageUrlId** | **string** | ID ссылки. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageDownloadURLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ImageDownloadResponse**](ImageDownloadResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageDownloadURLs

> ImageDownloadsResponse GetImageDownloadURLs(ctx, imageId).Limit(limit).Offset(offset).Execute()

Получение информации о ссылках на скачивание образов



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
    imageId := "imageId_example" // string | ID образа.
    limit := int32(56) // int32 |  (optional) (default to 100)
    offset := int32(56) // int32 |  (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesAPI.GetImageDownloadURLs(context.Background(), imageId).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.GetImageDownloadURLs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetImageDownloadURLs`: ImageDownloadsResponse
    fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.GetImageDownloadURLs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string** | ID образа. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageDownloadURLsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 100]
 **offset** | **int32** |  | [default to 0]

### Return type

[**ImageDownloadsResponse**](ImageDownloadsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImages

> ImagesOutResponse GetImages(ctx).Limit(limit).Offset(offset).Execute()

Получение списка образов



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
    limit := int32(56) // int32 |  (optional) (default to 100)
    offset := int32(56) // int32 |  (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesAPI.GetImages(context.Background()).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.GetImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetImages`: ImagesOutResponse
    fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.GetImages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 100]
 **offset** | **int32** |  | [default to 0]

### Return type

[**ImagesOutResponse**](ImagesOutResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateImage

> ImageOutResponse UpdateImage(ctx, imageId).ImageUpdateAPI(imageUpdateAPI).Execute()

Обновление информации о образе



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
    imageId := "imageId_example" // string | ID образа.
    imageUpdateAPI := *openapiclient.NewImageUpdateAPI() // ImageUpdateAPI | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesAPI.UpdateImage(context.Background(), imageId).ImageUpdateAPI(imageUpdateAPI).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.UpdateImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateImage`: ImageOutResponse
    fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.UpdateImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string** | ID образа. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **imageUpdateAPI** | [**ImageUpdateAPI**](ImageUpdateAPI.md) |  | 

### Return type

[**ImageOutResponse**](ImageOutResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadImage

> UploadSuccessfulResponse UploadImage(ctx, imageId).ContentDisposition(contentDisposition).Execute()

Загрузка образа



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
    imageId := "imageId_example" // string | 
    contentDisposition := "contentDisposition_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesAPI.UploadImage(context.Background(), imageId).ContentDisposition(contentDisposition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.UploadImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadImage`: UploadSuccessfulResponse
    fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.UploadImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentDisposition** | **string** |  | 

### Return type

[**UploadSuccessfulResponse**](UploadSuccessfulResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

