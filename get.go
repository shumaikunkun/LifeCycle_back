require 'json'
require 'aws-sdk'

def lambda_handler(event:, context:)

    #パラメタからnameを取得
    name=event["queryStringParameters"]["name"]
    #name="test"

    dynamoDB = Aws::DynamoDB::Resource.new(region: 'ap-northeast-1')
    table = dynamoDB.table('Data')


    #データを取得
    resp = table.get_item({
        key: { 'name' => name }
    })

    { statusCode: 200, body: resp }
end
