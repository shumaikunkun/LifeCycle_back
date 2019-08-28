require 'json'

def lambda_handler(event:, context:)
    print event
    uid=event["queryStringParameters"]["uid"].to_i

    { statusCode: 200, body: uid }
end
