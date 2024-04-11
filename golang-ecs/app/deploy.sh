#!/bin/bash

# Dockerイメージをビルドする
docker build -t go-ecs-app --platform linux/x86_64 .

# ECRにログインする
aws ecr get-login-password --region ap-northeast-1 | docker login --username AWS --password-stdin 879853972315.dkr.ecr.ap-northeast-1.amazonaws.com

# タグを付ける
docker tag go-ecs-app:latest 879853972315.dkr.ecr.ap-northeast-1.amazonaws.com/go-ecs-app:latest

# ECRにプッシュする
docker push 879853972315.dkr.ecr.ap-northeast-1.amazonaws.com/go-ecs-app:latest

# ECSにデプロイする
aws ecs update-service --cluster golang-ecs-cluster --service golang-ecs-app-service --force-new-deployment