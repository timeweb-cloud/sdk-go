# CreateAppV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderId** | Pointer to **string** | ID провайдера. Обязателен при создании из Git-репозитория. Не требуется для готового Docker-образа. | [optional] 
**Type** | **string** | Тип приложения: frontend — статический сайт, backend — серверное приложение или Docker-образ. | 
**RepositoryId** | Pointer to **string** | ID репозитория. Обязателен при создании из Git-репозитория. Не требуется для готового Docker-образа. | [optional] 
**BuildCmd** | Pointer to **string** | Команда сборки приложения из Git-репозитория. Обязательна для Git; можно передать пустую строку, если сборка не нужна. | [optional] 
**Envs** | Pointer to **map[string]string** | Переменные окружения приложения. Ключи и значения — строки. Длина ключа — не более 255 символов, значения — не более 10240. По умолчанию пустой объект. | [optional] 
**BranchName** | Pointer to **string** | Ветка Git-репозитория, из которой необходимо собрать приложение. Обязательна при создании из Git. | [optional] 
**IsAutoDeploy** | Pointer to **bool** | Автоматически запускать деплой при изменении Git-репозитория. Поле обязательно при создании из Git. Для провайдера git и готового Docker-образа автодеплой выключен. | [optional] 
**CommitSha** | Pointer to **string** | Хэш коммита для первого деплоя. Обязателен при создании из Git. | [optional] 
**Name** | **string** | Имя приложения, от 1 до 255 символов. Допустимы латинские и русские буквы, цифры, пробелы и символы - . ( ) , &amp; # [ ] _. | 
**Comment** | **string** | Комментарий к приложению. Обязательное поле; можно передать пустую строку. Допустимы те же символы, что и в имени приложения. | 
**PresetId** | Pointer to **int32** | ID тарифа приложения. Обязателен для frontend. Для backend передайте либо preset_id, либо configuration. Нельзя передавать вместе с configuration. Доступные тарифы: GET /api/v1/presets/apps. | [optional] 
**EnvVersion** | Pointer to **NullableString** | Версия окружения для сборки и запуска, например версия Node.js. Можно передать null. | [optional] 
**Framework** | Pointer to **string** | Фреймворк приложения, строка с учётом регистра. Обязателен при создании из Git. Для frontend: &#x60;angular&#x60;, &#x60;astro&#x60;, &#x60;ember&#x60;, &#x60;next.js&#x60;, &#x60;nuxt&#x60;, &#x60;preact&#x60;, &#x60;react&#x60;, &#x60;remix&#x60;, &#x60;solid&#x60;, &#x60;solid-start&#x60;, &#x60;svelte&#x60;, &#x60;sveltekit&#x60;, &#x60;tanstack-start&#x60;, &#x60;vue&#x60;, &#x60;static-nobuild&#x60;, &#x60;another&#x60;. Для backend: &#x60;django&#x60;, &#x60;express&#x60;, &#x60;phoenix&#x60;, &#x60;Spring&#x60;, &#x60;laravel&#x60;, &#x60;beego&#x60;, &#x60;fastapi&#x60;, &#x60;ASP.NET Core&#x60;, &#x60;hapi&#x60;, &#x60;celery&#x60;, &#x60;flask&#x60;, &#x60;gin&#x60;, &#x60;docker&#x60;, &#x60;docker-compose&#x60;, &#x60;fastify&#x60;, &#x60;nest&#x60;, &#x60;symfony&#x60;, &#x60;yii&#x60;, &#x60;next.js&#x60;, &#x60;nuxt&#x60;, &#x60;astro&#x60;, &#x60;tanstack-start&#x60;, &#x60;solid-start&#x60;, &#x60;ktor&#x60;, &#x60;axum&#x60;, &#x60;actix&#x60;, &#x60;tornado&#x60;, &#x60;aiohttp&#x60;, &#x60;fiber&#x60;, &#x60;echo&#x60;, &#x60;Micronaut&#x60;, &#x60;Quarkus&#x60;, &#x60;remix&#x60;, &#x60;sveltekit&#x60;. Для готового Docker-образа не требуется. | [optional] 
**IndexDir** | Pointer to **string** | Путь к каталогу с собранными статическими файлами, например /dist. Обязателен для frontend. Не используется для backend. | [optional] 
**RunCmd** | Pointer to **string** | Команда запуска backend-приложения из Git. Обязательна для backend из Git. Не требуется для готового Docker-образа и frontend. | [optional] 
**SystemDependencies** | Pointer to **[]string** | Системные пакеты для установки в окружение приложения. Допустимы строчные латинские буквы, цифры и символы . + - &#x3D;. По умолчанию пустой массив. | [optional] 
**ProjectId** | Pointer to **int32** | ID проекта, в который необходимо добавить приложение. Проект должен принадлежать аккаунту. | [optional] 
**Configuration** | Pointer to [**CreateAppV2Configuration**](CreateAppV2Configuration.md) |  | [optional] 
**Language** | **string** | Язык или способ развёртывания. Обязателен в v2. Для frontend: &#x60;javascript&#x60;. Для backend: &#x60;node&#x60;, &#x60;elixir&#x60;, &#x60;go&#x60;, &#x60;php&#x60;, &#x60;python&#x60;, &#x60;ruby&#x60;, &#x60;docker&#x60;, &#x60;docker-compose&#x60;, &#x60;dotnet&#x60;, &#x60;java&#x60;, &#x60;kotlin&#x60;, &#x60;rust&#x60;, &#x60;bun&#x60;, &#x60;deno&#x60;, &#x60;custom-image&#x60;. Для SSR-приложений из Git используйте node, для готового Docker-образа — custom-image. | 
**BuildEntryPath** | Pointer to **string** | Каталог приложения в Git-репозитории, относительно которого выполняется сборка. Непустой путь начинается с /, не содержит .. или двойной начальный /. Допустимы латинские буквы, цифры, . - _ /. Можно передать пустую строку. | [optional] 
**Network** | Pointer to [**CreateAppV2Network**](CreateAppV2Network.md) |  | [optional] 
**HealthCheck** | Pointer to **string** | HTTP-путь для проверки работоспособности backend-приложения. | [optional] 
**GlobalEnvIds** | Pointer to **[]float32** | ID глобальных переменных окружения аккаунта, которые необходимо связать с приложением. | [optional] 
**DockerImage** | Pointer to **string** | Имя готового Docker-образа для backend, без тега. Передавайте вместе с docker_image_tag. При указании обоих полей Git-параметры не требуются. | [optional] 
**DockerImageTag** | Pointer to **string** | Тег готового Docker-образа. Передавайте вместе с docker_image. | [optional] 
**IsRussianTrustedCa** | Pointer to **bool** | Добавить доверенные корневые сертификаты Минцифры в окружение backend-приложения. Применяется в российском регионе. | [optional] [default to false]
**IsTciTrustedCa** | Pointer to **bool** | Добавить доверенные корневые сертификаты ТЦИ в окружение backend-приложения. Применяется в российском регионе. | [optional] [default to false]

