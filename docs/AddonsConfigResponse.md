# AddonsConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | Pointer to **string** | ID запроса | [optional] 
**Meta** | [**SchemasMeta**](SchemasMeta.md) |  | 
**K8sAddons** | [**[]AddonConfigOut**](AddonConfigOut.md) | Массив конфигураций дополнений k8s | 

## Methods

### NewAddonsConfigResponse

`func NewAddonsConfigResponse(meta SchemasMeta, k8sAddons []AddonConfigOut, ) *AddonsConfigResponse`

NewAddonsConfigResponse instantiates a new AddonsConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonsConfigResponseWithDefaults

`func NewAddonsConfigResponseWithDefaults() *AddonsConfigResponse`

NewAddonsConfigResponseWithDefaults instantiates a new AddonsConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *AddonsConfigResponse) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *AddonsConfigResponse) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *AddonsConfigResponse) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.

### HasResponseId

`func (o *AddonsConfigResponse) HasResponseId() bool`

HasResponseId returns a boolean if a field has been set.

### GetMeta

`func (o *AddonsConfigResponse) GetMeta() SchemasMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AddonsConfigResponse) GetMetaOk() (*SchemasMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AddonsConfigResponse) SetMeta(v SchemasMeta)`

SetMeta sets Meta field to given value.


### GetK8sAddons

`func (o *AddonsConfigResponse) GetK8sAddons() []AddonConfigOut`

GetK8sAddons returns the K8sAddons field if non-nil, zero value otherwise.

### GetK8sAddonsOk

`func (o *AddonsConfigResponse) GetK8sAddonsOk() (*[]AddonConfigOut, bool)`

GetK8sAddonsOk returns a tuple with the K8sAddons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sAddons

`func (o *AddonsConfigResponse) SetK8sAddons(v []AddonConfigOut)`

SetK8sAddons sets K8sAddons field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


