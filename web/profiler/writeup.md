
まず、アプリケーションの通信を眺めるとGraphQLを用いていることがわかる。

そこで、GraphQLのエンドポイント(`/api`)に対して以下のようなイントロスペクションクエリを送信しスキーマを確認する。受信したスキーマは、[GraphQL Voyager](https://apis.guru/graphql-voyager/)等で可視化できる。

```
query IntrospectionQuery {
  __schema {
    queryType {
      name
    }
    mutationType {
      name
    }
    subscriptionType {
      name
    }
    types {
      ...FullType
    }
    directives {
      name
      description
      locations
      args {
        ...InputValue
      }
    }
  }
}

fragment FullType on __Type {
  kind
  name
  description
  fields(includeDeprecated: true) {
    name
    description
    args {
      ...InputValue
    }
    type {
      ...TypeRef
    }
    isDeprecated
    deprecationReason
  }
  inputFields {
    ...InputValue
  }
  interfaces {
    ...TypeRef
  }
  enumValues(includeDeprecated: true) {
    name
    description
    isDeprecated
    deprecationReason
  }
  possibleTypes {
    ...TypeRef
  }
}

fragment InputValue on __InputValue {
  name
  description
  type {
    ...TypeRef
  }
  defaultValue
}

fragment TypeRef on __Type {
  kind
  name
  ofType {
    kind
    name
    ofType {
      kind
      name
      ofType {
        kind
        name
        ofType {
          kind
          name
          ofType {
            kind
            name
            ofType {
              kind
              name
              ofType {
                kind
                name
              }
            }
          }
        }
      }
    }
  }
}
```

上記から`someone(uid: String!)`クエリを用いて特定のユーザの情報を取得できることが推察される。また、Flagページの内容からadminの`token`を取得することが目的であるとわかるため、`uid`には`admin`を指定する。

上記からadminの`token`がわかるため、`updateToken(token: "adminのtoken")`を送信し、自身の`token`をadminのものに置き換える。

最後にFlagのページにアクセスすればFlagを取得できる。
