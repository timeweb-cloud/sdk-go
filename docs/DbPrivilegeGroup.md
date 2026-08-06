# DbPrivilegeGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Privilege** | [**PropertiesMysql**](PropertiesMysql.md) |  | 
**Group** | **string** | Группа доступа, к которой относится привилегия: &#x60;GLOBAL&#x60; — привилегии уровня всего кластера, &#x60;DATA&#x60; — доступ к данным, &#x60;STRUCTURE&#x60; — изменение структуры, &#x60;OTHERS&#x60; — прочие привилегии. | 

## Methods

### NewDbPrivilegeGroup

`func NewDbPrivilegeGroup(privilege PropertiesMysql, group string, ) *DbPrivilegeGroup`

NewDbPrivilegeGroup instantiates a new DbPrivilegeGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbPrivilegeGroupWithDefaults

`func NewDbPrivilegeGroupWithDefaults() *DbPrivilegeGroup`

NewDbPrivilegeGroupWithDefaults instantiates a new DbPrivilegeGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivilege

`func (o *DbPrivilegeGroup) GetPrivilege() PropertiesMysql`

GetPrivilege returns the Privilege field if non-nil, zero value otherwise.

### GetPrivilegeOk

`func (o *DbPrivilegeGroup) GetPrivilegeOk() (*PropertiesMysql, bool)`

GetPrivilegeOk returns a tuple with the Privilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilege

`func (o *DbPrivilegeGroup) SetPrivilege(v PropertiesMysql)`

SetPrivilege sets Privilege field to given value.


### GetGroup

`func (o *DbPrivilegeGroup) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *DbPrivilegeGroup) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *DbPrivilegeGroup) SetGroup(v string)`

SetGroup sets Group field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


