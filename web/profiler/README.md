# profiler
## 問題文
Let's edit your profile with profiler!

(Hint: You don't need to deobfuscate *.js.)

{{ URL }}

## 難易度
普通(GraphQLの知識があるなら易)。

## 概要
GraphQLに関する問題。

当該アプリケーションでは以下の機能が提供されている。
- アカウント作成機能
  - ID、パスワード、ユーザ名を入力する。
  - アカウントを作成するとランダムなトークンが通知される。
- ログイン機能
  - IDとパスワードを用いてログインする。
- プロフィール更新機能
  - 更新したいプロフィールの内容とアカウント作成時に通知されるトークンを送信し、プロフィールを更新する。
  - トークンが正しくないとプロフィールは更新されない。
- Flag取得機能
  - アカウントに設定されているトークンがadminのものであればFlagを取得できる。

当該アプリケーションでは、アカウント情報の取得やプロフィールの更新のためにGraphQLを用いている。
GraphQLのスキーマは以下の通りである。

```
type User {
    uid: ID!
    name: String!
    profile: String!
    token: String!
}

type Query {
    me: User!
    someone(uid: ID!): User
    flag: String!
}

type Mutation {
    updateProfile(profile: String!, token: String!): Boolean!
    updateToken(token: String!): Boolean!
}
```

上記のスキーマから分かる通り、当該GraphQLエンドポイントにおいては、本来必要ないはずの`updateToken`ミューテーションを実行できたり、各ユーザの`token`を`me`クエリや`someone`クエリで取得したりすることができる(これは、現実の開発において起こり得る不必要なクエリやパラメータの公開を想定している)。

加えて、当該GraphQLエンドポイントにおいてはイントロスペクションクエリの実行が許可されている。そのため、細工したイントロスペクションを送信することで、上記のスキーマを閲覧することが可能である(これは、現実の開発において起こり得る不必要なイントロスペクションクエリ実行の許可を想定している)。

よって、Flagを取得するためには以下のようにすればよい。
1. アカウント登録とログインを行い、プロフィールページに遷移する。ここで通信内容を確認すれば、GraphQLのAPIとやり取りしていることがわかる。
1. イントロスペクションクエリを送信しスキーマを確認する。
2. `someone(uid: "admin")`を送信し、adminの`token`を取得する。
3. `updateToken(token: "adminのtoken")`を送信し、自身の`token`をadminのものに置き換える。
4. Flag取得機能にアクセスする。

## 備考
配布ファイルはありません。

## 参考
[GraphQL — Common vulnerabilities & how to exploit them](https://medium.com/@the.bilal.rizwan/graphql-common-vulnerabilities-how-to-exploit-them-464f9fdce696)