# AppV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID для каждого экземпляра приложения. Автоматически генерируется при создании. | 
**Type** | **string** | Тип приложения. | 
**Name** | **string** | Удобочитаемое имя, установленное для приложения. | 
**Status** | **string** | Текущее состояние приложения. После принятия запроса на создание — new. Ответ 201 не означает завершение развёртывания. | 
**Provider** | [**AppV2Provider**](AppV2Provider.md) |  | 
**Ip** | Pointer to **string** | Публичный IPv4-адрес. Отсутствует, пока сервер и IP-адрес не созданы. | [optional] 
**Domains** | [**[]AppV2DomainsInner**](AppV2DomainsInner.md) | Домены приложения. В начальном ответе на создание массив пуст. | 
**Framework** | Pointer to **string** | Фреймворк приложения. Поле может отсутствовать до создания инфраструктуры. | [optional] 
**Location** | Pointer to **string** | Локация сервера. Может отсутствовать до создания сервера. | [optional] 
**Repository** | Pointer to [**Repository**](Repository.md) |  | [optional] 
**EnvVersion** | Pointer to **NullableString** | Версия окружения. Поле может отсутствовать до создания инфраструктуры. | [optional] 
**Envs** | Pointer to **map[string]string** | Переменные окружения приложения. Объект с ключами и значениями типа string. Поле может отсутствовать до создания инфраструктуры. | [optional] 
**IsAutoDeploy** | **bool** | Включен ли автоматический деплой. | 
**CommitSha** | Pointer to **string** | Хэш коммита, из которого собрано приложение. | [optional] 
**Comment** | Pointer to **string** | Комментарий к приложению. | [optional] 
**PresetId** | **float32** | ID тарифа приложения; 0 при использовании конфигуратора. | 
**IndexDir** | Pointer to **NullableString** | Каталог со статическими файлами frontend-приложения. До создания инфраструктуры поле отсутствует. Для backend — null. | [optional] 
**BuildCmd** | Pointer to **string** | Команда сборки приложения. Поле может отсутствовать до создания инфраструктуры. | [optional] 
**AvatarLink** | **NullableString** | Ссылка на аватар приложения. | 
**RunCmd** | Pointer to **NullableString** | Команда запуска backend-приложения. До создания инфраструктуры поле отсутствует. Для frontend — null. | [optional] 
**Configuration** | Pointer to [**NullableAppV2Configuration**](AppV2Configuration.md) |  | [optional] 
**DiskStatus** | Pointer to [**NullableAppV2DiskStatus**](AppV2DiskStatus.md) |  | [optional] 
**IsQemuAgent** | Pointer to **NullableBool** | Включён ли агент QEMU. Для frontend — null. Для backend поле отсутствует до создания сервера. | [optional] 
**Language** | Pointer to **string** | Язык или способ развёртывания. Может отсутствовать до создания инфраструктуры. Поле может отсутствовать до создания инфраструктуры. | [optional] 
**StartTime** | **time.Time** | Время последнего изменения записи приложения. | 
**Branch** | Pointer to **string** | Ветка Git-репозитория. В ответе v2 поле называется branch. | [optional] 
**ConfiguratorId** | **NullableInt32** | ID конфигуратора приложения; null при использовании готового тарифа. | 
**ProjectId** | Pointer to **int32** | ID проекта приложения. Может отсутствовать, если проект не задан. | [optional] 
**SystemDependencies** | Pointer to **[]string** | Системные пакеты приложения. Поле отсутствует, если зависимости не заданы. | [optional] 
**CreatedAt** | **time.Time** | Время создания записи приложения. | 
**AvailabilityZone** | Pointer to **NullableString** | Зона доступности сервера. Может быть null или отсутствовать до создания сервера. | [optional] 
**RequestsLimit** | **NullableInt32** | Ограничение количества запросов; null, если индивидуальный лимит не задан. | 
**BuildEntryPath** | Pointer to **string** | Каталог приложения в Git-репозитории, относительно которого выполняется сборка. Может отсутствовать до создания инфраструктуры. | [optional] 
**HealthCheck** | Pointer to **NullableString** | HTTP-путь для проверки работоспособности backend-приложения. Может отсутствовать до создания инфраструктуры. | [optional] 
**DockerImage** | Pointer to **string** | Имя готового Docker-образа без тега. Присутствует для backend из готового Docker-образа. | [optional] 
**DockerImageTag** | Pointer to **string** | Тег готового Docker-образа. Присутствует для backend из готового Docker-образа. | [optional] 
**IsRussianTrustedCa** | Pointer to **bool** | Добавлены ли доверенные корневые сертификаты Минцифры в окружение backend-приложения. Поле возвращается в российском регионе после создания инфраструктуры. | [optional] 
**IsTciTrustedCa** | Pointer to **bool** | Добавлены ли доверенные корневые сертификаты ТЦИ в окружение backend-приложения. Поле возвращается в российском регионе после создания инфраструктуры. | [optional] 
**BlockedPorts** | Pointer to **[]int32** | Заблокированные порты backend-приложения. Поле может отсутствовать до создания сервера. | [optional] 
**ReverseProxy** | **map[string]interface{}** | Настройки обратного прокси. В начальном ответе на создание — null. | 

