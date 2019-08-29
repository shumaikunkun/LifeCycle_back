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

    #データベースを特定
    table = Aws::DynamoDB::Resource.new(region: 'ap-northeast-1').table('Data')

    #追加する案件
    data={ posted: Time.new.strftime("%m/%d-%H:%M"), st: st, en: en, tag: tag, content: content }
    #一致するユーザのデータベース上のデータ
    db=table.get_item({key: { 'name' => name } })["item"]

    #ユーザが登録がまだされていなかったら新規登録、されていたら案件を追加
    plan_arr= db.nil? ? [data] : db["plan"].push(data)

    #データを入力
    table.put_item(json={ item: { name: name, plan: plan_arr } })

    { statusCode: 200, body: json }
end
