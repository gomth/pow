# Pow a number using golang

[![CircleCI](https://circleci.com/gh/gomth/pow/tree/main.svg?style=svg)](https://circleci.com/gh/gomth/pow/tree/main)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/be861bf57e5c433e93ac22ffb89e5158)](https://www.codacy.com/gh/gomth/pow/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=gomth/pow&amp;utm_campaign=Badge_Grade)

##### Usage:

```golang
import "github.com/gomth/pow"

val, err := pow.Get(2, 3)
if err != nil {
    return nil, err
}
```

##### Package contents:
| API | Engine | time complexity | space complexity | Details |
| ------ | ------ | ------ | ------ | ------ |
| Get | | | | |
| GetUsingFactorOfPower | | | | Convert power to a product of prime numbers. |

##### TODO:
- automate displaying benchmarks