## Methods

### NewAppV2

`func NewAppV2(id int32, type_ string, name string, status string, provider AppV2Provider, domains []AppV2DomainsInner, isAutoDeploy bool, presetId float32, avatarLink NullableString, startTime time.Time, configuratorId NullableInt32, createdAt time.Time, requestsLimit NullableInt32, reverseProxy map[string]interface{}, ) *AppV2`

NewAppV2 instantiates a new AppV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppV2WithDefaults

`func NewAppV2WithDefaults() *AppV2`

NewAppV2WithDefaults instantiates a new AppV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppV2) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppV2) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppV2) SetId(v int32)`

SetId sets Id field to given value.


### GetType

`func (o *AppV2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppV2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppV2) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *AppV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppV2) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *AppV2) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppV2) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppV2) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetProvider

`func (o *AppV2) GetProvider() AppV2Provider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AppV2) GetProviderOk() (*AppV2Provider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AppV2) SetProvider(v AppV2Provider)`

SetProvider sets Provider field to given value.


### GetIp

`func (o *AppV2) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *AppV2) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *AppV2) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *AppV2) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetDomains

`func (o *AppV2) GetDomains() []AppV2DomainsInner`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *AppV2) GetDomainsOk() (*[]AppV2DomainsInner, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *AppV2) SetDomains(v []AppV2DomainsInner)`

SetDomains sets Domains field to given value.


### GetFramework

`func (o *AppV2) GetFramework() string`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *AppV2) GetFrameworkOk() (*string, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *AppV2) SetFramework(v string)`

SetFramework sets Framework field to given value.

### HasFramework

`func (o *AppV2) HasFramework() bool`

HasFramework returns a boolean if a field has been set.

### GetLocation

`func (o *AppV2) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AppV2) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AppV2) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *AppV2) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetRepository

