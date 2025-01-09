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

// checks if the CreateBalancer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateBalancer{}

// CreateBalancer struct for CreateBalancer
type CreateBalancer struct {
	// Удобочитаемое имя, установленное для балансировщика.
	Name string `json:"name"`
	// Алгоритм переключений балансировщика.
	Algo string `json:"algo"`
	// Это логическое значение, которое показывает, сохраняется ли сессия.
	IsSticky bool `json:"is_sticky"`
	// Это логическое значение, которое показывает, выступает ли балансировщик в качестве прокси.
	IsUseProxy bool `json:"is_use_proxy"`
	// Это логическое значение, которое показывает, требуется ли перенаправление на SSL.
	IsSsl bool `json:"is_ssl"`
	// Это логическое значение, которое показывает, выдает ли балансировщик сигнал о проверке жизнеспособности.
	IsKeepalive bool `json:"is_keepalive"`
	// Протокол.
	Proto string `json:"proto"`
	// Порт балансировщика.
	Port float32 `json:"port"`
	// Адрес балансировщика.
	Path string `json:"path"`
	// Интервал проверки.
	Inter float32 `json:"inter"`
	// Таймаут ответа балансировщика.
	Timeout float32 `json:"timeout"`
	// Порог количества ошибок.
	Fall float32 `json:"fall"`
	// Порог количества успешных ответов.
	Rise float32 `json:"rise"`
	// ID тарифа.
	PresetId float32 `json:"preset_id"`
	Network *Network `json:"network,omitempty"`
	AvailabilityZone *AvailabilityZone `json:"availability_zone,omitempty"`
}

// NewCreateBalancer instantiates a new CreateBalancer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateBalancer(name string, algo string, isSticky bool, isUseProxy bool, isSsl bool, isKeepalive bool, proto string, port float32, path string, inter float32, timeout float32, fall float32, rise float32, presetId float32) *CreateBalancer {
	this := CreateBalancer{}
	this.Name = name
	this.Algo = algo
	this.IsSticky = isSticky
	this.IsUseProxy = isUseProxy
	this.IsSsl = isSsl
	this.IsKeepalive = isKeepalive
	this.Proto = proto
	this.Port = port
	this.Path = path
	this.Inter = inter
	this.Timeout = timeout
	this.Fall = fall
	this.Rise = rise
	this.PresetId = presetId
	return &this
}

// NewCreateBalancerWithDefaults instantiates a new CreateBalancer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateBalancerWithDefaults() *CreateBalancer {
	this := CreateBalancer{}
	return &this
}

// GetName returns the Name field value
func (o *CreateBalancer) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateBalancer) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateBalancer) SetName(v string) {
	o.Name = v
}

// GetAlgo returns the Algo field value
func (o *CreateBalancer) GetAlgo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Algo
}

// GetAlgoOk returns a tuple with the Algo field value
// and a boolean to check if the value has been set.
func (o *CreateBalancer) GetAlgoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Algo, true
}

// SetAlgo sets field value
func (o *CreateBalancer) SetAlgo(v string) {
	o.Algo = v
}

// GetIsSticky returns the IsSticky field value
func (o *CreateBalancer) GetIsSticky() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSticky
}

// GetIsStickyOk returns a tuple with the IsSticky field value
// and a boolean to check if the value has been set.
func (o *CreateBalancer) GetIsStickyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSticky, true
}

// SetIsSticky sets field value
func (o *CreateBalancer) SetIsSticky(v bool) {
	o.IsSticky = v
}

// GetIsUseProxy returns the IsUseProxy field value
func (o *CreateBalancer) GetIsUseProxy() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsUseProxy
}

// GetIsUseProxyOk returns a tuple with the IsUseProxy field value
// and a boolean to check if the value has been set.
func (o *CreateBalancer) GetIsUseProxyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsUseProxy, true
}

// SetIsUseProxy sets field value
func (o *CreateBalancer) SetIsUseProxy(v bool) {
	o.IsUseProxy = v
}

// GetIsSsl returns the IsSsl field value
func (o *CreateBalancer) GetIsSsl() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSsl
}

// GetIsSslOk returns a tuple with the IsSsl field value
// and a boolean to check if the value has been set.
func (o *CreateBalancer) GetIsSslOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSsl, true
}

// SetIsSsl sets field value
func (o *CreateBalancer) SetIsSsl(v bool) {
	o.IsSsl = v
}

// GetIsKeepalive returns the IsKeepalive field value
func (o *CreateBalancer) GetIsKeepalive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsKeepalive
}

// GetIsKeepaliveOk returns a tuple with the IsKeepalive field value
// and a boolean to check if the value has been set.
func (o *CreateBalancer) GetIsKeepaliveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsKeepalive, true
}

// SetIsKeepalive sets field value
func (o *CreateBalancer) SetIsKeepalive(v bool) {
	o.IsKeepalive = v
}

// GetProto returns the Proto field value
func (o *CreateBalancer) GetProto() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Proto
}

