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

// checks if the DedicatedServer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DedicatedServer{}

// DedicatedServer Выделенный сервер
type DedicatedServer struct {
	// ID для каждого экземпляра выделенного сервера. Автоматически генерируется при создании.
	Id float32 `json:"id"`
	// Описание параметров процессора выделенного сервера.
	CpuDescription string `json:"cpu_description"`
	// Описание параметров жёсткого диска выделенного сервера.
	HddDescription string `json:"hdd_description"`
	// Описание ОЗУ выделенного сервера.
	RamDescription string `json:"ram_description"`
	// Значение времени, указанное в комбинированном формате даты и времени ISO8601, которое представляет, когда был создан выделенный сервер.
	CreatedAt time.Time `json:"created_at"`
	// IP-адрес сетевого интерфейса IPv4.
	Ip NullableString `json:"ip"`
	// IP-адрес сетевого интерфейса IPMI.
	IpmiIp NullableString `json:"ipmi_ip"`
	// Логин, используемый для входа в IPMI-консоль.
	IpmiLogin NullableString `json:"ipmi_login"`
	// Пароль, используемый для входа в IPMI-консоль.
	IpmiPassword NullableString `json:"ipmi_password"`
	// IP-адрес сетевого интерфейса IPv6.
	Ipv6 NullableString `json:"ipv6"`
	// Внутренний дополнительный ID сервера.
	NodeId NullableFloat32 `json:"node_id"`
	// Удобочитаемое имя, установленное для выделенного сервера.
	Name string `json:"name"`
	// Комментарий к выделенному серверу.
	Comment string `json:"comment"`
	// Пароль для подключения к VNC-консоли выделенного сервера.
	VncPass NullableString `json:"vnc_pass"`
	// Строка состояния, указывающая состояние выделенного сервера. Может быть «installing», «installed», «on» или «off».
	Status string `json:"status"`
	// ID операционной системы, установленной на выделенный сервер.
	OsId NullableFloat32 `json:"os_id"`
	// ID панели управления, установленной на выделенный сервер.
	CpId NullableFloat32 `json:"cp_id"`
	// ID интернет-канала, установленного на выделенный сервер.
	BandwidthId NullableFloat32 `json:"bandwidth_id"`
	// Массив уникальных ID сетевых дисков, подключенных к выделенному серверу.
	NetworkDriveId []float32 `json:"network_drive_id"`
	// Массив уникальных ID дополнительных IP-адресов, подключенных к выделенному серверу.
	AdditionalIpAddrId []float32 `json:"additional_ip_addr_id"`
	// ID списка дополнительных услуг выделенного сервера.
	PlanId NullableFloat32 `json:"plan_id"`
	// Стоимость выделенного сервера.
	Price float32 `json:"price"`
	// Локация сервера.
	Location string `json:"location"`
	// Количество готовых к автоматической выдаче серверов. Если значение равно 0, сервер будет установлен через инженеров.
	AutoinstallReady float32 `json:"autoinstall_ready"`
	// Пароль root сервера или пароль Администратора для серверов Windows.
	Password NullableString `json:"password"`
}

// NewDedicatedServer instantiates a new DedicatedServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDedicatedServer(id float32, cpuDescription string, hddDescription string, ramDescription string, createdAt time.Time, ip NullableString, ipmiIp NullableString, ipmiLogin NullableString, ipmiPassword NullableString, ipv6 NullableString, nodeId NullableFloat32, name string, comment string, vncPass NullableString, status string, osId NullableFloat32, cpId NullableFloat32, bandwidthId NullableFloat32, networkDriveId []float32, additionalIpAddrId []float32, planId NullableFloat32, price float32, location string, autoinstallReady float32, password NullableString) *DedicatedServer {
	this := DedicatedServer{}
	this.Id = id
	this.CpuDescription = cpuDescription
	this.HddDescription = hddDescription
	this.RamDescription = ramDescription
	this.CreatedAt = createdAt
	this.Ip = ip
	this.IpmiIp = ipmiIp
	this.IpmiLogin = ipmiLogin
	this.IpmiPassword = ipmiPassword
	this.Ipv6 = ipv6
	this.NodeId = nodeId
	this.Name = name
	this.Comment = comment
	this.VncPass = vncPass
	this.Status = status
	this.OsId = osId
	this.CpId = cpId
	this.BandwidthId = bandwidthId
	this.NetworkDriveId = networkDriveId
	this.AdditionalIpAddrId = additionalIpAddrId
	this.PlanId = planId
	this.Price = price
	this.Location = location
	this.AutoinstallReady = autoinstallReady
	this.Password = password
	return &this
}

