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

// checks if the AppV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppV2{}

// AppV2 Приложение v2. Создание и первый деплой выполняются асинхронно. Поля сервера и инфраструктуры могут отсутствовать, пока соответствующие ресурсы не созданы.
type AppV2 struct {
	// ID для каждого экземпляра приложения. Автоматически генерируется при создании.
	Id int32 `json:"id"`
	// Тип приложения.
	Type string `json:"type"`
	// Удобочитаемое имя, установленное для приложения.
	Name string `json:"name"`
	// Текущее состояние приложения. После принятия запроса на создание — new. Ответ 201 не означает завершение развёртывания.
	Status string `json:"status"`
	Provider AppV2Provider `json:"provider"`
	// Публичный IPv4-адрес. Отсутствует, пока сервер и IP-адрес не созданы.
	Ip *string `json:"ip,omitempty"`
	// Домены приложения. В начальном ответе на создание массив пуст.
	Domains []AppV2DomainsInner `json:"domains"`
	// Фреймворк приложения. Поле может отсутствовать до создания инфраструктуры.
	Framework *string `json:"framework,omitempty"`
	// Локация сервера. Может отсутствовать до создания сервера.
	Location *string `json:"location,omitempty"`
	Repository *Repository `json:"repository,omitempty"`
	// Версия окружения. Поле может отсутствовать до создания инфраструктуры.
	EnvVersion NullableString `json:"env_version,omitempty"`
	// Переменные окружения приложения. Объект с ключами и значениями типа string. Поле может отсутствовать до создания инфраструктуры.
	Envs *map[string]string `json:"envs,omitempty"`
	// Включен ли автоматический деплой.
	IsAutoDeploy bool `json:"is_auto_deploy"`
	// Хэш коммита, из которого собрано приложение.
	CommitSha *string `json:"commit_sha,omitempty"`
	// Комментарий к приложению.
	Comment *string `json:"comment,omitempty"`
	// ID тарифа приложения; 0 при использовании конфигуратора.
	PresetId float32 `json:"preset_id"`
	// Каталог со статическими файлами frontend-приложения. До создания инфраструктуры поле отсутствует. Для backend — null.
	IndexDir NullableString `json:"index_dir,omitempty"`
	// Команда сборки приложения. Поле может отсутствовать до создания инфраструктуры.
	BuildCmd *string `json:"build_cmd,omitempty"`
	// Ссылка на аватар приложения.
	AvatarLink NullableString `json:"avatar_link"`
	// Команда запуска backend-приложения. До создания инфраструктуры поле отсутствует. Для frontend — null.
	RunCmd NullableString `json:"run_cmd,omitempty"`
	Configuration NullableAppV2Configuration `json:"configuration,omitempty"`
	DiskStatus NullableAppV2DiskStatus `json:"disk_status,omitempty"`
	// Включён ли агент QEMU. Для frontend — null. Для backend поле отсутствует до создания сервера.
	IsQemuAgent NullableBool `json:"is_qemu_agent,omitempty"`
	// Язык или способ развёртывания. Может отсутствовать до создания инфраструктуры. Поле может отсутствовать до создания инфраструктуры.
	Language *string `json:"language,omitempty"`
	// Время последнего изменения записи приложения.
	StartTime time.Time `json:"start_time"`
	// Ветка Git-репозитория. В ответе v2 поле называется branch.
	Branch *string `json:"branch,omitempty"`
	// ID конфигуратора приложения; null при использовании готового тарифа.
	ConfiguratorId NullableInt32 `json:"configurator_id"`
	// ID проекта приложения. Может отсутствовать, если проект не задан.
	ProjectId *int32 `json:"project_id,omitempty"`
	// Системные пакеты приложения. Поле отсутствует, если зависимости не заданы.
	SystemDependencies []string `json:"system_dependencies,omitempty"`
	// Время создания записи приложения.
	CreatedAt time.Time `json:"created_at"`
	// Зона доступности сервера. Может быть null или отсутствовать до создания сервера.
	AvailabilityZone NullableString `json:"availability_zone,omitempty"`
	// Ограничение количества запросов; null, если индивидуальный лимит не задан.
	RequestsLimit NullableInt32 `json:"requests_limit"`
	// Каталог приложения в Git-репозитории, относительно которого выполняется сборка. Может отсутствовать до создания инфраструктуры.
	BuildEntryPath *string `json:"build_entry_path,omitempty"`
	// HTTP-путь для проверки работоспособности backend-приложения. Может отсутствовать до создания инфраструктуры.
	HealthCheck NullableString `json:"health_check,omitempty"`
	// Имя готового Docker-образа без тега. Присутствует для backend из готового Docker-образа.
	DockerImage *string `json:"docker_image,omitempty"`
	// Тег готового Docker-образа. Присутствует для backend из готового Docker-образа.
	DockerImageTag *string `json:"docker_image_tag,omitempty"`
	// Добавлены ли доверенные корневые сертификаты Минцифры в окружение backend-приложения. Поле возвращается в российском регионе после создания инфраструктуры.
	IsRussianTrustedCa *bool `json:"is_russian_trusted_ca,omitempty"`
	// Добавлены ли доверенные корневые сертификаты ТЦИ в окружение backend-приложения. Поле возвращается в российском регионе после создания инфраструктуры.
	IsTciTrustedCa *bool `json:"is_tci_trusted_ca,omitempty"`
	// Заблокированные порты backend-приложения. Поле может отсутствовать до создания сервера.
	BlockedPorts []int32 `json:"blocked_ports,omitempty"`
	// Настройки обратного прокси. В начальном ответе на создание — null.
	ReverseProxy map[string]interface{} `json:"reverse_proxy"`
}

