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

// checks if the UpdateClusterV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateClusterV2{}

// UpdateClusterV2 Изменение кластера базы данных (v2)
type UpdateClusterV2 struct {
	// Название кластера базы данных.
	Name *string `json:"name,omitempty"`
	// ID тарифа. Нельзя передавать вместе с `configuration`
	PresetId *int32 `json:"preset_id,omitempty"`
	Configuration *UpdateClusterConfiguration `json:"configuration,omitempty"`
	ConfigParameters *Mysql `json:"config_parameters,omitempty"`
	// Тип хеширования базы данных (mysql | postgres).
	HashType *string `json:"hash_type,omitempty"`
	// Описание кластера базы данных
	Description *string `json:"description,omitempty"`
	// Доступность публичного IP-адреса
	IsEnabledPublicNetwork *bool `json:"is_enabled_public_network,omitempty"`
	// Использование публичного IPv6-адреса.
	IsEnabledPublicIpv6 *bool `json:"is_enabled_public_ipv6,omitempty"`
	// Плавающий IP-адрес, который нужно привязать к кластеру базы данных. Передается сам адрес, а не его ID; адрес должен быть свободен (не привязан к другому сервису).
	FloatingIp *string `json:"floating_ip,omitempty"`
	// Включить защищенное подключение к кластеру базы данных. Обратите внимание: в ответе это же значение возвращается под ключом `is_secure_connection_enabled`.
	IsSecureConnectionEnable *bool `json:"is_secure_connection_enable,omitempty"`
	MaintenanceSlot *UpdateClusterV2MaintenanceSlot `json:"maintenance_slot,omitempty"`
	DiskAutoscaling *UpdateClusterV2DiskAutoscaling `json:"disk_autoscaling,omitempty"`
}

// NewUpdateClusterV2 instantiates a new UpdateClusterV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateClusterV2() *UpdateClusterV2 {
	this := UpdateClusterV2{}
	return &this
}

// NewUpdateClusterV2WithDefaults instantiates a new UpdateClusterV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateClusterV2WithDefaults() *UpdateClusterV2 {
	this := UpdateClusterV2{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateClusterV2) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateClusterV2) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateClusterV2) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateClusterV2) SetName(v string) {
	o.Name = &v
}

// GetPresetId returns the PresetId field value if set, zero value otherwise.
func (o *UpdateClusterV2) GetPresetId() int32 {
	if o == nil || IsNil(o.PresetId) {
		var ret int32
		return ret
	}
	return *o.PresetId
}

// GetPresetIdOk returns a tuple with the PresetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateClusterV2) GetPresetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.PresetId) {
		return nil, false
	}
	return o.PresetId, true
}

// HasPresetId returns a boolean if a field has been set.
func (o *UpdateClusterV2) HasPresetId() bool {
	if o != nil && !IsNil(o.PresetId) {
		return true
	}

	return false
}

// SetPresetId gets a reference to the given int32 and assigns it to the PresetId field.
func (o *UpdateClusterV2) SetPresetId(v int32) {
	o.PresetId = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *UpdateClusterV2) GetConfiguration() UpdateClusterConfiguration {
	if o == nil || IsNil(o.Configuration) {
		var ret UpdateClusterConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateClusterV2) GetConfigurationOk() (*UpdateClusterConfiguration, bool) {
	if o == nil || IsNil(o.Configuration) {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *UpdateClusterV2) HasConfiguration() bool {
	if o != nil && !IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given UpdateClusterConfiguration and assigns it to the Configuration field.
func (o *UpdateClusterV2) SetConfiguration(v UpdateClusterConfiguration) {
	o.Configuration = &v
}

// GetConfigParameters returns the ConfigParameters field value if set, zero value otherwise.
func (o *UpdateClusterV2) GetConfigParameters() Mysql {
	if o == nil || IsNil(o.ConfigParameters) {
		var ret Mysql
		return ret
	}
	return *o.ConfigParameters
}

// GetConfigParametersOk returns a tuple with the ConfigParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateClusterV2) GetConfigParametersOk() (*Mysql, bool) {
	if o == nil || IsNil(o.ConfigParameters) {
		return nil, false
	}
	return o.ConfigParameters, true
}

// HasConfigParameters returns a boolean if a field has been set.
func (o *UpdateClusterV2) HasConfigParameters() bool {
	if o != nil && !IsNil(o.ConfigParameters) {
		return true
	}

	return false
}

// SetConfigParameters gets a reference to the given Mysql and assigns it to the ConfigParameters field.
func (o *UpdateClusterV2) SetConfigParameters(v Mysql) {
	o.ConfigParameters = &v
}

// GetHashType returns the HashType field value if set, zero value otherwise.
func (o *UpdateClusterV2) GetHashType() string {
	if o == nil || IsNil(o.HashType) {
		var ret string
		return ret
	}
	return *o.HashType
}

// GetHashTypeOk returns a tuple with the HashType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateClusterV2) GetHashTypeOk() (*string, bool) {
	if o == nil || IsNil(o.HashType) {
		return nil, false
	}
	return o.HashType, true
}

// HasHashType returns a boolean if a field has been set.
func (o *UpdateClusterV2) HasHashType() bool {
	if o != nil && !IsNil(o.HashType) {
		return true
	}

	return false
}

// SetHashType gets a reference to the given string and assigns it to the HashType field.
func (o *UpdateClusterV2) SetHashType(v string) {
	o.HashType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateClusterV2) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateClusterV2) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateClusterV2) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateClusterV2) SetDescription(v string) {
	o.Description = &v
}