// NewDedicatedServerWithDefaults instantiates a new DedicatedServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDedicatedServerWithDefaults() *DedicatedServer {
	this := DedicatedServer{}
	return &this
}

// GetId returns the Id field value
func (o *DedicatedServer) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DedicatedServer) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DedicatedServer) SetId(v float32) {
	o.Id = v
}

// GetCpuDescription returns the CpuDescription field value
func (o *DedicatedServer) GetCpuDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CpuDescription
}

// GetCpuDescriptionOk returns a tuple with the CpuDescription field value
// and a boolean to check if the value has been set.
func (o *DedicatedServer) GetCpuDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CpuDescription, true
}

// SetCpuDescription sets field value
func (o *DedicatedServer) SetCpuDescription(v string) {
	o.CpuDescription = v
}

// GetHddDescription returns the HddDescription field value
func (o *DedicatedServer) GetHddDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HddDescription
}

// GetHddDescriptionOk returns a tuple with the HddDescription field value
// and a boolean to check if the value has been set.
func (o *DedicatedServer) GetHddDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HddDescription, true
}

// SetHddDescription sets field value
func (o *DedicatedServer) SetHddDescription(v string) {
	o.HddDescription = v
}

// GetRamDescription returns the RamDescription field value
func (o *DedicatedServer) GetRamDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RamDescription
}

// GetRamDescriptionOk returns a tuple with the RamDescription field value
// and a boolean to check if the value has been set.
func (o *DedicatedServer) GetRamDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RamDescription, true
}

// SetRamDescription sets field value
func (o *DedicatedServer) SetRamDescription(v string) {
	o.RamDescription = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *DedicatedServer) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DedicatedServer) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *DedicatedServer) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetIp returns the Ip field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DedicatedServer) GetIp() string {
	if o == nil || o.Ip.Get() == nil {
		var ret string
		return ret
	}

	return *o.Ip.Get()
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DedicatedServer) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ip.Get(), o.Ip.IsSet()
}

// SetIp sets field value
func (o *DedicatedServer) SetIp(v string) {
	o.Ip.Set(&v)
}

// GetIpmiIp returns the IpmiIp field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DedicatedServer) GetIpmiIp() string {
	if o == nil || o.IpmiIp.Get() == nil {
		var ret string
		return ret
	}

	return *o.IpmiIp.Get()
}

// GetIpmiIpOk returns a tuple with the IpmiIp field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DedicatedServer) GetIpmiIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpmiIp.Get(), o.IpmiIp.IsSet()
}

// SetIpmiIp sets field value
func (o *DedicatedServer) SetIpmiIp(v string) {
	o.IpmiIp.Set(&v)
}

// GetIpmiLogin returns the IpmiLogin field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DedicatedServer) GetIpmiLogin() string {
	if o == nil || o.IpmiLogin.Get() == nil {
		var ret string
		return ret
	}

	return *o.IpmiLogin.Get()
}

// GetIpmiLoginOk returns a tuple with the IpmiLogin field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DedicatedServer) GetIpmiLoginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpmiLogin.Get(), o.IpmiLogin.IsSet()
}

// SetIpmiLogin sets field value
func (o *DedicatedServer) SetIpmiLogin(v string) {
	o.IpmiLogin.Set(&v)
}

// GetIpmiPassword returns the IpmiPassword field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DedicatedServer) GetIpmiPassword() string {
	if o == nil || o.IpmiPassword.Get() == nil {
		var ret string
		return ret
	}

	return *o.IpmiPassword.Get()
}

// GetIpmiPasswordOk returns a tuple with the IpmiPassword field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DedicatedServer) GetIpmiPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpmiPassword.Get(), o.IpmiPassword.IsSet()
}

// SetIpmiPassword sets field value
func (o *DedicatedServer) SetIpmiPassword(v string) {
	o.IpmiPassword.Set(&v)
}

// GetIpv6 returns the Ipv6 field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DedicatedServer) GetIpv6() string {
	if o == nil || o.Ipv6.Get() == nil {
		var ret string
		return ret
	}

	return *o.Ipv6.Get()
}

// GetIpv6Ok returns a tuple with the Ipv6 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DedicatedServer) GetIpv6Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ipv6.Get(), o.Ipv6.IsSet()
}