// NewAppV2 instantiates a new AppV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppV2(id int32, type_ string, name string, status string, provider AppV2Provider, domains []AppV2DomainsInner, isAutoDeploy bool, presetId float32, avatarLink NullableString, startTime time.Time, configuratorId NullableInt32, createdAt time.Time, requestsLimit NullableInt32, reverseProxy map[string]interface{}) *AppV2 {
	this := AppV2{}
	this.Id = id
	this.Type = type_
	this.Name = name
	this.Status = status
	this.Provider = provider
	this.Domains = domains
	this.IsAutoDeploy = isAutoDeploy
	this.PresetId = presetId
	this.AvatarLink = avatarLink
	this.StartTime = startTime
	this.ConfiguratorId = configuratorId
	this.CreatedAt = createdAt
	this.RequestsLimit = requestsLimit
	this.ReverseProxy = reverseProxy
	return &this
}

// NewAppV2WithDefaults instantiates a new AppV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppV2WithDefaults() *AppV2 {
	this := AppV2{}
	return &this
}

// GetId returns the Id field value
func (o *AppV2) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppV2) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppV2) SetId(v int32) {
	o.Id = v
}

// GetType returns the Type field value
func (o *AppV2) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppV2) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppV2) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *AppV2) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AppV2) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AppV2) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value
func (o *AppV2) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AppV2) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AppV2) SetStatus(v string) {
	o.Status = v
}

// GetProvider returns the Provider field value
func (o *AppV2) GetProvider() AppV2Provider {
	if o == nil {
		var ret AppV2Provider
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *AppV2) GetProviderOk() (*AppV2Provider, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *AppV2) SetProvider(v AppV2Provider) {
	o.Provider = v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *AppV2) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppV2) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *AppV2) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *AppV2) SetIp(v string) {
	o.Ip = &v
}

// GetDomains returns the Domains field value
func (o *AppV2) GetDomains() []AppV2DomainsInner {
	if o == nil {
		var ret []AppV2DomainsInner
		return ret
	}

	return o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value
// and a boolean to check if the value has been set.
func (o *AppV2) GetDomainsOk() ([]AppV2DomainsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Domains, true
}

// SetDomains sets field value
func (o *AppV2) SetDomains(v []AppV2DomainsInner) {
	o.Domains = v
}

// GetFramework returns the Framework field value if set, zero value otherwise.
func (o *AppV2) GetFramework() string {
	if o == nil || IsNil(o.Framework) {
		var ret string
		return ret
	}
	return *o.Framework
}

// GetFrameworkOk returns a tuple with the Framework field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppV2) GetFrameworkOk() (*string, bool) {
	if o == nil || IsNil(o.Framework) {
		return nil, false
	}
	return o.Framework, true
}

// HasFramework returns a boolean if a field has been set.
func (o *AppV2) HasFramework() bool {
	if o != nil && !IsNil(o.Framework) {
		return true
	}

	return false
}

