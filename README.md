## gtwitter  
gwitter is a CLI for the Twitter API. Please, check usage what can do with this tool.
To use this CLI tool, "ConsumerKey", "ConsumerSecret", "AccessToken", "AccessSecret" are necessary.Visit [Twitter Developer's site](https://developer.twitter.com/en/docs/basics/authentication/guides/access-tokens.html) to get Tokens.
After you get Tokens, add token to "token.toml" file or set token by using "config" subcommand.

### Usage

#### Install 
```
go install github.com/ShogoTomioka/gtwitter                
```

#### Set Config
Check [Twitter Developer's site](https://developer.twitter.com/en/docs/basics/authentication/guides/access-tokens.html) to know how to get `Accsess Token` and `Consumer Keys`.

Add your Tokens ("ConsumerKey", "ConsumerSecret", "AccessToken", "AccessSecret") to config/token.tomk or 
```
gtwitter config -consumer       [your consumer Key]       \
                -consumerSecret [your consumerSecret Key] \
                -accsessToken   [your accsessToken]       \
                -accessSecret   [your accessSecret Key]   
```


#### Tweeting
`Tweet` subcommand provides you tweeting feature through command line with your tweet text.
``` 
gtwitter tweet ["you tweet text"]
```

#### Show Timeline
`Timeline` subcommand show your timeline.You can select number of timelines with `-count` or `-c` option.
```
// Show 10 timeline tweets, default is 20
gtwitter timeline  -c 10
```

#### Show Trends
`Trends` subcommand show you trends with `WOID=1110809(Tokyo)`.You can select number of showing trends with `-count` or `-c` option.
```
// Show 10 trends, default is 20
gtwitter trends -c 10
```

#### Search Tweets with word
`Search` subcommand provides you searching feature by word you want to search. You can select numnber of searching tweets with `-count` or `-c` option.
```
// Show 10 tweets searched out with "Bodybuilding", default is 20
gtwitter search -w Bodybuilding -c 10
```

#### Contributing
If you find mistake, feel free to send PR or Issue

#### LICENSE
MIT LICENSE