// GetIsEnabledPublicNetwork returns the IsEnabledPublicNetwork field value if set, zero value otherwise.
func (o *UpdateClusterV2) GetIsEnabledPublicNetwork() bool {
	if o == nil || IsNil(o.IsEnabledPublicNetwork) {
		var ret bool
		return ret
	}
	return *o.IsEnabledPublicNetwork
}

// GetIsEnabledPublicNetworkOk returns a tuple with the IsEnabledPublicNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateClusterV2) GetIsEnabledPublicNetworkOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabledPublicNetwork) {
		return nil, false
	}
	return o.IsEnabledPublicNetwork, true
}

// HasIsEnabledPublicNetwork returns a boolean if a field has been set.
func (o *UpdateClusterV2) HasIsEnabledPublicNetwork() bool {
	if o != nil && !IsNil(o.IsEnabledPublicNetwork) {
		return true
	}

	return false
}

// SetIsEnabledPublicNetwork gets a reference to the given bool and assigns it to the IsEnabledPublicNetwork field.
func (o *UpdateClusterV2) SetIsEnabledPublicNetwork(v bool) {
	o.IsEnabledPublicNetwork = &v
}

// GetIsEnabledPublicIpv6 returns the IsEnabledPublicIpv6 field value if set, zero value otherwise.
func (o *UpdateClusterV2) GetIsEnabledPublicIpv6() bool {
	if o == nil || IsNil(o.IsEnabledPublicIpv6) {
		var ret bool
		return ret
	}
	return *o.IsEnabledPublicIpv6
}

// GetIsEnabledPublicIpv6Ok returns a tuple with the IsEnabledPublicIpv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateClusterV2) GetIsEnabledPublicIpv6Ok() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabledPublicIpv6) {
		return nil, false
	}
	return o.IsEnabledPublicIpv6, true
}

// HasIsEnabledPublicIpv6 returns a boolean if a field has been set.
func (o *UpdateClusterV2) HasIsEnabledPublicIpv6() bool {
	if o != nil && !IsNil(o.IsEnabledPublicIpv6) {
		return true
	}

	return false
}

// SetIsEnabledPublicIpv6 gets a reference to the given bool and assigns it to the IsEnabledPublicIpv6 field.
func (o *UpdateClusterV2) SetIsEnabledPublicIpv6(v bool) {
	o.IsEnabledPublicIpv6 = &v
}

// GetFloatingIp returns the FloatingIp field value if set, zero value otherwise.
func (o *UpdateClusterV2) GetFloatingIp() string {
	if o == nil || IsNil(o.FloatingIp) {
		var ret string
		return ret
	}
	return *o.FloatingIp
}

// GetFloatingIpOk returns a tuple with the FloatingIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateClusterV2) GetFloatingIpOk() (*string, bool) {
	if o == nil || IsNil(o.FloatingIp) {
		return nil, false
	}
	return o.FloatingIp, true
}

// HasFloatingIp returns a boolean if a field has been set.
func (o *UpdateClusterV2) HasFloatingIp() bool {
	if o != nil && !IsNil(o.FloatingIp) {
		return true
	}

	return false
}

// SetFloatingIp gets a reference to the given string and assigns it to the FloatingIp field.
func (o *UpdateClusterV2) SetFloatingIp(v string) {
	o.FloatingIp = &v
}

// GetIsSecureConnectionEnable returns the IsSecureConnectionEnable field value if set, zero value otherwise.
func (o *UpdateClusterV2) GetIsSecureConnectionEnable() bool {
	if o == nil || IsNil(o.IsSecureConnectionEnable) {
		var ret bool
		return ret
	}
	return *o.IsSecureConnectionEnable
}

// GetIsSecureConnectionEnableOk returns a tuple with the IsSecureConnectionEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateClusterV2) GetIsSecureConnectionEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSecureConnectionEnable) {
		return nil, false
	}
	return o.IsSecureConnectionEnable, true
}

