APP_NAME=food-delivery
DEPLOY_CONNECT=root@g05.custohub.com


echo "Downloading packages ..."
go mod download
echo "Compiling..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app

echo "Docker building..."
docker build -t ${APP_NAME} -f ./Dockerfile .
echo "Docker saving..."
docker save -o ${APP_NAME}.tar ${APP_NAME}

echo "Deploying..."
scp -o StrictHostKeyChecking=no ${APP_NAME}.tar  ${DEPLOY_CONNECT}:~
scp -o StrictHostKeyChecking=no ${DEPLOY_CONNECT} 'bash -s' < ./deploy/stg.sh

echo "Cleaning..."
rm -f ./${APP_NAME}.tar
#docker rmi $(docker images -qa -f 'dangling=true')
echo "Done"