// SetIpv6 sets field value
func (o *DedicatedServer) SetIpv6(v string) {
	o.Ipv6.Set(&v)
}

// GetNodeId returns the NodeId field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *DedicatedServer) GetNodeId() float32 {
	if o == nil || o.NodeId.Get() == nil {
		var ret float32
		return ret
	}

	return *o.NodeId.Get()
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DedicatedServer) GetNodeIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NodeId.Get(), o.NodeId.IsSet()
}

// SetNodeId sets field value
func (o *DedicatedServer) SetNodeId(v float32) {
	o.NodeId.Set(&v)
}

// GetName returns the Name field value
func (o *DedicatedServer) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DedicatedServer) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DedicatedServer) SetName(v string) {
	o.Name = v
}

// GetComment returns the Comment field value
func (o *DedicatedServer) GetComment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Comment
}

// GetCommentOk returns a tuple with the Comment field value
// and a boolean to check if the value has been set.
func (o *DedicatedServer) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comment, true
}

// SetComment sets field value
func (o *DedicatedServer) SetComment(v string) {
	o.Comment = v
}

// GetVncPass returns the VncPass field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DedicatedServer) GetVncPass() string {
	if o == nil || o.VncPass.Get() == nil {
		var ret string
		return ret
	}

	return *o.VncPass.Get()
}

// GetVncPassOk returns a tuple with the VncPass field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DedicatedServer) GetVncPassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VncPass.Get(), o.VncPass.IsSet()
}

// SetVncPass sets field value
func (o *DedicatedServer) SetVncPass(v string) {
	o.VncPass.Set(&v)
}

// GetStatus returns the Status field value
func (o *DedicatedServer) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DedicatedServer) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DedicatedServer) SetStatus(v string) {
	o.Status = v
}

// GetOsId returns the OsId field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *DedicatedServer) GetOsId() float32 {
	if o == nil || o.OsId.Get() == nil {
		var ret float32
		return ret
	}

	return *o.OsId.Get()
}

// GetOsIdOk returns a tuple with the OsId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DedicatedServer) GetOsIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OsId.Get(), o.OsId.IsSet()
}

// SetOsId sets field value
func (o *DedicatedServer) SetOsId(v float32) {
	o.OsId.Set(&v)
}

// GetCpId returns the CpId field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *DedicatedServer) GetCpId() float32 {
	if o == nil || o.CpId.Get() == nil {
		var ret float32
		return ret
	}

	return *o.CpId.Get()
}

// GetCpIdOk returns a tuple with the CpId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DedicatedServer) GetCpIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CpId.Get(), o.CpId.IsSet()
}

// SetCpId sets field value
func (o *DedicatedServer) SetCpId(v float32) {
	o.CpId.Set(&v)
}

// GetBandwidthId returns the BandwidthId field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *DedicatedServer) GetBandwidthId() float32 {
	if o == nil || o.BandwidthId.Get() == nil {
		var ret float32
		return ret
	}

	return *o.BandwidthId.Get()
}

// GetBandwidthIdOk returns a tuple with the BandwidthId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DedicatedServer) GetBandwidthIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.BandwidthId.Get(), o.BandwidthId.IsSet()
}

// SetBandwidthId sets field value
func (o *DedicatedServer) SetBandwidthId(v float32) {
	o.BandwidthId.Set(&v)
}

// GetNetworkDriveId returns the NetworkDriveId field value
// If the value is explicit nil, the zero value for []float32 will be returned
func (o *DedicatedServer) GetNetworkDriveId() []float32 {
	if o == nil {
		var ret []float32
		return ret
	}

	return o.NetworkDriveId
}

// GetNetworkDriveIdOk returns a tuple with the NetworkDriveId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DedicatedServer) GetNetworkDriveIdOk() ([]float32, bool) {
	if o == nil || IsNil(o.NetworkDriveId) {
		return nil, false
	}
	return o.NetworkDriveId, true
}

// SetNetworkDriveId sets field value
func (o *DedicatedServer) SetNetworkDriveId(v []float32) {
	o.NetworkDriveId = v
}

// GetAdditionalIpAddrId returns the AdditionalIpAddrId field value
// If the value is explicit nil, the zero value for []float32 will be returned
func (o *DedicatedServer) GetAdditionalIpAddrId() []float32 {
	if o == nil {
		var ret []float32
		return ret
	}

	return o.AdditionalIpAddrId
}

