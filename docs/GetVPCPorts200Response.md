# GetVPCPorts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**Meta**](Meta.md) |  | 
**VpcPorts** | [**[]VpcPort**](VpcPort.md) |  | 
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 

## Methods

### NewGetVPCPorts200Response

`func NewGetVPCPorts200Response(meta Meta, vpcPorts []VpcPort, responseId string, ) *GetVPCPorts200Response`

NewGetVPCPorts200Response instantiates a new GetVPCPorts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVPCPorts200ResponseWithDefaults

`func NewGetVPCPorts200ResponseWithDefaults() *GetVPCPorts200Response`

NewGetVPCPorts200ResponseWithDefaults instantiates a new GetVPCPorts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetVPCPorts200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetVPCPorts200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetVPCPorts200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetVpcPorts

`func (o *GetVPCPorts200Response) GetVpcPorts() []VpcPort`

GetVpcPorts returns the VpcPorts field if non-nil, zero value otherwise.

### GetVpcPortsOk

`func (o *GetVPCPorts200Response) GetVpcPortsOk() (*[]VpcPort, bool)`

GetVpcPortsOk returns a tuple with the VpcPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcPorts

`func (o *GetVPCPorts200Response) SetVpcPorts(v []VpcPort)`

SetVpcPorts sets VpcPorts field to given value.


### GetResponseId

`func (o *GetVPCPorts200Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *GetVPCPorts200Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *GetVPCPorts200Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