`func (o *AppV2) GetRepository() Repository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *AppV2) GetRepositoryOk() (*Repository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *AppV2) SetRepository(v Repository)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *AppV2) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetEnvVersion

`func (o *AppV2) GetEnvVersion() string`

GetEnvVersion returns the EnvVersion field if non-nil, zero value otherwise.

### GetEnvVersionOk

`func (o *AppV2) GetEnvVersionOk() (*string, bool)`

GetEnvVersionOk returns a tuple with the EnvVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVersion

`func (o *AppV2) SetEnvVersion(v string)`

SetEnvVersion sets EnvVersion field to given value.

### HasEnvVersion

`func (o *AppV2) HasEnvVersion() bool`

HasEnvVersion returns a boolean if a field has been set.

### SetEnvVersionNil

`func (o *AppV2) SetEnvVersionNil(b bool)`

 SetEnvVersionNil sets the value for EnvVersion to be an explicit nil

### UnsetEnvVersion
`func (o *AppV2) UnsetEnvVersion()`

UnsetEnvVersion ensures that no value is present for EnvVersion, not even an explicit nil
### GetEnvs

`func (o *AppV2) GetEnvs() map[string]string`

GetEnvs returns the Envs field if non-nil, zero value otherwise.

### GetEnvsOk

`func (o *AppV2) GetEnvsOk() (*map[string]string, bool)`

GetEnvsOk returns a tuple with the Envs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvs

`func (o *AppV2) SetEnvs(v map[string]string)`

SetEnvs sets Envs field to given value.

### HasEnvs

`func (o *AppV2) HasEnvs() bool`

HasEnvs returns a boolean if a field has been set.

### GetIsAutoDeploy

`func (o *AppV2) GetIsAutoDeploy() bool`

GetIsAutoDeploy returns the IsAutoDeploy field if non-nil, zero value otherwise.

### GetIsAutoDeployOk

`func (o *AppV2) GetIsAutoDeployOk() (*bool, bool)`

GetIsAutoDeployOk returns a tuple with the IsAutoDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoDeploy

`func (o *AppV2) SetIsAutoDeploy(v bool)`

SetIsAutoDeploy sets IsAutoDeploy field to given value.


### GetCommitSha

`func (o *AppV2) GetCommitSha() string`

GetCommitSha returns the CommitSha field if non-nil, zero value otherwise.

### GetCommitShaOk

`func (o *AppV2) GetCommitShaOk() (*string, bool)`

GetCommitShaOk returns a tuple with the CommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitSha

`func (o *AppV2) SetCommitSha(v string)`

SetCommitSha sets CommitSha field to given value.

### HasCommitSha

`func (o *AppV2) HasCommitSha() bool`

HasCommitSha returns a boolean if a field has been set.

### GetComment

`func (o *AppV2) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *AppV2) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *AppV2) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *AppV2) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetPresetId

`func (o *AppV2) GetPresetId() float32`

GetPresetId returns the PresetId field if non-nil, zero value otherwise.

### GetPresetIdOk

`func (o *AppV2) GetPresetIdOk() (*float32, bool)`

GetPresetIdOk returns a tuple with the PresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetId

`func (o *AppV2) SetPresetId(v float32)`

SetPresetId sets PresetId field to given value.


### GetIndexDir

`func (o *AppV2) GetIndexDir() string`

GetIndexDir returns the IndexDir field if non-nil, zero value otherwise.

### GetIndexDirOk

`func (o *AppV2) GetIndexDirOk() (*string, bool)`

GetIndexDirOk returns a tuple with the IndexDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexDir

`func (o *AppV2) SetIndexDir(v string)`

SetIndexDir sets IndexDir field to given value.

### HasIndexDir

`func (o *AppV2) HasIndexDir() bool`

HasIndexDir returns a boolean if a field has been set.

### SetIndexDirNil

`func (o *AppV2) SetIndexDirNil(b bool)`

 SetIndexDirNil sets the value for IndexDir to be an explicit nil

### UnsetIndexDir
`func (o *AppV2) UnsetIndexDir()`

UnsetIndexDir ensures that no value is present for IndexDir, not even an explicit nil
### GetBuildCmd

`func (o *AppV2) GetBuildCmd() string`

GetBuildCmd returns the BuildCmd field if non-nil, zero value otherwise.

### GetBuildCmdOk

`func (o *AppV2) GetBuildCmdOk() (*string, bool)`

GetBuildCmdOk returns a tuple with the BuildCmd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildCmd

`func (o *AppV2) SetBuildCmd(v string)`

SetBuildCmd sets BuildCmd field to given value.

### HasBuildCmd

`func (o *AppV2) HasBuildCmd() bool`

HasBuildCmd returns a boolean if a field has been set.

### GetAvatarLink

`func (o *AppV2) GetAvatarLink() string`

GetAvatarLink returns the AvatarLink field if non-nil, zero value otherwise.

### GetAvatarLinkOk

`func (o *AppV2) GetAvatarLinkOk() (*string, bool)`

GetAvatarLinkOk returns a tuple with the AvatarLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarLink

`func (o *AppV2) SetAvatarLink(v string)`

SetAvatarLink sets AvatarLink field to given value.


### SetAvatarLinkNil

`func (o *AppV2) SetAvatarLinkNil(b bool)`

 SetAvatarLinkNil sets the value for AvatarLink to be an explicit nil

### UnsetAvatarLink
`func (o *AppV2) UnsetAvatarLink()`

UnsetAvatarLink ensures that no value is present for AvatarLink, not even an explicit nil
### GetRunCmd

