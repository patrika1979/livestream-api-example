# Kaiko livestream API - Python example


## Requirements
* Kaiko API key [contact us](https://www.kaiko.com/pages/contact-1)
* Python 2.7+
* pip 9+

## Getting started
```bash
$ pip install -r requirements.txt
$ export KAIKO_API_KEY=__REPLACE_WITH_YOUR_KEY__
$ python example.py
Received a trade from okex:spot:let-btc
[
    {
        "timestamp": 1552297148594, 
        "trade_id": "3037750", 
        "taker_side_sell": false, 
        "price": "0.0000011378", 
        "amount": "9264"
    }
]
...

```