// SetFramework gets a reference to the given string and assigns it to the Framework field.
func (o *AppV2) SetFramework(v string) {
	o.Framework = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *AppV2) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppV2) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *AppV2) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *AppV2) SetLocation(v string) {
	o.Location = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *AppV2) GetRepository() Repository {
	if o == nil || IsNil(o.Repository) {
		var ret Repository
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppV2) GetRepositoryOk() (*Repository, bool) {
	if o == nil || IsNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *AppV2) HasRepository() bool {
	if o != nil && !IsNil(o.Repository) {
		return true
	}

	return false
}

// SetRepository gets a reference to the given Repository and assigns it to the Repository field.
func (o *AppV2) SetRepository(v Repository) {
	o.Repository = &v
}

// GetEnvVersion returns the EnvVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppV2) GetEnvVersion() string {
	if o == nil || IsNil(o.EnvVersion.Get()) {
		var ret string
		return ret
	}
	return *o.EnvVersion.Get()
}

// GetEnvVersionOk returns a tuple with the EnvVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppV2) GetEnvVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnvVersion.Get(), o.EnvVersion.IsSet()
}

// HasEnvVersion returns a boolean if a field has been set.
func (o *AppV2) HasEnvVersion() bool {
	if o != nil && o.EnvVersion.IsSet() {
		return true
	}

	return false
}

// SetEnvVersion gets a reference to the given NullableString and assigns it to the EnvVersion field.
func (o *AppV2) SetEnvVersion(v string) {
	o.EnvVersion.Set(&v)
}
// SetEnvVersionNil sets the value for EnvVersion to be an explicit nil
func (o *AppV2) SetEnvVersionNil() {
	o.EnvVersion.Set(nil)
}

// UnsetEnvVersion ensures that no value is present for EnvVersion, not even an explicit nil
func (o *AppV2) UnsetEnvVersion() {
	o.EnvVersion.Unset()
}

// GetEnvs returns the Envs field value if set, zero value otherwise.
func (o *AppV2) GetEnvs() map[string]string {
	if o == nil || IsNil(o.Envs) {
		var ret map[string]string
		return ret
	}
	return *o.Envs
}

// GetEnvsOk returns a tuple with the Envs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppV2) GetEnvsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Envs) {
		return nil, false
	}
	return o.Envs, true
}

// HasEnvs returns a boolean if a field has been set.
func (o *AppV2) HasEnvs() bool {
	if o != nil && !IsNil(o.Envs) {
		return true
	}

	return false
}

// SetEnvs gets a reference to the given map[string]string and assigns it to the Envs field.
func (o *AppV2) SetEnvs(v map[string]string) {
	o.Envs = &v
}

// GetIsAutoDeploy returns the IsAutoDeploy field value
func (o *AppV2) GetIsAutoDeploy() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsAutoDeploy
}

// GetIsAutoDeployOk returns a tuple with the IsAutoDeploy field value
// and a boolean to check if the value has been set.
func (o *AppV2) GetIsAutoDeployOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsAutoDeploy, true
}

// SetIsAutoDeploy sets field value
func (o *AppV2) SetIsAutoDeploy(v bool) {
	o.IsAutoDeploy = v
}

// GetCommitSha returns the CommitSha field value if set, zero value otherwise.
func (o *AppV2) GetCommitSha() string {
	if o == nil || IsNil(o.CommitSha) {
		var ret string
		return ret
	}
	return *o.CommitSha
}

// GetCommitShaOk returns a tuple with the CommitSha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppV2) GetCommitShaOk() (*string, bool) {
	if o == nil || IsNil(o.CommitSha) {
		return nil, false
	}
	return o.CommitSha, true
}

// HasCommitSha returns a boolean if a field has been set.
func (o *AppV2) HasCommitSha() bool {
	if o != nil && !IsNil(o.CommitSha) {
		return true
	}

	return false
}

// SetCommitSha gets a reference to the given string and assigns it to the CommitSha field.
func (o *AppV2) SetCommitSha(v string) {
	o.CommitSha = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *AppV2) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppV2) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *AppV2) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *AppV2) SetComment(v string) {
	o.Comment = &v
}

// GetPresetId returns the PresetId field value
func (o *AppV2) GetPresetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PresetId
}

// GetPresetIdOk returns a tuple with the PresetId field value
// and a boolean to check if the value has been set.
func (o *AppV2) GetPresetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PresetId, true
}

// SetPresetId sets field value
func (o *AppV2) SetPresetId(v float32) {
	o.PresetId = v
}

// GetIndexDir returns the IndexDir field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppV2) GetIndexDir() string {
	if o == nil || IsNil(o.IndexDir.Get()) {
		var ret string
		return ret
	}
	return *o.IndexDir.Get()
}

// GetIndexDirOk returns a tuple with the IndexDir field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppV2) GetIndexDirOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IndexDir.Get(), o.IndexDir.IsSet()
}

