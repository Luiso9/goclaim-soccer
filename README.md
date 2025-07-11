# GoClaim Soccer
Automatically claim SoccerGuru's Card.

![image](https://github.com/user-attachments/assets/72dc7970-7db7-4d4e-8169-50f4c6a64cf9)


---
## Build Steps

### Prerequisites

- Go (version 1.24 or higher)
- Git
- Discord (with Rich Presence enabled)

### Steps
 1. **Clone the Repository**
```bash
git clone https://github.com/Luiso9/goclaim-soccer.git
cd goclaim-soccer
```
 2. **Initialize Go module**
```bash
go mod init github.com/Luiso9/goclaim-soccer
```

 3. **Download Dependencies**
```bash
go mod tidy
```

 4. **Create a .env file with value**
```bash
echo 'AUTH_KEY=your_token_value' > .env
```

 5. **Change Discord Webhook URL in `internal/webhook/webhook.go`**
```go
var webhookUrl = "ur webhook url"
```

 6. **Build**
```bash
go build
```

 7. **Run**
```bash
./goclaim-soccer
```

---

## How To Get AUTH_KEY

1. **Head up to https://soccerguru.live/dashboard and enable Developer Tools (F12)**
2. **Goto "Network" tab and do something like Claim Daily or Watch Video (AD)**
3. **Find the POST Request and check the Request Header. You will found authorization with bunch of random strings**
![image](https://github.com/user-attachments/assets/d2829c1d-1ef9-487b-9bd1-10ee45b91210)
4. **Copy that and paste it into your .env file as AUTH_KEY=yourkey**

---

## Known Issue

1. Unreliable JSON Parsing from API Response
Occasionally, the response from `https://api.soccerguru.live/claim` returns unexpected or malformed data, causing a JSON parsing error `(invalid character 'a' looking for beginning of value)`.
Well, we still able to claim the card tho :D
