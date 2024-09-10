build: generate
	@CGO_ENABLED=0 go build -o benchmark

generate:
	@regexp2go -pkg re -fn RxEmail '[\w\.+-]+@[\w\.-]+\.[\w\.-]+' > internal/re/rx-email.go
	@regexp2go -pkg re -fn RxURI '[\w]+://[^/\s?#]+[^\s?#]+(?:\?[^\s#]*)?(?:#[^\s]*)?' > internal/re/rx-uri.go
	@regexp2go -pkg re -fn RxIP '(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9])\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9])' > internal/re/rx-ip.go
	@regexp2go -pkg re -fn RxLongDate '(?i)(January?|February?|March|A[pv]ril|Ma[iy]|Jun[ei]|Jul[iy]|August|September|O[ck]tober|November|De[csz]ember|Jan|Feb|M[aä]r|Apr|Jun|Jul|Aug|Sep|O[ck]t|Nov|De[cz]|Januari|Februari|Maret|Mei|Agustus|Jänner|Feber|März|janvier|février|mars|juin|juillet|aout|septembre|octobre|novembre|décembre|Ocak|Şubat|Mart|Nisan|Mayıs|Haziran|Temmuz|Ağustos|Eylül|Ekim|Kasım|Aralık|Oca|Şub|Mar|Nis|Haz|Tem|Ağu|Eyl|Eki|Kas|Ara)\s([0-3]?[0-9])(?:st|nd|rd|th)?,?\s(199[0-9]|20[0-3][0-9])|([0-3]?[0-9])(?:st|nd|rd|th|\.)?\s(?:of\s)?(January?|February?|March|A[pv]ril|Ma[iy]|Jun[ei]|Jul[iy]|August|September|O[ck]tober|November|De[csz]ember|Jan|Feb|M[aä]r|Apr|Jun|Jul|Aug|Sep|O[ck]t|Nov|De[cz]|Januari|Februari|Maret|Mei|Agustus|Jänner|Feber|März|janvier|février|mars|juin|juillet|aout|septembre|octobre|novembre|décembre|Ocak|Şubat|Mart|Nisan|Mayıs|Haziran|Temmuz|Ağustos|Eylül|Ekim|Kasım|Aralık|Oca|Şub|Mar|Nis|Haz|Tem|Ağu|Eyl|Eki|Kas|Ara)[,.]?\s(199[0-9]|20[0-3][0-9])' > internal/re/rx-longdate.go
