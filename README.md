# Go API client for openapi

# Введение
API Timeweb Cloud позволяет вам управлять ресурсами в облаке программным способом с использованием обычных HTTP-запросов.

Множество функций, которые доступны в панели управления Timeweb Cloud, также доступны через API, что позволяет вам автоматизировать ваши собственные сценарии.

В этой документации сперва будет описан общий дизайн и принципы работы API, а после этого конкретные конечные точки. Также будут приведены примеры запросов к ним.


## Запросы
Запросы должны выполняться по протоколу `HTTPS`, чтобы гарантировать шифрование транзакций. Поддерживаются следующие методы запроса:
|Метод|Применение|
|--- |--- |
|GET|Извлекает данные о коллекциях и отдельных ресурсах.|
|POST|Для коллекций создает новый ресурс этого типа. Также используется для выполнения действий с конкретным ресурсом.|
|PUT|Обновляет существующий ресурс.|
|PATCH|Некоторые ресурсы поддерживают частичное обновление, то есть обновление только части атрибутов ресурса, в этом случае вместо метода PUT будет использован PATCH.|
|DELETE|Удаляет ресурс.|

Методы `POST`, `PUT` и `PATCH` могут включать объект в тело запроса с типом содержимого `application/json`.

### Параметры в запросах
Некоторые коллекции поддерживают пагинацию, поиск или сортировку в запросах. В параметрах запроса требуется передать:
- `limit` — обозначает количество записей, которое необходимо вернуть
 - `offset` — указывает на смещение, относительно начала списка
 - `search` — позволяет указать набор символов для поиска
 - `sort` — можно задать правило сортировки коллекции

## Ответы
Запросы вернут один из следующих кодов состояния ответа HTTP:

|Статус|Описание|
|--- |--- |
|200 OK|Действие с ресурсом было выполнено успешно.|
|201 Created|Ресурс был успешно создан. При этом ресурс может быть как уже готовым к использованию, так и находиться в процессе запуска.|
|204 No Content|Действие с ресурсом было выполнено успешно, и ответ не содержит дополнительной информации в теле.|
|400 Bad Request|Был отправлен неверный запрос, например, в нем отсутствуют обязательные параметры и т. д. Тело ответа будет содержать дополнительную информацию об ошибке.|
|401 Unauthorized|Ошибка аутентификации.|
|403 Forbidden|Аутентификация прошла успешно, но недостаточно прав для выполнения действия.|
|404 Not Found|Запрашиваемый ресурс не найден.|
|409 Conflict|Запрос конфликтует с текущим состоянием.|
|423 Locked|Ресурс из запроса заблокирован от применения к нему указанного метода.|
|429 Too Many Requests|Был достигнут лимит по количеству запросов в единицу времени.|
|500 Internal Server Error|При выполнении запроса произошла какая-то внутренняя ошибка. Чтобы решить эту проблему, лучше всего создать тикет в панели управления.|

### Структура успешного ответа
Все конечные точки будут возвращать данные в формате `JSON`. Ответы на `GET`-запросы будут иметь на верхнем уровне следующую структуру атрибутов: 
|Название поля|Тип|Описание|
|--- |--- |--- |
|[entity_name]|object, object[], string[], number[], boolean|Динамическое поле, которое будет меняться в зависимости от запрашиваемого ресурса и будет содержать все атрибуты, необходимые для описания этого ресурса. Например, при запросе списка баз данных будет возвращаться поле `dbs`, а при запросе конкретного облачного сервера `server`. Для некоторых конечных точек в ответе может возвращаться сразу несколько ресурсов.|
|meta|object|Опционально. Объект, который содержит вспомогательную информацию о ресурсе. Чаще всего будет встречаться при запросе коллекций и содержать поле `total`, которое будет указывать на количество элементов в коллекции.|
|response_id|string|Опционально. В большинстве случаев в ответе будет содержаться ID ответа в формате UUIDv4, который однозначно указывает на ваш запрос внутри нашей системы. Если вам потребуется задать вопрос нашей поддержке, приложите к вопросу этот ID— так мы сможем найти ответ на него намного быстрее. Также вы можете использовать этот ID, чтобы убедиться, что это новый ответ на запрос и результат не был получен из кэша.|

Пример запроса на получение списка SSH-ключей:
```
    HTTP/2.0 200 OK
    {
      \"ssh_keys\":[
          {
            \"body\":\"ssh-rsa AAAAB3NzaC1sdfghjkOAsBwWhs= example@device.local\",
            \"created_at\":\"2021-09-15T19:52:27Z\",
            \"expired_at\":null,
            \"id\":5297,
            \"is_default\":false,
            \"name\":\"example@device.local\",
            \"used_at\":null,
            \"used_by\":[]
          }
      ],
      \"meta\":{
          \"total\":1
      },
      \"response_id\":\"94608d15-8672-4eed-8ab6-28bd6fa3cdf7\"
    }
```

### Структура ответа с ошибкой
|Название поля|Тип|Описание|
|--- |--- |--- |
|status_code|number|Короткий числовой идентификатор ошибки.|
|error_code|string|Короткий текстовый идентификатор ошибки, который уточняет числовой идентификатор и удобен для программной обработки. Самый простой пример — это код `not_found` для ошибки 404.|
|message|string, string[]|Опционально. В большинстве случаев в ответе будет содержаться человекочитаемое подробное описание ошибки или ошибок, которые помогут понять, что нужно исправить.|
|response_id|string|Опционально. В большинстве случае в ответе будет содержаться ID ответа в формате UUIDv4, который однозначно указывает на ваш запрос внутри нашей системы. Если вам потребуется задать вопрос нашей поддержке, приложите к вопросу этот ID — так мы сможем найти ответ на него намного быстрее.|

Пример:
```
    HTTP/2.0 403 Forbidden
    {
      \"status_code\": 403,
      \"error_code\":  \"forbidden\",
      \"message\":     \"You do not have access for the attempted action\",
      \"response_id\": \"94608d15-8672-4eed-8ab6-28bd6fa3cdf7\"
    }
```

## Статусы ресурсов
Важно учесть, что при создании большинства ресурсов внутри платформы вам будет сразу возвращен ответ от сервера со статусом `200 OK` или `201 Created` и ID созданного ресурса в теле ответа, но при этом этот ресурс может быть ещё в *состоянии запуска*.

Для того чтобы понять, в каком состоянии сейчас находится ваш ресурс, мы добавили поле `status` в ответ на получение информации о ресурсе.

Список статусов будет отличаться в зависимости от типа ресурса. Увидеть поддерживаемый список статусов вы сможете в описании каждого конкретного ресурса.

 

## Ограничение скорости запросов (Rate Limiting)
Чтобы обеспечить стабильность для всех пользователей, Timeweb Cloud защищает API от всплесков входящего трафика, анализируя количество запросов c каждого аккаунта к каждой конечной точке.

Если ваше приложение отправляет более 20 запросов в секунду на одну конечную точку, то для этого запроса API может вернуть код состояния HTTP `429 Too Many Requests`.


## Аутентификация
Доступ к API осуществляется с помощью JWT-токена. Токенами можно управлять внутри панели управления Timeweb Cloud в разделе *API и Terraform*.

Токен необходимо передавать в заголовке каждого запроса в формате:
```
  Authorization: Bearer $TIMEWEB_CLOUD_TOKEN
```

## Формат примеров API
Примеры в этой документации описаны с помощью `curl`, HTTP-клиента командной строки. На компьютерах `Linux` и `macOS` обычно по умолчанию установлен `curl`, и он доступен для загрузки на всех популярных платформах, включая `Windows`.

Каждый пример разделен на несколько строк символом `\\`, который совместим с `bash`. Типичный пример выглядит так:
```
  curl -X PATCH 
    -H \"Content-Type: application/json\" 
    -H \"Authorization: Bearer $TIMEWEB_CLOUD_TOKEN\" 
    -d '{\"name\":\"Cute Corvus\",\"comment\":\"Development Server\"}' 
    \"https://api.timeweb.cloud/api/v1/dedicated/1051\"
```
- Параметр `-X` задает метод запроса. Для согласованности метод будет указан во всех примерах, даже если он явно не требуется для методов `GET`.
- Строки `-H` задают требуемые HTTP-заголовки.
- Примеры, для которых требуется объект JSON в теле запроса, передают требуемые данные через параметр `-d`.

Чтобы использовать приведенные примеры, не подставляя каждый раз в них свой токен, вы можете добавить токен один раз в переменные окружения в вашей консоли. Например, на `Linux` это можно сделать с помощью команды:

```
TIMEWEB_CLOUD_TOKEN=\"token\"
```

После этого токен будет автоматически подставляться в ваши запросы.

Обратите внимание, что все значения в этой документации являются примерами. Не полагайтесь на IDы операционных систем, тарифов и т.д., используемые в примерах. Используйте соответствующую конечную точку для получения значений перед созданием ресурсов.