// HasIndexDir returns a boolean if a field has been set.
func (o *AppV2) HasIndexDir() bool {
	if o != nil && o.IndexDir.IsSet() {
		return true
	}

	return false
}

// SetIndexDir gets a reference to the given NullableString and assigns it to the IndexDir field.
func (o *AppV2) SetIndexDir(v string) {
	o.IndexDir.Set(&v)
}
// SetIndexDirNil sets the value for IndexDir to be an explicit nil
func (o *AppV2) SetIndexDirNil() {
	o.IndexDir.Set(nil)
}

// UnsetIndexDir ensures that no value is present for IndexDir, not even an explicit nil
func (o *AppV2) UnsetIndexDir() {
	o.IndexDir.Unset()
}

// GetBuildCmd returns the BuildCmd field value if set, zero value otherwise.
func (o *AppV2) GetBuildCmd() string {
	if o == nil || IsNil(o.BuildCmd) {
		var ret string
		return ret
	}
	return *o.BuildCmd
}

// GetBuildCmdOk returns a tuple with the BuildCmd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppV2) GetBuildCmdOk() (*string, bool) {
	if o == nil || IsNil(o.BuildCmd) {
		return nil, false
	}
	return o.BuildCmd, true
}

// HasBuildCmd returns a boolean if a field has been set.
func (o *AppV2) HasBuildCmd() bool {
	if o != nil && !IsNil(o.BuildCmd) {
		return true
	}

	return false
}

// SetBuildCmd gets a reference to the given string and assigns it to the BuildCmd field.
func (o *AppV2) SetBuildCmd(v string) {
	o.BuildCmd = &v
}

// GetAvatarLink returns the AvatarLink field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AppV2) GetAvatarLink() string {
	if o == nil || o.AvatarLink.Get() == nil {
		var ret string
		return ret
	}

	return *o.AvatarLink.Get()
}

// GetAvatarLinkOk returns a tuple with the AvatarLink field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppV2) GetAvatarLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvatarLink.Get(), o.AvatarLink.IsSet()
}

// SetAvatarLink sets field value
func (o *AppV2) SetAvatarLink(v string) {
	o.AvatarLink.Set(&v)
}

// GetRunCmd returns the RunCmd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppV2) GetRunCmd() string {
	if o == nil || IsNil(o.RunCmd.Get()) {
		var ret string
		return ret
	}
	return *o.RunCmd.Get()
}

// GetRunCmdOk returns a tuple with the RunCmd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppV2) GetRunCmdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RunCmd.Get(), o.RunCmd.IsSet()
}

// HasRunCmd returns a boolean if a field has been set.
func (o *AppV2) HasRunCmd() bool {
	if o != nil && o.RunCmd.IsSet() {
		return true
	}

	return false
}

// SetRunCmd gets a reference to the given NullableString and assigns it to the RunCmd field.
func (o *AppV2) SetRunCmd(v string) {
	o.RunCmd.Set(&v)
}
// SetRunCmdNil sets the value for RunCmd to be an explicit nil
func (o *AppV2) SetRunCmdNil() {
	o.RunCmd.Set(nil)
}

// UnsetRunCmd ensures that no value is present for RunCmd, not even an explicit nil
func (o *AppV2) UnsetRunCmd() {
	o.RunCmd.Unset()
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppV2) GetConfiguration() AppV2Configuration {
	if o == nil || IsNil(o.Configuration.Get()) {
		var ret AppV2Configuration
		return ret
	}
	return *o.Configuration.Get()
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppV2) GetConfigurationOk() (*AppV2Configuration, bool) {
	if o == nil {
		return nil, false
	}
	return o.Configuration.Get(), o.Configuration.IsSet()
}

// HasConfiguration returns a boolean if a field has been set.
func (o *AppV2) HasConfiguration() bool {
	if o != nil && o.Configuration.IsSet() {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given NullableAppV2Configuration and assigns it to the Configuration field.
func (o *AppV2) SetConfiguration(v AppV2Configuration) {
	o.Configuration.Set(&v)
}
// SetConfigurationNil sets the value for Configuration to be an explicit nil
func (o *AppV2) SetConfigurationNil() {
	o.Configuration.Set(nil)
}

// UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil
func (o *AppV2) UnsetConfiguration() {
	o.Configuration.Unset()
}

// GetDiskStatus returns the DiskStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppV2) GetDiskStatus() AppV2DiskStatus {
	if o == nil || IsNil(o.DiskStatus.Get()) {
		var ret AppV2DiskStatus
		return ret
	}
	return *o.DiskStatus.Get()
}

// GetDiskStatusOk returns a tuple with the DiskStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppV2) GetDiskStatusOk() (*AppV2DiskStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiskStatus.Get(), o.DiskStatus.IsSet()
}

