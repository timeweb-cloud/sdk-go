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
)

// checks if the CreateServer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateServer{}

// CreateServer struct for CreateServer
type CreateServer struct {
	Configuration *CreateServerConfiguration `json:"configuration,omitempty"`
	// Защита от DDoS. Серверу выдается защищенный IP-адрес с защитой уровня L3 / L4. Для включения защиты уровня L7 необходимо создать тикет в техническую поддержку.
	IsDdosGuard *bool `json:"is_ddos_guard,omitempty"`
	// ID операционной системы, которая будет установлена на облачный сервер. Нельзя передавать вместе с `image_id`.
	OsId *float32 `json:"os_id,omitempty"`
	// ID образа, который будет установлен на облачный сервер. Нельзя передавать вместе с `os_id`.
	ImageId *string `json:"image_id,omitempty"`
	// ID программного обеспечения сервера.
	SoftwareId *float32 `json:"software_id,omitempty"`
	// ID тарифа сервера. Нельзя передавать вместе с ключом `configurator`.
	PresetId *float32 `json:"preset_id,omitempty"`
	// Пропускная способность тарифа. Доступные значения от 100 до 1000 с шагом 100.
	Bandwidth *float32 `json:"bandwidth,omitempty"`
	// Имя облачного сервера. Максимальная длина — 255 символов, имя должно быть уникальным.
	Name string `json:"name"`
	// ID аватара сервера. Описание методов работы с аватарами появится позднее.
	AvatarId *string `json:"avatar_id,omitempty"`
	// Комментарий к облачному серверу. Максимальная длина — 255 символов.
	Comment *string `json:"comment,omitempty"`
	// Список SSH-ключей.
	SshKeysIds []float32 `json:"ssh_keys_ids,omitempty"`
	// Локальная сеть.
	// Deprecated
	IsLocalNetwork *bool `json:"is_local_network,omitempty"`
	Network *Network `json:"network,omitempty"`
	// Cloud-init скрипт
	CloudInit *string `json:"cloud_init,omitempty"`
	AvailabilityZone *AvailabilityZone `json:"availability_zone,omitempty"`
}

// NewCreateServer instantiates a new CreateServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateServer(name string) *CreateServer {
	this := CreateServer{}
	this.Name = name
	return &this
}

// NewCreateServerWithDefaults instantiates a new CreateServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateServerWithDefaults() *CreateServer {
	this := CreateServer{}
	return &this
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *CreateServer) GetConfiguration() CreateServerConfiguration {
	if o == nil || IsNil(o.Configuration) {
		var ret CreateServerConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServer) GetConfigurationOk() (*CreateServerConfiguration, bool) {
	if o == nil || IsNil(o.Configuration) {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *CreateServer) HasConfiguration() bool {
	if o != nil && !IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given CreateServerConfiguration and assigns it to the Configuration field.
func (o *CreateServer) SetConfiguration(v CreateServerConfiguration) {
	o.Configuration = &v
}

// GetIsDdosGuard returns the IsDdosGuard field value if set, zero value otherwise.
func (o *CreateServer) GetIsDdosGuard() bool {
	if o == nil || IsNil(o.IsDdosGuard) {
		var ret bool
		return ret
	}
	return *o.IsDdosGuard
}

// GetIsDdosGuardOk returns a tuple with the IsDdosGuard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServer) GetIsDdosGuardOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDdosGuard) {
		return nil, false
	}
	return o.IsDdosGuard, true
}

// HasIsDdosGuard returns a boolean if a field has been set.
func (o *CreateServer) HasIsDdosGuard() bool {
	if o != nil && !IsNil(o.IsDdosGuard) {
		return true
	}

	return false
}

// SetIsDdosGuard gets a reference to the given bool and assigns it to the IsDdosGuard field.
func (o *CreateServer) SetIsDdosGuard(v bool) {
	o.IsDdosGuard = &v
}

// GetOsId returns the OsId field value if set, zero value otherwise.
func (o *CreateServer) GetOsId() float32 {
	if o == nil || IsNil(o.OsId) {
		var ret float32
		return ret
	}
	return *o.OsId
}

// GetOsIdOk returns a tuple with the OsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServer) GetOsIdOk() (*float32, bool) {
	if o == nil || IsNil(o.OsId) {
		return nil, false
	}
	return o.OsId, true
}

// HasOsId returns a boolean if a field has been set.
func (o *CreateServer) HasOsId() bool {
	if o != nil && !IsNil(o.OsId) {
		return true
	}

	return false
}

// SetOsId gets a reference to the given float32 and assigns it to the OsId field.
func (o *CreateServer) SetOsId(v float32) {
	o.OsId = &v
}

// GetImageId returns the ImageId field value if set, zero value otherwise.
func (o *CreateServer) GetImageId() string {
	if o == nil || IsNil(o.ImageId) {
		var ret string
		return ret
	}
	return *o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServer) GetImageIdOk() (*string, bool) {
	if o == nil || IsNil(o.ImageId) {
		return nil, false
	}
	return o.ImageId, true
}

