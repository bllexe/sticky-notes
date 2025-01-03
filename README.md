# Sticky Notes Application

A simple and efficient sticky notes application written in Go, designed for Ubuntu systems. This application allows you to manage your notes with different colors and provides basic note management functionality.

## Features

- Create, read, update, and delete sticky notes
- Support for 5 different note colors:
  - Yellow
  - Blue
  - Green
  - Pink
  - Orange
- Automatic timestamp tracking for creation and updates
- File-based storage system for persistence
- Search functionality to find specific notes
- Thread-safe operations for concurrent access

## Project Structure

The project follows a clean, layered architecture:

```
sticky-notes/
├── cmd/app/                    # Application entry points
├── internal/              # Internal packages
│   ├── model/            # Data models
│   ├── repository/       # Data storage layer
│   ├── service/          # Business logic layer
│   └── handler/          # User interface layer
├── pkg/                   # Reusable packages
├── data/                  # Storage directory for notes
└── README.md             # This file
```

## Prerequisites

- Go 1.21 or higher
- Ubuntu operating system

## Installation

1. Clone the repository:
```bash
git clone https://github.com/bllexe/sticky-notes.git
cd sticky-notes
```

2. Install dependencies:
```bash
go mod download
```

## Building

To build the application:

```bash
go build -o sticky-notes ./cmd/main.go
```

## Features in Detail

### Note Management
- Create new notes with custom content and color
- Update existing notes
- Delete unwanted notes
- View all notes or search for specific ones

### Data Persistence
- Notes are automatically saved to files
- Each note is stored as a separate JSON file
- Thread-safe operations for concurrent access


## todo 
-  make basic ui for sticky note 

## License
