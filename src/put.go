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
    #タグごとの総時間をハッシュに格納
    resp["item"]["plan"].each{|a|hash[a["tag"]]+=a["duar"].to_i}
    #タグごとの時間のパーセンテージを格納
    hash_rate=Hash.new
    hash.each{|k,v|hash_rate[k]=v*100.0/hash.values.inject(:+)}
    #得点を計算
    point=0
    hash_rate.each do |k,v|  #開発と勉強はそのまま得点化
        if k=="1"||k=="2"
            point+=v
        elsif k=="0"||k=="4"||k=="7"  #睡眠と運動と食事は半分だけ加算
            point+=v/2
        end
    end

    hash["point"]=point


    return { statusCode: 200, body: hash.to_json }
end
