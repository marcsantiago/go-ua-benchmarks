
# go-ua-benchmarks
 
This benchmarks 3 Go UserAgent Parsing Packages:

1) github.com/avct/uasurfer

2) github.com/ua-parser/uap-go

3) github.com/xojoc/useragent

  

# uasurfer benchmarks

|benchmark |operations  | ns/op |bytes/op | allocs/op
|--|--|--|--|--|
|BenchmarkIsValidBadUserAgent-8 |500000  |3111 |144 |2 |
|BenchmarkIsValidIOS-8 |500000  |3043 |144 |2 |
|BenchmarkIsValidAndroid-8 |500000  |3022 |144 |2 |
  

# useragent

|benchmark |operations  | ns/op |bytes/op | allocs/op
|--|--|--|--|--|
|BenchmarkIsValidBadUserAgent-8 |500000  |2151 |2400 |26 |
|BenchmarkIsValidIOS-8 |1000000  |2146 |2400 |26 |
|BenchmarkIsValidAndroid-8 |1000000  |2148 |2400 |26 |
    

# uap-go

|benchmark |operations  | ns/op |bytes/op | allocs/op
|--|--|--|--|--|
|BenchmarkIsValidBadUserAgent-8 |50  |22466297 |33484449 |95635 |
|BenchmarkIsValidIOS-8 |50  |22671333 |33484689 |95637 |
|BenchmarkIsValidAndroid-8 |50  |22971360 |33484803 |95636 |