// HasIsSecureConnectionEnable returns a boolean if a field has been set.
func (o *UpdateClusterV2) HasIsSecureConnectionEnable() bool {
	if o != nil && !IsNil(o.IsSecureConnectionEnable) {
		return true
	}

	return false
}

// SetIsSecureConnectionEnable gets a reference to the given bool and assigns it to the IsSecureConnectionEnable field.
func (o *UpdateClusterV2) SetIsSecureConnectionEnable(v bool) {
	o.IsSecureConnectionEnable = &v
}

// GetMaintenanceSlot returns the MaintenanceSlot field value if set, zero value otherwise.
func (o *UpdateClusterV2) GetMaintenanceSlot() UpdateClusterV2MaintenanceSlot {
	if o == nil || IsNil(o.MaintenanceSlot) {
		var ret UpdateClusterV2MaintenanceSlot
		return ret
	}
	return *o.MaintenanceSlot
}

// GetMaintenanceSlotOk returns a tuple with the MaintenanceSlot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateClusterV2) GetMaintenanceSlotOk() (*UpdateClusterV2MaintenanceSlot, bool) {
	if o == nil || IsNil(o.MaintenanceSlot) {
		return nil, false
	}
	return o.MaintenanceSlot, true
}

// HasMaintenanceSlot returns a boolean if a field has been set.
func (o *UpdateClusterV2) HasMaintenanceSlot() bool {
	if o != nil && !IsNil(o.MaintenanceSlot) {
		return true
	}

	return false
}

// SetMaintenanceSlot gets a reference to the given UpdateClusterV2MaintenanceSlot and assigns it to the MaintenanceSlot field.
func (o *UpdateClusterV2) SetMaintenanceSlot(v UpdateClusterV2MaintenanceSlot) {
	o.MaintenanceSlot = &v
}

// GetDiskAutoscaling returns the DiskAutoscaling field value if set, zero value otherwise.
func (o *UpdateClusterV2) GetDiskAutoscaling() UpdateClusterV2DiskAutoscaling {
	if o == nil || IsNil(o.DiskAutoscaling) {
		var ret UpdateClusterV2DiskAutoscaling
		return ret
	}
	return *o.DiskAutoscaling
}

// GetDiskAutoscalingOk returns a tuple with the DiskAutoscaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateClusterV2) GetDiskAutoscalingOk() (*UpdateClusterV2DiskAutoscaling, bool) {
	if o == nil || IsNil(o.DiskAutoscaling) {
		return nil, false
	}
	return o.DiskAutoscaling, true
}

// HasDiskAutoscaling returns a boolean if a field has been set.
func (o *UpdateClusterV2) HasDiskAutoscaling() bool {
	if o != nil && !IsNil(o.DiskAutoscaling) {
		return true
	}

	return false
}

// SetDiskAutoscaling gets a reference to the given UpdateClusterV2DiskAutoscaling and assigns it to the DiskAutoscaling field.
func (o *UpdateClusterV2) SetDiskAutoscaling(v UpdateClusterV2DiskAutoscaling) {
	o.DiskAutoscaling = &v
}

func (o UpdateClusterV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateClusterV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PresetId) {
		toSerialize["preset_id"] = o.PresetId
	}
	if !IsNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}
	if !IsNil(o.ConfigParameters) {
		toSerialize["config_parameters"] = o.ConfigParameters
	}
	if !IsNil(o.HashType) {
		toSerialize["hash_type"] = o.HashType
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.IsEnabledPublicNetwork) {
		toSerialize["is_enabled_public_network"] = o.IsEnabledPublicNetwork
	}
	if !IsNil(o.IsEnabledPublicIpv6) {
		toSerialize["is_enabled_public_ipv6"] = o.IsEnabledPublicIpv6
	}
	if !IsNil(o.FloatingIp) {
		toSerialize["floating_ip"] = o.FloatingIp
	}
	if !IsNil(o.IsSecureConnectionEnable) {
		toSerialize["is_secure_connection_enable"] = o.IsSecureConnectionEnable
	}
	if !IsNil(o.MaintenanceSlot) {
		toSerialize["maintenance_slot"] = o.MaintenanceSlot
	}
	if !IsNil(o.DiskAutoscaling) {
		toSerialize["disk_autoscaling"] = o.DiskAutoscaling
	}
	return toSerialize, nil
}

type NullableUpdateClusterV2 struct {
	value *UpdateClusterV2
	isSet bool
}

func (v NullableUpdateClusterV2) Get() *UpdateClusterV2 {
	return v.value
}

func (v *NullableUpdateClusterV2) Set(val *UpdateClusterV2) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateClusterV2) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateClusterV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateClusterV2(val *UpdateClusterV2) *NullableUpdateClusterV2 {
	return &NullableUpdateClusterV2{value: val, isSet: true}
}

func (v NullableUpdateClusterV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateClusterV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