`func (o *AppV2) GetRunCmd() string`

GetRunCmd returns the RunCmd field if non-nil, zero value otherwise.

### GetRunCmdOk

`func (o *AppV2) GetRunCmdOk() (*string, bool)`

GetRunCmdOk returns a tuple with the RunCmd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunCmd

`func (o *AppV2) SetRunCmd(v string)`

SetRunCmd sets RunCmd field to given value.

### HasRunCmd

`func (o *AppV2) HasRunCmd() bool`

HasRunCmd returns a boolean if a field has been set.

### SetRunCmdNil

`func (o *AppV2) SetRunCmdNil(b bool)`

 SetRunCmdNil sets the value for RunCmd to be an explicit nil

### UnsetRunCmd
`func (o *AppV2) UnsetRunCmd()`

UnsetRunCmd ensures that no value is present for RunCmd, not even an explicit nil
### GetConfiguration

`func (o *AppV2) GetConfiguration() AppV2Configuration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *AppV2) GetConfigurationOk() (*AppV2Configuration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *AppV2) SetConfiguration(v AppV2Configuration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *AppV2) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### SetConfigurationNil

`func (o *AppV2) SetConfigurationNil(b bool)`

 SetConfigurationNil sets the value for Configuration to be an explicit nil

### UnsetConfiguration
`func (o *AppV2) UnsetConfiguration()`

UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil
### GetDiskStatus

`func (o *AppV2) GetDiskStatus() AppV2DiskStatus`

GetDiskStatus returns the DiskStatus field if non-nil, zero value otherwise.

### GetDiskStatusOk

`func (o *AppV2) GetDiskStatusOk() (*AppV2DiskStatus, bool)`

GetDiskStatusOk returns a tuple with the DiskStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskStatus

`func (o *AppV2) SetDiskStatus(v AppV2DiskStatus)`

SetDiskStatus sets DiskStatus field to given value.

### HasDiskStatus

`func (o *AppV2) HasDiskStatus() bool`

HasDiskStatus returns a boolean if a field has been set.

### SetDiskStatusNil

`func (o *AppV2) SetDiskStatusNil(b bool)`

 SetDiskStatusNil sets the value for DiskStatus to be an explicit nil

### UnsetDiskStatus
`func (o *AppV2) UnsetDiskStatus()`

UnsetDiskStatus ensures that no value is present for DiskStatus, not even an explicit nil
### GetIsQemuAgent

`func (o *AppV2) GetIsQemuAgent() bool`

GetIsQemuAgent returns the IsQemuAgent field if non-nil, zero value otherwise.

### GetIsQemuAgentOk

`func (o *AppV2) GetIsQemuAgentOk() (*bool, bool)`

GetIsQemuAgentOk returns a tuple with the IsQemuAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQemuAgent

`func (o *AppV2) SetIsQemuAgent(v bool)`

SetIsQemuAgent sets IsQemuAgent field to given value.

### HasIsQemuAgent

`func (o *AppV2) HasIsQemuAgent() bool`

HasIsQemuAgent returns a boolean if a field has been set.

### SetIsQemuAgentNil

`func (o *AppV2) SetIsQemuAgentNil(b bool)`

 SetIsQemuAgentNil sets the value for IsQemuAgent to be an explicit nil

### UnsetIsQemuAgent
`func (o *AppV2) UnsetIsQemuAgent()`

UnsetIsQemuAgent ensures that no value is present for IsQemuAgent, not even an explicit nil
### GetLanguage

`func (o *AppV2) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *AppV2) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *AppV2) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *AppV2) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetStartTime

`func (o *AppV2) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *AppV2) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *AppV2) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetBranch

`func (o *AppV2) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *AppV2) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *AppV2) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *AppV2) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetConfiguratorId

`func (o *AppV2) GetConfiguratorId() int32`

GetConfiguratorId returns the ConfiguratorId field if non-nil, zero value otherwise.

### GetConfiguratorIdOk

`func (o *AppV2) GetConfiguratorIdOk() (*int32, bool)`

GetConfiguratorIdOk returns a tuple with the ConfiguratorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguratorId

`func (o *AppV2) SetConfiguratorId(v int32)`

SetConfiguratorId sets ConfiguratorId field to given value.


