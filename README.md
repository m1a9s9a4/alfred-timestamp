# Alfred Timestamp Converter

Alfred workflow for converting between various datetime and timestamp formats.

## Download

**[Download alfred-timestamp.alfredworkflow](https://github.com/masafumikato/alfred-timestamp/raw/main/releases/alfred-timestamp.alfredworkflow)**

Click the link above to download the workflow file, then double-click it to install in Alfred.

## Features

- Convert current time or specified datetime to multiple timestamp formats
- Support for Unix timestamps, ISO 8601, and various date formats
- Quick copy to clipboard functionality

## Installation

### Option 1: Download (Recommended)
1. Click the download link above
2. Double-click the downloaded `alfred-timestamp.alfredworkflow` file
3. Alfred will automatically import the workflow

### Option 2: Build from Source
1. Clone this repository
2. Double-click the `alfred-timestamp.alfredworkflow` file to install
3. Or manually add the workflow to Alfred

## Usage

Type `ts` in Alfred followed by:

- `now` - Convert current time
- Unix timestamp (e.g., `1705123456`)
- Date formats:
  - `YYYYMMDD` (e.g., `20240101`)
  - `YYYY-MM-DD` (e.g., `2024-01-01`)
  - `YYYYMMDDHHmmss` (e.g., `20240101150430`)
  - ISO 8601 format (e.g., `2024-01-13T14:24:16+09:00`)
  - Common datetime formats (e.g., `2024-01-13 14:24:16`)

## Output Formats

The workflow displays multiple format options:

1. **YYYYMMDDHHmmss** - Compact timestamp format
2. **Unix timestamp** - Seconds since epoch
3. **ISO 8601** - International standard format
4. **Human readable** - YYYY-MM-DD HH:mm:ss
5. **Date only** - YYYY-MM-DD
6. **RFC822** - Email date format

Select any format to copy it to your clipboard.

## Examples

```
ts now
→ Shows current time in all formats

ts 1705123456
→ Converts Unix timestamp to all formats

ts 20240101
→ Converts date (assumes time 01:01:01) to all formats

ts 2024-01-15
→ Converts date with dashes to all formats
```

## Development

Built with Go. To build from source:

```bash
go build -o main main.go
```

## Requirements

- Alfred 4 or 5
- macOS

## License

MIT