## Methods

### NewCreateAppV2

`func NewCreateAppV2(type_ string, name string, comment string, language string, ) *CreateAppV2`

NewCreateAppV2 instantiates a new CreateAppV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAppV2WithDefaults

`func NewCreateAppV2WithDefaults() *CreateAppV2`

NewCreateAppV2WithDefaults instantiates a new CreateAppV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderId

`func (o *CreateAppV2) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *CreateAppV2) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *CreateAppV2) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *CreateAppV2) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetType

`func (o *CreateAppV2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateAppV2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateAppV2) SetType(v string)`

SetType sets Type field to given value.


### GetRepositoryId

`func (o *CreateAppV2) GetRepositoryId() string`

GetRepositoryId returns the RepositoryId field if non-nil, zero value otherwise.

### GetRepositoryIdOk

`func (o *CreateAppV2) GetRepositoryIdOk() (*string, bool)`

GetRepositoryIdOk returns a tuple with the RepositoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryId

`func (o *CreateAppV2) SetRepositoryId(v string)`

SetRepositoryId sets RepositoryId field to given value.

### HasRepositoryId

`func (o *CreateAppV2) HasRepositoryId() bool`

HasRepositoryId returns a boolean if a field has been set.

### GetBuildCmd

`func (o *CreateAppV2) GetBuildCmd() string`

GetBuildCmd returns the BuildCmd field if non-nil, zero value otherwise.

### GetBuildCmdOk

`func (o *CreateAppV2) GetBuildCmdOk() (*string, bool)`

GetBuildCmdOk returns a tuple with the BuildCmd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildCmd

`func (o *CreateAppV2) SetBuildCmd(v string)`

SetBuildCmd sets BuildCmd field to given value.

### HasBuildCmd

`func (o *CreateAppV2) HasBuildCmd() bool`

HasBuildCmd returns a boolean if a field has been set.

### GetEnvs

`func (o *CreateAppV2) GetEnvs() map[string]string`

GetEnvs returns the Envs field if non-nil, zero value otherwise.

### GetEnvsOk

`func (o *CreateAppV2) GetEnvsOk() (*map[string]string, bool)`

GetEnvsOk returns a tuple with the Envs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvs

`func (o *CreateAppV2) SetEnvs(v map[string]string)`

SetEnvs sets Envs field to given value.

### HasEnvs

`func (o *CreateAppV2) HasEnvs() bool`

HasEnvs returns a boolean if a field has been set.

### GetBranchName

`func (o *CreateAppV2) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *CreateAppV2) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *CreateAppV2) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.

