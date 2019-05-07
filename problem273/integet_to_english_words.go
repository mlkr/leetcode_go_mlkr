package problem273

var lessThan21 = []string{
	"",
	"One",
	"Two",
	"Three",
	"Four",
	"Five",
	"Six",
	"Seven",
	"Eight",
	"Nine",
	"Ten",
	"Eleven",
	"Twelve",
	"Thirteen",
	"Fourteen",
	"Fifteen",
	"Sixteen",
	"Seventeen",
	"Eighteen",
	"Nineteen",
	"Twenty",
}

var ten = []string{
	"",
	"",
	"Twenty",
	"Thirty",
	"Forty",
	"Fifty",
	"Sixty",
	"Seventy",
	"Eighty",
	"Ninety",
}

var thousand = []string{
	"",
	"Thousand",
	"Million",
	"Billion",
}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

}

// 处理小于1k的数
func lessK(num int) string {
	if num == 0 {
		return ""
	}

	if num <= 20 {
		return lessThan21[num] + " "
	}

	if num < 100 {
		return ten[num/10] + " " + lessK(num%10)
	}

	return lessThan21[num/100] + " Hundred " + lessK(num%100)

}
