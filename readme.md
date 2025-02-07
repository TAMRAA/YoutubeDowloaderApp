# YouTube Downloader

This project is a simple YouTube video downloader using Go and Python.

## Prerequisites

Make sure you have the following installed:

- [Go](https://go.dev/dl/)
- [Python](https://www.python.org/downloads/)
- `pip` (Python package manager)

## Installation

Follow these steps to set up the project:

### 1. Initialize Go Module

```sh
go mod init youtube-downloader
```

### 2. Install Dependencies

```sh
go get -u github.com/kkdai/youtube/v2
```

```sh
pip install yt-dlp
```

## Usage

Run the Go program with the YouTube video URL:

```sh
go run main.go <YouTube Video URL>
```

Replace `<YouTube Video URL>` with the actual link of the YouTube video you want to download.

## License

This project is open-source and available under the [MIT License](LICENSE).

@buildwithXStudio
