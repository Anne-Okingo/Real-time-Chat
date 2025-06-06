# Real-Time WebSocket Chat Application

![Chat Application Demo](screenshot.png)

A simple real-time chat application built with Go (Golang) backend and modern HTML5/JavaScript frontend.Exploring the possibilities of golang websocket standard library.

## Features

- **Real-time messaging** using WebSocket protocol
- **Multiple concurrent users** with message broadcasting
- **Clean, responsive interface** works on desktop and mobile
- **Connection status indicators** (online/offline)
- **Message history** in the chat window
- **Modern UI** with visual message differentiation

## Tech Stack

- **Backend**: 
  - Go (Golang)
  - `golang.org/x/net/websocket` package
- **Frontend**:
  - Vanilla JavaScript
  - Modern CSS3 with Flexbox
  - WebSocket API

## Getting Started

### Prerequisites

- Go 1.16+ installed
- Modern web browser (Chrome, Firefox, Edge, Safari)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/websocket-chat.git
   cd websocket-chat
   ```
   Run the server:
bash

    go run server.go

    Open the chat interface:

        Open index.html in your browser

        Or visit http://localhost:8080 if serving via Go

📖 Usage

    Open the chat interface in your browser

    The system will automatically connect to the chat server

    Type messages in the input box at the bottom

    Press Enter or click Send to broadcast your message

    All connected users will see messages in real-time

🏗️ Project Structure

websocket-chat/
├── server.go         # WebSocket server implementation
├── index.html        # Chat client interface
├── README.md         # Documentation file
└── screenshot.png    # Application screenshot

🌈 Customization
Server Configuration

    Change port by editing :8080 in server.go

    Adjust message buffer size (default: 1024 bytes)

Client Customization

    Modify colors in CSS variables:
    css

    :root {
      --primary-color: #4361ee;
      --secondary-color: #3a0ca3;
    }

    Add features like:

        Username support

        Message timestamps

        Emoji picker

🐛 Troubleshooting
Issue	Solution
Can't connect to server	Ensure server is running (go run server.go)
Messages not appearing	Refresh the page to reconnect
Connection drops	Check server logs for errors
📜 License

MIT License - see LICENSE for details
🤝 Contributing

Pull requests are welcome! For major changes, please open an issue first to discuss what you'd like to change.