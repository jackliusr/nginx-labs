# dynamic ssi lab


A working example of dynamic ssi

## test 

```bash

docker compose up

# test existing

curl   --resolve 1203.com:8080:127.0.0.1 http://1203.com:8080/index.html
curl   --resolve 1303.com:8080:127.0.0.1 http://1303.com:8080/index.html
curl   --resolve 1403.com:8080:127.0.0.1 http://1403.com:8080/index.html

# test not existing
curl   --resolve 1209.com:8080:127.0.0.1 http://1209.com:8080/index.html

```