# FirewallGroupInAPI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Имя группы правил. | 
**Description** | Pointer to **string** | Описание группы правил. | [optional] 

## Methods

### NewFirewallGroupInAPI

`func NewFirewallGroupInAPI(name string, ) *FirewallGroupInAPI`

NewFirewallGroupInAPI instantiates a new FirewallGroupInAPI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallGroupInAPIWithDefaults

`func NewFirewallGroupInAPIWithDefaults() *FirewallGroupInAPI`

NewFirewallGroupInAPIWithDefaults instantiates a new FirewallGroupInAPI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FirewallGroupInAPI) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FirewallGroupInAPI) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FirewallGroupInAPI) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *FirewallGroupInAPI) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FirewallGroupInAPI) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FirewallGroupInAPI) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FirewallGroupInAPI) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


