# App

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | ID для каждого экземпляра приложения. Автоматически генерируется при создании. | 
**Type** | **string** | Тип приложения. | 
**Name** | **string** | Удобочитаемое имя, установленное для приложения. | 
**Status** | **string** | Статус приложения. | 
**Provider** | [**AppProvider**](AppProvider.md) |  | 
**Ip** | **string** | IPv4-адрес приложения. | 
**Domains** | [**[]AppDomainsInner**](AppDomainsInner.md) |  | 
**Framework** | [**Frameworks**](Frameworks.md) |  | 
**Location** | **string** | Локация сервера. | 
**Repository** | [**Repository**](Repository.md) |  | 
**EnvVersion** | **NullableString** | Версия окружения. | 
**Envs** | **map[string]interface{}** | Переменные окружения приложения. Объект с ключами и значениями типа string. | 
**BranchName** | **string** | Название ветки репозитория из которой собрано приложение. | 
**IsAutoDeploy** | **bool** | Включен ли автоматический деплой. | 
**CommitSha** | **string** | Хэш коммита из которого собрано приложеие. | 
**Comment** | **string** | Комментарий к приложению. | 
**PresetId** | **float32** | ID тарифа. | 
**IndexDir** | **NullableString** | Путь к директории с индексным файлом. Определен для приложений &#x60;type: frontend&#x60;. Для приложений &#x60;type: backend&#x60; всегда null. | 
**BuildCmd** | **string** | Команда сборки приложения. | 
**RunCmd** | **NullableString** | Команда для запуска приложения. Определена для приложений &#x60;type: backend&#x60;. Для приложений &#x60;type: frontend&#x60; всегда null. | 
**Configuration** | [**NullableAppConfiguration**](AppConfiguration.md) |  | 
**DiskStatus** | [**NullableAppDiskStatus**](AppDiskStatus.md) |  | 
**IsQemuAgent** | **bool** | Включен ли агент QEMU. | 
**Language** | **string** | Язык программирования приложения. | 
**StartTime** | **time.Time** | Время запуска приложения. | 

## Methods

### NewApp

`func NewApp(id float32, type_ string, name string, status string, provider AppProvider, ip string, domains []AppDomainsInner, framework Frameworks, location string, repository Repository, envVersion NullableString, envs map[string]interface{}, branchName string, isAutoDeploy bool, commitSha string, comment string, presetId float32, indexDir NullableString, buildCmd string, runCmd NullableString, configuration NullableAppConfiguration, diskStatus NullableAppDiskStatus, isQemuAgent bool, language string, startTime time.Time, ) *App`

NewApp instantiates a new App object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppWithDefaults

`func NewAppWithDefaults() *App`

NewAppWithDefaults instantiates a new App object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *App) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *App) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *App) SetId(v float32)`

SetId sets Id field to given value.


### GetType

`func (o *App) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *App) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *App) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *App) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *App) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *App) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *App) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *App) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *App) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetProvider

`func (o *App) GetProvider() AppProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *App) GetProviderOk() (*AppProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *App) SetProvider(v AppProvider)`

SetProvider sets Provider field to given value.


### GetIp

`func (o *App) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *App) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *App) SetIp(v string)`

SetIp sets Ip field to given value.


### GetDomains

`func (o *App) GetDomains() []AppDomainsInner`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *App) GetDomainsOk() (*[]AppDomainsInner, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *App) SetDomains(v []AppDomainsInner)`

SetDomains sets Domains field to given value.


### GetFramework

`func (o *App) GetFramework() Frameworks`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *App) GetFrameworkOk() (*Frameworks, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *App) SetFramework(v Frameworks)`

SetFramework sets Framework field to given value.


### GetLocation

`func (o *App) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *App) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *App) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetRepository

`func (o *App) GetRepository() Repository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *App) GetRepositoryOk() (*Repository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *App) SetRepository(v Repository)`

SetRepository sets Repository field to given value.


### GetEnvVersion

`func (o *App) GetEnvVersion() string`

GetEnvVersion returns the EnvVersion field if non-nil, zero value otherwise.

### GetEnvVersionOk

`func (o *App) GetEnvVersionOk() (*string, bool)`

GetEnvVersionOk returns a tuple with the EnvVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVersion

`func (o *App) SetEnvVersion(v string)`

SetEnvVersion sets EnvVersion field to given value.


### SetEnvVersionNil

`func (o *App) SetEnvVersionNil(b bool)`

 SetEnvVersionNil sets the value for EnvVersion to be an explicit nil

### UnsetEnvVersion
`func (o *App) UnsetEnvVersion()`

UnsetEnvVersion ensures that no value is present for EnvVersion, not even an explicit nil
### GetEnvs

`func (o *App) GetEnvs() map[string]interface{}`

GetEnvs returns the Envs field if non-nil, zero value otherwise.

### GetEnvsOk

`func (o *App) GetEnvsOk() (*map[string]interface{}, bool)`

GetEnvsOk returns a tuple with the Envs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvs

`func (o *App) SetEnvs(v map[string]interface{})`

SetEnvs sets Envs field to given value.


### GetBranchName

`func (o *App) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *App) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *App) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.


### GetIsAutoDeploy

