## Intro
This is an app written in Go, helping you to convert money from one currency to another.

It uses the [ECB foreign exchange rates](https://www.ecb.europa.eu/stats/policy_and_exchange_rates/euro_reference_exchange_rates/html/index.en.html).

## Request Example
API Endpoint:
http://127.0.0.1:9090/api/v1/convert

Body Params:

```
{
  "from_currency": "CAD",
  "to_currency": "USD",
  "amount": 129,65,
  "date": "2022-06-14" -- The exchange rates are based on this date.
}
```


The supported currencies are listed [here](https://www.ecb.europa.eu/stats/policy_and_exchange_rates/euro_reference_exchange_rates/html/index.en.html).