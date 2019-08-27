require 'json'

def lambda_handler(event:, context:)
    res = {
        event: {
            class: event.class,
            inspect: event.inspect,
        },
        context: {
            class: context.class,
            inspect: context.inspect,
        },
    }
    { statusCode: 200, body: res.to_json }
end
