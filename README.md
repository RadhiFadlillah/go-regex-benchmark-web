# Go Regex Benchmark

This repo is a benchmark for various Golang's regular expressions library. Based on benchmark by Rustem Kamalov [here][original-benchmark].

Unlike the original repository, here I only focus on Go language without caring about its performance compared to the other languages.

## Table of Contents

- [Input Text](#input-text)
- [Regex Patterns](#regex-patterns)
  - [Short Regex](#short-regex)
  - [Long Regex](#long-regex)
- [Measurement](#measurement)
- [Used Packages](#used-packages)
  - [Native Go Packages](#native-go-packages)
  - [Regex with CGO Binding](#regex-with-cgo-binding)
  - [Regex with Web Assembly Binding](#regex-with-web-assembly-binding)
  - [Regex Compiler](#regex-compiler)
- [Result](#result)
  - [Short Regex](#short-regex-1)
  - [Long Regex](#long-regex-1)
- [License](#license)

## Input Text

Like the original, the [input text](./input-text.txt) is a concatenation of [Learn X in Y minutes][x-in-y] repository.

## Regex Patterns

In this benchmark, there are 2 kinds of regexes that will be tested: short regex and long regex.

### Short Regex

For the short regex, we use 3 patterns from the original benchmark:

- Email: `[\w\.+-]+@[\w\.-]+\.[\w\.-]+`
- URI: `[\w]+://[^/\s?#]+[^\s?#]+(?:\?[^\s#]*)?(?:#[^\s]*)?`
- IPv4: `(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9])\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9])`

### Long Regex

For the long regex, we use pattern for detecting multilingual long date texts. Given the following patterns:

- day:

  ```
  [0-3]?[0-9]
  ```

- year:

  ```
  199[0-9]|20[0-3][0-9]
  ```

- month:

  ```
  January?|February?|March|A[pv]ril|Ma[iy]|Jun[ei]|Jul[iy]|August|September|O[ck]tober|November|De[csz]ember|Jan|Feb|M[aä]r|Apr|Jun|Jul|Aug|Sep|O[ck]t|Nov|De[cz]|Januari|Februari|Maret|Mei|Agustus|Jänner|Feber|März|janvier|février|mars|juin|juillet|aout|septembre|octobre|novembre|décembre|Ocak|Şubat|Mart|Nisan|Mayıs|Haziran|Temmuz|Ağustos|Eylül|Ekim|Kasım|Aralık|Oca|Şub|Mar|Nis|Haz|Tem|Ağu|Eyl|Eki|Kas|Ara
  ```

The final pattern is defined as:

```
(?i)({month})\s({day})(?:st|nd|rd|th)?,?\s({year})|({day})(?:st|nd|rd|th|\.)?\s(?:of\s)?({month})[,.]?\s({year})
```

## Measurement

Unlike the original, measuring is done without including regex compilation. So the measurement only focused on pattern matching.

The measurement are done 10 times, then the smallest durations are used as the final durations.

## Used Packages

There are 12 regular expressions that used in this benchmark:

- 5 are regex packages in native Go code.
- 3 are regex packages with CGO binding.
- 1 is regex packages with WASM binding.
- 3 are regex compiler to compile regular expressions to native Go code.

### Native Go Packages

1. **Go** is the `regexp` package from Go's standard library.
2. **Grafana** is the package from [`github.com/grafana/regexp@speedup`][grafana] by Bryan Boreham that improve the standard library with several optimations. Actively maintained at the time this benchmark is written.
3. **Modernc** is the package from [`modernc.org/regexp`][modernc] that implements experimental DFA support. However, it doesn't use DFA implementation from `google/re2` and instead uses its own implementation.
4. **Regexp2** is the package from [`github.com/dlclark/regexp2`][regexp2] that ports regex engine from .NET frameworks.
5. **Code Search** is the package from [`github.com/google/codesearch/regexp`][codesearch]. It uses DFA algorithm with trigram indexing and was used by Google Code Search.

   Since it's used for search engine, currently its API only supports grep-like matching as in for every matching patterns it will returns the entire line instead of only returning the matching phrase.

   It's not thread safe, so every regex can only be used one goroutine at a time. And since its API only supports grep matching, currently it's not suitable for daily use. However it's still benchmarked here to glimpse the possible performance that can be achieved by native Go's regex engine in the future.

### Regex with CGO Binding

1. **RE2 CGO** is the package from [`github.com/wasilibs/go-re2`][go-re2] that binds [`google/re2`][google-re2] regex engine using cgo. Since Go's regex also use `google/re2` syntax, it can be used as drop in replacement for Go's native regex.
2. **Hyperscan** is the package from [`github.com/flier/gohs`][go-hyperscan] that binds Intel's [`hyperscan`][hyperscan] regex engine using cgo.
3. **PCRE** is the package from [`github.com/GRbit/go-pcre`][go-pcre] that binds PCRE regex engine using cgo.

### Regex with Web Assembly Binding

1. **RE2 WASM** is the package from [`github.com/wasilibs/go-re2`][go-re2] that binds [`google/re2`][google-re2] regex engine using Web Assembly. Like its cgo counterpart, it can be used as drop in replacement for Go's native regex.

### Regex Compiler

1. **re2go** is a regex compiler from [`re2c.org`][re2c] that compiles regular expressions into a native Go codes. Despite its name, it's not related with `google/re2` and uses its own lookahead TDFA algorithm.

   Since originally it only supports C, `re2go` use its own regex syntax which is not really compatible with Go's regex syntax. Fortunately it also supports [Flex regex syntax][flex] which is kinda similar with Go's syntax, so modifying the existing pattern is pretty easy.

   It's actively maintained and has been used in production and many open source projects, e.g. PHP and Ninja.

2. **Regexp2go** is a regex compiler from [`github.com/CAFxX/regexp2go`][regexp2go]. Despite its name, it's not related with `github.com/dlclark/regexp2` package mentioned above.

   It's similar in spirit to `re2go`, but aiming for compatibility with Go's regex syntax. At the time this benchmark written, this compiler hasn't been updated for 2 years and its documentation mentioned that it's not recommended to use in production.

   However for basic regex and with enough testing I reckon it should be good enough to use.

3. **Regexp2cg** is a regex compiler from [`github.com/dlclark/regexp2cg`][regexp2cg]. It's related with `github.com/dlclark/regexp2` package mentioned above.

   It will compile regular expressions into Go codes which can be used by the `dlclark/regexp2` package.

## Result

The benchmark was run on Linux with Intel i7-8550U with RAM 16 GB.

### Short Regex

|   Package   |     Type      | Email (ms) | URI (ms) | IP (ms) | Total (ms) |  Times |
| :---------: | :-----------: | ---------: | -------: | ------: | ---------: | -----: |
|   RE2 CGO   |      CGO      |      13.71 |    15.39 |   10.81 |      39.91 | 22.31x |
| Code Search | Native (Grep) |      14.89 |    14.70 |   12.66 |      42.25 | 21.07x |
|  Hyperscan  |      CGO      |      34.68 |    28.15 |    0.83 |      63.65 | 13.99x |
|    PCRE     |      CGO      |      31.67 |    25.75 |    9.19 |      66.61 | 13.37x |
|    re2go    |   Compiler    |      79.54 |    45.90 |   21.48 |     146.91 |  6.06x |
|  RE2 WASM   |     WASM      |      52.37 |    57.68 |   50.26 |     160.31 |  5.55x |
|  Regexp2Go  |   Compiler    |      78.91 |   610.26 |  173.29 |     862.46 |  1.03x |
| Go std lib  |    Native     |     250.84 |   245.60 |  393.85 |     890.28 |  1.00x |
|   Modernc   |    Native     |     255.67 |   242.96 |  392.64 |     891.28 |  1.00x |
|   Grafana   |    Native     |     268.57 |   248.11 |  402.69 |     919.37 |  0.97x |
|  Regexp2cg  |   Compiler    |    1896.41 |  1718.47 |   66.01 |    3680.89 |  0.24x |
|   Regexp2   |    Native     |    2229.39 |  2051.56 |   76.79 |    4357.75 |  0.20x |

Some interesting points:

- It's amazing to see how fast Code Search is. It's almost as fast as RE2 with cgo.
- For code without cgo, regex that compiled by re2go has the best performance. However, since some Go's regex syntaxes are not supported by re2go, there are needs to modify the regex patterns before using it.
- For code without cgo but with full regex compatibility, RE2 WASM has the best performance albeit a bit slower than re2go.

### Long Regex

|   Package   |    Type     | Long Date (ms) |     Times |
| :---------: | :---------: | -------------: | --------: |
|  Hyperscan  |     CGO     |           1.10 | 11766.05x |
|   RE2 CGO   |     CGO     |          13.60 |   949.04x |
| Code Search | Native Grep |          46.00 |   280.55x |
|    re2go    |  Compiler   |          47.20 |   273.43x |
|  RE2 WASM   |    WASM     |          60.80 |   212.25x |
|    PCRE     |     CGO     |         167.13 |    77.22x |
|   Grafana   |   Native    |        3330.87 |     3.87x |
|   Regexp2   |   Native    |        4319.13 |     2.99x |
|  Regexp2cg  |  Compiler   |        5031.48 |     2.56x |
|  Regexp2Go  |  Compiler   |        7435.22 |     1.74x |
|   Modernc   |   Native    |       12639.61 |     1.02x |
| Go std lib  |   Native    |       12905.60 |     1.00x |

Some interesting points:

- Hyperscan is really fast at handling long regex pattern. It's even faster than when it's used for short regex.
- PCRE's performance become a lot slower for long regex compared to RE2 and Hyperscan. It's even slower than RE2 WASM.
- For native Go code, Grafana is pretty fast at handling long regex pattern. It's even faster than regex that compiled by Regexp2Go.
- For code without cgo, regex that compiled by re2go has the best performance. It has similar performance as Code Search.
- For code without cgo but with full regex compatibility, RE2 WASM has the best performance (since Code Search currently can't be used as daily regex engine).

## License

Like the original benchmark, this benchmark is also released under MIT license.

[original-benchmark]: https://github.com/karust/regex-benchmark
[x-in-y]: https://github.com/adambard/learnxinyminutes-docs
[grafana]: https://github.com/grafana/regexp/tree/speedup?tab=readme-ov-file
[modernc]: https://gitlab.com/cznic/regexp
[regexp2]: https://github.com/dlclark/regexp2
[codesearch]: https://github.com/google/codesearch
[go-re2]: https://github.com/wasilibs/go-re2
[google-re2]: https://github.com/google/re2
[go-hyperscan]: https://github.com/flier/gohs
[hyperscan]: https://www.intel.com/content/www/us/en/developer/articles/technical/introduction-to-hyperscan.html
[go-pcre]: https://github.com/GRbit/go-pcre
[re2c]: https://re2c.org/manual/manual_go.html
[flex]: https://github.com/westes/flex
[regexp2go]: https://github.com/CAFxX/regexp2go
[regexp2cg]: https://github.com/dlclark/regexp2cg
