# File Sorter

Утилита для автоматической сортировки файлов в указанной директории по категориям.

## Описание

File Sorter - это программа на Go, которая автоматически сортирует файлы в указанной директории по категориям, создавая соответствующие папки и перемещая файлы в зависимости от их расширений.

## Возможности

- Параллельная обработка файлов для повышения производительности
- Автоматическое создание категорий папок
- Сортировка файлов по следующим категориям:
  - Executable (исполняемые файлы)
  - Images (изображения)
  - Documents (документы)
  - Audio (аудио файлы)
  - Video (видео файлы)
  - Archives (архивы)
  - Code (исходный код)
  - Others (прочие файлы)
- Поддержка широкого спектра форматов файлов
- Автоматическое создание подпапок по расширениям файлов
- Логирование всех операций

## Установка

1. Клонируйте репозиторий:
```bash
git clone https://github.com/vyantik/file-sorter.git
```

2. Перейдите в директорию проекта:
```bash
cd file-sorter
```

3. Соберите проект:
```bash
go build
```

## Использование

1. Измените путь к директории в файле `main.go`:
```go
folderPath := "путь/к/вашей/директории"
```

2. Запустите программу:
```bash
./file-sorter
```

## Поддерживаемые форматы файлов

### Executable
- .exe, .msi, .dmg, .app, .deb, .rpm, .apk

### Images
- .jpg, .jpeg, .png, .gif, .bmp, .webp, .tiff, .ico, .svg, .psd, .ai

### Documents
- .pdf, .doc, .docx, .txt, .rtf, .xls, .xlsx, .ppt, .pptx, .odt, .ods, .odp

### Audio
- .mp3, .wav, .m4a, .ogg, .flac, .aac, .wma, .m4b

### Video
- .mp4, .avi, .mkv, .mov, .wmv, .flv, .mpeg, .mpg, .m4v, .webm

### Archives
- .zip, .rar, .7z, .tar, .gz, .bz2, .iso

### Code
- .py, .js, .ts, .go, .java, .cpp, .h, .c, .css, .html, .php, .sql, .json, .xml, .yaml, .yml, .toml, .conf, .env, .log, .md

## Лицензия

MIT

## Вклад в проект

Мы приветствуем ваши предложения и улучшения! Пожалуйста, создавайте pull request'ы для внесения изменений.
