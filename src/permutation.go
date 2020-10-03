package permutation

import(
	"fmt"
	"unicode"
	"strings"
	"golang.org/x/net/publicsuffix"
)

func RunPermutations(target string) {
		sanitizedDomain, tld ,sub := ProcessInput(target)
		printReport(additionAttack(sub),sanitizedDomain, tld)
		printReport( omissionAttack(sub),sanitizedDomain, tld)
		printReport( homographAttack(sub),sanitizedDomain, tld)
		printReport( subdomainAttack(sub),sanitizedDomain, tld)
		printReport( vowelswapAttack(sub),sanitizedDomain, tld)
		printReport( repetitionAttack(sub),sanitizedDomain, tld)
		printReport( hyphenationAttack(sub), sanitizedDomain, tld)
		printReport( replacementAttack(sub), sanitizedDomain, tld)
		printReport( bitsquattingAttack(sub),sanitizedDomain, tld)
		printReport( transpositionAttack(sub), sanitizedDomain, tld)
}
func printReport(results []string, sanitizedDomain string, tld string) {
		for _, result := range results {
			fmt.Println(result + "." +sanitizedDomain+"."+ tld)
		}


}
// returns a count of characters in a word
func countChar(word string) map[rune]int {
	count := make(map[rune]int)
	for _, r := range []rune(word) {
		count[r]++
	}
	return count
}
//Return the domain,tld, subdomain
func ProcessInput(input string) (sanitizedDomain, tld string, sub string) {
		tldPlusOne, _ := publicsuffix.EffectiveTLDPlusOne(input)
		tld, _ = publicsuffix.PublicSuffix(tldPlusOne)
		sanitizedDomain = strings.Replace(tldPlusOne, "."+tld, "", -1)
		sub = strings.Replace(input, "."+sanitizedDomain+"."+tld, "",-1)

	return sanitizedDomain, tld, sub
}

// performs an addition attack adding a single character to the domain
func additionAttack(domain string) []string {
	results := []string{}

	for i := 97; i < 123; i++ {
		results = append(results, fmt.Sprintf("%s%c", domain, i))
	}
	return results
}

// performs a vowel swap attack
func vowelswapAttack(domain string) []string {
	results := []string{}
	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'y'}
	runes := []rune(domain)

	for i := 0; i < len(runes); i++ {
		for _, v := range vowels {
			switch runes[i] {
			case 'a', 'e', 'i', 'o', 'u', 'y':
				if runes[i] != v {
					results = append(results, fmt.Sprintf("%s%c%s", string(runes[:i]), v, string(runes[i+1:])))
				}
			default:
			}
		}
	}
	return results
}

// performs a transposition attack swapping adjacent characters in the domain
func transpositionAttack(domain string) []string {
	results := []string{}
	for i := 0; i < len(domain)-1; i++ {
		if domain[i+1] != domain[i] {
			results = append(results, fmt.Sprintf("%s%c%c%s", domain[:i], domain[i+1], domain[i], domain[i+2:]))
		}
	}
	return results
}

// performs a subdomain attack by inserting dots between characters, effectively turning the
// domain in a subdomain
func subdomainAttack(domain string) []string {
	results := []string{}
	runes := []rune(domain)

	for i := 1; i < len(runes); i++ {
		if (rune(runes[i]) != '-' || rune(runes[i]) != '.') && (rune(runes[i-1]) != '-' || rune(runes[i-1]) != '.') {
			results = append(results, fmt.Sprintf("%s.%s", string(runes[:i]), string(runes[i:])))
		}
	}
	return results
}

