# ResourceTransfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToProject** | **float32** | ID проекта, куда переносится ресурс. | 
**ResourceId** | **float32** | ID перемещаемого ресурса (сервера, хранилища, кластера, балансировщика, базы данных или выделенного сервера). | 
**ResourceType** | **string** | Тип перемещаемого ресурса. | 

## Methods

### NewResourceTransfer

`func NewResourceTransfer(toProject float32, resourceId float32, resourceType string, ) *ResourceTransfer`

NewResourceTransfer instantiates a new ResourceTransfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceTransferWithDefaults

`func NewResourceTransferWithDefaults() *ResourceTransfer`

NewResourceTransferWithDefaults instantiates a new ResourceTransfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToProject

`func (o *ResourceTransfer) GetToProject() float32`

GetToProject returns the ToProject field if non-nil, zero value otherwise.

### GetToProjectOk

`func (o *ResourceTransfer) GetToProjectOk() (*float32, bool)`

GetToProjectOk returns a tuple with the ToProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToProject

`func (o *ResourceTransfer) SetToProject(v float32)`

SetToProject sets ToProject field to given value.


### GetResourceId

`func (o *ResourceTransfer) GetResourceId() float32`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ResourceTransfer) GetResourceIdOk() (*float32, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ResourceTransfer) SetResourceId(v float32)`

SetResourceId sets ResourceId field to given value.


### GetResourceType

`func (o *ResourceTransfer) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ResourceTransfer) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ResourceTransfer) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


