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

// checks if the UpdateBalancer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateBalancer{}

// UpdateBalancer struct for UpdateBalancer
type UpdateBalancer struct {
	// Удобочитаемое имя, установленное для балансировщика.
	Name *string `json:"name,omitempty"`
	// Алгоритм переключений балансировщика.
	Algo *string `json:"algo,omitempty"`
	// Это логическое значение, которое показывает, сохраняется ли сессия.
	IsSticky *bool `json:"is_sticky,omitempty"`
	// Это логическое значение, которое показывает, выступает ли балансировщик в качестве прокси.
	IsUseProxy *bool `json:"is_use_proxy,omitempty"`
	// Это логическое значение, которое показывает, требуется ли перенаправление на SSL.
	IsSsl *bool `json:"is_ssl,omitempty"`
	// Это логическое значение, которое показывает, выдает ли балансировщик сигнал о проверке жизнеспособности.
	IsKeepalive *bool `json:"is_keepalive,omitempty"`
	// Протокол.
	Proto *string `json:"proto,omitempty"`
	// Порт балансировщика.
	Port *float32 `json:"port,omitempty"`
	// Адрес балансировщика.
	Path *string `json:"path,omitempty"`
	// Интервал проверки.
	Inter *float32 `json:"inter,omitempty"`
	// Таймаут ответа балансировщика.
	Timeout *float32 `json:"timeout,omitempty"`
	// Порог количества ошибок.
	Fall *float32 `json:"fall,omitempty"`
	// Порог количества успешных ответов.
	Rise *float32 `json:"rise,omitempty"`
}

// NewUpdateBalancer instantiates a new UpdateBalancer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateBalancer() *UpdateBalancer {
	this := UpdateBalancer{}
	return &this
}

// NewUpdateBalancerWithDefaults instantiates a new UpdateBalancer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateBalancerWithDefaults() *UpdateBalancer {
	this := UpdateBalancer{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateBalancer) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBalancer) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateBalancer) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateBalancer) SetName(v string) {
	o.Name = &v
}

// GetAlgo returns the Algo field value if set, zero value otherwise.
func (o *UpdateBalancer) GetAlgo() string {
	if o == nil || IsNil(o.Algo) {
		var ret string
		return ret
	}
	return *o.Algo
}

// GetAlgoOk returns a tuple with the Algo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBalancer) GetAlgoOk() (*string, bool) {
	if o == nil || IsNil(o.Algo) {
		return nil, false
	}
	return o.Algo, true
}

// HasAlgo returns a boolean if a field has been set.
func (o *UpdateBalancer) HasAlgo() bool {
	if o != nil && !IsNil(o.Algo) {
		return true
	}

	return false
}

// SetAlgo gets a reference to the given string and assigns it to the Algo field.
func (o *UpdateBalancer) SetAlgo(v string) {
	o.Algo = &v
}

// GetIsSticky returns the IsSticky field value if set, zero value otherwise.
func (o *UpdateBalancer) GetIsSticky() bool {
	if o == nil || IsNil(o.IsSticky) {
		var ret bool
		return ret
	}
	return *o.IsSticky
}

// GetIsStickyOk returns a tuple with the IsSticky field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBalancer) GetIsStickyOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSticky) {
		return nil, false
	}
	return o.IsSticky, true
}

// HasIsSticky returns a boolean if a field has been set.
func (o *UpdateBalancer) HasIsSticky() bool {
	if o != nil && !IsNil(o.IsSticky) {
		return true
	}

	return false
}

// SetIsSticky gets a reference to the given bool and assigns it to the IsSticky field.
func (o *UpdateBalancer) SetIsSticky(v bool) {
	o.IsSticky = &v
}

// GetIsUseProxy returns the IsUseProxy field value if set, zero value otherwise.
func (o *UpdateBalancer) GetIsUseProxy() bool {
	if o == nil || IsNil(o.IsUseProxy) {
		var ret bool
		return ret
	}
	return *o.IsUseProxy
}

