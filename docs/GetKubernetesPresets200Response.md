# GetKubernetesPresets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 
**Meta** | [**Meta**](Meta.md) |  | 
**K8sPresets** | [**[]K8SPresetsInner**](K8SPresetsInner.md) | Массив тарифов k8s | 

## Methods

### NewGetKubernetesPresets200Response

`func NewGetKubernetesPresets200Response(responseId string, meta Meta, k8sPresets []K8SPresetsInner, ) *GetKubernetesPresets200Response`

NewGetKubernetesPresets200Response instantiates a new GetKubernetesPresets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetKubernetesPresets200ResponseWithDefaults

`func NewGetKubernetesPresets200ResponseWithDefaults() *GetKubernetesPresets200Response`

NewGetKubernetesPresets200ResponseWithDefaults instantiates a new GetKubernetesPresets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *GetKubernetesPresets200Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *GetKubernetesPresets200Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *GetKubernetesPresets200Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.


### GetMeta

`func (o *GetKubernetesPresets200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetKubernetesPresets200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetKubernetesPresets200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetK8sPresets

`func (o *GetKubernetesPresets200Response) GetK8sPresets() []K8SPresetsInner`

GetK8sPresets returns the K8sPresets field if non-nil, zero value otherwise.

### GetK8sPresetsOk

`func (o *GetKubernetesPresets200Response) GetK8sPresetsOk() (*[]K8SPresetsInner, bool)`

GetK8sPresetsOk returns a tuple with the K8sPresets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sPresets

`func (o *GetKubernetesPresets200Response) SetK8sPresets(v []K8SPresetsInner)`

SetK8sPresets sets K8sPresets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


