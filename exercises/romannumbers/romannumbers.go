package romannumbers

import "fmt"

func IntToRoman(number int64) string {
	struct RomanDigit {
		symbol: String,
		value: u32,
	}
	
	impl RomanDigit {
		fn new(symbol: &str, value: u32) -> RomanDigit {
			RomanDigit {
				symbol: String::from(symbol),
				value,
			}
		}
	}
	
	struct RomanNumber {
		digits: Vec<RomanDigit>,
	}
	
	impl RomanNumber {
		fn new() -> RomanNumber {
			RomanNumber { digits: Vec::new() }
		}
	}
	
	impl From<u32> for RomanNumber {
		fn from(mut num: u32) -> Self {
			let mut roman_number = RomanNumber::new();
			let roman_digits = vec![
				RomanDigit::new("M", 1000),
				RomanDigit::new("CM", 900),
				RomanDigit::new("D", 500),
				RomanDigit::new("CD", 400),
				RomanDigit::new("C", 100),
				RomanDigit::new("XC", 90),
				RomanDigit::new("L", 50),
				RomanDigit::new("XL", 40),
				RomanDigit::new("X", 10),
				RomanDigit::new("IX", 9),
				RomanDigit::new("V", 5),
				RomanDigit::new("IV", 4),
				RomanDigit::new("I", 1),
			];
	
			for roman_digit in &roman_digits {
				while num >= roman_digit.value {
					roman_number.digits.push(roman_digit.clone());
					num -= roman_digit.value;
				}
			}
	
			roman_number
		}
	}
	
	
}


func main() {
	
}
