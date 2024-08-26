Run the project by :
docker-compose up --build

Download httpie cli or any api-testing tool

http post http://localhost:8080/shorten long_url="https://www.google.com/search?q=httpie+linux&oq=httpie+linux&gs_lcrp=EgZjaHJvbWUyBggAEEUYOTIGCAEQRRg8MgYIAhBFGDwyBggDEEUYPNIBCDQzMDNqMGoxqAIAsAIA&sourceid=chrome&ie=UTF-8"
HTTP/1.1 200 OK
Content-Length: 0
Date: Mon, 26 Aug 2024 16:54:27 GMT

http post http://localhost:8080/redirect/6SRKsjKd                                                                                        ✔ 
HTTP/1.1 302 Found
Content-Length: 0
Date: Mon, 26 Aug 2024 16:55:47 GMT
Location: https://www.google.com/search?q=httpie+linux&oq=httpie+linux&gs_lcrp=EgZjaHJvbWUyBggAEEUYOTIGCAEQRRg8MgYIAhBFGDwyBggDEEUYPNIBCDQzMDNqMGoxqAIAsAIA&sourceid=chrome&ie=UTF-8


Scope for improvement ::

Can define type of http request for each path i.e --> shorten, redirect; currently it is not defined :)