### HasBranchName

`func (o *CreateAppV2) HasBranchName() bool`

HasBranchName returns a boolean if a field has been set.

### GetIsAutoDeploy

`func (o *CreateAppV2) GetIsAutoDeploy() bool`

GetIsAutoDeploy returns the IsAutoDeploy field if non-nil, zero value otherwise.

### GetIsAutoDeployOk

`func (o *CreateAppV2) GetIsAutoDeployOk() (*bool, bool)`

GetIsAutoDeployOk returns a tuple with the IsAutoDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoDeploy

`func (o *CreateAppV2) SetIsAutoDeploy(v bool)`

SetIsAutoDeploy sets IsAutoDeploy field to given value.

### HasIsAutoDeploy

`func (o *CreateAppV2) HasIsAutoDeploy() bool`

HasIsAutoDeploy returns a boolean if a field has been set.

### GetCommitSha

`func (o *CreateAppV2) GetCommitSha() string`

GetCommitSha returns the CommitSha field if non-nil, zero value otherwise.

### GetCommitShaOk

`func (o *CreateAppV2) GetCommitShaOk() (*string, bool)`

GetCommitShaOk returns a tuple with the CommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitSha

`func (o *CreateAppV2) SetCommitSha(v string)`

SetCommitSha sets CommitSha field to given value.

### HasCommitSha

`func (o *CreateAppV2) HasCommitSha() bool`

HasCommitSha returns a boolean if a field has been set.

### GetName

`func (o *CreateAppV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAppV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAppV2) SetName(v string)`

SetName sets Name field to given value.


### GetComment

`func (o *CreateAppV2) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateAppV2) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateAppV2) SetComment(v string)`

SetComment sets Comment field to given value.


### GetPresetId

`func (o *CreateAppV2) GetPresetId() int32`

GetPresetId returns the PresetId field if non-nil, zero value otherwise.

### GetPresetIdOk

`func (o *CreateAppV2) GetPresetIdOk() (*int32, bool)`

GetPresetIdOk returns a tuple with the PresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetId

`func (o *CreateAppV2) SetPresetId(v int32)`

SetPresetId sets PresetId field to given value.

### HasPresetId

`func (o *CreateAppV2) HasPresetId() bool`

HasPresetId returns a boolean if a field has been set.

### GetEnvVersion

`func (o *CreateAppV2) GetEnvVersion() string`

GetEnvVersion returns the EnvVersion field if non-nil, zero value otherwise.

### GetEnvVersionOk

`func (o *CreateAppV2) GetEnvVersionOk() (*string, bool)`

GetEnvVersionOk returns a tuple with the EnvVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVersion

`func (o *CreateAppV2) SetEnvVersion(v string)`

SetEnvVersion sets EnvVersion field to given value.

### HasEnvVersion

`func (o *CreateAppV2) HasEnvVersion() bool`

