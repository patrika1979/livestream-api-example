const WebSocket = require('ws');

const apiKey = process.env.KAIKO_API_KEY || '__REPLACE_ME__';

const exchange = 'bnce'; // Binance exchange
const instrument_class = 'spot'; // Spot market
const instrument = '*-btc'; // all instrument traded against BTC

const pattern = [exchange, instrument_class, instrument].join(':');

const websocketUrl = `wss://us.market-ws.kaiko.io/v2/data/trades_ws.latest/${pattern}`;

const ws = new WebSocket(websocketUrl, ['api_key', apiKey]);

ws.onmessage = ({ data }) => {
  const message = JSON.parse(data);
  if (message.event === 'update') {
    console.log(`Received a trade from ${message.payload.subscription.exchange}:${message.payload.subscription.instrument_class}:${message.payload.subscription.instrument}`);
    console.log(message.payload.data);
  }
};

ws.onclose = () => console.log('Connection closed');
