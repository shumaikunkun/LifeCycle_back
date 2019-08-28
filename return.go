require 'json'
require 'aws-sdk'

def lambda_handler(event:, context:)
    print event
    uid=event["queryStringParameters"]["uid"].to_i
    #パラメタからuidを取得


    dynamoDB = Aws::DynamoDB::Resource.new(region: 'ap-northeast-1')
    table = dynamoDB.table('Datas')
    # resp = table.get_item({
    #     key: { 'uid' => 0 }  #成功すれば0=>uidに変える
    # })

    table.put_item(
      {
      item: {
        uid: uid,
        usr: "name#{uid}",
        plan:
          [
            { posted:"2017", st:"2017", en:"2017", tag:1 },
            { posted:"2017", st:"2017", en:"2017", tag:2 },
            { posted:"2017", st:"2017", en:"2017", tag:3 },
          ]
      }
    }
    )

    { statusCode: 200, body: uid }  #成功すればuid=>respに変える
end

// https://qm2ju9z5y5.execute-api.ap-northeast-1.amazonaws.com/dev/test?uid=1234