// GetIsUseProxyOk returns a tuple with the IsUseProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBalancer) GetIsUseProxyOk() (*bool, bool) {
	if o == nil || IsNil(o.IsUseProxy) {
		return nil, false
	}
	return o.IsUseProxy, true
}

// HasIsUseProxy returns a boolean if a field has been set.
func (o *UpdateBalancer) HasIsUseProxy() bool {
	if o != nil && !IsNil(o.IsUseProxy) {
		return true
	}

	return false
}

// SetIsUseProxy gets a reference to the given bool and assigns it to the IsUseProxy field.
func (o *UpdateBalancer) SetIsUseProxy(v bool) {
	o.IsUseProxy = &v
}

// GetIsSsl returns the IsSsl field value if set, zero value otherwise.
func (o *UpdateBalancer) GetIsSsl() bool {
	if o == nil || IsNil(o.IsSsl) {
		var ret bool
		return ret
	}
	return *o.IsSsl
}

// GetIsSslOk returns a tuple with the IsSsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBalancer) GetIsSslOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSsl) {
		return nil, false
	}
	return o.IsSsl, true
}

// HasIsSsl returns a boolean if a field has been set.
func (o *UpdateBalancer) HasIsSsl() bool {
	if o != nil && !IsNil(o.IsSsl) {
		return true
	}

	return false
}

// SetIsSsl gets a reference to the given bool and assigns it to the IsSsl field.
func (o *UpdateBalancer) SetIsSsl(v bool) {
	o.IsSsl = &v
}

// GetIsKeepalive returns the IsKeepalive field value if set, zero value otherwise.
func (o *UpdateBalancer) GetIsKeepalive() bool {
	if o == nil || IsNil(o.IsKeepalive) {
		var ret bool
		return ret
	}
	return *o.IsKeepalive
}

// GetIsKeepaliveOk returns a tuple with the IsKeepalive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBalancer) GetIsKeepaliveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsKeepalive) {
		return nil, false
	}
	return o.IsKeepalive, true
}

// HasIsKeepalive returns a boolean if a field has been set.
func (o *UpdateBalancer) HasIsKeepalive() bool {
	if o != nil && !IsNil(o.IsKeepalive) {
		return true
	}

	return false
}

// SetIsKeepalive gets a reference to the given bool and assigns it to the IsKeepalive field.
func (o *UpdateBalancer) SetIsKeepalive(v bool) {
	o.IsKeepalive = &v
}

// GetProto returns the Proto field value if set, zero value otherwise.
func (o *UpdateBalancer) GetProto() string {
	if o == nil || IsNil(o.Proto) {
		var ret string
		return ret
	}
	return *o.Proto
}

// GetProtoOk returns a tuple with the Proto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBalancer) GetProtoOk() (*string, bool) {
	if o == nil || IsNil(o.Proto) {
		return nil, false
	}
	return o.Proto, true
}

// HasProto returns a boolean if a field has been set.
func (o *UpdateBalancer) HasProto() bool {
	if o != nil && !IsNil(o.Proto) {
		return true
	}

	return false
}

// SetProto gets a reference to the given string and assigns it to the Proto field.
func (o *UpdateBalancer) SetProto(v string) {
	o.Proto = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *UpdateBalancer) GetPort() float32 {
	if o == nil || IsNil(o.Port) {
		var ret float32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBalancer) GetPortOk() (*float32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *UpdateBalancer) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given float32 and assigns it to the Port field.
func (o *UpdateBalancer) SetPort(v float32) {
	o.Port = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *UpdateBalancer) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBalancer) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *UpdateBalancer) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *UpdateBalancer) SetPath(v string) {
	o.Path = &v
}

// GetInter returns the Inter field value if set, zero value otherwise.
func (o *UpdateBalancer) GetInter() float32 {
	if o == nil || IsNil(o.Inter) {
		var ret float32
		return ret
	}
	return *o.Inter
}

// GetInterOk returns a tuple with the Inter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBalancer) GetInterOk() (*float32, bool) {
	if o == nil || IsNil(o.Inter) {
		return nil, false
	}
	return o.Inter, true
}

