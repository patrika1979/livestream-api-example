import json
import websocket
from os import environ

def on_message(ws, data):
    message = json.loads(data)
    if message["event"] == "update":
        print('Received a trade from {}:{}:{}'.format(
            message["payload"]["subscription"]["exchange"],
            message["payload"]["subscription"]["instrument_class"],
            message["payload"]["subscription"]["instrument"])
        )
        print(json.dumps(message["payload"]["data"], indent=4))

def on_error(ws, error):
    print(error)

def on_close(ws):
    print("### closed ###")

if __name__ == "__main__":
    env = environ.get('KAIKO_API_KEY')
    api_key = env if env is not None else '__REPLACE_ME__'

    exchange = 'bnce' # Binance exchange
    instrument_class = 'spot' # Spot market
    instrument = '*-btc' # all instrument traded against BTC
    pattern = ":".join([exchange, instrument_class, instrument])

    websocketUrl = 'wss://us.market-ws.kaiko.io/v2/data/trades_ws.latest/{}'.format(pattern);
    ws = websocket.WebSocketApp(websocketUrl,
                              on_message = on_message,
                              on_error = on_error,
                              on_close = on_close,
                              header = ['Sec-WebSocket-Protocol: api_key, {}'.format(api_key)])
    ws.run_forever()
