require 'json'

def lambda_handler(event:, context:)
    res = {
    plus: event["A"]+event["B"],
    minus: event["A"]-event["B"]
    }

    { statusCode: 200, body: res }
end