// HasDiskStatus returns a boolean if a field has been set.
func (o *AppV2) HasDiskStatus() bool {
	if o != nil && o.DiskStatus.IsSet() {
		return true
	}

	return false
}

// SetDiskStatus gets a reference to the given NullableAppV2DiskStatus and assigns it to the DiskStatus field.
func (o *AppV2) SetDiskStatus(v AppV2DiskStatus) {
	o.DiskStatus.Set(&v)
}
// SetDiskStatusNil sets the value for DiskStatus to be an explicit nil
func (o *AppV2) SetDiskStatusNil() {
	o.DiskStatus.Set(nil)
}

// UnsetDiskStatus ensures that no value is present for DiskStatus, not even an explicit nil
func (o *AppV2) UnsetDiskStatus() {
	o.DiskStatus.Unset()
}

// GetIsQemuAgent returns the IsQemuAgent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppV2) GetIsQemuAgent() bool {
	if o == nil || IsNil(o.IsQemuAgent.Get()) {
		var ret bool
		return ret
	}
	return *o.IsQemuAgent.Get()
}

// GetIsQemuAgentOk returns a tuple with the IsQemuAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppV2) GetIsQemuAgentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsQemuAgent.Get(), o.IsQemuAgent.IsSet()
}

// HasIsQemuAgent returns a boolean if a field has been set.
func (o *AppV2) HasIsQemuAgent() bool {
	if o != nil && o.IsQemuAgent.IsSet() {
		return true
	}

	return false
}

// SetIsQemuAgent gets a reference to the given NullableBool and assigns it to the IsQemuAgent field.
func (o *AppV2) SetIsQemuAgent(v bool) {
	o.IsQemuAgent.Set(&v)
}
// SetIsQemuAgentNil sets the value for IsQemuAgent to be an explicit nil
func (o *AppV2) SetIsQemuAgentNil() {
	o.IsQemuAgent.Set(nil)
}

// UnsetIsQemuAgent ensures that no value is present for IsQemuAgent, not even an explicit nil
func (o *AppV2) UnsetIsQemuAgent() {
	o.IsQemuAgent.Unset()
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *AppV2) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppV2) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *AppV2) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *AppV2) SetLanguage(v string) {
	o.Language = &v
}

// GetStartTime returns the StartTime field value
func (o *AppV2) GetStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *AppV2) GetStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *AppV2) SetStartTime(v time.Time) {
	o.StartTime = v
}

// GetBranch returns the Branch field value if set, zero value otherwise.
func (o *AppV2) GetBranch() string {
	if o == nil || IsNil(o.Branch) {
		var ret string
		return ret
	}
	return *o.Branch
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppV2) GetBranchOk() (*string, bool) {
	if o == nil || IsNil(o.Branch) {
		return nil, false
	}
	return o.Branch, true
}

// HasBranch returns a boolean if a field has been set.
func (o *AppV2) HasBranch() bool {
	if o != nil && !IsNil(o.Branch) {
		return true
	}

	return false
}

// SetBranch gets a reference to the given string and assigns it to the Branch field.
func (o *AppV2) SetBranch(v string) {
	o.Branch = &v
}

// GetConfiguratorId returns the ConfiguratorId field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *AppV2) GetConfiguratorId() int32 {
	if o == nil || o.ConfiguratorId.Get() == nil {
		var ret int32
		return ret
	}

	return *o.ConfiguratorId.Get()
}

// GetConfiguratorIdOk returns a tuple with the ConfiguratorId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppV2) GetConfiguratorIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfiguratorId.Get(), o.ConfiguratorId.IsSet()
}

// SetConfiguratorId sets field value
func (o *AppV2) SetConfiguratorId(v int32) {
	o.ConfiguratorId.Set(&v)
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *AppV2) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppV2) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *AppV2) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *AppV2) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetSystemDependencies returns the SystemDependencies field value if set, zero value otherwise.
func (o *AppV2) GetSystemDependencies() []string {
	if o == nil || IsNil(o.SystemDependencies) {
		var ret []string
		return ret
	}
	return o.SystemDependencies
}

