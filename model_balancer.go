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

// checks if the Balancer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Balancer{}

// Balancer Балансировщик
type Balancer struct {
	// ID для каждого экземпляра балансировщика. Автоматически генерируется при создании.
	Id float32 `json:"id"`
	// Алгоритм переключений балансировщика.
	Algo string `json:"algo"`
	// Значение времени, указанное в комбинированном формате даты и времени ISO8601, которое представляет, когда был создан балансировщик.
	CreatedAt time.Time `json:"created_at"`
	// Порог количества ошибок.
	Fall float32 `json:"fall"`
	// Интервал проверки.
	Inter float32 `json:"inter"`
	// IP-адрес сетевого интерфейса IPv4.
	Ip NullableString `json:"ip"`
	// Локальный IP-адрес сетевого интерфейса IPv4.
	LocalIp NullableString `json:"local_ip"`
	// Это логическое значение, которое показывает, выдает ли балансировщик сигнал о проверке жизнеспособности.
	IsKeepalive bool `json:"is_keepalive"`
	// Удобочитаемое имя, установленное для балансировщика.
	Name string `json:"name"`
	// Адрес балансировщика.
	Path string `json:"path"`
	// Порт балансировщика.
	Port float32 `json:"port"`
	// Протокол.
	Proto string `json:"proto"`
	// Порог количества успешных ответов.
	Rise float32 `json:"rise"`
	// ID тарифа.
	PresetId float32 `json:"preset_id"`
	// Это логическое значение, которое показывает, требуется ли перенаправление на SSL.
	IsSsl bool `json:"is_ssl"`
	// Статус балансировщика.
	Status string `json:"status"`
	// Это логическое значение, которое показывает, сохраняется ли сессия.
	IsSticky bool `json:"is_sticky"`
	// Таймаут ответа балансировщика.
	Timeout float32 `json:"timeout"`
	// Это логическое значение, которое показывает, выступает ли балансировщик в качестве прокси.
	IsUseProxy bool `json:"is_use_proxy"`
	Rules []Rule `json:"rules"`
	// Список IP-адресов, привязанных к балансировщику
	Ips []string `json:"ips"`
	// Географическое расположение балансировщика
	Location string `json:"location"`
	AvailabilityZone AvailabilityZone `json:"availability_zone"`
}

// NewBalancer instantiates a new Balancer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalancer(id float32, algo string, createdAt time.Time, fall float32, inter float32, ip NullableString, localIp NullableString, isKeepalive bool, name string, path string, port float32, proto string, rise float32, presetId float32, isSsl bool, status string, isSticky bool, timeout float32, isUseProxy bool, rules []Rule, ips []string, location string, availabilityZone AvailabilityZone) *Balancer {
	this := Balancer{}
	this.Id = id
	this.Algo = algo
	this.CreatedAt = createdAt
	this.Fall = fall
	this.Inter = inter
	this.Ip = ip
	this.LocalIp = localIp
	this.IsKeepalive = isKeepalive
	this.Name = name
	this.Path = path
	this.Port = port
	this.Proto = proto
	this.Rise = rise
	this.PresetId = presetId
	this.IsSsl = isSsl
	this.Status = status
	this.IsSticky = isSticky
	this.Timeout = timeout
	this.IsUseProxy = isUseProxy
	this.Rules = rules
	this.Ips = ips
	this.Location = location
	this.AvailabilityZone = availabilityZone
	return &this
}

// NewBalancerWithDefaults instantiates a new Balancer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalancerWithDefaults() *Balancer {
	this := Balancer{}
	return &this
}

// GetId returns the Id field value
func (o *Balancer) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Balancer) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Balancer) SetId(v float32) {
	o.Id = v
}

// GetAlgo returns the Algo field value
func (o *Balancer) GetAlgo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Algo
}

// GetAlgoOk returns a tuple with the Algo field value
// and a boolean to check if the value has been set.
func (o *Balancer) GetAlgoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Algo, true
}

// SetAlgo sets field value
func (o *Balancer) SetAlgo(v string) {
	o.Algo = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Balancer) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Balancer) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Balancer) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetFall returns the Fall field value
