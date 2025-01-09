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

// checks if the ServersConfiguratorRequirements type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServersConfiguratorRequirements{}

// ServersConfiguratorRequirements struct for ServersConfiguratorRequirements
type ServersConfiguratorRequirements struct {
	// Минимальное количество ядер процессора.
	CpuMin float32 `json:"cpu_min"`
	// Размер шага ядер процессора.
	CpuStep float32 `json:"cpu_step"`
	// Максимальное количество ядер процессора.
	CpuMax float32 `json:"cpu_max"`
	// Минимальное количество оперативной памяти (в Мб).
	RamMin float32 `json:"ram_min"`
	// Размер шага оперативной памяти.
	RamStep float32 `json:"ram_step"`
	// Максимальное количество оперативной памяти (в Мб).
	RamMax float32 `json:"ram_max"`
	// Минимальный размер диска (в Мб).
	DiskMin float32 `json:"disk_min"`
	// Размер шага диска
	DiskStep float32 `json:"disk_step"`
	// Максимальный размер диска (в Мб).
	DiskMax float32 `json:"disk_max"`
	// Минимальныая пропускная способноть интернет-канала (в Мб)
	NetworkBandwidthMin float32 `json:"network_bandwidth_min"`
	// Размер шага пропускной способноти интернет-канала (в Мб)
	NetworkBandwidthStep float32 `json:"network_bandwidth_step"`
	// Максимальная пропускная способноть интернет-канала (в Мб)
	NetworkBandwidthMax float32 `json:"network_bandwidth_max"`
	// Минимальное количество видеокарт
	GpuMin NullableFloat32 `json:"gpu_min"`
	// Максимальное количество видеокарт
	GpuMax NullableFloat32 `json:"gpu_max"`
	// Размер шага видеокарт
	GpuStep NullableFloat32 `json:"gpu_step"`
}

// NewServersConfiguratorRequirements instantiates a new ServersConfiguratorRequirements object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServersConfiguratorRequirements(cpuMin float32, cpuStep float32, cpuMax float32, ramMin float32, ramStep float32, ramMax float32, diskMin float32, diskStep float32, diskMax float32, networkBandwidthMin float32, networkBandwidthStep float32, networkBandwidthMax float32, gpuMin NullableFloat32, gpuMax NullableFloat32, gpuStep NullableFloat32) *ServersConfiguratorRequirements {
	this := ServersConfiguratorRequirements{}
	this.CpuMin = cpuMin
	this.CpuStep = cpuStep
	this.CpuMax = cpuMax
	this.RamMin = ramMin
	this.RamStep = ramStep
	this.RamMax = ramMax
	this.DiskMin = diskMin
	this.DiskStep = diskStep
	this.DiskMax = diskMax
	this.NetworkBandwidthMin = networkBandwidthMin
	this.NetworkBandwidthStep = networkBandwidthStep
	this.NetworkBandwidthMax = networkBandwidthMax
	this.GpuMin = gpuMin
	this.GpuMax = gpuMax
	this.GpuStep = gpuStep
	return &this
}

// NewServersConfiguratorRequirementsWithDefaults instantiates a new ServersConfiguratorRequirements object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServersConfiguratorRequirementsWithDefaults() *ServersConfiguratorRequirements {
	this := ServersConfiguratorRequirements{}
	return &this
}

// GetCpuMin returns the CpuMin field value
func (o *ServersConfiguratorRequirements) GetCpuMin() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CpuMin
}

// GetCpuMinOk returns a tuple with the CpuMin field value
// and a boolean to check if the value has been set.
func (o *ServersConfiguratorRequirements) GetCpuMinOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CpuMin, true
}

// SetCpuMin sets field value
func (o *ServersConfiguratorRequirements) SetCpuMin(v float32) {
	o.CpuMin = v
}

// GetCpuStep returns the CpuStep field value
func (o *ServersConfiguratorRequirements) GetCpuStep() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CpuStep
}

// GetCpuStepOk returns a tuple with the CpuStep field value
// and a boolean to check if the value has been set.
func (o *ServersConfiguratorRequirements) GetCpuStepOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CpuStep, true
}

// SetCpuStep sets field value
func (o *ServersConfiguratorRequirements) SetCpuStep(v float32) {
	o.CpuStep = v
}

// GetCpuMax returns the CpuMax field value
func (o *ServersConfiguratorRequirements) GetCpuMax() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CpuMax
}

