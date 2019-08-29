require 'json'
require 'aws-sdk'

def lambda_handler(event:, context:)
    # TODO implement

    name=event["queryStringParameters"]["name"]
    #タグごとそれぞれの総時間を格納するためのハッシュ
    hash=Hash.new
    10.times{|i|hash[i.to_s]=0}

    dynamoDB = Aws::DynamoDB::Resource.new(region: 'ap-northeast-1')
    table = dynamoDB.table('Data')


    #データを取得
    resp = table.get_item({ key: { 'name' => name }})

    arr=resp["item"]["plan"]

    arr.each{|a|hash[a["tag"]]+= a["duar"].to_i}

    return { statusCode: 200, body: hash.to_json }
end
