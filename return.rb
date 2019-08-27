require 'json'

def lambda_handler(event:, context:)
    res=event["message"]

    { statusCode: 200, body: res }
end
