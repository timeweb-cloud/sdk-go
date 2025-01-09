# AppProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID провайдера. | 
**Type** | [**Providers**](Providers.md) |  | 

## Methods

### NewAppProvider

`func NewAppProvider(id string, type_ Providers, ) *AppProvider`

NewAppProvider instantiates a new AppProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppProviderWithDefaults

`func NewAppProviderWithDefaults() *AppProvider`

NewAppProviderWithDefaults instantiates a new AppProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppProvider) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppProvider) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppProvider) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *AppProvider) GetType() Providers`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppProvider) GetTypeOk() (*Providers, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppProvider) SetType(v Providers)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


