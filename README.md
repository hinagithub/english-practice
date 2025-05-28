# English Practice API 🇺🇸→🇯🇵

英語フレーズとその日本語訳をランダムで返す API。Cloud Run にデプロイしています。

---

## 🌐 URL

[Cloud Run 本番環境](https://english-practice-1064724819814.us-central1.run.app/)

例：

- トップページ:  
  `https://english-practice-1064724819814.us-central1.run.app/`

- ランダム英語フレーズ:  
  `https://english-practice-1064724819814.us-central1.run.app/random`

---

## 🛠 セットアップ手順

### Google Cloud SDK のインストール

```sh
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


## 🏃‍♀️ ローカル実行
```
go run main.go
```

```
curl http://localhost:8080
```

## ✨ デプロイ

```
make deploy
```

