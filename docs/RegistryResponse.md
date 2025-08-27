# RegistryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | Pointer to **string** | ID запроса | [optional] 
**ContainerRegistry** | [**RegistryOut**](RegistryOut.md) |  | 

## Methods

### NewRegistryResponse

`func NewRegistryResponse(containerRegistry RegistryOut, ) *RegistryResponse`

NewRegistryResponse instantiates a new RegistryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryResponseWithDefaults

`func NewRegistryResponseWithDefaults() *RegistryResponse`

NewRegistryResponseWithDefaults instantiates a new RegistryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *RegistryResponse) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *RegistryResponse) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *RegistryResponse) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.

### HasResponseId

`func (o *RegistryResponse) HasResponseId() bool`

HasResponseId returns a boolean if a field has been set.

### GetContainerRegistry

`func (o *RegistryResponse) GetContainerRegistry() RegistryOut`

GetContainerRegistry returns the ContainerRegistry field if non-nil, zero value otherwise.

### GetContainerRegistryOk

`func (o *RegistryResponse) GetContainerRegistryOk() (*RegistryOut, bool)`

GetContainerRegistryOk returns a tuple with the ContainerRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerRegistry

`func (o *RegistryResponse) SetContainerRegistry(v RegistryOut)`

SetContainerRegistry sets ContainerRegistry field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