// GetSystemDependenciesOk returns a tuple with the SystemDependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppV2) GetSystemDependenciesOk() ([]string, bool) {
	if o == nil || IsNil(o.SystemDependencies) {
		return nil, false
	}
	return o.SystemDependencies, true
}

// HasSystemDependencies returns a boolean if a field has been set.
func (o *AppV2) HasSystemDependencies() bool {
	if o != nil && !IsNil(o.SystemDependencies) {
		return true
	}

	return false
}

// SetSystemDependencies gets a reference to the given []string and assigns it to the SystemDependencies field.
func (o *AppV2) SetSystemDependencies(v []string) {
	o.SystemDependencies = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *AppV2) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AppV2) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AppV2) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetAvailabilityZone returns the AvailabilityZone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppV2) GetAvailabilityZone() string {
	if o == nil || IsNil(o.AvailabilityZone.Get()) {
		var ret string
		return ret
	}
	return *o.AvailabilityZone.Get()
}

// GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppV2) GetAvailabilityZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailabilityZone.Get(), o.AvailabilityZone.IsSet()
}

// HasAvailabilityZone returns a boolean if a field has been set.
func (o *AppV2) HasAvailabilityZone() bool {
	if o != nil && o.AvailabilityZone.IsSet() {
		return true
	}

	return false
}

// SetAvailabilityZone gets a reference to the given NullableString and assigns it to the AvailabilityZone field.
func (o *AppV2) SetAvailabilityZone(v string) {
	o.AvailabilityZone.Set(&v)
}
// SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil
func (o *AppV2) SetAvailabilityZoneNil() {
	o.AvailabilityZone.Set(nil)
}

// UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
func (o *AppV2) UnsetAvailabilityZone() {
	o.AvailabilityZone.Unset()
}

// GetRequestsLimit returns the RequestsLimit field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *AppV2) GetRequestsLimit() int32 {
	if o == nil || o.RequestsLimit.Get() == nil {
		var ret int32
		return ret
	}

	return *o.RequestsLimit.Get()
}

// GetRequestsLimitOk returns a tuple with the RequestsLimit field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppV2) GetRequestsLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestsLimit.Get(), o.RequestsLimit.IsSet()
}

// SetRequestsLimit sets field value
func (o *AppV2) SetRequestsLimit(v int32) {
	o.RequestsLimit.Set(&v)
}

// GetBuildEntryPath returns the BuildEntryPath field value if set, zero value otherwise.
func (o *AppV2) GetBuildEntryPath() string {
	if o == nil || IsNil(o.BuildEntryPath) {
		var ret string
		return ret
	}
	return *o.BuildEntryPath
}

// GetBuildEntryPathOk returns a tuple with the BuildEntryPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppV2) GetBuildEntryPathOk() (*string, bool) {
	if o == nil || IsNil(o.BuildEntryPath) {
		return nil, false
	}
	return o.BuildEntryPath, true
}

// HasBuildEntryPath returns a boolean if a field has been set.
func (o *AppV2) HasBuildEntryPath() bool {
	if o != nil && !IsNil(o.BuildEntryPath) {
		return true
	}

	return false
}

// SetBuildEntryPath gets a reference to the given string and assigns it to the BuildEntryPath field.
func (o *AppV2) SetBuildEntryPath(v string) {
	o.BuildEntryPath = &v
}

// GetHealthCheck returns the HealthCheck field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppV2) GetHealthCheck() string {
	if o == nil || IsNil(o.HealthCheck.Get()) {
		var ret string
		return ret
	}
	return *o.HealthCheck.Get()
}

// GetHealthCheckOk returns a tuple with the HealthCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppV2) GetHealthCheckOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HealthCheck.Get(), o.HealthCheck.IsSet()
}

// HasHealthCheck returns a boolean if a field has been set.
func (o *AppV2) HasHealthCheck() bool {
	if o != nil && o.HealthCheck.IsSet() {
		return true
	}

	return false
}

// SetHealthCheck gets a reference to the given NullableString and assigns it to the HealthCheck field.
func (o *AppV2) SetHealthCheck(v string) {
	o.HealthCheck.Set(&v)
}
// SetHealthCheckNil sets the value for HealthCheck to be an explicit nil
func (o *AppV2) SetHealthCheckNil() {
	o.HealthCheck.Set(nil)
}

// UnsetHealthCheck ensures that no value is present for HealthCheck, not even an explicit nil
func (o *AppV2) UnsetHealthCheck() {
	o.HealthCheck.Unset()
}

