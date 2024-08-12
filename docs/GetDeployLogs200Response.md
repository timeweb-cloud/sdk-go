# GetDeployLogs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**DeployLogs** | **[]string** |  | 

## Methods

### NewGetDeployLogs200Response

`func NewGetDeployLogs200Response(meta Meta, deployLogs []string, ) *GetDeployLogs200Response`

NewGetDeployLogs200Response instantiates a new GetDeployLogs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeployLogs200ResponseWithDefaults

`func NewGetDeployLogs200ResponseWithDefaults() *GetDeployLogs200Response`

NewGetDeployLogs200ResponseWithDefaults instantiates a new GetDeployLogs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetDeployLogs200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetDeployLogs200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetDeployLogs200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetDeployLogs

`func (o *GetDeployLogs200Response) GetDeployLogs() []string`

GetDeployLogs returns the DeployLogs field if non-nil, zero value otherwise.

### GetDeployLogsOk

`func (o *GetDeployLogs200Response) GetDeployLogsOk() (*[]string, bool)`

GetDeployLogsOk returns a tuple with the DeployLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployLogs

`func (o *GetDeployLogs200Response) SetDeployLogs(v []string)`

SetDeployLogs sets DeployLogs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


