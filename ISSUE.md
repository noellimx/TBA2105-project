


# Behavior


during ```cT.twitterExampleFullArchiveSearchV1("hello")```

Output: ```
{"errors":[{"message":"You currently have Essential access which includes access to Twitter API v2 endpoints only. If you need access to this endpoint, you’ll need to apply for Elevated access via the Developer Portal. You can learn more here: https://developer.twitter.com/en/docs/twitter-api/getting-started/about-twitter-api#v2-access-leve","code":453}]}```






# Notes


curl --request POST \
  --url https://api.twitter.com/1.1/tweets/search/30day/prod/counts.json \
  --header 'authorization: Bearer a%a' \
  --header 'content-type: application/json' \
  --data '{
                "query":"from:TwitterDev lang:en",
                "fromDate":"201811010000", 
                "toDate":"201811060000"
                "bucket": "day"
                }'
