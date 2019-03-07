# Kaiko livestream API - Ruby example


## Requirements
* Kaiko API key [contact us](https://www.kaiko.com/pages/contact-1)
* Ruby 2.5.1+
* bundler 1.17+

## Getting started
```bash
$ bundler isntall
$ export KAIKO_API_KEY=__REPLACE_WITH_YOUR_KEY__
$ ruby example.rb
Received a trade from okex spot int-btc
{"timestamp"=>1552297292548, "trade_id"=>"26506842", "price"=>"0.00000396", "amount"=>"3776.43", "taker_side_sell"=>true}
...

```