// GetCpuMaxOk returns a tuple with the CpuMax field value
// and a boolean to check if the value has been set.
func (o *ServersConfiguratorRequirements) GetCpuMaxOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CpuMax, true
}

// SetCpuMax sets field value
func (o *ServersConfiguratorRequirements) SetCpuMax(v float32) {
	o.CpuMax = v
}

// GetRamMin returns the RamMin field value
func (o *ServersConfiguratorRequirements) GetRamMin() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.RamMin
}

// GetRamMinOk returns a tuple with the RamMin field value
// and a boolean to check if the value has been set.
func (o *ServersConfiguratorRequirements) GetRamMinOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RamMin, true
}

// SetRamMin sets field value
func (o *ServersConfiguratorRequirements) SetRamMin(v float32) {
	o.RamMin = v
}

// GetRamStep returns the RamStep field value
func (o *ServersConfiguratorRequirements) GetRamStep() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.RamStep
}

// GetRamStepOk returns a tuple with the RamStep field value
// and a boolean to check if the value has been set.
func (o *ServersConfiguratorRequirements) GetRamStepOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RamStep, true
}

// SetRamStep sets field value
func (o *ServersConfiguratorRequirements) SetRamStep(v float32) {
	o.RamStep = v
}

// GetRamMax returns the RamMax field value
func (o *ServersConfiguratorRequirements) GetRamMax() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.RamMax
}

// GetRamMaxOk returns a tuple with the RamMax field value
// and a boolean to check if the value has been set.
func (o *ServersConfiguratorRequirements) GetRamMaxOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RamMax, true
}

// SetRamMax sets field value
func (o *ServersConfiguratorRequirements) SetRamMax(v float32) {
	o.RamMax = v
}

// GetDiskMin returns the DiskMin field value
func (o *ServersConfiguratorRequirements) GetDiskMin() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DiskMin
}

// GetDiskMinOk returns a tuple with the DiskMin field value
// and a boolean to check if the value has been set.
func (o *ServersConfiguratorRequirements) GetDiskMinOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiskMin, true
}

// SetDiskMin sets field value
func (o *ServersConfiguratorRequirements) SetDiskMin(v float32) {
	o.DiskMin = v
}

// GetDiskStep returns the DiskStep field value
func (o *ServersConfiguratorRequirements) GetDiskStep() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DiskStep
}

// GetDiskStepOk returns a tuple with the DiskStep field value
// and a boolean to check if the value has been set.
func (o *ServersConfiguratorRequirements) GetDiskStepOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiskStep, true
}

// SetDiskStep sets field value
func (o *ServersConfiguratorRequirements) SetDiskStep(v float32) {
	o.DiskStep = v
}

// GetDiskMax returns the DiskMax field value
func (o *ServersConfiguratorRequirements) GetDiskMax() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DiskMax
}

// GetDiskMaxOk returns a tuple with the DiskMax field value
// and a boolean to check if the value has been set.
func (o *ServersConfiguratorRequirements) GetDiskMaxOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiskMax, true
}

// SetDiskMax sets field value
func (o *ServersConfiguratorRequirements) SetDiskMax(v float32) {
	o.DiskMax = v
}

// GetNetworkBandwidthMin returns the NetworkBandwidthMin field value
func (o *ServersConfiguratorRequirements) GetNetworkBandwidthMin() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.NetworkBandwidthMin
}

// GetNetworkBandwidthMinOk returns a tuple with the NetworkBandwidthMin field value
// and a boolean to check if the value has been set.
func (o *ServersConfiguratorRequirements) GetNetworkBandwidthMinOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkBandwidthMin, true
}

// SetNetworkBandwidthMin sets field value
func (o *ServersConfiguratorRequirements) SetNetworkBandwidthMin(v float32) {
	o.NetworkBandwidthMin = v
}

// GetNetworkBandwidthStep returns the NetworkBandwidthStep field value
func (o *ServersConfiguratorRequirements) GetNetworkBandwidthStep() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.NetworkBandwidthStep
}

// GetNetworkBandwidthStepOk returns a tuple with the NetworkBandwidthStep field value
// and a boolean to check if the value has been set.
func (o *ServersConfiguratorRequirements) GetNetworkBandwidthStepOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkBandwidthStep, true
}

// SetNetworkBandwidthStep sets field value
func (o *ServersConfiguratorRequirements) SetNetworkBandwidthStep(v float32) {
	o.NetworkBandwidthStep = v
}

