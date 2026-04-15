# AddonsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | Pointer to **string** | ID запроса | [optional] 
**Meta** | [**SchemasMeta**](SchemasMeta.md) |  | 
**Addons** | [**[]AddonOut**](AddonOut.md) | Массив дополнений k8s | 

## Methods

### NewAddonsResponse

`func NewAddonsResponse(meta SchemasMeta, addons []AddonOut, ) *AddonsResponse`

NewAddonsResponse instantiates a new AddonsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonsResponseWithDefaults

`func NewAddonsResponseWithDefaults() *AddonsResponse`

NewAddonsResponseWithDefaults instantiates a new AddonsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *AddonsResponse) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *AddonsResponse) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *AddonsResponse) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.

### HasResponseId

`func (o *AddonsResponse) HasResponseId() bool`

HasResponseId returns a boolean if a field has been set.

### GetMeta

`func (o *AddonsResponse) GetMeta() SchemasMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AddonsResponse) GetMetaOk() (*SchemasMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AddonsResponse) SetMeta(v SchemasMeta)`

SetMeta sets Meta field to given value.


### GetAddons

`func (o *AddonsResponse) GetAddons() []AddonOut`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *AddonsResponse) GetAddonsOk() (*[]AddonOut, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *AddonsResponse) SetAddons(v []AddonOut)`

SetAddons sets Addons field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