// GetDockerImage returns the DockerImage field value if set, zero value otherwise.
func (o *AppV2) GetDockerImage() string {
	if o == nil || IsNil(o.DockerImage) {
		var ret string
		return ret
	}
	return *o.DockerImage
}

// GetDockerImageOk returns a tuple with the DockerImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppV2) GetDockerImageOk() (*string, bool) {
	if o == nil || IsNil(o.DockerImage) {
		return nil, false
	}
	return o.DockerImage, true
}

// HasDockerImage returns a boolean if a field has been set.
func (o *AppV2) HasDockerImage() bool {
	if o != nil && !IsNil(o.DockerImage) {
		return true
	}

	return false
}

// SetDockerImage gets a reference to the given string and assigns it to the DockerImage field.
func (o *AppV2) SetDockerImage(v string) {
	o.DockerImage = &v
}

// GetDockerImageTag returns the DockerImageTag field value if set, zero value otherwise.
func (o *AppV2) GetDockerImageTag() string {
	if o == nil || IsNil(o.DockerImageTag) {
		var ret string
		return ret
	}
	return *o.DockerImageTag
}

// GetDockerImageTagOk returns a tuple with the DockerImageTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppV2) GetDockerImageTagOk() (*string, bool) {
	if o == nil || IsNil(o.DockerImageTag) {
		return nil, false
	}
	return o.DockerImageTag, true
}

// HasDockerImageTag returns a boolean if a field has been set.
func (o *AppV2) HasDockerImageTag() bool {
	if o != nil && !IsNil(o.DockerImageTag) {
		return true
	}

	return false
}

// SetDockerImageTag gets a reference to the given string and assigns it to the DockerImageTag field.
func (o *AppV2) SetDockerImageTag(v string) {
	o.DockerImageTag = &v
}

// GetIsRussianTrustedCa returns the IsRussianTrustedCa field value if set, zero value otherwise.
func (o *AppV2) GetIsRussianTrustedCa() bool {
	if o == nil || IsNil(o.IsRussianTrustedCa) {
		var ret bool
		return ret
	}
	return *o.IsRussianTrustedCa
}

// GetIsRussianTrustedCaOk returns a tuple with the IsRussianTrustedCa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppV2) GetIsRussianTrustedCaOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRussianTrustedCa) {
		return nil, false
	}
	return o.IsRussianTrustedCa, true
}

// HasIsRussianTrustedCa returns a boolean if a field has been set.
func (o *AppV2) HasIsRussianTrustedCa() bool {
	if o != nil && !IsNil(o.IsRussianTrustedCa) {
		return true
	}

	return false
}

// SetIsRussianTrustedCa gets a reference to the given bool and assigns it to the IsRussianTrustedCa field.
func (o *AppV2) SetIsRussianTrustedCa(v bool) {
	o.IsRussianTrustedCa = &v
}

// GetIsTciTrustedCa returns the IsTciTrustedCa field value if set, zero value otherwise.
func (o *AppV2) GetIsTciTrustedCa() bool {
	if o == nil || IsNil(o.IsTciTrustedCa) {
		var ret bool
		return ret
	}
	return *o.IsTciTrustedCa
}

// GetIsTciTrustedCaOk returns a tuple with the IsTciTrustedCa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppV2) GetIsTciTrustedCaOk() (*bool, bool) {
	if o == nil || IsNil(o.IsTciTrustedCa) {
		return nil, false
	}
	return o.IsTciTrustedCa, true
}

// HasIsTciTrustedCa returns a boolean if a field has been set.
func (o *AppV2) HasIsTciTrustedCa() bool {
	if o != nil && !IsNil(o.IsTciTrustedCa) {
		return true
	}

	return false
}

// SetIsTciTrustedCa gets a reference to the given bool and assigns it to the IsTciTrustedCa field.
func (o *AppV2) SetIsTciTrustedCa(v bool) {
	o.IsTciTrustedCa = &v
}

// GetBlockedPorts returns the BlockedPorts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppV2) GetBlockedPorts() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.BlockedPorts
}

// GetBlockedPortsOk returns a tuple with the BlockedPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppV2) GetBlockedPortsOk() ([]int32, bool) {
	if o == nil || IsNil(o.BlockedPorts) {
		return nil, false
	}
	return o.BlockedPorts, true
}

// HasBlockedPorts returns a boolean if a field has been set.
func (o *AppV2) HasBlockedPorts() bool {
	if o != nil && IsNil(o.BlockedPorts) {
		return true
	}

	return false
}