// HasImageId returns a boolean if a field has been set.
func (o *CreateServer) HasImageId() bool {
	if o != nil && !IsNil(o.ImageId) {
		return true
	}

	return false
}

// SetImageId gets a reference to the given string and assigns it to the ImageId field.
func (o *CreateServer) SetImageId(v string) {
	o.ImageId = &v
}

// GetSoftwareId returns the SoftwareId field value if set, zero value otherwise.
func (o *CreateServer) GetSoftwareId() float32 {
	if o == nil || IsNil(o.SoftwareId) {
		var ret float32
		return ret
	}
	return *o.SoftwareId
}

// GetSoftwareIdOk returns a tuple with the SoftwareId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServer) GetSoftwareIdOk() (*float32, bool) {
	if o == nil || IsNil(o.SoftwareId) {
		return nil, false
	}
	return o.SoftwareId, true
}

// HasSoftwareId returns a boolean if a field has been set.
func (o *CreateServer) HasSoftwareId() bool {
	if o != nil && !IsNil(o.SoftwareId) {
		return true
	}

	return false
}

// SetSoftwareId gets a reference to the given float32 and assigns it to the SoftwareId field.
func (o *CreateServer) SetSoftwareId(v float32) {
	o.SoftwareId = &v
}

// GetPresetId returns the PresetId field value if set, zero value otherwise.
func (o *CreateServer) GetPresetId() float32 {
	if o == nil || IsNil(o.PresetId) {
		var ret float32
		return ret
	}
	return *o.PresetId
}

// GetPresetIdOk returns a tuple with the PresetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServer) GetPresetIdOk() (*float32, bool) {
	if o == nil || IsNil(o.PresetId) {
		return nil, false
	}
	return o.PresetId, true
}

// HasPresetId returns a boolean if a field has been set.
func (o *CreateServer) HasPresetId() bool {
	if o != nil && !IsNil(o.PresetId) {
		return true
	}

	return false
}

// SetPresetId gets a reference to the given float32 and assigns it to the PresetId field.
func (o *CreateServer) SetPresetId(v float32) {
	o.PresetId = &v
}

// GetBandwidth returns the Bandwidth field value if set, zero value otherwise.
func (o *CreateServer) GetBandwidth() float32 {
	if o == nil || IsNil(o.Bandwidth) {
		var ret float32
		return ret
	}
	return *o.Bandwidth
}

// GetBandwidthOk returns a tuple with the Bandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServer) GetBandwidthOk() (*float32, bool) {
	if o == nil || IsNil(o.Bandwidth) {
		return nil, false
	}
	return o.Bandwidth, true
}

// HasBandwidth returns a boolean if a field has been set.
func (o *CreateServer) HasBandwidth() bool {
	if o != nil && !IsNil(o.Bandwidth) {
		return true
	}

	return false
}

// SetBandwidth gets a reference to the given float32 and assigns it to the Bandwidth field.
func (o *CreateServer) SetBandwidth(v float32) {
	o.Bandwidth = &v
}

// GetName returns the Name field value
func (o *CreateServer) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateServer) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateServer) SetName(v string) {
	o.Name = v
}

// GetAvatarId returns the AvatarId field value if set, zero value otherwise.
func (o *CreateServer) GetAvatarId() string {
	if o == nil || IsNil(o.AvatarId) {
		var ret string
		return ret
	}
	return *o.AvatarId
}

// GetAvatarIdOk returns a tuple with the AvatarId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServer) GetAvatarIdOk() (*string, bool) {
	if o == nil || IsNil(o.AvatarId) {
		return nil, false
	}
	return o.AvatarId, true
}

// HasAvatarId returns a boolean if a field has been set.
func (o *CreateServer) HasAvatarId() bool {
	if o != nil && !IsNil(o.AvatarId) {
		return true
	}

	return false
}

// SetAvatarId gets a reference to the given string and assigns it to the AvatarId field.
func (o *CreateServer) SetAvatarId(v string) {
	o.AvatarId = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *CreateServer) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServer) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *CreateServer) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *CreateServer) SetComment(v string) {
	o.Comment = &v
}

// GetSshKeysIds returns the SshKeysIds field value if set, zero value otherwise.
func (o *CreateServer) GetSshKeysIds() []float32 {
	if o == nil || IsNil(o.SshKeysIds) {
		var ret []float32
		return ret
	}
	return o.SshKeysIds
}

// GetSshKeysIdsOk returns a tuple with the SshKeysIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServer) GetSshKeysIdsOk() ([]float32, bool) {
	if o == nil || IsNil(o.SshKeysIds) {
		return nil, false
	}
	return o.SshKeysIds, true
}

// HasSshKeysIds returns a boolean if a field has been set.
func (o *CreateServer) HasSshKeysIds() bool {
	if o != nil && !IsNil(o.SshKeysIds) {
		return true
	}

	return false
}

