# InfoServicePrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Название сервиса | [optional] 
**Description** | Pointer to **string** | Описание сервиса | [optional] 
**IsMounted** | Pointer to **bool** | Флаг, указывающий, подключен ли сервис | [optional] 
**Size** | Pointer to **float32** | Размер в ГБ | [optional] 
**Type** | Pointer to **string** | Тип кластера базы данных | [optional] 
**Count** | Pointer to **float32** | Количество | [optional] 

## Methods

### NewInfoServicePrice

`func NewInfoServicePrice() *InfoServicePrice`

NewInfoServicePrice instantiates a new InfoServicePrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfoServicePriceWithDefaults

`func NewInfoServicePriceWithDefaults() *InfoServicePrice`

NewInfoServicePriceWithDefaults instantiates a new InfoServicePrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InfoServicePrice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InfoServicePrice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InfoServicePrice) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InfoServicePrice) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InfoServicePrice) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InfoServicePrice) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InfoServicePrice) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InfoServicePrice) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsMounted

`func (o *InfoServicePrice) GetIsMounted() bool`

GetIsMounted returns the IsMounted field if non-nil, zero value otherwise.

### GetIsMountedOk

`func (o *InfoServicePrice) GetIsMountedOk() (*bool, bool)`

GetIsMountedOk returns a tuple with the IsMounted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMounted

`func (o *InfoServicePrice) SetIsMounted(v bool)`

SetIsMounted sets IsMounted field to given value.

### HasIsMounted

`func (o *InfoServicePrice) HasIsMounted() bool`

HasIsMounted returns a boolean if a field has been set.

### GetSize

`func (o *InfoServicePrice) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *InfoServicePrice) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *InfoServicePrice) SetSize(v float32)`

SetSize sets Size field to given value.

### HasSize

`func (o *InfoServicePrice) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *InfoServicePrice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InfoServicePrice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InfoServicePrice) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InfoServicePrice) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCount

`func (o *InfoServicePrice) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *InfoServicePrice) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *InfoServicePrice) SetCount(v float32)`

SetCount sets Count field to given value.

### HasCount

`func (o *InfoServicePrice) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


