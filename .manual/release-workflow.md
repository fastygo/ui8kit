Отличный вопрос — для Go-модуля, который раздаётся через `go get`, версионирование критически важно, потому что Go proxy кеширует теги навсегда.

## Как работает версионирование Go-модулей

Go module system привязан к **git tags**. Когда пользователь делает:

```bash
go get github.com/fastygo/ui8kit@latest
```

Go proxy (`proxy.golang.org`) ищет **semver-тег** (`v0.1.0`, `v1.2.3`) в репозитории. Без тега — fallback на pseudo-version (`v0.0.0-20260325...`), что выглядит непрофессионально и нестабильно.

## Стратегия версионирования

### 1. Semantic Versioning (semver)

```
v{MAJOR}.{MINOR}.{PATCH}
```

| Изменение | Когда инкрементировать |
|-----------|----------------------|
| PATCH `v0.1.1` | Баг-фиксы, не меняющие API |
| MINOR `v0.2.0` | Новый компонент, новые поля в Props (обратно совместимо) |
| MAJOR `v1.0.0` | Ломающие изменения: переименование Props, удаление компонентов |

### 2. Пока `v0.x.x` — свобода

Пока major = 0, Go считает API нестабильным. Можно ломать обратную совместимость без смены module path. Это ваш текущий этап (`v0.1.0`).

### 3. После `v1.0.0` — жёсткие правила

Ломающее изменение после `v1` **требует** смены module path:

```go
// v1
module github.com/fastygo/ui8kit

// v2 (если когда-нибудь потребуется)
module github.com/fastygo/ui8kit/v2
```

Поэтому не спешите с `v1.0.0` — выпускайте его только когда API стабилизирован.

## Практическая настройка

### Шаг 1: Константа версии в коде

Уже есть в `ui8kit.go`:

```1:13:ui8kit.go
// Package ui8kit provides a component kit for Go + templ + Tailwind CSS
// in the style of shadcn/ui.
//
// Import subpackages directly:
//
//	import "github.com/fastygo/ui8kit/ui"
//	import "github.com/fastygo/ui8kit/layout"
//	import "github.com/fastygo/ui8kit/utils"
//	import "github.com/fastygo/ui8kit/styles"
package ui8kit

const Version = "0.1.0"
```

Держите `Version` в синхронизации с git-тегами (это можно автоматизировать — см. ниже).

### Шаг 2: Git tags — формат и создание

```bash
# Первый релиз
git tag v0.1.0
git push origin v0.1.0

# Баг-фикс
git tag v0.1.1
git push origin v0.1.1

# Новый компонент (Card, Alert, etc.)
git tag v0.2.0
git push origin v0.2.0
```

Важные правила:

- **Всегда** префикс `v` — Go требует это.
- **Никогда** не удаляйте и не перезаписывайте опубликованный тег — proxy уже закешировал его.
- **Lightweight tags** достаточно, но **annotated tags** предпочтительнее (попадают в GitHub Releases):

```bash
git tag -a v0.1.0 -m "v0.1.0: initial release"
```

### Шаг 3: GitHub Release workflow

Ваш `release.yml` уже настроен — при пуше тега `v*` создаётся GitHub Release с `generate_release_notes: true`. Workflow также запускает `go test` перед релизом.

### Шаг 4: Go proxy

После первого `git push origin v0.1.0` модуль автоматически появится на `pkg.go.dev` в течение нескольких минут. Для принудительного обновления:

```bash
GOPROXY=proxy.golang.org go list -m github.com/fastygo/ui8kit@v0.1.0
```

## Рекомендуемый release workflow

```
1. Обновить Version в ui8kit.go
2. go test ./...
3. git add -A && git commit -m "chore: release v0.2.0"
4. git tag -a v0.2.0 -m "v0.2.0"
5. git push origin main --tags
6. → CI прогоняет тесты
7. → release.yml создаёт GitHub Release
8. → proxy.golang.org индексирует тег
9. → Потребители: go get github.com/fastygo/ui8kit@v0.2.0
```

## Что нельзя делать после публикации тега

| Действие | Почему опасно |
|----------|---------------|
| `git tag -d v0.1.0 && git push --delete origin v0.1.0` | Proxy уже закешировал; `go get` у пользователей сломается |
| `git push --force` на тегированный коммит | Checksum (`go.sum`) у пользователей не совпадёт — build fails |
| Переименование пакетов без нового MINOR | Сломает `import` у всех, кто уже использует |

Если в теге баг — **не удаляйте тег**, выпустите PATCH:

```bash
# v0.1.0 содержит баг → выпускаем v0.1.1
git tag -a v0.1.1 -m "v0.1.1: fix ..."
git push origin v0.1.1
```

## Retract (Go 1.16+)

Если тег выпущен по ошибке, можно пометить его как **retracted** — `go get` будет игнорировать эту версию:

```go
// go.mod
module github.com/fastygo/ui8kit

go 1.23.0

retract v0.1.0 // accidentally published with broken imports
```

Выпускаете новый тег с этим `retract` — и proxy перестаёт рекомендовать старую версию.

## Итого: чеклист

1. **Semver теги** (`v0.x.y`) — единственный источник правды для Go proxy.
2. **`Version` константа** в `ui8kit.go` — для runtime-проверок (держать в sync с тегами).
3. **Annotated tags** + **GitHub Releases** — changelog через `generate_release_notes`.
4. **Никогда не удалять/перезаписывать** опубликованные теги.
5. **`retract`** — если нужно "отозвать" плохой релиз.
6. **Оставаться на `v0.x`** до стабилизации API — это даёт свободу ломать без `/v2`.

Если хотите, чтобы я настроил автоматическую синхронизацию `Version` с git tag (через `go generate` или pre-tag script) — переключайтесь в Agent mode.

=================

Все 5 пакетов, 39 тестов — зелёные.

## Что добавлено

### `scripts/release.sh` — единая точка входа для релизов

```bash
./scripts/release.sh 0.2.0
```

Скрипт проверяет:
- Формат semver (без `v` — скрипт добавит сам)
- Тег ещё не существует
- Working tree чистый
- Ветка `main` (иначе warning с подтверждением)
- Тесты проходят (`go test -race`)

Затем атомарно: обновляет `Version` в `ui8kit.go`, коммитит, создаёт **annotated tag**. Остаётся только `git push origin main --tags`.

### `release.yml` — усиленный CI pipeline

Добавлены три защиты:
1. **Version consistency check** — если git tag `v0.2.0`, а в `ui8kit.go` написано `0.1.0` — pipeline падает с ошибкой
2. **`go vet`** — статический анализ перед релизом
3. **Go proxy indexing** — принудительно пингует `proxy.golang.org` чтобы `go get` работал сразу

### `ui8kit_test.go` — runtime-защита

Два теста: `TestVersionFormat` (валидный semver X.Y.Z) и `TestVersionNotEmpty`. Если кто-то случайно сломает константу — тесты упадут до того, как изменение попадёт в тег.

### `docs/versioning.md` — документация для контрибьюторов и пользователей

Полное описание: как создать релиз, как подключить конкретную версию, что нельзя делать с тегами, как использовать `retract`.