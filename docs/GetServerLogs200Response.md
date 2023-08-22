# GetServerLogs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**ServerLogs** | [**[]ServerLog**](ServerLog.md) |  | 

## Methods

### NewGetServerLogs200Response

`func NewGetServerLogs200Response(meta Meta, serverLogs []ServerLog, ) *GetServerLogs200Response`

NewGetServerLogs200Response instantiates a new GetServerLogs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServerLogs200ResponseWithDefaults

`func NewGetServerLogs200ResponseWithDefaults() *GetServerLogs200Response`

NewGetServerLogs200ResponseWithDefaults instantiates a new GetServerLogs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetServerLogs200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetServerLogs200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetServerLogs200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetServerLogs

`func (o *GetServerLogs200Response) GetServerLogs() []ServerLog`

GetServerLogs returns the ServerLogs field if non-nil, zero value otherwise.

### GetServerLogsOk

`func (o *GetServerLogs200Response) GetServerLogsOk() (*[]ServerLog, bool)`

GetServerLogsOk returns a tuple with the ServerLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerLogs

`func (o *GetServerLogs200Response) SetServerLogs(v []ServerLog)`

SetServerLogs sets ServerLogs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