func (o *Balancer) GetFall() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Fall
}

// GetFallOk returns a tuple with the Fall field value
// and a boolean to check if the value has been set.
func (o *Balancer) GetFallOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fall, true
}

// SetFall sets field value
func (o *Balancer) SetFall(v float32) {
	o.Fall = v
}

// GetInter returns the Inter field value
func (o *Balancer) GetInter() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Inter
}

// GetInterOk returns a tuple with the Inter field value
// and a boolean to check if the value has been set.
func (o *Balancer) GetInterOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inter, true
}

// SetInter sets field value
func (o *Balancer) SetInter(v float32) {
	o.Inter = v
}

// GetIp returns the Ip field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Balancer) GetIp() string {
	if o == nil || o.Ip.Get() == nil {
		var ret string
		return ret
	}

	return *o.Ip.Get()
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Balancer) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ip.Get(), o.Ip.IsSet()
}

// SetIp sets field value
func (o *Balancer) SetIp(v string) {
	o.Ip.Set(&v)
}

// GetLocalIp returns the LocalIp field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Balancer) GetLocalIp() string {
	if o == nil || o.LocalIp.Get() == nil {
		var ret string
		return ret
	}

	return *o.LocalIp.Get()
}

// GetLocalIpOk returns a tuple with the LocalIp field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Balancer) GetLocalIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LocalIp.Get(), o.LocalIp.IsSet()
}

// SetLocalIp sets field value
func (o *Balancer) SetLocalIp(v string) {
	o.LocalIp.Set(&v)
}

// GetIsKeepalive returns the IsKeepalive field value
func (o *Balancer) GetIsKeepalive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsKeepalive
}

// GetIsKeepaliveOk returns a tuple with the IsKeepalive field value
// and a boolean to check if the value has been set.
func (o *Balancer) GetIsKeepaliveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsKeepalive, true
}

// SetIsKeepalive sets field value
func (o *Balancer) SetIsKeepalive(v bool) {
	o.IsKeepalive = v
}

// GetName returns the Name field value
func (o *Balancer) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Balancer) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Balancer) SetName(v string) {
	o.Name = v
}

// GetPath returns the Path field value
func (o *Balancer) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *Balancer) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *Balancer) SetPath(v string) {
	o.Path = v
}

// GetPort returns the Port field value
func (o *Balancer) GetPort() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *Balancer) GetPortOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *Balancer) SetPort(v float32) {
	o.Port = v
}

// GetProto returns the Proto field value
func (o *Balancer) GetProto() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Proto
}

// GetProtoOk returns a tuple with the Proto field value
// and a boolean to check if the value has been set.
func (o *Balancer) GetProtoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Proto, true
}

// SetProto sets field value
func (o *Balancer) SetProto(v string) {
	o.Proto = v
}

// GetRise returns the Rise field value
func (o *Balancer) GetRise() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Rise
}

// GetRiseOk returns a tuple with the Rise field value
// and a boolean to check if the value has been set.
func (o *Balancer) GetRiseOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rise, true
}

// SetRise sets field value
func (o *Balancer) SetRise(v float32) {
	o.Rise = v
}

// GetPresetId returns the PresetId field value
func (o *Balancer) GetPresetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PresetId
}

// GetPresetIdOk returns a tuple with the PresetId field value
// and a boolean to check if the value has been set.
func (o *Balancer) GetPresetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PresetId, true
}

// SetPresetId sets field value
func (o *Balancer) SetPresetId(v float32) {
	o.PresetId = v
}

// GetIsSsl returns the IsSsl field value
func (o *Balancer) GetIsSsl() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSsl
}

// GetIsSslOk returns a tuple with the IsSsl field value
// and a boolean to check if the value has been set.
func (o *Balancer) GetIsSslOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSsl, true
}

// SetIsSsl sets field value
func (o *Balancer) SetIsSsl(v bool) {
	o.IsSsl = v
}

// GetStatus returns the Status field value
func (o *Balancer) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Balancer) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Balancer) SetStatus(v string) {
	o.Status = v
}

// GetIsSticky returns the IsSticky field value
func (o *Balancer) GetIsSticky() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSticky
}

