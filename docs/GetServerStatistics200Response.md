# GetServerStatistics200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | [**[]GetServerStatistics200ResponseAllOfCpuInner**](GetServerStatistics200ResponseAllOfCpuInner.md) |  | 
**NetworkTraffic** | [**[]GetServerStatistics200ResponseAllOfNetworkTrafficInner**](GetServerStatistics200ResponseAllOfNetworkTrafficInner.md) |  | 
**Disk** | [**[]GetServerStatistics200ResponseAllOfDiskInner**](GetServerStatistics200ResponseAllOfDiskInner.md) | Статистика основного диска | 
**Ram** | [**[]GetServerStatistics200ResponseAllOfRamInner**](GetServerStatistics200ResponseAllOfRamInner.md) |  | 
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 

## Methods

### NewGetServerStatistics200Response

`func NewGetServerStatistics200Response(cpu []GetServerStatistics200ResponseAllOfCpuInner, networkTraffic []GetServerStatistics200ResponseAllOfNetworkTrafficInner, disk []GetServerStatistics200ResponseAllOfDiskInner, ram []GetServerStatistics200ResponseAllOfRamInner, responseId string, ) *GetServerStatistics200Response`

NewGetServerStatistics200Response instantiates a new GetServerStatistics200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServerStatistics200ResponseWithDefaults

`func NewGetServerStatistics200ResponseWithDefaults() *GetServerStatistics200Response`

NewGetServerStatistics200ResponseWithDefaults instantiates a new GetServerStatistics200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *GetServerStatistics200Response) GetCpu() []GetServerStatistics200ResponseAllOfCpuInner`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *GetServerStatistics200Response) GetCpuOk() (*[]GetServerStatistics200ResponseAllOfCpuInner, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *GetServerStatistics200Response) SetCpu(v []GetServerStatistics200ResponseAllOfCpuInner)`

SetCpu sets Cpu field to given value.


### GetNetworkTraffic

`func (o *GetServerStatistics200Response) GetNetworkTraffic() []GetServerStatistics200ResponseAllOfNetworkTrafficInner`

GetNetworkTraffic returns the NetworkTraffic field if non-nil, zero value otherwise.

### GetNetworkTrafficOk

`func (o *GetServerStatistics200Response) GetNetworkTrafficOk() (*[]GetServerStatistics200ResponseAllOfNetworkTrafficInner, bool)`

GetNetworkTrafficOk returns a tuple with the NetworkTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTraffic

`func (o *GetServerStatistics200Response) SetNetworkTraffic(v []GetServerStatistics200ResponseAllOfNetworkTrafficInner)`

SetNetworkTraffic sets NetworkTraffic field to given value.


### GetDisk

`func (o *GetServerStatistics200Response) GetDisk() []GetServerStatistics200ResponseAllOfDiskInner`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *GetServerStatistics200Response) GetDiskOk() (*[]GetServerStatistics200ResponseAllOfDiskInner, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *GetServerStatistics200Response) SetDisk(v []GetServerStatistics200ResponseAllOfDiskInner)`

SetDisk sets Disk field to given value.


### GetRam

`func (o *GetServerStatistics200Response) GetRam() []GetServerStatistics200ResponseAllOfRamInner`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *GetServerStatistics200Response) GetRamOk() (*[]GetServerStatistics200ResponseAllOfRamInner, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *GetServerStatistics200Response) SetRam(v []GetServerStatistics200ResponseAllOfRamInner)`

SetRam sets Ram field to given value.


### GetResponseId

`func (o *GetServerStatistics200Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *GetServerStatistics200Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *GetServerStatistics200Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


