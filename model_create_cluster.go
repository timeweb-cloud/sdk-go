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

// checks if the CreateCluster type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCluster{}

// CreateCluster struct for CreateCluster
type CreateCluster struct {
	// Название кластера базы данных.
	Name string `json:"name"`
	Type DbType `json:"type"`
	Admin *CreateClusterAdmin `json:"admin,omitempty"`
	Instance *CreateClusterInstance `json:"instance,omitempty"`
	// Тип хеширования базы данных (mysql5 | mysql | postgres).
	HashType *string `json:"hash_type,omitempty"`
	// ID тарифа.
	PresetId int32 `json:"preset_id"`
	ConfigParameters *ConfigParameters `json:"config_parameters,omitempty"`
	Network *Network `json:"network,omitempty"`
	// Описание кластера базы данных
	Description *string `json:"description,omitempty"`
	AvailabilityZone *AvailabilityZone `json:"availability_zone,omitempty"`
	AutoBackups *CreateDbAutoBackups `json:"auto_backups,omitempty"`
}

// NewCreateCluster instantiates a new CreateCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCluster(name string, type_ DbType, presetId int32) *CreateCluster {
	this := CreateCluster{}
	this.Name = name
	this.Type = type_
	this.PresetId = presetId
	return &this
}

// NewCreateClusterWithDefaults instantiates a new CreateCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateClusterWithDefaults() *CreateCluster {
	this := CreateCluster{}
	return &this
}

// GetName returns the Name field value
func (o *CreateCluster) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateCluster) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateCluster) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *CreateCluster) GetType() DbType {
	if o == nil {
		var ret DbType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateCluster) GetTypeOk() (*DbType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateCluster) SetType(v DbType) {
	o.Type = v
}

// GetAdmin returns the Admin field value if set, zero value otherwise.
func (o *CreateCluster) GetAdmin() CreateClusterAdmin {
	if o == nil || IsNil(o.Admin) {
		var ret CreateClusterAdmin
		return ret
	}
	return *o.Admin
}

// GetAdminOk returns a tuple with the Admin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCluster) GetAdminOk() (*CreateClusterAdmin, bool) {
	if o == nil || IsNil(o.Admin) {
		return nil, false
	}
	return o.Admin, true
}

// HasAdmin returns a boolean if a field has been set.
func (o *CreateCluster) HasAdmin() bool {
	if o != nil && !IsNil(o.Admin) {
		return true
	}

	return false
}

// SetAdmin gets a reference to the given CreateClusterAdmin and assigns it to the Admin field.
func (o *CreateCluster) SetAdmin(v CreateClusterAdmin) {
	o.Admin = &v
}

// GetInstance returns the Instance field value if set, zero value otherwise.
func (o *CreateCluster) GetInstance() CreateClusterInstance {
	if o == nil || IsNil(o.Instance) {
		var ret CreateClusterInstance
		return ret
	}
	return *o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCluster) GetInstanceOk() (*CreateClusterInstance, bool) {
	if o == nil || IsNil(o.Instance) {
		return nil, false
	}
	return o.Instance, true
}

// HasInstance returns a boolean if a field has been set.
func (o *CreateCluster) HasInstance() bool {
	if o != nil && !IsNil(o.Instance) {
		return true
	}

	return false
}

// SetInstance gets a reference to the given CreateClusterInstance and assigns it to the Instance field.
func (o *CreateCluster) SetInstance(v CreateClusterInstance) {
	o.Instance = &v
}

// GetHashType returns the HashType field value if set, zero value otherwise.
func (o *CreateCluster) GetHashType() string {
	if o == nil || IsNil(o.HashType) {
		var ret string
		return ret
	}
	return *o.HashType
}

// GetHashTypeOk returns a tuple with the HashType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCluster) GetHashTypeOk() (*string, bool) {
	if o == nil || IsNil(o.HashType) {
		return nil, false
	}
	return o.HashType, true
}

// HasHashType returns a boolean if a field has been set.
func (o *CreateCluster) HasHashType() bool {
	if o != nil && !IsNil(o.HashType) {
		return true
	}

	return false
}

// SetHashType gets a reference to the given string and assigns it to the HashType field.
func (o *CreateCluster) SetHashType(v string) {
	o.HashType = &v
}

// GetPresetId returns the PresetId field value
func (o *CreateCluster) GetPresetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PresetId
}

// GetPresetIdOk returns a tuple with the PresetId field value
// and a boolean to check if the value has been set.
func (o *CreateCluster) GetPresetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PresetId, true
}

// SetPresetId sets field value
func (o *CreateCluster) SetPresetId(v int32) {
	o.PresetId = v
}

