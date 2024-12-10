package secondassignment

func GetNumberInWord(n int) string {

	var (
		result    = ""
		units     = []string{"", "bir", "ikki", "uch", "to'rt", "besh", "olti", "yetti", "sakkiz", "to'qqiz"}
		tens      = []string{"", "o'n", "yigirma", "o'ttiz", "qirq", "ellik", "oltmish", "yetmish", "sakson", "to'qson"}
		hundreds  = []string{"", "bir yuz", "ikki yuz", "uch yuz", "to'rt yuz", "besh yuz", "olti yuz", "yetti yuz", "sakkiz yuz", "to'qqiz yuz"}
		thousands = []string{"", "bir ming", "ikki ming", "uch ming", "to'rt ming", "besh ming", "olti ming", "yetti ming", "sakkiz ming", "to'qqiz ming"}
	)

	if n >= 1000 {
		result += thousands[n/1000] + " "
		n %= 1000
	}
	if n >= 100 {
		result += hundreds[n/100] + " "
		n %= 100
	}
	if n >= 10 {
		result += tens[n/10] + " "
		n %= 10
	}
	if n > 0 {
		result += units[n]
	}

	return result
}
