# DeploySettingsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Framework** | **string** | Название фреймворка. | 
**BuildCmd** | **string** | Команда для сборки проекта. | 
**IndexDir** | **string** | Директория с индексными файлами. | 
**RunCmd** | **string** | Команда для запуска проекта. | 

## Methods

### NewDeploySettingsInner

`func NewDeploySettingsInner(framework string, buildCmd string, indexDir string, runCmd string, ) *DeploySettingsInner`

NewDeploySettingsInner instantiates a new DeploySettingsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploySettingsInnerWithDefaults

`func NewDeploySettingsInnerWithDefaults() *DeploySettingsInner`

NewDeploySettingsInnerWithDefaults instantiates a new DeploySettingsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFramework

`func (o *DeploySettingsInner) GetFramework() string`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *DeploySettingsInner) GetFrameworkOk() (*string, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *DeploySettingsInner) SetFramework(v string)`

SetFramework sets Framework field to given value.


### GetBuildCmd

`func (o *DeploySettingsInner) GetBuildCmd() string`

GetBuildCmd returns the BuildCmd field if non-nil, zero value otherwise.

### GetBuildCmdOk

`func (o *DeploySettingsInner) GetBuildCmdOk() (*string, bool)`

GetBuildCmdOk returns a tuple with the BuildCmd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildCmd

`func (o *DeploySettingsInner) SetBuildCmd(v string)`

SetBuildCmd sets BuildCmd field to given value.


### GetIndexDir

`func (o *DeploySettingsInner) GetIndexDir() string`

GetIndexDir returns the IndexDir field if non-nil, zero value otherwise.

### GetIndexDirOk

`func (o *DeploySettingsInner) GetIndexDirOk() (*string, bool)`

GetIndexDirOk returns a tuple with the IndexDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexDir

`func (o *DeploySettingsInner) SetIndexDir(v string)`

SetIndexDir sets IndexDir field to given value.


### GetRunCmd

`func (o *DeploySettingsInner) GetRunCmd() string`

GetRunCmd returns the RunCmd field if non-nil, zero value otherwise.

### GetRunCmdOk

`func (o *DeploySettingsInner) GetRunCmdOk() (*string, bool)`

GetRunCmdOk returns a tuple with the RunCmd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunCmd

`func (o *DeploySettingsInner) SetRunCmd(v string)`

SetRunCmd sets RunCmd field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


