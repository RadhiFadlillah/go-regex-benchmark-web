# Go Regex Benchmark Web

This repo is a benchmark for various Golang's regular expressions library. Based on my benchmark [here][benchmark-old], which is based on benchmark by Rustem Kamalov [here][original-benchmark].

For its input, this benchmark use 1,058 web pages in [`000-input-files`](./000-input-files) directory. Hopefully it's good enough as representation of real world content.

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

The benchmark was run on Linux with Intel i7-8550U with RAM 16 GB.

### Short Regex

|   Package   |     Type      | Email (ms) |  URI (ms) |  IP (ms) | Total (ms) |
| :---------: | :-----------: | ---------: | --------: | -------: | ---------: |
|   RE2 CGO   |      CGO      |     221.29 |    345.34 |   224.52 |     791.15 |
| Code Search | Native (Grep) |     263.00 |    300.81 |   257.84 |     821.65 |
|    PCRE     |      CGO      |     764.56 |    486.42 |   462.49 |    1713.47 |
|  Hyperscan  |      CGO      |     595.38 |   2654.13 |    41.14 |    3290.65 |
|  RE2 WASM   |     WASM      |    1189.12 |   1533.30 |  1187.08 |    3909.50 |
| Go std lib  |    Native     |    6544.45 |   6926.63 | 10321.36 |   23792.44 |
|   Grafana   |    Native     |    6600.02 |   7148.91 | 10169.21 |   23918.14 |
|   Modernc   |    Native     |    6715.16 |   7216.19 | 10309.89 |   24241.23 |
|    re2go    |   Compiler    |   31545.11 |   7748.29 |   467.68 |   39761.08 |
|  Regexp2cg  |   Compiler    |  452673.57 | 191651.48 |  2899.08 |  647224.13 |
|   Regexp2   |    Native     |  568837.37 | 239814.81 |  3717.81 |  812369.98 |
|  Regexp2Go  |   Compiler    |            |           |          |            |

Some interesting points:

- It's amazing to see how fast Code Search is. It's almost as fast as RE2 with cgo.
- For some reasons code that generated by re2go has inconsistent performance. For the shortest pattern (email) it's slower than Go's standard library. However for the longest pattern (IP) it's faster than PCRE that uses cgo binding.
- For code without cgo but with full regex compatibility and consistent performance, RE2 WASM is the fastest.
- For some reasons code that generated by Regexp2Go doesn't work, so it's skipped for now.

### Long Regex

|   Package   |     Type      | Long Date (ms) |    Times |
| :---------: | :-----------: | -------------: | -------: |
|  Hyperscan  |      CGO      |          52.81 | 6338.55x |
|   RE2 CGO   |      CGO      |         233.61 | 1432.99x |
| Code Search | Native (Grep) |         365.35 |  916.30x |
|    re2go    |   Compiler    |         903.54 |  370.50x |
|  RE2 WASM   |     WASM      |        1201.96 |  278.51x |
|    PCRE     |      CGO      |        7421.20 |   45.11x |
|   Grafana   |    Native     |       83004.71 |    4.03x |
|   Regexp2   |    Native     |      134163.71 |    2.50x |
|  Regexp2cg  |   Compiler    |      152104.35 |    2.20x |
|   Modernc   |    Native     |      330399.26 |    1.01x |
| Go std lib  |    Native     |      334764.53 |    1.00x |
|  Regexp2Go  |   Compiler    |                |          |

Some interesting points:

- Hyperscan is really fast at handling long regex pattern.
- PCRE's performance become a lot slower for long regex compared to RE2 and Hyperscan. It's even slower than RE2 WASM.
- For native Go code, Grafana is pretty fast at handling long regex pattern. It's even faster than regex that compiled by Regexp2Go.
- For code without cgo, regex that compiled by re2go has the best performance.
- For code without cgo but with full regex compatibility, RE2 WASM has the best performance.
- For some reasons code that generated by Regexp2Go doesn't work, so it's skipped for now.

## License

Like the original benchmark, this benchmark is also released under MIT license.

[original-benchmark]: https://github.com/karust/regex-benchmark
[benchmark-old]: https://github.com/RadhiFadlillah/go-regex-benchmark
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