HasEnvVersion returns a boolean if a field has been set.

### SetEnvVersionNil

`func (o *CreateAppV2) SetEnvVersionNil(b bool)`

 SetEnvVersionNil sets the value for EnvVersion to be an explicit nil

### UnsetEnvVersion
`func (o *CreateAppV2) UnsetEnvVersion()`

UnsetEnvVersion ensures that no value is present for EnvVersion, not even an explicit nil
### GetFramework

`func (o *CreateAppV2) GetFramework() string`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *CreateAppV2) GetFrameworkOk() (*string, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *CreateAppV2) SetFramework(v string)`

SetFramework sets Framework field to given value.

### HasFramework

`func (o *CreateAppV2) HasFramework() bool`

HasFramework returns a boolean if a field has been set.

### GetIndexDir

`func (o *CreateAppV2) GetIndexDir() string`

GetIndexDir returns the IndexDir field if non-nil, zero value otherwise.

### GetIndexDirOk

`func (o *CreateAppV2) GetIndexDirOk() (*string, bool)`

GetIndexDirOk returns a tuple with the IndexDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexDir

`func (o *CreateAppV2) SetIndexDir(v string)`

SetIndexDir sets IndexDir field to given value.

### HasIndexDir

`func (o *CreateAppV2) HasIndexDir() bool`

HasIndexDir returns a boolean if a field has been set.

### GetRunCmd

`func (o *CreateAppV2) GetRunCmd() string`

GetRunCmd returns the RunCmd field if non-nil, zero value otherwise.

### GetRunCmdOk

`func (o *CreateAppV2) GetRunCmdOk() (*string, bool)`

GetRunCmdOk returns a tuple with the RunCmd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunCmd

`func (o *CreateAppV2) SetRunCmd(v string)`

SetRunCmd sets RunCmd field to given value.

### HasRunCmd

`func (o *CreateAppV2) HasRunCmd() bool`

HasRunCmd returns a boolean if a field has been set.

### GetSystemDependencies

`func (o *CreateAppV2) GetSystemDependencies() []string`

GetSystemDependencies returns the SystemDependencies field if non-nil, zero value otherwise.

### GetSystemDependenciesOk

`func (o *CreateAppV2) GetSystemDependenciesOk() (*[]string, bool)`

GetSystemDependenciesOk returns a tuple with the SystemDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDependencies

`func (o *CreateAppV2) SetSystemDependencies(v []string)`

SetSystemDependencies sets SystemDependencies field to given value.

### HasSystemDependencies

`func (o *CreateAppV2) HasSystemDependencies() bool`

HasSystemDependencies returns a boolean if a field has been set.

### GetProjectId

`func (o *CreateAppV2) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateAppV2) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateAppV2) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *CreateAppV2) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetConfiguration

`func (o *CreateAppV2) GetConfiguration() CreateAppV2Configuration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *CreateAppV2) GetConfigurationOk() (*CreateAppV2Configuration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *CreateAppV2) SetConfiguration(v CreateAppV2Configuration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *CreateAppV2) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetLanguage

`func (o *CreateAppV2) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *CreateAppV2) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *CreateAppV2) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetBuildEntryPath

`func (o *CreateAppV2) GetBuildEntryPath() string`

GetBuildEntryPath returns the BuildEntryPath field if non-nil, zero value otherwise.

### GetBuildEntryPathOk

`func (o *CreateAppV2) GetBuildEntryPathOk() (*string, bool)`

GetBuildEntryPathOk returns a tuple with the BuildEntryPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildEntryPath

`func (o *CreateAppV2) SetBuildEntryPath(v string)`

SetBuildEntryPath sets BuildEntryPath field to given value.

### HasBuildEntryPath

`func (o *CreateAppV2) HasBuildEntryPath() bool`

HasBuildEntryPath returns a boolean if a field has been set.

### GetNetwork