### SetConfiguratorIdNil

`func (o *AppV2) SetConfiguratorIdNil(b bool)`

 SetConfiguratorIdNil sets the value for ConfiguratorId to be an explicit nil

### UnsetConfiguratorId
`func (o *AppV2) UnsetConfiguratorId()`

UnsetConfiguratorId ensures that no value is present for ConfiguratorId, not even an explicit nil
### GetProjectId

`func (o *AppV2) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AppV2) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AppV2) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *AppV2) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetSystemDependencies

`func (o *AppV2) GetSystemDependencies() []string`

GetSystemDependencies returns the SystemDependencies field if non-nil, zero value otherwise.

### GetSystemDependenciesOk

`func (o *AppV2) GetSystemDependenciesOk() (*[]string, bool)`

GetSystemDependenciesOk returns a tuple with the SystemDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDependencies

`func (o *AppV2) SetSystemDependencies(v []string)`

SetSystemDependencies sets SystemDependencies field to given value.

### HasSystemDependencies

`func (o *AppV2) HasSystemDependencies() bool`

HasSystemDependencies returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AppV2) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AppV2) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AppV2) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetAvailabilityZone

`func (o *AppV2) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *AppV2) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *AppV2) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *AppV2) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### SetAvailabilityZoneNil

`func (o *AppV2) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *AppV2) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetRequestsLimit

`func (o *AppV2) GetRequestsLimit() int32`

GetRequestsLimit returns the RequestsLimit field if non-nil, zero value otherwise.

### GetRequestsLimitOk

`func (o *AppV2) GetRequestsLimitOk() (*int32, bool)`

GetRequestsLimitOk returns a tuple with the RequestsLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestsLimit

`func (o *AppV2) SetRequestsLimit(v int32)`

SetRequestsLimit sets RequestsLimit field to given value.


### SetRequestsLimitNil

`func (o *AppV2) SetRequestsLimitNil(b bool)`

 SetRequestsLimitNil sets the value for RequestsLimit to be an explicit nil

### UnsetRequestsLimit
`func (o *AppV2) UnsetRequestsLimit()`

UnsetRequestsLimit ensures that no value is present for RequestsLimit, not even an explicit nil
### GetBuildEntryPath

`func (o *AppV2) GetBuildEntryPath() string`

GetBuildEntryPath returns the BuildEntryPath field if non-nil, zero value otherwise.

### GetBuildEntryPathOk

`func (o *AppV2) GetBuildEntryPathOk() (*string, bool)`

GetBuildEntryPathOk returns a tuple with the BuildEntryPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildEntryPath

`func (o *AppV2) SetBuildEntryPath(v string)`

SetBuildEntryPath sets BuildEntryPath field to given value.

### HasBuildEntryPath

`func (o *AppV2) HasBuildEntryPath() bool`

HasBuildEntryPath returns a boolean if a field has been set.

### GetHealthCheck

`func (o *AppV2) GetHealthCheck() string`

GetHealthCheck returns the HealthCheck field if non-nil, zero value otherwise.

### GetHealthCheckOk

`func (o *AppV2) GetHealthCheckOk() (*string, bool)`

GetHealthCheckOk returns a tuple with the HealthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheck

`func (o *AppV2) SetHealthCheck(v string)`

SetHealthCheck sets HealthCheck field to given value.

### HasHealthCheck

`func (o *AppV2) HasHealthCheck() bool`

HasHealthCheck returns a boolean if a field has been set.

### SetHealthCheckNil

`func (o *AppV2) SetHealthCheckNil(b bool)`

 SetHealthCheckNil sets the value for HealthCheck to be an explicit nil

### UnsetHealthCheck
`func (o *AppV2) UnsetHealthCheck()`

UnsetHealthCheck ensures that no value is present for HealthCheck, not even an explicit nil
### GetDockerImage

`func (o *AppV2) GetDockerImage() string`

GetDockerImage returns the DockerImage field if non-nil, zero value otherwise.

### GetDockerImageOk

`func (o *AppV2) GetDockerImageOk() (*string, bool)`

GetDockerImageOk returns a tuple with the DockerImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImage

`func (o *AppV2) SetDockerImage(v string)`

SetDockerImage sets DockerImage field to given value.

### HasDockerImage

`func (o *AppV2) HasDockerImage() bool`

HasDockerImage returns a boolean if a field has been set.

### GetDockerImageTag

`func (o *AppV2) GetDockerImageTag() string`

GetDockerImageTag returns the DockerImageTag field if non-nil, zero value otherwise.

### GetDockerImageTagOk

`func (o *AppV2) GetDockerImageTagOk() (*string, bool)`

GetDockerImageTagOk returns a tuple with the DockerImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImageTag

`func (o *AppV2) SetDockerImageTag(v string)`

SetDockerImageTag sets DockerImageTag field to given value.

### HasDockerImageTag

`func (o *AppV2) HasDockerImageTag() bool`

HasDockerImageTag returns a boolean if a field has been set.

### GetIsRussianTrustedCa

`func (o *AppV2) GetIsRussianTrustedCa() bool`

GetIsRussianTrustedCa returns the IsRussianTrustedCa field if non-nil, zero value otherwise.

### GetIsRussianTrustedCaOk

`func (o *AppV2) GetIsRussianTrustedCaOk() (*bool, bool)`

GetIsRussianTrustedCaOk returns a tuple with the IsRussianTrustedCa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRussianTrustedCa

`func (o *AppV2) SetIsRussianTrustedCa(v bool)`

SetIsRussianTrustedCa sets IsRussianTrustedCa field to given value.

### HasIsRussianTrustedCa

`func (o *AppV2) HasIsRussianTrustedCa() bool`

HasIsRussianTrustedCa returns a boolean if a field has been set.

### GetIsTciTrustedCa

`func (o *AppV2) GetIsTciTrustedCa() bool`

GetIsTciTrustedCa returns the IsTciTrustedCa field if non-nil, zero value otherwise.

### GetIsTciTrustedCaOk

`func (o *AppV2) GetIsTciTrustedCaOk() (*bool, bool)`

GetIsTciTrustedCaOk returns a tuple with the IsTciTrustedCa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTciTrustedCa

`func (o *AppV2) SetIsTciTrustedCa(v bool)`

SetIsTciTrustedCa sets IsTciTrustedCa field to given value.

### HasIsTciTrustedCa

`func (o *AppV2) HasIsTciTrustedCa() bool`

HasIsTciTrustedCa returns a boolean if a field has been set.

### GetBlockedPorts

`func (o *AppV2) GetBlockedPorts() []int32`

GetBlockedPorts returns the BlockedPorts field if non-nil, zero value otherwise.

### GetBlockedPortsOk

`func (o *AppV2) GetBlockedPortsOk() (*[]int32, bool)`

GetBlockedPortsOk returns a tuple with the BlockedPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedPorts

`func (o *AppV2) SetBlockedPorts(v []int32)`

SetBlockedPorts sets BlockedPorts field to given value.

### HasBlockedPorts

`func (o *AppV2) HasBlockedPorts() bool`

HasBlockedPorts returns a boolean if a field has been set.

### SetBlockedPortsNil

`func (o *AppV2) SetBlockedPortsNil(b bool)`

 SetBlockedPortsNil sets the value for BlockedPorts to be an explicit nil

### UnsetBlockedPorts
`func (o *AppV2) UnsetBlockedPorts()`

UnsetBlockedPorts ensures that no value is present for BlockedPorts, not even an explicit nil
### GetReverseProxy

`func (o *AppV2) GetReverseProxy() map[string]interface{}`

GetReverseProxy returns the ReverseProxy field if non-nil, zero value otherwise.

### GetReverseProxyOk

`func (o *AppV2) GetReverseProxyOk() (*map[string]interface{}, bool)`

GetReverseProxyOk returns a tuple with the ReverseProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseProxy

`func (o *AppV2) SetReverseProxy(v map[string]interface{})`

SetReverseProxy sets ReverseProxy field to given value.


### SetReverseProxyNil

`func (o *AppV2) SetReverseProxyNil(b bool)`

 SetReverseProxyNil sets the value for ReverseProxy to be an explicit nil

### UnsetReverseProxy
`func (o *AppV2) UnsetReverseProxy()`

UnsetReverseProxy ensures that no value is present for ReverseProxy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


