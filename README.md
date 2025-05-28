## URL

```
https://english-practice-1064724819814.us-central1.run.app/
```


## セットアップ

Google Cloud SDKが必要

```
brew install --cask google-cloud-sdk
gcloud version

>
Google Cloud SDK 523.0.1
bq 2.1.17
core 2025.05.22
gcloud-crc32c 1.0.0
gsutil 5.34

gcloud init
```
→サインインや認証などする


## ローカル実行
```
go run main.go
```

```
curl http://localhost:8080
```

## デプロイ

```
gcloud run deploy english-practice \
  --source . \
  --region us-central1 \
  --project hinaproject-460723 \
  --platform managed \
  --allow-unauthenticated
```

