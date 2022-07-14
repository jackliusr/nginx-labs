# dynamic ssi reactjs lab


A working example of dynamic ssi reactjs lab

## test 

```bash

docker compose up

# test existing

curl   --resolve 1208.com:8080:127.0.0.1 http://1208.com:8080/index.html
curl   --resolve www.1208.com:8080:127.0.0.1 http://www.1208.com:8080/index.html

# test not existing
curl   --resolve 1209.com:8080:127.0.0.1 http://1209.com:8080/index.html

```