// GetConfigParameters returns the ConfigParameters field value if set, zero value otherwise.
func (o *CreateCluster) GetConfigParameters() ConfigParameters {
	if o == nil || IsNil(o.ConfigParameters) {
		var ret ConfigParameters
		return ret
	}
	return *o.ConfigParameters
}

// GetConfigParametersOk returns a tuple with the ConfigParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCluster) GetConfigParametersOk() (*ConfigParameters, bool) {
	if o == nil || IsNil(o.ConfigParameters) {
		return nil, false
	}
	return o.ConfigParameters, true
}

// HasConfigParameters returns a boolean if a field has been set.
func (o *CreateCluster) HasConfigParameters() bool {
	if o != nil && !IsNil(o.ConfigParameters) {
		return true
	}

	return false
}

// SetConfigParameters gets a reference to the given ConfigParameters and assigns it to the ConfigParameters field.
func (o *CreateCluster) SetConfigParameters(v ConfigParameters) {
	o.ConfigParameters = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *CreateCluster) GetNetwork() Network {
	if o == nil || IsNil(o.Network) {
		var ret Network
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCluster) GetNetworkOk() (*Network, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *CreateCluster) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given Network and assigns it to the Network field.
func (o *CreateCluster) SetNetwork(v Network) {
	o.Network = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateCluster) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCluster) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateCluster) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateCluster) SetDescription(v string) {
	o.Description = &v
}

// GetAvailabilityZone returns the AvailabilityZone field value if set, zero value otherwise.
func (o *CreateCluster) GetAvailabilityZone() AvailabilityZone {
	if o == nil || IsNil(o.AvailabilityZone) {
		var ret AvailabilityZone
		return ret
	}
	return *o.AvailabilityZone
}

// GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCluster) GetAvailabilityZoneOk() (*AvailabilityZone, bool) {
	if o == nil || IsNil(o.AvailabilityZone) {
		return nil, false
	}
	return o.AvailabilityZone, true
}

// HasAvailabilityZone returns a boolean if a field has been set.
func (o *CreateCluster) HasAvailabilityZone() bool {
	if o != nil && !IsNil(o.AvailabilityZone) {
		return true
	}

	return false
}

// SetAvailabilityZone gets a reference to the given AvailabilityZone and assigns it to the AvailabilityZone field.
func (o *CreateCluster) SetAvailabilityZone(v AvailabilityZone) {
	o.AvailabilityZone = &v
}

// GetAutoBackups returns the AutoBackups field value if set, zero value otherwise.
func (o *CreateCluster) GetAutoBackups() CreateDbAutoBackups {
	if o == nil || IsNil(o.AutoBackups) {
		var ret CreateDbAutoBackups
		return ret
	}
	return *o.AutoBackups
}

// GetAutoBackupsOk returns a tuple with the AutoBackups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCluster) GetAutoBackupsOk() (*CreateDbAutoBackups, bool) {
	if o == nil || IsNil(o.AutoBackups) {
		return nil, false
	}
	return o.AutoBackups, true
}

// HasAutoBackups returns a boolean if a field has been set.
func (o *CreateCluster) HasAutoBackups() bool {
	if o != nil && !IsNil(o.AutoBackups) {
		return true
	}

	return false
}

// SetAutoBackups gets a reference to the given CreateDbAutoBackups and assigns it to the AutoBackups field.
func (o *CreateCluster) SetAutoBackups(v CreateDbAutoBackups) {
	o.AutoBackups = &v
}

func (o CreateCluster) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCluster) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	if !IsNil(o.Admin) {
		toSerialize["admin"] = o.Admin
	}
	if !IsNil(o.Instance) {
		toSerialize["instance"] = o.Instance
	}
	if !IsNil(o.HashType) {
		toSerialize["hash_type"] = o.HashType
	}
	toSerialize["preset_id"] = o.PresetId
	if !IsNil(o.ConfigParameters) {
		toSerialize["config_parameters"] = o.ConfigParameters
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.AvailabilityZone) {
		toSerialize["availability_zone"] = o.AvailabilityZone
	}
	if !IsNil(o.AutoBackups) {
		toSerialize["auto_backups"] = o.AutoBackups
	}
	return toSerialize, nil
}

type NullableCreateCluster struct {
	value *CreateCluster
	isSet bool
}

func (v NullableCreateCluster) Get() *CreateCluster {
	return v.value
}

func (v *NullableCreateCluster) Set(val *CreateCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCluster(val *CreateCluster) *NullableCreateCluster {
	return &NullableCreateCluster{value: val, isSet: true}
}

func (v NullableCreateCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


