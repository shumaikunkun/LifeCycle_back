require 'json'
require 'aws-sdk'

def lambda_handler(event:, context:)

    #パラメタからユーザ名やタグを取得

    #文字列の場合、ハッシュに変換して取得
    # json=event.split("&").to_a.map{|a|a.split("=")}.to_h
    # name=json["name"]
    # tag=json["tag"]
    # st=json["st"]
    # en=json["en"]

    #JSONの場合
    name=event["name"]
    st=event["st"]
    en=event["en"]
    tag=event["tag"]


    dynamoDB = Aws::DynamoDB::Resource.new(region: 'ap-northeast-1')
    table = dynamoDB.table('Data')

    #データを入力
    table.put_item(
    {
      item: {
        name: name,
        plan:
          [
            { posted: "aaa", st: st, en: en, tag: tag },
            { posted:"2017", st:"2017", en:"2017", tag: "1" },
            { posted:"2017", st:"2017", en:"2017", tag: "1" },
          ]
      }
    }
    )

    #データを取得
    resp = table.get_item({
        key: { 'name' => name }
    })

    { statusCode: 200, body: resp }
end
