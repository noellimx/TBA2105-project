# Status

Found fix but not yet tested

# Behavior

during `cT.twitterExampleFullArchiveSearchV1("hello")`

Output: ` {"errors":[{"message":"You currently have Essential access which includes access to Twitter API v2 endpoints only. If you need access to this endpoint, you’ll need to apply for Elevated access via the Developer Portal. You can learn more here: https://developer.twitter.com/en/docs/twitter-api/getting-started/about-twitter-api#v2-access-leve","code":453}]}`

# Notes

## Search

curl --request POST \
 --url https://api.twitter.com/1.1/tweets/search/fullarchive/production.json \
 --header 'authorization: Bearer aaa' \
 --header 'content-type: application/json' \
 --data '{
"query":"from:TwitterDev lang:en",
"maxResults": "100",
"fromDate":"201802010000",
"toDate":"201802282359"
}'
