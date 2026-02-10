# Go Email Dispatcher

A high-performance, concurrent email dispatcher written in Go that uses the producer-consumer pattern to send personalized emails to multiple recipients efficiently.

## 📋 Overview

This project demonstrates the power of Go's concurrency features by implementing a robust email dispatcher system. It reads recipient data from a CSV file, processes emails concurrently using multiple worker goroutines, and sends personalized emails using customizable templates.

## ✨ Features

- **Concurrent Processing**: Utilizes Go's goroutines and channels for efficient parallel email processing
- **Producer-Consumer Pattern**: Implements a clean separation between data loading and email sending
- **Template-Based Emails**: Uses Go's `html/template` package for personalized email content
- **CSV Data Source**: Easily manage recipients through a simple CSV file
- **Worker Pool**: Configurable number of concurrent email workers
- **SMTP Support**: Compatible with any SMTP server (includes Mailpit setup for local testing)

## 🚀 Prerequisites

- Go 1.22.2 or higher
- Docker (optional, for running Mailpit locally)

## 📦 Installation

1. Clone the repository:
```bash
git clone https://github.com/Rudraksh121a/Go-Email-Dispatcher-.git
cd Go-Email-Dispatcher-
```

2. Install dependencies:
```bash
go mod download
```

3. (Optional) Set up Mailpit for local email testing:
```bash
docker run -d \
  --restart unless-stopped \
  --name=mailpit \
  -p 8025:8025 \
  -p 1025:1025 \
  axllent/mailpit
```

Access the Mailpit web interface at `http://localhost:8025` to view sent emails.

## 🎯 Usage

1. **Prepare your recipient data**: Edit `email-data.csv` with your recipients:
```csv
Name,Email
John Doe,john@example.com
Jane Smith,jane@example.com
```

2. **Customize the email template**: Modify `email.tmpl` to your needs:
```
To: {{.Email}}
Subject: Hello {{.Name}}!

Hi {{.Name}},

Your personalized message here...

Thanks,
Your Team
```

3. **Run the dispatcher**:
```bash
go run .
```

## ⚙️ Configuration

### Worker Count
Adjust the number of concurrent workers in `main.go`:
```go
workerCount := 5  // Change this value
```

### SMTP Settings
Configure SMTP settings in `consumer.go`:
```go
smtpHost := "LocalHost"  // Your SMTP host
smtpPort := "1025"        // Your SMTP port
```

## 📁 Project Structure

```
.
├── main.go           # Entry point, template execution
├── producer.go       # CSV reading and recipient loading
├── consumer.go       # Email worker implementation
├── email.tmpl        # Email template
├── email-data.csv    # Recipient data
├── go.mod            # Go module definition
├── info.md           # Docker setup instructions
└── README.md         # This file
```

## 🏗️ How It Works

1. **Producer**: The `loadRecipient` function reads the CSV file and sends recipient data through an unbuffered channel
2. **Channel**: Acts as a queue connecting producers and consumers
3. **Consumer Pool**: Multiple `emailWorker` goroutines process recipients concurrently
4. **Template Rendering**: Each worker renders the email template with recipient data
5. **Email Sending**: Workers send emails via SMTP
6. **Synchronization**: Uses `sync.WaitGroup` to ensure all emails are sent before program exits

### Architecture Diagram

```
┌─────────────┐
│ CSV File    │
└──────┬──────┘
       │
       ▼
┌─────────────┐      ┌──────────────┐
│  Producer   │─────▶│   Channel    │
│ (goroutine) │      │  (unbuffered)│
└─────────────┘      └──────┬───────┘
                            │
                ┌───────────┴───────────┐
                │                       │
         ┌──────▼──────┐        ┌──────▼──────┐
         │   Worker 1  │  ...   │   Worker N  │
         │ (goroutine) │        │ (goroutine) │
         └──────┬──────┘        └──────┬──────┘
                │                       │
                └───────────┬───────────┘
                            ▼
                    ┌──────────────┐
                    │ SMTP Server  │
                    └──────────────┘
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

## 👤 Author

**Rudraksh121a**

- GitHub: [@Rudraksh121a](https://github.com/Rudraksh121a)

## 🙏 Acknowledgments

- Thanks to the Go community for excellent concurrency primitives
- Mailpit for providing an easy-to-use email testing solution
