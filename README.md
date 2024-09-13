# Go Regex Benchmark Web - Concurrent

This repo is a benchmark for various Golang's regular expressions library. Based on my benchmark [here][benchmark-old], which is based on benchmark by Rustem Kamalov [here][original-benchmark].

For its input, this benchmark use 1,058 web pages in [`000-input-files`](./000-input-files) directory. Hopefully it's good enough as representation of real world content.

The benchmark in this branch run concurrently using goroutine. If you are interested in the performance for single thread regex matching, please check out the [`master`][master-branch] branch in this repository.

## Table of Contents

- [How to Run](#how-to-run)
- [Regex Patterns](#regex-patterns)
  - [Short Regex](#short-regex)
  - [Long Regex](#long-regex)
- [Used Packages](#used-packages)
  - [Native Go Packages](#native-go-packages)
  - [Regex with CGO Binding](#regex-with-cgo-binding)
  - [Regex with Web Assembly Binding](#regex-with-web-assembly-binding)
  - [Regex Compiler](#regex-compiler)
- [Result](#result)
  - [Short Regex](#short-regex-1)
  - [Long Regex](#long-regex-1)
- [License](#license)

## How to Run

If you are using GNU Make, you can simply run `make` in the root directory of this benchmark:

```
make clean-all    # clean the old build
make build-all    # rebuild the benchmark executable
make              # run the benchmark
```

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

Since re2go is a compiler, it has its own way to handle regular expressions. Thanks to this, in this benchmark its result for email and URI is a bit "unfair" compared to the other packages. This is because for those cases re2go templates in this benchmark is optimized with multiple rules to handle regex with quantifier groups. For more details, please check out [this issue][re2go-issue] where I discuss re2go performance with its maintainer.

Another note: Code Search is explicitly mentioned in its documentation as not thread safe, so it can't be used concurrently. However, I still keep it in this benchmark as comparison for other package.

With that out of the way, here is the benchmark result. The benchmark was run on my Linux PC with Intel i7-8550U and RAM 16 GB.

### Short Regex

|    Package    |     Type      | Email (ms) |  URI (ms) | IP (ms) | Total (ms) |  Times |
| :-----------: | :-----------: | ---------: | --------: | ------: | ---------: | -----: |
|     re2go     |   Compiler    |     117.16 |    107.22 |   86.08 |     310.45 | 31.63x |
|    RE2 CGO    |      CGO      |      93.86 |    148.31 |   97.56 |     339.73 | 28.90x |
|     PCRE      |      CGO      |     216.15 |    159.67 |  159.42 |     535.23 | 18.34x |
| Code Search\* | Native (Grep) |     269.48 |    314.27 |  258.55 |     842.30 | 11.66x |
|   RE2 WASM    |     WASM      |     414.44 |    649.14 |  505.15 |    1568.74 |  6.26x |
|    Grafana    |    Native     |    1969.71 |   2149.27 | 3112.42 |    7231.40 |  1.36x |
|      Go       |    Native     |    2760.24 |   3073.57 | 3984.71 |    9818.52 |  1.00x |
|    Modernc    |    Native     |    2937.42 |   3279.70 | 4375.58 |   10592.70 |  0.93x |
|   Regexp2cg   |   Compiler    |  207302.07 | 106816.08 |  831.05 |  314949.20 |  0.03x |
|    Regexp2    |    Native     |  284351.98 | 145852.15 | 1116.03 |  431320.16 |  0.02x |
|   Hyperscan   |               |            |           |         |            |        |
|   Regexp2Go   |               |            |           |         |            |        |

Some interesting points:

- The optimized re2go template generates a very fast native Go code. It's as fast as RE2 with cgo. However since some Go's regex syntaxes are not supported by re2go, there are needs to modify the regex patterns before using it.
- It's amazing to see how fast Code Search is. Even though it's running in a single thread, it's still faster than most package, including RE2 with WASM.
- For code without cgo but with full regex compatibility and consistent performance, RE2 WASM is still the fastest though (since Code Search can't be used for daily use).
- Apparently Hyperscan is not thread safe so it's failed to run.
- For some reasons code that generated by Regexp2Go doesn't work, so it's skipped for now.

### Long Regex

|    Package    |     Type      | Long Date (ms) |   Times |
| :-----------: | :-----------: | -------------: | ------: |
|    RE2 CGO    |      CGO      |         121.14 | 963.72x |
|     re2go     |   Compiler    |         267.98 | 435.65x |
| Code Search\* | Native (Grep) |         359.43 | 324.80x |
|   RE2 WASM    |     WASM      |         528.07 | 221.08x |
|     PCRE      |      CGO      |        2854.28 |  40.90x |
|    Grafana    |    Native     |       32614.29 |   3.58x |
|   Regexp2cg   |   Compiler    |       50775.91 |   2.30x |
|    Regexp2    |    Native     |       51528.74 |   2.27x |
|      Go       |    Native     |      116743.85 |   1.00x |
|    Modernc    |    Native     |      117094.39 |   1.00x |
|   Hyperscan   |               |                |         |
|   Regexp2Go   |               |                |         |

Some interesting points:

- It's amazing to see how fast Code Search is. Even though it's running in a single thread, it's still faster than most package.
- The optimized re2go template generates a very fast native Go code. It's almost as fast as RE2 with cgo.
- PCRE's performance become a lot slower for long regex. It's even slower than RE2 WASM.
- For native Go code, Grafana is pretty fast at handling long regex pattern.
- For code without cgo, regex that compiled by re2go has the best performance.
- For code without cgo but with full regex compatibility, RE2 WASM has the best performance.
- Hyperscan is not thread safe so it's failed to run.
- For some reasons code that generated by Regexp2Go doesn't work, so it's skipped for now.

## License

Like the original benchmark, this benchmark is also released under MIT license.

[original-benchmark]: https://github.com/karust/regex-benchmark
[benchmark-old]: https://github.com/RadhiFadlillah/go-regex-benchmark
[master-branch]: https://github.com/RadhiFadlillah/go-regex-benchmark-web/tree/master
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
[re2go-issue]: https://github.com/skvadrik/re2c/issues/487
