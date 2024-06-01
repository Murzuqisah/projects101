# math-skills

## Description

This project entails a math program written in Go language; that reads inputs from a file, calculates the statistics of the population within the file and gives an integer value output displayed in the terminal.

This program consists of functions used to perform statistical population operations.

_*Statistical population refers to a set of similar items, data or events subject to an evaluation or statistical operation to atain a specific statistical value*_

They are contained within the functions directory. The functions include :

+ Average(mean) : Refers to the measure of the central tendency of a distribution characterised by that distribution.

To calculate average, calculate the sum of the values and their total occurences.

    average = sum/total number of values

+ Median : Referred to median of medians, it refers to middle value of a data set arranged in an ascending or descending order.

* The data set is sort in either an ascending or descending order before the middle digit value is determined.

In the case of even number of values:

    median = ((digits[(length/2)-1] + digits[length/2])) / 2

In the case of odd number of values:

    median = digits[length/2]


+ Standard Deviation : This function consists of two statistical operations .i.e., variation and standard deviation.

Variance is the measure of dispersion; how far a set of numbers is spread out from their average value.

    variance, σ2 = (Σ(xi - μ)^2)/N

Standard deviation is calculated by getting the square root of the variance of a given data set.

    standard deviation, σ = (σ^2)^-2


## Usage

To run the program, a command is passed on the command-line terminal

```bash
go run your-program.go data.txt
```
The file <data.txt> contains the data set required for calculation. 
The sample below shows how the data is structured within the file.

```bash
189
113
121
114
145
110
```

After getting the file from the second index of the argument, the data values are read line by line and then parsed into floating point numbers.

```bash
scanner := bufio.NewScanner(file)

	// initialize array to store the elements
	var digits []float64

	// read the file line by line
	for scanner.Scan() {
		value := scanner.Text()
		// convert the strings to integer values
		number, err := strconv.ParseFloat(value, 64)
		CheckNilError(err, "Error converting string to integer")

		digits = append(digits, number)
	}
```

The script above displays a use of standard library packages bufio and strconv to convert the parsed input into a float64 data type.

After conversion of the input's data type, defined calculations are performed on them and displayed on terminal as shown below:

```bash
Average: 132
Median: 121
Variance: 785 
Standard Deviation: 28
```

## Support

For any questions or suggestions, please reach me on my gitea @[jamos](https://learn.zone01kisumu.ke/git/jamos)


## License
The license of this project is trademarked by [MIT](LICENSE)