// GetAdditionalIpAddrIdOk returns a tuple with the AdditionalIpAddrId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DedicatedServer) GetAdditionalIpAddrIdOk() ([]float32, bool) {
	if o == nil || IsNil(o.AdditionalIpAddrId) {
		return nil, false
	}
	return o.AdditionalIpAddrId, true
}

// SetAdditionalIpAddrId sets field value
func (o *DedicatedServer) SetAdditionalIpAddrId(v []float32) {
	o.AdditionalIpAddrId = v
}

// GetPlanId returns the PlanId field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *DedicatedServer) GetPlanId() float32 {
	if o == nil || o.PlanId.Get() == nil {
		var ret float32
		return ret
	}

	return *o.PlanId.Get()
}

// GetPlanIdOk returns a tuple with the PlanId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DedicatedServer) GetPlanIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlanId.Get(), o.PlanId.IsSet()
}

// SetPlanId sets field value
func (o *DedicatedServer) SetPlanId(v float32) {
	o.PlanId.Set(&v)
}

// GetPrice returns the Price field value
func (o *DedicatedServer) GetPrice() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *DedicatedServer) GetPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *DedicatedServer) SetPrice(v float32) {
	o.Price = v
}

// GetLocation returns the Location field value
func (o *DedicatedServer) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *DedicatedServer) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *DedicatedServer) SetLocation(v string) {
	o.Location = v
}

// GetAutoinstallReady returns the AutoinstallReady field value
func (o *DedicatedServer) GetAutoinstallReady() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AutoinstallReady
}

// GetAutoinstallReadyOk returns a tuple with the AutoinstallReady field value
// and a boolean to check if the value has been set.
func (o *DedicatedServer) GetAutoinstallReadyOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoinstallReady, true
}

// SetAutoinstallReady sets field value
func (o *DedicatedServer) SetAutoinstallReady(v float32) {
	o.AutoinstallReady = v
}

// GetPassword returns the Password field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DedicatedServer) GetPassword() string {
	if o == nil || o.Password.Get() == nil {
		var ret string
		return ret
	}

	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DedicatedServer) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// SetPassword sets field value
func (o *DedicatedServer) SetPassword(v string) {
	o.Password.Set(&v)
}

func (o DedicatedServer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DedicatedServer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["cpu_description"] = o.CpuDescription
	toSerialize["hdd_description"] = o.HddDescription
	toSerialize["ram_description"] = o.RamDescription
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["ip"] = o.Ip.Get()
	toSerialize["ipmi_ip"] = o.IpmiIp.Get()
	toSerialize["ipmi_login"] = o.IpmiLogin.Get()
	toSerialize["ipmi_password"] = o.IpmiPassword.Get()
	toSerialize["ipv6"] = o.Ipv6.Get()
	toSerialize["node_id"] = o.NodeId.Get()
	toSerialize["name"] = o.Name
	toSerialize["comment"] = o.Comment
	toSerialize["vnc_pass"] = o.VncPass.Get()
	toSerialize["status"] = o.Status
	toSerialize["os_id"] = o.OsId.Get()
	toSerialize["cp_id"] = o.CpId.Get()
	toSerialize["bandwidth_id"] = o.BandwidthId.Get()
	if o.NetworkDriveId != nil {
		toSerialize["network_drive_id"] = o.NetworkDriveId
	}
	if o.AdditionalIpAddrId != nil {
		toSerialize["additional_ip_addr_id"] = o.AdditionalIpAddrId
	}
	toSerialize["plan_id"] = o.PlanId.Get()
	toSerialize["price"] = o.Price
	toSerialize["location"] = o.Location
	toSerialize["autoinstall_ready"] = o.AutoinstallReady
	toSerialize["password"] = o.Password.Get()
	return toSerialize, nil
}

type NullableDedicatedServer struct {
	value *DedicatedServer
	isSet bool
}

func (v NullableDedicatedServer) Get() *DedicatedServer {
	return v.value
}

func (v *NullableDedicatedServer) Set(val *DedicatedServer) {
	v.value = val
	v.isSet = true
}

func (v NullableDedicatedServer) IsSet() bool {
	return v.isSet
}

func (v *NullableDedicatedServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDedicatedServer(val *DedicatedServer) *NullableDedicatedServer {
	return &NullableDedicatedServer{value: val, isSet: true}
}

func (v NullableDedicatedServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDedicatedServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


