

IMG_NAME=postgres-test-img
APP_NAME=pg-test-app

build:
	docker image build . -t $(IMG_NAME)
	docker container run -d --rm -p 5432:5432 -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=ziyan --name $(APP_NAME) $(IMG_NAME)
