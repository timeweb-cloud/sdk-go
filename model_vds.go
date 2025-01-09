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

// checks if the Vds type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Vds{}

// Vds Сервер
type Vds struct {
	// ID для каждого экземпляра сервера. Автоматически генерируется при создании.
	Id float32 `json:"id"`
	// Удобочитаемое имя, установленное для сервера.
	Name string `json:"name"`
	// Комментарий к серверу.
	Comment string `json:"comment"`
	// Дата создания сервера в формате ISO8061.
	CreatedAt string `json:"created_at"`
	Os VdsOs `json:"os"`
	Software NullableVdsSoftware `json:"software"`
	// ID тарифа сервера.
	PresetId NullableFloat32 `json:"preset_id"`
	// Локация сервера.
	Location string `json:"location"`
	// ID конфигуратора сервера.
	ConfiguratorId NullableFloat32 `json:"configurator_id"`
	// Режим загрузки ОС сервера.
	BootMode string `json:"boot_mode"`
	// Статус сервера.
	Status string `json:"status"`
	// Значение времени, указанное в комбинированном формате даты и времени ISO8601, которое представляет, когда был запущен сервер.
	StartAt NullableTime `json:"start_at"`
	// Это логическое значение, которое показывает, включена ли защита от DDoS у данного сервера.
	IsDdosGuard bool `json:"is_ddos_guard"`
	// Это логическое значение, которое показывает, доступно ли подключение по SSH для поддержки.
	IsMasterSsh bool `json:"is_master_ssh"`
	// Это логическое значение, которое показывает, является ли CPU выделенным.
	IsDedicatedCpu bool `json:"is_dedicated_cpu"`
	// Количество видеокарт сервера.
	Gpu float32 `json:"gpu"`
	// Количество ядер процессора сервера.
	Cpu float32 `json:"cpu"`
	// Частота ядер процессора сервера.
	CpuFrequency string `json:"cpu_frequency"`
	// Размер (в Мб) ОЗУ сервера.
	Ram float32 `json:"ram"`
	// Список дисков сервера.
	Disks []VdsDisksInner `json:"disks"`
	// ID аватара сервера. Описание методов работы с аватарами появится позднее.
	AvatarId NullableString `json:"avatar_id"`
	// Пароль от VNC.
	VncPass string `json:"vnc_pass"`
	// Пароль root сервера или пароль Администратора для серверов Windows.
	RootPass NullableString `json:"root_pass"`
	Image NullableVdsImage `json:"image"`
	// Список сетей сервера.
	Networks []VdsNetworksInner `json:"networks"`
	// Cloud-init скрипт.
	CloudInit NullableString `json:"cloud_init"`
	// Это логическое значение, которое показывает, включен ли QEMU-agent на сервере.
	IsQemuAgent bool `json:"is_qemu_agent"`
	AvailabilityZone AvailabilityZone `json:"availability_zone"`
}

// NewVds instantiates a new Vds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVds(id float32, name string, comment string, createdAt string, os VdsOs, software NullableVdsSoftware, presetId NullableFloat32, location string, configuratorId NullableFloat32, bootMode string, status string, startAt NullableTime, isDdosGuard bool, isMasterSsh bool, isDedicatedCpu bool, gpu float32, cpu float32, cpuFrequency string, ram float32, disks []VdsDisksInner, avatarId NullableString, vncPass string, rootPass NullableString, image NullableVdsImage, networks []VdsNetworksInner, cloudInit NullableString, isQemuAgent bool, availabilityZone AvailabilityZone) *Vds {
	this := Vds{}
	this.Id = id
	this.Name = name
	this.Comment = comment
	this.CreatedAt = createdAt
	this.Os = os
	this.Software = software
	this.PresetId = presetId
	this.Location = location
	this.ConfiguratorId = configuratorId
	this.BootMode = bootMode
	this.Status = status
	this.StartAt = startAt
	this.IsDdosGuard = isDdosGuard
	this.IsMasterSsh = isMasterSsh
	this.IsDedicatedCpu = isDedicatedCpu
	this.Gpu = gpu
	this.Cpu = cpu
	this.CpuFrequency = cpuFrequency
	this.Ram = ram
	this.Disks = disks
	this.AvatarId = avatarId
	this.VncPass = vncPass
	this.RootPass = rootPass
	this.Image = image
	this.Networks = networks
	this.CloudInit = cloudInit
	this.IsQemuAgent = isQemuAgent
	this.AvailabilityZone = availabilityZone
	return &this
}