// performs a replacement attack simulating a user pressing the wrong keys
func replacementAttack(domain string) []string {
	results := []string{}
	keyboards := make([]map[rune]string, 0)
	count := make(map[string]int)
	keyboardEn := map[rune]string{'q': "12wa", '2': "3wq1", '3': "4ew2", '4': "5re3", '5': "6tr4", '6': "7yt5", '7': "8uy6", '8': "9iu7", '9': "0oi8", '0': "po9",
		'w': "3esaq2", 'e': "4rdsw3", 'r': "5tfde4", 't': "6ygfr5", 'y': "7uhgt6", 'u': "8ijhy7", 'i': "9okju8", 'o': "0plki9", 'p': "lo0",
		'a': "qwsz", 's': "edxzaw", 'd': "rfcxse", 'f': "tgvcdr", 'g': "yhbvft", 'h': "ujnbgy", 'j': "ikmnhu", 'k': "olmji", 'l': "kop",
		'z': "asx", 'x': "zsdc", 'c': "xdfv", 'v': "cfgb", 'b': "vghn", 'n': "bhjm", 'm': "njk"}
	keyboardDe := map[rune]string{'q': "12wa", 'w': "23esaq", 'e': "34rdsw", 'r': "45tfde", 't': "56zgfr", 'z': "67uhgt", 'u': "78ijhz", 'i': "89okju",
		'o': "90plki", 'p': "0ßüölo", 'ü': "ß+äöp", 'a': "qwsy", 's': "wedxya", 'd': "erfcxs", 'f': "rtgvcd", 'g': "tzhbvf", 'h': "zujnbg", 'j': "uikmnh",
		'k': "iolmj", 'l': "opök", 'ö': "püäl-", 'ä': "ü-ö", 'y': "asx", 'x': "sdcy", 'c': "dfvx", 'v': "fgbc", 'b': "ghnv", 'n': "hjmb", 'm': "jkn",
		'1': "2q", '2': "13wq", '3': "24ew", '4': "35re", '5': "46tr", '6': "57zt", '7': "68uz", '8': "79iu", '9': "80oi", '0': "9ßpo", 'ß': "0üp"}
	keyboardEs := map[rune]string{'q': "12wa", 'w': "23esaq", 'e': "34rdsw", 'r': "45tfde", 't': "56ygfr", 'y': "67uhgt", 'u': "78ijhy", 'i': "89okju",
		'o': "90plki", 'p': "0loñ", 'a': "qwsz", 's': "wedxza", 'd': "erfcxs", 'f': "rtgvcd", 'g': "tyhbvf", 'h': "yujnbg", 'j': "uikmnh", 'k': "iolmj",
		'l': "opkñ", 'ñ': "pl", 'z': "asx", 'x': "sdcz", 'c': "dfvx", 'v': "fgbc", 'b': "ghnv", 'n': "hjmb", 'm': "jkn", '1': "2q", '2': "13wq",
		'3': "24ew", '4': "35re", '5': "46tr", '6': "57yt", '7': "68uy", '8': "79iu", '9': "80oi", '0': "9po"}
	keyboardFr := map[rune]string{'a': "12zqé", 'z': "23eésaq", 'e': "34rdsz", 'r': "45tfde", 't': "56ygfr-", 'y': "67uhgtè-", 'u': "78ijhyè",
		'i': "89okjuç", 'o': "90plkiçà", 'p': "0àlo", 'q': "azsw", 's': "zedxwq", 'd': "erfcxs", 'f': "rtgvcd", 'g': "tzhbvf", 'h': "zujnbg",
		'j': "uikmnh", 'k': "iolmj", 'l': "opmk", 'm': "pùl", 'w': "qsx", 'x': "sdcw", 'c': "dfvx", 'v': "fgbc", 'b': "ghnv", 'n': "hjb",
		'1': "2aé", '2': "13azé", '3': "24ewé", '4': "35re", '5': "46tr", '6': "57ytè", '7': "68uyè", '8': "79iuèç", '9': "80oiçà", '0': "9àçpo"}
	keyboards = append(keyboards, keyboardEn, keyboardDe, keyboardEs, keyboardFr)
	for i, c := range domain {
		for _, keyboard := range keyboards {
			for _, char := range []rune(keyboard[c]) {
				result := fmt.Sprintf("%s%c%s", domain[:i], char, domain[i+1:])
				// remove duplicates
				count[result]++
				if count[result] < 2 {
					results = append(results, result)
				}
			}
		}
	}
	return results
}

// performs a repetition attack simulating a user pressing a key twice
func repetitionAttack(domain string) []string {
	results := []string{}
	count := make(map[string]int)
	for i, c := range domain {
		if unicode.IsLetter(c) {
			result := fmt.Sprintf("%s%c%c%s", domain[:i], domain[i], domain[i], domain[i+1:])
			// remove duplicates
			count[result]++
			if count[result] < 2 {
				results = append(results, result)
			}
		}
	}
	return results
}

// performs an omission attack removing characters across the domain name
func omissionAttack(domain string) []string {
	results := []string{}
	for i := range domain {
		results = append(results, fmt.Sprintf("%s%s", domain[:i], domain[i+1:]))
	}
	return results
}

// performs a hyphenation attack adding hyphens between characters
func hyphenationAttack(domain string) []string {
	results := []string{}
	for i := 1; i < len(domain); i++ {
		if (rune(domain[i]) != '-' || rune(domain[i]) != '.') && (rune(domain[i-1]) != '-' || rune(domain[i-1]) != '.') {
			results = append(results, fmt.Sprintf("%s-%s", domain[:i], domain[i:]))
		}
	}
	return results
}