## Версионирование
API построено согласно принципам [семантического версионирования](https://semver.org/lang/ru). Это значит, что мы гарантируем обратную совместимость всех изменений в пределах одной мажорной версии.

Мажорная версия каждой конечной точки обозначается в пути запроса, например, запрос `/api/v1/servers` указывает, что этот метод имеет версию 1.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.timeweb.cloud*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AIAgentIntegrationsAPI* | [**CreateAgentBotMaxBot**](docs/AIAgentIntegrationsAPI.md#createagentbotmaxbot) | **Post** /api/v1/cloud-ai/max-bots | Подключение бота MAX к агенту
*AIAgentIntegrationsAPI* | [**CreateAgentBotTelegramBot**](docs/AIAgentIntegrationsAPI.md#createagentbottelegrambot) | **Post** /api/v1/cloud-ai/bots | Подключение Telegram-бота к агенту
*AIAgentIntegrationsAPI* | [**CreateAgentIntegrationCrisp**](docs/AIAgentIntegrationsAPI.md#createagentintegrationcrisp) | **Post** /api/v1/cloud-ai/agent/{agent_id}/crisp | Подключение агента к Crisp
*AIAgentIntegrationsAPI* | [**CreateAgentIntegrationJivo**](docs/AIAgentIntegrationsAPI.md#createagentintegrationjivo) | **Post** /api/v1/cloud-ai/agent/{agent_id}/jivo | Подключение агента к Jivo
*AIAgentIntegrationsAPI* | [**DeleteAgentBotMaxBot**](docs/AIAgentIntegrationsAPI.md#deleteagentbotmaxbot) | **Delete** /api/v1/cloud-ai/max-bots/{bot_id} | Отключение бота MAX
*AIAgentIntegrationsAPI* | [**DeleteAgentBotTelegramBot**](docs/AIAgentIntegrationsAPI.md#deleteagentbottelegrambot) | **Delete** /api/v1/cloud-ai/bots/{bot_id} | Отключение Telegram-бота
*AIAgentIntegrationsAPI* | [**DeleteAgentIntegrationCrisp**](docs/AIAgentIntegrationsAPI.md#deleteagentintegrationcrisp) | **Delete** /api/v1/cloud-ai/agent/{agent_id}/crisp | Отключение агента от Crisp
*AIAgentIntegrationsAPI* | [**DeleteAgentIntegrationJivo**](docs/AIAgentIntegrationsAPI.md#deleteagentintegrationjivo) | **Delete** /api/v1/cloud-ai/agent/{agent_id}/jivo | Отключение агента от Jivo
*AIAgentIntegrationsAPI* | [**GetAgentBotMaxBot**](docs/AIAgentIntegrationsAPI.md#getagentbotmaxbot) | **Get** /api/v1/cloud-ai/max-bots/{bot_id} | Получение информации о боте MAX
*AIAgentIntegrationsAPI* | [**GetAgentBotMaxBots**](docs/AIAgentIntegrationsAPI.md#getagentbotmaxbots) | **Get** /api/v1/cloud-ai/max-bots | Получение списка ботов MAX
*AIAgentIntegrationsAPI* | [**GetAgentBotMaxBotsByAgent**](docs/AIAgentIntegrationsAPI.md#getagentbotmaxbotsbyagent) | **Get** /api/v1/cloud-ai/agents/{agent_id}/max-bots | Получение списка ботов MAX агента
*AIAgentIntegrationsAPI* | [**GetAgentBotTelegramBot**](docs/AIAgentIntegrationsAPI.md#getagentbottelegrambot) | **Get** /api/v1/cloud-ai/bots/{bot_id} | Получение информации о Telegram-боте
*AIAgentIntegrationsAPI* | [**GetAgentBotTelegramBots**](docs/AIAgentIntegrationsAPI.md#getagentbottelegrambots) | **Get** /api/v1/cloud-ai/bots | Получение списка Telegram-ботов
*AIAgentIntegrationsAPI* | [**GetAgentBotTelegramBotsByAgent**](docs/AIAgentIntegrationsAPI.md#getagentbottelegrambotsbyagent) | **Get** /api/v1/cloud-ai/agents/{agent_id}/bots | Получение списка Telegram-ботов агента
*AIAgentIntegrationsAPI* | [**GetAgentIntegrationCrisp**](docs/AIAgentIntegrationsAPI.md#getagentintegrationcrisp) | **Get** /api/v1/cloud-ai/agent/{agent_id}/crisp | Получение интеграции агента с Crisp
*AIAgentIntegrationsAPI* | [**GetAgentIntegrationJivo**](docs/AIAgentIntegrationsAPI.md#getagentintegrationjivo) | **Get** /api/v1/cloud-ai/agent/{agent_id}/jivo | Получение интеграции агента с Jivo
*AIAgentIntegrationsAPI* | [**UpdateAgentBotMaxBot**](docs/AIAgentIntegrationsAPI.md#updateagentbotmaxbot) | **Patch** /api/v1/cloud-ai/max-bots/{bot_id} | Изменение бота MAX
*AIAgentIntegrationsAPI* | [**UpdateAgentBotTelegramBot**](docs/AIAgentIntegrationsAPI.md#updateagentbottelegrambot) | **Patch** /api/v1/cloud-ai/bots/{bot_id} | Изменение Telegram-бота
*AIAgentSkillsAPI* | [**ApplyAgentSkillCatalogUpdate**](docs/AIAgentSkillsAPI.md#applyagentskillcatalogupdate) | **Post** /api/v1/cloud-ai/skills/{skill_id}/apply-catalog-update | Применение обновления навыка из каталога
*AIAgentSkillsAPI* | [**ConnectAgentSkillsToAgent**](docs/AIAgentSkillsAPI.md#connectagentskillstoagent) | **Post** /api/v1/cloud-ai/agents/{agent_id}/skills/connect | Подключение нескольких навыков к агенту
*AIAgentSkillsAPI* | [**CreateAgentSkill**](docs/AIAgentSkillsAPI.md#createagentskill) | **Post** /api/v1/cloud-ai/skills | Создание навыка
*AIAgentSkillsAPI* | [**DeleteAgentSkill**](docs/AIAgentSkillsAPI.md#deleteagentskill) | **Delete** /api/v1/cloud-ai/skills/{skill_id} | Удаление навыка
*AIAgentSkillsAPI* | [**DisconnectAgentSkillFromAgent**](docs/AIAgentSkillsAPI.md#disconnectagentskillfromagent) | **Delete** /api/v1/cloud-ai/agents/{agent_id}/skills/{skill_id} | Отключение навыка от агента
*AIAgentSkillsAPI* | [**DuplicateAgentSkill**](docs/AIAgentSkillsAPI.md#duplicateagentskill) | **Post** /api/v1/cloud-ai/skills/{skill_id}/duplicate | Копирование навыка
*AIAgentSkillsAPI* | [**GetAgentSkill**](docs/AIAgentSkillsAPI.md#getagentskill) | **Get** /api/v1/cloud-ai/skills/{skill_id} | Получение навыка
*AIAgentSkillsAPI* | [**GetAgentSkillCatalog**](docs/AIAgentSkillsAPI.md#getagentskillcatalog) | **Get** /api/v1/cloud-ai/skill-catalog | Получение каталога навыков
*AIAgentSkillsAPI* | [**GetAgentSkillLibrary**](docs/AIAgentSkillsAPI.md#getagentskilllibrary) | **Get** /api/v1/cloud-ai/skills | Получение списка навыков
*AIAgentSkillsAPI* | [**GetAgentSkillVersion**](docs/AIAgentSkillsAPI.md#getagentskillversion) | **Get** /api/v1/cloud-ai/skills/{skill_id}/versions/{version} | Получение версии навыка
*AIAgentSkillsAPI* | [**GetAgentSkillVersions**](docs/AIAgentSkillsAPI.md#getagentskillversions) | **Get** /api/v1/cloud-ai/skills/{skill_id}/versions | Получение истории версий навыка
*AIAgentSkillsAPI* | [**GetAgentSkillsForAgent**](docs/AIAgentSkillsAPI.md#getagentskillsforagent) | **Get** /api/v1/cloud-ai/agents/{agent_id}/skills | Получение навыков агента
*AIAgentSkillsAPI* | [**InstallAgentSkillFromCatalog**](docs/AIAgentSkillsAPI.md#installagentskillfromcatalog) | **Post** /api/v1/cloud-ai/skill-catalog/{catalog_item_id}/install | Установка навыка из каталога
*AIAgentSkillsAPI* | [**InstallAndConnectAgentSkill**](docs/AIAgentSkillsAPI.md#installandconnectagentskill) | **Post** /api/v1/cloud-ai/agents/{agent_id}/skills/install-and-connect | Установка навыка из каталога и включение у агента
*AIAgentSkillsAPI* | [**RollbackAgentSkillVersion**](docs/AIAgentSkillsAPI.md#rollbackagentskillversion) | **Post** /api/v1/cloud-ai/skills/{skill_id}/versions/{version}/rollback | Переключение навыка на версию из истории
*AIAgentSkillsAPI* | [**SetAgentSkillForAgent**](docs/AIAgentSkillsAPI.md#setagentskillforagent) | **Put** /api/v1/cloud-ai/agents/{agent_id}/skills/{skill_id} | Включение и выключение навыка у агента
*AIAgentSkillsAPI* | [**SyncAgentSkillAgentsVersions**](docs/AIAgentSkillsAPI.md#syncagentskillagentsversions) | **Post** /api/v1/cloud-ai/skills/{skill_id}/agents/sync-versions | Перевод всех агентов навыка на текущую версию
*AIAgentSkillsAPI* | [**SyncAgentSkillVersionsForAgent**](docs/AIAgentSkillsAPI.md#syncagentskillversionsforagent) | **Post** /api/v1/cloud-ai/agents/{agent_id}/skills/sync-versions | Перевод всех навыков агента на текущие версии
*AIAgentSkillsAPI* | [**UpdateAgentSkill**](docs/AIAgentSkillsAPI.md#updateagentskill) | **Patch** /api/v1/cloud-ai/skills/{skill_id} | Изменение навыка
*AIAgentSkillsAPI* | [**UpdateAgentSkillVersionForAgent**](docs/AIAgentSkillsAPI.md#updateagentskillversionforagent) | **Post** /api/v1/cloud-ai/agents/{agent_id}/skills/{skill_id}/update-version | Перевод агента на текущую версию навыка
*AIAgentToolsAPI* | [**ConnectAgentToolToAgent**](docs/AIAgentToolsAPI.md#connectagenttooltoagent) | **Post** /api/v1/cloud-ai/agents/{agent_id}/tools | Подключение MCP-инструмента к агенту
*AIAgentToolsAPI* | [**CreateAgentTool**](docs/AIAgentToolsAPI.md#createagenttool) | **Post** /api/v1/cloud-ai/tools | Подключение собственного MCP-сервера
*AIAgentToolsAPI* | [**DeleteAgentTool**](docs/AIAgentToolsAPI.md#deleteagenttool) | **Delete** /api/v1/cloud-ai/tools/{id} | Удаление MCP-инструмента
*AIAgentToolsAPI* | [**DisconnectAgentToolFromAgent**](docs/AIAgentToolsAPI.md#disconnectagenttoolfromagent) | **Delete** /api/v1/cloud-ai/agents/{agent_id}/tools/{tool_id} | Отключение MCP-инструмента от агента
*AIAgentToolsAPI* | [**GetAgentTool**](docs/AIAgentToolsAPI.md#getagenttool) | **Get** /api/v1/cloud-ai/tools/{id} | Получение MCP-инструмента
*AIAgentToolsAPI* | [**GetAgentToolOAuth2AuthorizeUrl**](docs/AIAgentToolsAPI.md#getagenttooloauth2authorizeurl) | **Get** /api/v1/cloud-ai/tools/oauth2/authorize-url/{tool_id} | Получение ссылки на OAuth2-авторизацию инструмента
*AIAgentToolsAPI* | [**GetAgentToolTemplatePrices**](docs/AIAgentToolsAPI.md#getagenttooltemplateprices) | **Get** /api/v1/cloud-ai/tools/templates/prices | Получение цен системных инструментов
*AIAgentToolsAPI* | [**GetAgentToolTemplates**](docs/AIAgentToolsAPI.md#getagenttooltemplates) | **Get** /api/v1/cloud-ai/tools/templates/marketplace | Получение шаблонов MCP-инструментов из маркетплейса
*AIAgentToolsAPI* | [**GetAgentTools**](docs/AIAgentToolsAPI.md#getagenttools) | **Get** /api/v1/cloud-ai/tools | Получение списка MCP-инструментов
*AIAgentToolsAPI* | [**GetAgentToolsByAgent**](docs/AIAgentToolsAPI.md#getagenttoolsbyagent) | **Get** /api/v1/cloud-ai/agents/{agent_id}/tools | Получение MCP-инструментов агента
*AIAgentToolsAPI* | [**InstallAgentToolFromTemplate**](docs/AIAgentToolsAPI.md#installagenttoolfromtemplate) | **Post** /api/v1/cloud-ai/tools/install | Установка MCP-инструмента из маркетплейса
*AIAgentToolsAPI* | [**RefreshAgentToolMethods**](docs/AIAgentToolsAPI.md#refreshagenttoolmethods) | **Post** /api/v1/cloud-ai/tools/{id}/refresh-methods | Обновление списка методов MCP-инструмента
*AIAgentToolsAPI* | [**UpdateAgentTool**](docs/AIAgentToolsAPI.md#updateagenttool) | **Put** /api/v1/cloud-ai/tools/{id} | Изменение MCP-инструмента
*AIAgentToolsAPI* | [**UpdateAgentToolConnection**](docs/AIAgentToolsAPI.md#updateagenttoolconnection) | **Put** /api/v1/cloud-ai/agents/{agent_id}/tools/{tool_id} | Изменение разрешённых методов MCP-инструмента у агента
*AIAgentsAPI* | [**AddAdditionalTokenPackage**](docs/AIAgentsAPI.md#addadditionaltokenpackage) | **Post** /api/v1/cloud-ai/agents/{id}/add-additional-token-package | Добавление дополнительного пакета токенов
*AIAgentsAPI* | [**ChangeAgentModel**](docs/AIAgentsAPI.md#changeagentmodel) | **Patch** /api/v1/cloud-ai/agents/{id}/model | Смена модели AI агента
*AIAgentsAPI* | [**CreateAgent**](docs/AIAgentsAPI.md#createagent) | **Post** /api/v1/cloud-ai/agents | Создание AI агента
*AIAgentsAPI* | [**CreateAgentEntityLimit**](docs/AIAgentsAPI.md#createagententitylimit) | **Post** /api/v1/cloud-ai/agents/{id}/entity-limits | Создание лимита агента
*AIAgentsAPI* | [**CreateAgentV2**](docs/AIAgentsAPI.md#createagentv2) | **Post** /api/v2/cloud-ai/agents | Создание AI агента
*AIAgentsAPI* | [**DeleteAgent**](docs/AIAgentsAPI.md#deleteagent) | **Delete** /api/v1/cloud-ai/agents/{id} | Удаление AI агента
*AIAgentsAPI* | [**DeleteAgentEntityLimit**](docs/AIAgentsAPI.md#deleteagententitylimit) | **Delete** /api/v1/cloud-ai/agents/{id}/entity-limits/{limit_id} | Удаление лимита агента
*AIAgentsAPI* | [**ExportAgentChatsCsv**](docs/AIAgentsAPI.md#exportagentchatscsv) | **Get** /api/v1/cloud-ai/agents/{id}/chats/docs | Выгрузка чатов агента в CSV
*AIAgentsAPI* | [**GetAgent**](docs/AIAgentsAPI.md#getagent) | **Get** /api/v1/cloud-ai/agents/{id} | Получение AI агента
*AIAgentsAPI* | [**GetAgentChatList**](docs/AIAgentsAPI.md#getagentchatlist) | **Get** /api/v1/cloud-ai/agents/{id}/chat-list | Получение списка чатов агента
*AIAgentsAPI* | [**GetAgentChatMessages**](docs/AIAgentsAPI.md#getagentchatmessages) | **Get** /api/v1/cloud-ai/agents/{id}/chat-list/{chat_id}/messages | Получение сообщений чата агента
*AIAgentsAPI* | [**GetAgentEntityLimit**](docs/AIAgentsAPI.md#getagententitylimit) | **Get** /api/v1/cloud-ai/agents/{id}/entity-limits/{limit_id} | Получение лимита агента
*AIAgentsAPI* | [**GetAgentEntityLimits**](docs/AIAgentsAPI.md#getagententitylimits) | **Get** /api/v1/cloud-ai/agents/{id}/entity-limits | Получение лимитов агента
*AIAgentsAPI* | [**GetAgentStatistics**](docs/AIAgentsAPI.md#getagentstatistics) | **Get** /api/v1/cloud-ai/agents/{id}/statistic | Получение статистики использования токенов агента
*AIAgentsAPI* | [**GetAgentV3**](docs/AIAgentsAPI.md#getagentv3) | **Get** /api/v3/cloud-ai/agents/{id} | Получение AI агента (v3)
*AIAgentsAPI* | [**GetAgents**](docs/AIAgentsAPI.md#getagents) | **Get** /api/v1/cloud-ai/agents | Получение списка AI агентов
*AIAgentsAPI* | [**GetAgentsV3**](docs/AIAgentsAPI.md#getagentsv3) | **Get** /api/v3/cloud-ai/agents | Получение списка AI агентов (v3)
*AIAgentsAPI* | [**SyncAgentEntityLimits**](docs/AIAgentsAPI.md#syncagententitylimits) | **Put** /api/v1/cloud-ai/agents/{id}/entity-limits/sync | Синхронизация лимитов агента за период
*AIAgentsAPI* | [**UpdateAgent**](docs/AIAgentsAPI.md#updateagent) | **Patch** /api/v1/cloud-ai/agents/{id} | Обновление AI агента
*AIAgentsAPI* | [**UpdateAgentEntityLimit**](docs/AIAgentsAPI.md#updateagententitylimit) | **Patch** /api/v1/cloud-ai/agents/{id}/entity-limits/{limit_id} | Изменение лимита агента
*AIGatewayAPI* | [**CreateAiKey**](docs/AIGatewayAPI.md#createaikey) | **Post** /api/v1/cloud-ai/ai-keys | Создание ключа AI Gateway
*AIGatewayAPI* | [**CreateAiKeyLimit**](docs/AIGatewayAPI.md#createaikeylimit) | **Post** /api/v1/cloud-ai/ai-keys/{key_id}/entity-limits | Создание лимита ключа AI Gateway
*AIGatewayAPI* | [**DeleteAiKey**](docs/AIGatewayAPI.md#deleteaikey) | **Delete** /api/v1/cloud-ai/ai-keys/{key_id} | Удаление ключа AI Gateway
*AIGatewayAPI* | [**DeleteAiKeyLimit**](docs/AIGatewayAPI.md#deleteaikeylimit) | **Delete** /api/v1/cloud-ai/ai-keys/{key_id}/entity-limits/{limit_id} | Удаление лимита ключа AI Gateway
*AIGatewayAPI* | [**GetAiKeyCallsStatistic**](docs/AIGatewayAPI.md#getaikeycallsstatistic) | **Get** /api/v1/cloud-ai/ai-keys/statistic/calls | Получение статистики запросов ключа AI Gateway
*AIGatewayAPI* | [**GetAiKeyLimit**](docs/AIGatewayAPI.md#getaikeylimit) | **Get** /api/v1/cloud-ai/ai-keys/{key_id}/entity-limits/{limit_id} | Получение лимита ключа AI Gateway
*AIGatewayAPI* | [**GetAiKeyLimits**](docs/AIGatewayAPI.md#getaikeylimits) | **Get** /api/v1/cloud-ai/ai-keys/{key_id}/entity-limits | Получение списка лимитов ключа AI Gateway
*AIGatewayAPI* | [**GetAiKeyModelsStatistic**](docs/AIGatewayAPI.md#getaikeymodelsstatistic) | **Get** /api/v1/cloud-ai/ai-keys/statistic/models | Получение списка моделей, использованных ключом AI Gateway
*AIGatewayAPI* | [**GetAiKeyTokenUsageStatistic**](docs/AIGatewayAPI.md#getaikeytokenusagestatistic) | **Get** /api/v1/cloud-ai/ai-keys/statistic/token-usage | Получение статистики расхода токенов ключа AI Gateway
*AIGatewayAPI* | [**GetAiKeys**](docs/AIGatewayAPI.md#getaikeys) | **Get** /api/v1/cloud-ai/ai-keys | Получение списка ключей AI Gateway
*AIGatewayAPI* | [**RegenerateAiKey**](docs/AIGatewayAPI.md#regenerateaikey) | **Put** /api/v1/cloud-ai/ai-keys/{key_id}/regenerate | Перевыпуск ключа AI Gateway
*AIGatewayAPI* | [**SyncAiKeyLimits**](docs/AIGatewayAPI.md#syncaikeylimits) | **Put** /api/v1/cloud-ai/ai-keys/{key_id}/entity-limits/sync | Синхронизация лимитов ключа AI Gateway
*AIGatewayAPI* | [**UpdateAiKey**](docs/AIGatewayAPI.md#updateaikey) | **Patch** /api/v1/cloud-ai/ai-keys/{key_id} | Изменение ключа AI Gateway
*AIGatewayAPI* | [**UpdateAiKeyLimit**](docs/AIGatewayAPI.md#updateaikeylimit) | **Patch** /api/v1/cloud-ai/ai-keys/{key_id}/entity-limits/{limit_id} | Изменение лимита ключа AI Gateway
*AIGatewayInferenceAPI* | [**CancelAiGatewayBatch**](docs/AIGatewayInferenceAPI.md#cancelaigatewaybatch) | **Post** /v1/batches/{batch_id}/cancel | Отмена батча
*AIGatewayInferenceAPI* | [**CancelAiGatewayResponse**](docs/AIGatewayInferenceAPI.md#cancelaigatewayresponse) | **Post** /v1/responses/{response_id}/cancel | Отмена фонового ответа (Responses API)
*AIGatewayInferenceAPI* | [**CreateAiGatewayBatch**](docs/AIGatewayInferenceAPI.md#createaigatewaybatch) | **Post** /v1/batches | Создание батча
*AIGatewayInferenceAPI* | [**CreateAiGatewayChatCompletion**](docs/AIGatewayInferenceAPI.md#createaigatewaychatcompletion) | **Post** /v1/chat/completions | Генерация ответа в диалоге (Chat Completions)
*AIGatewayInferenceAPI* | [**CreateAiGatewayCompletion**](docs/AIGatewayInferenceAPI.md#createaigatewaycompletion) | **Post** /v1/completions | Продолжение текста (Completions)
*AIGatewayInferenceAPI* | [**CreateAiGatewayEmbedding**](docs/AIGatewayInferenceAPI.md#createaigatewayembedding) | **Post** /v1/embeddings | Получение эмбеддингов
*AIGatewayInferenceAPI* | [**CreateAiGatewayImage**](docs/AIGatewayInferenceAPI.md#createaigatewayimage) | **Post** /v1/images/generations | Генерация изображений
*AIGatewayInferenceAPI* | [**CreateAiGatewayImageEdit**](docs/AIGatewayInferenceAPI.md#createaigatewayimageedit) | **Post** /v1/images/edits | Редактирование изображений
*AIGatewayInferenceAPI* | [**CreateAiGatewayResponse**](docs/AIGatewayInferenceAPI.md#createaigatewayresponse) | **Post** /v1/responses | Создание ответа (Responses API)
*AIGatewayInferenceAPI* | [**CreateAiGatewaySpeech**](docs/AIGatewayInferenceAPI.md#createaigatewayspeech) | **Post** /v1/audio/speech | Синтез речи
*AIGatewayInferenceAPI* | [**CreateAiGatewayTranscription**](docs/AIGatewayInferenceAPI.md#createaigatewaytranscription) | **Post** /v1/audio/transcriptions | Распознавание речи
*AIGatewayInferenceAPI* | [**DeleteAiGatewayFile**](docs/AIGatewayInferenceAPI.md#deleteaigatewayfile) | **Delete** /v1/files/{file_id} | Удаление файла
*AIGatewayInferenceAPI* | [**DeleteAiGatewayResponse**](docs/AIGatewayInferenceAPI.md#deleteaigatewayresponse) | **Delete** /v1/responses/{response_id} | Удаление ответа (Responses API)
*AIGatewayInferenceAPI* | [**DownloadAiGatewayFileContent**](docs/AIGatewayInferenceAPI.md#downloadaigatewayfilecontent) | **Get** /v1/files/{file_id}/content | Скачивание файла
*AIGatewayInferenceAPI* | [**GetAiGatewayBatch**](docs/AIGatewayInferenceAPI.md#getaigatewaybatch) | **Get** /v1/batches/{batch_id} | Получение батча
*AIGatewayInferenceAPI* | [**GetAiGatewayBatches**](docs/AIGatewayInferenceAPI.md#getaigatewaybatches) | **Get** /v1/batches | Получение списка батчей
*AIGatewayInferenceAPI* | [**GetAiGatewayFile**](docs/AIGatewayInferenceAPI.md#getaigatewayfile) | **Get** /v1/files/{file_id} | Получение информации о файле
*AIGatewayInferenceAPI* | [**GetAiGatewayModels**](docs/AIGatewayInferenceAPI.md#getaigatewaymodels) | **Get** /v1/models | Получение списка моделей AI Gateway
*AIGatewayInferenceAPI* | [**GetAiGatewayResponse**](docs/AIGatewayInferenceAPI.md#getaigatewayresponse) | **Get** /v1/responses/{response_id} | Получение ответа (Responses API)
*AIGatewayInferenceAPI* | [**UploadAiGatewayFile**](docs/AIGatewayInferenceAPI.md#uploadaigatewayfile) | **Post** /v1/files | Загрузка файла
*AIModelsAPI* | [**GetAgentsTokenPackages**](docs/AIModelsAPI.md#getagentstokenpackages) | **Get** /api/v1/cloud-ai/token-packages/agents | Получение списка пакетов токенов для агентов
*AIModelsAPI* | [**GetAiModelParameters**](docs/AIModelsAPI.md#getaimodelparameters) | **Get** /api/v3/cloud-ai/model-parameters | Получение справочника характеристик моделей
*AIModelsAPI* | [**GetAiModelTokenPackages**](docs/AIModelsAPI.md#getaimodeltokenpackages) | **Get** /api/v1/cloud-ai/models/{id}/token-packages | Получение пакетов токенов модели
*AIModelsAPI* | [**GetAiProvider**](docs/AIModelsAPI.md#getaiprovider) | **Get** /api/v1/cloud-ai/providers/{id} | Получение провайдера моделей
*AIModelsAPI* | [**GetAiProviders**](docs/AIModelsAPI.md#getaiproviders) | **Get** /api/v1/cloud-ai/providers | Получение списка провайдеров моделей
*AIModelsAPI* | [**GetKnowledgebasesTokenPackages**](docs/AIModelsAPI.md#getknowledgebasestokenpackages) | **Get** /api/v1/cloud-ai/token-packages/knowledge-bases | Получение списка пакетов токенов для баз знаний
*AIModelsAPI* | [**GetModels**](docs/AIModelsAPI.md#getmodels) | **Get** /api/v1/cloud-ai/models | Получение списка моделей (v1)
*AIModelsAPI* | [**GetModelsV3**](docs/AIModelsAPI.md#getmodelsv3) | **Get** /api/v3/cloud-ai/models | Получение списка моделей
*AISolutionsAPI* | [**AddAiSolutionTelegramBot**](docs/AISolutionsAPI.md#addaisolutiontelegrambot) | **Post** /api/v1/cloud-ai/solutions/{solution_id}/telegram/bots | Подключение Telegram-бота к AI-решению
*AISolutionsAPI* | [**ApproveAiSolutionTelegramPairing**](docs/AISolutionsAPI.md#approveaisolutiontelegrampairing) | **Post** /api/v1/cloud-ai/solutions/{solution_id}/telegram/bots/{account_key}/pairing/approve | Подтверждение владельца Telegram-бота AI-решения
*AISolutionsAPI* | [**CreateAiSolution**](docs/AISolutionsAPI.md#createaisolution) | **Post** /api/v1/cloud-ai/solutions | Создание AI-решения
*AISolutionsAPI* | [**DeleteAiSolution**](docs/AISolutionsAPI.md#deleteaisolution) | **Delete** /api/v1/cloud-ai/solutions/{solution_id} | Удаление AI-решения
*AISolutionsAPI* | [**GetAiSolution**](docs/AISolutionsAPI.md#getaisolution) | **Get** /api/v1/cloud-ai/solutions/{solution_id} | Получение AI-решения
*AISolutionsAPI* | [**GetAiSolutionChatFullHistory**](docs/AISolutionsAPI.md#getaisolutionchatfullhistory) | **Get** /api/v1/cloud-ai/solutions/{solution_id}/chat/history | Получение всей истории чата AI-решения
*AISolutionsAPI* | [**GetAiSolutionChatHistory**](docs/AISolutionsAPI.md#getaisolutionchathistory) | **Get** /api/v1/cloud-ai/solutions/{solution_id}/chat/conversations/{conversation_id}/messages | Получение истории разговора AI-решения
*AISolutionsAPI* | [**GetAiSolutionChatStatus**](docs/AISolutionsAPI.md#getaisolutionchatstatus) | **Get** /api/v1/cloud-ai/solutions/{solution_id}/chat/status | Получение статуса чата AI-решения
*AISolutionsAPI* | [**GetAiSolutionMainChatHistory**](docs/AISolutionsAPI.md#getaisolutionmainchathistory) | **Get** /api/v1/cloud-ai/solutions/{solution_id}/chat/messages | Получение истории основного разговора AI-решения
*AISolutionsAPI* | [**ListAiSolutionConversations**](docs/AISolutionsAPI.md#listaisolutionconversations) | **Get** /api/v1/cloud-ai/solutions/{solution_id}/chat/conversations | Получение списка разговоров AI-решения
*AISolutionsAPI* | [**ListAiSolutionTelegramBots**](docs/AISolutionsAPI.md#listaisolutiontelegrambots) | **Get** /api/v1/cloud-ai/solutions/{solution_id}/telegram/bots | Получение списка Telegram-ботов AI-решения
*AISolutionsAPI* | [**ListAiSolutionTypes**](docs/AISolutionsAPI.md#listaisolutiontypes) | **Get** /api/v1/cloud-ai/solutions/types | Получение списка типов AI-решений
*AISolutionsAPI* | [**ListAiSolutions**](docs/AISolutionsAPI.md#listaisolutions) | **Get** /api/v1/cloud-ai/solutions | Получение списка AI-решений
*AISolutionsAPI* | [**RegenerateAiSolutionAiKey**](docs/AISolutionsAPI.md#regenerateaisolutionaikey) | **Post** /api/v1/cloud-ai/solutions/{solution_id}/ai-key/regenerate | Перевыпуск ключа AI Gateway решения
*AISolutionsAPI* | [**RemoveAiSolutionTelegramBot**](docs/AISolutionsAPI.md#removeaisolutiontelegrambot) | **Delete** /api/v1/cloud-ai/solutions/{solution_id}/telegram/bots/{account_key} | Отключение Telegram-бота от AI-решения
*AISolutionsAPI* | [**ResetAiSolutionChat**](docs/AISolutionsAPI.md#resetaisolutionchat) | **Post** /api/v1/cloud-ai/solutions/{solution_id}/chat/reset | Сброс контекста разговора AI-решения
*AISolutionsAPI* | [**SendAiSolutionMessage**](docs/AISolutionsAPI.md#sendaisolutionmessage) | **Post** /api/v1/cloud-ai/solutions/{solution_id}/chat/messages | Отправка сообщения агенту AI-решения
*AISolutionsAPI* | [**ShutdownAiSolution**](docs/AISolutionsAPI.md#shutdownaisolution) | **Post** /api/v1/cloud-ai/solutions/{solution_id}/shutdown | Выключение AI-решения
*AISolutionsAPI* | [**StartAiSolution**](docs/AISolutionsAPI.md#startaisolution) | **Post** /api/v1/cloud-ai/solutions/{solution_id}/start | Запуск AI-решения
*AISolutionsAPI* | [**StopAiSolutionChat**](docs/AISolutionsAPI.md#stopaisolutionchat) | **Post** /api/v1/cloud-ai/solutions/{solution_id}/chat/stop | Остановка хода агента AI-решения
*AISolutionsAPI* | [**StreamAiSolutionChat**](docs/AISolutionsAPI.md#streamaisolutionchat) | **Post** /api/v1/cloud-ai/solutions/{solution_id}/chat/stream | Отправка сообщения агенту AI-решения с потоковым ответом
*AISolutionsAPI* | [**UpdateAiSolution**](docs/AISolutionsAPI.md#updateaisolution) | **Patch** /api/v1/cloud-ai/solutions/{solution_id} | Изменение AI-решения
*AccountAPI* | [**AddCountriesToAllowedList**](docs/AccountAPI.md#addcountriestoallowedlist) | **Post** /api/v1/auth/access/countries | Добавление стран в список разрешенных
*AccountAPI* | [**AddIPsToAllowedList**](docs/AccountAPI.md#addipstoallowedlist) | **Post** /api/v1/auth/access/ips | Добавление IP-адресов в список разрешенных
*AccountAPI* | [**DeleteCountriesFromAllowedList**](docs/AccountAPI.md#deletecountriesfromallowedlist) | **Delete** /api/v1/auth/access/countries | Удаление стран из списка разрешенных
*AccountAPI* | [**DeleteIPsFromAllowedList**](docs/AccountAPI.md#deleteipsfromallowedlist) | **Delete** /api/v1/auth/access/ips | Удаление IP-адресов из списка разрешенных
*AccountAPI* | [**GetAccountStatus**](docs/AccountAPI.md#getaccountstatus) | **Get** /api/v1/account/status | Получение статуса аккаунта
*AccountAPI* | [**GetAuthAccessSettings**](docs/AccountAPI.md#getauthaccesssettings) | **Get** /api/v1/auth/access | Получить информацию о ограничениях авторизации пользователя
*AccountAPI* | [**GetCountries**](docs/AccountAPI.md#getcountries) | **Get** /api/v1/auth/access/countries | Получение списка стран
*AccountAPI* | [**GetNotificationSettings**](docs/AccountAPI.md#getnotificationsettings) | **Get** /api/v1/account/notification-settings | Получение настроек уведомлений аккаунта
*AccountAPI* | [**UpdateAuthRestrictionsByCountries**](docs/AccountAPI.md#updateauthrestrictionsbycountries) | **Post** /api/v1/auth/access/countries/enabled | Включение/отключение ограничений по стране
*AccountAPI* | [**UpdateAuthRestrictionsByIP**](docs/AccountAPI.md#updateauthrestrictionsbyip) | **Post** /api/v1/auth/access/ips/enabled | Включение/отключение ограничений по IP-адресу
*AccountAPI* | [**UpdateNotificationSettings**](docs/AccountAPI.md#updatenotificationsettings) | **Patch** /api/v1/account/notification-settings | Изменение настроек уведомлений аккаунта
*AppsAPI* | [**AddProvider**](docs/AppsAPI.md#addprovider) | **Post** /api/v1/vcs-provider | Привязка vcs провайдера
*AppsAPI* | [**CreateApp**](docs/AppsAPI.md#createapp) | **Post** /api/v1/apps | Создание приложения
*AppsAPI* | [**CreateDeploy**](docs/AppsAPI.md#createdeploy) | **Post** /api/v1/apps/{app_id}/deploy | Запуск деплоя приложения
*AppsAPI* | [**DeleteApp**](docs/AppsAPI.md#deleteapp) | **Delete** /api/v1/apps/{app_id} | Удаление приложения
*AppsAPI* | [**DeleteProvider**](docs/AppsAPI.md#deleteprovider) | **Delete** /api/v1/vcs-provider/{provider_id} | Отвязка vcs провайдера от аккаунта
*AppsAPI* | [**DeployAction**](docs/AppsAPI.md#deployaction) | **Post** /api/v1/apps/{app_id}/deploy/{deploy_id}/stop | Остановка деплоя приложения
*AppsAPI* | [**GetApp**](docs/AppsAPI.md#getapp) | **Get** /api/v1/apps/{app_id} | Получение приложения по id
*AppsAPI* | [**GetAppDeploys**](docs/AppsAPI.md#getappdeploys) | **Get** /api/v1/apps/{app_id}/deploys | Получение списка деплоев приложения
*AppsAPI* | [**GetAppLogs**](docs/AppsAPI.md#getapplogs) | **Get** /api/v1/apps/{app_id}/logs | Получение логов приложения
*AppsAPI* | [**GetAppStatistics**](docs/AppsAPI.md#getappstatistics) | **Get** /api/v1/apps/{app_id}/statistics | Получение статистики приложения
*AppsAPI* | [**GetApps**](docs/AppsAPI.md#getapps) | **Get** /api/v1/apps | Получение списка приложений
*AppsAPI* | [**GetAppsPresets**](docs/AppsAPI.md#getappspresets) | **Get** /api/v1/presets/apps | Получение списка доступных тарифов для приложения
*AppsAPI* | [**GetBranches**](docs/AppsAPI.md#getbranches) | **Get** /api/v1/vcs-provider/{provider_id}/repository/{repository_id} | Получение списка веток репозитория
*AppsAPI* | [**GetCommits**](docs/AppsAPI.md#getcommits) | **Get** /api/v1/vcs-provider/{provider_id}/repository/{repository_id}/branch | Получение списка коммитов ветки репозитория
*AppsAPI* | [**GetDeployLogs**](docs/AppsAPI.md#getdeploylogs) | **Get** /api/v1/apps/{app_id}/deploy/{deploy_id}/logs | Получение логов деплоя приложения
*AppsAPI* | [**GetDeploySettings**](docs/AppsAPI.md#getdeploysettings) | **Get** /api/v1/deploy-settings/apps | Получение списка дефолтных настроек деплоя для приложения
*AppsAPI* | [**GetFrameworks**](docs/AppsAPI.md#getframeworks) | **Get** /api/v1/frameworks/apps | Получение списка доступных фреймворков для приложения
*AppsAPI* | [**GetProviders**](docs/AppsAPI.md#getproviders) | **Get** /api/v1/vcs-provider | Получение списка vcs провайдеров
*AppsAPI* | [**GetRepositories**](docs/AppsAPI.md#getrepositories) | **Get** /api/v1/vcs-provider/{provider_id} | Получение списка репозиториев vcs провайдера
*AppsAPI* | [**UpdateAppSettings**](docs/AppsAPI.md#updateappsettings) | **Patch** /api/v1/apps/{app_id} | Изменение настроек приложения
*AppsAPI* | [**UpdateAppState**](docs/AppsAPI.md#updateappstate) | **Patch** /api/v1/apps/{app_id}/action/{action} | Изменение состояния приложения
*BalancersAPI* | [**AddIPsToBalancer**](docs/BalancersAPI.md#addipstobalancer) | **Post** /api/v1/balancers/{balancer_id}/ips | Добавление IP-адресов к балансировщику
*BalancersAPI* | [**CreateBalancer**](docs/BalancersAPI.md#createbalancer) | **Post** /api/v1/balancers | Создание бaлансировщика
*BalancersAPI* | [**CreateBalancerRule**](docs/BalancersAPI.md#createbalancerrule) | **Post** /api/v1/balancers/{balancer_id}/rules | Создание правила для балансировщика
*BalancersAPI* | [**DeleteBalancer**](docs/BalancersAPI.md#deletebalancer) | **Delete** /api/v1/balancers/{balancer_id} | Удаление балансировщика
*BalancersAPI* | [**DeleteBalancerRule**](docs/BalancersAPI.md#deletebalancerrule) | **Delete** /api/v1/balancers/{balancer_id}/rules/{rule_id} | Удаление правила для балансировщика
*BalancersAPI* | [**DeleteIPsFromBalancer**](docs/BalancersAPI.md#deleteipsfrombalancer) | **Delete** /api/v1/balancers/{balancer_id}/ips | Удаление IP-адресов из балансировщика
*BalancersAPI* | [**GetBalancer**](docs/BalancersAPI.md#getbalancer) | **Get** /api/v1/balancers/{balancer_id} | Получение бaлансировщика
*BalancersAPI* | [**GetBalancerIPs**](docs/BalancersAPI.md#getbalancerips) | **Get** /api/v1/balancers/{balancer_id}/ips | Получение списка IP-адресов балансировщика
*BalancersAPI* | [**GetBalancerRules**](docs/BalancersAPI.md#getbalancerrules) | **Get** /api/v1/balancers/{balancer_id}/rules | Получение правил балансировщика
*BalancersAPI* | [**GetBalancers**](docs/BalancersAPI.md#getbalancers) | **Get** /api/v1/balancers | Получение списка всех бaлансировщиков
*BalancersAPI* | [**GetBalancersPresets**](docs/BalancersAPI.md#getbalancerspresets) | **Get** /api/v1/presets/balancers | Получение списка тарифов для балансировщика
*BalancersAPI* | [**UpdateBalancer**](docs/BalancersAPI.md#updatebalancer) | **Patch** /api/v1/balancers/{balancer_id} | Обновление балансировщика
*BalancersAPI* | [**UpdateBalancerRule**](docs/BalancersAPI.md#updatebalancerrule) | **Patch** /api/v1/balancers/{balancer_id}/rules/{rule_id} | Обновление правила для балансировщика
*CDNAPI* | [**AddCdnCertificate**](docs/CDNAPI.md#addcdncertificate) | **Post** /api/v1/cdn/certificates | Загрузка собственного сертификата CDN
*CDNAPI* | [**ArchiveCdnCertificateTask**](docs/CDNAPI.md#archivecdncertificatetask) | **Post** /api/v1/cdn/certificates/tasks/{task_id}/archive | Архивация задачи на выпуск сертификата
*CDNAPI* | [**ClearCdnResourceCache**](docs/CDNAPI.md#clearcdnresourcecache) | **Post** /api/v1/cdn/http-resources/{resource_id}/clear-cache | Очистка кэша CDN-ресурса
*CDNAPI* | [**CreateCdnResource**](docs/CDNAPI.md#createcdnresource) | **Post** /api/v1/cdn/http-resources | Создание CDN-ресурса
*CDNAPI* | [**DeleteCdnCertificate**](docs/CDNAPI.md#deletecdncertificate) | **Delete** /api/v1/cdn/certificates/{certificate_id} | Удаление сертификата CDN
*CDNAPI* | [**DeleteCdnResource**](docs/CDNAPI.md#deletecdnresource) | **Delete** /api/v1/cdn/http-resources/{resource_id} | Удаление CDN-ресурса
*CDNAPI* | [**GetCdnCertificateTasks**](docs/CDNAPI.md#getcdncertificatetasks) | **Get** /api/v1/cdn/certificates/tasks | Получение списка задач на выпуск сертификатов
*CDNAPI* | [**GetCdnCertificates**](docs/CDNAPI.md#getcdncertificates) | **Get** /api/v1/cdn/certificates | Получение списка сертификатов CDN
*CDNAPI* | [**GetCdnOriginNodes**](docs/CDNAPI.md#getcdnoriginnodes) | **Get** /api/v1/cdn/nodes/origin | Получение списка подсетей узлов CDN
*CDNAPI* | [**GetCdnPresets**](docs/CDNAPI.md#getcdnpresets) | **Get** /api/v1/cdn/presets | Получение списка тарифов CDN
*CDNAPI* | [**GetCdnResource**](docs/CDNAPI.md#getcdnresource) | **Get** /api/v1/cdn/http-resources/{resource_id} | Получение CDN-ресурса
*CDNAPI* | [**GetCdnResourceConfiguration**](docs/CDNAPI.md#getcdnresourceconfiguration) | **Get** /api/v1/cdn/http-resources/{resource_id}/configuration | Получение конфигурации CDN-ресурса
*CDNAPI* | [**GetCdnResourceNodes**](docs/CDNAPI.md#getcdnresourcenodes) | **Get** /api/v1/cdn/nodes/http-resources/{resource_id} | Получение списка раздающих узлов CDN-ресурса
*CDNAPI* | [**GetCdnResourceStatistics**](docs/CDNAPI.md#getcdnresourcestatistics) | **Get** /api/v1/cdn/http-resources/{resource_id}/statistics | Получение статистики CDN-ресурса
*CDNAPI* | [**GetCdnResources**](docs/CDNAPI.md#getcdnresources) | **Get** /api/v1/cdn/http-resources | Получение списка CDN-ресурсов
*CDNAPI* | [**IssueCdnCertificate**](docs/CDNAPI.md#issuecdncertificate) | **Post** /api/v1/cdn/certificates/issue | Выпуск сертификата Let&#39;s Encrypt для CDN-ресурса
*CDNAPI* | [**PreloadCdnResourceCache**](docs/CDNAPI.md#preloadcdnresourcecache) | **Post** /api/v1/cdn/http-resources/{resource_id}/preload-cache | Предварительная загрузка кэша CDN-ресурса
*CDNAPI* | [**ResumeCdnResource**](docs/CDNAPI.md#resumecdnresource) | **Post** /api/v1/cdn/http-resources/{resource_id}/resume | Возобновление раздачи CDN-ресурса
*CDNAPI* | [**SuspendCdnResource**](docs/CDNAPI.md#suspendcdnresource) | **Post** /api/v1/cdn/http-resources/{resource_id}/suspend | Приостановка раздачи CDN-ресурса
*CDNAPI* | [**UpdateCdnResource**](docs/CDNAPI.md#updatecdnresource) | **Patch** /api/v1/cdn/http-resources/{resource_id} | Изменение CDN-ресурса
*ContainerRegistryAPI* | [**CreateRegistry**](docs/ContainerRegistryAPI.md#createregistry) | **Post** /api/v1/container-registry | Создание реестра
*ContainerRegistryAPI* | [**DeleteRegistry**](docs/ContainerRegistryAPI.md#deleteregistry) | **Delete** /api/v1/container-registry/{registry_id} | Удаление реестра
*ContainerRegistryAPI* | [**GetRegistries**](docs/ContainerRegistryAPI.md#getregistries) | **Get** /api/v1/container-registry | Получение списка реестров контейнеров
*ContainerRegistryAPI* | [**GetRegistry**](docs/ContainerRegistryAPI.md#getregistry) | **Get** /api/v1/container-registry/{registry_id} | Получение информации о реестре
*ContainerRegistryAPI* | [**GetRegistryPresets**](docs/ContainerRegistryAPI.md#getregistrypresets) | **Get** /api/v1/container-registry/presets | Получение списка тарифов
*ContainerRegistryAPI* | [**GetRegistryRepositories**](docs/ContainerRegistryAPI.md#getregistryrepositories) | **Get** /api/v1/container-registry/{registry_id}/repositories | Получение списка репозиториев
*ContainerRegistryAPI* | [**UpdateRegistry**](docs/ContainerRegistryAPI.md#updateregistry) | **Patch** /api/v1/container-registry/{registry_id} | Обновление информации о реестре
*DatabasesAPI* | [**CreateDatabaseBackup**](docs/DatabasesAPI.md#createdatabasebackup) | **Post** /api/v1/dbs/{db_id}/backups | Создание бэкапа базы данных
*DatabasesAPI* | [**CreateDatabaseBackupDownloadUrl**](docs/DatabasesAPI.md#createdatabasebackupdownloadurl) | **Post** /api/v1/dbs/{db_id}/backups/{backup_id}/download-url | Получение ссылки для скачивания бэкапа базы данных
*DatabasesAPI* | [**CreateDatabaseCluster**](docs/DatabasesAPI.md#createdatabasecluster) | **Post** /api/v1/databases | Создание кластера базы данных
*DatabasesAPI* | [**CreateDatabaseInstance**](docs/DatabasesAPI.md#createdatabaseinstance) | **Post** /api/v1/databases/{db_cluster_id}/instances | Создание инстанса базы данных
*DatabasesAPI* | [**CreateDatabaseS3Backup**](docs/DatabasesAPI.md#createdatabases3backup) | **Post** /api/v2/databases/{db_id}/backups | Создание S3-бэкапа базы данных
*DatabasesAPI* | [**CreateDatabaseUser**](docs/DatabasesAPI.md#createdatabaseuser) | **Post** /api/v1/databases/{db_cluster_id}/admins | Создание пользователя базы данных
*DatabasesAPI* | [**DeleteDatabaseBackup**](docs/DatabasesAPI.md#deletedatabasebackup) | **Delete** /api/v1/dbs/{db_id}/backups/{backup_id} | Удаление бэкапа базы данных
*DatabasesAPI* | [**DeleteDatabaseCluster**](docs/DatabasesAPI.md#deletedatabasecluster) | **Delete** /api/v1/databases/{db_cluster_id} | Удаление кластера базы данных
*DatabasesAPI* | [**DeleteDatabaseInstance**](docs/DatabasesAPI.md#deletedatabaseinstance) | **Delete** /api/v1/databases/{db_cluster_id}/instances/{instance_id} | Удаление инстанса базы данных
*DatabasesAPI* | [**DeleteDatabaseS3Backup**](docs/DatabasesAPI.md#deletedatabases3backup) | **Delete** /api/v2/databases/{db_id}/backups/{backup_id} | Удаление S3-бэкапа базы данных
*DatabasesAPI* | [**DeleteDatabaseUser**](docs/DatabasesAPI.md#deletedatabaseuser) | **Delete** /api/v1/databases/{db_cluster_id}/admins/{admin_id} | Удаление пользователя базы данных
*DatabasesAPI* | [**GetDatabaseAutoBackupsSettings**](docs/DatabasesAPI.md#getdatabaseautobackupssettings) | **Get** /api/v1/dbs/{db_id}/auto-backups | Получение настроек автобэкапов базы данных
*DatabasesAPI* | [**GetDatabaseBackup**](docs/DatabasesAPI.md#getdatabasebackup) | **Get** /api/v1/dbs/{db_id}/backups/{backup_id} | Получение бэкапа базы данных
*DatabasesAPI* | [**GetDatabaseBackups**](docs/DatabasesAPI.md#getdatabasebackups) | **Get** /api/v1/dbs/{db_id}/backups | Список бэкапов базы данных
*DatabasesAPI* | [**GetDatabaseCluster**](docs/DatabasesAPI.md#getdatabasecluster) | **Get** /api/v1/databases/{db_cluster_id} | Получение кластера базы данных
*DatabasesAPI* | [**GetDatabaseClusterReplicas**](docs/DatabasesAPI.md#getdatabaseclusterreplicas) | **Get** /api/v1/databases/{db_cluster_id}/replicas | Получение списка реплик кластера базы данных
*DatabasesAPI* | [**GetDatabaseClusterTypes**](docs/DatabasesAPI.md#getdatabaseclustertypes) | **Get** /api/v1/database-types | Получение списка типов кластеров баз данных
*DatabasesAPI* | [**GetDatabaseClusters**](docs/DatabasesAPI.md#getdatabaseclusters) | **Get** /api/v1/databases | Получение списка кластеров баз данных
*DatabasesAPI* | [**GetDatabaseConfigurators**](docs/DatabasesAPI.md#getdatabaseconfigurators) | **Get** /api/v1/configurator/databases | Получение списка конфигураторов баз данных
*DatabasesAPI* | [**GetDatabaseDefaultParameters**](docs/DatabasesAPI.md#getdatabasedefaultparameters) | **Get** /api/v1/dbs/default-parameters | Получение рекомендуемых значений параметров баз данных
*DatabasesAPI* | [**GetDatabaseInstance**](docs/DatabasesAPI.md#getdatabaseinstance) | **Get** /api/v1/databases/{db_cluster_id}/instances/{instance_id} | Получение инстанса базы данных
*DatabasesAPI* | [**GetDatabaseInstances**](docs/DatabasesAPI.md#getdatabaseinstances) | **Get** /api/v1/databases/{db_cluster_id}/instances | Получение списка инстансов баз данных
*DatabasesAPI* | [**GetDatabaseParameters**](docs/DatabasesAPI.md#getdatabaseparameters) | **Get** /api/v1/dbs/parameters | Получение списка параметров баз данных
*DatabasesAPI* | [**GetDatabasePreset**](docs/DatabasesAPI.md#getdatabasepreset) | **Get** /api/v2/dbs/presets/{preset_id} | Получение тарифа для базы данных
*DatabasesAPI* | [**GetDatabasePrivileges**](docs/DatabasesAPI.md#getdatabaseprivileges) | **Get** /api/v1/databases/{db_cluster_id}/privileges | Получение привилегий кластера базы данных
*DatabasesAPI* | [**GetDatabaseS3Backup**](docs/DatabasesAPI.md#getdatabases3backup) | **Get** /api/v2/databases/{db_id}/backups/{backup_id} | Получение S3-бэкапа базы данных
*DatabasesAPI* | [**GetDatabaseS3Backups**](docs/DatabasesAPI.md#getdatabases3backups) | **Get** /api/v2/databases/{db_id}/backups | Список S3-бэкапов базы данных
*DatabasesAPI* | [**GetDatabaseUser**](docs/DatabasesAPI.md#getdatabaseuser) | **Get** /api/v1/databases/{db_cluster_id}/admins/{admin_id} | Получение пользователя базы данных
*DatabasesAPI* | [**GetDatabaseUsers**](docs/DatabasesAPI.md#getdatabaseusers) | **Get** /api/v1/databases/{db_cluster_id}/admins | Получение списка пользователей базы данных
*DatabasesAPI* | [**GetDatabasesPresets**](docs/DatabasesAPI.md#getdatabasespresets) | **Get** /api/v2/presets/dbs | Получение списка тарифов для баз данных
*DatabasesAPI* | [**PerformDatabaseClusterAction**](docs/DatabasesAPI.md#performdatabaseclusteraction) | **Post** /api/v1/databases/{db_cluster_id}/action | Выполнение действия над кластером базы данных
*DatabasesAPI* | [**RestoreDatabaseFromBackup**](docs/DatabasesAPI.md#restoredatabasefrombackup) | **Put** /api/v1/dbs/{db_id}/backups/{backup_id} | Восстановление базы данных из бэкапа
*DatabasesAPI* | [**RestoreDatabaseFromS3Backup**](docs/DatabasesAPI.md#restoredatabasefroms3backup) | **Post** /api/v2/databases/{db_id}/backups/{backup_id}/restore | Восстановление базы данных из S3-бэкапа
*DatabasesAPI* | [**UpdateDatabaseAutoBackupsSettings**](docs/DatabasesAPI.md#updatedatabaseautobackupssettings) | **Patch** /api/v1/dbs/{db_id}/auto-backups | Изменение настроек автобэкапов базы данных
*DatabasesAPI* | [**UpdateDatabaseBackup**](docs/DatabasesAPI.md#updatedatabasebackup) | **Patch** /api/v1/dbs/{db_id}/backups/{backup_id} | Изменение комментария к бэкапу базы данных
*DatabasesAPI* | [**UpdateDatabaseCluster**](docs/DatabasesAPI.md#updatedatabasecluster) | **Patch** /api/v1/databases/{db_cluster_id} | Изменение кластера базы данных
*DatabasesAPI* | [**UpdateDatabaseClusterV2**](docs/DatabasesAPI.md#updatedatabaseclusterv2) | **Patch** /api/v2/databases/{db_cluster_id} | Изменение кластера базы данных (v2)
*DatabasesAPI* | [**UpdateDatabaseInstance**](docs/DatabasesAPI.md#updatedatabaseinstance) | **Patch** /api/v1/databases/{db_cluster_id}/instances/{instance_id} | Изменение инстанса базы данных
*DatabasesAPI* | [**UpdateDatabaseS3Backup**](docs/DatabasesAPI.md#updatedatabases3backup) | **Patch** /api/v2/databases/{db_id}/backups/{backup_id} | Изменение комментария S3-бэкапа базы данных
*DatabasesAPI* | [**UpdateDatabaseUser**](docs/DatabasesAPI.md#updatedatabaseuser) | **Patch** /api/v1/databases/{db_cluster_id}/admins/{admin_id} | Изменение пользователя базы данных
*DedicatedServersAPI* | [**CreateDedicatedServer**](docs/DedicatedServersAPI.md#creatededicatedserver) | **Post** /api/v1/dedicated-servers | Создание выделенного сервера
*DedicatedServersAPI* | [**DeleteDedicatedServer**](docs/DedicatedServersAPI.md#deletededicatedserver) | **Delete** /api/v1/dedicated-servers/{dedicated_id} | Удаление выделенного сервера
*DedicatedServersAPI* | [**GetDedicatedServer**](docs/DedicatedServersAPI.md#getdedicatedserver) | **Get** /api/v1/dedicated-servers/{dedicated_id} | Получение выделенного сервера
*DedicatedServersAPI* | [**GetDedicatedServerPresetAdditionalServices**](docs/DedicatedServersAPI.md#getdedicatedserverpresetadditionalservices) | **Get** /api/v1/presets/dedicated-servers/{preset_id}/additional-services | Получение дополнительных услуг для выделенного сервера
*DedicatedServersAPI* | [**GetDedicatedServers**](docs/DedicatedServersAPI.md#getdedicatedservers) | **Get** /api/v1/dedicated-servers | Получение списка выделенных серверов
*DedicatedServersAPI* | [**GetDedicatedServersPresets**](docs/DedicatedServersAPI.md#getdedicatedserverspresets) | **Get** /api/v1/presets/dedicated-servers | Получение списка тарифов для выделенного сервера
*DedicatedServersAPI* | [**UpdateDedicatedServer**](docs/DedicatedServersAPI.md#updatededicatedserver) | **Patch** /api/v1/dedicated-servers/{dedicated_id} | Обновление выделенного сервера
*DomainsAPI* | [**AddDomain**](docs/DomainsAPI.md#adddomain) | **Post** /api/v1/add-domain/{fqdn} | Добавление домена на аккаунт
*DomainsAPI* | [**AddSubdomain**](docs/DomainsAPI.md#addsubdomain) | **Post** /api/v1/domains/{fqdn}/subdomains/{subdomain} | Добавление поддомена
*DomainsAPI* | [**CheckDomain**](docs/DomainsAPI.md#checkdomain) | **Get** /api/v1/check-domain/{fqdn} | Проверить, доступен ли домен для регистрации
*DomainsAPI* | [**CreateDomainDNSRecord**](docs/DomainsAPI.md#createdomaindnsrecord) | **Post** /api/v1/domains/{fqdn}/dns-records | Добавить информацию о DNS-записи для домена или поддомена
*DomainsAPI* | [**CreateDomainDNSRecordV2**](docs/DomainsAPI.md#createdomaindnsrecordv2) | **Post** /api/v2/domains/{fqdn}/dns-records | Добавить информацию о DNS-записи для домена или поддомена
*DomainsAPI* | [**CreateDomainRequest**](docs/DomainsAPI.md#createdomainrequest) | **Post** /api/v1/domains-requests | Создание заявки на регистрацию/продление/трансфер домена
*DomainsAPI* | [**CreatePerson**](docs/DomainsAPI.md#createperson) | **Post** /api/v1/persons | Создание администратора доменов
*DomainsAPI* | [**DeleteDomain**](docs/DomainsAPI.md#deletedomain) | **Delete** /api/v1/domains/{fqdn} | Удаление домена
*DomainsAPI* | [**DeleteDomainDNSRecord**](docs/DomainsAPI.md#deletedomaindnsrecord) | **Delete** /api/v1/domains/{fqdn}/dns-records/{record_id} | Удалить информацию о DNS-записи для домена или поддомена
*DomainsAPI* | [**DeleteDomainDNSRecordV2**](docs/DomainsAPI.md#deletedomaindnsrecordv2) | **Delete** /api/v2/domains/{fqdn}/dns-records/{record_id} | Удалить информацию о DNS-записи для домена или поддомена
*DomainsAPI* | [**DeletePerson**](docs/DomainsAPI.md#deleteperson) | **Delete** /api/v1/persons/{person_id} | Удаление администратора доменов
*DomainsAPI* | [**DeleteSubdomain**](docs/DomainsAPI.md#deletesubdomain) | **Delete** /api/v1/domains/{fqdn}/subdomains/{subdomain} | Удаление поддомена
*DomainsAPI* | [**GetDomain**](docs/DomainsAPI.md#getdomain) | **Get** /api/v1/domains/{fqdn} | Получение информации о домене
*DomainsAPI* | [**GetDomainDNSRecords**](docs/DomainsAPI.md#getdomaindnsrecords) | **Get** /api/v1/domains/{fqdn}/dns-records | Получить информацию обо всех пользовательских DNS-записях домена или поддомена
*DomainsAPI* | [**GetDomainDefaultDNSRecords**](docs/DomainsAPI.md#getdomaindefaultdnsrecords) | **Get** /api/v1/domains/{fqdn}/default-dns-records | Получить информацию обо всех DNS-записях по умолчанию домена или поддомена
*DomainsAPI* | [**GetDomainNameServers**](docs/DomainsAPI.md#getdomainnameservers) | **Get** /api/v1/domains/{fqdn}/name-servers | Получение списка name-серверов домена
*DomainsAPI* | [**GetDomainRequest**](docs/DomainsAPI.md#getdomainrequest) | **Get** /api/v1/domains-requests/{request_id} | Получение заявки на регистрацию/продление/трансфер домена
*DomainsAPI* | [**GetDomainRequests**](docs/DomainsAPI.md#getdomainrequests) | **Get** /api/v1/domains-requests | Получение списка заявок на регистрацию/продление/трансфер домена
*DomainsAPI* | [**GetDomains**](docs/DomainsAPI.md#getdomains) | **Get** /api/v1/domains | Получение списка всех доменов
*DomainsAPI* | [**GetPerson**](docs/DomainsAPI.md#getperson) | **Get** /api/v1/persons/{person_id} | Получение администратора доменов
*DomainsAPI* | [**GetPersons**](docs/DomainsAPI.md#getpersons) | **Get** /api/v1/persons | Получение списка администраторов доменов
*DomainsAPI* | [**GetTLD**](docs/DomainsAPI.md#gettld) | **Get** /api/v1/tlds/{tld_id} | Получить информацию о доменной зоне по ID
*DomainsAPI* | [**GetTLDs**](docs/DomainsAPI.md#gettlds) | **Get** /api/v1/tlds | Получить информацию о доменных зонах
*DomainsAPI* | [**UpdateDomainAutoProlongation**](docs/DomainsAPI.md#updatedomainautoprolongation) | **Patch** /api/v1/domains/{fqdn} | Включение/выключение автопродления домена
*DomainsAPI* | [**UpdateDomainDNSRecord**](docs/DomainsAPI.md#updatedomaindnsrecord) | **Patch** /api/v1/domains/{fqdn}/dns-records/{record_id} | Обновить информацию о DNS-записи домена или поддомена
*DomainsAPI* | [**UpdateDomainDNSRecordV2**](docs/DomainsAPI.md#updatedomaindnsrecordv2) | **Patch** /api/v2/domains/{fqdn}/dns-records/{record_id} | Обновить информацию о DNS-записи домена или поддомена
*DomainsAPI* | [**UpdateDomainNameServers**](docs/DomainsAPI.md#updatedomainnameservers) | **Put** /api/v1/domains/{fqdn}/name-servers | Изменение name-серверов домена
*DomainsAPI* | [**UpdateDomainRequest**](docs/DomainsAPI.md#updatedomainrequest) | **Patch** /api/v1/domains-requests/{request_id} | Оплата/обновление заявки на регистрацию/продление/трансфер домена
*DomainsAPI* | [**UpdatePerson**](docs/DomainsAPI.md#updateperson) | **Put** /api/v1/persons/{person_id} | Обновление контактных данных администратора доменов
*FirewallAPI* | [**AddResourceToGroup**](docs/FirewallAPI.md#addresourcetogroup) | **Post** /api/v1/firewall/groups/{group_id}/resources/{resource_id} | Линковка ресурса в firewall group
*FirewallAPI* | [**CreateGroup**](docs/FirewallAPI.md#creategroup) | **Post** /api/v1/firewall/groups | Создание группы правил
*FirewallAPI* | [**CreateGroupRule**](docs/FirewallAPI.md#creategrouprule) | **Post** /api/v1/firewall/groups/{group_id}/rules | Создание firewall правила
*FirewallAPI* | [**DeleteGroup**](docs/FirewallAPI.md#deletegroup) | **Delete** /api/v1/firewall/groups/{group_id} | Удаление группы правил
*FirewallAPI* | [**DeleteGroupRule**](docs/FirewallAPI.md#deletegrouprule) | **Delete** /api/v1/firewall/groups/{group_id}/rules/{rule_id} | Удаление firewall правила
*FirewallAPI* | [**DeleteResourceFromGroup**](docs/FirewallAPI.md#deleteresourcefromgroup) | **Delete** /api/v1/firewall/groups/{group_id}/resources/{resource_id} | Отлинковка ресурса из firewall group
*FirewallAPI* | [**GetGroup**](docs/FirewallAPI.md#getgroup) | **Get** /api/v1/firewall/groups/{group_id} | Получение информации о группе правил
*FirewallAPI* | [**GetGroupResources**](docs/FirewallAPI.md#getgroupresources) | **Get** /api/v1/firewall/groups/{group_id}/resources | Получение слинкованных ресурсов
*FirewallAPI* | [**GetGroupRule**](docs/FirewallAPI.md#getgrouprule) | **Get** /api/v1/firewall/groups/{group_id}/rules/{rule_id} | Получение информации о правиле
*FirewallAPI* | [**GetGroupRules**](docs/FirewallAPI.md#getgrouprules) | **Get** /api/v1/firewall/groups/{group_id}/rules | Получение списка правил
*FirewallAPI* | [**GetGroups**](docs/FirewallAPI.md#getgroups) | **Get** /api/v1/firewall/groups | Получение групп правил
*FirewallAPI* | [**GetRulesForResource**](docs/FirewallAPI.md#getrulesforresource) | **Get** /api/v1/firewall/service/{resource_type}/{resource_id} | Получение групп правил для ресурса
*FirewallAPI* | [**UpdateGroup**](docs/FirewallAPI.md#updategroup) | **Patch** /api/v1/firewall/groups/{group_id} | Обновление группы правил
*FirewallAPI* | [**UpdateGroupRule**](docs/FirewallAPI.md#updategrouprule) | **Patch** /api/v1/firewall/groups/{group_id}/rules/{rule_id} | Обновление firewall правила
*FloatingIPAPI* | [**BindFloatingIp**](docs/FloatingIPAPI.md#bindfloatingip) | **Post** /api/v1/floating-ips/{floating_ip_id}/bind | Привязать IP к сервису
*FloatingIPAPI* | [**CreateFloatingIp**](docs/FloatingIPAPI.md#createfloatingip) | **Post** /api/v1/floating-ips | Создание плавающего IP
*FloatingIPAPI* | [**DeleteFloatingIP**](docs/FloatingIPAPI.md#deletefloatingip) | **Delete** /api/v1/floating-ips/{floating_ip_id} | Удаление плавающего IP по ID
*FloatingIPAPI* | [**GetFloatingIp**](docs/FloatingIPAPI.md#getfloatingip) | **Get** /api/v1/floating-ips/{floating_ip_id} | Получение плавающего IP
*FloatingIPAPI* | [**GetFloatingIps**](docs/FloatingIPAPI.md#getfloatingips) | **Get** /api/v1/floating-ips | Получение списка плавающих IP
*FloatingIPAPI* | [**UnbindFloatingIp**](docs/FloatingIPAPI.md#unbindfloatingip) | **Post** /api/v1/floating-ips/{floating_ip_id}/unbind | Отвязать IP от сервиса
*FloatingIPAPI* | [**UpdateFloatingIP**](docs/FloatingIPAPI.md#updatefloatingip) | **Patch** /api/v1/floating-ips/{floating_ip_id} | Изменение плавающего IP по ID
*ImagesAPI* | [**CreateImage**](docs/ImagesAPI.md#createimage) | **Post** /api/v1/images | Создание образа
*ImagesAPI* | [**CreateImageDownloadUrl**](docs/ImagesAPI.md#createimagedownloadurl) | **Post** /api/v1/images/{image_id}/download-url | Создание ссылки на скачивание образа
*ImagesAPI* | [**DeleteImage**](docs/ImagesAPI.md#deleteimage) | **Delete** /api/v1/images/{image_id} | Удаление образа
*ImagesAPI* | [**DeleteImageDownloadURL**](docs/ImagesAPI.md#deleteimagedownloadurl) | **Delete** /api/v1/images/{image_id}/download-url/{image_url_id} | Удаление ссылки на образ
*ImagesAPI* | [**GetImage**](docs/ImagesAPI.md#getimage) | **Get** /api/v1/images/{image_id} | Получение информации о образе
*ImagesAPI* | [**GetImageDownloadURL**](docs/ImagesAPI.md#getimagedownloadurl) | **Get** /api/v1/images/{image_id}/download-url/{image_url_id} | Получение информации о ссылке на скачивание образа
*ImagesAPI* | [**GetImageDownloadURLs**](docs/ImagesAPI.md#getimagedownloadurls) | **Get** /api/v1/images/{image_id}/download-url | Получение информации о ссылках на скачивание образов
*ImagesAPI* | [**GetImages**](docs/ImagesAPI.md#getimages) | **Get** /api/v1/images | Получение списка образов
*ImagesAPI* | [**UpdateImage**](docs/ImagesAPI.md#updateimage) | **Patch** /api/v1/images/{image_id} | Обновление информации о образе
*ImagesAPI* | [**UploadImage**](docs/ImagesAPI.md#uploadimage) | **Post** /api/v1/images/{image_id} | Загрузка образа
*KnowledgeBasesAPI* | [**AddAdditionalTokenPackageToKnowledgebase**](docs/KnowledgeBasesAPI.md#addadditionaltokenpackagetoknowledgebase) | **Post** /api/v1/cloud-ai/knowledge-bases/{id}/add-additional-token-package | Покупка дополнительных пакетов токенов
*KnowledgeBasesAPI* | [**CreateKnowledgebase**](docs/KnowledgeBasesAPI.md#createknowledgebase) | **Post** /api/v1/cloud-ai/knowledge-bases | Создание базы знаний
*KnowledgeBasesAPI* | [**CreateKnowledgebaseDocument**](docs/KnowledgeBasesAPI.md#createknowledgebasedocument) | **Post** /api/v1/cloud-ai/knowledge-bases/{id}/documents | Добавление документа по ссылке
*KnowledgeBasesAPI* | [**DeleteDocument**](docs/KnowledgeBasesAPI.md#deletedocument) | **Delete** /api/v1/cloud-ai/knowledge-bases/{id}/documents/{document_id} | Удаление документа из базы знаний
*KnowledgeBasesAPI* | [**DeleteKnowledgebase**](docs/KnowledgeBasesAPI.md#deleteknowledgebase) | **Delete** /api/v1/cloud-ai/knowledge-bases/{id} | Удаление базы знаний
*KnowledgeBasesAPI* | [**DownloadDocument**](docs/KnowledgeBasesAPI.md#downloaddocument) | **Get** /api/v1/cloud-ai/knowledge-bases/{id}/documents/{document_id}/download | Скачивание файла документа
*KnowledgeBasesAPI* | [**GetKnowledgebase**](docs/KnowledgeBasesAPI.md#getknowledgebase) | **Get** /api/v1/cloud-ai/knowledge-bases/{id} | Получение базы знаний со списком документов
*KnowledgeBasesAPI* | [**GetKnowledgebaseDocumentsV2**](docs/KnowledgeBasesAPI.md#getknowledgebasedocumentsv2) | **Get** /api/v2/cloud-ai/knowledge-bases/{id}/documents | Получение списка документов базы знаний
*KnowledgeBasesAPI* | [**GetKnowledgebaseV3**](docs/KnowledgeBasesAPI.md#getknowledgebasev3) | **Get** /api/v3/cloud-ai/knowledge-bases/{id} | Получение базы знаний
*KnowledgeBasesAPI* | [**GetKnowledgebases**](docs/KnowledgeBasesAPI.md#getknowledgebases) | **Get** /api/v1/cloud-ai/knowledge-bases | Получение списка баз знаний (v1)
*KnowledgeBasesAPI* | [**GetKnowledgebasesV2**](docs/KnowledgeBasesAPI.md#getknowledgebasesv2) | **Get** /api/v2/cloud-ai/knowledge-bases | Получение списка баз знаний (v2)
*KnowledgeBasesAPI* | [**GetKnowledgebasesV3**](docs/KnowledgeBasesAPI.md#getknowledgebasesv3) | **Get** /api/v3/cloud-ai/knowledge-bases | Получение списка баз знаний
*KnowledgeBasesAPI* | [**LinkKnowledgebaseToAgent**](docs/KnowledgeBasesAPI.md#linkknowledgebasetoagent) | **Post** /api/v1/cloud-ai/knowledge-bases/{id}/link/{agent_id} | Привязка базы знаний к агенту
*KnowledgeBasesAPI* | [**ReindexDocument**](docs/KnowledgeBasesAPI.md#reindexdocument) | **Post** /api/v1/cloud-ai/knowledge-bases/{id}/documents/{document_id}/reindex | Переиндексация документа
*KnowledgeBasesAPI* | [**UnlinkKnowledgebaseFromAgent**](docs/KnowledgeBasesAPI.md#unlinkknowledgebasefromagent) | **Delete** /api/v1/cloud-ai/knowledge-bases/{id}/link/{agent_id} | Отвязка базы знаний от агента
*KnowledgeBasesAPI* | [**UpdateKnowledgebase**](docs/KnowledgeBasesAPI.md#updateknowledgebase) | **Patch** /api/v1/cloud-ai/knowledge-bases/{id} | Изменение базы знаний
*KnowledgeBasesAPI* | [**UpdateKnowledgebaseAgentLink**](docs/KnowledgeBasesAPI.md#updateknowledgebaseagentlink) | **Patch** /api/v1/cloud-ai/knowledge-bases/{id}/link/{agent_id} | Изменение режима поиска по базе знаний для агента
*KnowledgeBasesAPI* | [**UpdateKnowledgebaseDocument**](docs/KnowledgeBasesAPI.md#updateknowledgebasedocument) | **Patch** /api/v1/cloud-ai/knowledge-bases/{id}/documents/{document_id} | Изменение документа
*KnowledgeBasesAPI* | [**UploadFilesToKnowledgebase**](docs/KnowledgeBasesAPI.md#uploadfilestoknowledgebase) | **Post** /api/v1/cloud-ai/knowledge-bases/{id}/upload | Загрузка файлов в базу знаний
*KubernetesAPI* | [**CreateCluster**](docs/KubernetesAPI.md#createcluster) | **Post** /api/v1/k8s/clusters | Создание кластера
*KubernetesAPI* | [**CreateClusterNodeGroup**](docs/KubernetesAPI.md#createclusternodegroup) | **Post** /api/v1/k8s/clusters/{cluster_id}/groups | Создание группы нод
*KubernetesAPI* | [**DeleteCluster**](docs/KubernetesAPI.md#deletecluster) | **Delete** /api/v1/k8s/clusters/{cluster_id} | Удаление кластера
*KubernetesAPI* | [**DeleteClusterNode**](docs/KubernetesAPI.md#deleteclusternode) | **Delete** /api/v1/k8s/clusters/{cluster_id}/nodes/{node_id} | Удаление ноды
*KubernetesAPI* | [**DeleteClusterNodeGroup**](docs/KubernetesAPI.md#deleteclusternodegroup) | **Delete** /api/v1/k8s/clusters/{cluster_id}/groups/{group_id} | Удаление группы нод
*KubernetesAPI* | [**DeleteKubernetesAddons**](docs/KubernetesAPI.md#deletekubernetesaddons) | **Delete** /api/v1/k8s/clusters/{cluster_id}/addons/{addon_id} | Удаление дополнения
*KubernetesAPI* | [**GetCluster**](docs/KubernetesAPI.md#getcluster) | **Get** /api/v1/k8s/clusters/{cluster_id} | Получение информации о кластере
*KubernetesAPI* | [**GetClusterKubeconfig**](docs/KubernetesAPI.md#getclusterkubeconfig) | **Get** /api/v1/k8s/clusters/{cluster_id}/kubeconfig | Получение файла kubeconfig
*KubernetesAPI* | [**GetClusterNodeGroup**](docs/KubernetesAPI.md#getclusternodegroup) | **Get** /api/v1/k8s/clusters/{cluster_id}/groups/{group_id} | Получение информации о группе нод
*KubernetesAPI* | [**GetClusterNodeGroups**](docs/KubernetesAPI.md#getclusternodegroups) | **Get** /api/v1/k8s/clusters/{cluster_id}/groups | Получение групп нод кластера
*KubernetesAPI* | [**GetClusterNodes**](docs/KubernetesAPI.md#getclusternodes) | **Get** /api/v1/k8s/clusters/{cluster_id}/nodes | Получение списка нод
*KubernetesAPI* | [**GetClusterNodesFromGroup**](docs/KubernetesAPI.md#getclusternodesfromgroup) | **Get** /api/v1/k8s/clusters/{cluster_id}/groups/{group_id}/nodes | Получение списка нод, принадлежащих группе
*KubernetesAPI* | [**GetClusterResources**](docs/KubernetesAPI.md#getclusterresources) | **Get** /api/v1/k8s/clusters/{cluster_id}/resources | Получение ресурсов кластера
*KubernetesAPI* | [**GetClusters**](docs/KubernetesAPI.md#getclusters) | **Get** /api/v1/k8s/clusters | Получение списка кластеров
*KubernetesAPI* | [**GetK8SNetworkDrivers**](docs/KubernetesAPI.md#getk8snetworkdrivers) | **Get** /api/v1/k8s/network-drivers | Получение списка сетевых драйверов k8s
*KubernetesAPI* | [**GetK8SVersions**](docs/KubernetesAPI.md#getk8sversions) | **Get** /api/v1/k8s/k8s-versions | Получение списка версий k8s
*KubernetesAPI* | [**GetKubernetesAddons**](docs/KubernetesAPI.md#getkubernetesaddons) | **Get** /api/v1/k8s/clusters/{cluster_id}/addons | Получение списка установленных дополнений
*KubernetesAPI* | [**GetKubernetesAddonsConfig**](docs/KubernetesAPI.md#getkubernetesaddonsconfig) | **Get** /api/v1/k8s/clusters/{cluster_id}/addons-configs | Получение списка конфигураций дополнений
*KubernetesAPI* | [**GetKubernetesPresets**](docs/KubernetesAPI.md#getkubernetespresets) | **Get** /api/v1/presets/k8s | Получение списка тарифов
*KubernetesAPI* | [**IncreaseCountOfNodesInGroup**](docs/KubernetesAPI.md#increasecountofnodesingroup) | **Post** /api/v1/k8s/clusters/{cluster_id}/groups/{group_id}/nodes | Увеличение количества нод в группе на указанное количество
*KubernetesAPI* | [**PostKubernetesAddons**](docs/KubernetesAPI.md#postkubernetesaddons) | **Post** /api/v1/k8s/clusters/{cluster_id}/addons | Установка дополнения
*KubernetesAPI* | [**PostKubernetesAddonsUpdate**](docs/KubernetesAPI.md#postkubernetesaddonsupdate) | **Post** /api/v1/k8s/clusters/{cluster_id}/addons/{addon_id} | Изменение конфигурации дополнения
*KubernetesAPI* | [**ReduceCountOfNodesInGroup**](docs/KubernetesAPI.md#reducecountofnodesingroup) | **Delete** /api/v1/k8s/clusters/{cluster_id}/groups/{group_id}/nodes | Уменьшение количества нод в группе на указанное количество
*KubernetesAPI* | [**UpdateCluster**](docs/KubernetesAPI.md#updatecluster) | **Patch** /api/v1/k8s/clusters/{cluster_id} | Обновление информации о кластере
*KubernetesAPI* | [**UpdateClusterVersion**](docs/KubernetesAPI.md#updateclusterversion) | **Patch** /api/v1/k8s/clusters/{cluster_id}/versions/update | Обновление версии кластера
*LocationsAPI* | [**GetLocations**](docs/LocationsAPI.md#getlocations) | **Get** /api/v2/locations | Получение списка локаций
*MCPAPI* | [**DeleteMcpSearchSession**](docs/MCPAPI.md#deletemcpsearchsession) | **Delete** /api/v1/mcp/search | Завершение сессии сервера MCP с поиском инструментов
*MCPAPI* | [**DeleteMcpSession**](docs/MCPAPI.md#deletemcpsession) | **Delete** /api/v1/mcp | Завершение сессии сервера MCP
*MCPAPI* | [**OpenMcpSearchStream**](docs/MCPAPI.md#openmcpsearchstream) | **Get** /api/v1/mcp/search | Открытие потока событий сервера MCP с поиском инструментов
*MCPAPI* | [**OpenMcpStream**](docs/MCPAPI.md#openmcpstream) | **Get** /api/v1/mcp | Открытие потока событий сервера MCP
*MCPAPI* | [**SendMcpMessage**](docs/MCPAPI.md#sendmcpmessage) | **Post** /api/v1/mcp | Отправка JSON-RPC-сообщения серверу MCP
*MCPAPI* | [**SendMcpSearchMessage**](docs/MCPAPI.md#sendmcpsearchmessage) | **Post** /api/v1/mcp/search | Отправка JSON-RPC-сообщения серверу MCP с поиском инструментов
*MailAPI* | [**CreateDomainMailbox**](docs/MailAPI.md#createdomainmailbox) | **Post** /api/v1/mail/domains/{domain} | Создание почтового ящика
*MailAPI* | [**CreateDomainMailboxV2**](docs/MailAPI.md#createdomainmailboxv2) | **Post** /api/v2/mail/domains/{domain} | Создание почтового ящика
*MailAPI* | [**CreateMultipleDomainMailboxes**](docs/MailAPI.md#createmultipledomainmailboxes) | **Post** /api/v1/mail/domains/{domain}/batch | Множественное создание почтовых ящиков
*MailAPI* | [**CreateMultipleDomainMailboxesV2**](docs/MailAPI.md#createmultipledomainmailboxesv2) | **Post** /api/v2/mail/domains/{domain}/batch | Множественное создание почтовых ящиков
*MailAPI* | [**DeleteMailbox**](docs/MailAPI.md#deletemailbox) | **Delete** /api/v1/mail/domains/{domain}/mailboxes/{mailbox} | Удаление почтового ящика
*MailAPI* | [**GetAllMailboxesV2**](docs/MailAPI.md#getallmailboxesv2) | **Get** /api/v2/mail | Получение списка всех почтовых ящиков аккаунта
*MailAPI* | [**GetDomainMailInfo**](docs/MailAPI.md#getdomainmailinfo) | **Get** /api/v1/mail/domains/{domain}/info | Получение почтовой информации о домене
*MailAPI* | [**GetDomainMailboxes**](docs/MailAPI.md#getdomainmailboxes) | **Get** /api/v1/mail/domains/{domain} | Получение списка почтовых ящиков домена
*MailAPI* | [**GetMailbox**](docs/MailAPI.md#getmailbox) | **Get** /api/v1/mail/domains/{domain}/mailboxes/{mailbox} | Получение почтового ящика
*MailAPI* | [**GetMailboxV2**](docs/MailAPI.md#getmailboxv2) | **Get** /api/v2/mail/domains/{domain}/mailboxes/{mailbox} | Получение почтового ящика
*MailAPI* | [**GetMailboxes**](docs/MailAPI.md#getmailboxes) | **Get** /api/v1/mail | Получение списка почтовых ящиков аккаунта
*MailAPI* | [**UpdateDomainMailInfo**](docs/MailAPI.md#updatedomainmailinfo) | **Patch** /api/v1/mail/domains/{domain}/info | Изменение почтовой информации о домене
*MailAPI* | [**UpdateMailbox**](docs/MailAPI.md#updatemailbox) | **Patch** /api/v1/mail/domains/{domain}/mailboxes/{mailbox} | Изменение почтового ящика
*MailAPI* | [**UpdateMailboxV2**](docs/MailAPI.md#updatemailboxv2) | **Patch** /api/v2/mail/domains/{domain}/mailboxes/{mailbox} | Изменение почтового ящика
*NetworkDrivesAPI* | [**CreateNetworkDrive**](docs/NetworkDrivesAPI.md#createnetworkdrive) | **Post** /api/v1/network-drives | Создание сетевого диска
*NetworkDrivesAPI* | [**DeleteNetworkDrive**](docs/NetworkDrivesAPI.md#deletenetworkdrive) | **Delete** /api/v1/network-drives/{network_drive_id} | Удаление сетевого диска по идентификатору
*NetworkDrivesAPI* | [**GetNetworkDrive**](docs/NetworkDrivesAPI.md#getnetworkdrive) | **Get** /api/v1/network-drives/{network_drive_id} | Получение сетевого диска
*NetworkDrivesAPI* | [**GetNetworkDrives**](docs/NetworkDrivesAPI.md#getnetworkdrives) | **Get** /api/v1/network-drives | Получение списка cетевых дисков
*NetworkDrivesAPI* | [**GetNetworkDrivesAvailableResources**](docs/NetworkDrivesAPI.md#getnetworkdrivesavailableresources) | **Get** /api/v1/network-drives/available-resources | Получение списка сервисов доступных для подключения диска
*NetworkDrivesAPI* | [**GetNetworkDrivesPresets**](docs/NetworkDrivesAPI.md#getnetworkdrivespresets) | **Get** /api/v1/presets/network-drives | Получение списка доступных тарифов для сетевого диска
*NetworkDrivesAPI* | [**MountNetworkDrive**](docs/NetworkDrivesAPI.md#mountnetworkdrive) | **Post** /api/v1/network-drives/{network_drive_id}/mount | Подключить сетевой диск к сервису
*NetworkDrivesAPI* | [**UnmountNetworkDrive**](docs/NetworkDrivesAPI.md#unmountnetworkdrive) | **Post** /api/v1/network-drives/{network_drive_id}/unmount | Отключить сетевой диск от сервиса
*NetworkDrivesAPI* | [**UpdateNetworkDrive**](docs/NetworkDrivesAPI.md#updatenetworkdrive) | **Patch** /api/v1/network-drives/{network_drive_id} | Изменение сетевого диска по ID
*PaymentsAPI* | [**GetFinances**](docs/PaymentsAPI.md#getfinances) | **Get** /api/v1/account/finances | Получение платежной информации
*PaymentsAPI* | [**GetServicePrices**](docs/PaymentsAPI.md#getserviceprices) | **Get** /api/v1/account/services/cost | Получение стоимости сервисов
*ProjectsAPI* | [**AddBalancerToProject**](docs/ProjectsAPI.md#addbalancertoproject) | **Post** /api/v1/projects/{project_id}/resources/balancers | Добавление балансировщика в проект
*ProjectsAPI* | [**AddClusterToProject**](docs/ProjectsAPI.md#addclustertoproject) | **Post** /api/v1/projects/{project_id}/resources/clusters | Добавление кластера в проект
*ProjectsAPI* | [**AddDatabaseToProject**](docs/ProjectsAPI.md#adddatabasetoproject) | **Post** /api/v1/projects/{project_id}/resources/databases | Добавление базы данных в проект
*ProjectsAPI* | [**AddDedicatedServerToProject**](docs/ProjectsAPI.md#adddedicatedservertoproject) | **Post** /api/v1/projects/{project_id}/resources/dedicated | Добавление выделенного сервера в проект
*ProjectsAPI* | [**AddServerToProject**](docs/ProjectsAPI.md#addservertoproject) | **Post** /api/v1/projects/{project_id}/resources/servers | Добавление сервера в проект
*ProjectsAPI* | [**AddStorageToProject**](docs/ProjectsAPI.md#addstoragetoproject) | **Post** /api/v1/projects/{project_id}/resources/buckets | Добавление хранилища в проект
*ProjectsAPI* | [**CreateProject**](docs/ProjectsAPI.md#createproject) | **Post** /api/v1/projects | Создание проекта
*ProjectsAPI* | [**DeleteProject**](docs/ProjectsAPI.md#deleteproject) | **Delete** /api/v1/projects/{project_id} | Удаление проекта
*ProjectsAPI* | [**GetAccountBalancers**](docs/ProjectsAPI.md#getaccountbalancers) | **Get** /api/v1/projects/resources/balancers | Получение списка всех балансировщиков на аккаунте
*ProjectsAPI* | [**GetAccountClusters**](docs/ProjectsAPI.md#getaccountclusters) | **Get** /api/v1/projects/resources/clusters | Получение списка всех кластеров на аккаунте
*ProjectsAPI* | [**GetAccountDatabases**](docs/ProjectsAPI.md#getaccountdatabases) | **Get** /api/v1/projects/resources/databases | Получение списка всех баз данных на аккаунте
*ProjectsAPI* | [**GetAccountDedicatedServers**](docs/ProjectsAPI.md#getaccountdedicatedservers) | **Get** /api/v1/projects/resources/dedicated | Получение списка всех выделенных серверов на аккаунте
*ProjectsAPI* | [**GetAccountServers**](docs/ProjectsAPI.md#getaccountservers) | **Get** /api/v1/projects/resources/servers | Получение списка всех серверов на аккаунте
*ProjectsAPI* | [**GetAccountStorages**](docs/ProjectsAPI.md#getaccountstorages) | **Get** /api/v1/projects/resources/buckets | Получение списка всех хранилищ на аккаунте
*ProjectsAPI* | [**GetAllProjectResources**](docs/ProjectsAPI.md#getallprojectresources) | **Get** /api/v1/projects/{project_id}/resources | Получение всех ресурсов проекта
*ProjectsAPI* | [**GetProject**](docs/ProjectsAPI.md#getproject) | **Get** /api/v1/projects/{project_id} | Получение проекта по ID
*ProjectsAPI* | [**GetProjectBalancers**](docs/ProjectsAPI.md#getprojectbalancers) | **Get** /api/v1/projects/{project_id}/resources/balancers | Получение списка балансировщиков проекта
*ProjectsAPI* | [**GetProjectClusters**](docs/ProjectsAPI.md#getprojectclusters) | **Get** /api/v1/projects/{project_id}/resources/clusters | Получение списка кластеров проекта
*ProjectsAPI* | [**GetProjectDatabases**](docs/ProjectsAPI.md#getprojectdatabases) | **Get** /api/v1/projects/{project_id}/resources/databases | Получение списка баз данных проекта
*ProjectsAPI* | [**GetProjectDedicatedServers**](docs/ProjectsAPI.md#getprojectdedicatedservers) | **Get** /api/v1/projects/{project_id}/resources/dedicated | Получение списка выделенных серверов проекта
*ProjectsAPI* | [**GetProjectServers**](docs/ProjectsAPI.md#getprojectservers) | **Get** /api/v1/projects/{project_id}/resources/servers | Получение списка серверов проекта
*ProjectsAPI* | [**GetProjectStorages**](docs/ProjectsAPI.md#getprojectstorages) | **Get** /api/v1/projects/{project_id}/resources/buckets | Получение списка хранилищ проекта
*ProjectsAPI* | [**GetProjects**](docs/ProjectsAPI.md#getprojects) | **Get** /api/v1/projects | Получение списка проектов
*ProjectsAPI* | [**TransferResourceToAnotherProject**](docs/ProjectsAPI.md#transferresourcetoanotherproject) | **Put** /api/v1/projects/{project_id}/resources/transfer | Перенести ресурс в другой проект
*ProjectsAPI* | [**UpdateProject**](docs/ProjectsAPI.md#updateproject) | **Put** /api/v1/projects/{project_id} | Изменение проекта
*RoutersAPI* | [**AddNetworks**](docs/RoutersAPI.md#addnetworks) | **Post** /api/v1/routers/{router_id}/networks | Подключение сетей к роутеру
*RoutersAPI* | [**CreateRouter**](docs/RoutersAPI.md#createrouter) | **Post** /api/v1/routers | Создание роутера
*RoutersAPI* | [**DeleteDnat**](docs/RoutersAPI.md#deletednat) | **Delete** /api/v1/routers/{router_id}/dnat-rules/{dnat_id} | Удаление правила проброса портов
*RoutersAPI* | [**DeleteRouter**](docs/RoutersAPI.md#deleterouter) | **Delete** /api/v1/routers/{router_id} | Удаление роутера
*RoutersAPI* | [**DeleteRouterNat**](docs/RoutersAPI.md#deleterouternat) | **Delete** /api/v1/routers/{router_id}/networks/{network_name}/nat | Выключение NAT для сети
*RoutersAPI* | [**DeleteRouterNetwork**](docs/RoutersAPI.md#deleterouternetwork) | **Delete** /api/v1/routers/{router_id}/networks/{network_name} | Удаление сети из роутера
*RoutersAPI* | [**DeleteStaticRoute**](docs/RoutersAPI.md#deletestaticroute) | **Delete** /api/v1/routers/{router_id}/static-routes/{static_route_id} | Удаление статического маршрута
*RoutersAPI* | [**GetAvailableStaticRoutes**](docs/RoutersAPI.md#getavailablestaticroutes) | **Get** /api/v1/routers/{router_id}/static-routes/available | Получение доступных подсетей для статических маршрутов
*RoutersAPI* | [**GetDnat**](docs/RoutersAPI.md#getdnat) | **Get** /api/v1/routers/{router_id}/dnat-rules | Получение списка правил проброса портов
*RoutersAPI* | [**GetDnatRule**](docs/RoutersAPI.md#getdnatrule) | **Get** /api/v1/routers/{router_id}/dnat-rules/{dnat_id} | Получение правила проброса портов
*RoutersAPI* | [**GetNetworks**](docs/RoutersAPI.md#getnetworks) | **Get** /api/v1/routers/{router_id}/networks | Получение списка сетей роутера
*RoutersAPI* | [**GetRouter**](docs/RoutersAPI.md#getrouter) | **Get** /api/v1/routers/{router_id} | Получение информации о роутере
*RoutersAPI* | [**GetRouterAvailableNetworks**](docs/RoutersAPI.md#getrouteravailablenetworks) | **Get** /api/v1/routers/networks/available | Получение списка доступных сетей
*RoutersAPI* | [**GetRouterPresets**](docs/RoutersAPI.md#getrouterpresets) | **Get** /api/v1/presets/routers | Получение списка тарифов роутеров
*RoutersAPI* | [**GetRouterStatistics**](docs/RoutersAPI.md#getrouterstatistics) | **Get** /api/v1/routers/{router_id}/statistics/{time_from}/{period}/{keys} | Получение статистики роутера
*RoutersAPI* | [**GetRouters**](docs/RoutersAPI.md#getrouters) | **Get** /api/v1/routers | Получение списка роутеров
*RoutersAPI* | [**GetStaticRoutes**](docs/RoutersAPI.md#getstaticroutes) | **Get** /api/v1/routers/{router_id}/static-routes | Получение списка статических маршрутов
*RoutersAPI* | [**PatchNetwork**](docs/RoutersAPI.md#patchnetwork) | **Patch** /api/v1/routers/{router_id}/networks/{network_name} | Обновление информации о сети
*RoutersAPI* | [**PatchNetworks**](docs/RoutersAPI.md#patchnetworks) | **Patch** /api/v1/routers/{router_id}/networks | Обновление сетей роутера
*RoutersAPI* | [**PostDnat**](docs/RoutersAPI.md#postdnat) | **Post** /api/v1/routers/{router_id}/dnat-rules | Добавление правила проброса портов
*RoutersAPI* | [**PostStaticRoute**](docs/RoutersAPI.md#poststaticroute) | **Post** /api/v1/routers/{router_id}/static-routes | Добавление статического маршрута
*RoutersAPI* | [**UpdateRouter**](docs/RoutersAPI.md#updaterouter) | **Patch** /api/v1/routers/{router_id} | Обновление информации о роутере
*RoutersAPI* | [**UpdateRouterNat**](docs/RoutersAPI.md#updaterouternat) | **Patch** /api/v1/routers/{router_id}/networks/{network_name}/nat | Включение NAT для сети
*S3API* | [**AddStorageSubdomainCertificate**](docs/S3API.md#addstoragesubdomaincertificate) | **Post** /api/v1/storages/certificates/generate | Добавление сертификата для поддомена хранилища
*S3API* | [**AddStorageSubdomains**](docs/S3API.md#addstoragesubdomains) | **Post** /api/v1/storages/buckets/{bucket_id}/subdomains | Добавление поддоменов для хранилища
*S3API* | [**CreateStorage**](docs/S3API.md#createstorage) | **Post** /api/v1/storages/buckets | Создание хранилища
*S3API* | [**DeleteStorage**](docs/S3API.md#deletestorage) | **Delete** /api/v1/storages/buckets/{bucket_id} | Удаление хранилища на аккаунте
*S3API* | [**DeleteStorageSubdomains**](docs/S3API.md#deletestoragesubdomains) | **Delete** /api/v1/storages/buckets/{bucket_id}/subdomains | Удаление поддоменов хранилища
*S3API* | [**GetStorage**](docs/S3API.md#getstorage) | **Get** /api/v1/storages/buckets/{bucket_id} | Получение хранилища по ID
*S3API* | [**GetStorageSubdomains**](docs/S3API.md#getstoragesubdomains) | **Get** /api/v1/storages/buckets/{bucket_id}/subdomains | Получение списка поддоменов хранилища
*S3API* | [**GetStorageTransferStatus**](docs/S3API.md#getstoragetransferstatus) | **Get** /api/v1/storages/buckets/{bucket_id}/transfer-status | Получение статуса переноса хранилища от стороннего S3 в Timeweb Cloud
*S3API* | [**GetStorageUsers**](docs/S3API.md#getstorageusers) | **Get** /api/v1/storages/users | Получение списка пользователей хранилищ аккаунта
*S3API* | [**GetStorages**](docs/S3API.md#getstorages) | **Get** /api/v1/storages/buckets | Получение списка хранилищ аккаунта
*S3API* | [**GetStoragesPresets**](docs/S3API.md#getstoragespresets) | **Get** /api/v1/presets/storages | Получение списка тарифов для хранилищ
*S3API* | [**TransferStorage**](docs/S3API.md#transferstorage) | **Post** /api/v1/storages/transfer | Перенос хранилища от стороннего провайдера S3 в Timeweb Cloud
*S3API* | [**UpdateStorage**](docs/S3API.md#updatestorage) | **Patch** /api/v1/storages/buckets/{bucket_id} | Изменение хранилища на аккаунте
*S3API* | [**UpdateStorageUser**](docs/S3API.md#updatestorageuser) | **Patch** /api/v1/storages/users/{user_id} | Изменение пароля пользователя-администратора хранилища
*SSHAPI* | [**AddKeyToServer**](docs/SSHAPI.md#addkeytoserver) | **Post** /api/v1/servers/{server_id}/ssh-keys | Добавление SSH-ключей на сервер
*SSHAPI* | [**CreateKey**](docs/SSHAPI.md#createkey) | **Post** /api/v1/ssh-keys | Создание SSH-ключа
*SSHAPI* | [**DeleteKey**](docs/SSHAPI.md#deletekey) | **Delete** /api/v1/ssh-keys/{ssh_key_id} | Удаление SSH-ключа по ID
*SSHAPI* | [**DeleteKeyFromServer**](docs/SSHAPI.md#deletekeyfromserver) | **Delete** /api/v1/servers/{server_id}/ssh-keys/{ssh_key_id} | Удаление SSH-ключей с сервера
*SSHAPI* | [**GetKey**](docs/SSHAPI.md#getkey) | **Get** /api/v1/ssh-keys/{ssh_key_id} | Получение SSH-ключа по ID
*SSHAPI* | [**GetKeys**](docs/SSHAPI.md#getkeys) | **Get** /api/v1/ssh-keys | Получение списка SSH-ключей
*SSHAPI* | [**UpdateKey**](docs/SSHAPI.md#updatekey) | **Patch** /api/v1/ssh-keys/{ssh_key_id} | Изменение SSH-ключа по ID
*ServersAPI* | [**AddServerIP**](docs/ServersAPI.md#addserverip) | **Post** /api/v1/servers/{server_id}/ips | Добавление IP-адреса сервера
*ServersAPI* | [**CloneServer**](docs/ServersAPI.md#cloneserver) | **Post** /api/v1/servers/{server_id}/clone | Клонирование сервера
*ServersAPI* | [**CommitRestorePoint**](docs/ServersAPI.md#commitrestorepoint) | **Post** /api/v1/restore-points/{vds_id}/commit | Фиксация снапшота
*ServersAPI* | [**CreateRestorePoint**](docs/ServersAPI.md#createrestorepoint) | **Post** /api/v1/restore-points/{vds_id}/create | Создание снапшота
*ServersAPI* | [**CreateServer**](docs/ServersAPI.md#createserver) | **Post** /api/v1/servers | Создание сервера
*ServersAPI* | [**CreateServerDisk**](docs/ServersAPI.md#createserverdisk) | **Post** /api/v1/servers/{server_id}/disks | Создание диска сервера
*ServersAPI* | [**CreateServerDiskBackup**](docs/ServersAPI.md#createserverdiskbackup) | **Post** /api/v1/servers/{server_id}/disks/{disk_id}/backups | Создание бэкапа диска сервера
*ServersAPI* | [**DeleteServer**](docs/ServersAPI.md#deleteserver) | **Delete** /api/v1/servers/{server_id} | Удаление сервера
*ServersAPI* | [**DeleteServerDisk**](docs/ServersAPI.md#deleteserverdisk) | **Delete** /api/v1/servers/{server_id}/disks/{disk_id} | Удаление диска сервера
*ServersAPI* | [**DeleteServerDiskBackup**](docs/ServersAPI.md#deleteserverdiskbackup) | **Delete** /api/v1/servers/{server_id}/disks/{disk_id}/backups/{backup_id} | Удаление бэкапа диска сервера
*ServersAPI* | [**DeleteServerIP**](docs/ServersAPI.md#deleteserverip) | **Delete** /api/v1/servers/{server_id}/ips | Удаление IP-адреса сервера
*ServersAPI* | [**GetConfigurators**](docs/ServersAPI.md#getconfigurators) | **Get** /api/v1/configurator/servers | Получение списка конфигураторов серверов
*ServersAPI* | [**GetOsList**](docs/ServersAPI.md#getoslist) | **Get** /api/v1/os/servers | Получение списка операционных систем
*ServersAPI* | [**GetRestorePoint**](docs/ServersAPI.md#getrestorepoint) | **Get** /api/v1/restore-points/{vds_id} | Получение снапшота сервера
*ServersAPI* | [**GetRestorePoints**](docs/ServersAPI.md#getrestorepoints) | **Get** /api/v1/restore-points | Получение списка снапшотов
*ServersAPI* | [**GetServer**](docs/ServersAPI.md#getserver) | **Get** /api/v1/servers/{server_id} | Получение сервера
*ServersAPI* | [**GetServerDisk**](docs/ServersAPI.md#getserverdisk) | **Get** /api/v1/servers/{server_id}/disks/{disk_id} | Получение диска сервера
*ServersAPI* | [**GetServerDiskAutoBackupSettings**](docs/ServersAPI.md#getserverdiskautobackupsettings) | **Get** /api/v1/servers/{server_id}/disks/{disk_id}/auto-backups | Получить настройки автобэкапов диска сервера
*ServersAPI* | [**GetServerDiskBackup**](docs/ServersAPI.md#getserverdiskbackup) | **Get** /api/v1/servers/{server_id}/disks/{disk_id}/backups/{backup_id} | Получение бэкапа диска сервера
*ServersAPI* | [**GetServerDiskBackups**](docs/ServersAPI.md#getserverdiskbackups) | **Get** /api/v1/servers/{server_id}/disks/{disk_id}/backups | Получение списка бэкапов диска сервера
*ServersAPI* | [**GetServerDisks**](docs/ServersAPI.md#getserverdisks) | **Get** /api/v1/servers/{server_id}/disks | Получение списка дисков сервера
*ServersAPI* | [**GetServerIPs**](docs/ServersAPI.md#getserverips) | **Get** /api/v1/servers/{server_id}/ips | Получение списка IP-адресов сервера
*ServersAPI* | [**GetServerLogs**](docs/ServersAPI.md#getserverlogs) | **Get** /api/v1/servers/{server_id}/logs | Получение списка логов сервера
*ServersAPI* | [**GetServerStatistics**](docs/ServersAPI.md#getserverstatistics) | **Get** /api/v1/servers/{server_id}/statistics | Получение статистики сервера
*ServersAPI* | [**GetServerStatisticsNew**](docs/ServersAPI.md#getserverstatisticsnew) | **Get** /api/v1/servers/{server_id}/statistics/{time_from}/{period}/{keys} | Получение статистики сервера
*ServersAPI* | [**GetServers**](docs/ServersAPI.md#getservers) | **Get** /api/v1/servers | Получение списка серверов
*ServersAPI* | [**GetServersPresets**](docs/ServersAPI.md#getserverspresets) | **Get** /api/v1/presets/servers | Получение списка тарифов серверов
*ServersAPI* | [**GetSoftware**](docs/ServersAPI.md#getsoftware) | **Get** /api/v1/software/servers | Получение списка ПО из маркетплейса
*ServersAPI* | [**HardShutdownServer**](docs/ServersAPI.md#hardshutdownserver) | **Post** /api/v1/servers/{server_id}/hard-shutdown | Принудительное выключение сервера
*ServersAPI* | [**ImageUnmountAndServerReload**](docs/ServersAPI.md#imageunmountandserverreload) | **Post** /api/v1/servers/{server_id}/image-unmount | Отмонтирование ISO образа и перезагрузка сервера
*ServersAPI* | [**PerformActionOnBackup**](docs/ServersAPI.md#performactiononbackup) | **Post** /api/v1/servers/{server_id}/disks/{disk_id}/backups/{backup_id}/action | Выполнение действия над бэкапом диска сервера
*ServersAPI* | [**PerformActionOnServer**](docs/ServersAPI.md#performactiononserver) | **Post** /api/v1/servers/{server_id}/action | Выполнение действия над сервером
*ServersAPI* | [**RebootServer**](docs/ServersAPI.md#rebootserver) | **Post** /api/v1/servers/{server_id}/reboot | Перезагрузка сервера
*ServersAPI* | [**RebootServerHard**](docs/ServersAPI.md#rebootserverhard) | **Post** /api/v1/servers/{server_id}/hard-reboot | Принудительная перезагрузка сервера
*ServersAPI* | [**ResetServerPassword**](docs/ServersAPI.md#resetserverpassword) | **Post** /api/v1/servers/{server_id}/reset-password | Сброс пароля сервера
*ServersAPI* | [**RollbackRestorePoint**](docs/ServersAPI.md#rollbackrestorepoint) | **Post** /api/v1/restore-points/{vds_id}/rollback | Откат к снапшоту
*ServersAPI* | [**ShutdownServer**](docs/ServersAPI.md#shutdownserver) | **Post** /api/v1/servers/{server_id}/shutdown | Выключение сервера
*ServersAPI* | [**StartServer**](docs/ServersAPI.md#startserver) | **Post** /api/v1/servers/{server_id}/start | Запуск сервера
*ServersAPI* | [**UpdateServer**](docs/ServersAPI.md#updateserver) | **Patch** /api/v1/servers/{server_id} | Изменение сервера
*ServersAPI* | [**UpdateServerDisk**](docs/ServersAPI.md#updateserverdisk) | **Patch** /api/v1/servers/{server_id}/disks/{disk_id} | Изменение параметров диска сервера
*ServersAPI* | [**UpdateServerDiskAutoBackupSettings**](docs/ServersAPI.md#updateserverdiskautobackupsettings) | **Patch** /api/v1/servers/{server_id}/disks/{disk_id}/auto-backups | Изменение настроек автобэкапов диска сервера
*ServersAPI* | [**UpdateServerDiskBackup**](docs/ServersAPI.md#updateserverdiskbackup) | **Patch** /api/v1/servers/{server_id}/disks/{disk_id}/backups/{backup_id} | Изменение бэкапа диска сервера
*ServersAPI* | [**UpdateServerIP**](docs/ServersAPI.md#updateserverip) | **Patch** /api/v1/servers/{server_id}/ips | Изменение IP-адреса сервера
*ServersAPI* | [**UpdateServerNAT**](docs/ServersAPI.md#updateservernat) | **Patch** /api/v1/servers/{server_id}/local-networks/nat-mode | Изменение правил маршрутизации трафика сервера (NAT)
*ServersAPI* | [**UpdateServerOSBootMode**](docs/ServersAPI.md#updateserverosbootmode) | **Post** /api/v1/servers/{server_id}/boot-mode | Выбор типа загрузки операционной системы сервера
*VPCAPI* | [**CreateVPC**](docs/VPCAPI.md#createvpc) | **Post** /api/v2/vpcs | Создание VPC
*VPCAPI* | [**DeleteVPC**](docs/VPCAPI.md#deletevpc) | **Delete** /api/v1/vpcs/{vpc_id} | Удаление VPC по ID сети
*VPCAPI* | [**GetVPC**](docs/VPCAPI.md#getvpc) | **Get** /api/v2/vpcs/{vpc_id} | Получение VPC
*VPCAPI* | [**GetVPCPorts**](docs/VPCAPI.md#getvpcports) | **Get** /api/v1/vpcs/{vpc_id}/ports | Получение списка портов для VPC
*VPCAPI* | [**GetVPCServices**](docs/VPCAPI.md#getvpcservices) | **Get** /api/v2/vpcs/{vpc_id}/services | Получение списка сервисов в VPC
*VPCAPI* | [**GetVPCs**](docs/VPCAPI.md#getvpcs) | **Get** /api/v2/vpcs | Получение списка VPCs
*VPCAPI* | [**UpdateVPCs**](docs/VPCAPI.md#updatevpcs) | **Patch** /api/v2/vpcs/{vpc_id} | Изменение VPC по ID сети


## Documentation For Models

 - [AddAiSolutionTelegramBot201Response](docs/AddAiSolutionTelegramBot201Response.md)
 - [AddBalancerToProject200Response](docs/AddBalancerToProject200Response.md)
 - [AddBalancerToProjectRequest](docs/AddBalancerToProjectRequest.md)
 - [AddBitbucket](docs/AddBitbucket.md)
 - [AddCdnCertificate422Response](docs/AddCdnCertificate422Response.md)
 - [AddCertificate](docs/AddCertificate.md)
 - [AddClusterToProjectRequest](docs/AddClusterToProjectRequest.md)
 - [AddCountries](docs/AddCountries.md)
 - [AddCountriesToAllowedList201Response](docs/AddCountriesToAllowedList201Response.md)
 - [AddCountriesToAllowedListRequest](docs/AddCountriesToAllowedListRequest.md)
 - [AddDatabaseToProjectRequest](docs/AddDatabaseToProjectRequest.md)
 - [AddDedicatedServerToProjectRequest](docs/AddDedicatedServerToProjectRequest.md)
 - [AddGit](docs/AddGit.md)
 - [AddGithub](docs/AddGithub.md)
 - [AddGitlab](docs/AddGitlab.md)
 - [AddIPsToAllowedList201Response](docs/AddIPsToAllowedList201Response.md)
 - [AddIPsToAllowedListRequest](docs/AddIPsToAllowedListRequest.md)
 - [AddIPsToBalancerRequest](docs/AddIPsToBalancerRequest.md)
 - [AddIps](docs/AddIps.md)
 - [AddKeyToServerRequest](docs/AddKeyToServerRequest.md)
 - [AddProvider201Response](docs/AddProvider201Response.md)
 - [AddServerIP201Response](docs/AddServerIP201Response.md)
 - [AddServerIPRequest](docs/AddServerIPRequest.md)
 - [AddServerToProjectRequest](docs/AddServerToProjectRequest.md)
 - [AddStorageSubdomainCertificateRequest](docs/AddStorageSubdomainCertificateRequest.md)
 - [AddStorageSubdomains200Response](docs/AddStorageSubdomains200Response.md)
 - [AddStorageSubdomainsRequest](docs/AddStorageSubdomainsRequest.md)
 - [AddStorageToProjectRequest](docs/AddStorageToProjectRequest.md)
 - [AddSubdomain201Response](docs/AddSubdomain201Response.md)
 - [AddTokenPackage](docs/AddTokenPackage.md)
 - [AddedSubdomain](docs/AddedSubdomain.md)
 - [AddonConfigOut](docs/AddonConfigOut.md)
 - [AddonOut](docs/AddonOut.md)
 - [AddonsConfigResponse](docs/AddonsConfigResponse.md)
 - [AddonsResponse](docs/AddonsResponse.md)
 - [Agent](docs/Agent.md)
 - [AgentAdditionalModel](docs/AgentAdditionalModel.md)
 - [AgentEntityLimit](docs/AgentEntityLimit.md)
 - [AgentModelSettings](docs/AgentModelSettings.md)
 - [AgentSettings](docs/AgentSettings.md)
 - [AgentSkill](docs/AgentSkill.md)
 - [AgentTokenStatistic](docs/AgentTokenStatistic.md)
 - [AgentToolsInner](docs/AgentToolsInner.md)
 - [AgentV3](docs/AgentV3.md)
 - [AgentV3KnowledgeBasesInner](docs/AgentV3KnowledgeBasesInner.md)
 - [AgentWidgetSettings](docs/AgentWidgetSettings.md)
 - [AiGatewayBatch](docs/AiGatewayBatch.md)
 - [AiGatewayBatchCreateRequest](docs/AiGatewayBatchCreateRequest.md)
 - [AiGatewayBatchList](docs/AiGatewayBatchList.md)
 - [AiGatewayBatchRequestCounts](docs/AiGatewayBatchRequestCounts.md)
 - [AiGatewayChatCompletion](docs/AiGatewayChatCompletion.md)
 - [AiGatewayChatCompletionChoicesInner](docs/AiGatewayChatCompletionChoicesInner.md)
 - [AiGatewayChatCompletionChunk](docs/AiGatewayChatCompletionChunk.md)
 - [AiGatewayChatCompletionChunkChoicesInner](docs/AiGatewayChatCompletionChunkChoicesInner.md)
 - [AiGatewayChatCompletionRequest](docs/AiGatewayChatCompletionRequest.md)
 - [AiGatewayChatCompletionRequestStreamOptions](docs/AiGatewayChatCompletionRequestStreamOptions.md)
 - [AiGatewayChatMessage](docs/AiGatewayChatMessage.md)
 - [AiGatewayCompletion](docs/AiGatewayCompletion.md)
 - [AiGatewayCompletionChoicesInner](docs/AiGatewayCompletionChoicesInner.md)
 - [AiGatewayCompletionRequest](docs/AiGatewayCompletionRequest.md)
 - [AiGatewayEmbeddingList](docs/AiGatewayEmbeddingList.md)
 - [AiGatewayEmbeddingListDataInner](docs/AiGatewayEmbeddingListDataInner.md)
 - [AiGatewayEmbeddingListUsage](docs/AiGatewayEmbeddingListUsage.md)
 - [AiGatewayEmbeddingRequest](docs/AiGatewayEmbeddingRequest.md)
 - [AiGatewayError](docs/AiGatewayError.md)
 - [AiGatewayErrorError](docs/AiGatewayErrorError.md)
 - [AiGatewayFile](docs/AiGatewayFile.md)
 - [AiGatewayFileDeleted](docs/AiGatewayFileDeleted.md)
 - [AiGatewayImageGenerationRequest](docs/AiGatewayImageGenerationRequest.md)
 - [AiGatewayImagesResponse](docs/AiGatewayImagesResponse.md)
 - [AiGatewayImagesResponseDataInner](docs/AiGatewayImagesResponseDataInner.md)
 - [AiGatewayModel](docs/AiGatewayModel.md)
 - [AiGatewayModelList](docs/AiGatewayModelList.md)
 - [AiGatewayResponse](docs/AiGatewayResponse.md)
 - [AiGatewayResponseCreateRequest](docs/AiGatewayResponseCreateRequest.md)
 - [AiGatewayResponseCreateRequestReasoning](docs/AiGatewayResponseCreateRequestReasoning.md)
 - [AiGatewayResponseDeleted](docs/AiGatewayResponseDeleted.md)
 - [AiGatewayResponseIncompleteDetails](docs/AiGatewayResponseIncompleteDetails.md)
 - [AiGatewayResponseUsage](docs/AiGatewayResponseUsage.md)
 - [AiGatewaySpeechRequest](docs/AiGatewaySpeechRequest.md)
 - [AiGatewayToolCall](docs/AiGatewayToolCall.md)
 - [AiGatewayToolCallFunction](docs/AiGatewayToolCallFunction.md)
 - [AiGatewayTranscription](docs/AiGatewayTranscription.md)
 - [AiGatewayUsage](docs/AiGatewayUsage.md)
 - [AiGatewayUsageCompletionTokensDetails](docs/AiGatewayUsageCompletionTokensDetails.md)
 - [AiGatewayUsagePromptTokensDetails](docs/AiGatewayUsagePromptTokensDetails.md)
 - [AiGatewayValidationError](docs/AiGatewayValidationError.md)
 - [AiKey](docs/AiKey.md)
 - [AiKeyCallsStatisticPoint](docs/AiKeyCallsStatisticPoint.md)
 - [AiKeyLimit](docs/AiKeyLimit.md)
 - [AiKeyTokenStatisticPoint](docs/AiKeyTokenStatisticPoint.md)
 - [AiProvider](docs/AiProvider.md)
 - [AiSolution](docs/AiSolution.md)
 - [AiSolutionAiKey](docs/AiSolutionAiKey.md)
 - [AiSolutionChatMessage](docs/AiSolutionChatMessage.md)
 - [AiSolutionChatReply](docs/AiSolutionChatReply.md)
 - [AiSolutionChatStatus](docs/AiSolutionChatStatus.md)
 - [AiSolutionConversation](docs/AiSolutionConversation.md)
 - [AiSolutionTelegramBot](docs/AiSolutionTelegramBot.md)
 - [AiSolutionType](docs/AiSolutionType.md)
 - [AiSolutionTypeRequirements](docs/AiSolutionTypeRequirements.md)
 - [App](docs/App.md)
 - [AppConfiguration](docs/AppConfiguration.md)
 - [AppDiskStatus](docs/AppDiskStatus.md)
 - [AppDomainsInner](docs/AppDomainsInner.md)
 - [AppProvider](docs/AppProvider.md)
 - [ApproveTelegramPairing](docs/ApproveTelegramPairing.md)
 - [AppsPresets](docs/AppsPresets.md)
 - [AppsPresetsBackendPresetsInner](docs/AppsPresetsBackendPresetsInner.md)
 - [AppsPresetsFrontendPresetsInner](docs/AppsPresetsFrontendPresetsInner.md)
 - [AuthorizeUrl](docs/AuthorizeUrl.md)
 - [AutoBackup](docs/AutoBackup.md)
 - [AutoReplyIsDisabled](docs/AutoReplyIsDisabled.md)
 - [AutoReplyIsEnabled](docs/AutoReplyIsEnabled.md)
 - [AutoreplyIsDisabled](docs/AutoreplyIsDisabled.md)
 - [AutoreplyIsEnabled](docs/AutoreplyIsEnabled.md)
 - [AvailabilityZone](docs/AvailabilityZone.md)
 - [AvailableFrameworks](docs/AvailableFrameworks.md)
 - [AvailableFrameworksBackendFrameworksInner](docs/AvailableFrameworksBackendFrameworksInner.md)
 - [AvailableFrameworksFrontendFrameworksInner](docs/AvailableFrameworksFrontendFrameworksInner.md)
 - [AvailableNetwork](docs/AvailableNetwork.md)
 - [AvailableNetworksResponse](docs/AvailableNetworksResponse.md)
 - [AvailableStaticRoute](docs/AvailableStaticRoute.md)
 - [AvailableStaticRouteSubnetsInner](docs/AvailableStaticRouteSubnetsInner.md)
 - [AvailableStaticRoutesResponse](docs/AvailableStaticRoutesResponse.md)
 - [Backup](docs/Backup.md)
 - [BackupDownloadUrl](docs/BackupDownloadUrl.md)
 - [BackupDownloadUrlRequest](docs/BackupDownloadUrlRequest.md)
 - [Balancer](docs/Balancer.md)
 - [BalancerNetworksInner](docs/BalancerNetworksInner.md)
 - [BalancerNetworksInnerIpsInner](docs/BalancerNetworksInnerIpsInner.md)
 - [BaseError](docs/BaseError.md)
 - [BindFloatingIp](docs/BindFloatingIp.md)
 - [Bonus](docs/Bonus.md)
 - [Bot](docs/Bot.md)
 - [Branch](docs/Branch.md)
 - [Bucket](docs/Bucket.md)
 - [BucketDiskStats](docs/BucketDiskStats.md)
 - [BucketUser](docs/BucketUser.md)
 - [BucketWebsiteConfig](docs/BucketWebsiteConfig.md)
 - [BucketWebsiteConfigErrorPagesInner](docs/BucketWebsiteConfigErrorPagesInner.md)
 - [Certificate](docs/Certificate.md)
 - [CertificateTask](docs/CertificateTask.md)
 - [ChangeAgentModel](docs/ChangeAgentModel.md)
 - [Chat](docs/Chat.md)
 - [ChatFirstMessage](docs/ChatFirstMessage.md)
 - [ChatListMetrics](docs/ChatListMetrics.md)
 - [ChatMessage](docs/ChatMessage.md)
 - [ChatMessageFile](docs/ChatMessageFile.md)
 - [ChatsCsv](docs/ChatsCsv.md)
 - [CheckDomain200Response](docs/CheckDomain200Response.md)
 - [ClearCache](docs/ClearCache.md)
 - [Clickhouse](docs/Clickhouse.md)
 - [ClusterAction](docs/ClusterAction.md)
 - [ClusterEdit](docs/ClusterEdit.md)
 - [ClusterEditOidcProvider](docs/ClusterEditOidcProvider.md)
 - [ClusterIn](docs/ClusterIn.md)
 - [ClusterIn1](docs/ClusterIn1.md)
 - [ClusterInClusterNetworkCidr](docs/ClusterInClusterNetworkCidr.md)
 - [ClusterInConfiguration](docs/ClusterInConfiguration.md)
 - [ClusterInMaintenanceSlot](docs/ClusterInMaintenanceSlot.md)
 - [ClusterInOidcProvider](docs/ClusterInOidcProvider.md)
 - [ClusterOut](docs/ClusterOut.md)
 - [ClusterResponse](docs/ClusterResponse.md)
 - [ClusterVersionEdit](docs/ClusterVersionEdit.md)
 - [Clusterk8s](docs/Clusterk8s.md)
 - [ClustersResponse](docs/ClustersResponse.md)
 - [Commit](docs/Commit.md)
 - [ComponentsSchemasBaseError](docs/ComponentsSchemasBaseError.md)
 - [ComponentsSchemasMeta](docs/ComponentsSchemasMeta.md)
 - [ConfigAccess](docs/ConfigAccess.md)
 - [ConfigCache](docs/ConfigCache.md)
 - [ConfigCacheAlwaysOnline](docs/ConfigCacheAlwaysOnline.md)
 - [ConfigCacheBrowser](docs/ConfigCacheBrowser.md)
 - [ConfigCacheCdn](docs/ConfigCacheCdn.md)
 - [ConfigCacheCdnTtl](docs/ConfigCacheCdnTtl.md)
 - [ConfigCacheQueryArgs](docs/ConfigCacheQueryArgs.md)
 - [ConfigDelivery](docs/ConfigDelivery.md)
 - [ConfigDeliveryPackaging](docs/ConfigDeliveryPackaging.md)
 - [ConfigDeliveryPackagingMp4](docs/ConfigDeliveryPackagingMp4.md)
 - [ConfigDomains](docs/ConfigDomains.md)
 - [ConfigHttpHeaders](docs/ConfigHttpHeaders.md)
 - [ConfigHttpHeadersCors](docs/ConfigHttpHeadersCors.md)
 - [ConfigOrigin](docs/ConfigOrigin.md)
 - [ConfigOriginAws](docs/ConfigOriginAws.md)
 - [ConfigRobots](docs/ConfigRobots.md)
 - [ConfigSecurity](docs/ConfigSecurity.md)
 - [ConfigSecuritySecureToken](docs/ConfigSecuritySecureToken.md)
 - [ConnectAgentTool](docs/ConnectAgentTool.md)
 - [ConnectSkillItem](docs/ConnectSkillItem.md)
 - [ConnectSkills](docs/ConnectSkills.md)
 - [ConnectTelegram](docs/ConnectTelegram.md)
 - [ContainerRegistryPresetsInner](docs/ContainerRegistryPresetsInner.md)
 - [ContainerRegistryRepositoriesInner](docs/ContainerRegistryRepositoriesInner.md)
 - [CreateAdmin](docs/CreateAdmin.md)
 - [CreateAgent](docs/CreateAgent.md)
 - [CreateAgent201Response](docs/CreateAgent201Response.md)
 - [CreateAgentBotTelegramBot201Response](docs/CreateAgentBotTelegramBot201Response.md)
 - [CreateAgentEntityLimit](docs/CreateAgentEntityLimit.md)
 - [CreateAgentEntityLimit201Response](docs/CreateAgentEntityLimit201Response.md)
 - [CreateAgentSkill201Response](docs/CreateAgentSkill201Response.md)
 - [CreateAgentTool201Response](docs/CreateAgentTool201Response.md)
 - [CreateAgentV2](docs/CreateAgentV2.md)
 - [CreateAiKey](docs/CreateAiKey.md)
 - [CreateAiKey201Response](docs/CreateAiKey201Response.md)
 - [CreateAiKeyLimit](docs/CreateAiKeyLimit.md)
 - [CreateAiKeyLimit201Response](docs/CreateAiKeyLimit201Response.md)
 - [CreateAiSolution](docs/CreateAiSolution.md)
 - [CreateAiSolution201Response](docs/CreateAiSolution201Response.md)
 - [CreateApp](docs/CreateApp.md)
 - [CreateApp201Response](docs/CreateApp201Response.md)
 - [CreateBalancer](docs/CreateBalancer.md)
 - [CreateBalancer200Response](docs/CreateBalancer200Response.md)
 - [CreateBalancerCertificates](docs/CreateBalancerCertificates.md)
 - [CreateBalancerRule200Response](docs/CreateBalancerRule200Response.md)
 - [CreateBot](docs/CreateBot.md)
 - [CreateCdnResource201Response](docs/CreateCdnResource201Response.md)
 - [CreateCluster](docs/CreateCluster.md)
 - [CreateClusterAdmin](docs/CreateClusterAdmin.md)
 - [CreateClusterBackupSchedule](docs/CreateClusterBackupSchedule.md)
 - [CreateClusterConfiguration](docs/CreateClusterConfiguration.md)
 - [CreateClusterDiskAutoscaling](docs/CreateClusterDiskAutoscaling.md)
 - [CreateClusterInstance](docs/CreateClusterInstance.md)
 - [CreateClusterMaintenanceSlot](docs/CreateClusterMaintenanceSlot.md)
 - [CreateCrispIntegration](docs/CreateCrispIntegration.md)
 - [CreateDatabaseBackup201Response](docs/CreateDatabaseBackup201Response.md)
 - [CreateDatabaseBackupDownloadUrl201Response](docs/CreateDatabaseBackupDownloadUrl201Response.md)
 - [CreateDatabaseCluster201Response](docs/CreateDatabaseCluster201Response.md)
 - [CreateDatabaseInstance201Response](docs/CreateDatabaseInstance201Response.md)
 - [CreateDatabaseS3Backup201Response](docs/CreateDatabaseS3Backup201Response.md)
 - [CreateDatabaseUser201Response](docs/CreateDatabaseUser201Response.md)
 - [CreateDbAutoBackups](docs/CreateDbAutoBackups.md)
 - [CreateDedicatedServer](docs/CreateDedicatedServer.md)
 - [CreateDedicatedServer201Response](docs/CreateDedicatedServer201Response.md)
 - [CreateDeploy201Response](docs/CreateDeploy201Response.md)
 - [CreateDeployRequest](docs/CreateDeployRequest.md)
 - [CreateDns](docs/CreateDns.md)
 - [CreateDnsV2](docs/CreateDnsV2.md)
 - [CreateDocument](docs/CreateDocument.md)
 - [CreateDomainDNSRecord201Response](docs/CreateDomainDNSRecord201Response.md)
 - [CreateDomainDNSRecordV2201Response](docs/CreateDomainDNSRecordV2201Response.md)
 - [CreateDomainMailbox201Response](docs/CreateDomainMailbox201Response.md)
 - [CreateDomainMailboxRequest](docs/CreateDomainMailboxRequest.md)
 - [CreateDomainMailboxV2201Response](docs/CreateDomainMailboxV2201Response.md)
 - [CreateDomainMailboxV2Request](docs/CreateDomainMailboxV2Request.md)
 - [CreateDomainRequest201Response](docs/CreateDomainRequest201Response.md)
 - [CreateFloatingIp](docs/CreateFloatingIp.md)
 - [CreateFloatingIp201Response](docs/CreateFloatingIp201Response.md)
 - [CreateHttpResource](docs/CreateHttpResource.md)
 - [CreateInstance](docs/CreateInstance.md)
 - [CreateJivoIntegration](docs/CreateJivoIntegration.md)
 - [CreateKey201Response](docs/CreateKey201Response.md)
 - [CreateKeyRequest](docs/CreateKeyRequest.md)
 - [CreateKnowledgebase](docs/CreateKnowledgebase.md)
 - [CreateKnowledgebase201Response](docs/CreateKnowledgebase201Response.md)
 - [CreateKnowledgebaseDocument201Response](docs/CreateKnowledgebaseDocument201Response.md)
 - [CreateMaxBot](docs/CreateMaxBot.md)
 - [CreateMultipleDomainMailboxes201Response](docs/CreateMultipleDomainMailboxes201Response.md)
 - [CreateMultipleDomainMailboxesRequest](docs/CreateMultipleDomainMailboxesRequest.md)
 - [CreateMultipleDomainMailboxesRequestMailboxesInner](docs/CreateMultipleDomainMailboxesRequestMailboxesInner.md)
 - [CreateMultipleDomainMailboxesV2201Response](docs/CreateMultipleDomainMailboxesV2201Response.md)
 - [CreateMultipleDomainMailboxesV2RequestInner](docs/CreateMultipleDomainMailboxesV2RequestInner.md)
 - [CreateNetworkDrive](docs/CreateNetworkDrive.md)
 - [CreateNetworkDrive201Response](docs/CreateNetworkDrive201Response.md)
 - [CreatePerson201Response](docs/CreatePerson201Response.md)
 - [CreateProject](docs/CreateProject.md)
 - [CreateProject201Response](docs/CreateProject201Response.md)
 - [CreateRule](docs/CreateRule.md)
 - [CreateS3Backup](docs/CreateS3Backup.md)
 - [CreateServer](docs/CreateServer.md)
 - [CreateServer201Response](docs/CreateServer201Response.md)
 - [CreateServerConfiguration](docs/CreateServerConfiguration.md)
 - [CreateServerDisk201Response](docs/CreateServerDisk201Response.md)
 - [CreateServerDiskBackup201Response](docs/CreateServerDiskBackup201Response.md)
 - [CreateServerDiskBackupRequest](docs/CreateServerDiskBackupRequest.md)
 - [CreateServerDiskRequest](docs/CreateServerDiskRequest.md)
 - [CreateServerNetwork](docs/CreateServerNetwork.md)
 - [CreateSkill](docs/CreateSkill.md)
 - [CreateStorage201Response](docs/CreateStorage201Response.md)
 - [CreateStorageRequest](docs/CreateStorageRequest.md)
 - [CreateStorageRequestConfigurator](docs/CreateStorageRequestConfigurator.md)
 - [CreateTool](docs/CreateTool.md)
 - [CreateVPC201Response](docs/CreateVPC201Response.md)
 - [CreateVpc](docs/CreateVpc.md)
 - [CrispIntegration](docs/CrispIntegration.md)
 - [DatabaseAdmin](docs/DatabaseAdmin.md)
 - [DatabaseAdminInstancesInner](docs/DatabaseAdminInstancesInner.md)
 - [DatabaseCluster](docs/DatabaseCluster.md)
 - [DatabaseClusterChildServicesInner](docs/DatabaseClusterChildServicesInner.md)
 - [DatabaseClusterDisk](docs/DatabaseClusterDisk.md)
 - [DatabaseClusterDiskAutoscaling](docs/DatabaseClusterDiskAutoscaling.md)
 - [DatabaseClusterDomainsInner](docs/DatabaseClusterDomainsInner.md)
 - [DatabaseClusterMaintenanceSlot](docs/DatabaseClusterMaintenanceSlot.md)
 - [DatabaseClusterNetworksInner](docs/DatabaseClusterNetworksInner.md)
 - [DatabaseClusterNetworksInnerIpsInner](docs/DatabaseClusterNetworksInnerIpsInner.md)
 - [DatabaseClusterParentServicesInner](docs/DatabaseClusterParentServicesInner.md)
 - [DatabaseClusterReplicaListInner](docs/DatabaseClusterReplicaListInner.md)
 - [DatabaseClusterReplicaListInnerDisk](docs/DatabaseClusterReplicaListInnerDisk.md)
 - [DatabaseConfigurator](docs/DatabaseConfigurator.md)
 - [DatabaseConfiguratorPrices](docs/DatabaseConfiguratorPrices.md)
 - [DatabaseConfiguratorRequirements](docs/DatabaseConfiguratorRequirements.md)
 - [DatabaseExtensions](docs/DatabaseExtensions.md)
 - [DatabaseInstance](docs/DatabaseInstance.md)
 - [DatabaseType](docs/DatabaseType.md)
 - [DatabaseTypeRequirements](docs/DatabaseTypeRequirements.md)
 - [Db](docs/Db.md)
 - [DbDiskStats](docs/DbDiskStats.md)
 - [DbExtension](docs/DbExtension.md)
 - [DbParametersByType](docs/DbParametersByType.md)
 - [DbPrivilegeGroup](docs/DbPrivilegeGroup.md)
 - [DbReplication](docs/DbReplication.md)
 - [DbType](docs/DbType.md)
 - [DbsCreateBackup](docs/DbsCreateBackup.md)
 - [DbsUpdateBackup](docs/DbsUpdateBackup.md)
 - [DedicatedServer](docs/DedicatedServer.md)
 - [DedicatedServerAdditionalService](docs/DedicatedServerAdditionalService.md)
 - [DedicatedServerPreset](docs/DedicatedServerPreset.md)
 - [DedicatedServerPresetCpu](docs/DedicatedServerPresetCpu.md)
 - [DedicatedServerPresetDisk](docs/DedicatedServerPresetDisk.md)
 - [DedicatedServerPresetMemory](docs/DedicatedServerPresetMemory.md)
 - [DeleteAgent](docs/DeleteAgent.md)
 - [DeleteBalancer200Response](docs/DeleteBalancer200Response.md)
 - [DeleteCluster200Response](docs/DeleteCluster200Response.md)
 - [DeleteCountriesFromAllowedList200Response](docs/DeleteCountriesFromAllowedList200Response.md)
 - [DeleteCountriesFromAllowedListRequest](docs/DeleteCountriesFromAllowedListRequest.md)
 - [DeleteDatabaseCluster200Response](docs/DeleteDatabaseCluster200Response.md)
 - [DeleteDatabaseCluster200ResponseDatabaseDelete](docs/DeleteDatabaseCluster200ResponseDatabaseDelete.md)
 - [DeleteIPsFromAllowedList200Response](docs/DeleteIPsFromAllowedList200Response.md)
 - [DeleteIPsFromAllowedListRequest](docs/DeleteIPsFromAllowedListRequest.md)
 - [DeleteServer200Response](docs/DeleteServer200Response.md)
 - [DeleteServerIPRequest](docs/DeleteServerIPRequest.md)
 - [DeleteServiceResponse](docs/DeleteServiceResponse.md)
 - [DeleteStorage200Response](docs/DeleteStorage200Response.md)
 - [Deploy](docs/Deploy.md)
 - [DeploySettingsInner](docs/DeploySettingsInner.md)
 - [DeployStatus](docs/DeployStatus.md)
 - [DnatIn](docs/DnatIn.md)
 - [DnatInLocal](docs/DnatInLocal.md)
 - [DnatInPublic](docs/DnatInPublic.md)
 - [DnatRuleOut](docs/DnatRuleOut.md)
 - [DnatRuleResponse](docs/DnatRuleResponse.md)
 - [DnatRulesResponse](docs/DnatRulesResponse.md)
 - [DnsRecord](docs/DnsRecord.md)
 - [DnsRecordData](docs/DnsRecordData.md)
 - [DnsRecordV2](docs/DnsRecordV2.md)
 - [DnsRecordV2Data](docs/DnsRecordV2Data.md)
 - [Document](docs/Document.md)
 - [DocumentStatusInfo](docs/DocumentStatusInfo.md)
 - [DocumentStatusInfoDetails](docs/DocumentStatusInfoDetails.md)
 - [Domain](docs/Domain.md)
 - [DomainAllowedBuyPeriodsInner](docs/DomainAllowedBuyPeriodsInner.md)
 - [DomainInfo](docs/DomainInfo.md)
 - [DomainNameServer](docs/DomainNameServer.md)
 - [DomainNameServerItemsInner](docs/DomainNameServerItemsInner.md)
 - [DomainPaymentPeriod](docs/DomainPaymentPeriod.md)
 - [DomainPrimeType](docs/DomainPrimeType.md)
 - [DomainProlong](docs/DomainProlong.md)
 - [DomainRegister](docs/DomainRegister.md)
 - [DomainRequest](docs/DomainRequest.md)
 - [DomainTransfer](docs/DomainTransfer.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [ExportAgentChatsCsv200Response](docs/ExportAgentChatsCsv200Response.md)
 - [Finances](docs/Finances.md)
 - [FirewallGroup](docs/FirewallGroup.md)
 - [FirewallGroupInAPI](docs/FirewallGroupInAPI.md)
 - [FirewallGroupOutResponse](docs/FirewallGroupOutResponse.md)
 - [FirewallGroupResource](docs/FirewallGroupResource.md)
 - [FirewallGroupResourceOutResponse](docs/FirewallGroupResourceOutResponse.md)
 - [FirewallGroupResourcesOutResponse](docs/FirewallGroupResourcesOutResponse.md)
 - [FirewallGroupsOutResponse](docs/FirewallGroupsOutResponse.md)
 - [FirewallRule](docs/FirewallRule.md)
 - [FirewallRuleDirection](docs/FirewallRuleDirection.md)
 - [FirewallRuleInAPI](docs/FirewallRuleInAPI.md)
 - [FirewallRuleOutResponse](docs/FirewallRuleOutResponse.md)
 - [FirewallRuleProtocol](docs/FirewallRuleProtocol.md)
 - [FirewallRulesOutResponse](docs/FirewallRulesOutResponse.md)
 - [FloatingIp](docs/FloatingIp.md)
 - [ForwardIsDisabled](docs/ForwardIsDisabled.md)
 - [ForwardIsEnabled](docs/ForwardIsEnabled.md)
 - [ForwardingIncomingIsDisabled](docs/ForwardingIncomingIsDisabled.md)
 - [ForwardingIncomingIsEnabled](docs/ForwardingIncomingIsEnabled.md)
 - [ForwardingOutgoingIsDisabled](docs/ForwardingOutgoingIsDisabled.md)
 - [ForwardingOutgoingIsEnabled](docs/ForwardingOutgoingIsEnabled.md)
 - [Frameworks](docs/Frameworks.md)
 - [Free](docs/Free.md)
 - [GetAccountStatus200Response](docs/GetAccountStatus200Response.md)
 - [GetAccountStatus403Response](docs/GetAccountStatus403Response.md)
 - [GetAgentBotMaxBots200Response](docs/GetAgentBotMaxBots200Response.md)
 - [GetAgentBotMaxBots200ResponseMeta](docs/GetAgentBotMaxBots200ResponseMeta.md)
 - [GetAgentBotMaxBotsByAgent200Response](docs/GetAgentBotMaxBotsByAgent200Response.md)
 - [GetAgentBotMaxBotsByAgent200ResponseMeta](docs/GetAgentBotMaxBotsByAgent200ResponseMeta.md)
 - [GetAgentBotTelegramBots200Response](docs/GetAgentBotTelegramBots200Response.md)
 - [GetAgentBotTelegramBots200ResponseMeta](docs/GetAgentBotTelegramBots200ResponseMeta.md)
 - [GetAgentBotTelegramBotsByAgent200Response](docs/GetAgentBotTelegramBotsByAgent200Response.md)
 - [GetAgentBotTelegramBotsByAgent200ResponseMeta](docs/GetAgentBotTelegramBotsByAgent200ResponseMeta.md)
 - [GetAgentChatList200Response](docs/GetAgentChatList200Response.md)
 - [GetAgentChatList200ResponseMeta](docs/GetAgentChatList200ResponseMeta.md)
 - [GetAgentChatMessages200Response](docs/GetAgentChatMessages200Response.md)
 - [GetAgentChatMessages200ResponseMeta](docs/GetAgentChatMessages200ResponseMeta.md)
 - [GetAgentEntityLimits200Response](docs/GetAgentEntityLimits200Response.md)
 - [GetAgentEntityLimits200ResponseMeta](docs/GetAgentEntityLimits200ResponseMeta.md)
 - [GetAgentIntegrationCrisp200Response](docs/GetAgentIntegrationCrisp200Response.md)
 - [GetAgentIntegrationJivo200Response](docs/GetAgentIntegrationJivo200Response.md)
 - [GetAgentSkill200Response](docs/GetAgentSkill200Response.md)
 - [GetAgentSkillCatalog200Response](docs/GetAgentSkillCatalog200Response.md)
 - [GetAgentSkillCatalog200ResponseMeta](docs/GetAgentSkillCatalog200ResponseMeta.md)
 - [GetAgentSkillLibrary200Response](docs/GetAgentSkillLibrary200Response.md)
 - [GetAgentSkillLibrary200ResponseMeta](docs/GetAgentSkillLibrary200ResponseMeta.md)
 - [GetAgentSkillVersion200Response](docs/GetAgentSkillVersion200Response.md)
 - [GetAgentSkillVersions200Response](docs/GetAgentSkillVersions200Response.md)
 - [GetAgentSkillVersions200ResponseMeta](docs/GetAgentSkillVersions200ResponseMeta.md)
 - [GetAgentSkillsForAgent200Response](docs/GetAgentSkillsForAgent200Response.md)
 - [GetAgentSkillsForAgent200ResponseMeta](docs/GetAgentSkillsForAgent200ResponseMeta.md)
 - [GetAgentStatistics200Response](docs/GetAgentStatistics200Response.md)
 - [GetAgentStatistics200ResponseMeta](docs/GetAgentStatistics200ResponseMeta.md)
 - [GetAgentToolTemplatePrices200Response](docs/GetAgentToolTemplatePrices200Response.md)
 - [GetAgentToolTemplates200Response](docs/GetAgentToolTemplates200Response.md)
 - [GetAgentToolTemplates200ResponseMeta](docs/GetAgentToolTemplates200ResponseMeta.md)
 - [GetAgentTools200Response](docs/GetAgentTools200Response.md)
 - [GetAgentTools200ResponseMeta](docs/GetAgentTools200ResponseMeta.md)
 - [GetAgentToolsByAgent200Response](docs/GetAgentToolsByAgent200Response.md)
 - [GetAgentToolsByAgent200ResponseMeta](docs/GetAgentToolsByAgent200ResponseMeta.md)
 - [GetAgentV3200Response](docs/GetAgentV3200Response.md)
 - [GetAgents200Response](docs/GetAgents200Response.md)
 - [GetAgents200ResponseMeta](docs/GetAgents200ResponseMeta.md)
 - [GetAgentsTokenPackages200Response](docs/GetAgentsTokenPackages200Response.md)
 - [GetAgentsTokenPackages200ResponseMeta](docs/GetAgentsTokenPackages200ResponseMeta.md)
 - [GetAgentsV3200Response](docs/GetAgentsV3200Response.md)
 - [GetAiKeyCallsStatistic200Response](docs/GetAiKeyCallsStatistic200Response.md)
 - [GetAiKeyLimits200Response](docs/GetAiKeyLimits200Response.md)
 - [GetAiKeyLimits200ResponseMeta](docs/GetAiKeyLimits200ResponseMeta.md)
 - [GetAiKeyModelsStatistic200Response](docs/GetAiKeyModelsStatistic200Response.md)
 - [GetAiKeyModelsStatistic200ResponseMeta](docs/GetAiKeyModelsStatistic200ResponseMeta.md)
 - [GetAiKeyTokenUsageStatistic200Response](docs/GetAiKeyTokenUsageStatistic200Response.md)
 - [GetAiKeyTokenUsageStatistic200ResponseMeta](docs/GetAiKeyTokenUsageStatistic200ResponseMeta.md)
 - [GetAiKeys200Response](docs/GetAiKeys200Response.md)
 - [GetAiKeys200ResponseMeta](docs/GetAiKeys200ResponseMeta.md)
 - [GetAiModelParameters200Response](docs/GetAiModelParameters200Response.md)
 - [GetAiModelParameters200ResponseMeta](docs/GetAiModelParameters200ResponseMeta.md)
 - [GetAiModelTokenPackages200Response](docs/GetAiModelTokenPackages200Response.md)
 - [GetAiModelTokenPackages200ResponseMeta](docs/GetAiModelTokenPackages200ResponseMeta.md)
 - [GetAiProvider200Response](docs/GetAiProvider200Response.md)
 - [GetAiProviders200Response](docs/GetAiProviders200Response.md)
 - [GetAiProviders200ResponseMeta](docs/GetAiProviders200ResponseMeta.md)
 - [GetAiSolutionChatStatus200Response](docs/GetAiSolutionChatStatus200Response.md)
 - [GetAiSolutionMainChatHistory200Response](docs/GetAiSolutionMainChatHistory200Response.md)
 - [GetAllMailboxesV2200Response](docs/GetAllMailboxesV2200Response.md)
 - [GetAllProjectResources200Response](docs/GetAllProjectResources200Response.md)
 - [GetAppDeploys200Response](docs/GetAppDeploys200Response.md)
 - [GetAppLogs200Response](docs/GetAppLogs200Response.md)
 - [GetApps200Response](docs/GetApps200Response.md)
 - [GetAuthAccessSettings200Response](docs/GetAuthAccessSettings200Response.md)
 - [GetAuthAccessSettings200ResponseWhiteList](docs/GetAuthAccessSettings200ResponseWhiteList.md)
 - [GetBalancerIPs200Response](docs/GetBalancerIPs200Response.md)
 - [GetBalancerRules200Response](docs/GetBalancerRules200Response.md)
 - [GetBalancers200Response](docs/GetBalancers200Response.md)
 - [GetBalancersPresets200Response](docs/GetBalancersPresets200Response.md)
 - [GetBranches200Response](docs/GetBranches200Response.md)
 - [GetCdnCertificateTasks200Response](docs/GetCdnCertificateTasks200Response.md)
 - [GetCdnCertificates200Response](docs/GetCdnCertificates200Response.md)
 - [GetCdnOriginNodes200Response](docs/GetCdnOriginNodes200Response.md)
 - [GetCdnPresets200Response](docs/GetCdnPresets200Response.md)
 - [GetCdnResourceConfiguration200Response](docs/GetCdnResourceConfiguration200Response.md)
 - [GetCdnResourceNodes200Response](docs/GetCdnResourceNodes200Response.md)
 - [GetCdnResourceStatistics200Response](docs/GetCdnResourceStatistics200Response.md)
 - [GetCdnResources200Response](docs/GetCdnResources200Response.md)
 - [GetCommits200Response](docs/GetCommits200Response.md)
 - [GetConfigurators200Response](docs/GetConfigurators200Response.md)
 - [GetCountries200Response](docs/GetCountries200Response.md)
 - [GetDatabaseAutoBackupsSettings200Response](docs/GetDatabaseAutoBackupsSettings200Response.md)
 - [GetDatabaseBackup200Response](docs/GetDatabaseBackup200Response.md)
 - [GetDatabaseBackups200Response](docs/GetDatabaseBackups200Response.md)
 - [GetDatabaseClusterReplicas200Response](docs/GetDatabaseClusterReplicas200Response.md)
 - [GetDatabaseClusterTypes200Response](docs/GetDatabaseClusterTypes200Response.md)
 - [GetDatabaseClusters200Response](docs/GetDatabaseClusters200Response.md)
 - [GetDatabaseConfigurators200Response](docs/GetDatabaseConfigurators200Response.md)
 - [GetDatabaseDefaultParameters200Response](docs/GetDatabaseDefaultParameters200Response.md)
 - [GetDatabaseDefaultParameters200ResponseConfigParams](docs/GetDatabaseDefaultParameters200ResponseConfigParams.md)
 - [GetDatabaseInstances200Response](docs/GetDatabaseInstances200Response.md)
 - [GetDatabasePreset200Response](docs/GetDatabasePreset200Response.md)
 - [GetDatabasePrivileges200Response](docs/GetDatabasePrivileges200Response.md)
 - [GetDatabaseS3Backups200Response](docs/GetDatabaseS3Backups200Response.md)
 - [GetDatabaseUsers200Response](docs/GetDatabaseUsers200Response.md)
 - [GetDatabasesPresets200Response](docs/GetDatabasesPresets200Response.md)
 - [GetDedicatedServerPresetAdditionalServices200Response](docs/GetDedicatedServerPresetAdditionalServices200Response.md)
 - [GetDedicatedServers200Response](docs/GetDedicatedServers200Response.md)
 - [GetDedicatedServersPresets200Response](docs/GetDedicatedServersPresets200Response.md)
 - [GetDeployLogs200Response](docs/GetDeployLogs200Response.md)
 - [GetDeploySettings200Response](docs/GetDeploySettings200Response.md)
 - [GetDomain200Response](docs/GetDomain200Response.md)
 - [GetDomainDNSRecords200Response](docs/GetDomainDNSRecords200Response.md)
 - [GetDomainMailInfo200Response](docs/GetDomainMailInfo200Response.md)
 - [GetDomainNameServers200Response](docs/GetDomainNameServers200Response.md)
 - [GetDomainRequests200Response](docs/GetDomainRequests200Response.md)
 - [GetDomains200Response](docs/GetDomains200Response.md)
 - [GetFinances200Response](docs/GetFinances200Response.md)
 - [GetFinances400Response](docs/GetFinances400Response.md)
 - [GetFinances401Response](docs/GetFinances401Response.md)
 - [GetFinances429Response](docs/GetFinances429Response.md)
 - [GetFinances500Response](docs/GetFinances500Response.md)
 - [GetFloatingIps200Response](docs/GetFloatingIps200Response.md)
 - [GetImage404Response](docs/GetImage404Response.md)
 - [GetKey200Response](docs/GetKey200Response.md)
 - [GetKeys200Response](docs/GetKeys200Response.md)
 - [GetKnowledgebaseDocumentsV2200Response](docs/GetKnowledgebaseDocumentsV2200Response.md)
 - [GetKnowledgebaseDocumentsV2200ResponseMeta](docs/GetKnowledgebaseDocumentsV2200ResponseMeta.md)
 - [GetKnowledgebaseV3200Response](docs/GetKnowledgebaseV3200Response.md)
 - [GetKnowledgebases200Response](docs/GetKnowledgebases200Response.md)
 - [GetKnowledgebasesTokenPackages200Response](docs/GetKnowledgebasesTokenPackages200Response.md)
 - [GetKnowledgebasesTokenPackages200ResponseMeta](docs/GetKnowledgebasesTokenPackages200ResponseMeta.md)
 - [GetKnowledgebasesV2200Response](docs/GetKnowledgebasesV2200Response.md)
 - [GetKnowledgebasesV3200Response](docs/GetKnowledgebasesV3200Response.md)
 - [GetKnowledgebasesV3200ResponseMeta](docs/GetKnowledgebasesV3200ResponseMeta.md)
 - [GetLocations200Response](docs/GetLocations200Response.md)
 - [GetMailboxes200Response](docs/GetMailboxes200Response.md)
 - [GetModels200Response](docs/GetModels200Response.md)
 - [GetModelsV3200Response](docs/GetModelsV3200Response.md)
 - [GetModelsV3200ResponseMeta](docs/GetModelsV3200ResponseMeta.md)
 - [GetNetworkDrives200Response](docs/GetNetworkDrives200Response.md)
 - [GetNetworkDrivesAvailableResources200Response](docs/GetNetworkDrivesAvailableResources200Response.md)
 - [GetNetworkDrivesPresets200Response](docs/GetNetworkDrivesPresets200Response.md)
 - [GetNotificationSettings200Response](docs/GetNotificationSettings200Response.md)
 - [GetOsList200Response](docs/GetOsList200Response.md)
 - [GetPersons200Response](docs/GetPersons200Response.md)
 - [GetProjectBalancers200Response](docs/GetProjectBalancers200Response.md)
 - [GetProjectClusters200Response](docs/GetProjectClusters200Response.md)
 - [GetProjectDatabases200Response](docs/GetProjectDatabases200Response.md)
 - [GetProjectDedicatedServers200Response](docs/GetProjectDedicatedServers200Response.md)
 - [GetProjectServers200Response](docs/GetProjectServers200Response.md)
 - [GetProjectStorages200Response](docs/GetProjectStorages200Response.md)
 - [GetProjects200Response](docs/GetProjects200Response.md)
 - [GetProviders200Response](docs/GetProviders200Response.md)
 - [GetRepositories200Response](docs/GetRepositories200Response.md)
 - [GetRestorePoint200Response](docs/GetRestorePoint200Response.md)
 - [GetRestorePoints200Response](docs/GetRestorePoints200Response.md)
 - [GetServerDiskBackup200Response](docs/GetServerDiskBackup200Response.md)
 - [GetServerDiskBackups200Response](docs/GetServerDiskBackups200Response.md)
 - [GetServerDisks200Response](docs/GetServerDisks200Response.md)
 - [GetServerIPs200Response](docs/GetServerIPs200Response.md)
 - [GetServerLogs200Response](docs/GetServerLogs200Response.md)
 - [GetServerStatistics200Response](docs/GetServerStatistics200Response.md)
 - [GetServerStatistics200ResponseCpuInner](docs/GetServerStatistics200ResponseCpuInner.md)
 - [GetServerStatistics200ResponseDiskInner](docs/GetServerStatistics200ResponseDiskInner.md)
 - [GetServerStatistics200ResponseNetworkTrafficInner](docs/GetServerStatistics200ResponseNetworkTrafficInner.md)
 - [GetServerStatistics200ResponseRamInner](docs/GetServerStatistics200ResponseRamInner.md)
 - [GetServerStatisticsNew200Response](docs/GetServerStatisticsNew200Response.md)
 - [GetServers200Response](docs/GetServers200Response.md)
 - [GetServersPresets200Response](docs/GetServersPresets200Response.md)
 - [GetServicePrices200Response](docs/GetServicePrices200Response.md)
 - [GetServicePrices200ResponseMeta](docs/GetServicePrices200ResponseMeta.md)
 - [GetSoftware200Response](docs/GetSoftware200Response.md)
 - [GetStorageSubdomains200Response](docs/GetStorageSubdomains200Response.md)
 - [GetStorageTransferStatus200Response](docs/GetStorageTransferStatus200Response.md)
 - [GetStorageUsers200Response](docs/GetStorageUsers200Response.md)
 - [GetStoragesPresets200Response](docs/GetStoragesPresets200Response.md)
 - [GetTLD200Response](docs/GetTLD200Response.md)
 - [GetTLDs200Response](docs/GetTLDs200Response.md)
 - [GetVPCPorts200Response](docs/GetVPCPorts200Response.md)
 - [GetVPCServices200Response](docs/GetVPCServices200Response.md)
 - [GetVPCs200Response](docs/GetVPCs200Response.md)
 - [HttpResource](docs/HttpResource.md)
 - [HttpResourceConfig](docs/HttpResourceConfig.md)
 - [Image](docs/Image.md)
 - [ImageDownload](docs/ImageDownload.md)
 - [ImageDownloadResponse](docs/ImageDownloadResponse.md)
 - [ImageDownloadsResponse](docs/ImageDownloadsResponse.md)
 - [ImageInAPI](docs/ImageInAPI.md)
 - [ImageOutResponse](docs/ImageOutResponse.md)
 - [ImageStatus](docs/ImageStatus.md)
 - [ImageUpdateAPI](docs/ImageUpdateAPI.md)
 - [ImageUrlAuth](docs/ImageUrlAuth.md)
 - [ImageUrlIn](docs/ImageUrlIn.md)
 - [ImagesOutResponse](docs/ImagesOutResponse.md)
 - [IncreaseNodes](docs/IncreaseNodes.md)
 - [IndexingSchedule](docs/IndexingSchedule.md)
 - [InfoServicePrice](docs/InfoServicePrice.md)
 - [InitializeParams](docs/InitializeParams.md)
 - [InitializeParamsClientInfo](docs/InitializeParamsClientInfo.md)
 - [InitializeResult](docs/InitializeResult.md)
 - [InitializeResultServerInfo](docs/InitializeResultServerInfo.md)
 - [InstallAgentSkillFromCatalog200Response](docs/InstallAgentSkillFromCatalog200Response.md)
 - [InstallAgentSkillFromCatalog200ResponseMeta](docs/InstallAgentSkillFromCatalog200ResponseMeta.md)
 - [InstallAndConnectAgentSkill200Response](docs/InstallAndConnectAgentSkill200Response.md)
 - [InstallAndConnectSkill](docs/InstallAndConnectSkill.md)
 - [InstallTool](docs/InstallTool.md)
 - [Invoice](docs/Invoice.md)
 - [Ip](docs/Ip.md)
 - [IssueCertificate](docs/IssueCertificate.md)
 - [JivoIntegration](docs/JivoIntegration.md)
 - [JsonRpcError](docs/JsonRpcError.md)
 - [JsonRpcErrorResponse](docs/JsonRpcErrorResponse.md)
 - [JsonRpcNotification](docs/JsonRpcNotification.md)
 - [JsonRpcRequest](docs/JsonRpcRequest.md)
 - [JsonRpcResultResponse](docs/JsonRpcResultResponse.md)
 - [K8SVersionsResponse](docs/K8SVersionsResponse.md)
 - [Kafka](docs/Kafka.md)
 - [KafkaConfigParameters](docs/KafkaConfigParameters.md)
 - [Knowledgebase](docs/Knowledgebase.md)
 - [KnowledgebaseAddTokenPackage](docs/KnowledgebaseAddTokenPackage.md)
 - [KnowledgebaseV2](docs/KnowledgebaseV2.md)
 - [KnowledgebaseV3](docs/KnowledgebaseV3.md)
 - [LinkKnowledgebase](docs/LinkKnowledgebase.md)
 - [ListAiSolutionConversations200Response](docs/ListAiSolutionConversations200Response.md)
 - [ListAiSolutionTelegramBots200Response](docs/ListAiSolutionTelegramBots200Response.md)
 - [ListAiSolutionTypes200Response](docs/ListAiSolutionTypes200Response.md)
 - [ListAiSolutions200Response](docs/ListAiSolutions200Response.md)
 - [Location](docs/Location.md)
 - [LocationDto](docs/LocationDto.md)
 - [Mailbox](docs/Mailbox.md)
 - [MailboxAutoReply](docs/MailboxAutoReply.md)
 - [MailboxForwardingIncoming](docs/MailboxForwardingIncoming.md)
 - [MailboxForwardingOutgoing](docs/MailboxForwardingOutgoing.md)
 - [MailboxResponse](docs/MailboxResponse.md)
 - [MailboxSpamFilter](docs/MailboxSpamFilter.md)
 - [MailboxV2](docs/MailboxV2.md)
 - [MailboxesBatchV2](docs/MailboxesBatchV2.md)
 - [MasterPresetOutApi](docs/MasterPresetOutApi.md)
 - [Meta](docs/Meta.md)
 - [Meta1](docs/Meta1.md)
 - [Model](docs/Model.md)
 - [ModelBatch](docs/ModelBatch.md)
 - [ModelParamInfo](docs/ModelParamInfo.md)
 - [ModelParameter](docs/ModelParameter.md)
 - [ModelPriceTier](docs/ModelPriceTier.md)
 - [ModelV3](docs/ModelV3.md)
 - [Mongodb](docs/Mongodb.md)
 - [MountNetworkDrive](docs/MountNetworkDrive.md)
 - [Mysql](docs/Mysql.md)
 - [NatIn](docs/NatIn.md)
 - [Network](docs/Network.md)
 - [NetworkDrive](docs/NetworkDrive.md)
 - [NetworkDriveAvailableResource](docs/NetworkDriveAvailableResource.md)
 - [NetworkDrivePreset](docs/NetworkDrivePreset.md)
 - [NetworkDrivePresetRead](docs/NetworkDrivePresetRead.md)
 - [NetworkDrivePresetWrite](docs/NetworkDrivePresetWrite.md)
 - [NetworkDriveServiceListInner](docs/NetworkDriveServiceListInner.md)
 - [NetworkDriversResponse](docs/NetworkDriversResponse.md)
 - [NetworkEdit](docs/NetworkEdit.md)
 - [NetworkIn](docs/NetworkIn.md)
 - [NetworkOut](docs/NetworkOut.md)
 - [NetworkResponse](docs/NetworkResponse.md)
 - [NetworksResponse](docs/NetworksResponse.md)
 - [NodeGroupIn](docs/NodeGroupIn.md)
 - [NodeGroupInConfiguration](docs/NodeGroupInConfiguration.md)
 - [NodeGroupOut](docs/NodeGroupOut.md)
 - [NodeGroupResponse](docs/NodeGroupResponse.md)
 - [NodeGroupsResponse](docs/NodeGroupsResponse.md)
 - [NodeOut](docs/NodeOut.md)
 - [NodesResponse](docs/NodesResponse.md)
 - [NotificationSetting](docs/NotificationSetting.md)
 - [NotificationSettingChannel](docs/NotificationSettingChannel.md)
 - [NotificationSettingChannels](docs/NotificationSettingChannels.md)
 - [NotificationSettingType](docs/NotificationSettingType.md)
 - [OS](docs/OS.md)
 - [Opensearch](docs/Opensearch.md)
 - [Org](docs/Org.md)
 - [OriginNode](docs/OriginNode.md)
 - [OriginServer](docs/OriginServer.md)
 - [OutgoingIsDisabled](docs/OutgoingIsDisabled.md)
 - [OutgoingIsEnabled](docs/OutgoingIsEnabled.md)
 - [PerformActionOnBackupRequest](docs/PerformActionOnBackupRequest.md)
 - [PerformActionOnServerRequest](docs/PerformActionOnServerRequest.md)
 - [Person](docs/Person.md)
 - [Person2](docs/Person2.md)
 - [Policy](docs/Policy.md)
 - [Postgres](docs/Postgres.md)
 - [Postgresql](docs/Postgresql.md)
 - [PreloadCache](docs/PreloadCache.md)
 - [Preset](docs/Preset.md)
 - [PresetsBalancer](docs/PresetsBalancer.md)
 - [PresetsDbs](docs/PresetsDbs.md)
 - [PresetsResponse](docs/PresetsResponse.md)
 - [PresetsStorage](docs/PresetsStorage.md)
 - [Project](docs/Project.md)
 - [ProjectResource](docs/ProjectResource.md)
 - [PropertiesMysql](docs/PropertiesMysql.md)
 - [PropertiesValkey](docs/PropertiesValkey.md)
 - [Provider](docs/Provider.md)
 - [Providers](docs/Providers.md)
 - [Rabbitmq](docs/Rabbitmq.md)
 - [ReduceNodes](docs/ReduceNodes.md)
 - [RegenerateAiSolutionAiKey200Response](docs/RegenerateAiSolutionAiKey200Response.md)
 - [RegistriesResponse](docs/RegistriesResponse.md)
 - [RegistryEdit](docs/RegistryEdit.md)
 - [RegistryIn](docs/RegistryIn.md)
 - [RegistryInConfiguration](docs/RegistryInConfiguration.md)
 - [RegistryOut](docs/RegistryOut.md)
 - [RegistryOutDiskStats](docs/RegistryOutDiskStats.md)
 - [RegistryResponse](docs/RegistryResponse.md)
 - [RemoveCountries](docs/RemoveCountries.md)
 - [RemoveIps](docs/RemoveIps.md)
 - [Replica](docs/Replica.md)
 - [ReplicaDisk](docs/ReplicaDisk.md)
 - [RepositoriesResponse](docs/RepositoriesResponse.md)
 - [Repository](docs/Repository.md)
 - [ResetAiSolutionChat](docs/ResetAiSolutionChat.md)
 - [ResetAiSolutionChat200Response](docs/ResetAiSolutionChat200Response.md)
 - [Resource](docs/Resource.md)
 - [ResourceTransfer](docs/ResourceTransfer.md)
 - [ResourceType](docs/ResourceType.md)
 - [Resources](docs/Resources.md)
 - [ResourcesListResult](docs/ResourcesListResult.md)
 - [ResourcesReadParams](docs/ResourcesReadParams.md)
 - [ResourcesReadResult](docs/ResourcesReadResult.md)
 - [ResourcesReadResultContentsInner](docs/ResourcesReadResultContentsInner.md)
 - [ResourcesResponse](docs/ResourcesResponse.md)
 - [RestorePoint](docs/RestorePoint.md)
 - [RetrievalMode](docs/RetrievalMode.md)
 - [RetrievalModeRef](docs/RetrievalModeRef.md)
 - [RouterEdit](docs/RouterEdit.md)
 - [RouterIn](docs/RouterIn.md)
 - [RouterInIpsInner](docs/RouterInIpsInner.md)
 - [RouterInIpsInnerNat](docs/RouterInIpsInnerNat.md)
 - [RouterInNetworksInner](docs/RouterInNetworksInner.md)
 - [RouterInParentService](docs/RouterInParentService.md)
 - [RouterNetworkMeta](docs/RouterNetworkMeta.md)
 - [RouterNetworkMetaDhcp](docs/RouterNetworkMetaDhcp.md)
 - [RouterOut](docs/RouterOut.md)
 - [RouterOutIpsInner](docs/RouterOutIpsInner.md)
 - [RouterOutIpsInnerNat](docs/RouterOutIpsInnerNat.md)
 - [RouterOutNodesInner](docs/RouterOutNodesInner.md)
 - [RouterOutParentServicesInner](docs/RouterOutParentServicesInner.md)
 - [RouterPreset](docs/RouterPreset.md)
 - [RouterPresetsResponse](docs/RouterPresetsResponse.md)
 - [RouterResponse](docs/RouterResponse.md)
 - [RouterStatistic](docs/RouterStatistic.md)
 - [RouterStatisticListInner](docs/RouterStatisticListInner.md)
 - [RouterStatisticMeta](docs/RouterStatisticMeta.md)
 - [RouterStatisticsResponse](docs/RouterStatisticsResponse.md)
 - [RoutersResponse](docs/RoutersResponse.md)
 - [Rule](docs/Rule.md)
 - [S3Backup](docs/S3Backup.md)
 - [S3Subdomain](docs/S3Subdomain.md)
 - [SchemasBaseError](docs/SchemasBaseError.md)
 - [SchemasMeta](docs/SchemasMeta.md)
 - [SchemasPresetsResponse](docs/SchemasPresetsResponse.md)
 - [SendAiSolutionMessage](docs/SendAiSolutionMessage.md)
 - [SendAiSolutionMessage200Response](docs/SendAiSolutionMessage200Response.md)
 - [ServerBackup](docs/ServerBackup.md)
 - [ServerDisk](docs/ServerDisk.md)
 - [ServerIp](docs/ServerIp.md)
 - [ServerLog](docs/ServerLog.md)
 - [ServersConfigurator](docs/ServersConfigurator.md)
 - [ServersConfiguratorRequirements](docs/ServersConfiguratorRequirements.md)
 - [ServersOs](docs/ServersOs.md)
 - [ServersOsRequirements](docs/ServersOsRequirements.md)
 - [ServersPreset](docs/ServersPreset.md)
 - [ServersSoftware](docs/ServersSoftware.md)
 - [ServersSoftwareRequirements](docs/ServersSoftwareRequirements.md)
 - [ServersStatistics](docs/ServersStatistics.md)
 - [ServersStatisticsListInner](docs/ServersStatisticsListInner.md)
 - [ServersStatisticsMeta](docs/ServersStatisticsMeta.md)
 - [ServiceCostType](docs/ServiceCostType.md)
 - [ServicePrice](docs/ServicePrice.md)
 - [ServicePriceConfiguration](docs/ServicePriceConfiguration.md)
 - [ServicePriceType](docs/ServicePriceType.md)
 - [ServiceServicePrice](docs/ServiceServicePrice.md)
 - [ServiceServicePriceNodeGroupsInner](docs/ServiceServicePriceNodeGroupsInner.md)
 - [SetAgentSkill](docs/SetAgentSkill.md)
 - [SetAgentSkillForAgent200Response](docs/SetAgentSkillForAgent200Response.md)
 - [SetLabels](docs/SetLabels.md)
 - [SettingCondition](docs/SettingCondition.md)
 - [Skill](docs/Skill.md)
 - [SkillAgent](docs/SkillAgent.md)
 - [SkillCatalogItem](docs/SkillCatalogItem.md)
 - [SkillSaveMeta](docs/SkillSaveMeta.md)
 - [SkillScript](docs/SkillScript.md)
 - [SkillScriptInput](docs/SkillScriptInput.md)
 - [SkillVersion](docs/SkillVersion.md)
 - [SpamFilterIsDisabled](docs/SpamFilterIsDisabled.md)
 - [SpamFilterIsEnabled](docs/SpamFilterIsEnabled.md)
 - [SpamProtectionIsDisabled](docs/SpamProtectionIsDisabled.md)
 - [SpamProtectionIsEnabled](docs/SpamProtectionIsEnabled.md)
 - [SshKey](docs/SshKey.md)
 - [SshKeyUsedByInner](docs/SshKeyUsedByInner.md)
 - [StaticRouteIn](docs/StaticRouteIn.md)
 - [StaticRouteOut](docs/StaticRouteOut.md)
 - [StaticRouteResponse](docs/StaticRouteResponse.md)
 - [StaticRoutesResponse](docs/StaticRoutesResponse.md)
 - [Statistics](docs/Statistics.md)
 - [StatisticsRequestsInner](docs/StatisticsRequestsInner.md)
 - [StatisticsTrafficInner](docs/StatisticsTrafficInner.md)
 - [Status](docs/Status.md)
 - [StatusCompanyInfo](docs/StatusCompanyInfo.md)
 - [StopAiSolutionChat](docs/StopAiSolutionChat.md)
 - [StreamError](docs/StreamError.md)
 - [Subdomain](docs/Subdomain.md)
 - [SyncAgentEntityLimits](docs/SyncAgentEntityLimits.md)
 - [SyncAgentEntityLimitsLimitsInner](docs/SyncAgentEntityLimitsLimitsInner.md)
 - [SyncAgentSkillAgentsVersions200Response](docs/SyncAgentSkillAgentsVersions200Response.md)
 - [SyncAgentSkillAgentsVersions200ResponseMeta](docs/SyncAgentSkillAgentsVersions200ResponseMeta.md)
 - [SyncAgentSkillVersionsForAgent200Response](docs/SyncAgentSkillVersionsForAgent200Response.md)
 - [SyncAiKeyLimits](docs/SyncAiKeyLimits.md)
 - [SyncAiKeyLimitsLimitsInner](docs/SyncAiKeyLimitsLimitsInner.md)
 - [Tags](docs/Tags.md)
 - [TextContent](docs/TextContent.md)
 - [TokenPackage](docs/TokenPackage.md)
 - [Tool](docs/Tool.md)
 - [Tool2](docs/Tool2.md)
 - [Tool2Annotations](docs/Tool2Annotations.md)
 - [ToolAuthData](docs/ToolAuthData.md)
 - [ToolAuthDataParams](docs/ToolAuthDataParams.md)
 - [ToolConfigField](docs/ToolConfigField.md)
 - [ToolMethod](docs/ToolMethod.md)
 - [ToolPrice](docs/ToolPrice.md)
 - [ToolTemplate](docs/ToolTemplate.md)
 - [ToolTemplateConfigSchema](docs/ToolTemplateConfigSchema.md)
 - [ToolsCallParams](docs/ToolsCallParams.md)
 - [ToolsCallResult](docs/ToolsCallResult.md)
 - [ToolsListResult](docs/ToolsListResult.md)
 - [TopLevelDomain](docs/TopLevelDomain.md)
 - [TopLevelDomainAllowedBuyPeriodsInner](docs/TopLevelDomainAllowedBuyPeriodsInner.md)
 - [TrafficUsage](docs/TrafficUsage.md)
 - [TransferStatus](docs/TransferStatus.md)
 - [TransferStatusErrorsInner](docs/TransferStatusErrorsInner.md)
 - [TransferStorageRequest](docs/TransferStorageRequest.md)
 - [URLType](docs/URLType.md)
 - [UpdateAdmin](docs/UpdateAdmin.md)
 - [UpdateAgent](docs/UpdateAgent.md)
 - [UpdateAgentEntityLimit](docs/UpdateAgentEntityLimit.md)
 - [UpdateAgentTool](docs/UpdateAgentTool.md)
 - [UpdateAiKey](docs/UpdateAiKey.md)
 - [UpdateAiKey200Response](docs/UpdateAiKey200Response.md)
 - [UpdateAiKeyLimit](docs/UpdateAiKeyLimit.md)
 - [UpdateAiSolution](docs/UpdateAiSolution.md)
 - [UpdateAppSettings200Response](docs/UpdateAppSettings200Response.md)
 - [UpdateAuthRestrictionsByCountriesRequest](docs/UpdateAuthRestrictionsByCountriesRequest.md)
 - [UpdateAutoBackup](docs/UpdateAutoBackup.md)
 - [UpdateBalancer](docs/UpdateBalancer.md)
 - [UpdateBot](docs/UpdateBot.md)
 - [UpdateCluster](docs/UpdateCluster.md)
 - [UpdateClusterConfiguration](docs/UpdateClusterConfiguration.md)
 - [UpdateClusterV2](docs/UpdateClusterV2.md)
 - [UpdateClusterV2DiskAutoscaling](docs/UpdateClusterV2DiskAutoscaling.md)
 - [UpdateClusterV2MaintenanceSlot](docs/UpdateClusterV2MaintenanceSlot.md)
 - [UpdateDatabaseCluster200Response](docs/UpdateDatabaseCluster200Response.md)
 - [UpdateDatabaseInstance409Response](docs/UpdateDatabaseInstance409Response.md)
 - [UpdateDedicatedServerRequest](docs/UpdateDedicatedServerRequest.md)
 - [UpdateDocument](docs/UpdateDocument.md)
 - [UpdateDomain](docs/UpdateDomain.md)
 - [UpdateDomainAutoProlongation200Response](docs/UpdateDomainAutoProlongation200Response.md)
 - [UpdateDomainMailInfoRequest](docs/UpdateDomainMailInfoRequest.md)
 - [UpdateDomainNameServers](docs/UpdateDomainNameServers.md)
 - [UpdateDomainNameServersNameServersInner](docs/UpdateDomainNameServersNameServersInner.md)
 - [UpdateFloatingIp](docs/UpdateFloatingIp.md)
 - [UpdateHttpResource](docs/UpdateHttpResource.md)
 - [UpdateInstance](docs/UpdateInstance.md)
 - [UpdateKafkaConfigParameters](docs/UpdateKafkaConfigParameters.md)
 - [UpdateKeyRequest](docs/UpdateKeyRequest.md)
 - [UpdateKnowledgebase](docs/UpdateKnowledgebase.md)
 - [UpdateMailbox](docs/UpdateMailbox.md)
 - [UpdateMailboxV2](docs/UpdateMailboxV2.md)
 - [UpdateMailboxV2200Response](docs/UpdateMailboxV2200Response.md)
 - [UpdateMaxBot](docs/UpdateMaxBot.md)
 - [UpdateNetworkDrive](docs/UpdateNetworkDrive.md)
 - [UpdateNotificationSettingsRequest](docs/UpdateNotificationSettingsRequest.md)
 - [UpdateNotificationSettingsRequestSettingsInner](docs/UpdateNotificationSettingsRequestSettingsInner.md)
 - [UpdateNotificationSettingsRequestSettingsInnerChannels](docs/UpdateNotificationSettingsRequestSettingsInnerChannels.md)
 - [UpdatePerson](docs/UpdatePerson.md)
 - [UpdateProject](docs/UpdateProject.md)
 - [UpdateRule](docs/UpdateRule.md)
 - [UpdateS3Backup](docs/UpdateS3Backup.md)
 - [UpdateServer](docs/UpdateServer.md)
 - [UpdateServerConfigurator](docs/UpdateServerConfigurator.md)
 - [UpdateServerDiskBackupRequest](docs/UpdateServerDiskBackupRequest.md)
 - [UpdateServerDiskRequest](docs/UpdateServerDiskRequest.md)
 - [UpdateServerIPRequest](docs/UpdateServerIPRequest.md)
 - [UpdateServerNATRequest](docs/UpdateServerNATRequest.md)
 - [UpdateServerOSBootModeRequest](docs/UpdateServerOSBootModeRequest.md)
 - [UpdateSettings](docs/UpdateSettings.md)
 - [UpdateSkill](docs/UpdateSkill.md)
 - [UpdateStorageRequest](docs/UpdateStorageRequest.md)
 - [UpdateStorageRequestConfigurator](docs/UpdateStorageRequestConfigurator.md)
 - [UpdateStorageUser200Response](docs/UpdateStorageUser200Response.md)
 - [UpdateStorageUserRequest](docs/UpdateStorageUserRequest.md)
 - [UpdateTool](docs/UpdateTool.md)
 - [UpdateVpc](docs/UpdateVpc.md)
 - [UploadResult](docs/UploadResult.md)
 - [UploadSuccessful](docs/UploadSuccessful.md)
 - [UploadSuccessfulResponse](docs/UploadSuccessfulResponse.md)
 - [UrlStatus](docs/UrlStatus.md)
 - [Use](docs/Use.md)
 - [UserNode](docs/UserNode.md)
 - [Valkey](docs/Valkey.md)
 - [Vds](docs/Vds.md)
 - [VdsDisksInner](docs/VdsDisksInner.md)
 - [VdsImage](docs/VdsImage.md)
 - [VdsNetworksInner](docs/VdsNetworksInner.md)
 - [VdsNetworksInnerIpsInner](docs/VdsNetworksInnerIpsInner.md)
 - [VdsOs](docs/VdsOs.md)
 - [VdsSoftware](docs/VdsSoftware.md)
 - [Vpc](docs/Vpc.md)
 - [VpcPort](docs/VpcPort.md)
 - [VpcPortService](docs/VpcPortService.md)
 - [VpcService](docs/VpcService.md)
 - [WorkerPresetOutApi](docs/WorkerPresetOutApi.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### Bearer

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```

### AiGatewayKey

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

info@timeweb.cloud