// NewVdsWithDefaults instantiates a new Vds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVdsWithDefaults() *Vds {
	this := Vds{}
	return &this
}

// GetId returns the Id field value
func (o *Vds) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Vds) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Vds) SetId(v float32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Vds) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Vds) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Vds) SetName(v string) {
	o.Name = v
}

// GetComment returns the Comment field value
func (o *Vds) GetComment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Comment
}

// GetCommentOk returns a tuple with the Comment field value
// and a boolean to check if the value has been set.
func (o *Vds) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comment, true
}

// SetComment sets field value
func (o *Vds) SetComment(v string) {
	o.Comment = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Vds) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Vds) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Vds) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetOs returns the Os field value
func (o *Vds) GetOs() VdsOs {
	if o == nil {
		var ret VdsOs
		return ret
	}

	return o.Os
}

// GetOsOk returns a tuple with the Os field value
// and a boolean to check if the value has been set.
func (o *Vds) GetOsOk() (*VdsOs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Os, true
}

// SetOs sets field value
func (o *Vds) SetOs(v VdsOs) {
	o.Os = v
}

// GetSoftware returns the Software field value
// If the value is explicit nil, the zero value for VdsSoftware will be returned
func (o *Vds) GetSoftware() VdsSoftware {
	if o == nil || o.Software.Get() == nil {
		var ret VdsSoftware
		return ret
	}

	return *o.Software.Get()
}

// GetSoftwareOk returns a tuple with the Software field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Vds) GetSoftwareOk() (*VdsSoftware, bool) {
	if o == nil {
		return nil, false
	}
	return o.Software.Get(), o.Software.IsSet()
}

// SetSoftware sets field value
func (o *Vds) SetSoftware(v VdsSoftware) {
	o.Software.Set(&v)
}

// GetPresetId returns the PresetId field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *Vds) GetPresetId() float32 {
	if o == nil || o.PresetId.Get() == nil {
		var ret float32
		return ret
	}

	return *o.PresetId.Get()
}

// GetPresetIdOk returns a tuple with the PresetId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Vds) GetPresetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PresetId.Get(), o.PresetId.IsSet()
}

// SetPresetId sets field value
func (o *Vds) SetPresetId(v float32) {
	o.PresetId.Set(&v)
}

// GetLocation returns the Location field value
func (o *Vds) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *Vds) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *Vds) SetLocation(v string) {
	o.Location = v
}

// GetConfiguratorId returns the ConfiguratorId field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *Vds) GetConfiguratorId() float32 {
	if o == nil || o.ConfiguratorId.Get() == nil {
		var ret float32
		return ret
	}

	return *o.ConfiguratorId.Get()
}

// GetConfiguratorIdOk returns a tuple with the ConfiguratorId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Vds) GetConfiguratorIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfiguratorId.Get(), o.ConfiguratorId.IsSet()
}

// SetConfiguratorId sets field value
func (o *Vds) SetConfiguratorId(v float32) {
	o.ConfiguratorId.Set(&v)
}

// GetBootMode returns the BootMode field value
func (o *Vds) GetBootMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BootMode
}

// GetBootModeOk returns a tuple with the BootMode field value
// and a boolean to check if the value has been set.
func (o *Vds) GetBootModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BootMode, true
}

// SetBootMode sets field value
func (o *Vds) SetBootMode(v string) {
	o.BootMode = v
}

// GetStatus returns the Status field value
func (o *Vds) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Vds) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Vds) SetStatus(v string) {
	o.Status = v
}

// GetStartAt returns the StartAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Vds) GetStartAt() time.Time {
	if o == nil || o.StartAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.StartAt.Get()
}

// GetStartAtOk returns a tuple with the StartAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Vds) GetStartAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartAt.Get(), o.StartAt.IsSet()
}

// SetStartAt sets field value
func (o *Vds) SetStartAt(v time.Time) {
	o.StartAt.Set(&v)
}

// GetIsDdosGuard returns the IsDdosGuard field value
func (o *Vds) GetIsDdosGuard() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDdosGuard
}

// GetIsDdosGuardOk returns a tuple with the IsDdosGuard field value
// and a boolean to check if the value has been set.
func (o *Vds) GetIsDdosGuardOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDdosGuard, true
}

// SetIsDdosGuard sets field value
func (o *Vds) SetIsDdosGuard(v bool) {
	o.IsDdosGuard = v
}

