require 'websocket-client-simple'
require 'json'
require 'eventmachine'

api_key = ENV.fetch('KAIKO_API_KEY', '__REPLACE_ME__')

exchange = 'bnce' # Okex exchange
instrument_class = 'spot' # Spot market
instrument = '*-btc' # all instrument traded against BTC

url = 'wss://us.market-ws.kaiko.io/v2/data/trades_ws.latest/%s:%s:%s' % [exchange, instrument_class, instrument]

EM.run {
  ws = WebSocket::Client::Simple.connect url, :headers => {'Sec-WebSocket-Protocol' => 'api_key, %s' % [api_key]}

  ws.on :message do |msg|
    message = JSON.parse(msg.data)
    if message['event'] == 'update'
      puts 'Received a trade from %s %s %s' % [
        message['payload']['subscription']['exchange'],
        message['payload']['subscription']['instrument_class'],
        message['payload']['subscription']['instrument']
      ]
      puts message['payload']['data']
    end
  end

  ws.on :close do |e|
    p e
    exit 1
  end

  ws.on :error do |e|
    p e
  end
}
