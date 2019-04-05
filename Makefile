all:
	CGO_ENABLED=0 go build -a -installsuffix cgo -o login .

prod:
	GOOS=linux CGO_ENABLED=0 go build -a -installsuffix cgo -o login .
	- docker image rm reg.urantiatech.com/auth/login 
	docker build -t reg.urantiatech.com/auth/login .
	docker push reg.urantiatech.com/auth/login

dev:
	GOOS=linux CGO_ENABLED=0 go build -a -installsuffix cgo -o login .
	- docker image rm localhost:5000/auth/login 
	docker build -t localhost:5000/auth/login .
	docker push localhost:5000/auth/login
