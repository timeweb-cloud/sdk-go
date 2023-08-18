# DomainNameServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDelegationAllowed** | **bool** | Это логическое значение, которое показывает включена ли услуга разрешено ли делегирование домена. | 
**Items** | [**[]DomainNameServerItemsInner**](DomainNameServerItemsInner.md) | Список name-серверов | 
**TaskStatus** | **string** | Статус добавления name-сервера. | 

## Methods

### NewDomainNameServer

`func NewDomainNameServer(isDelegationAllowed bool, items []DomainNameServerItemsInner, taskStatus string, ) *DomainNameServer`

NewDomainNameServer instantiates a new DomainNameServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainNameServerWithDefaults

`func NewDomainNameServerWithDefaults() *DomainNameServer`

NewDomainNameServerWithDefaults instantiates a new DomainNameServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDelegationAllowed

`func (o *DomainNameServer) GetIsDelegationAllowed() bool`

GetIsDelegationAllowed returns the IsDelegationAllowed field if non-nil, zero value otherwise.

### GetIsDelegationAllowedOk

`func (o *DomainNameServer) GetIsDelegationAllowedOk() (*bool, bool)`

GetIsDelegationAllowedOk returns a tuple with the IsDelegationAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDelegationAllowed

`func (o *DomainNameServer) SetIsDelegationAllowed(v bool)`

SetIsDelegationAllowed sets IsDelegationAllowed field to given value.


### GetItems

`func (o *DomainNameServer) GetItems() []DomainNameServerItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *DomainNameServer) GetItemsOk() (*[]DomainNameServerItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *DomainNameServer) SetItems(v []DomainNameServerItemsInner)`

SetItems sets Items field to given value.


### GetTaskStatus

`func (o *DomainNameServer) GetTaskStatus() string`

GetTaskStatus returns the TaskStatus field if non-nil, zero value otherwise.

### GetTaskStatusOk

`func (o *DomainNameServer) GetTaskStatusOk() (*string, bool)`

GetTaskStatusOk returns a tuple with the TaskStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskStatus

`func (o *DomainNameServer) SetTaskStatus(v string)`

SetTaskStatus sets TaskStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


