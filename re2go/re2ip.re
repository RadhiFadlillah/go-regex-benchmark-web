package main

func findIPs(bytes []byte) int {
	var count int
	var cur, mar int
	lim := len(bytes)

	// Peek function
	peek := func(bytes []byte, cur int, lim int) byte {
		if cur < lim {
			return bytes[cur]
		}
		return 0
	}

	for { /*!re2c
		re2c:eof = 0;
		re2c:yyfill:enable = 0;
		re2c:posix-captures = 0;
		re2c:case-insensitive = 0;
		re2c:define:YYCTYPE     = byte;
		re2c:define:YYPEEK      = "peek(bytes, cur, lim)";
		re2c:define:YYSKIP      = "cur += 1";
		re2c:define:YYBACKUP    = "mar = cur";
		re2c:define:YYRESTORE   = "cur = mar";
		re2c:define:YYLESSTHAN  = "lim <= cur";
		re2c:define:YYSTAGP     = "@@{tag} = cur";
		re2c:define:YYSTAGN     = "@@{tag} = -1";
		re2c:define:YYSHIFTSTAG = "@@{tag} += @@{shift}";

		ip = ((2(!5[0-5]|[0-4][0-9])|[01]?[0-9][0-9])[\.])((2(!5[0-5]|[0-4][0-9])|[01]?[0-9][0-9])[\.])((2(!5[0-5]|[0-4][0-9])|[01]?[0-9][0-9])[\.])(!2(!5[0-5]|[0-4][0-9])|[01]?[0-9][0-9]);

		{ip} { count += 1; continue }
		*    { continue }
		$    { return count }
		*/
	}
}
