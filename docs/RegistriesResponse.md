# RegistriesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | Pointer to **string** | ID запроса | [optional] 
**ContainerRegistryList** | Pointer to [**[]RegistryOut**](RegistryOut.md) | Реестр контейнеров | [optional] 

## Methods

### NewRegistriesResponse

`func NewRegistriesResponse() *RegistriesResponse`

NewRegistriesResponse instantiates a new RegistriesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistriesResponseWithDefaults

`func NewRegistriesResponseWithDefaults() *RegistriesResponse`

NewRegistriesResponseWithDefaults instantiates a new RegistriesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *RegistriesResponse) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *RegistriesResponse) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *RegistriesResponse) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.

### HasResponseId

`func (o *RegistriesResponse) HasResponseId() bool`

HasResponseId returns a boolean if a field has been set.

### GetContainerRegistryList

`func (o *RegistriesResponse) GetContainerRegistryList() []RegistryOut`

GetContainerRegistryList returns the ContainerRegistryList field if non-nil, zero value otherwise.

### GetContainerRegistryListOk

`func (o *RegistriesResponse) GetContainerRegistryListOk() (*[]RegistryOut, bool)`

GetContainerRegistryListOk returns a tuple with the ContainerRegistryList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerRegistryList

`func (o *RegistriesResponse) SetContainerRegistryList(v []RegistryOut)`

SetContainerRegistryList sets ContainerRegistryList field to given value.

### HasContainerRegistryList

`func (o *RegistriesResponse) HasContainerRegistryList() bool`

HasContainerRegistryList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


