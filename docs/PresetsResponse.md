# PresetsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | Pointer to **string** | ID запроса | [optional] 
**Meta** | [**SchemasMeta**](SchemasMeta.md) |  | 
**K8sPresets** | [**[]WorkerPresetOutApi**](WorkerPresetOutApi.md) | Массив тарифов k8s | 

## Methods

### NewPresetsResponse

`func NewPresetsResponse(meta SchemasMeta, k8sPresets []WorkerPresetOutApi, ) *PresetsResponse`

NewPresetsResponse instantiates a new PresetsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPresetsResponseWithDefaults

`func NewPresetsResponseWithDefaults() *PresetsResponse`

NewPresetsResponseWithDefaults instantiates a new PresetsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *PresetsResponse) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *PresetsResponse) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *PresetsResponse) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.

### HasResponseId

`func (o *PresetsResponse) HasResponseId() bool`

HasResponseId returns a boolean if a field has been set.

### GetMeta

`func (o *PresetsResponse) GetMeta() SchemasMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PresetsResponse) GetMetaOk() (*SchemasMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PresetsResponse) SetMeta(v SchemasMeta)`

SetMeta sets Meta field to given value.


### GetK8sPresets

`func (o *PresetsResponse) GetK8sPresets() []WorkerPresetOutApi`

GetK8sPresets returns the K8sPresets field if non-nil, zero value otherwise.

### GetK8sPresetsOk

`func (o *PresetsResponse) GetK8sPresetsOk() (*[]WorkerPresetOutApi, bool)`

GetK8sPresetsOk returns a tuple with the K8sPresets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sPresets

`func (o *PresetsResponse) SetK8sPresets(v []WorkerPresetOutApi)`

SetK8sPresets sets K8sPresets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


