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
    content=event["content"]


    dynamoDB = Aws::DynamoDB::Resource.new(region: 'ap-northeast-1')
    table = dynamoDB.table('Data')

    #データを入力

    #追加する案件
    data={ posted: Time.new.strftime("%m/%d-%H:%M"), st: st, en: en, tag: tag, content: content }

    #ユーザが登録がまだされていなかったら
    if table.get_item({key: { 'name' => name } })["item"].nil?
      json={
        item: {
          name: name,
          plan: [data]
        }
      }
      table.put_item(json)
    else
      json={
        item: {
          name: name,
          plan: table.get_item({key: { 'name' => name } })["item"]["plan"].push(data)
        }
      }
      table.put_item(json)
    end

    #データを取得
    resp = table.get_item({
        key: { 'name' => name }
    })

    { statusCode: 200, body: resp }
end
