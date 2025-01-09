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

// checks if the NodeGroupIn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NodeGroupIn{}

// NodeGroupIn struct for NodeGroupIn
type NodeGroupIn struct {
	// Название группы
	Name string `json:"name"`
	// ID тарифа воркер-ноды. Нельзя передавать вместе с `configuration`. Локация воркер-нод должна совпадать с локацией кластера
	PresetId *int32 `json:"preset_id,omitempty"`
	Configuration *NodeGroupInConfiguration `json:"configuration,omitempty"`
	// Количество нод в группе
	NodeCount int32 `json:"node_count"`
	// Лейблы для группы нод
	Labels []SetLabels `json:"labels,omitempty"`
	// Автомасштабирование. Автоматическое увеличение и уменьшение количества нод в группе в зависимости от текущей нагрузки
	IsAutoscaling *bool `json:"is_autoscaling,omitempty"`
	// Минимальное количество нод. Передавать в связке с параметрами `is_autoscaling` и `max_size`
	MinSize *int32 `json:"min-size,omitempty"`
	// Максимальное количество нод. Передавать в связке с параметрами `is_autoscaling` и `min_size`. Максимальное количество нод ограничено тарифом кластера
	MaxSize *int32 `json:"max-size,omitempty"`
}

// NewNodeGroupIn instantiates a new NodeGroupIn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeGroupIn(name string, nodeCount int32) *NodeGroupIn {
	this := NodeGroupIn{}
	this.Name = name
	this.NodeCount = nodeCount
	return &this
}

// NewNodeGroupInWithDefaults instantiates a new NodeGroupIn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeGroupInWithDefaults() *NodeGroupIn {
	this := NodeGroupIn{}
	return &this
}

// GetName returns the Name field value
func (o *NodeGroupIn) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NodeGroupIn) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NodeGroupIn) SetName(v string) {
	o.Name = v
}

// GetPresetId returns the PresetId field value if set, zero value otherwise.
func (o *NodeGroupIn) GetPresetId() int32 {
	if o == nil || IsNil(o.PresetId) {
		var ret int32
		return ret
	}
	return *o.PresetId
}

// GetPresetIdOk returns a tuple with the PresetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeGroupIn) GetPresetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.PresetId) {
		return nil, false
	}
	return o.PresetId, true
}

// HasPresetId returns a boolean if a field has been set.
func (o *NodeGroupIn) HasPresetId() bool {
	if o != nil && !IsNil(o.PresetId) {
		return true
	}

	return false
}

// SetPresetId gets a reference to the given int32 and assigns it to the PresetId field.
func (o *NodeGroupIn) SetPresetId(v int32) {
	o.PresetId = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *NodeGroupIn) GetConfiguration() NodeGroupInConfiguration {
	if o == nil || IsNil(o.Configuration) {
		var ret NodeGroupInConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeGroupIn) GetConfigurationOk() (*NodeGroupInConfiguration, bool) {
	if o == nil || IsNil(o.Configuration) {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *NodeGroupIn) HasConfiguration() bool {
	if o != nil && !IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given NodeGroupInConfiguration and assigns it to the Configuration field.
func (o *NodeGroupIn) SetConfiguration(v NodeGroupInConfiguration) {
	o.Configuration = &v
}

// GetNodeCount returns the NodeCount field value
func (o *NodeGroupIn) GetNodeCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NodeCount
}

// GetNodeCountOk returns a tuple with the NodeCount field value
// and a boolean to check if the value has been set.
func (o *NodeGroupIn) GetNodeCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeCount, true
}

// SetNodeCount sets field value
func (o *NodeGroupIn) SetNodeCount(v int32) {
	o.NodeCount = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *NodeGroupIn) GetLabels() []SetLabels {
	if o == nil || IsNil(o.Labels) {
		var ret []SetLabels
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeGroupIn) GetLabelsOk() ([]SetLabels, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *NodeGroupIn) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []SetLabels and assigns it to the Labels field.
func (o *NodeGroupIn) SetLabels(v []SetLabels) {
	o.Labels = v
}

// GetIsAutoscaling returns the IsAutoscaling field value if set, zero value otherwise.
func (o *NodeGroupIn) GetIsAutoscaling() bool {
	if o == nil || IsNil(o.IsAutoscaling) {
		var ret bool
		return ret
	}
	return *o.IsAutoscaling
}

// GetIsAutoscalingOk returns a tuple with the IsAutoscaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeGroupIn) GetIsAutoscalingOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAutoscaling) {
		return nil, false
	}
	return o.IsAutoscaling, true
}

// HasIsAutoscaling returns a boolean if a field has been set.
func (o *NodeGroupIn) HasIsAutoscaling() bool {
	if o != nil && !IsNil(o.IsAutoscaling) {
		return true
	}

	return false
}

// SetIsAutoscaling gets a reference to the given bool and assigns it to the IsAutoscaling field.
func (o *NodeGroupIn) SetIsAutoscaling(v bool) {
	o.IsAutoscaling = &v
}

// GetMinSize returns the MinSize field value if set, zero value otherwise.
func (o *NodeGroupIn) GetMinSize() int32 {
	if o == nil || IsNil(o.MinSize) {
		var ret int32
		return ret
	}
	return *o.MinSize
}

// GetMinSizeOk returns a tuple with the MinSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeGroupIn) GetMinSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.MinSize) {
		return nil, false
	}
	return o.MinSize, true
}

// HasMinSize returns a boolean if a field has been set.
func (o *NodeGroupIn) HasMinSize() bool {
	if o != nil && !IsNil(o.MinSize) {
		return true
	}

	return false
}

// SetMinSize gets a reference to the given int32 and assigns it to the MinSize field.
func (o *NodeGroupIn) SetMinSize(v int32) {
	o.MinSize = &v
}

// GetMaxSize returns the MaxSize field value if set, zero value otherwise.
func (o *NodeGroupIn) GetMaxSize() int32 {
	if o == nil || IsNil(o.MaxSize) {
		var ret int32
		return ret
	}
	return *o.MaxSize
}

// GetMaxSizeOk returns a tuple with the MaxSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeGroupIn) GetMaxSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxSize) {
		return nil, false
	}
	return o.MaxSize, true
}

// HasMaxSize returns a boolean if a field has been set.
func (o *NodeGroupIn) HasMaxSize() bool {
	if o != nil && !IsNil(o.MaxSize) {
		return true
	}

	return false
}

// SetMaxSize gets a reference to the given int32 and assigns it to the MaxSize field.
func (o *NodeGroupIn) SetMaxSize(v int32) {
	o.MaxSize = &v
}

func (o NodeGroupIn) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NodeGroupIn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.PresetId) {
		toSerialize["preset_id"] = o.PresetId
	}
	if !IsNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}
	toSerialize["node_count"] = o.NodeCount
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.IsAutoscaling) {
		toSerialize["is_autoscaling"] = o.IsAutoscaling
	}
	if !IsNil(o.MinSize) {
		toSerialize["min-size"] = o.MinSize
	}
	if !IsNil(o.MaxSize) {
		toSerialize["max-size"] = o.MaxSize
	}
	return toSerialize, nil
}

type NullableNodeGroupIn struct {
	value *NodeGroupIn
	isSet bool
}

func (v NullableNodeGroupIn) Get() *NodeGroupIn {
	return v.value
}

func (v *NullableNodeGroupIn) Set(val *NodeGroupIn) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeGroupIn) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeGroupIn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeGroupIn(val *NodeGroupIn) *NullableNodeGroupIn {
	return &NullableNodeGroupIn{value: val, isSet: true}
}

func (v NullableNodeGroupIn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeGroupIn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