// GetIsMasterSsh returns the IsMasterSsh field value
func (o *Vds) GetIsMasterSsh() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMasterSsh
}

// GetIsMasterSshOk returns a tuple with the IsMasterSsh field value
// and a boolean to check if the value has been set.
func (o *Vds) GetIsMasterSshOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMasterSsh, true
}

// SetIsMasterSsh sets field value
func (o *Vds) SetIsMasterSsh(v bool) {
	o.IsMasterSsh = v
}

// GetIsDedicatedCpu returns the IsDedicatedCpu field value
func (o *Vds) GetIsDedicatedCpu() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDedicatedCpu
}

// GetIsDedicatedCpuOk returns a tuple with the IsDedicatedCpu field value
// and a boolean to check if the value has been set.
func (o *Vds) GetIsDedicatedCpuOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDedicatedCpu, true
}

// SetIsDedicatedCpu sets field value
func (o *Vds) SetIsDedicatedCpu(v bool) {
	o.IsDedicatedCpu = v
}

// GetGpu returns the Gpu field value
func (o *Vds) GetGpu() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Gpu
}

// GetGpuOk returns a tuple with the Gpu field value
// and a boolean to check if the value has been set.
func (o *Vds) GetGpuOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gpu, true
}

// SetGpu sets field value
func (o *Vds) SetGpu(v float32) {
	o.Gpu = v
}

// GetCpu returns the Cpu field value
func (o *Vds) GetCpu() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value
// and a boolean to check if the value has been set.
func (o *Vds) GetCpuOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cpu, true
}

// SetCpu sets field value
func (o *Vds) SetCpu(v float32) {
	o.Cpu = v
}

// GetCpuFrequency returns the CpuFrequency field value
func (o *Vds) GetCpuFrequency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CpuFrequency
}

// GetCpuFrequencyOk returns a tuple with the CpuFrequency field value
// and a boolean to check if the value has been set.
func (o *Vds) GetCpuFrequencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CpuFrequency, true
}

// SetCpuFrequency sets field value
func (o *Vds) SetCpuFrequency(v string) {
	o.CpuFrequency = v
}

// GetRam returns the Ram field value
func (o *Vds) GetRam() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Ram
}

// GetRamOk returns a tuple with the Ram field value
// and a boolean to check if the value has been set.
func (o *Vds) GetRamOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ram, true
}

// SetRam sets field value
func (o *Vds) SetRam(v float32) {
	o.Ram = v
}

// GetDisks returns the Disks field value
func (o *Vds) GetDisks() []VdsDisksInner {
	if o == nil {
		var ret []VdsDisksInner
		return ret
	}

	return o.Disks
}

// GetDisksOk returns a tuple with the Disks field value
// and a boolean to check if the value has been set.
func (o *Vds) GetDisksOk() ([]VdsDisksInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Disks, true
}

// SetDisks sets field value
func (o *Vds) SetDisks(v []VdsDisksInner) {
	o.Disks = v
}

// GetAvatarId returns the AvatarId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Vds) GetAvatarId() string {
	if o == nil || o.AvatarId.Get() == nil {
		var ret string
		return ret
	}

	return *o.AvatarId.Get()
}

// GetAvatarIdOk returns a tuple with the AvatarId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Vds) GetAvatarIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvatarId.Get(), o.AvatarId.IsSet()
}

// SetAvatarId sets field value
func (o *Vds) SetAvatarId(v string) {
	o.AvatarId.Set(&v)
}

// GetVncPass returns the VncPass field value
func (o *Vds) GetVncPass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VncPass
}

// GetVncPassOk returns a tuple with the VncPass field value
// and a boolean to check if the value has been set.
func (o *Vds) GetVncPassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VncPass, true
}

// SetVncPass sets field value
func (o *Vds) SetVncPass(v string) {
	o.VncPass = v
}

// GetRootPass returns the RootPass field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Vds) GetRootPass() string {
	if o == nil || o.RootPass.Get() == nil {
		var ret string
		return ret
	}

	return *o.RootPass.Get()
}

// GetRootPassOk returns a tuple with the RootPass field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Vds) GetRootPassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RootPass.Get(), o.RootPass.IsSet()
}

// SetRootPass sets field value
func (o *Vds) SetRootPass(v string) {
	o.RootPass.Set(&v)
}

