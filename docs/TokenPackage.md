# TokenPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID пакета | 
**ModelId** | **int32** | ID модели, к которой относится пакет | 
**ServiceTypeId** | **int32** | ID тарифа пакета в биллинге | 
**Name** | **string** | Название пакета | 
**PackageType** | **string** | Тип пакета: &#x60;base&#x60; — пакет подписки, &#x60;additional&#x60; — дополнительный пакет, &#x60;promo&#x60; — промо-пакет | 
**Type** | **string** | Для чего пакет: &#x60;agent&#x60; — AI агент, &#x60;knowledgebase&#x60; — база знаний | 
**TokenAmount** | **int32** | Количество токенов в пакете | 
**Price** | **float32** | Цена пакета в рублях | 
**DurationDays** | **NullableInt32** | Период подписки в днях. &#x60;null&#x60; для дополнительных пакетов | 
**IsAvailable** | **bool** | Пакет доступен для покупки | 

## Methods

### NewTokenPackage

`func NewTokenPackage(id int32, modelId int32, serviceTypeId int32, name string, packageType string, type_ string, tokenAmount int32, price float32, durationDays NullableInt32, isAvailable bool, ) *TokenPackage`

NewTokenPackage instantiates a new TokenPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenPackageWithDefaults

`func NewTokenPackageWithDefaults() *TokenPackage`

NewTokenPackageWithDefaults instantiates a new TokenPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TokenPackage) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TokenPackage) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TokenPackage) SetId(v int32)`

SetId sets Id field to given value.


### GetModelId

`func (o *TokenPackage) GetModelId() int32`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *TokenPackage) GetModelIdOk() (*int32, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *TokenPackage) SetModelId(v int32)`

SetModelId sets ModelId field to given value.


### GetServiceTypeId

`func (o *TokenPackage) GetServiceTypeId() int32`

GetServiceTypeId returns the ServiceTypeId field if non-nil, zero value otherwise.

### GetServiceTypeIdOk

`func (o *TokenPackage) GetServiceTypeIdOk() (*int32, bool)`

GetServiceTypeIdOk returns a tuple with the ServiceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTypeId

`func (o *TokenPackage) SetServiceTypeId(v int32)`

SetServiceTypeId sets ServiceTypeId field to given value.


### GetName

`func (o *TokenPackage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenPackage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenPackage) SetName(v string)`

SetName sets Name field to given value.


### GetPackageType

`func (o *TokenPackage) GetPackageType() string`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *TokenPackage) GetPackageTypeOk() (*string, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *TokenPackage) SetPackageType(v string)`

SetPackageType sets PackageType field to given value.


### GetType

`func (o *TokenPackage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TokenPackage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TokenPackage) SetType(v string)`

SetType sets Type field to given value.


### GetTokenAmount

`func (o *TokenPackage) GetTokenAmount() int32`

GetTokenAmount returns the TokenAmount field if non-nil, zero value otherwise.

### GetTokenAmountOk

`func (o *TokenPackage) GetTokenAmountOk() (*int32, bool)`

GetTokenAmountOk returns a tuple with the TokenAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAmount

`func (o *TokenPackage) SetTokenAmount(v int32)`

SetTokenAmount sets TokenAmount field to given value.


### GetPrice

`func (o *TokenPackage) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *TokenPackage) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *TokenPackage) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetDurationDays

`func (o *TokenPackage) GetDurationDays() int32`

GetDurationDays returns the DurationDays field if non-nil, zero value otherwise.

### GetDurationDaysOk

`func (o *TokenPackage) GetDurationDaysOk() (*int32, bool)`

GetDurationDaysOk returns a tuple with the DurationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationDays

`func (o *TokenPackage) SetDurationDays(v int32)`

SetDurationDays sets DurationDays field to given value.


### SetDurationDaysNil

`func (o *TokenPackage) SetDurationDaysNil(b bool)`

 SetDurationDaysNil sets the value for DurationDays to be an explicit nil

### UnsetDurationDays
`func (o *TokenPackage) UnsetDurationDays()`

UnsetDurationDays ensures that no value is present for DurationDays, not even an explicit nil
### GetIsAvailable

`func (o *TokenPackage) GetIsAvailable() bool`

GetIsAvailable returns the IsAvailable field if non-nil, zero value otherwise.

### GetIsAvailableOk

`func (o *TokenPackage) GetIsAvailableOk() (*bool, bool)`

GetIsAvailableOk returns a tuple with the IsAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAvailable

`func (o *TokenPackage) SetIsAvailable(v bool)`

SetIsAvailable sets IsAvailable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


