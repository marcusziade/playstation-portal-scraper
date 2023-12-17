# Playstation Portal sales availability monitor

Example output:
```
$ go run .

2023/12/17 13:39:45 Requesting URL: https://direct.playstation.com/en-us/buy-accessories/playstation-portal-remote-player
2023/12/17 13:39:45 Received response: 200 from https://direct.playstation.com/en-us/buy-accessories/playstation-portal-remote-player
2023/12/17 13:39:45 Found element with text: 'Sign In'
2023/12/17 13:39:45 Found element with text: 'Currently Unavailable'
2023/12/17 13:39:45 Found element with text: 'Low stock'
2023/12/17 13:39:45 Found element with text: 'Coming Soon'
2023/12/17 13:39:45 Found element with text: '11/15/2023'
The page has the 'Currently Unavailable' label.
```