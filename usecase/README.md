- アプリケーションレイヤでシステム使用上のユースケースを表現
  - ユーザ登録、ユーザ覧表示など
- handlerから呼びされる
  - 1つのhandlerに対応する専用usecaseが1つ存在します。
- 基本的にはdomainを触る
  - データの取得や保存などでDBにアクセスするときもdomain層のrepositoryを介してアクセスすることによって、infrastructure層ではなくdomain層のみに依存させています。