`func (o *CreateAppV2) GetNetwork() CreateAppV2Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *CreateAppV2) GetNetworkOk() (*CreateAppV2Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *CreateAppV2) SetNetwork(v CreateAppV2Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *CreateAppV2) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetHealthCheck

`func (o *CreateAppV2) GetHealthCheck() string`

GetHealthCheck returns the HealthCheck field if non-nil, zero value otherwise.

### GetHealthCheckOk

`func (o *CreateAppV2) GetHealthCheckOk() (*string, bool)`

GetHealthCheckOk returns a tuple with the HealthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheck

`func (o *CreateAppV2) SetHealthCheck(v string)`

SetHealthCheck sets HealthCheck field to given value.

### HasHealthCheck

`func (o *CreateAppV2) HasHealthCheck() bool`

HasHealthCheck returns a boolean if a field has been set.

### GetGlobalEnvIds

`func (o *CreateAppV2) GetGlobalEnvIds() []float32`

GetGlobalEnvIds returns the GlobalEnvIds field if non-nil, zero value otherwise.

### GetGlobalEnvIdsOk

`func (o *CreateAppV2) GetGlobalEnvIdsOk() (*[]float32, bool)`

GetGlobalEnvIdsOk returns a tuple with the GlobalEnvIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalEnvIds

`func (o *CreateAppV2) SetGlobalEnvIds(v []float32)`

SetGlobalEnvIds sets GlobalEnvIds field to given value.

### HasGlobalEnvIds

`func (o *CreateAppV2) HasGlobalEnvIds() bool`

HasGlobalEnvIds returns a boolean if a field has been set.

### GetDockerImage

`func (o *CreateAppV2) GetDockerImage() string`

GetDockerImage returns the DockerImage field if non-nil, zero value otherwise.

### GetDockerImageOk

`func (o *CreateAppV2) GetDockerImageOk() (*string, bool)`

GetDockerImageOk returns a tuple with the DockerImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImage

`func (o *CreateAppV2) SetDockerImage(v string)`

SetDockerImage sets DockerImage field to given value.

### HasDockerImage

`func (o *CreateAppV2) HasDockerImage() bool`

HasDockerImage returns a boolean if a field has been set.

### GetDockerImageTag

`func (o *CreateAppV2) GetDockerImageTag() string`

GetDockerImageTag returns the DockerImageTag field if non-nil, zero value otherwise.

### GetDockerImageTagOk

`func (o *CreateAppV2) GetDockerImageTagOk() (*string, bool)`

GetDockerImageTagOk returns a tuple with the DockerImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImageTag

`func (o *CreateAppV2) SetDockerImageTag(v string)`

SetDockerImageTag sets DockerImageTag field to given value.

### HasDockerImageTag

`func (o *CreateAppV2) HasDockerImageTag() bool`

HasDockerImageTag returns a boolean if a field has been set.

### GetIsRussianTrustedCa

`func (o *CreateAppV2) GetIsRussianTrustedCa() bool`

GetIsRussianTrustedCa returns the IsRussianTrustedCa field if non-nil, zero value otherwise.

### GetIsRussianTrustedCaOk

`func (o *CreateAppV2) GetIsRussianTrustedCaOk() (*bool, bool)`

GetIsRussianTrustedCaOk returns a tuple with the IsRussianTrustedCa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRussianTrustedCa

`func (o *CreateAppV2) SetIsRussianTrustedCa(v bool)`

SetIsRussianTrustedCa sets IsRussianTrustedCa field to given value.

### HasIsRussianTrustedCa

`func (o *CreateAppV2) HasIsRussianTrustedCa() bool`

HasIsRussianTrustedCa returns a boolean if a field has been set.

### GetIsTciTrustedCa

`func (o *CreateAppV2) GetIsTciTrustedCa() bool`

GetIsTciTrustedCa returns the IsTciTrustedCa field if non-nil, zero value otherwise.

### GetIsTciTrustedCaOk

`func (o *CreateAppV2) GetIsTciTrustedCaOk() (*bool, bool)`

GetIsTciTrustedCaOk returns a tuple with the IsTciTrustedCa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTciTrustedCa

`func (o *CreateAppV2) SetIsTciTrustedCa(v bool)`

SetIsTciTrustedCa sets IsTciTrustedCa field to given value.

### HasIsTciTrustedCa

`func (o *CreateAppV2) HasIsTciTrustedCa() bool`

HasIsTciTrustedCa returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


