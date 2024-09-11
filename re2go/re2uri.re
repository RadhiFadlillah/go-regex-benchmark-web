package main

func findURIs(bytes []byte) int {
	var count int
	var cur, mar int
	bytes = append(bytes, byte(0))
	lim := len(bytes) - 1

	for { /*!re2c
		re2c:eof = 0;
		re2c:yyfill:enable = 0;
		re2c:posix-captures = 0;
		re2c:case-insensitive = 0;
		re2c:define:YYCTYPE     = byte;
		re2c:define:YYPEEK      = "bytes[cur]";
		re2c:define:YYSKIP      = "cur += 1";
		re2c:define:YYBACKUP    = "mar = cur";
		re2c:define:YYRESTORE   = "cur = mar";
		re2c:define:YYLESSTHAN  = "lim <= cur";
		re2c:define:YYSTAGP     = "@@{tag} = cur";
		re2c:define:YYSTAGN     = "@@{tag} = -1";
		re2c:define:YYSHIFTSTAG = "@@{tag} += @@{shift}";

		uri = [0-9A-Z_a-z]+:[/][/][^\t\n\f\r #/\?]+[^\t\n\f\r #\?]+(![?][^\t\n\f\r #]*)?(!#[^\t\n\f\r ]*)?;

		{uri} { count += 1; continue }
		*     { continue }
		$     { return count }
		*/
	}
}