`func (o *App) GetIsAutoDeploy() bool`

GetIsAutoDeploy returns the IsAutoDeploy field if non-nil, zero value otherwise.

### GetIsAutoDeployOk

`func (o *App) GetIsAutoDeployOk() (*bool, bool)`

GetIsAutoDeployOk returns a tuple with the IsAutoDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoDeploy

`func (o *App) SetIsAutoDeploy(v bool)`

SetIsAutoDeploy sets IsAutoDeploy field to given value.


### GetCommitSha

`func (o *App) GetCommitSha() string`

GetCommitSha returns the CommitSha field if non-nil, zero value otherwise.

### GetCommitShaOk

`func (o *App) GetCommitShaOk() (*string, bool)`

GetCommitShaOk returns a tuple with the CommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitSha

`func (o *App) SetCommitSha(v string)`

SetCommitSha sets CommitSha field to given value.


### GetComment

`func (o *App) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *App) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *App) SetComment(v string)`

SetComment sets Comment field to given value.


### GetPresetId

`func (o *App) GetPresetId() float32`

GetPresetId returns the PresetId field if non-nil, zero value otherwise.

### GetPresetIdOk

`func (o *App) GetPresetIdOk() (*float32, bool)`

GetPresetIdOk returns a tuple with the PresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetId

`func (o *App) SetPresetId(v float32)`

SetPresetId sets PresetId field to given value.


### GetIndexDir

`func (o *App) GetIndexDir() string`

GetIndexDir returns the IndexDir field if non-nil, zero value otherwise.

### GetIndexDirOk

`func (o *App) GetIndexDirOk() (*string, bool)`

GetIndexDirOk returns a tuple with the IndexDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexDir

`func (o *App) SetIndexDir(v string)`

SetIndexDir sets IndexDir field to given value.


### SetIndexDirNil

`func (o *App) SetIndexDirNil(b bool)`

 SetIndexDirNil sets the value for IndexDir to be an explicit nil

### UnsetIndexDir
`func (o *App) UnsetIndexDir()`

UnsetIndexDir ensures that no value is present for IndexDir, not even an explicit nil
### GetBuildCmd

`func (o *App) GetBuildCmd() string`

GetBuildCmd returns the BuildCmd field if non-nil, zero value otherwise.

### GetBuildCmdOk

`func (o *App) GetBuildCmdOk() (*string, bool)`

GetBuildCmdOk returns a tuple with the BuildCmd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildCmd

`func (o *App) SetBuildCmd(v string)`

SetBuildCmd sets BuildCmd field to given value.


### GetRunCmd

`func (o *App) GetRunCmd() string`

GetRunCmd returns the RunCmd field if non-nil, zero value otherwise.

### GetRunCmdOk

`func (o *App) GetRunCmdOk() (*string, bool)`

GetRunCmdOk returns a tuple with the RunCmd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunCmd

`func (o *App) SetRunCmd(v string)`

SetRunCmd sets RunCmd field to given value.


### SetRunCmdNil

`func (o *App) SetRunCmdNil(b bool)`

 SetRunCmdNil sets the value for RunCmd to be an explicit nil

### UnsetRunCmd
`func (o *App) UnsetRunCmd()`

UnsetRunCmd ensures that no value is present for RunCmd, not even an explicit nil
### GetConfiguration

`func (o *App) GetConfiguration() AppConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *App) GetConfigurationOk() (*AppConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *App) SetConfiguration(v AppConfiguration)`

SetConfiguration sets Configuration field to given value.


### SetConfigurationNil

`func (o *App) SetConfigurationNil(b bool)`

 SetConfigurationNil sets the value for Configuration to be an explicit nil

### UnsetConfiguration
`func (o *App) UnsetConfiguration()`

UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil
### GetDiskStatus

`func (o *App) GetDiskStatus() AppDiskStatus`

GetDiskStatus returns the DiskStatus field if non-nil, zero value otherwise.

### GetDiskStatusOk

`func (o *App) GetDiskStatusOk() (*AppDiskStatus, bool)`

GetDiskStatusOk returns a tuple with the DiskStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskStatus

`func (o *App) SetDiskStatus(v AppDiskStatus)`

SetDiskStatus sets DiskStatus field to given value.


### SetDiskStatusNil

`func (o *App) SetDiskStatusNil(b bool)`

 SetDiskStatusNil sets the value for DiskStatus to be an explicit nil

### UnsetDiskStatus
`func (o *App) UnsetDiskStatus()`

UnsetDiskStatus ensures that no value is present for DiskStatus, not even an explicit nil
### GetIsQemuAgent

`func (o *App) GetIsQemuAgent() bool`

GetIsQemuAgent returns the IsQemuAgent field if non-nil, zero value otherwise.

### GetIsQemuAgentOk

`func (o *App) GetIsQemuAgentOk() (*bool, bool)`

GetIsQemuAgentOk returns a tuple with the IsQemuAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQemuAgent

`func (o *App) SetIsQemuAgent(v bool)`

SetIsQemuAgent sets IsQemuAgent field to given value.


### GetLanguage

`func (o *App) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *App) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *App) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetStartTime

`func (o *App) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *App) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *App) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