// GetImage returns the Image field value
// If the value is explicit nil, the zero value for VdsImage will be returned
func (o *Vds) GetImage() VdsImage {
	if o == nil || o.Image.Get() == nil {
		var ret VdsImage
		return ret
	}

	return *o.Image.Get()
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Vds) GetImageOk() (*VdsImage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Image.Get(), o.Image.IsSet()
}

// SetImage sets field value
func (o *Vds) SetImage(v VdsImage) {
	o.Image.Set(&v)
}

// GetNetworks returns the Networks field value
func (o *Vds) GetNetworks() []VdsNetworksInner {
	if o == nil {
		var ret []VdsNetworksInner
		return ret
	}

	return o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value
// and a boolean to check if the value has been set.
func (o *Vds) GetNetworksOk() ([]VdsNetworksInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Networks, true
}

// SetNetworks sets field value
func (o *Vds) SetNetworks(v []VdsNetworksInner) {
	o.Networks = v
}

// GetCloudInit returns the CloudInit field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Vds) GetCloudInit() string {
	if o == nil || o.CloudInit.Get() == nil {
		var ret string
		return ret
	}

	return *o.CloudInit.Get()
}

// GetCloudInitOk returns a tuple with the CloudInit field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Vds) GetCloudInitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CloudInit.Get(), o.CloudInit.IsSet()
}

// SetCloudInit sets field value
func (o *Vds) SetCloudInit(v string) {
	o.CloudInit.Set(&v)
}

// GetIsQemuAgent returns the IsQemuAgent field value
func (o *Vds) GetIsQemuAgent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsQemuAgent
}

// GetIsQemuAgentOk returns a tuple with the IsQemuAgent field value
// and a boolean to check if the value has been set.
func (o *Vds) GetIsQemuAgentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsQemuAgent, true
}

// SetIsQemuAgent sets field value
func (o *Vds) SetIsQemuAgent(v bool) {
	o.IsQemuAgent = v
}

// GetAvailabilityZone returns the AvailabilityZone field value
func (o *Vds) GetAvailabilityZone() AvailabilityZone {
	if o == nil {
		var ret AvailabilityZone
		return ret
	}

	return o.AvailabilityZone
}

// GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field value
// and a boolean to check if the value has been set.
func (o *Vds) GetAvailabilityZoneOk() (*AvailabilityZone, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvailabilityZone, true
}

// SetAvailabilityZone sets field value
func (o *Vds) SetAvailabilityZone(v AvailabilityZone) {
	o.AvailabilityZone = v
}

func (o Vds) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Vds) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["comment"] = o.Comment
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["os"] = o.Os
	toSerialize["software"] = o.Software.Get()
	toSerialize["preset_id"] = o.PresetId.Get()
	toSerialize["location"] = o.Location
	toSerialize["configurator_id"] = o.ConfiguratorId.Get()
	toSerialize["boot_mode"] = o.BootMode
	toSerialize["status"] = o.Status
	toSerialize["start_at"] = o.StartAt.Get()
	toSerialize["is_ddos_guard"] = o.IsDdosGuard
	toSerialize["is_master_ssh"] = o.IsMasterSsh
	toSerialize["is_dedicated_cpu"] = o.IsDedicatedCpu
	toSerialize["gpu"] = o.Gpu
	toSerialize["cpu"] = o.Cpu
	toSerialize["cpu_frequency"] = o.CpuFrequency
	toSerialize["ram"] = o.Ram
	toSerialize["disks"] = o.Disks
	toSerialize["avatar_id"] = o.AvatarId.Get()
	toSerialize["vnc_pass"] = o.VncPass
	toSerialize["root_pass"] = o.RootPass.Get()
	toSerialize["image"] = o.Image.Get()
	toSerialize["networks"] = o.Networks
	toSerialize["cloud_init"] = o.CloudInit.Get()
	toSerialize["is_qemu_agent"] = o.IsQemuAgent
	toSerialize["availability_zone"] = o.AvailabilityZone
	return toSerialize, nil
}

type NullableVds struct {
	value *Vds
	isSet bool
}

func (v NullableVds) Get() *Vds {
	return v.value
}

func (v *NullableVds) Set(val *Vds) {
	v.value = val
	v.isSet = true
}

func (v NullableVds) IsSet() bool {
	return v.isSet
}

func (v *NullableVds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVds(val *Vds) *NullableVds {
	return &NullableVds{value: val, isSet: true}
}

func (v NullableVds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


