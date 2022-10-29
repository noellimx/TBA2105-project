query:
location, value_string

fields obtained:
tweets, retweets (using tweet text itself instead of retweeted text)
TEXT || FULL TEXT = PROCESSED TEXT

table should look like

yyyymmddHHMMSS | text  
20221227001235 | hi i am in jb

processing:

1. lemmatization
   caveat: all words will be lemmatized as english words

2. Aggregation

table should look like

yyyymmddhh | word | occurrence
2022122700 | traffic | 4
2022122723 | jam | 7
