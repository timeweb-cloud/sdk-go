# ConfigDomains

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aliases** | Pointer to **[]string** | Список доменов ресурса. Первым в списке всегда идет технический домен &#x60;cdn_domain&#x60;, вторым может быть ваш собственный домен. Чтобы собственный домен заработал, направьте на &#x60;cdn_domain&#x60; CNAME-запись. &#x60;null&#x60; — оставить только технический домен. | [optional] 

## Methods

### NewConfigDomains

`func NewConfigDomains() *ConfigDomains`

NewConfigDomains instantiates a new ConfigDomains object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigDomainsWithDefaults

`func NewConfigDomainsWithDefaults() *ConfigDomains`

NewConfigDomainsWithDefaults instantiates a new ConfigDomains object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliases

`func (o *ConfigDomains) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *ConfigDomains) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *ConfigDomains) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *ConfigDomains) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### SetAliasesNil

`func (o *ConfigDomains) SetAliasesNil(b bool)`

 SetAliasesNil sets the value for Aliases to be an explicit nil

### UnsetAliases
`func (o *ConfigDomains) UnsetAliases()`

UnsetAliases ensures that no value is present for Aliases, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


