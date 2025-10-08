# ServicePrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cost** | **float32** | Стоимость сервиса | 
**TotalCost** | **float32** | Общая стоимость сервиса с учетом всех дополнительных услуг | 
**Type** | [**ServicePriceType**](ServicePriceType.md) |  | 
**ServiceId** | Pointer to **float32** | Идентификатор сервиса | [optional] 
**ProjectId** | Pointer to **NullableFloat32** | Идентификатор проекта | [optional] 
**Services** | Pointer to [**[]ServiceServicePrice**](ServiceServicePrice.md) | Список вложенных сервисов | [optional] 
**Info** | Pointer to [**InfoServicePrice**](InfoServicePrice.md) |  | [optional] 
**Configuration** | Pointer to [**ServicePriceConfiguration**](ServicePriceConfiguration.md) |  | [optional] 

## Methods

### NewServicePrice

`func NewServicePrice(cost float32, totalCost float32, type_ ServicePriceType, ) *ServicePrice`

NewServicePrice instantiates a new ServicePrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePriceWithDefaults

`func NewServicePriceWithDefaults() *ServicePrice`

NewServicePriceWithDefaults instantiates a new ServicePrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCost

`func (o *ServicePrice) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ServicePrice) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ServicePrice) SetCost(v float32)`

SetCost sets Cost field to given value.


### GetTotalCost

`func (o *ServicePrice) GetTotalCost() float32`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *ServicePrice) GetTotalCostOk() (*float32, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *ServicePrice) SetTotalCost(v float32)`

SetTotalCost sets TotalCost field to given value.


### GetType

`func (o *ServicePrice) GetType() ServicePriceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServicePrice) GetTypeOk() (*ServicePriceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServicePrice) SetType(v ServicePriceType)`

SetType sets Type field to given value.


### GetServiceId

`func (o *ServicePrice) GetServiceId() float32`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ServicePrice) GetServiceIdOk() (*float32, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ServicePrice) SetServiceId(v float32)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ServicePrice) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetProjectId

`func (o *ServicePrice) GetProjectId() float32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ServicePrice) GetProjectIdOk() (*float32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ServicePrice) SetProjectId(v float32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ServicePrice) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *ServicePrice) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *ServicePrice) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetServices

`func (o *ServicePrice) GetServices() []ServiceServicePrice`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ServicePrice) GetServicesOk() (*[]ServiceServicePrice, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ServicePrice) SetServices(v []ServiceServicePrice)`

SetServices sets Services field to given value.

### HasServices

`func (o *ServicePrice) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetInfo

`func (o *ServicePrice) GetInfo() InfoServicePrice`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ServicePrice) GetInfoOk() (*InfoServicePrice, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ServicePrice) SetInfo(v InfoServicePrice)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ServicePrice) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetConfiguration

`func (o *ServicePrice) GetConfiguration() ServicePriceConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ServicePrice) GetConfigurationOk() (*ServicePriceConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ServicePrice) SetConfiguration(v ServicePriceConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *ServicePrice) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