// SetSshKeysIds gets a reference to the given []float32 and assigns it to the SshKeysIds field.
func (o *CreateServer) SetSshKeysIds(v []float32) {
	o.SshKeysIds = v
}

// GetIsLocalNetwork returns the IsLocalNetwork field value if set, zero value otherwise.
// Deprecated
func (o *CreateServer) GetIsLocalNetwork() bool {
	if o == nil || IsNil(o.IsLocalNetwork) {
		var ret bool
		return ret
	}
	return *o.IsLocalNetwork
}

// GetIsLocalNetworkOk returns a tuple with the IsLocalNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CreateServer) GetIsLocalNetworkOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLocalNetwork) {
		return nil, false
	}
	return o.IsLocalNetwork, true
}

// HasIsLocalNetwork returns a boolean if a field has been set.
func (o *CreateServer) HasIsLocalNetwork() bool {
	if o != nil && !IsNil(o.IsLocalNetwork) {
		return true
	}

	return false
}

// SetIsLocalNetwork gets a reference to the given bool and assigns it to the IsLocalNetwork field.
// Deprecated
func (o *CreateServer) SetIsLocalNetwork(v bool) {
	o.IsLocalNetwork = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *CreateServer) GetNetwork() Network {
	if o == nil || IsNil(o.Network) {
		var ret Network
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServer) GetNetworkOk() (*Network, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *CreateServer) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given Network and assigns it to the Network field.
func (o *CreateServer) SetNetwork(v Network) {
	o.Network = &v
}

// GetCloudInit returns the CloudInit field value if set, zero value otherwise.
func (o *CreateServer) GetCloudInit() string {
	if o == nil || IsNil(o.CloudInit) {
		var ret string
		return ret
	}
	return *o.CloudInit
}

// GetCloudInitOk returns a tuple with the CloudInit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServer) GetCloudInitOk() (*string, bool) {
	if o == nil || IsNil(o.CloudInit) {
		return nil, false
	}
	return o.CloudInit, true
}

// HasCloudInit returns a boolean if a field has been set.
func (o *CreateServer) HasCloudInit() bool {
	if o != nil && !IsNil(o.CloudInit) {
		return true
	}

	return false
}

// SetCloudInit gets a reference to the given string and assigns it to the CloudInit field.
func (o *CreateServer) SetCloudInit(v string) {
	o.CloudInit = &v
}

// GetAvailabilityZone returns the AvailabilityZone field value if set, zero value otherwise.
func (o *CreateServer) GetAvailabilityZone() AvailabilityZone {
	if o == nil || IsNil(o.AvailabilityZone) {
		var ret AvailabilityZone
		return ret
	}
	return *o.AvailabilityZone
}

// GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServer) GetAvailabilityZoneOk() (*AvailabilityZone, bool) {
	if o == nil || IsNil(o.AvailabilityZone) {
		return nil, false
	}
	return o.AvailabilityZone, true
}

// HasAvailabilityZone returns a boolean if a field has been set.
func (o *CreateServer) HasAvailabilityZone() bool {
	if o != nil && !IsNil(o.AvailabilityZone) {
		return true
	}

	return false
}

// SetAvailabilityZone gets a reference to the given AvailabilityZone and assigns it to the AvailabilityZone field.
func (o *CreateServer) SetAvailabilityZone(v AvailabilityZone) {
	o.AvailabilityZone = &v
}

func (o CreateServer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateServer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}
	if !IsNil(o.IsDdosGuard) {
		toSerialize["is_ddos_guard"] = o.IsDdosGuard
	}
	if !IsNil(o.OsId) {
		toSerialize["os_id"] = o.OsId
	}
	if !IsNil(o.ImageId) {
		toSerialize["image_id"] = o.ImageId
	}
	if !IsNil(o.SoftwareId) {
		toSerialize["software_id"] = o.SoftwareId
	}
	if !IsNil(o.PresetId) {
		toSerialize["preset_id"] = o.PresetId
	}
	if !IsNil(o.Bandwidth) {
		toSerialize["bandwidth"] = o.Bandwidth
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.AvatarId) {
		toSerialize["avatar_id"] = o.AvatarId
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.SshKeysIds) {
		toSerialize["ssh_keys_ids"] = o.SshKeysIds
	}
	if !IsNil(o.IsLocalNetwork) {
		toSerialize["is_local_network"] = o.IsLocalNetwork
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.CloudInit) {
		toSerialize["cloud_init"] = o.CloudInit
	}
	if !IsNil(o.AvailabilityZone) {
		toSerialize["availability_zone"] = o.AvailabilityZone
	}
	return toSerialize, nil
}

type NullableCreateServer struct {
	value *CreateServer
	isSet bool
}

func (v NullableCreateServer) Get() *CreateServer {
	return v.value
}

func (v *NullableCreateServer) Set(val *CreateServer) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateServer) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateServer(val *CreateServer) *NullableCreateServer {
	return &NullableCreateServer{value: val, isSet: true}
}

func (v NullableCreateServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