// GetIsStickyOk returns a tuple with the IsSticky field value
// and a boolean to check if the value has been set.
func (o *Balancer) GetIsStickyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSticky, true
}

// SetIsSticky sets field value
func (o *Balancer) SetIsSticky(v bool) {
	o.IsSticky = v
}

// GetTimeout returns the Timeout field value
func (o *Balancer) GetTimeout() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value
// and a boolean to check if the value has been set.
func (o *Balancer) GetTimeoutOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timeout, true
}

// SetTimeout sets field value
func (o *Balancer) SetTimeout(v float32) {
	o.Timeout = v
}

// GetIsUseProxy returns the IsUseProxy field value
func (o *Balancer) GetIsUseProxy() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsUseProxy
}

// GetIsUseProxyOk returns a tuple with the IsUseProxy field value
// and a boolean to check if the value has been set.
func (o *Balancer) GetIsUseProxyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsUseProxy, true
}

// SetIsUseProxy sets field value
func (o *Balancer) SetIsUseProxy(v bool) {
	o.IsUseProxy = v
}

// GetRules returns the Rules field value
func (o *Balancer) GetRules() []Rule {
	if o == nil {
		var ret []Rule
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *Balancer) GetRulesOk() ([]Rule, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rules, true
}

// SetRules sets field value
func (o *Balancer) SetRules(v []Rule) {
	o.Rules = v
}

// GetIps returns the Ips field value
func (o *Balancer) GetIps() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Ips
}

// GetIpsOk returns a tuple with the Ips field value
// and a boolean to check if the value has been set.
func (o *Balancer) GetIpsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ips, true
}

// SetIps sets field value
func (o *Balancer) SetIps(v []string) {
	o.Ips = v
}

// GetLocation returns the Location field value
func (o *Balancer) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *Balancer) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *Balancer) SetLocation(v string) {
	o.Location = v
}

// GetAvailabilityZone returns the AvailabilityZone field value
func (o *Balancer) GetAvailabilityZone() AvailabilityZone {
	if o == nil {
		var ret AvailabilityZone
		return ret
	}

	return o.AvailabilityZone
}

// GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field value
// and a boolean to check if the value has been set.
func (o *Balancer) GetAvailabilityZoneOk() (*AvailabilityZone, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvailabilityZone, true
}

// SetAvailabilityZone sets field value
func (o *Balancer) SetAvailabilityZone(v AvailabilityZone) {
	o.AvailabilityZone = v
}

func (o Balancer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Balancer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["algo"] = o.Algo
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["fall"] = o.Fall
	toSerialize["inter"] = o.Inter
	toSerialize["ip"] = o.Ip.Get()
	toSerialize["local_ip"] = o.LocalIp.Get()
	toSerialize["is_keepalive"] = o.IsKeepalive
	toSerialize["name"] = o.Name
	toSerialize["path"] = o.Path
	toSerialize["port"] = o.Port
	toSerialize["proto"] = o.Proto
	toSerialize["rise"] = o.Rise
	toSerialize["preset_id"] = o.PresetId
	toSerialize["is_ssl"] = o.IsSsl
	toSerialize["status"] = o.Status
	toSerialize["is_sticky"] = o.IsSticky
	toSerialize["timeout"] = o.Timeout
	toSerialize["is_use_proxy"] = o.IsUseProxy
	toSerialize["rules"] = o.Rules
	toSerialize["ips"] = o.Ips
	toSerialize["location"] = o.Location
	toSerialize["availability_zone"] = o.AvailabilityZone
	return toSerialize, nil
}

type NullableBalancer struct {
	value *Balancer
	isSet bool
}

func (v NullableBalancer) Get() *Balancer {
	return v.value
}

func (v *NullableBalancer) Set(val *Balancer) {
	v.value = val
	v.isSet = true
}

func (v NullableBalancer) IsSet() bool {
	return v.isSet
}

func (v *NullableBalancer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalancer(val *Balancer) *NullableBalancer {
	return &NullableBalancer{value: val, isSet: true}
}

func (v NullableBalancer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalancer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


