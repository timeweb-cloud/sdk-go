# GetConfigurators200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**ServerConfigurators** | [**[]ServersConfigurator**](ServersConfigurator.md) |  | 
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 

## Methods

### NewGetConfigurators200Response

`func NewGetConfigurators200Response(meta Meta, serverConfigurators []ServersConfigurator, responseId string, ) *GetConfigurators200Response`

NewGetConfigurators200Response instantiates a new GetConfigurators200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConfigurators200ResponseWithDefaults

`func NewGetConfigurators200ResponseWithDefaults() *GetConfigurators200Response`

NewGetConfigurators200ResponseWithDefaults instantiates a new GetConfigurators200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetConfigurators200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetConfigurators200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetConfigurators200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetServerConfigurators

`func (o *GetConfigurators200Response) GetServerConfigurators() []ServersConfigurator`

GetServerConfigurators returns the ServerConfigurators field if non-nil, zero value otherwise.

### GetServerConfiguratorsOk

`func (o *GetConfigurators200Response) GetServerConfiguratorsOk() (*[]ServersConfigurator, bool)`

GetServerConfiguratorsOk returns a tuple with the ServerConfigurators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerConfigurators

`func (o *GetConfigurators200Response) SetServerConfigurators(v []ServersConfigurator)`

SetServerConfigurators sets ServerConfigurators field to given value.


### GetResponseId

`func (o *GetConfigurators200Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *GetConfigurators200Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *GetConfigurators200Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


