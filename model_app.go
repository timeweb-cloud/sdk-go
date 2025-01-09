/*
Timeweb Cloud API

# Введение API Timeweb Cloud позволяет вам управлять ресурсами в облаке программным способом с использованием обычных HTTP-запросов.  Множество функций, которые доступны в панели управления Timeweb Cloud, также доступны через API, что позволяет вам автоматизировать ваши собственные сценарии.  В этой документации сперва будет описан общий дизайн и принципы работы API, а после этого конкретные конечные точки. Также будут приведены примеры запросов к ним.   ## Запросы Запросы должны выполняться по протоколу `HTTPS`, чтобы гарантировать шифрование транзакций. Поддерживаются следующие методы запроса: |Метод|Применение| |--- |--- | |GET|Извлекает данные о коллекциях и отдельных ресурсах.| |POST|Для коллекций создает новый ресурс этого типа. Также используется для выполнения действий с конкретным ресурсом.| |PUT|Обновляет существующий ресурс.| |PATCH|Некоторые ресурсы поддерживают частичное обновление, то есть обновление только части атрибутов ресурса, в этом случае вместо метода PUT будет использован PATCH.| |DELETE|Удаляет ресурс.|  Методы `POST`, `PUT` и `PATCH` могут включать объект в тело запроса с типом содержимого `application/json`.  ### Параметры в запросах Некоторые коллекции поддерживают пагинацию, поиск или сортировку в запросах. В параметрах запроса требуется передать: - `limit` — обозначает количество записей, которое необходимо вернуть  - `offset` — указывает на смещение, относительно начала списка  - `search` — позволяет указать набор символов для поиска  - `sort` — можно задать правило сортировки коллекции  ## Ответы Запросы вернут один из следующих кодов состояния ответа HTTP:  |Статус|Описание| |--- |--- | |200 OK|Действие с ресурсом было выполнено успешно.| |201 Created|Ресурс был успешно создан. При этом ресурс может быть как уже готовым к использованию, так и находиться в процессе запуска.| |204 No Content|Действие с ресурсом было выполнено успешно, и ответ не содержит дополнительной информации в теле.| |400 Bad Request|Был отправлен неверный запрос, например, в нем отсутствуют обязательные параметры и т. д. Тело ответа будет содержать дополнительную информацию об ошибке.| |401 Unauthorized|Ошибка аутентификации.| |403 Forbidden|Аутентификация прошла успешно, но недостаточно прав для выполнения действия.| |404 Not Found|Запрашиваемый ресурс не найден.| |409 Conflict|Запрос конфликтует с текущим состоянием.| |423 Locked|Ресурс из запроса заблокирован от применения к нему указанного метода.| |429 Too Many Requests|Был достигнут лимит по количеству запросов в единицу времени.| |500 Internal Server Error|При выполнении запроса произошла какая-то внутренняя ошибка. Чтобы решить эту проблему, лучше всего создать тикет в панели управления.|  ### Структура успешного ответа Все конечные точки будут возвращать данные в формате `JSON`. Ответы на `GET`-запросы будут иметь на верхнем уровне следующую структуру атрибутов:  |Название поля|Тип|Описание| |--- |--- |--- | |[entity_name]|object, object[], string[], number[], boolean|Динамическое поле, которое будет меняться в зависимости от запрашиваемого ресурса и будет содержать все атрибуты, необходимые для описания этого ресурса. Например, при запросе списка баз данных будет возвращаться поле `dbs`, а при запросе конкретного облачного сервера `server`. Для некоторых конечных точек в ответе может возвращаться сразу несколько ресурсов.| |meta|object|Опционально. Объект, который содержит вспомогательную информацию о ресурсе. Чаще всего будет встречаться при запросе коллекций и содержать поле `total`, которое будет указывать на количество элементов в коллекции.| |response_id|string|Опционально. В большинстве случаев в ответе будет содержаться ID ответа в формате UUIDv4, который однозначно указывает на ваш запрос внутри нашей системы. Если вам потребуется задать вопрос нашей поддержке, приложите к вопросу этот ID— так мы сможем найти ответ на него намного быстрее. Также вы можете использовать этот ID, чтобы убедиться, что это новый ответ на запрос и результат не был получен из кэша.|  Пример запроса на получение списка SSH-ключей: ```     HTTP/2.0 200 OK     {       \"ssh_keys\":[           {             \"body\":\"ssh-rsa AAAAB3NzaC1sdfghjkOAsBwWhs= example@device.local\",             \"created_at\":\"2021-09-15T19:52:27Z\",             \"expired_at\":null,             \"id\":5297,             \"is_default\":false,             \"name\":\"example@device.local\",             \"used_at\":null,             \"used_by\":[]           }       ],       \"meta\":{           \"total\":1       },       \"response_id\":\"94608d15-8672-4eed-8ab6-28bd6fa3cdf7\"     } ```  ### Структура ответа с ошибкой |Название поля|Тип|Описание| |--- |--- |--- | |status_code|number|Короткий числовой идентификатор ошибки.| |error_code|string|Короткий текстовый идентификатор ошибки, который уточняет числовой идентификатор и удобен для программной обработки. Самый простой пример — это код `not_found` для ошибки 404.| |message|string, string[]|Опционально. В большинстве случаев в ответе будет содержаться человекочитаемое подробное описание ошибки или ошибок, которые помогут понять, что нужно исправить.| |response_id|string|Опционально. В большинстве случае в ответе будет содержаться ID ответа в формате UUIDv4, который однозначно указывает на ваш запрос внутри нашей системы. Если вам потребуется задать вопрос нашей поддержке, приложите к вопросу этот ID — так мы сможем найти ответ на него намного быстрее.|  Пример: ```     HTTP/2.0 403 Forbidden     {       \"status_code\": 403,       \"error_code\":  \"forbidden\",       \"message\":     \"You do not have access for the attempted action\",       \"response_id\": \"94608d15-8672-4eed-8ab6-28bd6fa3cdf7\"     } ```  ## Статусы ресурсов Важно учесть, что при создании большинства ресурсов внутри платформы вам будет сразу возвращен ответ от сервера со статусом `200 OK` или `201 Created` и ID созданного ресурса в теле ответа, но при этом этот ресурс может быть ещё в *состоянии запуска*.  Для того чтобы понять, в каком состоянии сейчас находится ваш ресурс, мы добавили поле `status` в ответ на получение информации о ресурсе.  Список статусов будет отличаться в зависимости от типа ресурса. Увидеть поддерживаемый список статусов вы сможете в описании каждого конкретного ресурса.     ## Ограничение скорости запросов (Rate Limiting) Чтобы обеспечить стабильность для всех пользователей, Timeweb Cloud защищает API от всплесков входящего трафика, анализируя количество запросов c каждого аккаунта к каждой конечной точке.  Если ваше приложение отправляет более 20 запросов в секунду на одну конечную точку, то для этого запроса API может вернуть код состояния HTTP `429 Too Many Requests`.   ## Аутентификация Доступ к API осуществляется с помощью JWT-токена. Токенами можно управлять внутри панели управления Timeweb Cloud в разделе *API и Terraform*.  Токен необходимо передавать в заголовке каждого запроса в формате: ```   Authorization: Bearer $TIMEWEB_CLOUD_TOKEN ```  ## Формат примеров API Примеры в этой документации описаны с помощью `curl`, HTTP-клиента командной строки. На компьютерах `Linux` и `macOS` обычно по умолчанию установлен `curl`, и он доступен для загрузки на всех популярных платформах, включая `Windows`.  Каждый пример разделен на несколько строк символом `\\`, который совместим с `bash`. Типичный пример выглядит так: ```   curl -X PATCH      -H \"Content-Type: application/json\"      -H \"Authorization: Bearer $TIMEWEB_CLOUD_TOKEN\"      -d '{\"name\":\"Cute Corvus\",\"comment\":\"Development Server\"}'      \"https://api.timeweb.cloud/api/v1/dedicated/1051\" ``` - Параметр `-X` задает метод запроса. Для согласованности метод будет указан во всех примерах, даже если он явно не требуется для методов `GET`. - Строки `-H` задают требуемые HTTP-заголовки. - Примеры, для которых требуется объект JSON в теле запроса, передают требуемые данные через параметр `-d`.  Чтобы использовать приведенные примеры, не подставляя каждый раз в них свой токен, вы можете добавить токен один раз в переменные окружения в вашей консоли. Например, на `Linux` это можно сделать с помощью команды:  ``` TIMEWEB_CLOUD_TOKEN=\"token\" ```  После этого токен будет автоматически подставляться в ваши запросы.  Обратите внимание, что все значения в этой документации являются примерами. Не полагайтесь на IDы операционных систем, тарифов и т.д., используемые в примерах. Используйте соответствующую конечную точку для получения значений перед созданием ресурсов.   ## Версионирование API построено согласно принципам [семантического версионирования](https://semver.org/lang/ru). Это значит, что мы гарантируем обратную совместимость всех изменений в пределах одной мажорной версии.  Мажорная версия каждой конечной точки обозначается в пути запроса, например, запрос `/api/v1/servers` указывает, что этот метод имеет версию 1.

API version: 1.0.0
Contact: info@timeweb.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the App type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &App{}

// App Экземпляр приложения.
type App struct {
	// ID для каждого экземпляра приложения. Автоматически генерируется при создании.
	Id float32 `json:"id"`
	// Тип приложения.
	Type string `json:"type"`
	// Удобочитаемое имя, установленное для приложения.
	Name string `json:"name"`
	// Статус приложения.
	Status string `json:"status"`
	Provider AppProvider `json:"provider"`
	// IPv4-адрес приложения.
	Ip string `json:"ip"`
	Domains []AppDomainsInner `json:"domains"`
	Framework Frameworks `json:"framework"`
	// Локация сервера.
	Location string `json:"location"`
	Repository Repository `json:"repository"`
	// Версия окружения.
	EnvVersion NullableString `json:"env_version"`
	// Переменные окружения приложения. Объект с ключами и значениями типа string.
	Envs map[string]interface{} `json:"envs"`
	// Название ветки репозитория из которой собрано приложение.
	BranchName string `json:"branch_name"`
	// Включен ли автоматический деплой.
	IsAutoDeploy bool `json:"is_auto_deploy"`
	// Хэш коммита из которого собрано приложеие.
	CommitSha string `json:"commit_sha"`
	// Комментарий к приложению.
	Comment string `json:"comment"`
	// ID тарифа.
	PresetId float32 `json:"preset_id"`
	// Путь к директории с индексным файлом. Определен для приложений `type: frontend`. Для приложений `type: backend` всегда null.
	IndexDir NullableString `json:"index_dir"`
	// Команда сборки приложения.
	BuildCmd string `json:"build_cmd"`
	// Команда для запуска приложения. Определена для приложений `type: backend`. Для приложений `type: frontend` всегда null.
	RunCmd NullableString `json:"run_cmd"`
	Configuration NullableAppConfiguration `json:"configuration"`
	DiskStatus NullableAppDiskStatus `json:"disk_status"`
	// Включен ли агент QEMU.
	IsQemuAgent bool `json:"is_qemu_agent"`
	// Язык программирования приложения.
	Language string `json:"language"`
	// Время запуска приложения.
	StartTime time.Time `json:"start_time"`
}

// NewApp instantiates a new App object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApp(id float32, type_ string, name string, status string, provider AppProvider, ip string, domains []AppDomainsInner, framework Frameworks, location string, repository Repository, envVersion NullableString, envs map[string]interface{}, branchName string, isAutoDeploy bool, commitSha string, comment string, presetId float32, indexDir NullableString, buildCmd string, runCmd NullableString, configuration NullableAppConfiguration, diskStatus NullableAppDiskStatus, isQemuAgent bool, language string, startTime time.Time) *App {
	this := App{}
	this.Id = id
	this.Type = type_
	this.Name = name
	this.Status = status
	this.Provider = provider
	this.Ip = ip
	this.Domains = domains
	this.Framework = framework
	this.Location = location
	this.Repository = repository
	this.EnvVersion = envVersion
	this.Envs = envs
	this.BranchName = branchName
	this.IsAutoDeploy = isAutoDeploy
	this.CommitSha = commitSha
	this.Comment = comment
	this.PresetId = presetId
	this.IndexDir = indexDir
	this.BuildCmd = buildCmd
	this.RunCmd = runCmd
	this.Configuration = configuration
	this.DiskStatus = diskStatus
	this.IsQemuAgent = isQemuAgent
	this.Language = language
	this.StartTime = startTime
	return &this
}

// NewAppWithDefaults instantiates a new App object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppWithDefaults() *App {
	this := App{}
	return &this
}

// GetId returns the Id field value
func (o *App) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *App) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *App) SetId(v float32) {
	o.Id = v
}

// GetType returns the Type field value
func (o *App) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *App) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *App) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *App) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *App) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *App) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value
func (o *App) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *App) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *App) SetStatus(v string) {
	o.Status = v
}

// GetProvider returns the Provider field value
func (o *App) GetProvider() AppProvider {
	if o == nil {
		var ret AppProvider
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *App) GetProviderOk() (*AppProvider, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *App) SetProvider(v AppProvider) {
	o.Provider = v
}

// GetIp returns the Ip field value
func (o *App) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *App) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *App) SetIp(v string) {
	o.Ip = v
}

// GetDomains returns the Domains field value
func (o *App) GetDomains() []AppDomainsInner {
	if o == nil {
		var ret []AppDomainsInner
		return ret
	}

	return o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value
// and a boolean to check if the value has been set.
func (o *App) GetDomainsOk() ([]AppDomainsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Domains, true
}

// SetDomains sets field value
func (o *App) SetDomains(v []AppDomainsInner) {
	o.Domains = v
}

// GetFramework returns the Framework field value
func (o *App) GetFramework() Frameworks {
	if o == nil {
		var ret Frameworks
		return ret
	}

	return o.Framework
}

// GetFrameworkOk returns a tuple with the Framework field value
// and a boolean to check if the value has been set.
func (o *App) GetFrameworkOk() (*Frameworks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Framework, true
}

// SetFramework sets field value
func (o *App) SetFramework(v Frameworks) {
	o.Framework = v
}

// GetLocation returns the Location field value
func (o *App) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *App) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *App) SetLocation(v string) {
	o.Location = v
}

// GetRepository returns the Repository field value
func (o *App) GetRepository() Repository {
	if o == nil {
		var ret Repository
		return ret
	}

	return o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value
// and a boolean to check if the value has been set.
func (o *App) GetRepositoryOk() (*Repository, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repository, true
}

// SetRepository sets field value
func (o *App) SetRepository(v Repository) {
	o.Repository = v
}

// GetEnvVersion returns the EnvVersion field value
// If the value is explicit nil, the zero value for string will be returned
func (o *App) GetEnvVersion() string {
	if o == nil || o.EnvVersion.Get() == nil {
		var ret string
		return ret
	}

	return *o.EnvVersion.Get()
}

// GetEnvVersionOk returns a tuple with the EnvVersion field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *App) GetEnvVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnvVersion.Get(), o.EnvVersion.IsSet()
}

// SetEnvVersion sets field value
func (o *App) SetEnvVersion(v string) {
	o.EnvVersion.Set(&v)
}

// GetEnvs returns the Envs field value
func (o *App) GetEnvs() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Envs
}

// GetEnvsOk returns a tuple with the Envs field value
// and a boolean to check if the value has been set.
func (o *App) GetEnvsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Envs, true
}

// SetEnvs sets field value
func (o *App) SetEnvs(v map[string]interface{}) {
	o.Envs = v
}

// GetBranchName returns the BranchName field value
func (o *App) GetBranchName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BranchName
}

// GetBranchNameOk returns a tuple with the BranchName field value
// and a boolean to check if the value has been set.
func (o *App) GetBranchNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BranchName, true
}

// SetBranchName sets field value
func (o *App) SetBranchName(v string) {
	o.BranchName = v
}

// GetIsAutoDeploy returns the IsAutoDeploy field value
func (o *App) GetIsAutoDeploy() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsAutoDeploy
}

// GetIsAutoDeployOk returns a tuple with the IsAutoDeploy field value
// and a boolean to check if the value has been set.
func (o *App) GetIsAutoDeployOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsAutoDeploy, true
}

// SetIsAutoDeploy sets field value
func (o *App) SetIsAutoDeploy(v bool) {
	o.IsAutoDeploy = v
}

// GetCommitSha returns the CommitSha field value
func (o *App) GetCommitSha() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommitSha
}

// GetCommitShaOk returns a tuple with the CommitSha field value
// and a boolean to check if the value has been set.
func (o *App) GetCommitShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommitSha, true
}

// SetCommitSha sets field value
func (o *App) SetCommitSha(v string) {
	o.CommitSha = v
}

// GetComment returns the Comment field value
func (o *App) GetComment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Comment
}

// GetCommentOk returns a tuple with the Comment field value
// and a boolean to check if the value has been set.
func (o *App) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comment, true
}

// SetComment sets field value
func (o *App) SetComment(v string) {
	o.Comment = v
}

// GetPresetId returns the PresetId field value
func (o *App) GetPresetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PresetId
}

// GetPresetIdOk returns a tuple with the PresetId field value
// and a boolean to check if the value has been set.
func (o *App) GetPresetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PresetId, true
}

// SetPresetId sets field value
func (o *App) SetPresetId(v float32) {
	o.PresetId = v
}

// GetIndexDir returns the IndexDir field value
// If the value is explicit nil, the zero value for string will be returned
func (o *App) GetIndexDir() string {
	if o == nil || o.IndexDir.Get() == nil {
		var ret string
		return ret
	}

	return *o.IndexDir.Get()
}

// GetIndexDirOk returns a tuple with the IndexDir field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *App) GetIndexDirOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IndexDir.Get(), o.IndexDir.IsSet()
}

// SetIndexDir sets field value
func (o *App) SetIndexDir(v string) {
	o.IndexDir.Set(&v)
}

// GetBuildCmd returns the BuildCmd field value
func (o *App) GetBuildCmd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BuildCmd
}

// GetBuildCmdOk returns a tuple with the BuildCmd field value
// and a boolean to check if the value has been set.
func (o *App) GetBuildCmdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BuildCmd, true
}

// SetBuildCmd sets field value
func (o *App) SetBuildCmd(v string) {
	o.BuildCmd = v
}

// GetRunCmd returns the RunCmd field value
// If the value is explicit nil, the zero value for string will be returned
func (o *App) GetRunCmd() string {
	if o == nil || o.RunCmd.Get() == nil {
		var ret string
		return ret
	}

	return *o.RunCmd.Get()
}

// GetRunCmdOk returns a tuple with the RunCmd field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *App) GetRunCmdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RunCmd.Get(), o.RunCmd.IsSet()
}

// SetRunCmd sets field value
func (o *App) SetRunCmd(v string) {
	o.RunCmd.Set(&v)
}

// GetConfiguration returns the Configuration field value
// If the value is explicit nil, the zero value for AppConfiguration will be returned
func (o *App) GetConfiguration() AppConfiguration {
	if o == nil || o.Configuration.Get() == nil {
		var ret AppConfiguration
		return ret
	}

	return *o.Configuration.Get()
}

// GetConfigurationOk returns a tuple with the Configuration field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *App) GetConfigurationOk() (*AppConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return o.Configuration.Get(), o.Configuration.IsSet()
}

// SetConfiguration sets field value
func (o *App) SetConfiguration(v AppConfiguration) {
	o.Configuration.Set(&v)
}

// GetDiskStatus returns the DiskStatus field value
// If the value is explicit nil, the zero value for AppDiskStatus will be returned
func (o *App) GetDiskStatus() AppDiskStatus {
	if o == nil || o.DiskStatus.Get() == nil {
		var ret AppDiskStatus
		return ret
	}

	return *o.DiskStatus.Get()
}

// GetDiskStatusOk returns a tuple with the DiskStatus field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *App) GetDiskStatusOk() (*AppDiskStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiskStatus.Get(), o.DiskStatus.IsSet()
}

// SetDiskStatus sets field value
func (o *App) SetDiskStatus(v AppDiskStatus) {
	o.DiskStatus.Set(&v)
}

// GetIsQemuAgent returns the IsQemuAgent field value
func (o *App) GetIsQemuAgent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsQemuAgent
}

// GetIsQemuAgentOk returns a tuple with the IsQemuAgent field value
// and a boolean to check if the value has been set.
func (o *App) GetIsQemuAgentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsQemuAgent, true
}

// SetIsQemuAgent sets field value
func (o *App) SetIsQemuAgent(v bool) {
	o.IsQemuAgent = v
}

// GetLanguage returns the Language field value
func (o *App) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *App) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *App) SetLanguage(v string) {
	o.Language = v
}

// GetStartTime returns the StartTime field value
func (o *App) GetStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *App) GetStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *App) SetStartTime(v time.Time) {
	o.StartTime = v
}

func (o App) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o App) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name
	toSerialize["status"] = o.Status
	toSerialize["provider"] = o.Provider
	toSerialize["ip"] = o.Ip
	toSerialize["domains"] = o.Domains
	toSerialize["framework"] = o.Framework
	toSerialize["location"] = o.Location
	toSerialize["repository"] = o.Repository
	toSerialize["env_version"] = o.EnvVersion.Get()
	toSerialize["envs"] = o.Envs
	toSerialize["branch_name"] = o.BranchName
	toSerialize["is_auto_deploy"] = o.IsAutoDeploy
	toSerialize["commit_sha"] = o.CommitSha
	toSerialize["comment"] = o.Comment
	toSerialize["preset_id"] = o.PresetId
	toSerialize["index_dir"] = o.IndexDir.Get()
	toSerialize["build_cmd"] = o.BuildCmd
	toSerialize["run_cmd"] = o.RunCmd.Get()
	toSerialize["configuration"] = o.Configuration.Get()
	toSerialize["disk_status"] = o.DiskStatus.Get()
	toSerialize["is_qemu_agent"] = o.IsQemuAgent
	toSerialize["language"] = o.Language
	toSerialize["start_time"] = o.StartTime
	return toSerialize, nil
}

type NullableApp struct {
	value *App
	isSet bool
}

func (v NullableApp) Get() *App {
	return v.value
}

func (v *NullableApp) Set(val *App) {
	v.value = val
	v.isSet = true
}

func (v NullableApp) IsSet() bool {
	return v.isSet
}

func (v *NullableApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApp(val *App) *NullableApp {
	return &NullableApp{value: val, isSet: true}
}

func (v NullableApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


