gcpの学習用リポジトリ

# セットアップ
1. App Engin extension for GOのインストール  
```
gcloud components install app-engine-go
```  
2. gcloud用のGoogleアカウント認証  
```
gcloud auth login
```
3. project 作成
```
mkdir gcpSample
cd gcpSample
```
4. app.yamlを作成  
5. main.goの作成

# ローカルでの動作確認
```
dev_appserver.py app.yaml
```

# サーバーにdeploy
```
gcloud app deploy --project gcpSample 
```

# サーバーの動作確認
```
gcloud app browse --project gcpSample 
```