// performs a bitsquat permutation attack
func bitsquattingAttack(domain string) []string {

	results := []string{}
	masks := []int32{1, 2, 4, 8, 16, 32, 64, 128}

	for i, c := range domain {
		for m := range masks {
			b := rune(int(c) ^ m)
			o := int(b)
			if (o >= 48 && o <= 57) || (o >= 97 && o <= 122) || o == 45 {
				results = append(results, fmt.Sprintf("%s%c%s", domain[:i], b, domain[i+1:]))
			}
		}
	}
	return results
}

// performs a homograph permutation attack
func homographAttack(domain string) []string {
	// set local variables
	glyphs := map[rune][]rune{
		'a': {'à', 'á', 'â', 'ã', 'ä', 'å', 'ɑ', 'а', 'ạ', 'ǎ', 'ă', 'ȧ', 'α', 'ａ'},
		'b': {'d', 'ʙ', 'Ь', 'ɓ', 'Б', 'ß', 'β', 'ᛒ', '\u1E05', '\u1E03', '\u1D6C'}, // 'lb', 'ib'
		'c': {'ϲ', 'с', 'ƈ', 'ċ', 'ć', 'ç', 'ｃ'},
		'd': {'b', 'ԁ', 'ժ', 'ɗ', 'đ'}, // 'cl', 'dl', 'di'
		'e': {'é', 'ê', 'ë', 'ē', 'ĕ', 'ě', 'ė', 'е', 'ẹ', 'ę', 'є', 'ϵ', 'ҽ'},
		'f': {'Ϝ', 'ƒ', 'Ғ'},
		'g': {'q', 'ɢ', 'ɡ', 'Ԍ', 'Ԍ', 'ġ', 'ğ', 'ց', 'ǵ', 'ģ'},
		'h': {'һ', 'հ', '\u13C2', 'н'}, // 'lh', 'ih'
		'i': {'1', 'l', '\u13A5', 'í', 'ï', 'ı', 'ɩ', 'ι', 'ꙇ', 'ǐ', 'ĭ'},
		'j': {'ј', 'ʝ', 'ϳ', 'ɉ'},
		'k': {'κ', 'κ'}, // 'lk', 'ik', 'lc'
		'l': {'1', 'i', 'ɫ', 'ł'},
		'm': {'n', 'ṃ', 'ᴍ', 'м', 'ɱ'}, // 'nn', 'rn', 'rr'
		'n': {'m', 'r', 'ń'},
		'o': {'0', 'Ο', 'ο', 'О', 'о', 'Օ', 'ȯ', 'ọ', 'ỏ', 'ơ', 'ó', 'ö', 'ӧ', 'ｏ'},
		'p': {'ρ', 'р', 'ƿ', 'Ϸ', 'Þ'},
		'q': {'g', 'զ', 'ԛ', 'գ', 'ʠ'},
		'r': {'ʀ', 'Г', 'ᴦ', 'ɼ', 'ɽ'},
		's': {'Ⴝ', '\u13DA', 'ʂ', 'ś', 'ѕ'},
		't': {'τ', 'т', 'ţ'},
		'u': {'μ', 'υ', 'Ս', 'ս', 'ц', 'ᴜ', 'ǔ', 'ŭ'},
		'v': {'ѵ', 'ν', '\u1E7F', '\u1E7D'},      // 'v̇'
		'w': {'ѡ', 'ա', 'ԝ'}, // 'vv'
		'x': {'х', 'ҳ', '\u1E8B'},
		'y': {'ʏ', 'γ', 'у', 'Ү', 'ý'},
		'z': {'ʐ', 'ż', 'ź', 'ʐ', 'ᴢ'},
	}
	doneCount := make(map[rune]bool)
	results := []string{}
	runes := []rune(domain)
	count := countChar(domain)

	for i, char := range runes {
		// perform attack against single character
		for _, glyph := range glyphs[char] {
			results = append(results, fmt.Sprintf("%s%c%s", string(runes[:i]), glyph, string(runes[i+1:])))
		}
		// determine if character is a duplicate
		// and if the attack has already been performed
		// against all characters at the same time
		if count[char] > 1 && doneCount[char] != true {
			doneCount[char] = true
			for _, glyph := range glyphs[char] {
				result := strings.Replace(domain, string(char), string(glyph), -1)
				results = append(results, result)
			}
		}
	}
	return results
}