// SetBlockedPorts gets a reference to the given []int32 and assigns it to the BlockedPorts field.
func (o *AppV2) SetBlockedPorts(v []int32) {
	o.BlockedPorts = v
}

// GetReverseProxy returns the ReverseProxy field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *AppV2) GetReverseProxy() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ReverseProxy
}

// GetReverseProxyOk returns a tuple with the ReverseProxy field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppV2) GetReverseProxyOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ReverseProxy) {
		return map[string]interface{}{}, false
	}
	return o.ReverseProxy, true
}

// SetReverseProxy sets field value
func (o *AppV2) SetReverseProxy(v map[string]interface{}) {
	o.ReverseProxy = v
}

func (o AppV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name
	toSerialize["status"] = o.Status
	toSerialize["provider"] = o.Provider
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	toSerialize["domains"] = o.Domains
	if !IsNil(o.Framework) {
		toSerialize["framework"] = o.Framework
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.Repository) {
		toSerialize["repository"] = o.Repository
	}
	if o.EnvVersion.IsSet() {
		toSerialize["env_version"] = o.EnvVersion.Get()
	}
	if !IsNil(o.Envs) {
		toSerialize["envs"] = o.Envs
	}
	toSerialize["is_auto_deploy"] = o.IsAutoDeploy
	if !IsNil(o.CommitSha) {
		toSerialize["commit_sha"] = o.CommitSha
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	toSerialize["preset_id"] = o.PresetId
	if o.IndexDir.IsSet() {
		toSerialize["index_dir"] = o.IndexDir.Get()
	}
	if !IsNil(o.BuildCmd) {
		toSerialize["build_cmd"] = o.BuildCmd
	}
	toSerialize["avatar_link"] = o.AvatarLink.Get()
	if o.RunCmd.IsSet() {
		toSerialize["run_cmd"] = o.RunCmd.Get()
	}
	if o.Configuration.IsSet() {
		toSerialize["configuration"] = o.Configuration.Get()
	}
	if o.DiskStatus.IsSet() {
		toSerialize["disk_status"] = o.DiskStatus.Get()
	}
	if o.IsQemuAgent.IsSet() {
		toSerialize["is_qemu_agent"] = o.IsQemuAgent.Get()
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	toSerialize["start_time"] = o.StartTime
	if !IsNil(o.Branch) {
		toSerialize["branch"] = o.Branch
	}
	toSerialize["configurator_id"] = o.ConfiguratorId.Get()
	if !IsNil(o.ProjectId) {
		toSerialize["project_id"] = o.ProjectId
	}
	if !IsNil(o.SystemDependencies) {
		toSerialize["system_dependencies"] = o.SystemDependencies
	}
	toSerialize["created_at"] = o.CreatedAt
	if o.AvailabilityZone.IsSet() {
		toSerialize["availability_zone"] = o.AvailabilityZone.Get()
	}
	toSerialize["requests_limit"] = o.RequestsLimit.Get()
	if !IsNil(o.BuildEntryPath) {
		toSerialize["build_entry_path"] = o.BuildEntryPath
	}
	if o.HealthCheck.IsSet() {
		toSerialize["health_check"] = o.HealthCheck.Get()
	}
	if !IsNil(o.DockerImage) {
		toSerialize["docker_image"] = o.DockerImage
	}
	if !IsNil(o.DockerImageTag) {
		toSerialize["docker_image_tag"] = o.DockerImageTag
	}
	if !IsNil(o.IsRussianTrustedCa) {
		toSerialize["is_russian_trusted_ca"] = o.IsRussianTrustedCa
	}
	if !IsNil(o.IsTciTrustedCa) {
		toSerialize["is_tci_trusted_ca"] = o.IsTciTrustedCa
	}
	if o.BlockedPorts != nil {
		toSerialize["blocked_ports"] = o.BlockedPorts
	}
	if o.ReverseProxy != nil {
		toSerialize["reverse_proxy"] = o.ReverseProxy
	}
	return toSerialize, nil
}

type NullableAppV2 struct {
	value *AppV2
	isSet bool
}

func (v NullableAppV2) Get() *AppV2 {
	return v.value
}

func (v *NullableAppV2) Set(val *AppV2) {
	v.value = val
	v.isSet = true
}

func (v NullableAppV2) IsSet() bool {
	return v.isSet
}

func (v *NullableAppV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppV2(val *AppV2) *NullableAppV2 {
	return &NullableAppV2{value: val, isSet: true}
}

func (v NullableAppV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


