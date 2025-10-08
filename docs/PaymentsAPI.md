# \PaymentsAPI

All URIs are relative to *https://api.timeweb.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFinances**](PaymentsAPI.md#GetFinances) | **Get** /api/v1/account/finances | Получение платежной информации
[**GetLinkCardPayment**](PaymentsAPI.md#GetLinkCardPayment) | **Post** /api/v1/account/payment-link | Получение ссылки на оплату
[**GetServicePrices**](PaymentsAPI.md#GetServicePrices) | **Get** /api/v1/account/services/cost | Получение стоимости сервисов



## GetFinances

> GetFinances200Response GetFinances(ctx).Execute()

Получение платежной информации



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
    resp, r, err := apiClient.PaymentsAPI.GetFinances(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.GetFinances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFinances`: GetFinances200Response
    fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.GetFinances`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinancesRequest struct via the builder pattern


### Return type

[**GetFinances200Response**](GetFinances200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinkCardPayment

> GetLinkCardPayment200Response GetLinkCardPayment(ctx).CreatePayment(createPayment).Execute()

Получение ссылки на оплату



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
    createPayment := *openapiclient.NewCreatePayment(float32(1000), openapiclient.payment-type("card")) // CreatePayment | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentsAPI.GetLinkCardPayment(context.Background()).CreatePayment(createPayment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.GetLinkCardPayment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLinkCardPayment`: GetLinkCardPayment200Response
    fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.GetLinkCardPayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLinkCardPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPayment** | [**CreatePayment**](CreatePayment.md) |  | 

### Return type

[**GetLinkCardPayment200Response**](GetLinkCardPayment200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServicePrices

> GetServicePrices200Response GetServicePrices(ctx).Execute()

Получение стоимости сервисов



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
    resp, r, err := apiClient.PaymentsAPI.GetServicePrices(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.GetServicePrices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServicePrices`: GetServicePrices200Response
    fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.GetServicePrices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetServicePricesRequest struct via the builder pattern


### Return type

[**GetServicePrices200Response**](GetServicePrices200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