// GetProtoOk returns a tuple with the Proto field value
// and a boolean to check if the value has been set.
func (o *CreateBalancer) GetProtoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Proto, true
}

// SetProto sets field value
func (o *CreateBalancer) SetProto(v string) {
	o.Proto = v
}

// GetPort returns the Port field value
func (o *CreateBalancer) GetPort() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *CreateBalancer) GetPortOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *CreateBalancer) SetPort(v float32) {
	o.Port = v
}

// GetPath returns the Path field value
func (o *CreateBalancer) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *CreateBalancer) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *CreateBalancer) SetPath(v string) {
	o.Path = v
}

// GetInter returns the Inter field value
func (o *CreateBalancer) GetInter() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Inter
}

// GetInterOk returns a tuple with the Inter field value
// and a boolean to check if the value has been set.
func (o *CreateBalancer) GetInterOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inter, true
}

// SetInter sets field value
func (o *CreateBalancer) SetInter(v float32) {
	o.Inter = v
}

// GetTimeout returns the Timeout field value
func (o *CreateBalancer) GetTimeout() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value
// and a boolean to check if the value has been set.
func (o *CreateBalancer) GetTimeoutOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timeout, true
}

// SetTimeout sets field value
func (o *CreateBalancer) SetTimeout(v float32) {
	o.Timeout = v
}

// GetFall returns the Fall field value
func (o *CreateBalancer) GetFall() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Fall
}

// GetFallOk returns a tuple with the Fall field value
// and a boolean to check if the value has been set.
func (o *CreateBalancer) GetFallOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fall, true
}

// SetFall sets field value
func (o *CreateBalancer) SetFall(v float32) {
	o.Fall = v
}

// GetRise returns the Rise field value
func (o *CreateBalancer) GetRise() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Rise
}

// GetRiseOk returns a tuple with the Rise field value
// and a boolean to check if the value has been set.
func (o *CreateBalancer) GetRiseOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rise, true
}

// SetRise sets field value
func (o *CreateBalancer) SetRise(v float32) {
	o.Rise = v
}

// GetPresetId returns the PresetId field value
func (o *CreateBalancer) GetPresetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PresetId
}

// GetPresetIdOk returns a tuple with the PresetId field value
// and a boolean to check if the value has been set.
func (o *CreateBalancer) GetPresetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PresetId, true
}

// SetPresetId sets field value
func (o *CreateBalancer) SetPresetId(v float32) {
	o.PresetId = v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *CreateBalancer) GetNetwork() Network {
	if o == nil || IsNil(o.Network) {
		var ret Network
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBalancer) GetNetworkOk() (*Network, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *CreateBalancer) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given Network and assigns it to the Network field.
func (o *CreateBalancer) SetNetwork(v Network) {
	o.Network = &v
}

// GetAvailabilityZone returns the AvailabilityZone field value if set, zero value otherwise.
func (o *CreateBalancer) GetAvailabilityZone() AvailabilityZone {
	if o == nil || IsNil(o.AvailabilityZone) {
		var ret AvailabilityZone
		return ret
	}
	return *o.AvailabilityZone
}

// GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBalancer) GetAvailabilityZoneOk() (*AvailabilityZone, bool) {
	if o == nil || IsNil(o.AvailabilityZone) {
		return nil, false
	}
	return o.AvailabilityZone, true
}

// HasAvailabilityZone returns a boolean if a field has been set.
func (o *CreateBalancer) HasAvailabilityZone() bool {
	if o != nil && !IsNil(o.AvailabilityZone) {
		return true
	}

	return false
}

// SetAvailabilityZone gets a reference to the given AvailabilityZone and assigns it to the AvailabilityZone field.
func (o *CreateBalancer) SetAvailabilityZone(v AvailabilityZone) {
	o.AvailabilityZone = &v
}

func (o CreateBalancer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateBalancer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["algo"] = o.Algo
	toSerialize["is_sticky"] = o.IsSticky
	toSerialize["is_use_proxy"] = o.IsUseProxy
	toSerialize["is_ssl"] = o.IsSsl
	toSerialize["is_keepalive"] = o.IsKeepalive
	toSerialize["proto"] = o.Proto
	toSerialize["port"] = o.Port
	toSerialize["path"] = o.Path
	toSerialize["inter"] = o.Inter
	toSerialize["timeout"] = o.Timeout
	toSerialize["fall"] = o.Fall
	toSerialize["rise"] = o.Rise
	toSerialize["preset_id"] = o.PresetId
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.AvailabilityZone) {
		toSerialize["availability_zone"] = o.AvailabilityZone
	}
	return toSerialize, nil
}

type NullableCreateBalancer struct {
	value *CreateBalancer
	isSet bool
}

func (v NullableCreateBalancer) Get() *CreateBalancer {
	return v.value
}

func (v *NullableCreateBalancer) Set(val *CreateBalancer) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateBalancer) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateBalancer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateBalancer(val *CreateBalancer) *NullableCreateBalancer {
	return &NullableCreateBalancer{value: val, isSet: true}
}

func (v NullableCreateBalancer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateBalancer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


