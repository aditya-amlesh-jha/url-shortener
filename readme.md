# URL Shortener

This is a simple URL shortener service built with Go. It allows users to shorten long URLs and provides a redirect service to access the original URLs.

**Features:**

* Shortens long URLs using a unique generated short code.
* Redirects short URLs to their corresponding long URLs.
* Uses Redis for caching shortened URLs for faster retrieval.
* Persistent storage of short URL and long URL mapping in a database.

**Technologies Used:**

* Go
* Redis
* MySQL (or any database with a Go driver)
* http (web framework - optional)


**Folder Structure**

server/
├── main.go
├── handler/
│   ├── url.go
├── model/
│   └── url.go
│   └── mysql.go
│   └── redis.go
├── utils/
│   └── shortener.go
├── config/
│   └── config.go
└── test/
    └── url_test.go

**Getting Started:**

1. **Clone the repository:**

git clone https://github.com/aditya-amlesh-jha/url-shortener.git

2. **Install dependencies:**

cd url-shortener
go get

3. **Run the server:**
docker-compose up --build

Download httpie cli or any api-testing tool

**API Endpoints:**
* **POST /shorten:** Shortens a long URL.
    * Request body: `{"long_url": "your_long_url_here"}`
    * Response: `{"short_url": "shortened_url"}`
* **GET /{short_url}:** Redirects to the original long URL.
