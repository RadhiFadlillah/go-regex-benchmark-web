package main

func findEmails(bytes []byte) int {
	var count int
	var cur, mar int
	bytes = append(bytes, byte(0)) // add terminating null
	lim := len(bytes) - 1 // lim points at the terminating null

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

		quant1 = [+\-.0-9A-Z_a-z]+;
		quant2 = [+\-.0-9A-Z_a-z]+@[\-.0-9A-Z_a-z]+;
		email  = [+\-.0-9A-Z_a-z]+@[\-.0-9A-Z_a-z]+[.][\-.0-9A-Z_a-z]+;

		{email}  { count += 1; continue }
		{quant1} { continue }
		{quant2} { continue }
		*        { continue }
		$        { return count }
		*/
	}
}
