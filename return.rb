require 'json'

def lambda_handler(event:, context:)
    res=event["A"]+event["B"]

    { statusCode: 200, body: res }
end