// HasInter returns a boolean if a field has been set.
func (o *UpdateBalancer) HasInter() bool {
	if o != nil && !IsNil(o.Inter) {
		return true
	}

	return false
}

// SetInter gets a reference to the given float32 and assigns it to the Inter field.
func (o *UpdateBalancer) SetInter(v float32) {
	o.Inter = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *UpdateBalancer) GetTimeout() float32 {
	if o == nil || IsNil(o.Timeout) {
		var ret float32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBalancer) GetTimeoutOk() (*float32, bool) {
	if o == nil || IsNil(o.Timeout) {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *UpdateBalancer) HasTimeout() bool {
	if o != nil && !IsNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given float32 and assigns it to the Timeout field.
func (o *UpdateBalancer) SetTimeout(v float32) {
	o.Timeout = &v
}

// GetFall returns the Fall field value if set, zero value otherwise.
func (o *UpdateBalancer) GetFall() float32 {
	if o == nil || IsNil(o.Fall) {
		var ret float32
		return ret
	}
	return *o.Fall
}

// GetFallOk returns a tuple with the Fall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBalancer) GetFallOk() (*float32, bool) {
	if o == nil || IsNil(o.Fall) {
		return nil, false
	}
	return o.Fall, true
}

// HasFall returns a boolean if a field has been set.
func (o *UpdateBalancer) HasFall() bool {
	if o != nil && !IsNil(o.Fall) {
		return true
	}

	return false
}

// SetFall gets a reference to the given float32 and assigns it to the Fall field.
func (o *UpdateBalancer) SetFall(v float32) {
	o.Fall = &v
}

// GetRise returns the Rise field value if set, zero value otherwise.
func (o *UpdateBalancer) GetRise() float32 {
	if o == nil || IsNil(o.Rise) {
		var ret float32
		return ret
	}
	return *o.Rise
}

// GetRiseOk returns a tuple with the Rise field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBalancer) GetRiseOk() (*float32, bool) {
	if o == nil || IsNil(o.Rise) {
		return nil, false
	}
	return o.Rise, true
}

// HasRise returns a boolean if a field has been set.
func (o *UpdateBalancer) HasRise() bool {
	if o != nil && !IsNil(o.Rise) {
		return true
	}

	return false
}

// SetRise gets a reference to the given float32 and assigns it to the Rise field.
func (o *UpdateBalancer) SetRise(v float32) {
	o.Rise = &v
}

func (o UpdateBalancer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateBalancer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Algo) {
		toSerialize["algo"] = o.Algo
	}
	if !IsNil(o.IsSticky) {
		toSerialize["is_sticky"] = o.IsSticky
	}
	if !IsNil(o.IsUseProxy) {
		toSerialize["is_use_proxy"] = o.IsUseProxy
	}
	if !IsNil(o.IsSsl) {
		toSerialize["is_ssl"] = o.IsSsl
	}
	if !IsNil(o.IsKeepalive) {
		toSerialize["is_keepalive"] = o.IsKeepalive
	}
	if !IsNil(o.Proto) {
		toSerialize["proto"] = o.Proto
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Inter) {
		toSerialize["inter"] = o.Inter
	}
	if !IsNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	if !IsNil(o.Fall) {
		toSerialize["fall"] = o.Fall
	}
	if !IsNil(o.Rise) {
		toSerialize["rise"] = o.Rise
	}
	return toSerialize, nil
}

type NullableUpdateBalancer struct {
	value *UpdateBalancer
	isSet bool
}

func (v NullableUpdateBalancer) Get() *UpdateBalancer {
	return v.value
}

func (v *NullableUpdateBalancer) Set(val *UpdateBalancer) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateBalancer) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateBalancer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateBalancer(val *UpdateBalancer) *NullableUpdateBalancer {
	return &NullableUpdateBalancer{value: val, isSet: true}
}

func (v NullableUpdateBalancer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateBalancer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