// GetNetworkBandwidthMax returns the NetworkBandwidthMax field value
func (o *ServersConfiguratorRequirements) GetNetworkBandwidthMax() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.NetworkBandwidthMax
}

// GetNetworkBandwidthMaxOk returns a tuple with the NetworkBandwidthMax field value
// and a boolean to check if the value has been set.
func (o *ServersConfiguratorRequirements) GetNetworkBandwidthMaxOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkBandwidthMax, true
}

// SetNetworkBandwidthMax sets field value
func (o *ServersConfiguratorRequirements) SetNetworkBandwidthMax(v float32) {
	o.NetworkBandwidthMax = v
}

// GetGpuMin returns the GpuMin field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *ServersConfiguratorRequirements) GetGpuMin() float32 {
	if o == nil || o.GpuMin.Get() == nil {
		var ret float32
		return ret
	}

	return *o.GpuMin.Get()
}

// GetGpuMinOk returns a tuple with the GpuMin field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServersConfiguratorRequirements) GetGpuMinOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.GpuMin.Get(), o.GpuMin.IsSet()
}

// SetGpuMin sets field value
func (o *ServersConfiguratorRequirements) SetGpuMin(v float32) {
	o.GpuMin.Set(&v)
}

// GetGpuMax returns the GpuMax field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *ServersConfiguratorRequirements) GetGpuMax() float32 {
	if o == nil || o.GpuMax.Get() == nil {
		var ret float32
		return ret
	}

	return *o.GpuMax.Get()
}

// GetGpuMaxOk returns a tuple with the GpuMax field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServersConfiguratorRequirements) GetGpuMaxOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.GpuMax.Get(), o.GpuMax.IsSet()
}

// SetGpuMax sets field value
func (o *ServersConfiguratorRequirements) SetGpuMax(v float32) {
	o.GpuMax.Set(&v)
}

// GetGpuStep returns the GpuStep field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *ServersConfiguratorRequirements) GetGpuStep() float32 {
	if o == nil || o.GpuStep.Get() == nil {
		var ret float32
		return ret
	}

	return *o.GpuStep.Get()
}

// GetGpuStepOk returns a tuple with the GpuStep field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServersConfiguratorRequirements) GetGpuStepOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.GpuStep.Get(), o.GpuStep.IsSet()
}

// SetGpuStep sets field value
func (o *ServersConfiguratorRequirements) SetGpuStep(v float32) {
	o.GpuStep.Set(&v)
}

func (o ServersConfiguratorRequirements) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServersConfiguratorRequirements) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cpu_min"] = o.CpuMin
	toSerialize["cpu_step"] = o.CpuStep
	toSerialize["cpu_max"] = o.CpuMax
	toSerialize["ram_min"] = o.RamMin
	toSerialize["ram_step"] = o.RamStep
	toSerialize["ram_max"] = o.RamMax
	toSerialize["disk_min"] = o.DiskMin
	toSerialize["disk_step"] = o.DiskStep
	toSerialize["disk_max"] = o.DiskMax
	toSerialize["network_bandwidth_min"] = o.NetworkBandwidthMin
	toSerialize["network_bandwidth_step"] = o.NetworkBandwidthStep
	toSerialize["network_bandwidth_max"] = o.NetworkBandwidthMax
	toSerialize["gpu_min"] = o.GpuMin.Get()
	toSerialize["gpu_max"] = o.GpuMax.Get()
	toSerialize["gpu_step"] = o.GpuStep.Get()
	return toSerialize, nil
}

type NullableServersConfiguratorRequirements struct {
	value *ServersConfiguratorRequirements
	isSet bool
}

func (v NullableServersConfiguratorRequirements) Get() *ServersConfiguratorRequirements {
	return v.value
}

func (v *NullableServersConfiguratorRequirements) Set(val *ServersConfiguratorRequirements) {
	v.value = val
	v.isSet = true
}

func (v NullableServersConfiguratorRequirements) IsSet() bool {
	return v.isSet
}

func (v *NullableServersConfiguratorRequirements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServersConfiguratorRequirements(val *ServersConfiguratorRequirements) *NullableServersConfiguratorRequirements {
	return &NullableServersConfiguratorRequirements{value: val, isSet: true}
}

func (v NullableServersConfiguratorRequirements) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServersConfiguratorRequirements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


