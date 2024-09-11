package main

func findLongDatePattern(bytes []byte) int {
	var count int
	var cur, mar int
	bytes = append(bytes, byte(0))
	lim := len(bytes) - 1

	// Capturing groups
	/*!maxnmatch:re2c*/	yypmatch := make([]int, YYMAXNMATCH*2)
	var yynmatch int
	var yyt1, yyt2, yyt3, yyt4, yyt5, yyt6, yyt7, yyt8, yyt9, yyt10 int
	_ = yynmatch

	for { /*!re2c
		re2c:eof = 0;
		re2c:yyfill:enable = 0;
		re2c:posix-captures = 1;
		re2c:case-insensitive = 1;
		re2c:define:YYCTYPE     = byte;
		re2c:define:YYPEEK      = "bytes[cur]";
		re2c:define:YYSKIP      = "cur += 1";
		re2c:define:YYBACKUP    = "mar = cur";
		re2c:define:YYRESTORE   = "cur = mar";
		re2c:define:YYLESSTHAN  = "lim <= cur";
		re2c:define:YYSTAGP     = "@@{tag} = cur";
		re2c:define:YYSTAGN     = "@@{tag} = -1";
		re2c:define:YYSHIFTSTAG = "@@{tag} += @@{shift}";

		rxDay = [0-3]?[0-9];
		rxYear = 199[0-9]|20[0-3][0-9];
		rxMonth = January?|February?|March|A[pv]ril|Ma[iy]|Ju(!n[ei]|l[iy])|August|September|O[ck]tober|November|De[csz]ember|Jan|Feb|M[aä]r|Apr|Ju[ln]|Aug|Sep|O[ck]t|Nov|De[cz]|Januari|Februari|M(!aret|ei)|Agustus|Jänner|Feber|März|janvier|février|mars|jui(!n|llet)|aout|septembre|octobre|novembre|décembre|Ocak|Şubat|Mart|Nisan|Mayıs|Haziran|Temmuz|Ağustos|E(!ylül|kim)|Kasım|Aralık|Oca|Şub|Mar|Nis|Haz|Tem|Ağu|E(!yl|ki)|Kas|Ara;
		rxLongPattern = ({rxMonth})[\t\n\f\r ]({rxDay})(!st|nd|rd|th)?,?[\t\n\f\r ]({rxYear})|({rxDay})(!st|nd|rd|th|[\.])?[\t\n\f\r ](!of[\t\n\f\r ])?({rxMonth})[,\.]?[\t\n\f\r ]({rxYear});

		{rxLongPattern} {
			if yynmatch != 7 {
				panic("expected 7 submatch groups")
			}
			count += 1
			continue
		}

		* { continue }
		$ { return count }
		*/
	}
}
