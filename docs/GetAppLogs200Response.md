# GetAppLogs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**AppLogs** | **[]string** |  | 

## Methods

### NewGetAppLogs200Response

`func NewGetAppLogs200Response(meta Meta, appLogs []string, ) *GetAppLogs200Response`

NewGetAppLogs200Response instantiates a new GetAppLogs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAppLogs200ResponseWithDefaults

`func NewGetAppLogs200ResponseWithDefaults() *GetAppLogs200Response`

NewGetAppLogs200ResponseWithDefaults instantiates a new GetAppLogs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetAppLogs200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAppLogs200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAppLogs200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetAppLogs

`func (o *GetAppLogs200Response) GetAppLogs() []string`

GetAppLogs returns the AppLogs field if non-nil, zero value otherwise.

### GetAppLogsOk

`func (o *GetAppLogs200Response) GetAppLogsOk() (*[]string, bool)`

GetAppLogsOk returns a tuple with the AppLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppLogs

`func (o *GetAppLogs200Response) SetAppLogs(v []string)`

SetAppLogs sets AppLogs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


